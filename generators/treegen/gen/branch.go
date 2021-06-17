package gen

import "fmt"

var _ Node = (*Branch)(nil)

// Branch is a node in the dependency tree.
// It represents a file which depends on a collection of children nodes.
type Branch struct {
	index    int
	group    *Group
	children []Node
}

// NewBranch creates a new branch node.
func NewBranch() *Branch {
	n := &Branch{
		index:    -1,
		group:    nil,
		children: nil,
	}
	return n
}

// Add will add the given nodes as dependencies to this branch.
func (n *Branch) Add(nodes ...Node) {
	n.children = append(n.children, nodes...)
}

// Group gets the collection of items this branch belongs to.
func (n *Branch) Group() *Group {
	return n.group
}

// setGroup should only be called by a group when this branch is being added to the group
// and is used to set the group this branch belongs to.
func (n *Branch) setGroup(group *Group, index int) {
	n.group = group
	n.index = index
}

// PrintTree prints this node and any children to the console as a tree.
func (n *Branch) PrintTree(indent string, last bool) {
	count := len(n.children)
	if last {
		fmt.Println(indent+` '--branch`, n)
		for i, item := range n.children {
			item.PrintTree(indent+`   `, i+1 == count)
		}
	} else {
		fmt.Println(indent+` |--branch`, n)
		for i, item := range n.children {
			item.PrintTree(indent+` | `, i+1 == count)
		}
	}
}

// String gets the file name name without any extension for this branch.
func (n *Branch) String() string {
	return fmt.Sprint(n.group, "_", n.index)
}

// Write will write this branch's file in the given base path.
// This base path is the folder that should contain the `lib` folder.
func (n *Branch) Write(dryRun bool, basePath, packageName string) {
	if len(n.children) <= 0 {
		panic(fmt.Errorf("may not write a branch with no children"))
	}

	out := NewItemOutput(dryRun, n, basePath, `lib`, `src`, n.group.String())
	defer out.Close()

	if n.group.IsLibrary() {
		out.WriteLine(`part of `, n.group, `;`)
		out.WriteLine()
	} else {
		hasImports := false
		for _, item := range n.children {
			if item.Group() != n.group {
				out.WriteLine(`import 'package:`, packageName, `/`, item, `.dart';`)
				hasImports = true
			}
		}
		if hasImports {
			out.WriteLine()
		}

		hasImports = false
		for _, item := range n.children {
			if item.Group() == n.group {
				out.WriteLine(`import '`, item, `.dart';`)
				hasImports = true
			}
		}
		if hasImports {
			out.WriteLine()
		}
	}
	maxIndex := len(n.children) - 1

	out.WriteLine(`class `, n, `{`)
	for i, item := range n.children {
		out.WriteLine(`   `, item, ` _item`, i, `;`)
	}
	out.WriteLine()
	out.WriteLine(`   `, n, `() {`)
	for i, item := range n.children {
		out.WriteLine(`      _item`, i, ` = new `, item, `();`)
	}
	out.WriteLine(`   }`)
	out.WriteLine()
	out.WriteLine(`   int get hash {`)
	out.WriteLine(`      int hashcode = 1430287;`)
	for i := range n.children {
		out.WriteLine(`      hashcode *= 7302013 ^ _item`, i, `.hash;`)
	}
	out.WriteLine(`      return hashcode;`)
	out.WriteLine(`   }`)
	out.WriteLine()
	out.WriteLine(`   int get sum =>`)
	for i := range n.children {
		if i == maxIndex {
			out.WriteLine(`      _item`, i, `.sum;`)
		} else {
			out.WriteLine(`      _item`, i, `.sum +`)
		}
	}
	out.WriteLine()
	out.WriteLine(`   int get count =>`)
	for i := range n.children {
		if i == maxIndex {
			out.WriteLine(`      _item`, i, `.count;`)
		} else {
			out.WriteLine(`      _item`, i, `.count +`)
		}
	}
	out.WriteLine(`}`)
}
