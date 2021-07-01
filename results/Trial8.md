# Trial 7

- Rate:   2.0
- Scalar: 2.0
- Depth:  12
- Group:  15
- Pub 2.13.4

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
1 1 1 Library_Dep 11.98652
2 1 0 File_Dep 11.36428
3 2 0 File_Dep 11.31432
4 2 1 Library_Dep 11.29713
5 3 0 File_Dep 11.74825
6 3 1 Library_Dep 11.16280
7 4 0 File_Dep 11.31584
8 4 1 Library_Dep 11.28413
9 5 1 Library_Dep 11.18245
10 5 0 File_Dep 11.32648
11 6 1 Library_Dep 11.42395
12 6 0 File_Dep 11.20961
13 7 1 Library_Dep 11.27401
14 7 0 File_Dep 11.50390
15 8 0 File_Dep 11.36590
16 8 1 Library_Dep 11.16191
17 9 0 File_Dep 11.32236
18 9 1 Library_Dep 11.27549
19 10 1 Library_Dep 11.22665
20 10 0 File_Dep 11.30645
21 11 0 File_Dep 11.12129
22 11 1 Library_Dep 11.10686
23 12 1 Library_Dep 11.49045
24 12 0 File_Dep 11.31152
25 13 0 File_Dep 11.33547
26 13 1 Library_Dep 11.28789
27 14 1 Library_Dep 11.23179
28 14 0 File_Dep 11.15074
29 15 0 File_Dep 11.28930
30 15 1 Library_Dep 11.07873
31 16 1 Library_Dep 11.26586
32 16 0 File_Dep 11.50928
33 17 1 Library_Dep 11.16521
34 17 0 File_Dep 11.22981
35 18 1 Library_Dep 11.31224
36 18 0 File_Dep 11.31807
37 19 0 File_Dep 11.17653
38 19 1 Library_Dep 11.32210
39 20 1 Library_Dep 11.12889
40 20 0 File_Dep 11.34634
41 21 0 File_Dep 11.47328
42 21 1 Library_Dep 11.02886
43 22 1 Library_Dep 11.25980
44 22 0 File_Dep 11.50587
45 23 1 Library_Dep 11.20058
46 23 0 File_Dep 11.25362
47 24 1 Library_Dep 11.27329
48 24 0 File_Dep 11.21778
49 25 0 File_Dep 11.20603
50 25 1 Library_Dep 11.29718
51 26 0 File_Dep 11.23740
52 26 1 Library_Dep 11.10372
53 27 0 File_Dep 10.95650
54 27 1 Library_Dep 10.99143
55 28 0 File_Dep 11.00197
56 28 1 Library_Dep 11.06481
57 29 0 File_Dep 11.06048
58 29 1 Library_Dep 11.03567
59 30 0 File_Dep 11.04138
60 30 1 Library_Dep 11.25652
61 31 0 File_Dep 11.06913
62 31 1 Library_Dep 11.02292
63 32 0 File_Dep 11.10379
64 32 1 Library_Dep 10.99582
65 33 1 Library_Dep 11.08763
66 33 0 File_Dep 11.14099
67 34 0 File_Dep 10.92189
68 34 1 Library_Dep 10.89443
69 35 1 Library_Dep 11.16738
70 35 0 File_Dep 11.07668
71 36 0 File_Dep 11.05353
72 36 1 Library_Dep 10.89013
73 37 0 File_Dep 11.02564
74 37 1 Library_Dep 11.00692
75 38 1 Library_Dep 10.99570
76 38 0 File_Dep 10.98324
77 39 0 File_Dep 11.02457
78 39 1 Library_Dep 11.43345
79 40 1 Library_Dep 11.04576
80 40 0 File_Dep 11.01644
81 41 1 Library_Dep 11.05053
82 41 0 File_Dep 10.98572
83 42 1 Library_Dep 11.06167
84 42 0 File_Dep 11.03250
85 43 0 File_Dep 10.90059
86 43 1 Library_Dep 10.89410
87 44 0 File_Dep 11.01374
88 44 1 Library_Dep 10.92052
89 45 1 Library_Dep 10.85065
90 45 0 File_Dep 11.06336
91 46 1 Library_Dep 11.10618
92 46 0 File_Dep 11.01158
93 47 0 File_Dep 11.00181
94 47 1 Library_Dep 10.95144
95 48 1 Library_Dep 10.92089
96 48 0 File_Dep 10.92724
97 49 0 File_Dep 11.34834
98 49 1 Library_Dep 11.03060
99 50 0 File_Dep 10.94184
100 50 1 Library_Dep 11.10478
101 51 0 File_Dep 10.98752
102 51 1 Library_Dep 10.97099
103 52 0 File_Dep 11.05441
104 52 1 Library_Dep 11.03947
105 53 0 File_Dep 10.93954
106 53 1 Library_Dep 11.18052
107 54 0 File_Dep 11.08557
108 54 1 Library_Dep 11.07692
109 55 1 Library_Dep 11.13002
110 55 0 File_Dep 11.04125
111 56 0 File_Dep 11.03059
112 56 1 Library_Dep 10.98364
113 57 0 File_Dep 11.13384
114 57 1 Library_Dep 11.16461
115 58 1 Library_Dep 11.23616
116 58 0 File_Dep 10.89229
117 59 1 Library_Dep 10.89656
118 59 0 File_Dep 10.95363
119 60 0 File_Dep 10.91843
120 60 1 Library_Dep 11.09598
121 61 1 Library_Dep 10.79209
122 61 0 File_Dep 10.96952
123 62 0 File_Dep 10.97302
124 62 1 Library_Dep 10.95014
125 63 0 File_Dep 11.02784
126 63 1 Library_Dep 11.08331
127 64 0 File_Dep 11.15725
128 64 1 Library_Dep 11.09529
129 65 1 Library_Dep 11.08711
130 65 0 File_Dep 10.98685
131 66 1 Library_Dep 10.99406
132 66 0 File_Dep 11.04481
133 67 0 File_Dep 10.92369
134 67 1 Library_Dep 11.38526
135 68 0 File_Dep 10.93185
136 68 1 Library_Dep 11.04193
137 69 1 Library_Dep 10.84111
138 69 0 File_Dep 11.02590
139 70 1 Library_Dep 10.99973
140 70 0 File_Dep 10.92391
141 71 1 Library_Dep 10.84568
142 71 0 File_Dep 11.03339
143 72 1 Library_Dep 11.27583
144 72 0 File_Dep 10.95535
145 73 1 Library_Dep 10.91620
146 73 0 File_Dep 11.07486
147 74 1 Library_Dep 10.86976
148 74 0 File_Dep 11.21293
149 75 1 Library_Dep 10.79933
150 75 0 File_Dep 11.10696
151 76 0 File_Dep 10.85104
152 76 1 Library_Dep 11.28434
153 77 1 Library_Dep 10.99943
154 77 0 File_Dep 11.06544
155 78 1 Library_Dep 11.02485
156 78 0 File_Dep 11.02137
157 79 0 File_Dep 11.07427
158 79 1 Library_Dep 10.96249
159 80 0 File_Dep 11.02175
160 80 1 Library_Dep 10.92440
161 81 1 Library_Dep 11.03285
162 81 0 File_Dep 10.99050
163 82 1 Library_Dep 11.03163
164 82 0 File_Dep 11.05917
165 83 0 File_Dep 11.05427
166 83 1 Library_Dep 11.02443
167 84 0 File_Dep 11.02179
168 84 1 Library_Dep 11.08862
169 85 0 File_Dep 11.01869
170 85 1 Library_Dep 11.06714
```

Something kicked on during the run because there was a sudden slow down.

```Plain
171 86 1 Library_Dep 11.12559
172 86 0 File_Dep 10.97269
173 87 1 Library_Dep 11.07048
174 87 0 File_Dep 11.41283
175 88 0 File_Dep 11.30771
176 88 1 Library_Dep 13.56406
177 89 1 Library_Dep 12.84520
178 89 0 File_Dep 12.89501
179 90 0 File_Dep 12.87029
180 90 1 Library_Dep 14.37345
181 91 0 File_Dep 13.95096
182 91 1 Library_Dep 13.07651
183 92 0 File_Dep 12.85416
184 92 1 Library_Dep 14.02273
185 93 1 Library_Dep 13.79570
186 93 0 File_Dep 13.92178
187 94 0 File_Dep 14.19925
188 94 1 Library_Dep 13.64527
189 95 1 Library_Dep 14.01940
190 95 0 File_Dep 13.65917
191 96 1 Library_Dep 12.85179
192 96 0 File_Dep 13.56400
193 97 0 File_Dep 12.85139
194 97 1 Library_Dep 11.72642
195 98 1 Library_Dep 11.40028
196 98 0 File_Dep 11.23900
197 99 0 File_Dep 11.29158
198 99 1 Library_Dep 11.17737
199 100 0 File_Dep 11.25893
200 100 1 Library_Dep 11.22047
```

## Results

These results are ignoring the latter half of the data which was slow.

```Plain
               sum_sq     df          F        PR(>F)
order        1.894360    1.0  93.480965  8.066052e-18
index        0.004926    1.0   0.243082  6.226404e-01
order:index  0.018028    1.0   0.889616  3.469521e-01
Residual     3.363933  166.0        NaN           NaN
```

The following is the analysis of just the second half.

```Plain
                sum_sq    df         F    PR(>F)
order         0.290582   1.0  0.188273  0.667937
index         0.093781   1.0  0.060762  0.807231
order:index   0.675510   1.0  0.437674  0.514068
Residual     40.128624  26.0       NaN       NaN
```
