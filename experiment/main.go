package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Grant-Nelson/dartCompTest/experiment/trial"
)

// main is the entry point for the experiment.
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Experiment failed: %v\n", r)
			os.Exit(1)
		}
	}()

	repetitions := 100
	flag.IntVar(&repetitions, "reps", repetitions,
		`The number of times to run the experiment.`)

	resultFile := `results.txt`
	flag.StringVar(&resultFile, "out", resultFile,
		`The result file to write the duration of the application to.`)

	flag.Parse()

	trial := trial.New(repetitions, resultFile)

	trial.AddTreatment().
		Name(`File_Dependencies`).
		Path(`treatments/filedeps`).
		RunCommand(`webdev`, `build`)

	// trial.AddTreatment().
	// 	Name(`Library_Dependencies`).
	// 	Path(`treatments/libdeps`).
	// 	RunCommand(`webdev`, `build`)

	trial.Run()

	fmt.Println("Experiment done")
	os.Exit(0)
}
