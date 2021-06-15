package treegen

import (
	"fmt"
)

var _ Node = (*Leaf)(nil)

type Leaf struct {
	index int
	group *Group
	value int
}

func NewLeaf() *Leaf {
	n := &Leaf{
		index: -1,
		group: nil,
		value: 0,
	}
	return n
}

func (n *Leaf) Group() *Group {
	return n.group
}

func (n *Leaf) SetGroup(group *Group, index int) {
	n.group = group
	n.index = index
}

func (n *Leaf) String() string {
	return fmt.Sprint(n.group, "_", n.index)
}

func (n *Leaf) Write(basePath, packageName string) {
	f := CreateFile(basePath, n)
	defer f.Close()

	if n.group.IsLibrary() {
		WriteLine(f, `part of `, n.group)
		WriteLine(f)
	}

	WriteLine(f, `class `, n, `{`)
	WriteLine(f, `   int _value;`)
	WriteLine(f)
	WriteLine(f, `   `, n, `() {`)
	WriteLine(f, `      _value = `, n.value)
	WriteLine(f, `   }`)
	WriteLine(f)
	WriteLine(f, `   int get hex => _value;`)
	WriteLine(f)
	WriteLine(f, `   int get sum => _value;`)
	WriteLine(f)
	WriteLine(f, `   int get count => 1;`)
	WriteLine(f, `}`)
}
