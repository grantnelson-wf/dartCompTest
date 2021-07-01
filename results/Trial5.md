# Trial 5

- Rate:   2.0
- Scalar: 2.0
- Depth:  12
- Group:  15

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

```Plain
order replicate index name seconds
1 1 1 Library_Dep 10.68807
2 1 0 File_Dep 10.85510
3 2 1 Library_Dep 10.46361
4 2 0 File_Dep 10.68946
5 3 0 File_Dep 10.69553
6 3 1 Library_Dep 10.34424
7 4 0 File_Dep 10.84327
8 4 1 Library_Dep 10.41243
9 5 0 File_Dep 10.69724
10 5 1 Library_Dep 10.35954
11 6 0 File_Dep 11.03748
12 6 1 Library_Dep 10.56282
13 7 0 File_Dep 10.81971
14 7 1 Library_Dep 10.54950
15 8 1 Library_Dep 10.56889
16 8 0 File_Dep 10.76138
17 9 0 File_Dep 10.73997
18 9 1 Library_Dep 10.31310
19 10 0 File_Dep 10.82850
20 10 1 Library_Dep 10.24381
```

| index | name        | Average (seconds) |
|-------|-------------|-------------------|
| 0     | File_Dep    | 10.796764         |
| 1     | Library_Dep | 10.450601         |
|       | diff        |  0.346163         |

## Results

```Plain
               sum_sq    df          F    PR(>F)
order        0.014045   1.0   0.947131  0.344935
index        0.592081   1.0  39.926724  0.000010
order:index  0.025535   1.0   1.721929  0.207959
Residual     0.237267  16.0        NaN       NaN
```
