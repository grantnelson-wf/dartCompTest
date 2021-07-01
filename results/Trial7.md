# Trial 7

- Rate:   2.0
- Scalar: 2.0
- Depth:  8
- Group:  15
- Pub 2.13.4

## Procedure

- Changed the main experiment to run using treatments from `addTreatments_FileVsLib_Generated_Dart2js`
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
1 1 0 File_Dep 4.36635
2 1 1 Library_Dep 4.25752
3 2 0 File_Dep 4.23206
4 2 1 Library_Dep 4.24161
5 3 1 Library_Dep 4.24824
6 3 0 File_Dep 4.19231
7 4 1 Library_Dep 4.21916
8 4 0 File_Dep 4.26501
9 5 0 File_Dep 4.19315
10 5 1 Library_Dep 4.31832
11 6 0 File_Dep 4.18505
12 6 1 Library_Dep 4.27870
13 7 1 Library_Dep 4.32383
14 7 0 File_Dep 4.20523
15 8 1 Library_Dep 4.33858
16 8 0 File_Dep 4.21456
17 9 0 File_Dep 4.29787
18 9 1 Library_Dep 4.24891
19 10 0 File_Dep 4.27876
20 10 1 Library_Dep 4.33558
21 11 1 Library_Dep 4.27388
22 11 0 File_Dep 4.22951
23 12 1 Library_Dep 4.23863
24 12 0 File_Dep 4.27315
25 13 0 File_Dep 4.21508
26 13 1 Library_Dep 4.16124
27 14 1 Library_Dep 4.25558
28 14 0 File_Dep 4.22799
29 15 0 File_Dep 4.30237
30 15 1 Library_Dep 4.25256
31 16 1 Library_Dep 4.32938
32 16 0 File_Dep 4.22885
33 17 0 File_Dep 4.28587
34 17 1 Library_Dep 4.35749
35 18 1 Library_Dep 4.23531
36 18 0 File_Dep 4.20657
37 19 0 File_Dep 4.15784
38 19 1 Library_Dep 4.27565
39 20 1 Library_Dep 4.23060
40 20 0 File_Dep 4.17634
41 21 0 File_Dep 4.20354
42 21 1 Library_Dep 4.26964
43 22 0 File_Dep 4.25080
44 22 1 Library_Dep 4.25572
45 23 1 Library_Dep 4.21862
46 23 0 File_Dep 4.17795
47 24 1 Library_Dep 4.22083
48 24 0 File_Dep 4.21301
49 25 1 Library_Dep 4.24812
50 25 0 File_Dep 4.17780
51 26 1 Library_Dep 4.21282
52 26 0 File_Dep 4.41015
53 27 0 File_Dep 4.20775
54 27 1 Library_Dep 4.29567
55 28 0 File_Dep 4.24156
56 28 1 Library_Dep 4.18915
57 29 0 File_Dep 4.20719
58 29 1 Library_Dep 4.23781
59 30 0 File_Dep 4.24782
60 30 1 Library_Dep 4.16742
61 31 0 File_Dep 4.28291
62 31 1 Library_Dep 4.25465
63 32 1 Library_Dep 4.19962
64 32 0 File_Dep 4.18667
65 33 1 Library_Dep 4.20393
66 33 0 File_Dep 4.20334
67 34 1 Library_Dep 4.24801
68 34 0 File_Dep 4.19233
69 35 1 Library_Dep 4.19591
70 35 0 File_Dep 4.14097
71 36 0 File_Dep 4.45058
72 36 1 Library_Dep 4.28241
73 37 1 Library_Dep 4.14239
74 37 0 File_Dep 4.23463
75 38 1 Library_Dep 4.27683
76 38 0 File_Dep 4.17903
77 39 1 Library_Dep 4.27885
78 39 0 File_Dep 4.18802
79 40 1 Library_Dep 4.23954
80 40 0 File_Dep 4.15527
81 41 0 File_Dep 4.21074
82 41 1 Library_Dep 4.24168
83 42 1 Library_Dep 4.17734
84 42 0 File_Dep 4.21998
85 43 0 File_Dep 4.12373
86 43 1 Library_Dep 4.29566
87 44 1 Library_Dep 4.16383
88 44 0 File_Dep 4.20718
89 45 1 Library_Dep 4.35549
90 45 0 File_Dep 4.56856
91 46 0 File_Dep 4.36866
92 46 1 Library_Dep 4.19725
93 47 1 Library_Dep 4.23632
94 47 0 File_Dep 4.22073
95 48 1 Library_Dep 4.21648
96 48 0 File_Dep 4.24909
97 49 0 File_Dep 4.17329
98 49 1 Library_Dep 4.22681
99 50 0 File_Dep 4.18471
100 50 1 Library_Dep 4.28077
101 51 1 Library_Dep 4.17142
102 51 0 File_Dep 4.17379
103 52 1 Library_Dep 4.20347
104 52 0 File_Dep 4.22867
105 53 1 Library_Dep 4.21545
106 53 0 File_Dep 4.15934
107 54 0 File_Dep 4.16400
108 54 1 Library_Dep 4.21083
109 55 1 Library_Dep 4.11758
110 55 0 File_Dep 4.20424
111 56 0 File_Dep 4.14549
112 56 1 Library_Dep 4.29193
113 57 0 File_Dep 4.17592
114 57 1 Library_Dep 4.22127
115 58 0 File_Dep 4.26923
116 58 1 Library_Dep 4.19840
117 59 1 Library_Dep 4.17249
118 59 0 File_Dep 4.23396
119 60 0 File_Dep 4.29207
120 60 1 Library_Dep 4.28451
121 61 0 File_Dep 4.19140
122 61 1 Library_Dep 4.18161
123 62 1 Library_Dep 4.20582
124 62 0 File_Dep 4.24381
125 63 0 File_Dep 4.20399
126 63 1 Library_Dep 4.20910
127 64 0 File_Dep 4.24375
128 64 1 Library_Dep 4.21146
129 65 0 File_Dep 4.48903
130 65 1 Library_Dep 4.26495
131 66 0 File_Dep 4.23331
132 66 1 Library_Dep 4.19868
133 67 1 Library_Dep 4.12710
134 67 0 File_Dep 4.23916
135 68 0 File_Dep 4.25243
136 68 1 Library_Dep 4.24561
137 69 1 Library_Dep 4.16009
138 69 0 File_Dep 4.20050
139 70 1 Library_Dep 4.27561
140 70 0 File_Dep 4.22048
141 71 0 File_Dep 4.17818
142 71 1 Library_Dep 4.19062
143 72 1 Library_Dep 4.18867
144 72 0 File_Dep 4.16854
145 73 0 File_Dep 4.13018
146 73 1 Library_Dep 4.22819
147 74 1 Library_Dep 4.18501
148 74 0 File_Dep 4.81778
149 75 0 File_Dep 4.18394
150 75 1 Library_Dep 4.26476
151 76 0 File_Dep 4.24953
152 76 1 Library_Dep 4.25539
153 77 1 Library_Dep 4.19760
154 77 0 File_Dep 4.23323
155 78 1 Library_Dep 4.26204
156 78 0 File_Dep 4.25423
157 79 1 Library_Dep 4.18342
158 79 0 File_Dep 4.25588
159 80 0 File_Dep 4.16505
160 80 1 Library_Dep 4.14635
161 81 0 File_Dep 4.22798
162 81 1 Library_Dep 4.20598
163 82 1 Library_Dep 4.17667
164 82 0 File_Dep 4.38653
165 83 1 Library_Dep 4.20883
166 83 0 File_Dep 4.23419
167 84 0 File_Dep 4.29270
168 84 1 Library_Dep 4.24996
169 85 1 Library_Dep 4.28437
170 85 0 File_Dep 4.23353
171 86 1 Library_Dep 4.23538
172 86 0 File_Dep 4.23177
173 87 0 File_Dep 4.20743
174 87 1 Library_Dep 4.23241
175 88 1 Library_Dep 4.21667
176 88 0 File_Dep 4.21574
177 89 0 File_Dep 4.24651
178 89 1 Library_Dep 4.16852
179 90 0 File_Dep 4.16381
180 90 1 Library_Dep 4.20040
181 91 0 File_Dep 4.21930
182 91 1 Library_Dep 4.29092
183 92 0 File_Dep 4.33184
184 92 1 Library_Dep 4.28207
185 93 1 Library_Dep 4.22241
186 93 0 File_Dep 4.20566
187 94 1 Library_Dep 4.22208
188 94 0 File_Dep 4.29002
189 95 0 File_Dep 4.09493
190 95 1 Library_Dep 4.22786
191 96 1 Library_Dep 4.12501
192 96 0 File_Dep 4.23111
193 97 0 File_Dep 4.13634
194 97 1 Library_Dep 4.17912
195 98 0 File_Dep 4.12647
196 98 1 Library_Dep 4.17086
197 99 1 Library_Dep 4.17209
198 99 0 File_Dep 4.19419
199 100 1 Library_Dep 4.25942
200 100 0 File_Dep 4.14872
```

| index | name        | Average (seconds) |
|-------|-------------|-------------------|
| 0     | File_Dep    | 4.23432           |
| 1     | Library_Dep | 4.23050           |
|       | diff        | 0.038             |

## Results

```Plain
               sum_sq     df         F    PR(>F)
order        0.023597    1.0  4.217357  0.041340
index        0.000727    1.0  0.129966  0.718855
order:index  0.012643    1.0  2.259554  0.134402
Residual     1.096681  196.0       NaN       NaN
```

No enough evidence to reject the null hypothesis.
