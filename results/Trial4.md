# Trial 3

## Setup

- Changed the main experiment to run using treatments from `addTreatments_FileVsLib_Generated_Change10_Dart2js` but changing the `frac` to `0.01`
- `go run ./generators/treegen/main.go -del -gen -pubget -out ./treatments/filedeps_gen -rate 2 -depth 8 -scalar 2 -group 15`
- `go run ./generators/treegen/main.go -del -gen -pubget -out ./treatments/libdeps_gen -lib -rate 2 -depth 8 -scalar 2 -group 15`
- `go run ./experiment/main.go`
- `(cd analysis && python3 anova.py)`

## Generated

```Plain
total nodes:      1021
branch nodes:     509
leaf nodes:       512
groups:           69
nodes per group:  14.7971
child per branch: 2.0039
```

## Data

```Plain
order replicate index name seconds
1 1 1 Library_Dep 4.32590
2 1 0 File_Dep 4.85879
3 2 1 Library_Dep 4.53293
4 2 0 File_Dep 4.77423
5 3 0 File_Dep 4.80333
6 3 1 Library_Dep 4.42559
7 4 1 Library_Dep 4.38946
8 4 0 File_Dep 4.99289
9 5 0 File_Dep 4.84239
10 5 1 Library_Dep 4.37916
11 6 1 Library_Dep 4.39356
12 6 0 File_Dep 4.81284
13 7 1 Library_Dep 4.33760
14 7 0 File_Dep 4.75039
15 8 0 File_Dep 4.75814
16 8 1 Library_Dep 4.37150
17 9 1 Library_Dep 4.35046
18 9 0 File_Dep 4.97128
19 10 0 File_Dep 4.82174
20 10 1 Library_Dep 4.42451
```

| index | name        | Average (seconds) |
|-------|-------------|-------------------|
| 0     | File_Dep    | 4.838602          |
| 1     | Library_Dep | 4.393067          |

## Results

```Plain
               sum_sq    df           F        PR(>F)
order        0.000508   1.0    0.088114  7.704070e-01
index        0.992988   1.0  172.181131  5.578636e-10
order:index  0.001323   1.0    0.229367  6.384735e-01
Residual     0.092274  16.0         NaN           NaN
```

There is significant evidence to reject the null hypothesis.
The null hypothesis is that there is no difference between File_Dep and Library_Dep treatments when building after 1% of the leaf nodes (5 leaves) were changed.
Meaning that file and library deps likely take a different amount of time to build for the current configuration.
