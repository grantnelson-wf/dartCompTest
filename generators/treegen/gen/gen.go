package gen

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	// maxLeafValue is the maximum integer value that the leaf can handle
	// since this will be run on JS via dart.
	maxLeafValue = 2 ^ 53 - 1
)

// Generator is the structure to fill out to configure the generation.
type Generator struct {
	DeleteOld bool
	Generate  bool
	Update    bool
	ShowTree  bool
	DryRun    bool
	RunPubGet bool

	BasePath    string
	PackageName string

	Scalar        float64
	Rate          float64
	MaxDepth      int
	ItemsPerGroup int
	UseLibraries  bool
	RandomSeed    bool
	UpdateFrac    float64
	UpdateConst   int
}

// Run will run the current configuration and potentially
// generate a dependency tree of files for a large treatment example code,
// update a selection of the generated leaf nodes,
// delete the old generated files and folders.
func (g Generator) Run() {
	g.validateConfig()

	if g.DeleteOld {
		g.deleteOldGen()
	}

	if !(g.Generate || g.Update || g.ShowTree) {
		// No reason to generate dependency tree so leave.
		return
	}

	root, allNodes, leaves := g.generateDependencyTree()
	groups := g.groupNodes(allNodes)

	if g.Generate {
		for _, group := range groups {
			group.Write(g.DryRun, g.BasePath, g.PackageName)
		}

		g.writeYaml()
		g.writeWeb(root)
		g.writeHtml(root)
		g.writeManifest(allNodes, groups)

		if !g.DryRun && g.RunPubGet {
			g.runPubGet()
		}
	}

	if g.Update {
		leafCount := len(leaves)
		rand.Shuffle(leafCount, func(i, j int) {
			leaves[i], leaves[j] = leaves[j], leaves[i]
		})
		amount := limit(bound(g.UpdateConst+int(g.UpdateFrac*float64(leafCount))), leafCount)
		update := leaves[0:amount]
		for _, leaf := range update {
			leaf.Write(g.DryRun, g.BasePath, g.PackageName)
		}
	}

	if g.ShowTree {
		fmt.Println(`dependencies`)
		root.PrintTree(``, true)
	}
}

// validateConfig checks that the generators configuration is valid before
// performing any actions. This will panic on any invalid configuration.
func (g *Generator) validateConfig() {
	if g.Scalar <= 0.0 {
		panic(fmt.Errorf(`scalar must be positive`))
	}
	if g.Rate < 1.0 {
		panic(fmt.Errorf(`rate must be one or greater`))
	}
	if len(g.BasePath) <= 0 {
		panic(fmt.Errorf(`must provide the base path to write to`))
	}
	if g.ItemsPerGroup < 1 {
		panic(fmt.Errorf(`items per group must be greater than one`))
	}
	if g.MaxDepth < 1 {
		panic(fmt.Errorf(`maximum depth must be greater than one`))
	}
	if len(g.PackageName) <= 0 {
		panic(fmt.Errorf(`must have a non-empty package name`))
	}
	if g.DeleteOld && !g.Generate && g.Update {
		panic(fmt.Errorf(`can not delete the old while also only updating`))
	}
	if g.UpdateFrac < 0.0 || g.UpdateFrac > 1.0 {
		panic(fmt.Errorf(`the update fraction must be between 0.0 and 1.0`))
	}
	if g.UpdateConst < 0 {
		panic(fmt.Errorf(`the update constant must be greater than zero`))
	}
}

// bound will return the given value or 1, whichever is greater.
func bound(val int) int {
	if val < 1 {
		return 1
	}
	return val
}

// limit will return the given value or the given maximum, whichever is less.
func limit(val, max int) int {
	if val > max {
		return max
	}
	return val
}

// deleteOldGen will remove all the old generated code.
// This will not remove the pubspec and manifest nor the base path under
// the assumption that the path and package will be regenerated and those
// will simply be overwritten.
func (g *Generator) deleteOldGen() {
	DeletePath(g.DryRun, g.BasePath, `lib`)
	DeletePath(g.DryRun, g.BasePath, `web`)
	DeletePath(g.DryRun, g.BasePath, `build`)
}

// createBreadth creates all the nodes for the entire breadth of the tree at the given depth
// using the given growth factors to determine the size of this depth's breadth.
// If the given depth is the maximum depth then the returned nodes will be all leaves,
// otherwise they will be branches.
func (g *Generator) createBreadth(depth int) []Node {
	count := bound(int(g.Scalar * math.Pow(g.Rate, float64(depth))))
	nodes := make([]Node, count)
	if depth == g.MaxDepth {
		seed := int64(0)
		if g.RandomSeed {
			seed = rand.Int63()
		}
		r := rand.New(rand.NewSource(seed))
		for i := range nodes {
			nodes[i] = NewLeaf(r.Intn(maxLeafValue))
		}
	} else {
		for i := range nodes {
			nodes[i] = NewBranch()
		}
	}
	return nodes
}

// assignParents will distribute the children, curNodes, as evenly as possible among the given parents,
// prevNodes. The parents are assumed to be branch nodes, otherwise this shouldn't be called.
func (g *Generator) assignParents(prevNodes, curNodes []Node) {
	curCount, prevCount := float64(len(curNodes)), float64(len(prevNodes))
	start := 0
	for i, parent := range prevNodes {
		stop := int(float64(i+1) * curCount / prevCount)
		parent.(*Branch).Add(curNodes[start:stop]...)
		start = stop
	}
}

// generateDependencyTree will generate the whole dependency tree.
// This returns the root, all the nodes in the whole tree, and the leaf nodes.
func (g *Generator) generateDependencyTree() (Node, []Node, []Node) {
	root := NewBranch()
	allNodes := []Node{root}
	prevNodes := []Node{root}
	for depth := 1; depth <= g.MaxDepth; depth++ {
		curNodes := g.createBreadth(depth)
		g.assignParents(prevNodes, curNodes)
		prevNodes = curNodes
		allNodes = append(allNodes, curNodes...)
	}
	return root, allNodes, prevNodes
}

// groupNodes will group the dependency tree's nodes into groups which can be written as folders or libraries.
func (g *Generator) groupNodes(allNodes []Node) []*Group {
	nodeCount := len(allNodes)
	groupCount := bound(nodeCount / g.ItemsPerGroup)
	if groupCount*g.ItemsPerGroup < nodeCount {
		groupCount++
	}
	groups := make([]*Group, groupCount)
	start := 0
	for i := range groups {
		stop := int(float64(i+1) * float64(nodeCount) / float64(groupCount))
		group := NewGroup(i, g.UseLibraries)
		group.Add(allNodes[start:stop]...)
		groups[i] = group
		start = stop
	}
	return groups
}

// writeYaml will write the pubspec.yaml file for this package.
func (g *Generator) writeYaml() {
	out := NewOutput(g.DryRun, g.BasePath, `pubspec.yaml`)
	defer out.Close()

	out.WriteLine(`name: `, g.PackageName)
	out.WriteLine(`version: 0.1.0`)
	out.WriteLine(`description: An experimental treatment which has been generated`)
	out.WriteLine()
	out.WriteLine(`environment:`)
	out.WriteLine(`  sdk: '>=2.7.0 <3.0.0'`)
	out.WriteLine()
	out.WriteLine(`dependencies:`)
	out.WriteLine(`  validators: ^2.0.1`)
	out.WriteLine()
	out.WriteLine(`dev_dependencies:`)
	out.WriteLine(`  dart_dev: ^3.6.1`)
	out.WriteLine(`  build_runner: ^1.10.0`)
	out.WriteLine(`  build_test: ^1.2.1`)
	out.WriteLine(`  build_web_compilers: ^2.9.0`)
}

// writeWeb will write the main dart entry point for this package.
func (g *Generator) writeWeb(root Node) {
	out := NewOutput(g.DryRun, g.BasePath, `web`, `main.dart`)
	defer out.Close()

	out.WriteLine(`import 'dart:html';`)
	out.WriteLine()
	out.WriteLine(`import 'package:`, g.PackageName, `/`, root.Group(), `.dart';`)
	out.WriteLine()
	out.WriteLine(`void main() {`)
	out.WriteLine(`  document.title = 'TreeGen - `, g.PackageName, `';`)
	out.WriteLine()
	out.WriteLine(`  final root = `, root, `();`)
	out.WriteLine(`  final div1 = DivElement()..innerText = 'Hash = ${root.hash}';`)
	out.WriteLine(`  final div2 = DivElement()..innerText = 'Sum = ${root.sum}';`)
	out.WriteLine(`  final div3 = DivElement()..innerText = 'Count = ${root.count}';`)
	out.WriteLine()
	out.WriteLine(`  document.body`)
	out.WriteLine(`    ..append(div1)`)
	out.WriteLine(`    ..append(div2)`)
	out.WriteLine(`    ..append(div3);`)
	out.WriteLine(`}`)
}

// writeHtml will write the html file for this package.
func (g *Generator) writeHtml(root Node) {
	out := NewOutput(g.DryRun, g.BasePath, `web`, `index.html`)
	defer out.Close()

	out.WriteLine(`<!DOCTYPE html>`)
	out.WriteLine(`<html>`)
	out.WriteLine(`  <head>`)
	out.WriteLine(`    <title>Loading Dart...</title>`)
	out.WriteLine(`  </head>`)
	out.WriteLine(`  <body>`)
	out.WriteLine(`    <script defer src="main.dart.js"></script>`)
	out.WriteLine(`  </body>`)
	out.WriteLine(`</html>`)
}

// writeManifest will write the configuration and
func (g *Generator) writeManifest(allNodes []Node, groups []*Group) {
	allCount := len(allNodes)
	groupCount := len(groups)
	branchCount := 0
	childSum := 0
	for _, node := range allNodes {
		if b, ok := node.(*Branch); ok {
			branchCount++
			childSum += len(b.Children())
		}
	}

	totalFiles := allCount + groupCount + 3
	if g.UseLibraries {
		totalFiles += groupCount
	}

	out := NewOutput(g.DryRun, g.BasePath, `manifest.txt`)
	defer out.Close()

	out.WriteLine(`This package was generated with treegen`)
	out.WriteLine(`command: go run ./generators/treegen/main.go `, strings.Join(os.Args[1:], ` `))
	out.WriteLine(`time:    `, time.Now().Format(time.UnixDate))
	out.WriteLine()

	out.WriteLine(`deleteOld: `, g.DeleteOld)
	out.WriteLine(`generate:  `, g.Generate)
	out.WriteLine(`update:    `, g.Update)
	out.WriteLine(`showTree:  `, g.ShowTree)
	out.WriteLine(`dryRun:    `, g.DryRun)
	out.WriteLine(`runPubGet: `, g.RunPubGet)
	out.WriteLine()

	out.WriteLine(`basePath:    `, g.BasePath)
	out.WriteLine(`packageName: `, g.PackageName)
	out.WriteLine()

	out.WriteLine(`scalar:        `, fmt.Sprintf("%.4f", g.Scalar))
	out.WriteLine(`rate:          `, fmt.Sprintf("%.4f", g.Rate))
	out.WriteLine(`maxDepth:      `, g.MaxDepth)
	out.WriteLine(`itemsPerGroup: `, g.ItemsPerGroup)
	out.WriteLine(`useLibraries:  `, g.UseLibraries)
	out.WriteLine(`randomSeed:    `, g.RandomSeed)
	out.WriteLine(`updateFrac:    `, fmt.Sprintf("%.4f", g.UpdateFrac))
	out.WriteLine()

	out.WriteLine(`total files:      `, totalFiles)
	out.WriteLine(`total nodes:      `, allCount)
	out.WriteLine(`branch nodes:     `, branchCount)
	out.WriteLine(`leaf nodes:       `, allCount-branchCount)
	out.WriteLine(`groups:           `, groupCount)
	out.WriteLine(`nodes per group:  `, fmt.Sprintf("%.4f", float64(allCount)/float64(groupCount)))
	out.WriteLine(`child per branch: `, fmt.Sprintf("%.4f", float64(childSum)/float64(branchCount)))
}

// runPubGet will run `pub get` on the create yaml to prepare the generated code to be built.
func (g *Generator) runPubGet() {
	cmd := exec.Command(`pub`, `get`)
	cmd.Dir = g.BasePath
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
