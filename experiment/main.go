package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/grantnelson-wf/dartCompTest/experiment/trial"
)

// addTreatments defines all the treatments which should be run in this experiment
// and configures them as needed. Note that the commands are relative to the given path.
func addTreatments(trial *trial.Trial) {

	// File Dependencies is example dart code where the imports or on each file
	// such that any dependency graph has to be made between all the files.
	trial.AddTreatment().
		Name(`File_Dependencies`).
		Path(`treatments/filedeps`).
		PrepareCommand(`rm`, `-rf`, `build`).
		RunCommand(`webdev`, `build`)

	// Library Dependencies is example dart code where the imports are on a library file
	// such that
	trial.AddTreatment().
		Name(`Library_Dependencies`).
		Path(`treatments/libdeps`).
		PrepareCommand(`rm`, `-rf`, `build`).
		RunCommand(`webdev`, `build`)
}

// main is the entry point for the experiment.
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Experiment failed: %v\n", r)
			os.Exit(1)
		}
	}()

	repetitions := 10
	flag.IntVar(&repetitions, "reps", repetitions,
		`The number of times to run the experiment.`)

	resultFile := `results.txt`
	flag.StringVar(&resultFile, "out", resultFile,
		`The result file to write the duration of the application to.`)

	flag.Parse()

	trial := trial.New(repetitions, resultFile)
	addTreatments(trial)
	trial.Run()

	fmt.Println("Experiment done")
	os.Exit(0)
}
