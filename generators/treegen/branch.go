package treegen

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

// SetGroup should only be called by a group when this branch is being added to the group
// and is used to set the group this branch belongs to.
func (n *Branch) SetGroup(group *Group, index int) {
	n.group = group
	n.index = index
}

// String gets the file name name without any extension for this branch.
func (n *Branch) String() string {
	return fmt.Sprint(n.group, "_", n.index)
}

// Write will write this branch's file in the given base path.
// This base path is the folder that should contain the `lib` folder.
func (n *Branch) Write(basePath string) {
	if len(n.children) <= 0 {
		panic(fmt.Errorf("may not write a branch with no children"))
	}

	f := CreateFile(n, basePath, `lib`, `src`)
	defer f.Close()

	if n.group.IsLibrary() {
		WriteLine(f, `part of `, n.group, `;`)
		WriteLine(f)
	} else {
		hasImports := false
		for _, item := range n.children {
			if item.Group() != n.group {
				WriteLine(f, `import 'package:`, item.Group(), `/`, item, `.dart';`)
				hasImports = true
			}
		}
		if hasImports {
			WriteLine(f)
		}

		hasImports = false
		for _, item := range n.children {
			if item.Group() == n.group {
				WriteLine(f, `import '`, item, `.dart';`)
				hasImports = true
			}
		}
		if hasImports {
			WriteLine(f)
		}
	}
	maxIndex := len(n.children) - 1

	WriteLine(f, `class `, n, `{`)
	for i, item := range n.children {
		WriteLine(f, `   `, item, ` _item`, i, `;`)
	}
	WriteLine(f)
	WriteLine(f, `   `, n, `() {`)
	for i, item := range n.children {
		WriteLine(f, `      _item`, i, ` = new `, item, `();`)
	}
	WriteLine(f, `   }`)
	WriteLine(f)
	WriteLine(f, `   int get hash {`)
	WriteLine(f, `      int hashcode = 1430287;`)
	for i := range n.children {
		WriteLine(f, `      hashcode *= 7302013 ^ _item`, i, `.hash;`)
	}
	WriteLine(f, `      return hashcode;`)
	WriteLine(f, `   }`)
	WriteLine(f)
	WriteLine(f, `   int get sum =>`)
	for i := range n.children {
		if i == maxIndex {
			WriteLine(f, `      _item`, i, `.sum +`)
		} else {
			WriteLine(f, `      _item`, i, `.sum;`)
		}
	}
	WriteLine(f)
	WriteLine(f, `   int get count =>`)
	for i := range n.children {
		if i == maxIndex {
			WriteLine(f, `      _item`, i, `.count +`)
		} else {
			WriteLine(f, `      _item`, i, `.count;`)
		}
	}
	WriteLine(f, `}`)
}
