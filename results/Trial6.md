# Trial 5

- Rate:   2.0
- Scalar: 2.0
- Depth:  10
- Group:  15
- Pub 2.13.4

## Data

Burn-in times

```Plain
order replicate index name seconds
1 1 0 File_Dep 62.75314
2 1 1 Library_Dep 49.60512
```

Run 1

```Plain
order replicate index name seconds
3 2 1 Library_Dep 5.75145
4 2 0 File_Dep 5.57508
5 3 0 File_Dep 5.59787
6 3 1 Library_Dep 5.54282
7 4 1 Library_Dep 5.56862
8 4 0 File_Dep 5.61536
9 5 1 Library_Dep 5.47166
10 5 0 File_Dep 5.60665
11 6 1 Library_Dep 5.66320
12 6 0 File_Dep 5.69560
13 7 0 File_Dep 5.56822
14 7 1 Library_Dep 5.61922
15 8 1 Library_Dep 5.60087
16 8 0 File_Dep 5.76217
17 9 0 File_Dep 5.61310
18 9 1 Library_Dep 5.57520
19 10 1 Library_Dep 5.64201
20 10 0 File_Dep 5.58898
```

Run 2

```Plain
order replicate index name seconds
1 1 1 Library_Dep 5.87178
2 1 0 File_Dep 5.73774
3 2 0 File_Dep 5.69418
4 2 1 Library_Dep 5.66578
5 3 0 File_Dep 5.71755
6 3 1 Library_Dep 5.66441
7 4 1 Library_Dep 5.64900
8 4 0 File_Dep 5.78539
9 5 0 File_Dep 5.78396
10 5 1 Library_Dep 5.66109
11 6 1 Library_Dep 5.76699
12 6 0 File_Dep 5.67340
13 7 0 File_Dep 5.78942
14 7 1 Library_Dep 5.64378
15 8 0 File_Dep 5.72723
16 8 1 Library_Dep 5.65301
17 9 0 File_Dep 5.79364
18 9 1 Library_Dep 5.71453
19 10 1 Library_Dep 5.64238
20 10 0 File_Dep 5.82747
```

## Results

Run 1

```Plain
               sum_sq    df         F    PR(>F)
order        0.000469   1.0  0.082942  0.777569
index        0.001900   1.0  0.335853  0.571441
order:index  0.002722   1.0  0.481087  0.499278
Residual     0.079200  14.0       NaN       NaN
```

Run 2

```Plain
               sum_sq    df         F    PR(>F)
order        0.000219   1.0  0.063906  0.803645
index        0.017760   1.0  5.186190  0.036850
order:index  0.016466   1.0  4.808144  0.043453
Residual     0.054793  16.0       NaN       NaN
```
