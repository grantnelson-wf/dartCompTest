# Trial 2

- Rate:   2.0
- Scalar: 2.0
- Depth:  8
- Group:  15

## Procedure

- Change the main experiment to run using treatments from `addTreatments_FileVsLib_Generated_Dart2js`.
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

Burn-in values.

```Plain
order replicate index name seconds
1 1 1 Library_Dep 60.16925
2 1 0 File_Dep 57.68370
```

```Plain
order replicate index name seconds
1 1 1 Library_Dep 5.05607
2 1 0 File_Dep 5.64965
3 2 0 File_Dep 5.66992
4 2 1 Library_Dep 5.06471
5 3 1 Library_Dep 5.01645
6 3 0 File_Dep 5.68681
7 4 1 Library_Dep 5.05373
8 4 0 File_Dep 5.74152
9 5 0 File_Dep 5.65698
10 5 1 Library_Dep 5.06381
11 6 0 File_Dep 5.85612
12 6 1 Library_Dep 5.57470
13 7 1 Library_Dep 5.06077
14 7 0 File_Dep 5.59839
15 8 0 File_Dep 5.80279
16 8 1 Library_Dep 5.01511
17 9 1 Library_Dep 4.98705
18 9 0 File_Dep 5.62232
19 10 0 File_Dep 5.61046
20 10 1 Library_Dep 5.10981
21 11 1 Library_Dep 4.99045
22 11 0 File_Dep 5.62501
23 12 0 File_Dep 5.63862
24 12 1 Library_Dep 5.03726
25 13 1 Library_Dep 5.00102
26 13 0 File_Dep 5.64417
27 14 0 File_Dep 5.62528
28 14 1 Library_Dep 5.19554
29 15 1 Library_Dep 5.01649
30 15 0 File_Dep 5.66564
31 16 0 File_Dep 5.63948
32 16 1 Library_Dep 5.11312
```

Dart build server locked up during replicate 7.
The remaining replicates came from a second run started shortly after the first once the locked dart was killed.

| index | name        | Average (seconds) |
|-------|-------------|-------------------|
| 0     | File_Dep    | 5.670822500       |
| 1     | Library_Dep | 5.084755625       |
|       | diff        |  0.586066875      |

## Results

```Plain
               sum_sq    df           F        PR(>F)
order        0.006132   1.0    0.469122  4.990227e-01
index        2.747795   1.0  210.206034  1.522660e-14
order:index  0.002507   1.0    0.191803  6.647802e-01
Residual     0.366014  28.0         NaN           NaN
```

There is significant evidence to reject the null hypothesis.
The null hypothesis is that there is no difference between File_Dep and Library_Dep treatments when building after deleting the `build` folder.
Meaning that file and library deps likely take a different amount of time to build for the current configuration.
