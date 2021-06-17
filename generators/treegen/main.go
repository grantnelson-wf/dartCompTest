package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/grantnelson-wf/dartCompTest/generators/treegen/gen"
)

// main is the entry point for the file dependency or library dependency generator.
// This will generate a large amount of dart code to test different build models.
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(`Tree generation failed:`, r)
			os.Exit(1)
		}
	}()

	scalar := 1.0
	flag.Float64Var(&scalar, `scalar`, scalar,
		`The scalar to apply to the growth factor of the tree. The 'a' in 'ax^e'. Must be positive.`)

	exponent := 1.0
	flag.Float64Var(&exponent, `exp`, exponent,
		`The exponent to apply to the growth factor of the tree. The 'b' in 'ax^b'. Must be one or greater.`)

	maxDepth := 4
	flag.IntVar(&maxDepth, `depth`, maxDepth,
		`The maximum depth to grow the dependency tree.`)

	itemsPerGroup := 10
	flag.IntVar(&itemsPerGroup, `group`, itemsPerGroup,
		`The maximum number of items that can be put into a group.`)

	useLibraries := false
	flag.BoolVar(&useLibraries, `lib`, useLibraries,
		`Indicates that the resulting files should be in libraries. File level dependencies will be used otherwise.`)

	dryRun := false
	flag.BoolVar(&dryRun, `dry`, dryRun,
		`Indicates that the files should not be written. Instead a dry run will be run and the outputs are written to the consoles.`)

	basePath := ``
	flag.StringVar(&basePath, `out`, basePath,
		`The required base path to write the generatred files to. This is the path that should have the 'lib' folder in it.`)

	packageName := `TreeGen`
	flag.StringVar(&packageName, `package`, packageName,
		`The name of the package to use for this generated code.`)

	flag.Parse()

	if len(basePath) <= 0 {
		panic(fmt.Errorf(`must provide the base path to write to`))
	}
	if itemsPerGroup < 1 {
		panic(fmt.Errorf(`items per group must be greater than one`))
	}
	if maxDepth < 1 {
		panic(fmt.Errorf(`maximum depth must be greater than one`))
	}
	if len(packageName) <= 0 {
		panic(fmt.Errorf(`must have a non-empty package name`))
	}

	gen.Generate(scalar, exponent, maxDepth, itemsPerGroup, dryRun, useLibraries, basePath, packageName)
	os.Exit(0)
}
