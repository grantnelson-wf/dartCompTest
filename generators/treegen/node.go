package treegen

import (
	"fmt"
	"os"
	"path"
)

type (
	// Item is the interface for anything which can produce a file.
	Item interface {

		// String gets the file name or package name without any extension,
		// e.g. `blockchain`, not `blockchain.dart`
		String() string

		// Write will write the file for this item in the given base path.
		// This base path is the folder that should contain the `lib` folder.
		Write(basePath string)
	}

	// Node is the interface for any part of the dependency tree being built.
	Node interface {
		Item

		// Group is the collection of items which share a folder or library.
		Group() *Group

		// SetGroup should only be called by a group when an item is being added to the group
		// and is used to set the group an item belongs to.
		SetGroup(group *Group, index int)
	}
)

// CreateFile will create the file for the given item.
// The given path parts are prepended to the the given item's file name.
func CreateFile(item Item, pathParts ...string) *os.File {
	pathParts = append(pathParts, item.String()+".dart")
	f, err := os.Create(path.Join(pathParts...))
	if err != nil {
		panic(err)
	}
	return f
}

// WriteLine will write a single line of code to the given file.
// This will not put spaces between given parts but will always end with a newline.
func WriteLine(f *os.File, parts ...interface{}) {
	_, err := f.WriteString(fmt.Sprint(parts...) + "\n")
	if err != nil {
		panic(err)
	}
}
