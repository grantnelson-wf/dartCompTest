package gen

import (
	"fmt"
)

var _ Node = (*Leaf)(nil)

// Leaf is a node at the end of a dependency tree.
// The leaf node does not depend on anyother item.
type Leaf struct {
	index int
	group *Group

	// The value used for the leaf could be used to rewrite the leafs
	// with random values to force a rebuild of a part of the whole dependency tree.
	value int
}

// NewLeaf creates a new leaf node.
func NewLeaf(value int) *Leaf {
	n := &Leaf{
		index: -1,
		group: nil,
		value: value,
	}
	return n
}

// Group gets the collection of items this leaf belongs to.
func (n *Leaf) Group() *Group {
	return n.group
}

// setGroup should only be called by a group when this leaf is being added to the group
// and is used to set the group this leaf belongs to.
func (n *Leaf) setGroup(group *Group, index int) {
	n.group = group
	n.index = index
}

// PrintTree prints this node and any children to the console as a tree.
func (n *Leaf) PrintTree(indent string, last bool) {
	if last {
		fmt.Println(indent+` '--leaf`, n, `:`, n.value)
	} else {
		fmt.Println(indent+` |--leaf`, n, `:`, n.value)
	}
}

// String gets the file name name without any extension for this leaf.
func (n *Leaf) String() string {
	return fmt.Sprint(n.group, "_", n.index)
}

// Write will write this leaf's file in the given base path.
// This base path is the folder that should contain the `lib` folder.
func (n *Leaf) Write(dryRun bool, basePath string) {
	out := NewItemOutput(dryRun, n, basePath, `lib`, `src`)
	defer out.Close()

	if n.group.IsLibrary() {
		out.WriteLine(`part of `, n.group, `;`)
		out.WriteLine()
	}

	out.WriteLine(`class `, n, `{`)
	out.WriteLine(`   int _value;`)
	out.WriteLine()
	out.WriteLine(`   `, n, `() {`)
	out.WriteLine(`      _value = `, n.value)
	out.WriteLine(`   }`)
	out.WriteLine()
	out.WriteLine(`   int get hash => _value;`)
	out.WriteLine()
	out.WriteLine(`   int get sum => _value;`)
	out.WriteLine()
	out.WriteLine(`   int get count => 1;`)
	out.WriteLine(`}`)
}
