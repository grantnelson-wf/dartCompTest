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

// SetGroup should only be called by a group when this leaf is being added to the group
// and is used to set the group this leaf belongs to.
func (n *Leaf) SetGroup(group *Group, index int) {
	n.group = group
	n.index = index
}

// String gets the file name name without any extension for this leaf.
func (n *Leaf) String() string {
	return fmt.Sprint(n.group, "_", n.index)
}

// Write will write this leaf's file in the given base path.
// This base path is the folder that should contain the `lib` folder.
func (n *Leaf) Write(basePath string) {
	f := CreateFile(n, basePath, `lib`, `src`)
	defer f.Close()

	if n.group.IsLibrary() {
		WriteLine(f, `part of `, n.group, `;`)
		WriteLine(f)
	}

	WriteLine(f, `class `, n, `{`)
	WriteLine(f, `   int _value;`)
	WriteLine(f)
	WriteLine(f, `   `, n, `() {`)
	WriteLine(f, `      _value = `, n.value)
	WriteLine(f, `   }`)
	WriteLine(f)
	WriteLine(f, `   int get hash => _value;`)
	WriteLine(f)
	WriteLine(f, `   int get sum => _value;`)
	WriteLine(f)
	WriteLine(f, `   int get count => 1;`)
	WriteLine(f, `}`)
}
