package treegen

import "fmt"

var _ Item = (*Group)(nil)

type Group struct {
	index   int
	isLib   bool
	members []Node
}

func NewGroup(index int, library bool) *Group {
	return &Group{
		index:   index,
		isLib:   library,
		members: []Node{},
	}
}

func (g *Group) IsLibrary() bool {
	return g.isLib
}

func (g *Group) String() string {
	return fmt.Sprint("group", g.index)
}

func (g *Group) Add(member Node) {
	member.SetGroup(g, len(g.members))
	g.members = append(g.members, member)
}

func (g *Group) Write(basePath, packageName string) {
	for _, node := range g.members {
		node.Write(basePath, packageName)
	}

	if g.isLib {
		g.writeLibraryFile(basePath, packageName)
	}
}

func (g *Group) writeLibraryFile(basePath, packageName string) {

}
