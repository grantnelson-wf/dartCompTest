package treegen

import (
	"fmt"
	"os"
	"path"
)

type (
	Item interface {
		String() string
		Write(basePath, packageName string)
	}

	Node interface {
		Item
		Group() *Group
		SetGroup(group *Group, index int)
	}
)

func CreateFile(basePath string, item Item) *os.File {
	f, err := os.Create(path.Join(basePath, item.String()+".dart"))
	if err != nil {
		panic(err)
	}
	return f
}

func WriteLine(f *os.File, parts ...interface{}) {
	_, err := f.WriteString(fmt.Sprint(parts...) + "\n")
	if err != nil {
		panic(err)
	}
}
