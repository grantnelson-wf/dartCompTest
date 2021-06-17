# Dart Compilation Experiment

Experimental test of dart compiling different type of code bases.

## Setting up the Experiment

> `./setup.sh`

Generating treatments with something like

> `go run ./generators/treegen/main.go -out ./treatments/filedeps_gen -pubget -exp 2 -depth 10 -scalar 100 -group 100`
>
> `go run ./generators/treegen/main.go -out ./treatments/libdeps_gen -lib -pubget -exp 2 -depth 10 -scalar 100 -group 100`

## Running Experiment

Configure the trial that should be run in [`./experiment/main.go`](./experiment/main.go),

Then run

> `go run ./experiment/main.go`

## Running Analysis on Results

You may need to add some python packages, use pip as needed.

This assumes the results are in `./results.txt`

> `cd analysis`
>
> `python3 anova.py`

## Running Individual Treatments

Each one may have it's own unique method for running but in general it will look like.

> `cd treatments/filedeps`
>
> `pub get`
>
> `webdev serve --no-injected-client`

## On Treatment Timeout or Failure

If a treatment times out it is possible its because the dart compiler got stuck (it does that)
use the following command to see if the dart compiler is still alive and needs to be killed.

> `lsof -P | grep localhost | grep dart`

Also, if you want to analysis the results, you may have to manually
remove any incomplete replicas from the results file prior to analysis.

## NOTICE

The code in the treatments was designed to be extra complicated to
help better test the build process.
Obviously the treatment code can be greatly simplified, improved, etc.
The goal is to test the Dart build speeds, not write good Dart code.
