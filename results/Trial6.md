# Trial 6

- Rate:   2.0
- Scalar: 2.0
- Depth:  10
- Group:  15
- Pub 2.13.4

## Procedure

- Changed the main experiment to run using treatments from `addTreatments_FileVsLib_Generated_Dart2js`
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

Run 3

```Plain
order replicate index name seconds
1 1 1 Library_Dep 7.09551
2 1 0 File_Dep 7.66610
3 2 0 File_Dep 7.75475
4 2 1 Library_Dep 7.30435
5 3 1 Library_Dep 8.46022
6 3 0 File_Dep 7.24435
7 4 0 File_Dep 7.75402
8 4 1 Library_Dep 7.58249
9 5 0 File_Dep 8.03429
10 5 1 Library_Dep 8.24396
11 6 0 File_Dep 7.00295
12 6 1 Library_Dep 6.78505
13 7 1 Library_Dep 6.52625
14 7 0 File_Dep 5.89287
15 8 0 File_Dep 5.93106
16 8 1 Library_Dep 6.82208
17 9 0 File_Dep 6.70651
18 9 1 Library_Dep 6.47169
19 10 1 Library_Dep 6.57685
20 10 0 File_Dep 6.25752
21 11 1 Library_Dep 6.38096
22 11 0 File_Dep 6.48295
23 12 0 File_Dep 6.34721
24 12 1 Library_Dep 6.93136
25 13 1 Library_Dep 6.59493
26 13 0 File_Dep 6.32709
27 14 0 File_Dep 7.09766
28 14 1 Library_Dep 6.54467
29 15 1 Library_Dep 6.63510
30 15 0 File_Dep 6.41650
31 16 0 File_Dep 7.10015
32 16 1 Library_Dep 6.82895
33 17 0 File_Dep 6.23882
34 17 1 Library_Dep 7.12426
35 18 1 Library_Dep 6.78557
36 18 0 File_Dep 6.64210
37 19 0 File_Dep 6.36287
38 19 1 Library_Dep 6.43663
39 20 1 Library_Dep 6.38661
40 20 0 File_Dep 6.71029
41 21 0 File_Dep 6.93613
42 21 1 Library_Dep 6.54361
43 22 0 File_Dep 6.65098
44 22 1 Library_Dep 6.03727
45 23 1 Library_Dep 6.69431
46 23 0 File_Dep 6.52451
47 24 0 File_Dep 6.83921
48 24 1 Library_Dep 6.35387
49 25 0 File_Dep 6.46794
50 25 1 Library_Dep 6.45372
51 26 1 Library_Dep 6.88066
52 26 0 File_Dep 6.93272
53 27 0 File_Dep 6.41921
54 27 1 Library_Dep 6.32352
55 28 1 Library_Dep 6.44638
56 28 0 File_Dep 6.74469
57 29 1 Library_Dep 6.91041
58 29 0 File_Dep 8.54345
59 30 0 File_Dep 8.26135
60 30 1 Library_Dep 7.88314
61 31 1 Library_Dep 6.94196
62 31 0 File_Dep 6.11020
63 32 0 File_Dep 8.26766
64 32 1 Library_Dep 6.19703
65 33 0 File_Dep 6.27349
66 33 1 Library_Dep 6.62704
67 34 0 File_Dep 8.04357
68 34 1 Library_Dep 7.77403
69 35 1 Library_Dep 7.60750
70 35 0 File_Dep 7.06659
71 36 0 File_Dep 6.62644
72 36 1 Library_Dep 6.50598
73 37 1 Library_Dep 6.56120
74 37 0 File_Dep 6.79681
75 38 1 Library_Dep 6.55773
76 38 0 File_Dep 6.05035
77 39 1 Library_Dep 5.78338
78 39 0 File_Dep 6.00535
79 40 0 File_Dep 5.94188
80 40 1 Library_Dep 5.91834
81 41 0 File_Dep 6.11343
82 41 1 Library_Dep 5.93247
83 42 0 File_Dep 6.05378
84 42 1 Library_Dep 5.88793
85 43 0 File_Dep 5.82503
86 43 1 Library_Dep 5.85036
87 44 1 Library_Dep 5.96529
88 44 0 File_Dep 6.55566
89 45 0 File_Dep 6.48045
90 45 1 Library_Dep 7.03050
91 46 0 File_Dep 6.77435
92 46 1 Library_Dep 8.20520
93 47 0 File_Dep 7.54903
94 47 1 Library_Dep 6.64561
95 48 0 File_Dep 6.53042
96 48 1 Library_Dep 6.62760
97 49 0 File_Dep 6.75985
98 49 1 Library_Dep 6.50441
99 50 1 Library_Dep 6.68018
100 50 0 File_Dep 6.63914
101 51 0 File_Dep 6.08512
102 51 1 Library_Dep 5.60934
103 52 1 Library_Dep 5.61682
104 52 0 File_Dep 5.72960
105 53 0 File_Dep 5.64526
106 53 1 Library_Dep 5.53658
107 54 1 Library_Dep 5.59920
108 54 0 File_Dep 5.69487
109 55 1 Library_Dep 5.47884
110 55 0 File_Dep 5.54691
111 56 1 Library_Dep 5.52702
112 56 0 File_Dep 5.93963
113 57 1 Library_Dep 5.60972
114 57 0 File_Dep 5.61305
115 58 0 File_Dep 5.58608
116 58 1 Library_Dep 5.51834
117 59 1 Library_Dep 5.58126
118 59 0 File_Dep 5.60478
119 60 0 File_Dep 5.58710
120 60 1 Library_Dep 5.73305
121 61 1 Library_Dep 5.57890
122 61 0 File_Dep 5.57807
123 62 1 Library_Dep 5.69491
124 62 0 File_Dep 5.59581
125 63 0 File_Dep 5.57820
126 63 1 Library_Dep 5.49114
127 64 0 File_Dep 5.71739
128 64 1 Library_Dep 5.88273
129 65 0 File_Dep 5.67007
130 65 1 Library_Dep 5.54883
131 66 1 Library_Dep 5.64631
132 66 0 File_Dep 5.59886
133 67 1 Library_Dep 5.49876
134 67 0 File_Dep 5.55137
135 68 1 Library_Dep 5.64269
136 68 0 File_Dep 5.58969
137 69 1 Library_Dep 5.51374
138 69 0 File_Dep 5.58416
139 70 1 Library_Dep 5.59766
140 70 0 File_Dep 5.58877
141 71 1 Library_Dep 5.52110
142 71 0 File_Dep 5.60240
143 72 1 Library_Dep 5.61208
144 72 0 File_Dep 5.86957
145 73 1 Library_Dep 5.60362
146 73 0 File_Dep 5.67780
147 74 0 File_Dep 5.54106
148 74 1 Library_Dep 5.60595
149 75 0 File_Dep 5.59726
150 75 1 Library_Dep 5.52004
151 76 0 File_Dep 5.67483
152 76 1 Library_Dep 5.59024
153 77 1 Library_Dep 5.50470
154 77 0 File_Dep 5.57508
155 78 1 Library_Dep 5.58692
156 78 0 File_Dep 5.59695
157 79 0 File_Dep 5.50957
158 79 1 Library_Dep 5.49589
159 80 1 Library_Dep 5.51530
160 80 0 File_Dep 5.56580
161 81 1 Library_Dep 5.51332
162 81 0 File_Dep 5.64899
163 82 1 Library_Dep 5.56641
164 82 0 File_Dep 5.57987
165 83 0 File_Dep 5.50390
166 83 1 Library_Dep 5.65948
167 84 1 Library_Dep 6.81050
168 84 0 File_Dep 6.93022
169 85 1 Library_Dep 6.53259
170 85 0 File_Dep 6.38063
171 86 0 File_Dep 5.78692
172 86 1 Library_Dep 5.90292
173 87 0 File_Dep 5.84558
174 87 1 Library_Dep 5.84486
175 88 1 Library_Dep 6.10398
176 88 0 File_Dep 5.77895
177 89 0 File_Dep 5.84412
178 89 1 Library_Dep 5.87464
179 90 1 Library_Dep 5.86576
180 90 0 File_Dep 5.90688
181 91 0 File_Dep 5.83902
182 91 1 Library_Dep 5.89348
183 92 0 File_Dep 5.87342
184 92 1 Library_Dep 5.83763
185 93 1 Library_Dep 5.86574
186 93 0 File_Dep 5.72303
187 94 0 File_Dep 5.89482
188 94 1 Library_Dep 5.84184
189 95 0 File_Dep 5.91712
190 95 1 Library_Dep 6.11502
191 96 1 Library_Dep 5.76181
192 96 0 File_Dep 5.84616
193 97 1 Library_Dep 5.82204
194 97 0 File_Dep 5.96322
195 98 0 File_Dep 5.76444
196 98 1 Library_Dep 5.81012
197 99 0 File_Dep 5.80524
198 99 1 Library_Dep 5.81201
199 100 1 Library_Dep 5.81785
200 100 0 File_Dep 5.87666
```

Run 4

```Plain
order replicate index name seconds
1 1 1 Library_Dep 7.25096
2 1 0 File_Dep 6.90370
3 2 1 Library_Dep 7.17884
4 2 0 File_Dep 6.86137
5 3 0 File_Dep 6.96157
6 3 1 Library_Dep 6.53684
7 4 0 File_Dep 6.72009
8 4 1 Library_Dep 6.76053
9 5 0 File_Dep 7.03315
10 5 1 Library_Dep 5.77272
11 6 0 File_Dep 6.84290
12 6 1 Library_Dep 6.78203
13 7 1 Library_Dep 6.36618
14 7 0 File_Dep 6.67854
15 8 1 Library_Dep 5.82083
16 8 0 File_Dep 5.75365
17 9 0 File_Dep 7.00490
18 9 1 Library_Dep 6.07025
19 10 0 File_Dep 6.59549
20 10 1 Library_Dep 7.31720
21 11 0 File_Dep 7.30897
22 11 1 Library_Dep 6.28299
23 12 1 Library_Dep 7.35243
24 12 0 File_Dep 6.47004
25 13 0 File_Dep 6.58453
26 13 1 Library_Dep 6.25541
27 14 1 Library_Dep 6.31978
28 14 0 File_Dep 6.39617
29 15 1 Library_Dep 6.32033
30 15 0 File_Dep 6.45071
31 16 0 File_Dep 6.73191
32 16 1 Library_Dep 6.13482
33 17 1 Library_Dep 6.22652
34 17 0 File_Dep 6.16739
35 18 0 File_Dep 6.12279
36 18 1 Library_Dep 6.23465
37 19 1 Library_Dep 6.15040
38 19 0 File_Dep 6.25857
39 20 0 File_Dep 6.24720
40 20 1 Library_Dep 6.09089
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

Run 3

This run showed a difference in order meaning something changed during the run.

```Plain
                sum_sq     df           F        PR(>F)
order        42.133695    1.0  142.913831  4.286981e-25
index         0.050572    1.0    0.171538  6.792019e-01
order:index   0.000208    1.0    0.000704  9.788601e-01
Residual     57.784500  196.0         NaN           NaN
```

Run 4

This one also had a ramp-up speed seen in the order value.

```Plain
               sum_sq    df          F    PR(>F)
order        1.761442   1.0  11.735307  0.001548
index        0.200598   1.0   1.336452  0.255275
order:index  0.000849   1.0   0.005655  0.940470
Residual     5.403516  36.0        NaN       NaN
```
