# Results

- orm-benchmark (with no flags)

```
   
Reports:

   200 times - Insert
  pgx_pool:     0.05s       245074 ns/op     265 B/op     10 allocs/op
     beego:     0.05s       256990 ns/op    2375 B/op     56 allocs/op
       pgx:     0.05s       265282 ns/op     244 B/op     10 allocs/op
      sqlc:     0.06s       278989 ns/op    3184 B/op     64 allocs/op
    reform:     0.06s       286501 ns/op    2091 B/op     53 allocs/op
 sqlboiler:     0.06s       291875 ns/op    1874 B/op     37 allocs/op
       raw:     0.06s       293506 ns/op    1007 B/op     15 allocs/op
      gorm:     0.06s       299928 ns/op    7182 B/op     90 allocs/op
       ent:     0.06s       310605 ns/op    4172 B/op     99 allocs/op
      gorp:     0.07s       325421 ns/op    2073 B/op     43 allocs/op
       dbr:     0.07s       343122 ns/op    2796 B/op     66 allocs/op
 gorm_prep:     0.07s       358492 ns/op    5755 B/op     70 allocs/op
        pg:     0.08s       410241 ns/op    6275 B/op     11 allocs/op
       bun:     0.08s       419736 ns/op    5290 B/op     14 allocs/op
       rel:     0.08s       419945 ns/op    2511 B/op     41 allocs/op
      godb:     0.09s       457772 ns/op    4706 B/op    116 allocs/op
      xorm:     0.09s       470192 ns/op    3440 B/op     90 allocs/op
      sqlx:     0.10s       485284 ns/op     856 B/op     19 allocs/op
       pop:     0.14s       691463 ns/op   10395 B/op    250 allocs/op
     upper:     0.62s      3091853 ns/op   14198 B/op    685 allocs/op

   200 times - MultiInsert 100 row
       pgx:     0.20s       986154 ns/op  114312 B/op     42 allocs/op
  pgx_pool:     0.20s       988054 ns/op  114397 B/op     43 allocs/op
       raw:     0.25s      1246689 ns/op  192353 B/op    938 allocs/op
    reform:     0.25s      1260349 ns/op  466428 B/op   2747 allocs/op
     beego:     0.26s      1324250 ns/op  179489 B/op   2746 allocs/op
 gorm_prep:     0.29s      1461352 ns/op  235339 B/op   2281 allocs/op
       ent:     0.30s      1505309 ns/op  412424 B/op   4902 allocs/op
       bun:     0.32s      1579157 ns/op   42361 B/op    217 allocs/op
        pg:     0.32s      1594911 ns/op    3316 B/op    112 allocs/op
      gorm:     0.36s      1780589 ns/op  272532 B/op   3827 allocs/op
      sqlx:     0.37s      1849694 ns/op  171538 B/op   1552 allocs/op
       rel:     0.40s      1985415 ns/op  303976 B/op   3261 allocs/op
      xorm:     0.42s      2075169 ns/op  254783 B/op   5416 allocs/op
      godb:     0.44s      2206863 ns/op  259949 B/op   5893 allocs/op
     upper:     0.55s      2756331 ns/op  547001 B/op  19610 allocs/op
       dbr:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert

   200 times - Update
       raw:     0.02s       124649 ns/op     721 B/op     13 allocs/op
       ent:     0.03s       128689 ns/op    4529 B/op     99 allocs/op
    reform:     0.05s       235000 ns/op    1747 B/op     51 allocs/op
 sqlboiler:     0.05s       242443 ns/op     924 B/op     17 allocs/op
      sqlc:     0.05s       248313 ns/op     891 B/op     14 allocs/op
  pgx_pool:     0.05s       253656 ns/op     302 B/op     10 allocs/op
       pgx:     0.05s       258459 ns/op     282 B/op     10 allocs/op
      sqlx:     0.05s       262839 ns/op     872 B/op     20 allocs/op
 gorm_prep:     0.06s       275227 ns/op    5132 B/op     59 allocs/op
     beego:     0.06s       277551 ns/op    1755 B/op     47 allocs/op
        pg:     0.06s       313877 ns/op     768 B/op      9 allocs/op
       bun:     0.07s       328530 ns/op    4728 B/op      5 allocs/op
       pop:     0.07s       328790 ns/op    6555 B/op    198 allocs/op
       dbr:     0.07s       330743 ns/op    2651 B/op     57 allocs/op
      gorm:     0.08s       412241 ns/op    6625 B/op     81 allocs/op
      xorm:     0.09s       464885 ns/op    3648 B/op    126 allocs/op
       rel:     0.10s       476978 ns/op    2531 B/op     41 allocs/op
      godb:     0.11s       534493 ns/op    5161 B/op    154 allocs/op
      gorp:     0.15s       733466 ns/op    1223 B/op     32 allocs/op
     upper:     0.21s      1072052 ns/op   33713 B/op   1510 allocs/op

   200 times - Read
       pgx:     0.02s       110430 ns/op     886 B/op      8 allocs/op
  pgx_pool:     0.02s       121668 ns/op    1078 B/op      9 allocs/op
      sqlc:     0.02s       124824 ns/op    2162 B/op     52 allocs/op
       raw:     0.03s       125597 ns/op    2098 B/op     50 allocs/op
     beego:     0.03s       126245 ns/op    2093 B/op     75 allocs/op
    reform:     0.03s       133248 ns/op    3232 B/op     86 allocs/op
       pop:     0.03s       137541 ns/op    3914 B/op     71 allocs/op
      gorp:     0.03s       138027 ns/op    3869 B/op    194 allocs/op
 gorm_prep:     0.03s       142732 ns/op    4800 B/op     84 allocs/op
        pg:     0.03s       147607 ns/op     873 B/op     20 allocs/op
       bun:     0.03s       150374 ns/op    5490 B/op     21 allocs/op
       ent:     0.03s       154386 ns/op    5424 B/op    149 allocs/op
       rel:     0.03s       154429 ns/op    1803 B/op     45 allocs/op
       dbr:     0.03s       155212 ns/op    2184 B/op     37 allocs/op
 sqlboiler:     0.03s       160117 ns/op    1302 B/op     14 allocs/op
      gorm:     0.04s       193500 ns/op    5157 B/op     95 allocs/op
      sqlx:     0.05s       269626 ns/op    1976 B/op     43 allocs/op
      godb:     0.06s       282277 ns/op    4112 B/op    102 allocs/op
      xorm:     0.06s       294818 ns/op    4615 B/op    125 allocs/op
     upper:     0.06s       297165 ns/op    8126 B/op    315 allocs/op

   200 times - MultiRead limit 100
    reform:     0.03s       125809 ns/op    2955 B/op     74 allocs/op
       pgx:     0.05s       234406 ns/op   42919 B/op    504 allocs/op
  pgx_pool:     0.05s       242220 ns/op   42993 B/op    504 allocs/op
     upper:     0.06s       299478 ns/op    7895 B/op    294 allocs/op
       raw:     0.07s       334488 ns/op   38352 B/op   1038 allocs/op
      sqlc:     0.07s       343411 ns/op   73129 B/op   1251 allocs/op
        pg:     0.07s       351259 ns/op   22258 B/op    629 allocs/op
 sqlboiler:     0.08s       398217 ns/op   58031 B/op   1259 allocs/op
      sqlx:     0.08s       399482 ns/op   39064 B/op   1516 allocs/op
       ent:     0.08s       401262 ns/op   76035 B/op   2039 allocs/op
       dbr:     0.09s       441272 ns/op   32416 B/op   1545 allocs/op
      gorp:     0.09s       456616 ns/op   57383 B/op   1494 allocs/op
       bun:     0.09s       473083 ns/op   33999 B/op   1124 allocs/op
     beego:     0.10s       476920 ns/op   55187 B/op   3077 allocs/op
       pop:     0.10s       483879 ns/op   75950 B/op   1511 allocs/op
 gorm_prep:     0.11s       561035 ns/op   56077 B/op   2078 allocs/op
      gorm:     0.12s       606233 ns/op   57202 B/op   2378 allocs/op
       rel:     0.13s       641483 ns/op   95321 B/op   2248 allocs/op
      godb:     0.14s       700470 ns/op   75263 B/op   3084 allocs/op
      xorm:     0.15s       774524 ns/op  119973 B/op   4687 allocs/op
```

- orm-benchmark -multi=5

```
  
Reports:

  1000 times - Insert
       raw:     0.23s       232371 ns/op     763 B/op     13 allocs/op
      gorp:     0.24s       239670 ns/op    1828 B/op     42 allocs/op
 sqlboiler:     0.25s       254119 ns/op    1622 B/op     35 allocs/op
        pg:     0.27s       267467 ns/op    1889 B/op     10 allocs/op
 gorm_prep:     0.27s       274103 ns/op    5357 B/op     68 allocs/op
       pgx:     0.30s       297665 ns/op     265 B/op     10 allocs/op
     beego:     0.32s       319199 ns/op    2375 B/op     56 allocs/op
      gorm:     0.33s       327029 ns/op    6974 B/op     90 allocs/op
      sqlc:     0.34s       341302 ns/op    2928 B/op     63 allocs/op
  pgx_pool:     0.35s       351228 ns/op     281 B/op     10 allocs/op
    reform:     0.36s       358293 ns/op    1837 B/op     51 allocs/op
       bun:     0.39s       387634 ns/op    5053 B/op     14 allocs/op
       dbr:     0.40s       401741 ns/op    2709 B/op     65 allocs/op
      godb:     0.40s       402019 ns/op    4575 B/op    116 allocs/op
     upper:     0.47s       465916 ns/op   13841 B/op    675 allocs/op
      sqlx:     0.47s       471531 ns/op     856 B/op     19 allocs/op
       ent:     0.48s       476309 ns/op    4214 B/op    100 allocs/op
      xorm:     0.49s       485804 ns/op    3343 B/op     89 allocs/op
       rel:     0.52s       523686 ns/op    2504 B/op     42 allocs/op
       pop:     0.68s       682291 ns/op   10146 B/op    249 allocs/op

  1000 times - MultiInsert 100 row
       pgx:     0.97s       972880 ns/op  114279 B/op     42 allocs/op
  pgx_pool:     1.06s      1060093 ns/op  114316 B/op     43 allocs/op
       raw:     1.07s      1066369 ns/op  191429 B/op    932 allocs/op
     beego:     1.18s      1183030 ns/op  179421 B/op   2746 allocs/op
    reform:     1.29s      1288599 ns/op  466306 B/op   2747 allocs/op
 gorm_prep:     1.38s      1376536 ns/op  235262 B/op   2283 allocs/op
       ent:     1.52s      1524252 ns/op  412165 B/op   4900 allocs/op
        pg:     1.55s      1545085 ns/op    4437 B/op    112 allocs/op
       bun:     1.67s      1666641 ns/op   42256 B/op    218 allocs/op
      sqlx:     1.75s      1751929 ns/op  171228 B/op   1552 allocs/op
      gorm:     1.84s      1844521 ns/op  272896 B/op   3829 allocs/op
      xorm:     2.01s      2010513 ns/op  254782 B/op   5416 allocs/op
       rel:     2.09s      2092650 ns/op  303992 B/op   3263 allocs/op
      godb:     2.17s      2168776 ns/op  259958 B/op   5894 allocs/op
     upper:     2.80s      2800796 ns/op  545451 B/op  19495 allocs/op
       pop:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert

  1000 times - Update
       raw:     0.11s       112503 ns/op     752 B/op     13 allocs/op
       ent:     0.13s       132288 ns/op    4561 B/op     99 allocs/op
       pgx:     0.22s       220180 ns/op     273 B/op     10 allocs/op
 sqlboiler:     0.24s       238785 ns/op     907 B/op     17 allocs/op
      sqlx:     0.26s       256295 ns/op     872 B/op     20 allocs/op
      gorp:     0.28s       279293 ns/op    1209 B/op     32 allocs/op
    reform:     0.30s       296641 ns/op    1777 B/op     51 allocs/op
 gorm_prep:     0.31s       309167 ns/op    5113 B/op     59 allocs/op
        pg:     0.32s       315328 ns/op     768 B/op      9 allocs/op
     beego:     0.32s       318470 ns/op    1752 B/op     47 allocs/op
      sqlc:     0.33s       326662 ns/op     880 B/op     14 allocs/op
       pop:     0.33s       330475 ns/op    6586 B/op    198 allocs/op
      gorm:     0.34s       344776 ns/op    6600 B/op     81 allocs/op
      godb:     0.42s       420720 ns/op    5160 B/op    154 allocs/op
       bun:     0.43s       428672 ns/op    4728 B/op      5 allocs/op
       dbr:     0.49s       492431 ns/op    2651 B/op     57 allocs/op
  pgx_pool:     0.51s       508551 ns/op     287 B/op     10 allocs/op
       rel:     0.56s       555202 ns/op    2528 B/op     41 allocs/op
      xorm:     0.73s       728055 ns/op    3649 B/op    126 allocs/op
     upper:     1.09s      1087893 ns/op   33515 B/op   1503 allocs/op

  1000 times - Read
       pgx:     0.11s       112759 ns/op     889 B/op      8 allocs/op
  pgx_pool:     0.12s       116308 ns/op    1076 B/op      9 allocs/op
      sqlc:     0.12s       120599 ns/op    2176 B/op     52 allocs/op
       raw:     0.13s       127287 ns/op    2057 B/op     50 allocs/op
    reform:     0.13s       127866 ns/op    3193 B/op     86 allocs/op
     beego:     0.13s       130152 ns/op    2089 B/op     75 allocs/op
       pop:     0.14s       136527 ns/op    3561 B/op     71 allocs/op
      gorp:     0.14s       137286 ns/op    3881 B/op    194 allocs/op
 gorm_prep:     0.14s       140273 ns/op    4807 B/op     84 allocs/op
        pg:     0.14s       144615 ns/op     872 B/op     20 allocs/op
       ent:     0.15s       147878 ns/op    5401 B/op    149 allocs/op
 sqlboiler:     0.15s       150328 ns/op     971 B/op     14 allocs/op
       bun:     0.15s       151230 ns/op    5489 B/op     21 allocs/op
       rel:     0.15s       153123 ns/op    1800 B/op     45 allocs/op
       dbr:     0.17s       167845 ns/op    2184 B/op     37 allocs/op
      gorm:     0.19s       192505 ns/op    5156 B/op     95 allocs/op
      sqlx:     0.28s       275558 ns/op    1976 B/op     43 allocs/op
      godb:     0.29s       290282 ns/op    4112 B/op    102 allocs/op
      xorm:     0.29s       293656 ns/op    4616 B/op    125 allocs/op
     upper:     0.30s       302490 ns/op    8136 B/op    315 allocs/op

  1000 times - MultiRead limit 100
    reform:     0.13s       130429 ns/op    2921 B/op     74 allocs/op
  pgx_pool:     0.23s       230688 ns/op   43009 B/op    504 allocs/op
       pgx:     0.23s       231621 ns/op   42949 B/op    504 allocs/op
     upper:     0.29s       294645 ns/op    7849 B/op    294 allocs/op
       raw:     0.30s       302510 ns/op   38342 B/op   1038 allocs/op
      sqlc:     0.35s       348298 ns/op   73158 B/op   1251 allocs/op
        pg:     0.35s       351714 ns/op   22259 B/op    629 allocs/op
      gorp:     0.39s       392158 ns/op   57371 B/op   1494 allocs/op
 sqlboiler:     0.40s       397040 ns/op   57895 B/op   1259 allocs/op
       ent:     0.40s       397313 ns/op   76049 B/op   2039 allocs/op
      sqlx:     0.40s       402028 ns/op   39064 B/op   1516 allocs/op
       dbr:     0.44s       440409 ns/op   32416 B/op   1545 allocs/op
       bun:     0.45s       449950 ns/op   33993 B/op   1124 allocs/op
     beego:     0.46s       462553 ns/op   55183 B/op   3077 allocs/op
       pop:     0.48s       482677 ns/op   75774 B/op   1511 allocs/op
 gorm_prep:     0.55s       545944 ns/op   56065 B/op   2078 allocs/op
      gorm:     0.60s       600034 ns/op   57175 B/op   2378 allocs/op
       rel:     0.67s       671510 ns/op   95322 B/op   2248 allocs/op
      godb:     0.70s       700961 ns/op   75263 B/op   3084 allocs/op
      xorm:     0.77s       767049 ns/op  119958 B/op   4687 allocs/op
```

- orm-benchmark -multi=10

```
  
Reports:

  2000 times - Insert
       raw:     0.47s       232552 ns/op     734 B/op     13 allocs/op
    reform:     0.50s       251327 ns/op    1807 B/op     51 allocs/op
       pgx:     0.51s       256782 ns/op     269 B/op     10 allocs/op
      gorp:     0.52s       259791 ns/op    1798 B/op     42 allocs/op
  pgx_pool:     0.55s       275199 ns/op     285 B/op     10 allocs/op
       ent:     0.56s       280591 ns/op    4215 B/op    100 allocs/op
     beego:     0.56s       281193 ns/op    2376 B/op     56 allocs/op
       bun:     0.59s       297469 ns/op    5022 B/op     14 allocs/op
 gorm_prep:     0.61s       302734 ns/op    5285 B/op     68 allocs/op
       dbr:     0.61s       306889 ns/op    2698 B/op     65 allocs/op
 sqlboiler:     0.64s       319640 ns/op    1592 B/op     35 allocs/op
      sqlc:     0.66s       329046 ns/op    2900 B/op     63 allocs/op
        pg:     0.70s       349487 ns/op     817 B/op     10 allocs/op
      gorm:     0.71s       354485 ns/op    6956 B/op     90 allocs/op
      sqlx:     0.73s       362817 ns/op     856 B/op     19 allocs/op
      godb:     0.84s       418691 ns/op    4564 B/op    116 allocs/op
      xorm:     0.85s       425751 ns/op    3335 B/op     89 allocs/op
     upper:     0.91s       453191 ns/op   13798 B/op    674 allocs/op
       rel:     1.03s       512537 ns/op    2504 B/op     42 allocs/op
       pop:     1.20s       600456 ns/op   10094 B/op    249 allocs/op

  2000 times - MultiInsert 100 row
  pgx_pool:     1.96s       979228 ns/op  114302 B/op     42 allocs/op
       pgx:     1.96s       979859 ns/op  114296 B/op     43 allocs/op
       raw:     2.16s      1081653 ns/op  191315 B/op    931 allocs/op
     beego:     2.35s      1173084 ns/op  179412 B/op   2746 allocs/op
    reform:     2.49s      1244495 ns/op  466332 B/op   2747 allocs/op
 gorm_prep:     2.64s      1317698 ns/op  235249 B/op   2283 allocs/op
       ent:     2.92s      1461360 ns/op  412139 B/op   4900 allocs/op
        pg:     3.27s      1636814 ns/op    3876 B/op    112 allocs/op
       bun:     3.32s      1662025 ns/op   42245 B/op    218 allocs/op
      sqlx:     3.51s      1753239 ns/op  171221 B/op   1552 allocs/op
      gorm:     3.76s      1877963 ns/op  272990 B/op   3829 allocs/op
      xorm:     4.01s      2007311 ns/op  254759 B/op   5416 allocs/op
       rel:     4.04s      2018683 ns/op  303994 B/op   3263 allocs/op
      godb:     4.31s      2155542 ns/op  259964 B/op   5895 allocs/op
     upper:     5.56s      2780583 ns/op  545254 B/op  19480 allocs/op
 sqlboiler:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert

  2000 times - Update
       raw:     0.24s       117654 ns/op     748 B/op     13 allocs/op
       ent:     0.26s       131200 ns/op    4556 B/op     99 allocs/op
      gorp:     0.48s       237989 ns/op    1205 B/op     32 allocs/op
    reform:     0.49s       244612 ns/op    1772 B/op     51 allocs/op
  pgx_pool:     0.51s       253959 ns/op     285 B/op     10 allocs/op
     beego:     0.51s       254801 ns/op    1752 B/op     47 allocs/op
       pgx:     0.54s       269085 ns/op     268 B/op     10 allocs/op
      sqlx:     0.55s       275285 ns/op     872 B/op     20 allocs/op
 gorm_prep:     0.57s       282884 ns/op    5121 B/op     59 allocs/op
 sqlboiler:     0.59s       296831 ns/op     901 B/op     17 allocs/op
       bun:     0.60s       298133 ns/op    4728 B/op      5 allocs/op
      sqlc:     0.61s       305460 ns/op     876 B/op     14 allocs/op
       dbr:     0.63s       316067 ns/op    2651 B/op     57 allocs/op
        pg:     0.71s       353794 ns/op     768 B/op      9 allocs/op
       pop:     0.71s       354609 ns/op    6582 B/op    198 allocs/op
      gorm:     0.80s       398948 ns/op    6604 B/op     81 allocs/op
      godb:     0.84s       421240 ns/op    5160 B/op    154 allocs/op
       rel:     0.95s       477004 ns/op    2528 B/op     41 allocs/op
      xorm:     1.09s       543835 ns/op    3648 B/op    126 allocs/op
     upper:     2.24s      1118376 ns/op   33495 B/op   1502 allocs/op

  2000 times - Read
       pgx:     0.23s       114139 ns/op     892 B/op      8 allocs/op
  pgx_pool:     0.23s       117247 ns/op    1080 B/op      9 allocs/op
       raw:     0.24s       119898 ns/op    2065 B/op     50 allocs/op
      sqlc:     0.24s       120391 ns/op    2172 B/op     52 allocs/op
     beego:     0.25s       126036 ns/op    2088 B/op     75 allocs/op
    reform:     0.26s       128793 ns/op    3197 B/op     86 allocs/op
      gorp:     0.27s       135568 ns/op    3876 B/op    194 allocs/op
       pop:     0.27s       136235 ns/op    3529 B/op     71 allocs/op
 gorm_prep:     0.28s       140668 ns/op    4801 B/op     84 allocs/op
        pg:     0.29s       144767 ns/op     872 B/op     20 allocs/op
 sqlboiler:     0.30s       147577 ns/op     975 B/op     14 allocs/op
       ent:     0.30s       148545 ns/op    5405 B/op    149 allocs/op
       bun:     0.30s       151071 ns/op    5489 B/op     21 allocs/op
       rel:     0.31s       153919 ns/op    1800 B/op     45 allocs/op
       dbr:     0.31s       154526 ns/op    2184 B/op     37 allocs/op
      gorm:     0.38s       192318 ns/op    5151 B/op     95 allocs/op
      sqlx:     0.55s       274827 ns/op    1976 B/op     43 allocs/op
      godb:     0.57s       286024 ns/op    4112 B/op    102 allocs/op
     upper:     0.61s       303265 ns/op    8135 B/op    315 allocs/op
      xorm:     0.63s       315924 ns/op    4616 B/op    125 allocs/op

  2000 times - MultiRead limit 100
    reform:     0.25s       123007 ns/op    2924 B/op     74 allocs/op
  pgx_pool:     0.46s       229694 ns/op   43008 B/op    504 allocs/op
       pgx:     0.47s       233466 ns/op   42949 B/op    504 allocs/op
     upper:     0.60s       300142 ns/op    7841 B/op    294 allocs/op
       raw:     0.62s       307586 ns/op   38342 B/op   1038 allocs/op
      sqlc:     0.70s       349380 ns/op   73158 B/op   1251 allocs/op
        pg:     0.71s       356888 ns/op   22259 B/op    629 allocs/op
 sqlboiler:     0.78s       389179 ns/op   57824 B/op   1259 allocs/op
      gorp:     0.80s       399941 ns/op   57372 B/op   1494 allocs/op
      sqlx:     0.80s       400438 ns/op   39064 B/op   1516 allocs/op
       ent:     0.80s       400567 ns/op   76049 B/op   2039 allocs/op
       bun:     0.90s       450218 ns/op   33993 B/op   1124 allocs/op
       dbr:     0.91s       452802 ns/op   32416 B/op   1545 allocs/op
     beego:     0.92s       459299 ns/op   55182 B/op   3077 allocs/op
       pop:     0.95s       477044 ns/op   75752 B/op   1510 allocs/op
 gorm_prep:     1.11s       554050 ns/op   56068 B/op   2078 allocs/op
      gorm:     1.26s       629364 ns/op   57171 B/op   2378 allocs/op
       rel:     1.32s       662341 ns/op   95322 B/op   2248 allocs/op
      godb:     1.40s       700660 ns/op   75264 B/op   3084 allocs/op
      xorm:     1.52s       761857 ns/op  119955 B/op   4687 allocs/op
```

- orm-benchmark -multi=20

```
  
Reports:

  4000 times - Insert
  pgx_pool:     0.85s       212302 ns/op     286 B/op     10 allocs/op
       raw:     0.92s       228881 ns/op     719 B/op     13 allocs/op
       ent:     0.96s       239200 ns/op    4216 B/op    100 allocs/op
      sqlc:     1.02s       256193 ns/op    2884 B/op     63 allocs/op
     beego:     1.06s       266070 ns/op    2375 B/op     56 allocs/op
 gorm_prep:     1.07s       266722 ns/op    5272 B/op     68 allocs/op
       dbr:     1.07s       268403 ns/op    2693 B/op     65 allocs/op
       bun:     1.15s       286370 ns/op    4998 B/op     14 allocs/op
       pgx:     1.16s       289026 ns/op     270 B/op     10 allocs/op
 sqlboiler:     1.22s       306151 ns/op    1574 B/op     35 allocs/op
      gorp:     1.23s       307011 ns/op    1782 B/op     42 allocs/op
      gorm:     1.27s       316616 ns/op    6949 B/op     90 allocs/op
        pg:     1.33s       331482 ns/op    1067 B/op     10 allocs/op
    reform:     1.39s       346603 ns/op    1790 B/op     51 allocs/op
      sqlx:     1.74s       433908 ns/op     856 B/op     19 allocs/op
      godb:     1.80s       451118 ns/op    4558 B/op    116 allocs/op
      xorm:     1.90s       475966 ns/op    3326 B/op     89 allocs/op
       rel:     1.97s       493178 ns/op    2504 B/op     42 allocs/op
     upper:     2.21s       552695 ns/op   13775 B/op    673 allocs/op
       pop:     2.33s       581785 ns/op   10079 B/op    249 allocs/op

  4000 times - MultiInsert 100 row
  pgx_pool:     3.71s       927203 ns/op  114309 B/op     43 allocs/op
       pgx:     3.78s       945373 ns/op  114290 B/op     43 allocs/op
       raw:     4.13s      1033308 ns/op  191258 B/op    931 allocs/op
     beego:     4.41s      1103580 ns/op  179425 B/op   2746 allocs/op
    reform:     4.87s      1218716 ns/op  466279 B/op   2747 allocs/op
 gorm_prep:     4.96s      1240876 ns/op  235238 B/op   2283 allocs/op
       ent:     5.66s      1414173 ns/op  412127 B/op   4900 allocs/op
        pg:     6.06s      1515313 ns/op    3857 B/op    112 allocs/op
       bun:     6.33s      1583501 ns/op   42242 B/op    219 allocs/op
      sqlx:     6.71s      1676436 ns/op  171169 B/op   1552 allocs/op
      gorm:     7.25s      1812285 ns/op  273045 B/op   3829 allocs/op
      xorm:     8.10s      2025792 ns/op  254767 B/op   5416 allocs/op
       rel:     8.13s      2033095 ns/op  303995 B/op   3263 allocs/op
      godb:     8.51s      2128653 ns/op  259967 B/op   5895 allocs/op
     upper:    11.08s      2770590 ns/op  545161 B/op  19473 allocs/op
 sqlboiler:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert

  4000 times - Update
       raw:     0.47s       116522 ns/op     750 B/op     13 allocs/op
       ent:     0.54s       133754 ns/op    4558 B/op     99 allocs/op
  pgx_pool:     0.94s       234987 ns/op     286 B/op     10 allocs/op
      sqlc:     1.02s       255754 ns/op     878 B/op     14 allocs/op
     beego:     1.03s       256764 ns/op    1752 B/op     47 allocs/op
      sqlx:     1.05s       262981 ns/op     872 B/op     20 allocs/op
       dbr:     1.06s       264428 ns/op    2651 B/op     57 allocs/op
       pop:     1.10s       273920 ns/op    6582 B/op    198 allocs/op
 sqlboiler:     1.11s       277650 ns/op     901 B/op     17 allocs/op
        pg:     1.16s       289809 ns/op     768 B/op      9 allocs/op
       bun:     1.22s       304865 ns/op    4728 B/op      5 allocs/op
       pgx:     1.28s       320898 ns/op     270 B/op     10 allocs/op
      gorp:     1.31s       326643 ns/op    1204 B/op     32 allocs/op
    reform:     1.34s       334255 ns/op    1774 B/op     51 allocs/op
      gorm:     1.37s       343641 ns/op    6604 B/op     81 allocs/op
 gorm_prep:     1.38s       344458 ns/op    5116 B/op     59 allocs/op
      godb:     1.76s       440759 ns/op    5160 B/op    154 allocs/op
      xorm:     1.78s       444703 ns/op    3648 B/op    126 allocs/op
       rel:     2.00s       499036 ns/op    2528 B/op     41 allocs/op
     upper:     4.16s      1039887 ns/op   33485 B/op   1502 allocs/op

  4000 times - Read
       pgx:     0.46s       114676 ns/op     894 B/op      8 allocs/op
  pgx_pool:     0.47s       116923 ns/op    1079 B/op      9 allocs/op
       raw:     0.48s       120428 ns/op    2062 B/op     50 allocs/op
      sqlc:     0.48s       120513 ns/op    2172 B/op     52 allocs/op
     beego:     0.52s       128972 ns/op    2088 B/op     75 allocs/op
    reform:     0.52s       129841 ns/op    3198 B/op     86 allocs/op
      gorp:     0.54s       136130 ns/op    3878 B/op    194 allocs/op
       pop:     0.55s       136685 ns/op    3507 B/op     71 allocs/op
 gorm_prep:     0.55s       137931 ns/op    4801 B/op     84 allocs/op
        pg:     0.57s       142932 ns/op     872 B/op     20 allocs/op
       ent:     0.60s       148989 ns/op    5404 B/op    149 allocs/op
       bun:     0.61s       152083 ns/op    5489 B/op     21 allocs/op
 sqlboiler:     0.61s       152498 ns/op     956 B/op     14 allocs/op
       rel:     0.61s       153743 ns/op    1800 B/op     45 allocs/op
       dbr:     0.62s       155237 ns/op    2184 B/op     37 allocs/op
      gorm:     0.78s       194716 ns/op    5150 B/op     95 allocs/op
      sqlx:     1.10s       273821 ns/op    1976 B/op     43 allocs/op
      godb:     1.15s       288490 ns/op    4112 B/op    102 allocs/op
      xorm:     1.17s       293593 ns/op    4616 B/op    125 allocs/op
     upper:     1.20s       300573 ns/op    8136 B/op    315 allocs/op

  4000 times - MultiRead limit 100
    reform:     0.49s       122555 ns/op    2926 B/op     74 allocs/op
       pgx:     0.95s       236353 ns/op   42949 B/op    504 allocs/op
  pgx_pool:     0.95s       237018 ns/op   43006 B/op    504 allocs/op
     upper:     1.18s       293822 ns/op    7841 B/op    294 allocs/op
       raw:     1.20s       300391 ns/op   38341 B/op   1038 allocs/op
      sqlc:     1.39s       346892 ns/op   73158 B/op   1251 allocs/op
        pg:     1.41s       351751 ns/op   23064 B/op    629 allocs/op
 sqlboiler:     1.56s       390432 ns/op   57827 B/op   1259 allocs/op
      sqlx:     1.59s       397532 ns/op   39064 B/op   1516 allocs/op
      gorp:     1.61s       402202 ns/op   57372 B/op   1494 allocs/op
       ent:     1.65s       411762 ns/op   76045 B/op   2039 allocs/op
       dbr:     1.77s       443524 ns/op   32416 B/op   1545 allocs/op
       bun:     1.80s       449678 ns/op   33993 B/op   1124 allocs/op
     beego:     1.90s       474795 ns/op   55184 B/op   3077 allocs/op
       pop:     1.95s       487137 ns/op   75733 B/op   1510 allocs/op
 gorm_prep:     2.19s       548137 ns/op   56051 B/op   2078 allocs/op
      gorm:     2.54s       635663 ns/op   57179 B/op   2378 allocs/op
       rel:     2.69s       673151 ns/op   95322 B/op   2248 allocs/op
      godb:     2.83s       707146 ns/op   75264 B/op   3084 allocs/op
      xorm:     3.09s       773654 ns/op  119953 B/op   4687 allocs/op
```
