# Trial 5

## Procedure

- Changed the main experiment to run using treatments from `addTreatments_FileVsLib_Generated_Dart2js`
- `go run ./generators/treegen/main.go -del -gen -pubget -out ./treatments/filedeps_gen -rate 2 -depth 12 -scalar 2 -group 15`
- `go run ./generators/treegen/main.go -del -gen -pubget -out ./treatments/libdeps_gen -lib -rate 2 -depth 12 -scalar 2 -group 15`
- `go run ./experiment/main.go`
- `(cd analysis && python3 anova.py)`

## Generated

```Plain
total nodes:      16381
branch nodes:     8189
leaf nodes:       8192
groups:           1093
nodes per group:  14.9872
child per branch: 2.0002
```

## Data
