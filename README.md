# Dart Compilation Experiment

Experimental test of dart compiling different type of code bases.

## References

- [Guide to creating dart library packages](https://dart.dev/guides/libraries/create-library-packages)
- [Write-up of results](https://staging.wdesk.org/a/QWNjb3VudB80ODMwNDcyMDE4MzI5NjAw/presentation/0dde949b3bc5478a80f197e01c5ab62f/r/-1/v/1/sec/0dde949b3bc5478a80f197e01c5ab62f_549755813969)
- [Export of write-up (Jul 2, 21)](./results/Dart%20Build%20Experiment%20Results.pdf)

## Setting up the Experiment

> `./setup.sh`

### Generating treatments

For some treatments the code has to be generated. They can be generated with something like:

> `go run ./generators/treegen/main.go -del -gen -pubget -out ./treatments/filedeps_gen -rate 2 -depth 10 -scalar 10 -group 15`
>
> `go run ./generators/treegen/main.go -del -gen -pubget -out ./treatments/libdeps_gen -lib -rate 2 -depth 10 -scalar 10 -group 15`

Based on what size and what you are testing will depend on how you want to configure that generation.
To see possible configurations get help with `go run ./generators/treegen/main.go -h`.

## Running Experiment

Configure the trial that should be run in [`./experiment/main.go`](./experiment/main.go),

Then run

> `go run ./experiment/main.go`

## Running Analysis on Results

You may need to add some python packages, use pip as needed.

This assumes the results are in `./results.txt`

> `(cd analysis && python3 anova.py)`

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

or

> `killall dart`

Also, if you want to analysis the results, you may have to manually
remove any incomplete replicas from the results file prior to analysis.

## NOTICE

The code in the treatments was designed to be extra complicated to
help better test the build process.
Obviously the treatment code can be greatly simplified, improved, etc.
The goal is to test the Dart build speeds, not write good Dart code.
