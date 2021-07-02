# Trial 10

- Block Chain Example
- Pub 2.13.4

## Procedure

- Changed the main experiment to run using treatments from `addTreatments_FileVsLib_Dart2js`
- `go run ./experiment/main.go`
- `(cd analysis && python3 anova.py)`

## Data

Burn-in

```Plain
order replicate index name seconds
1 1 0 File_Dep 19.77898
2 1 1 Library_Dep 17.57122
```

```Plain
order replicate index name seconds
3 2 0 File_Dep 3.82404
4 2 1 Library_Dep 3.57120
5 3 1 Library_Dep 3.56140
6 3 0 File_Dep 3.82973
7 4 1 Library_Dep 3.60557
8 4 0 File_Dep 4.42753
9 5 0 File_Dep 4.55783
10 5 1 Library_Dep 5.15812
11 6 0 File_Dep 5.58252
12 6 1 Library_Dep 4.47566
13 7 0 File_Dep 4.59063
14 7 1 Library_Dep 4.23904
15 8 1 Library_Dep 4.24687
16 8 0 File_Dep 4.51348
17 9 1 Library_Dep 4.26612
18 9 0 File_Dep 4.45709
19 10 0 File_Dep 4.43892
20 10 1 Library_Dep 4.20988
21 11 0 File_Dep 4.59822
22 11 1 Library_Dep 4.32558
23 12 0 File_Dep 4.45973
24 12 1 Library_Dep 4.32528
25 13 0 File_Dep 4.61447
26 13 1 Library_Dep 4.24446
27 14 1 Library_Dep 4.36783
28 14 0 File_Dep 4.44542
29 15 1 Library_Dep 4.28826
30 15 0 File_Dep 4.61570
31 16 0 File_Dep 4.45045
32 16 1 Library_Dep 4.33688
33 17 0 File_Dep 4.54002
34 17 1 Library_Dep 4.25350
35 18 1 Library_Dep 4.29440
36 18 0 File_Dep 4.44658
37 19 0 File_Dep 4.58446
38 19 1 Library_Dep 4.24323
39 20 0 File_Dep 4.43883
40 20 1 Library_Dep 4.87521
41 21 1 Library_Dep 4.16589
42 21 0 File_Dep 4.55827
43 22 0 File_Dep 4.48668
44 22 1 Library_Dep 4.26124
45 23 0 File_Dep 4.47193
46 23 1 Library_Dep 4.22711
47 24 0 File_Dep 4.38269
48 24 1 Library_Dep 4.28434
49 25 1 Library_Dep 4.22559
50 25 0 File_Dep 4.44997
51 26 1 Library_Dep 4.26744
52 26 0 File_Dep 4.37963
53 27 1 Library_Dep 4.33415
54 27 0 File_Dep 4.42890
55 28 0 File_Dep 4.37460
56 28 1 Library_Dep 4.30227
57 29 0 File_Dep 4.46572
58 29 1 Library_Dep 4.47257
59 30 0 File_Dep 4.44875
60 30 1 Library_Dep 4.20377
61 31 0 File_Dep 5.07889
62 31 1 Library_Dep 4.72598
63 32 1 Library_Dep 4.39172
64 32 0 File_Dep 4.35859
65 33 1 Library_Dep 4.27302
66 33 0 File_Dep 4.53145
67 34 0 File_Dep 4.58097
68 34 1 Library_Dep 4.25823
69 35 1 Library_Dep 4.31306
70 35 0 File_Dep 4.54541
71 36 0 File_Dep 4.45608
72 36 1 Library_Dep 4.33682
73 37 0 File_Dep 4.42652
74 37 1 Library_Dep 4.25153
75 38 0 File_Dep 4.43235
76 38 1 Library_Dep 4.20563
77 39 1 Library_Dep 4.34305
78 39 0 File_Dep 4.55426
79 40 1 Library_Dep 4.25739
80 40 0 File_Dep 4.44409
81 41 1 Library_Dep 4.32300
82 41 0 File_Dep 4.54855
83 42 0 File_Dep 4.48535
84 42 1 Library_Dep 4.28378
85 43 0 File_Dep 4.57816
86 43 1 Library_Dep 4.32258
87 44 0 File_Dep 4.39250
88 44 1 Library_Dep 4.29430
89 45 0 File_Dep 4.52763
90 45 1 Library_Dep 4.28787
91 46 0 File_Dep 4.43253
92 46 1 Library_Dep 4.23642
93 47 0 File_Dep 4.58817
94 47 1 Library_Dep 4.24807
95 48 1 Library_Dep 4.29145
96 48 0 File_Dep 4.48958
97 49 1 Library_Dep 4.23885
98 49 0 File_Dep 4.52195
99 50 0 File_Dep 4.54960
100 50 1 Library_Dep 4.28191
101 51 1 Library_Dep 4.30321
102 51 0 File_Dep 4.41457
103 52 0 File_Dep 4.50377
104 52 1 Library_Dep 4.32230
105 53 1 Library_Dep 4.37291
106 53 0 File_Dep 4.51833
107 54 1 Library_Dep 4.32343
108 54 0 File_Dep 4.41234
109 55 0 File_Dep 4.61027
110 55 1 Library_Dep 4.24206
111 56 0 File_Dep 4.40435
112 56 1 Library_Dep 4.21850
113 57 1 Library_Dep 4.21448
114 57 0 File_Dep 4.55372
115 58 1 Library_Dep 4.25011
116 58 0 File_Dep 4.37978
117 59 0 File_Dep 4.45644
118 59 1 Library_Dep 4.22614
119 60 0 File_Dep 4.48084
120 60 1 Library_Dep 4.30291
121 61 1 Library_Dep 4.25375
122 61 0 File_Dep 4.42744
123 62 1 Library_Dep 4.23224
124 62 0 File_Dep 4.43848
125 63 1 Library_Dep 4.34411
126 63 0 File_Dep 4.37844
127 64 0 File_Dep 4.46045
128 64 1 Library_Dep 4.33635
129 65 1 Library_Dep 4.27824
130 65 0 File_Dep 4.51273
131 66 0 File_Dep 4.48045
132 66 1 Library_Dep 4.21560
133 67 1 Library_Dep 4.29782
134 67 0 File_Dep 4.36148
135 68 1 Library_Dep 4.24973
136 68 0 File_Dep 4.54073
137 69 1 Library_Dep 4.26712
138 69 0 File_Dep 4.42758
139 70 0 File_Dep 4.51497
140 70 1 Library_Dep 4.24806
141 71 0 File_Dep 4.50585
142 71 1 Library_Dep 4.27572
143 72 1 Library_Dep 4.23707
144 72 0 File_Dep 4.45294
145 73 1 Library_Dep 4.21030
146 73 0 File_Dep 4.48946
147 74 1 Library_Dep 4.29487
148 74 0 File_Dep 4.32490
149 75 1 Library_Dep 4.27941
150 75 0 File_Dep 4.44366
151 76 1 Library_Dep 4.21909
152 76 0 File_Dep 4.54426
153 77 0 File_Dep 4.52384
154 77 1 Library_Dep 3.61962
155 78 0 File_Dep 3.67403
156 78 1 Library_Dep 3.88470
157 79 1 Library_Dep 3.76490
158 79 0 File_Dep 3.77329
159 80 1 Library_Dep 3.48084
160 80 0 File_Dep 3.69272
161 81 1 Library_Dep 3.47708
162 81 0 File_Dep 3.72273
163 82 1 Library_Dep 3.51734
164 82 0 File_Dep 3.68759
165 83 1 Library_Dep 3.54831
166 83 0 File_Dep 3.79192
167 84 0 File_Dep 3.71116
168 84 1 Library_Dep 3.46369
169 85 0 File_Dep 3.66128
170 85 1 Library_Dep 3.49283
171 86 0 File_Dep 3.88648
172 86 1 Library_Dep 4.27972
173 87 0 File_Dep 4.30410
174 87 1 Library_Dep 4.16171
175 88 1 Library_Dep 4.51670
176 88 0 File_Dep 4.45813
177 89 0 File_Dep 4.52610
178 89 1 Library_Dep 4.24954
179 90 1 Library_Dep 4.39768
180 90 0 File_Dep 4.50620
181 91 1 Library_Dep 4.26248
182 91 0 File_Dep 4.43065
183 92 0 File_Dep 4.51542
184 92 1 Library_Dep 4.24002
185 93 1 Library_Dep 4.29468
186 93 0 File_Dep 4.55303
187 94 1 Library_Dep 4.27126
188 94 0 File_Dep 4.40811
189 95 1 Library_Dep 4.28352
190 95 0 File_Dep 4.46227
191 96 0 File_Dep 4.47591
192 96 1 Library_Dep 4.03078
193 97 0 File_Dep 4.06101
194 97 1 Library_Dep 3.64187
195 98 0 File_Dep 3.68380
196 98 1 Library_Dep 3.55072
197 99 1 Library_Dep 3.52486
198 99 0 File_Dep 3.69857
199 100 0 File_Dep 3.68426
200 100 1 Library_Dep 3.52606
```

| index | name        | Average (seconds) |
| ----- | ----------- | ----------------- |
| 0     | File_Dep    | 4.387032828       |
| 1     | Library_Dep | 4.184787475       |
|       | diff        | 0.2022453535      |

## Results

```Plain
                sum_sq     df          F        PR(>F)
order         2.831682    1.0  36.323898  8.266298e-09
index         2.023438    1.0  25.956007  8.250686e-07
order:index   0.014726    1.0   0.188906  6.643122e-01
Residual     15.123551  194.0        NaN           NaN
```
