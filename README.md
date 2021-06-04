# Dart Compilation Experiment

Experimental test of dart compiling different type of code bases.

## Setting up the Experiment

> `./setup.sh`

## Running Experiment

> `go run ./experiment/main.go`

## Running Analysis on Results

You may need to add some python packages, use pip as needed.

This assumes the results are in `./results.txt`

> `cd analysis`
> `python3 anova.py`

## Running Individual Treatments

Each one may have it's own unique method for running but in general it will look like.

> `cd treatments/filedeps`
> `pub get`
> `webdev serve --no-injected-client`

## NOTICE

The code in the treatments was designed to be extra complicated to
help better test the build process.
Obviously the treatment code can be greatly simplified, improved, etc.
The goal is to test the Dart build speeds, not write good Dart code.
