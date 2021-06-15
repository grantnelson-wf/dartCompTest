package treegen

import "fmt"

var _ Node = (*Branch)(nil)

type Branch struct {
	index    int
	group    *Group
	children []Node
}

func NewBranch() *Branch {
	n := &Branch{
		index:    -1,
		group:    nil,
		children: nil,
	}
	return n
}

func (n *Branch) Add(nodes ...Node) {
	n.children = append(n.children, nodes...)
}

func (n *Branch) Group() *Group {
	return n.group
}

func (n *Branch) SetGroup(group *Group, index int) {
	n.group = group
	n.index = index
}

func (n *Branch) String() string {
	return fmt.Sprint(n.group, "_", n.index)
}

func (n *Branch) Write(basePath, packageName string) {
	f := CreateFile(basePath, n)
	defer f.Close()

	if n.group.IsLibrary() {
		WriteLine(f, `part of `, n.group)
		WriteLine(f)
	} else {
		for _, item := range n.children {
			if item.Group() != n.group {
				// TODO: FILL OUT
			}
		}
		for _, item := range n.children {
			if item.Group() == n.group {
				// TODO: FILL OUT
			}
		}
	}

	// TODO: FINISH
	WriteLine(f, `class `, n, `{`)
	//for i, item := n.children {
	//	WriteLine(f, `   final int value;`)
	//}
	WriteLine(f)
	WriteLine(f, `   `, n, `() {`)
	//WriteLine(f, `      value = `, n.value)
	WriteLine(f, `   }`)
	WriteLine(f)
	WriteLine(f, `   int get hex => value;`)
	WriteLine(f)
	WriteLine(f, `   int get sum => value;`)
	WriteLine(f)
	WriteLine(f, `   int get count => 1;`)
	WriteLine(f, `}`)
}
