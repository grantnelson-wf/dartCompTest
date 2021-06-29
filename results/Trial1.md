# Trial 1

- Rate:   2.0
- Scalar: 2.0
- Depth:  10
- Group:  15

## Procedure

- Change the main experiment to run using treatments from `addTreatments_FileVsLib_Generated_Dart2js`.
- `go run ./generators/treegen/main.go -del -gen -pubget -out ./treatments/filedeps_gen -rate 2 -depth 10 -scalar 2 -group 15`
- `go run ./generators/treegen/main.go -del -gen -pubget -out ./treatments/libdeps_gen -lib -rate 2 -depth 10 -scalar 2 -group 15`
- `go run ./experiment/main.go`
- `(cd analysis && python3 anova.py)`

## Generated

```Plain
total nodes:      4093
branch nodes:     2045
leaf nodes:       2048
groups:           273
nodes per group:  14.9927
child per branch: 2.0010
```

## Data

Performed first build burn-in so that these don't run the actions part of the build.

```Plain
order replicate index name seconds
1 1 0 File_Dep 6.51455
2 1 1 Library_Dep 6.15818
3 2 1 Library_Dep 5.85560
4 2 0 File_Dep 6.54592
5 3 1 Library_Dep 5.95343
6 3 0 File_Dep 6.60267
7 4 0 File_Dep 6.49151
8 4 1 Library_Dep 5.88836
```

Dart build server locked up during replicate 5 (and several other times).

| index | name        | Average (seconds) |
|-------|-------------|-------------------|
| 0     | File_Dep    | 6.5386625         |
| 1     | Library_Dep | 5.9638925         |

Not really enough replicas a strong result.
I reran this trial on a different day and got
different build speeds.

Burn in:

```Plain
1 1 0 File_Dep 65.26596
2 1 1 Library_Dep 62.68349
```

Additional measurements:

``` Plain
order replicate index name seconds
1 1 0 File_Dep 7.42410
2 1 1 Library_Dep 7.18000
3 2 1 Library_Dep 6.75744
4 2 0 File_Dep 6.60424
5 3 0 File_Dep 7.09020
6 3 1 Library_Dep 7.22092
7 4 0 File_Dep 7.21278
8 4 1 Library_Dep 7.18945
9 5 0 File_Dep 9.12963
10 5 1 Library_Dep 7.89248
11 6 0 File_Dep 7.37536
12 6 1 Library_Dep 7.29014
13 7 0 File_Dep 7.36424
14 7 1 Library_Dep 7.50907
15 8 0 File_Dep 7.06371
16 8 1 Library_Dep 7.73773
17 9 0 File_Dep 7.14663
18 9 1 Library_Dep 8.99935
19 10 0 File_Dep 9.32552
20 10 1 Library_Dep 8.46731
```

## Results

```Plain
               sum_sq   df          F    PR(>F)
order        0.006789  1.0   0.598926  0.482184
index        0.660721  1.0  58.289553  0.001581
order:index  0.010114  1.0   0.892302  0.398341
Residual     0.045341  4.0        NaN       NaN
```

There is significant evidence to reject the null hypothesis.
The null hypothesis is that there is no difference between File_Dep and Library_Dep treatments when building after deleting the `build` folder.
Meaning that file and library deps likely take a different amount of time to build for the current configuration.
