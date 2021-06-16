package treegen

import (
	"flag"
	"fmt"
	"math"
	"os"
)

// main is the entry point for the file dependency or library dependency generator.
// This will generate a large amount of dart code to test different build models.
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Tree generation failed: %v\n", r)
			os.Exit(1)
		}
	}()

	scalar := 1.0
	flag.Float64Var(&scalar, "scalar", scalar,
		`The scalar to apply to the growth factor of the tree. The 'a' in 'ax^e'. Must be positive.`)

	exponent := 1.0
	flag.Float64Var(&exponent, "exp", exponent,
		`The exponent to apply to the growth factor of the tree. The 'b' in 'ax^b'. Must be one or greater.`)

	maxDepth := 4
	flag.IntVar(&maxDepth, "depth", maxDepth,
		`The maximum depth to grow the dependency tree.`)

	itemsPerGroup := 10
	flag.IntVar(&itemsPerGroup, "group", itemsPerGroup,
		`The maximum number of items that can be put into a group.`)

	useLibraries := false
	flag.BoolVar(&useLibraries, "lib", useLibraries,
		`Indicates that the resulting files should be in libraries. File level dependencies will be used otherwise.`)

	basePath := ``
	flag.StringVar(&basePath, "out", basePath,
		`The required base path to write the generatred files to. This is the path that should have the 'lib' folder in it.`)

	flag.Parse()

	Generate(scalar, exponent, maxDepth, itemsPerGroup, useLibraries, basePath)
	os.Exit(0)
}

// Generate will generate a dependency tree of files for a large treatment example code.
func Generate(scalar, exponent float64, maxDepth, itemsPerGroup int, useLibraries bool, basePath string) {
	allNodes := generateDependencyTree(scalar, exponent, maxDepth)
	groups := groupNodes(allNodes, itemsPerGroup, useLibraries)
	for _, group := range groups {
		group.Write(basePath)
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

// createBreadth creates all the nodes for the entire breadth of the tree at the given depth
// using the given growth factors to determine the size of this depth's breadth.
// If the given depth is the maximum depth then the returned nodes will be all leaves,
// otherwise they will be branches.
func createBreadth(scalar, exponent float64, depth, maxDepth int) []Node {
	count := bound(int(scalar * math.Pow(float64(depth), exponent)))
	nodes := make([]Node, count)
	if depth == maxDepth {
		for i := range nodes {
			nodes[i] = NewLeaf()
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
func assignParents(prevNodes, curNodes []Node) {
	count := len(curNodes)
	childrenPerBranch := bound(count / len(prevNodes))
	for i, start, stop := 0, 0, childrenPerBranch; start < count; i, start, stop = i+1, stop, limit(stop+childrenPerBranch, count) {
		prevNodes[i].(*Branch).Add(curNodes[start:stop]...)
	}
}

// generateDependencyTree will generate the whole dependency tree.
// This returns all the nodes in the whole tree.
func generateDependencyTree(scalar, exponent float64, maxDepth int) []Node {
	root := NewBranch()
	allNodes := []Node{root}
	prevNodes := []Node{root}
	for depth := 1; depth <= maxDepth; depth++ {
		curNodes := createBreadth(scalar, exponent, depth, maxDepth)
		assignParents(prevNodes, curNodes)
		prevNodes = curNodes
		allNodes = append(allNodes, curNodes...)
	}
	return allNodes
}

// groupNodes will group the dependency tree's nodes into groups which can be written as folders or libraries.
func groupNodes(allNodes []Node, itemsPerGroup int, useLibraries bool) []*Group {
	nodeCount := len(allNodes)
	groupCount := bound(nodeCount / itemsPerGroup)
	groups := make([]*Group, groupCount)
	start, stop := 0, limit(itemsPerGroup, nodeCount)
	for i := range groups {
		group := NewGroup(i, useLibraries)
		group.Add(allNodes[start:stop]...)
		groups[i] = group
		start, stop = stop, limit(stop+itemsPerGroup, nodeCount)
	}
	return groups
}
