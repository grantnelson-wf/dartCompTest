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

	var (
		deleteOld = false
		generate  = false
		update    = false
		showTree  = false
		dryRun    = false
		runPubGet = false

		basePath    = ``
		packageName = `TreeGen`

		scalar        = 1.0
		rate          = 1.0
		maxDepth      = 4
		itemsPerGroup = 10
		useLibraries  = false
		randomSeed    = false
		updateFrac    = 0.0
	)

	flag.BoolVar(&deleteOld, `del`, deleteOld,
		`Indicates that any old generated code will be deleted before adding the new generated code. `+
			`Must delete manually or with this if changing the dependency trees growth or depth.`)
	flag.BoolVar(&generate, `gen`, generate,
		`Indicates that the dependency tree should be generated and written.`)
	flag.BoolVar(&update, `update`, update,
		`Indicates that some of the leaves will be rewritten. This should be accompanied with the fraction (frac) to update. `+
			`Recommend using a random seed (rand) when updating.`)
	flag.BoolVar(&showTree, `tree`, showTree,
		`Indicates that the dependency tree diagram should be printed to the console.`)
	flag.BoolVar(&dryRun, `dry`, dryRun,
		`Indicates that the files should not be written. Instead a dry run will be run and the outputs are written to the console.`)
	flag.BoolVar(&runPubGet, `pubget`, runPubGet,
		`Indicates that after generating the code "pub get" should be automatically run.`)

	flag.StringVar(&basePath, `out`, basePath,
		`The required base path to write the generatred files to. This is the path that should have the 'lib' folder in it.`)
	flag.StringVar(&packageName, `package`, packageName,
		`The name of the package to use for this generated code.`)

	flag.Float64Var(&scalar, `scalar`, scalar,
		`The scalar to apply to the growth factor of the tree. The 'a' in 'a(b^x)'. Must be positive.`)
	flag.Float64Var(&rate, `rate`, rate,
		`The rate to apply to the growth factor of the tree. The 'b' in 'a(b^x)'. Must be one or greater.`)
	flag.IntVar(&maxDepth, `depth`, maxDepth,
		`The maximum depth to grow the dependency tree. The depth is the 'x' in 'a(b^x)'. Must be one or greater.`)
	flag.IntVar(&itemsPerGroup, `group`, itemsPerGroup,
		`The maximum number of items that can be put into a group.`)
	flag.BoolVar(&useLibraries, `lib`, useLibraries,
		`Indicates that the resulting files should be in libraries. File level dependencies will be used otherwise.`)
	flag.BoolVar(&randomSeed, `rand`, randomSeed,
		`Indicates that the values used for the leafs should be fully randomized.`)
	flag.Float64Var(&updateFrac, `frac`, updateFrac,
		`The fraction of leaves, between 0.0 and 1.0, to update during the update.`)

	flag.Parse()

	gen.Generator{
		DeleteOld: deleteOld,
		Generate:  generate,
		Update:    update,
		ShowTree:  showTree,
		DryRun:    dryRun,
		RunPubGet: runPubGet,

		BasePath:    basePath,
		PackageName: packageName,

		Scalar:        scalar,
		Rate:          rate,
		MaxDepth:      maxDepth,
		ItemsPerGroup: itemsPerGroup,
		UseLibraries:  useLibraries,
		RandomSeed:    randomSeed,
		UpdateFrac:    updateFrac,
	}.Run()

	os.Exit(0)
}
