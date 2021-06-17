package gen

import (
	"math"
)

// Generate will generate a dependency tree of files for a large treatment example code.
func Generate(scalar, exponent float64, maxDepth, itemsPerGroup int, dryRun, useLibraries bool, basePath string) {
	allNodes := generateDependencyTree(scalar, exponent, maxDepth)
	groups := groupNodes(allNodes, itemsPerGroup, useLibraries)
	for _, group := range groups {
		group.Write(dryRun, basePath)
	}

	root := allNodes[0]
	// root.PrintTree(``, true)
	writeWeb(dryRun, basePath, root)
	writeHtml(dryRun, basePath, root)
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
			nodes[i] = NewLeaf(i)
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
	if groupCount*itemsPerGroup < nodeCount {
		groupCount++
	}
	groups := make([]*Group, groupCount)
	start := 0
	for i := range groups {
		stop := int(float64(i+1) * float64(nodeCount) / float64(groupCount))
		group := NewGroup(i, useLibraries)
		group.Add(allNodes[start:stop]...)
		groups[i] = group
		start = stop
	}
	return groups
}

func writeWeb(dryRun bool, basePath string, root Node) {
	out := NewOutput(dryRun, basePath, `web`, `main.dart`)
	defer out.Close()

	// TODO: Implement
}

func writeHtml(dryRun bool, basePath string, root Node) {
	out := NewOutput(dryRun, basePath, `web`, `index.html`)
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
