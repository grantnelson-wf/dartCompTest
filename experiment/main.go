package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/grantnelson-wf/dartCompTest/experiment/trial"
)

// addTreatments_FileVsLib_Dart2js defines the treatments which should be run in an experiment.
// This experiment compares using file dependencies against library dependencies. Both will compile with Dart2js.
// Note that the commands are relative to the given path.
func addTreatments_FileVsLib_Dart2js(trial *trial.Trial) {
	trial.AddTreatment().
		Name(`File_Dep`).
		Path(`treatments/filedeps1`).
		PrepareCommand(`rm`, `-rf`, `build`).
		RunCommand(`webdev`, `build`)

	trial.AddTreatment().
		Name(`Library_Dep`).
		Path(`treatments/libdeps1`).
		PrepareCommand(`rm`, `-rf`, `build`).
		RunCommand(`webdev`, `build`)
}

// addTreatments_Dart2jsVsDartDevC_File defines the treatments which should be run in an experiment.
// This experiment compares building the file dependencies example with Dart2js and DartDevC.
// This can not be combined with addTreatments_FileVsLib_Dart2js without updating ANOVA model
// to have additional grouping for the type of build.
// Note that the commands are relative to the given path.
func addTreatments_Dart2jsVsDartDevC_File(trial *trial.Trial) {
	trial.AddTreatment().
		Name(`File_Dep_Dart2js`).
		Path(`treatments/filedeps1`).
		PrepareCommand(`rm`, `-rf`, `build`).
		RunCommand(`webdev`, `build`)

	trial.AddTreatment().
		Name(`File_Dep_DartDevC`).
		Path(`treatments/filedeps1`).
		PrepareCommand(`rm`, `-rf`, `build`).
		RunCommand(`webdev`, `build`, `--no-release`)
}

// addTreatments_FileVsLib_Generated_Dart2js defines the treatments which should be run in an experiment.
// This experiment compares using the generated file dependencies against the generated library dependencies.
// Both will compile with Dart2js. Note that the commands are relative to the given path.
func addTreatments_FileVsLib_Generated_Dart2js(trial *trial.Trial) {
	trial.AddTreatment().
		Name(`File_Dep`).
		Path(`treatments/filedeps_gen`).
		PrepareCommand(`rm`, `-rf`, `build`).
		RunCommand(`webdev`, `build`)

	trial.AddTreatment().
		Name(`Library_Dep`).
		Path(`treatments/libdeps_gen`).
		PrepareCommand(`rm`, `-rf`, `build`).
		RunCommand(`webdev`, `build`)

	trial.SetTimeouts(30*time.Second, 5*time.Minute)
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

	// Pick Only One!
	//------------------
	//addTreatments_FileVsLib_Dart2js(trial)
	//addTreatments_Dart2jsVsDartDevC_File(trial)
	addTreatments_FileVsLib_Generated_Dart2js(trial)
	//------------------

	trial.Run()

	fmt.Println("Experiment done")
	os.Exit(0)
}
