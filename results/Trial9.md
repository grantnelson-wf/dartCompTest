# Trial 9

- Rate:   2.0
- Scalar: 2.0
- Depth:  12
- Group:  15
- Pub 2.13.4

## Procedure

- Changed the main experiment to run using treatments from `addTreatments_FileVsLib_Generated_UpdateLeaves_Dart2js` and made sure the `frac` was `0.01`
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
1 1 1 Library_Dep 11.76173
2 1 0 File_Dep 11.47892
3 2 0 File_Dep 12.48072
4 2 1 Library_Dep 12.10055
5 3 0 File_Dep 12.13168
6 3 1 Library_Dep 12.28765
7 4 0 File_Dep 12.81029
8 4 1 Library_Dep 12.04179
9 5 1 Library_Dep 12.64698
10 5 0 File_Dep 13.09926
11 6 0 File_Dep 12.58002
12 6 1 Library_Dep 13.02040
13 7 0 File_Dep 12.93177
14 7 1 Library_Dep 12.47227
15 8 1 Library_Dep 12.97193
16 8 0 File_Dep 12.38259
17 9 1 Library_Dep 12.26497
18 9 0 File_Dep 12.39472
19 10 1 Library_Dep 13.64802
20 10 0 File_Dep 12.54637
21 11 1 Library_Dep 12.08482
22 11 0 File_Dep 12.89598
23 12 0 File_Dep 12.30626
24 12 1 Library_Dep 13.67736
25 13 1 Library_Dep 12.69290
26 13 0 File_Dep 12.23458
27 14 1 Library_Dep 11.86187
28 14 0 File_Dep 11.19948
29 15 0 File_Dep 11.46685
30 15 1 Library_Dep 11.80803
31 16 1 Library_Dep 12.86180
32 16 0 File_Dep 12.62943
33 17 1 Library_Dep 12.94815
34 17 0 File_Dep 12.44432
35 18 0 File_Dep 12.47464
36 18 1 Library_Dep 12.42042
37 19 0 File_Dep 12.65425
38 19 1 Library_Dep 11.94234
39 20 1 Library_Dep 11.90828
40 20 0 File_Dep 12.78012
41 21 0 File_Dep 12.82522
42 21 1 Library_Dep 11.89962
43 22 0 File_Dep 12.81638
44 22 1 Library_Dep 11.89848
45 23 1 Library_Dep 11.76838
46 23 0 File_Dep 12.46297
47 24 1 Library_Dep 11.86811
48 24 0 File_Dep 11.60544
49 25 1 Library_Dep 12.28696
50 25 0 File_Dep 13.18756
51 26 1 Library_Dep 12.55224
52 26 0 File_Dep 13.12395
53 27 1 Library_Dep 13.14030
54 27 0 File_Dep 12.82809
55 28 0 File_Dep 12.25586
56 28 1 Library_Dep 12.12031
57 29 1 Library_Dep 11.93031
58 29 0 File_Dep 12.03876
59 30 1 Library_Dep 12.16072
60 30 0 File_Dep 12.22156
61 31 0 File_Dep 12.16487
62 31 1 Library_Dep 11.83018
63 32 1 Library_Dep 12.07956
64 32 0 File_Dep 12.15094
65 33 1 Library_Dep 11.95964
66 33 0 File_Dep 12.06391
67 34 0 File_Dep 12.06268
68 34 1 Library_Dep 12.73007
69 35 0 File_Dep 12.17175
70 35 1 Library_Dep 12.03718
71 36 0 File_Dep 12.01481
72 36 1 Library_Dep 12.62542
73 37 1 Library_Dep 14.52859
74 37 0 File_Dep 13.22862
75 38 1 Library_Dep 13.07724
76 38 0 File_Dep 12.74090
77 39 0 File_Dep 13.24251
78 39 1 Library_Dep 12.73890
79 40 1 Library_Dep 12.44958
80 40 0 File_Dep 12.46309
81 41 1 Library_Dep 12.49093
82 41 0 File_Dep 12.57564
83 42 0 File_Dep 12.14037
84 42 1 Library_Dep 12.98512
85 43 1 Library_Dep 12.42542
86 43 0 File_Dep 12.62592
87 44 1 Library_Dep 12.30314
88 44 0 File_Dep 12.54109
89 45 1 Library_Dep 12.39276
90 45 0 File_Dep 12.25174
91 46 1 Library_Dep 12.60339
92 46 0 File_Dep 12.38891
93 47 0 File_Dep 12.46694
94 47 1 Library_Dep 13.78509
95 48 0 File_Dep 12.69836
96 48 1 Library_Dep 12.48831
97 49 1 Library_Dep 12.22411
98 49 0 File_Dep 12.47940
99 50 0 File_Dep 13.01184
100 50 1 Library_Dep 12.99328
101 51 0 File_Dep 12.48494
102 51 1 Library_Dep 13.67867
103 52 1 Library_Dep 13.25497
104 52 0 File_Dep 12.95880
105 53 1 Library_Dep 13.11183
106 53 0 File_Dep 13.37248
107 54 1 Library_Dep 13.12855
108 54 0 File_Dep 13.27681
109 55 0 File_Dep 13.03362
110 55 1 Library_Dep 13.08153
111 56 0 File_Dep 14.29431
112 56 1 Library_Dep 19.16003
113 57 0 File_Dep 15.65498
114 57 1 Library_Dep 15.20968
115 58 1 Library_Dep 14.61851
116 58 0 File_Dep 15.53775
117 59 0 File_Dep 14.71234
118 59 1 Library_Dep 15.12342
119 60 1 Library_Dep 17.51832
120 60 0 File_Dep 14.72065
121 61 0 File_Dep 15.12028
122 61 1 Library_Dep 14.87350
123 62 0 File_Dep 17.99033
124 62 1 Library_Dep 18.93055
125 63 0 File_Dep 16.51911
126 63 1 Library_Dep 17.12308
127 64 0 File_Dep 15.02034
128 64 1 Library_Dep 16.72750
129 65 1 Library_Dep 16.69429
130 65 0 File_Dep 14.73108
131 66 1 Library_Dep 13.43973
132 66 0 File_Dep 14.12097
133 67 1 Library_Dep 14.45257
134 67 0 File_Dep 15.30506
135 68 1 Library_Dep 15.22668
136 68 0 File_Dep 15.12184
137 69 1 Library_Dep 15.02609
138 69 0 File_Dep 15.47697
139 70 0 File_Dep 15.17876
140 70 1 Library_Dep 14.76290
141 71 1 Library_Dep 15.13936
142 71 0 File_Dep 14.97743
143 72 0 File_Dep 15.03148
144 72 1 Library_Dep 14.78947
145 73 1 Library_Dep 14.84681
146 73 0 File_Dep 14.50287
147 74 0 File_Dep 14.42326
148 74 1 Library_Dep 15.01182
149 75 0 File_Dep 14.81733
150 75 1 Library_Dep 15.68146
151 76 1 Library_Dep 18.04358
152 76 0 File_Dep 18.91268
153 77 1 Library_Dep 15.46958
154 77 0 File_Dep 15.68176
155 78 1 Library_Dep 15.21996
156 78 0 File_Dep 15.64551
157 79 1 Library_Dep 17.87410
158 79 0 File_Dep 12.66296
159 80 0 File_Dep 12.78038
160 80 1 Library_Dep 14.00691
161 81 0 File_Dep 13.50045
162 81 1 Library_Dep 13.46498
163 82 1 Library_Dep 13.32574
164 82 0 File_Dep 14.02857
165 83 1 Library_Dep 13.33873
166 83 0 File_Dep 12.81621
167 84 1 Library_Dep 13.53262
168 84 0 File_Dep 14.06206
169 85 0 File_Dep 12.49580
170 85 1 Library_Dep 11.99398
171 86 1 Library_Dep 13.22060
172 86 0 File_Dep 13.33649
173 87 0 File_Dep 13.94548
174 87 1 Library_Dep 13.74115
175 88 1 Library_Dep 13.54364
176 88 0 File_Dep 14.11743
177 89 1 Library_Dep 13.64089
178 89 0 File_Dep 13.87784
179 90 1 Library_Dep 14.25172
180 90 0 File_Dep 13.60620
181 91 1 Library_Dep 13.44946
182 91 0 File_Dep 13.73230
183 92 1 Library_Dep 13.77604
184 92 0 File_Dep 13.54686
185 93 0 File_Dep 13.02160
186 93 1 Library_Dep 13.09105
187 94 1 Library_Dep 13.02481
188 94 0 File_Dep 13.97012
189 95 1 Library_Dep 13.19902
190 95 0 File_Dep 12.60570
191 96 0 File_Dep 13.54585
192 96 1 Library_Dep 13.71084
193 97 0 File_Dep 13.68474
194 97 1 Library_Dep 13.83296
195 98 1 Library_Dep 13.71392
196 98 0 File_Dep 13.51603
197 99 1 Library_Dep 12.38321
198 99 0 File_Dep 13.60786
199 100 0 File_Dep 12.36102
200 100 1 Library_Dep 12.71302
```

## Results

```Plain
                 sum_sq     df          F        PR(>F)
order        103.692577    1.0  60.050063  4.881253e-13
index          1.272163    1.0   0.736730  3.917588e-01
order:index    0.210486    1.0   0.121896  7.273621e-01
Residual     338.446692  196.0        NaN           NaN
```
