package main

import (
	"flag"
	"os"

	"./trial"
)

// main is the entry point for the experiment.
func main() {
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

	trial.AddTreatment().
		Name(`Library_Dependencies`).
		Path(`treatments/libdeps`).
		RunCommand(`webdev`, `build`)

	trial.Run()

	os.Exit(0)
}
