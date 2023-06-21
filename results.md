# Results

- orm-benchmark (with no flags)

```
   Reports:

Insert - 200 Times
gorm_prep:      1.032s  1046255 ns/op   5162 B/op       65 allocs/op
zorm:           1.049s  1314706 ns/op   3780 B/op       77 allocs/op
bun:            1.068s  888189 ns/op    5006 B/op       13 allocs/op
godb:           1.093s  964227 ns/op    4508 B/op       116 allocs/op
reform:         1.107s  996307 ns/op    1776 B/op       51 allocs/op
pg:             1.150s  956789 ns/op    797 B/op        10 allocs/op
gorm:           1.151s  1030971 ns/op   7163 B/op       105 allocs/op
pgx_pool:       1.186s  932640 ns/op    288 B/op        10 allocs/op
raw:            1.195s  1182115 ns/op   704 B/op        13 allocs/op
ent:            1.200s  946785 ns/op    4125 B/op       97 allocs/op
gorp:           1.205s  869243 ns/op    1766 B/op       41 allocs/op
sqlboiler:      1.213s  908759 ns/op    1563 B/op       35 allocs/op
pop:            1.220s  1300572 ns/op   9622 B/op       238 allocs/op
sqlx:           1.243s  1034900 ns/op   856 B/op        19 allocs/op
beego:          1.256s  1046496 ns/op   2387 B/op       57 allocs/op
dbr:            1.256s  1101932 ns/op   2688 B/op       65 allocs/op
pgx:            1.259s  1139603 ns/op   269 B/op        10 allocs/op
xorm:           1.279s  1027395 ns/op   3321 B/op       89 allocs/op
rel:            1.285s  1042427 ns/op   2606 B/op       45 allocs/op
sqlc:           1.300s  1024513 ns/op   2770 B/op       61 allocs/op
upper:          1.302s  1071928 ns/op   5877 B/op       126 allocs/op

InsertMulti - 200 Times
pgx_pool:       1.037s  1864418 ns/op   113065 B/op     43 allocs/op
zorm:           1.073s  3626637 ns/op   199929 B/op     2780 allocs/op
rel:            1.097s  2714652 ns/op   306898 B/op     3265 allocs/op
sqlx:           1.126s  3254170 ns/op   171318 B/op     1551 allocs/op
raw:            1.148s  2645783 ns/op   184113 B/op     932 allocs/op
beego:          1.163s  2517590 ns/op   177901 B/op     2745 allocs/op
pg:             1.167s  2701912 ns/op   5922 B/op       112 allocs/op
xorm:           1.190s  3216016 ns/op   248576 B/op     5415 allocs/op
godb:           1.190s  2938851 ns/op   254208 B/op     5894 allocs/op
gorm_prep:      1.208s  2649752 ns/op   251151 B/op     1890 allocs/op
upper:          1.253s  3056806 ns/op   322889 B/op     4204 allocs/op
ent:            1.261s  2718463 ns/op   386594 B/op     4599 allocs/op
bun:            1.312s  2909492 ns/op   42793 B/op      219 allocs/op
pgx:            1.438s  2149259 ns/op   113170 B/op     44 allocs/op
gorm:           1.460s  3678107 ns/op   291534 B/op     5231 allocs/op
reform:         1.909s  3378881 ns/op   459239 B/op     2746 allocs/op
pop:            bulk-insert is not supported
sqlc:           bulk-insert is not supported
sqlboiler:      bulk-insert is not supported
gorp:           bulk-insert is not supported
dbr:            bulk-insert is not supported

Update - 200 Times
gorm_prep:      1.090s  1078542 ns/op   5008 B/op       56 allocs/op
sqlc:           1.103s  83441 ns/op     877 B/op        14 allocs/op
beego:          1.157s  1040203 ns/op   1752 B/op       47 allocs/op
sqlboiler:      1.164s  967229 ns/op    899 B/op        17 allocs/op
pgx_pool:       1.178s  1055144 ns/op   284 B/op        10 allocs/op
gorp:           1.189s  975945 ns/op    1202 B/op       32 allocs/op
rel:            1.192s  1034444 ns/op   3048 B/op       45 allocs/op
pop:            1.194s  1145626 ns/op   6055 B/op       186 allocs/op
gorm:           1.199s  1115010 ns/op   6752 B/op       99 allocs/op
raw:            1.203s  80960 ns/op     749 B/op        13 allocs/op
sqlx:           1.205s  171257 ns/op    872 B/op        20 allocs/op
reform:         1.212s  1154305 ns/op   1767 B/op       51 allocs/op
ent:            1.215s  97139 ns/op     4678 B/op       97 allocs/op
zorm:           1.230s  1224039 ns/op   3024 B/op       59 allocs/op
upper:          1.263s  1543502 ns/op   16669 B/op      390 allocs/op
godb:           1.304s  1130580 ns/op   5116 B/op       154 allocs/op
bun:            1.407s  1374466 ns/op   4731 B/op       5 allocs/op
dbr:            1.422s  1163328 ns/op   2651 B/op       57 allocs/op
pgx:            1.422s  1334380 ns/op   270 B/op        10 allocs/op
xorm:           1.487s  1268467 ns/op   3948 B/op       132 allocs/op
pg:             1.511s  1250709 ns/op   768 B/op        9 allocs/op

Read - 200 Times
bun:            1.001s  100116 ns/op    5822 B/op       39 allocs/op
sqlc:           1.010s  102942 ns/op    2076 B/op       51 allocs/op
rel:            1.018s  101833 ns/op    2304 B/op       47 allocs/op
reform:         1.060s  105992 ns/op    3200 B/op       86 allocs/op
pg:             1.074s  100372 ns/op    872 B/op        20 allocs/op
godb:           1.100s  201215 ns/op    4067 B/op       102 allocs/op
pgx:            1.118s  81923 ns/op     893 B/op        8 allocs/op
xorm:           1.138s  236966 ns/op    4652 B/op       127 allocs/op
sqlx:           1.138s  173269 ns/op    1976 B/op       43 allocs/op
gorp:           1.153s  87443 ns/op     3877 B/op       194 allocs/op
dbr:            1.157s  137446 ns/op    2184 B/op       37 allocs/op
zorm:           1.163s  143814 ns/op    3016 B/op       64 allocs/op
pgx_pool:       1.193s  81918 ns/op     1079 B/op       9 allocs/op
gorm_prep:      1.206s  120648 ns/op    4415 B/op       87 allocs/op
gorm:           1.225s  181121 ns/op    4784 B/op       98 allocs/op
upper:          1.233s  171572 ns/op    5054 B/op       110 allocs/op
ent:            1.247s  86143 ns/op     5605 B/op       144 allocs/op
sqlboiler:      1.272s  119833 ns/op    958 B/op        14 allocs/op
raw:            1.311s  101606 ns/op    2061 B/op       50 allocs/op
pop:            1.323s  86975 ns/op     3195 B/op       67 allocs/op
beego:          1.338s  100541 ns/op    2098 B/op       76 allocs/op

ReadSlice - 200 Times
godb:           1.010s  518344 ns/op    75274 B/op      3084 allocs/op
bun:            1.049s  477320 ns/op    34089 B/op      1124 allocs/op
pg:             1.052s  329015 ns/op    23288 B/op      629 allocs/op
beego:          1.062s  492772 ns/op    55236 B/op      3078 allocs/op
raw:            1.078s  236953 ns/op    38343 B/op      1038 allocs/op
gorm:           1.092s  579993 ns/op    44433 B/op      2191 allocs/op
sqlx:           1.134s  369455 ns/op    37480 B/op      1225 allocs/op
gorm_prep:      1.165s  487295 ns/op    43270 B/op      2081 allocs/op
gorp:           1.198s  355779 ns/op    57424 B/op      1494 allocs/op
dbr:            1.214s  470179 ns/op    30816 B/op      1254 allocs/op
pgx_pool:       1.238s  185385 ns/op    43006 B/op      504 allocs/op
rel:            1.266s  593356 ns/op    100641 B/op     2253 allocs/op
upper:          1.282s  156585 ns/op    4791 B/op       90 allocs/op
zorm:           1.293s  716793 ns/op    161617 B/op     2949 allocs/op
sqlboiler:      1.302s  512286 ns/op    66886 B/op      2259 allocs/op
xorm:           1.314s  592055 ns/op    119476 B/op     4401 allocs/op
sqlc:           1.321s  324102 ns/op    62662 B/op      1150 allocs/op
ent:            1.345s  314630 ns/op    77191 B/op      2035 allocs/op
pop:            1.361s  414690 ns/op    69416 B/op      1307 allocs/op
reform:         1.407s  79721 ns/op     4512 B/op       112 allocs/op
pgx:            1.413s  144453 ns/op    42949 B/op      504 allocs/op
```

- orm-benchmark -multi=5

```
  Reports:

Insert - 1000 Times
sqlboiler:      1.045s  857764 ns/op    1562 B/op       34 allocs/op
gorm_prep:      1.052s  3370204 ns/op   5159 B/op       65 allocs/op
xorm:           1.069s  4050471 ns/op   3317 B/op       88 allocs/op
sqlc:           1.112s  3901431 ns/op   2760 B/op       61 allocs/op
sqlx:           1.172s  1077649 ns/op   856 B/op        19 allocs/op
reform:         1.173s  4146334 ns/op   1782 B/op       51 allocs/op
godb:           1.194s  1098885 ns/op   4511 B/op       116 allocs/op
pg:             1.197s  3885123 ns/op   4452 B/op       10 allocs/op
upper:          1.219s  3819957 ns/op   5886 B/op       125 allocs/op
dbr:            1.237s  952482 ns/op    2688 B/op       65 allocs/op
gorm:           1.240s  1126350 ns/op   7165 B/op       105 allocs/op
pgx_pool:       1.242s  1036527 ns/op   289 B/op        10 allocs/op
ent:            1.245s  4149224 ns/op   4127 B/op       96 allocs/op
pop:            1.250s  1059898 ns/op   9606 B/op       239 allocs/op
gorp:           1.259s  3959730 ns/op   1769 B/op       41 allocs/op
bun:            1.266s  4220592 ns/op   5001 B/op       13 allocs/op
zorm:           1.311s  4200959 ns/op   3776 B/op       77 allocs/op
pgx:            1.315s  1048426 ns/op   272 B/op        10 allocs/op
rel:            1.431s  1168813 ns/op   2606 B/op       45 allocs/op
beego:          2.137s  1840708 ns/op   2387 B/op       57 allocs/op
raw:            2.623s  2091522 ns/op   704 B/op        13 allocs/op

InsertMulti - 1000 Times
upper:          1.084s  6125645 ns/op   322915 B/op     4204 allocs/op
sqlx:           1.085s  3349365 ns/op   171764 B/op     1551 allocs/op
raw:            1.122s  5008898 ns/op   184577 B/op     935 allocs/op
ent:            1.127s  5498100 ns/op   386718 B/op     4602 allocs/op
zorm:           1.128s  6795121 ns/op   199928 B/op     2780 allocs/op
beego:          1.136s  5488898 ns/op   177804 B/op     2745 allocs/op
reform:         1.154s  5391724 ns/op   459194 B/op     2746 allocs/op
xorm:           1.211s  6374622 ns/op   248377 B/op     5414 allocs/op
pgx:            1.226s  1888497 ns/op   113140 B/op     43 allocs/op
gorm:           1.235s  3403051 ns/op   291561 B/op     5231 allocs/op
pgx_pool:       1.259s  1914008 ns/op   113120 B/op     43 allocs/op
pg:             1.265s  6485680 ns/op   3335 B/op       112 allocs/op
bun:            1.270s  6383731 ns/op   42786 B/op      219 allocs/op
godb:           1.392s  3888914 ns/op   254238 B/op     5894 allocs/op
gorm_prep:      1.405s  2681709 ns/op   251159 B/op     1890 allocs/op
rel:            1.436s  3912338 ns/op   306899 B/op     3265 allocs/op
gorp:           bulk-insert is not supported
pop:            bulk-insert is not supported
sqlc:           bulk-insert is not supported
sqlboiler:      bulk-insert is not supported
dbr:            bulk-insert is not supported

Update - 1000 Times
gorm_prep:      1.050s  1136273 ns/op   5008 B/op       56 allocs/op
rel:            1.065s  1171550 ns/op   3048 B/op       45 allocs/op
xorm:           1.067s  4200070 ns/op   3944 B/op       132 allocs/op
zorm:           1.085s  3959055 ns/op   3024 B/op       59 allocs/op
sqlx:           1.118s  172653 ns/op    872 B/op        20 allocs/op
pop:            1.132s  908856 ns/op    6057 B/op       186 allocs/op
dbr:            1.138s  998467 ns/op    2651 B/op       57 allocs/op
pgx:            1.149s  1044999 ns/op   269 B/op        10 allocs/op
raw:            1.161s  81784 ns/op     750 B/op        13 allocs/op
gorp:           1.165s  4001812 ns/op   1204 B/op       32 allocs/op
beego:          1.192s  3858789 ns/op   1752 B/op       47 allocs/op
upper:          1.217s  4680020 ns/op   16627 B/op      390 allocs/op
pg:             1.220s  3973775 ns/op   768 B/op        9 allocs/op
godb:           1.223s  1212247 ns/op   5116 B/op       154 allocs/op
sqlboiler:      1.223s  892969 ns/op    901 B/op        17 allocs/op
ent:            1.229s  97482 ns/op     4677 B/op       97 allocs/op
pgx_pool:       1.240s  1190771 ns/op   287 B/op        10 allocs/op
gorm:           1.245s  958112 ns/op    6752 B/op       99 allocs/op
sqlc:           1.263s  80654 ns/op     878 B/op        14 allocs/op
reform:         1.300s  4277615 ns/op   1772 B/op       51 allocs/op
bun:            3.461s  3471196 ns/op   4731 B/op       5 allocs/op

Read - 1000 Times
ent:            1.022s  83890 ns/op     5605 B/op       144 allocs/op
rel:            1.036s  103604 ns/op    2304 B/op       47 allocs/op
gorm_prep:      1.043s  117495 ns/op    4415 B/op       87 allocs/op
pgx_pool:       1.076s  76656 ns/op     1079 B/op       9 allocs/op
zorm:           1.084s  131776 ns/op    3016 B/op       64 allocs/op
gorm:           1.136s  128924 ns/op    4786 B/op       98 allocs/op
gorp:           1.166s  76256 ns/op     3877 B/op       194 allocs/op
dbr:            1.168s  134148 ns/op    2184 B/op       37 allocs/op
raw:            1.180s  89642 ns/op     2061 B/op       50 allocs/op
sqlx:           1.184s  173689 ns/op    1976 B/op       43 allocs/op
xorm:           1.185s  213653 ns/op    4652 B/op       127 allocs/op
reform:         1.190s  118951 ns/op    3201 B/op       86 allocs/op
bun:            1.209s  100576 ns/op    5967 B/op       48 allocs/op
sqlc:           1.235s  97306 ns/op     2076 B/op       51 allocs/op
sqlboiler:      1.236s  92112 ns/op     968 B/op        14 allocs/op
godb:           1.247s  208399 ns/op    4067 B/op       102 allocs/op
pop:            1.262s  81275 ns/op     3207 B/op       67 allocs/op
beego:          1.286s  104204 ns/op    2098 B/op       76 allocs/op
pgx:            1.310s  82059 ns/op     893 B/op        8 allocs/op
pg:             1.355s  95200 ns/op     1025 B/op       20 allocs/op
upper:          1.425s  230051 ns/op    5055 B/op       110 allocs/op

ReadSlice - 1000 Times
rel:            1.003s  644447 ns/op    100641 B/op     2253 allocs/op
zorm:           1.052s  686509 ns/op    161617 B/op     2949 allocs/op
dbr:            1.067s  422351 ns/op    30816 B/op      1254 allocs/op
ent:            1.077s  268170 ns/op    77187 B/op      2035 allocs/op
raw:            1.122s  274805 ns/op    38340 B/op      1038 allocs/op
gorp:           1.134s  331907 ns/op    57418 B/op      1494 allocs/op
reform:         1.136s  73090 ns/op     4511 B/op       112 allocs/op
xorm:           1.189s  743399 ns/op    119451 B/op     4401 allocs/op
bun:            1.210s  337702 ns/op    34086 B/op      1124 allocs/op
gorm_prep:      1.231s  452017 ns/op    43267 B/op      2081 allocs/op
gorm:           1.249s  450475 ns/op    44444 B/op      2191 allocs/op
pg:             1.249s  324502 ns/op    23972 B/op      629 allocs/op
godb:           1.266s  644039 ns/op    75283 B/op      3084 allocs/op
sqlx:           1.284s  347275 ns/op    37480 B/op      1225 allocs/op
pgx_pool:       1.330s  177762 ns/op    43006 B/op      504 allocs/op
pgx:            1.349s  177873 ns/op    42948 B/op      504 allocs/op
sqlc:           1.361s  320151 ns/op    62661 B/op      1150 allocs/op
upper:          1.368s  197943 ns/op    4791 B/op       90 allocs/op
beego:          1.400s  491728 ns/op    55235 B/op      3078 allocs/op
pop:            1.402s  385881 ns/op    69489 B/op      1307 allocs/op
sqlboiler:      1.418s  496878 ns/op    66981 B/op      2260 allocs/op
```

- orm-benchmark -multi=10

```
  Reports:

Insert - 2000 Times
pgx_pool:       1.066s  4024303 ns/op   286 B/op        10 allocs/op
rel:            1.073s  4097234 ns/op   2598 B/op       44 allocs/op
sqlboiler:      1.121s  3918674 ns/op   1577 B/op       34 allocs/op
dbr:            1.140s  4192268 ns/op   2688 B/op       65 allocs/op
xorm:           1.156s  4013244 ns/op   3320 B/op       88 allocs/op
sqlx:           1.162s  3724477 ns/op   856 B/op        19 allocs/op
gorm:           1.169s  4059171 ns/op   7160 B/op       105 allocs/op
pgx:            1.171s  3741924 ns/op   266 B/op        10 allocs/op
gorp:           1.195s  4192846 ns/op   1777 B/op       41 allocs/op
gorm_prep:      1.197s  4016029 ns/op   5160 B/op       65 allocs/op
godb:           1.202s  4485256 ns/op   4512 B/op       115 allocs/op
beego:          1.204s  3987087 ns/op   2386 B/op       57 allocs/op
bun:            1.215s  3703754 ns/op   5011 B/op       13 allocs/op
pg:             1.224s  3848723 ns/op   4337 B/op       10 allocs/op
pop:            1.226s  4087614 ns/op   9668 B/op       238 allocs/op
ent:            1.228s  3961758 ns/op   4126 B/op       96 allocs/op
zorm:           1.243s  4185223 ns/op   3775 B/op       76 allocs/op
raw:            1.252s  4245376 ns/op   699 B/op        13 allocs/op
reform:         1.263s  4194808 ns/op   1784 B/op       51 allocs/op
upper:          1.264s  4298168 ns/op   5898 B/op       125 allocs/op
sqlc:           1.265s  4159805 ns/op   2766 B/op       61 allocs/op

InsertMulti - 2000 Times
reform:         1.093s  5690397 ns/op   459178 B/op     2746 allocs/op
sqlx:           1.114s  5991905 ns/op   170932 B/op     1551 allocs/op
ent:            1.150s  6084009 ns/op   386693 B/op     4602 allocs/op
gorm:           1.162s  6600824 ns/op   291540 B/op     5231 allocs/op
beego:          1.171s  5578191 ns/op   177951 B/op     2745 allocs/op
zorm:           1.179s  6622564 ns/op   199940 B/op     2780 allocs/op
pg:             1.190s  6195609 ns/op   9182 B/op       112 allocs/op
raw:            1.210s  4800211 ns/op   184470 B/op     934 allocs/op
gorm_prep:      1.214s  5324381 ns/op   251181 B/op     1890 allocs/op
pgx_pool:       1.234s  4996258 ns/op   113079 B/op     43 allocs/op
bun:            1.261s  6062324 ns/op   42838 B/op      219 allocs/op
xorm:           1.268s  5870731 ns/op   248521 B/op     5415 allocs/op
pgx:            1.271s  4573300 ns/op   113142 B/op     43 allocs/op
upper:          1.306s  7508345 ns/op   322895 B/op     4204 allocs/op
rel:            1.332s  6531693 ns/op   306898 B/op     3265 allocs/op
godb:           1.409s  7298795 ns/op   254244 B/op     5894 allocs/op
sqlc:           bulk-insert is not supported
gorp:           bulk-insert is not supported
dbr:            bulk-insert is not supported
pop:            bulk-insert is not supported
sqlboiler:      bulk-insert is not supported

Update - 2000 Times
upper:          1.052s  4477767 ns/op   16636 B/op      390 allocs/op
raw:            1.087s  62915 ns/op     749 B/op        13 allocs/op
ent:            1.098s  109843 ns/op    4678 B/op       97 allocs/op
zorm:           1.119s  4304966 ns/op   3024 B/op       59 allocs/op
sqlx:           1.136s  157918 ns/op    872 B/op        20 allocs/op
pg:             1.146s  4091868 ns/op   768 B/op        9 allocs/op
beego:          1.147s  3771821 ns/op   1752 B/op       47 allocs/op
sqlc:           1.148s  76647 ns/op     878 B/op        14 allocs/op
sqlboiler:      1.154s  3994776 ns/op   900 B/op        17 allocs/op
gorm_prep:      1.175s  4078714 ns/op   5009 B/op       56 allocs/op
pgx:            1.180s  3688926 ns/op   291 B/op        10 allocs/op
gorp:           1.181s  4437980 ns/op   1206 B/op       32 allocs/op
xorm:           1.190s  4281430 ns/op   3944 B/op       132 allocs/op
dbr:            1.191s  4091601 ns/op   2651 B/op       57 allocs/op
pop:            1.202s  4217258 ns/op   6048 B/op       186 allocs/op
rel:            1.215s  4466215 ns/op   3048 B/op       45 allocs/op
bun:            1.249s  4164949 ns/op   4728 B/op       5 allocs/op
reform:         1.277s  4258071 ns/op   1771 B/op       51 allocs/op
gorm:           1.294s  4356899 ns/op   6753 B/op       99 allocs/op
pgx_pool:       1.300s  4691892 ns/op   291 B/op        10 allocs/op
godb:           1.531s  5171850 ns/op   5112 B/op       154 allocs/op

Read - 2000 Times
bun:            1.043s  104323 ns/op    5822 B/op       39 allocs/op
gorm:           1.066s  184617 ns/op    4785 B/op       98 allocs/op
reform:         1.088s  115575 ns/op    3201 B/op       86 allocs/op
sqlboiler:      1.129s  115277 ns/op    945 B/op        14 allocs/op
pg:             1.131s  91987 ns/op     1049 B/op       20 allocs/op
pgx_pool:       1.173s  81427 ns/op     1079 B/op       9 allocs/op
ent:            1.175s  130683 ns/op    5604 B/op       144 allocs/op
gorm_prep:      1.177s  116030 ns/op    4416 B/op       87 allocs/op
sqlc:           1.192s  88379 ns/op     2077 B/op       51 allocs/op
pgx:            1.195s  62694 ns/op     893 B/op        8 allocs/op
zorm:           1.221s  149065 ns/op    3016 B/op       64 allocs/op
rel:            1.230s  103826 ns/op    2304 B/op       47 allocs/op
sqlx:           1.244s  192212 ns/op    1976 B/op       43 allocs/op
gorp:           1.263s  83628 ns/op     3877 B/op       194 allocs/op
upper:          1.275s  211260 ns/op    5055 B/op       110 allocs/op
xorm:           1.275s  235242 ns/op    4652 B/op       127 allocs/op
dbr:            1.278s  133958 ns/op    2184 B/op       37 allocs/op
raw:            1.307s  97534 ns/op     2061 B/op       50 allocs/op
beego:          1.332s  115195 ns/op    2098 B/op       76 allocs/op
godb:           1.415s  234425 ns/op    4068 B/op       102 allocs/op
pop:            1.538s  71117 ns/op     3220 B/op       67 allocs/op

ReadSlice - 2000 Times
xorm:           1.026s  830006 ns/op    119475 B/op     4401 allocs/op
sqlx:           1.042s  375746 ns/op    37480 B/op      1225 allocs/op
sqlboiler:      1.080s  483776 ns/op    67034 B/op      2260 allocs/op
ent:            1.083s  353925 ns/op    77189 B/op      2035 allocs/op
gorp:           1.100s  341504 ns/op    57421 B/op      1494 allocs/op
raw:            1.105s  264395 ns/op    38343 B/op      1038 allocs/op
rel:            1.112s  619503 ns/op    100641 B/op     2253 allocs/op
beego:          1.125s  486620 ns/op    55241 B/op      3078 allocs/op
pgx_pool:       1.128s  175380 ns/op    43007 B/op      504 allocs/op
zorm:           1.168s  688440 ns/op    161618 B/op     2949 allocs/op
bun:            1.176s  434452 ns/op    34086 B/op      1124 allocs/op
sqlc:           1.184s  322787 ns/op    62662 B/op      1150 allocs/op
pgx:            1.184s  127665 ns/op    42949 B/op      504 allocs/op
godb:           1.212s  663957 ns/op    75281 B/op      3084 allocs/op
gorm_prep:      1.215s  472517 ns/op    43279 B/op      2081 allocs/op
pop:            1.229s  364915 ns/op    69572 B/op      1307 allocs/op
dbr:            1.231s  483132 ns/op    30816 B/op      1254 allocs/op
upper:          1.258s  208682 ns/op    4790 B/op       90 allocs/op
pg:             1.261s  399900 ns/op    24346 B/op      629 allocs/op
gorm:           1.435s  609279 ns/op    44427 B/op      2191 allocs/op
reform:         1.515s  86176 ns/op     4512 B/op       112 allocs/op
```

- orm-benchmark -multi=20

```
  Reports:

Insert - 4000 Times
upper:          1.055s  4072131 ns/op   5888 B/op       125 allocs/op
zorm:           1.092s  4182804 ns/op   3774 B/op       76 allocs/op
xorm:           1.100s  4149432 ns/op   3319 B/op       88 allocs/op
pop:            1.129s  4429315 ns/op   9596 B/op       237 allocs/op
sqlc:           1.153s  4132812 ns/op   2766 B/op       61 allocs/op
sqlboiler:      1.172s  889144 ns/op    1563 B/op       34 allocs/op
bun:            1.172s  4156180 ns/op   5006 B/op       13 allocs/op
pgx:            1.184s  3868235 ns/op   266 B/op        10 allocs/op
reform:         1.199s  1067096 ns/op   1776 B/op       51 allocs/op
pg:             1.208s  4068175 ns/op   4587 B/op       10 allocs/op
gorm:           1.214s  3952871 ns/op   7159 B/op       105 allocs/op
dbr:            1.219s  1010051 ns/op   2688 B/op       65 allocs/op
ent:            1.223s  1100968 ns/op   4127 B/op       97 allocs/op
sqlx:           1.224s  1170227 ns/op   856 B/op        19 allocs/op
gorp:           1.230s  927165 ns/op    1770 B/op       42 allocs/op
gorm_prep:      1.243s  1102225 ns/op   5163 B/op       65 allocs/op
pgx_pool:       1.279s  1015473 ns/op   282 B/op        10 allocs/op
godb:           1.281s  4172641 ns/op   4510 B/op       115 allocs/op
beego:          1.286s  3945980 ns/op   2388 B/op       57 allocs/op
rel:            1.333s  1257533 ns/op   2608 B/op       46 allocs/op
raw:            1.347s  1063001 ns/op   704 B/op        13 allocs/op

InsertMulti - 4000 Times
pgx_pool:       1.028s  1831914 ns/op   113102 B/op     43 allocs/op
rel:            1.075s  3411786 ns/op   306899 B/op     3265 allocs/op
gorm_prep:      1.146s  2368462 ns/op   251159 B/op     1890 allocs/op
beego:          1.148s  5886166 ns/op   177869 B/op     2745 allocs/op
sqlx:           1.155s  3200287 ns/op   171723 B/op     1551 allocs/op
bun:            1.158s  6161999 ns/op   42659 B/op      219 allocs/op
zorm:           1.161s  6596927 ns/op   199930 B/op     2780 allocs/op
raw:            1.183s  2151322 ns/op   184038 B/op     931 allocs/op
godb:           1.186s  5898749 ns/op   254227 B/op     5894 allocs/op
pg:             1.200s  6894410 ns/op   9789 B/op       112 allocs/op
xorm:           1.252s  6732289 ns/op   248460 B/op     5414 allocs/op
gorm:           1.254s  6396763 ns/op   291551 B/op     5231 allocs/op
reform:         1.298s  2560795 ns/op   459156 B/op     2746 allocs/op
ent:            1.378s  3055495 ns/op   386606 B/op     4599 allocs/op
upper:          1.387s  7969652 ns/op   322895 B/op     4204 allocs/op
pgx:            1.407s  5410978 ns/op   113139 B/op     43 allocs/op
pop:            bulk-insert is not supported
sqlc:           bulk-insert is not supported
gorp:           bulk-insert is not supported
dbr:            bulk-insert is not supported
sqlboiler:      bulk-insert is not supported

Update - 4000 Times
sqlx:           1.007s  163111 ns/op    872 B/op        20 allocs/op
pgx_pool:       1.012s  1022051 ns/op   279 B/op        10 allocs/op
upper:          1.074s  4947762 ns/op   16707 B/op      390 allocs/op
raw:            1.095s  82024 ns/op     750 B/op        13 allocs/op
zorm:           1.146s  4227833 ns/op   3024 B/op       59 allocs/op
rel:            1.157s  1267730 ns/op   3048 B/op       45 allocs/op
beego:          1.161s  4284934 ns/op   1752 B/op       47 allocs/op
sqlboiler:      1.162s  916038 ns/op    904 B/op        17 allocs/op
dbr:            1.171s  1066380 ns/op   2651 B/op       57 allocs/op
pg:             1.173s  3950513 ns/op   768 B/op        9 allocs/op
pgx:            1.178s  4237794 ns/op   269 B/op        10 allocs/op
sqlc:           1.213s  84154 ns/op     878 B/op        14 allocs/op
bun:            1.213s  4395570 ns/op   4728 B/op       5 allocs/op
pop:            1.221s  4016930 ns/op   6043 B/op       186 allocs/op
gorm:           1.223s  4246672 ns/op   6753 B/op       99 allocs/op
xorm:           1.250s  4253186 ns/op   3945 B/op       132 allocs/op
godb:           1.264s  4014103 ns/op   5112 B/op       154 allocs/op
gorp:           1.272s  979202 ns/op    1207 B/op       32 allocs/op
gorm_prep:      1.287s  1151362 ns/op   5008 B/op       56 allocs/op
ent:            1.292s  84659 ns/op     4678 B/op       97 allocs/op
reform:         1.331s  1242730 ns/op   1775 B/op       51 allocs/op

Read - 4000 Times
rel:            1.014s  109492 ns/op    2304 B/op       47 allocs/op
pop:            1.064s  71092 ns/op     3257 B/op       67 allocs/op
sqlx:           1.077s  171967 ns/op    1976 B/op       43 allocs/op
raw:            1.081s  108065 ns/op    2061 B/op       50 allocs/op
pg:             1.101s  85547 ns/op     1036 B/op       20 allocs/op
sqlc:           1.103s  84719 ns/op     2077 B/op       51 allocs/op
dbr:            1.105s  129811 ns/op    2184 B/op       37 allocs/op
pgx:            1.138s  78650 ns/op     893 B/op        8 allocs/op
gorp:           1.144s  79533 ns/op     3877 B/op       194 allocs/op
gorm:           1.172s  170852 ns/op    4785 B/op       98 allocs/op
xorm:           1.179s  248488 ns/op    4652 B/op       127 allocs/op
gorm_prep:      1.180s  119029 ns/op    4415 B/op       87 allocs/op
beego:          1.184s  100935 ns/op    2098 B/op       76 allocs/op
godb:           1.210s  201072 ns/op    4067 B/op       102 allocs/op
bun:            1.211s  98861 ns/op     5967 B/op       48 allocs/op
reform:         1.216s  100602 ns/op    3200 B/op       86 allocs/op
ent:            1.247s  94846 ns/op     5605 B/op       144 allocs/op
sqlboiler:      1.249s  94680 ns/op     952 B/op        14 allocs/op
upper:          1.303s  237077 ns/op    5054 B/op       110 allocs/op
pgx_pool:       1.313s  83616 ns/op     1079 B/op       9 allocs/op
zorm:           1.349s  148562 ns/op    3016 B/op       64 allocs/op

ReadSlice - 4000 Times
ent:            1.012s  287305 ns/op    77189 B/op      2035 allocs/op
zorm:           1.053s  652563 ns/op    161618 B/op     2949 allocs/op
pg:             1.084s  297692 ns/op    23720 B/op      629 allocs/op
sqlx:           1.150s  334439 ns/op    37480 B/op      1225 allocs/op
upper:          1.159s  215784 ns/op    4791 B/op       90 allocs/op
bun:            1.162s  429901 ns/op    34088 B/op      1124 allocs/op
pgx_pool:       1.176s  169150 ns/op    43006 B/op      504 allocs/op
sqlboiler:      1.207s  487534 ns/op    66852 B/op      2259 allocs/op
pgx:            1.209s  184935 ns/op    42949 B/op      504 allocs/op
gorm:           1.238s  561567 ns/op    44434 B/op      2191 allocs/op
raw:            1.251s  249294 ns/op    38342 B/op      1038 allocs/op
gorp:           1.261s  404763 ns/op    57414 B/op      1494 allocs/op
pop:            1.265s  284646 ns/op    69517 B/op      1307 allocs/op
gorm_prep:      1.267s  507897 ns/op    43266 B/op      2081 allocs/op
sqlc:           1.289s  316556 ns/op    62662 B/op      1150 allocs/op
xorm:           1.369s  800564 ns/op    119464 B/op     4401 allocs/op
beego:          1.370s  459080 ns/op    55238 B/op      3078 allocs/op
dbr:            1.411s  468380 ns/op    30816 B/op      1254 allocs/op
godb:           1.413s  653971 ns/op    75279 B/op      3084 allocs/op
reform:         1.425s  80897 ns/op     4511 B/op       112 allocs/op
rel:            1.473s  685741 ns/op    100641 B/op     2253 allocs/op
```
