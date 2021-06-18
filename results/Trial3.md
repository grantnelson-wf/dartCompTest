# Trial 3

## Setup

- Changed the main experiment to run using treatments from `addTreatments_FileVsLib_Generated_Change10_Dart2js`
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
1 1 0 File_Dep 4.93957
2 1 1 Library_Dep 4.78297
3 2 0 File_Dep 5.13756
4 2 1 Library_Dep 4.39256
5 3 1 Library_Dep 4.42902
6 3 0 File_Dep 4.88703
7 4 0 File_Dep 4.99906
8 4 1 Library_Dep 4.39942
9 5 1 Library_Dep 4.49695
10 5 0 File_Dep 4.86735
11 6 1 Library_Dep 4.36453
12 6 0 File_Dep 4.96404
13 7 0 File_Dep 5.02585
14 7 1 Library_Dep 4.40614
15 8 1 Library_Dep 4.43764
16 8 0 File_Dep 4.97858
17 9 1 Library_Dep 4.51348
18 9 0 File_Dep 4.97015
19 10 1 Library_Dep 4.48297
20 10 0 File_Dep 4.98856
```

| index | name        | Average (seconds) |
|-------|-------------|-------------------|
| 0     | File_Dep    | 4.975775          |
| 1     | Library_Dep | 4.470568          |

## Results

```Plain
               sum_sq    df           F        PR(>F)
order        0.006754   1.0    0.640073  4.354028e-01
index        1.279008   1.0  121.209072  7.092262e-09
order:index  0.004559   1.0    0.432031  5.203383e-01
Residual     0.168833  16.0         NaN           NaN
```

There is significant evidence to reject the null hypothesis.
The null hypothesis is that there is no difference between File_Dep and Library_Dep treatments when building after 10% of the leaf nodes (51 leaves) were changed.
Meaning that file and library deps likely take a different amount of time to build for the current configuration.
