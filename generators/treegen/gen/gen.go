package gen

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
)

const (
	// maxLeafValue is the maximum integer value that the leaf can handle
	// since this will be run on JS via dart.
	maxLeafValue = 2 ^ 53 - 1
)

// Generator is the structure to fill out to configure the generation.
type Generator struct {
	Scalar        float64
	Exponent      float64
	MaxDepth      int
	ItemsPerGroup int
	DryRun        bool
	RunPubGet     bool
	ShowTree      bool
	UseLibraries  bool
	BasePath      string
	PackageName   string
	RandomSeed    bool
}

// Generate will generate a dependency tree of files for a large treatment example code.
func (g *Generator) Generate() {
	allNodes := g.generateDependencyTree()
	groups := g.groupNodes(allNodes)
	for _, group := range groups {
		group.Write(g.DryRun, g.BasePath, g.PackageName)
	}

	root := allNodes[0]
	g.writeYaml()
	g.writeWeb(root)
	g.writeHtml(root)
	if !g.DryRun && g.RunPubGet {
		g.runPubGet()
	}

	if g.ShowTree {
		fmt.Println(`dependencies`)
		root.PrintTree(``, true)
	}

	fmt.Println(len(allNodes), `nodes,`, len(groups), `groups generated`)
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

// createBreadth creates all the nodes for the entire breadth of the tree at the given depth
// using the given growth factors to determine the size of this depth's breadth.
// If the given depth is the maximum depth then the returned nodes will be all leaves,
// otherwise they will be branches.
func (g *Generator) createBreadth(depth int) []Node {
	count := bound(int(g.Scalar * math.Pow(float64(depth), g.Exponent)))
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
// This returns all the nodes in the whole tree.
func (g *Generator) generateDependencyTree() []Node {
	root := NewBranch()
	allNodes := []Node{root}
	prevNodes := []Node{root}
	for depth := 1; depth <= g.MaxDepth; depth++ {
		curNodes := g.createBreadth(depth)
		g.assignParents(prevNodes, curNodes)
		prevNodes = curNodes
		allNodes = append(allNodes, curNodes...)
	}
	return allNodes
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
