package gen

import (
	"fmt"
	"os"
	"path"
)

const (
	// dirMode is the permissions mask used for any directly being created.
	dirMode = 0755
)

// Output provides a way for writting a code file.
type Output struct {
	dryRun bool
	f      *os.File
}

// NewItemOutput will create the file or output for the given item.
// The given path parts are prepended to the the given item's file name.
func NewItemOutput(dryRun bool, item Item, pathParts ...string) *Output {
	return NewOutput(dryRun, append(pathParts, item.String()+".dart")...)
}

// NewOutput will create the file or output.
func NewOutput(dryRun bool, pathParts ...string) *Output {
	filePath := path.Join(pathParts...)
	var f *os.File
	if dryRun {
		fmt.Println("+---------------------------------------------------------")
		fmt.Println("|", filePath)
		fmt.Println("+---------------------------------------------------------")
	} else {
		ensureDirectory(filePath)
		var err error
		if f, err = os.Create(filePath); err != nil {
			panic(err)
		}
	}

	return &Output{
		dryRun: dryRun,
		f:      f,
	}
}

// ensureDirectory will check that the directory for the given path exists
// and builds it recursively if it does not.
func ensureDirectory(p string) {
	dir := path.Dir(p)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		ensureDirectory(dir)
		if err := os.Mkdir(dir, dirMode); err != nil {
			panic(err)
		}
	}
}

// Close will close the open file being outputted.
func (out *Output) Close() {
	if out.f != nil {
		out.f.Close()
		out.f = nil
	}
	if out.dryRun {
		fmt.Println("+---------------------------------------------------------")
		fmt.Println()
	}
}

// WriteLine will write a single line of code to the output.
// This will not put spaces between given parts but will always end with a newline.
func (out *Output) WriteLine(parts ...interface{}) {
	line := fmt.Sprint(parts...)
	if out.dryRun {
		fmt.Println("|", line)
	} else {
		if _, err := out.f.WriteString(line + "\n"); err != nil {
			panic(err)
		}
	}
}

// DeletePath will delete the folder and any files within that folder
// at the given path.
func DeletePath(dryRun bool, pathParts ...string) {
	fullPath := path.Join(pathParts...)
	if dryRun {
		fmt.Println(`Delete`, fullPath)
	} else {
		if err := os.RemoveAll(fullPath); err != nil {
			panic(err)
		}
	}
}
