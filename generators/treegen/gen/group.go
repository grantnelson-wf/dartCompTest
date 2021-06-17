package gen

import (
	"fmt"
	"sort"
)

var _ Item = (*Group)(nil)

// Group is a collection of items which all live in the same path.
// The group may be a library, if desired.
type Group struct {
	index   int
	isLib   bool
	members []Node
}

// NewGroup creates a new group item.
func NewGroup(index int, library bool) *Group {
	return &Group{
		index:   index,
		isLib:   library,
		members: []Node{},
	}
}

// IsLibrary indicates if this group should be outputted as a library or not.
func (g *Group) IsLibrary() bool {
	return g.isLib
}

// String gets the file name or package name without any extension.
func (g *Group) String() string {
	return fmt.Sprint(`group`, g.index)
}

// Add adds the given nodes to this group.
// This will set the group and index for the given nodes.
func (g *Group) Add(members ...Node) {
	for _, member := range members {
		member.setGroup(g, len(g.members))
		g.members = append(g.members, member)
	}
}

// Write will write this group's files in the given base path.
// This will tell all the groups items to also write.
// This base path is the folder that should contain the `lib` folder.
func (g *Group) Write(dryRun bool, basePath, packageName string) {
	if len(g.members) <= 0 {
		panic(fmt.Errorf("may not write a group with no members"))
	}

	for _, node := range g.members {
		node.Write(dryRun, basePath, packageName)
	}

	if g.IsLibrary() {
		g.writeLibraryFile(dryRun, basePath, packageName)
	}
	g.writeExport(dryRun, basePath)
}

// ImportGroupNames gets the sorted names of all the groups which are not
// this group and therefore need to be imported for the given nodes.
func (g *Group) ImportGroupNames(nodes []Node) []string {
	imports := map[string]bool{}
	for _, node := range nodes {
		if node.Group() != g {
			imports[node.Group().String()] = true
		}
	}

	sortedImports := make([]string, len(imports))
	i := 0
	for groupName := range imports {
		sortedImports[i] = groupName
		i++
	}
	sort.Strings(sortedImports)
	return sortedImports
}

// allMemberChildren gets all the dependencies from all the members.
// This may contain duplicate nodes.
func (g *Group) allMemberChildren() []Node {
	nodes := []Node{}
	for _, item := range g.members {
		nodes = append(nodes, item.Children()...)
	}
	return nodes
}

// writeLibraryFile writes the library file for this group.
func (g *Group) writeLibraryFile(dryRun bool, basePath, packageName string) {
	out := NewItemOutput(dryRun, g, basePath, `lib`, `src`, g.String())
	defer out.Close()

	out.WriteLine(`library `, g, `;`)
	out.WriteLine()

	imports := g.ImportGroupNames(g.allMemberChildren())
	if len(imports) > 0 {
		for _, name := range imports {
			out.WriteLine(`import 'package:`, packageName, `/`, name, `.dart';`)
		}
		out.WriteLine()
	}

	for _, item := range g.members {
		out.WriteLine(`part '`, item, `.dart';`)
	}
	out.WriteLine()
}

// writeExport will write the export file for the group.
func (g *Group) writeExport(dryRun bool, basePath string) {
	out := NewItemOutput(dryRun, g, basePath, `lib`)
	defer out.Close()

	if g.IsLibrary() {
		out.WriteLine(`library `, g, `;`)
		out.WriteLine()
		out.WriteLine(`export 'src/`, g, `/`, g, `.dart';`)
	} else {
		for _, item := range g.members {
			out.WriteLine(`export 'src/`, g, `/`, item, `.dart';`)
		}
	}
}
