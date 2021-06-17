package gen

import (
	"fmt"
	"os"
	"path"
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
	path := path.Join(pathParts...)

	var f *os.File
	if dryRun {
		fmt.Println("+---------------------------------------------------------")
		fmt.Println("|", path)
		fmt.Println("+---------------------------------------------------------")
	} else {
		var err error
		if f, err = os.Create(path); err != nil {
			panic(err)
		}
	}

	return &Output{
		dryRun: dryRun,
		f:      f,
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
