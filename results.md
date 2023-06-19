# Results

- orm-benchmark (with no flags)

```
   
Reports:

   200 times - Insert
  pgx_pool:     0.05s       245223 ns/op     265 B/op     10 allocs/op
       ent:     0.06s       291065 ns/op    4140 B/op     96 allocs/op
       raw:     0.06s       312332 ns/op    1006 B/op     15 allocs/op
     beego:     0.06s       313152 ns/op    2381 B/op     57 allocs/op
       pgx:     0.06s       317422 ns/op     244 B/op     10 allocs/op
    reform:     0.06s       319363 ns/op    2083 B/op     53 allocs/op
       bun:     0.07s       333007 ns/op    5108 B/op     14 allocs/op
      gorp:     0.07s       337982 ns/op    2068 B/op     43 allocs/op
      gorm:     0.07s       348374 ns/op    7255 B/op    106 allocs/op
        pg:     0.07s       348745 ns/op    1029 B/op     10 allocs/op
      sqlc:     0.07s       356534 ns/op    3175 B/op     64 allocs/op
 sqlboiler:     0.07s       363480 ns/op    1865 B/op     36 allocs/op
       dbr:     0.07s       368020 ns/op    2803 B/op     66 allocs/op
 gorm_prep:     0.08s       404777 ns/op    5526 B/op     66 allocs/op
      sqlx:     0.08s       421342 ns/op     856 B/op     19 allocs/op
      zorm:     0.09s       462627 ns/op    3802 B/op     76 allocs/op
     upper:     0.09s       467551 ns/op    6169 B/op    130 allocs/op
      godb:     0.09s       469332 ns/op    4641 B/op    116 allocs/op
       rel:     0.10s       484309 ns/op    2614 B/op     44 allocs/op
      xorm:     0.10s       499582 ns/op    3447 B/op     90 allocs/op
       pop:     0.13s       648481 ns/op    9939 B/op    240 allocs/op

   200 times - MultiInsert 100 row
       pgx:     0.20s       975998 ns/op  113077 B/op     43 allocs/op
  pgx_pool:     0.23s      1143982 ns/op  112946 B/op     43 allocs/op
       raw:     0.24s      1193022 ns/op  184802 B/op    936 allocs/op
     beego:     0.24s      1217760 ns/op  177682 B/op   2745 allocs/op
        pg:     0.29s      1448012 ns/op    3316 B/op    112 allocs/op
    reform:     0.30s      1509847 ns/op  458873 B/op   2745 allocs/op
 gorm_prep:     0.30s      1511578 ns/op  251046 B/op   1888 allocs/op
       ent:     0.31s      1541119 ns/op  389237 B/op   4600 allocs/op
       bun:     0.31s      1560561 ns/op   42484 B/op    217 allocs/op
      gorm:     0.36s      1778054 ns/op  291329 B/op   5229 allocs/op
      zorm:     0.39s      1928663 ns/op  199927 B/op   2780 allocs/op
      xorm:     0.39s      1968804 ns/op  248094 B/op   5414 allocs/op
       rel:     0.41s      2071917 ns/op  306877 B/op   3262 allocs/op
     upper:     0.44s      2217522 ns/op  323194 B/op   4218 allocs/op
      sqlx:     0.51s      2556643 ns/op  169926 B/op   1551 allocs/op
      godb:     0.74s      3689974 ns/op  254012 B/op   5892 allocs/op
      sqlc:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert

   200 times - Update
       raw:     0.02s       110243 ns/op     722 B/op     13 allocs/op
       ent:     0.03s       136380 ns/op    4690 B/op     97 allocs/op
      sqlx:     0.05s       268245 ns/op     872 B/op     20 allocs/op
       pgx:     0.05s       271426 ns/op     282 B/op     10 allocs/op
      sqlc:     0.06s       283567 ns/op     891 B/op     14 allocs/op
    reform:     0.06s       292369 ns/op    1747 B/op     51 allocs/op
  pgx_pool:     0.06s       298128 ns/op     302 B/op     10 allocs/op
      gorp:     0.06s       305542 ns/op    1223 B/op     32 allocs/op
     beego:     0.06s       306466 ns/op    1758 B/op     47 allocs/op
 gorm_prep:     0.06s       310143 ns/op    5011 B/op     56 allocs/op
        pg:     0.06s       310155 ns/op     768 B/op      9 allocs/op
       pop:     0.07s       337070 ns/op    6022 B/op    186 allocs/op
       dbr:     0.07s       337963 ns/op    2651 B/op     57 allocs/op
 sqlboiler:     0.07s       349826 ns/op     923 B/op     17 allocs/op
       bun:     0.08s       379973 ns/op    4728 B/op      5 allocs/op
      gorm:     0.08s       382976 ns/op    6752 B/op     99 allocs/op
      godb:     0.09s       444685 ns/op    5128 B/op    154 allocs/op
      zorm:     0.09s       456262 ns/op    3024 B/op     59 allocs/op
      xorm:     0.10s       479949 ns/op    3944 B/op    132 allocs/op
       rel:     0.10s       480074 ns/op    3051 B/op     45 allocs/op
     upper:     0.21s      1061447 ns/op   16721 B/op    391 allocs/op

   200 times - Read
  pgx_pool:     0.02s       117110 ns/op    1078 B/op      9 allocs/op
       pgx:     0.02s       120265 ns/op     886 B/op      8 allocs/op
      sqlc:     0.02s       121168 ns/op    2162 B/op     52 allocs/op
       raw:     0.03s       125710 ns/op    2098 B/op     50 allocs/op
     beego:     0.03s       128705 ns/op    2101 B/op     76 allocs/op
    reform:     0.03s       134916 ns/op    3230 B/op     86 allocs/op
       ent:     0.03s       135714 ns/op    5602 B/op    144 allocs/op
      gorp:     0.03s       138954 ns/op    3869 B/op    194 allocs/op
       pop:     0.03s       139230 ns/op    3572 B/op     67 allocs/op
        pg:     0.03s       145845 ns/op     873 B/op     20 allocs/op
 sqlboiler:     0.03s       146416 ns/op    1118 B/op     14 allocs/op
 gorm_prep:     0.03s       146772 ns/op    4424 B/op     87 allocs/op
       bun:     0.03s       153566 ns/op    5522 B/op     21 allocs/op
       dbr:     0.03s       156098 ns/op    2184 B/op     37 allocs/op
       rel:     0.03s       158051 ns/op    2307 B/op     47 allocs/op
      zorm:     0.03s       167452 ns/op    3064 B/op     65 allocs/op
      gorm:     0.04s       188650 ns/op    4789 B/op     98 allocs/op
      sqlx:     0.06s       275886 ns/op    1976 B/op     43 allocs/op
      godb:     0.06s       282732 ns/op    4080 B/op    102 allocs/op
     upper:     0.06s       284034 ns/op    5040 B/op    110 allocs/op
      xorm:     0.06s       292813 ns/op    4649 B/op    127 allocs/op

   200 times - MultiRead limit 100
    reform:     0.03s       131927 ns/op    2955 B/op     74 allocs/op
       pgx:     0.05s       245972 ns/op   42919 B/op    504 allocs/op
     upper:     0.06s       275272 ns/op    4819 B/op     89 allocs/op
  pgx_pool:     0.06s       278945 ns/op   42993 B/op    504 allocs/op
       raw:     0.06s       316342 ns/op   38353 B/op   1038 allocs/op
        pg:     0.07s       348640 ns/op   27505 B/op    629 allocs/op
      sqlc:     0.07s       360162 ns/op   73128 B/op   1251 allocs/op
       ent:     0.08s       399997 ns/op   78022 B/op   2035 allocs/op
      sqlx:     0.08s       404314 ns/op   37480 B/op   1225 allocs/op
       pop:     0.08s       420775 ns/op   68502 B/op   1306 allocs/op
       bun:     0.09s       442310 ns/op   34026 B/op   1124 allocs/op
      gorp:     0.09s       450535 ns/op   57380 B/op   1494 allocs/op
       dbr:     0.09s       457247 ns/op   30816 B/op   1254 allocs/op
     beego:     0.09s       469019 ns/op   55200 B/op   3078 allocs/op
 gorm_prep:     0.10s       491733 ns/op   43160 B/op   2081 allocs/op
 sqlboiler:     0.11s       541715 ns/op   66397 B/op   2259 allocs/op
      gorm:     0.12s       579073 ns/op   44307 B/op   2191 allocs/op
      zorm:     0.13s       664450 ns/op  161664 B/op   2950 allocs/op
      godb:     0.14s       691782 ns/op   75235 B/op   3084 allocs/op
       rel:     0.14s       707032 ns/op  100640 B/op   2253 allocs/op
      xorm:     0.15s       772515 ns/op  119397 B/op   4401 allocs/op
```

- orm-benchmark -multi=5

```
  
Reports:

  1000 times - Insert
 sqlboiler:     0.26s       257307 ns/op    1620 B/op     35 allocs/op
       pgx:     0.27s       267682 ns/op     265 B/op     10 allocs/op
  pgx_pool:     0.27s       270634 ns/op     281 B/op     10 allocs/op
      gorp:     0.29s       286616 ns/op    1827 B/op     42 allocs/op
     beego:     0.30s       302955 ns/op    2383 B/op     57 allocs/op
       raw:     0.30s       303408 ns/op     763 B/op     13 allocs/op
    reform:     0.31s       308163 ns/op    1837 B/op     51 allocs/op
      sqlc:     0.31s       308444 ns/op    2928 B/op     63 allocs/op
       ent:     0.32s       315967 ns/op    4142 B/op     97 allocs/op
        pg:     0.32s       318680 ns/op    1890 B/op     10 allocs/op
 gorm_prep:     0.33s       329019 ns/op    5232 B/op     66 allocs/op
       bun:     0.34s       340273 ns/op    5054 B/op     14 allocs/op
      gorm:     0.36s       356024 ns/op    7216 B/op    105 allocs/op
       dbr:     0.36s       361807 ns/op    2711 B/op     65 allocs/op
      sqlx:     0.44s       438137 ns/op     856 B/op     19 allocs/op
      zorm:     0.44s       438740 ns/op    3786 B/op     77 allocs/op
       rel:     0.44s       444000 ns/op    2608 B/op     45 allocs/op
      xorm:     0.47s       470786 ns/op    3344 B/op     89 allocs/op
     upper:     0.48s       475162 ns/op    5937 B/op    126 allocs/op
      godb:     0.48s       478437 ns/op    4545 B/op    116 allocs/op
       pop:     0.64s       636313 ns/op    9654 B/op    239 allocs/op

  1000 times - MultiInsert 100 row
       pgx:     0.98s       984342 ns/op  112906 B/op     42 allocs/op
  pgx_pool:     1.01s      1005908 ns/op  112948 B/op     43 allocs/op
       raw:     1.13s      1132061 ns/op  183877 B/op    930 allocs/op
     beego:     1.20s      1196873 ns/op  177599 B/op   2745 allocs/op
 gorm_prep:     1.32s      1316238 ns/op  251004 B/op   1889 allocs/op
    reform:     1.39s      1387219 ns/op  458790 B/op   2745 allocs/op
        pg:     1.50s      1497764 ns/op    4437 B/op    112 allocs/op
       ent:     1.52s      1519804 ns/op  389013 B/op   4598 allocs/op
       bun:     1.60s      1604599 ns/op   42487 B/op    218 allocs/op
      sqlx:     1.75s      1748106 ns/op  169615 B/op   1551 allocs/op
      gorm:     1.84s      1840916 ns/op  291353 B/op   5231 allocs/op
      zorm:     1.97s      1969146 ns/op  199928 B/op   2780 allocs/op
       rel:     2.09s      2091313 ns/op  306893 B/op   3264 allocs/op
      xorm:     2.13s      2129732 ns/op  248096 B/op   5414 allocs/op
     upper:     2.21s      2212780 ns/op  322949 B/op   4207 allocs/op
      godb:     2.24s      2244876 ns/op  254022 B/op   5894 allocs/op
      gorp:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert

  1000 times - Update
       raw:     0.12s       121947 ns/op     753 B/op     13 allocs/op
       ent:     0.14s       136828 ns/op    4680 B/op     97 allocs/op
 sqlboiler:     0.26s       264770 ns/op     907 B/op     17 allocs/op
      sqlx:     0.27s       269446 ns/op     872 B/op     20 allocs/op
     beego:     0.28s       278333 ns/op    1752 B/op     47 allocs/op
      sqlc:     0.28s       280087 ns/op     880 B/op     14 allocs/op
      gorp:     0.29s       289028 ns/op    1209 B/op     32 allocs/op
  pgx_pool:     0.29s       294809 ns/op     287 B/op     10 allocs/op
       pgx:     0.30s       298221 ns/op     273 B/op     10 allocs/op
    reform:     0.30s       300539 ns/op    1777 B/op     51 allocs/op
       pop:     0.32s       319239 ns/op    6051 B/op    186 allocs/op
 gorm_prep:     0.33s       334165 ns/op    5009 B/op     56 allocs/op
        pg:     0.35s       348450 ns/op     768 B/op      9 allocs/op
       dbr:     0.35s       354382 ns/op    2651 B/op     57 allocs/op
       bun:     0.36s       359828 ns/op    4728 B/op      5 allocs/op
      gorm:     0.37s       366409 ns/op    6752 B/op     99 allocs/op
      godb:     0.46s       460261 ns/op    5128 B/op    154 allocs/op
      xorm:     0.46s       463270 ns/op    3944 B/op    132 allocs/op
      zorm:     0.48s       480330 ns/op    3024 B/op     59 allocs/op
       rel:     0.49s       488285 ns/op    3048 B/op     45 allocs/op
     upper:     1.03s      1028250 ns/op   16665 B/op    390 allocs/op

  1000 times - Read
  pgx_pool:     0.12s       117853 ns/op    1076 B/op      9 allocs/op
       pgx:     0.12s       118958 ns/op     889 B/op      8 allocs/op
       raw:     0.12s       122748 ns/op    2058 B/op     50 allocs/op
      sqlc:     0.12s       124274 ns/op    2176 B/op     52 allocs/op
    reform:     0.13s       128144 ns/op    3193 B/op     86 allocs/op
     beego:     0.13s       130860 ns/op    2097 B/op     76 allocs/op
       pop:     0.14s       137519 ns/op    3217 B/op     67 allocs/op
 gorm_prep:     0.14s       141614 ns/op    4408 B/op     87 allocs/op
      gorp:     0.14s       142293 ns/op    3881 B/op    194 allocs/op
        pg:     0.15s       146628 ns/op     872 B/op     20 allocs/op
       dbr:     0.15s       150318 ns/op    2184 B/op     37 allocs/op
 sqlboiler:     0.16s       156111 ns/op     971 B/op     14 allocs/op
       ent:     0.16s       156965 ns/op    5624 B/op    144 allocs/op
       bun:     0.16s       159112 ns/op    5521 B/op     21 allocs/op
       rel:     0.16s       160058 ns/op    2304 B/op     47 allocs/op
      zorm:     0.17s       165834 ns/op    3064 B/op     65 allocs/op
      gorm:     0.20s       198471 ns/op    4776 B/op     98 allocs/op
      sqlx:     0.28s       278604 ns/op    1976 B/op     43 allocs/op
      godb:     0.29s       294474 ns/op    4080 B/op    102 allocs/op
     upper:     0.30s       297088 ns/op    5055 B/op    110 allocs/op
      xorm:     0.30s       303759 ns/op    4649 B/op    127 allocs/op

  1000 times - MultiRead limit 100
    reform:     0.12s       122211 ns/op    2920 B/op     74 allocs/op
       pgx:     0.24s       237885 ns/op   42949 B/op    504 allocs/op
  pgx_pool:     0.25s       248893 ns/op   43009 B/op    504 allocs/op
     upper:     0.29s       290889 ns/op    4783 B/op     89 allocs/op
       raw:     0.31s       311398 ns/op   38342 B/op   1038 allocs/op
        pg:     0.35s       346615 ns/op   23381 B/op    629 allocs/op
      sqlc:     0.37s       365880 ns/op   73158 B/op   1251 allocs/op
      sqlx:     0.39s       391437 ns/op   37480 B/op   1225 allocs/op
      gorp:     0.40s       395313 ns/op   57377 B/op   1494 allocs/op
       ent:     0.43s       427307 ns/op   77993 B/op   2035 allocs/op
       dbr:     0.44s       442816 ns/op   30816 B/op   1254 allocs/op
       pop:     0.45s       451865 ns/op   68881 B/op   1307 allocs/op
       bun:     0.46s       457082 ns/op   34038 B/op   1124 allocs/op
     beego:     0.47s       471525 ns/op   55192 B/op   3078 allocs/op
 sqlboiler:     0.49s       489430 ns/op   66685 B/op   2259 allocs/op
 gorm_prep:     0.51s       508911 ns/op   43152 B/op   2081 allocs/op
      gorm:     0.58s       580695 ns/op   44308 B/op   2191 allocs/op
      zorm:     0.69s       693050 ns/op  161665 B/op   2950 allocs/op
       rel:     0.71s       710302 ns/op  100640 B/op   2253 allocs/op
      godb:     0.72s       721652 ns/op   75238 B/op   3084 allocs/op
      xorm:     0.78s       782485 ns/op  119381 B/op   4401 allocs/op
```

- orm-benchmark -multi=10

```
  
Reports:

  2000 times - Insert
     beego:     0.49s       246524 ns/op    2384 B/op     57 allocs/op
       pgx:     0.52s       260097 ns/op     269 B/op     10 allocs/op
      gorp:     0.54s       271995 ns/op    1798 B/op     42 allocs/op
       ent:     0.55s       273303 ns/op    4143 B/op     97 allocs/op
  pgx_pool:     0.55s       275674 ns/op     285 B/op     10 allocs/op
       raw:     0.58s       287501 ns/op     733 B/op     13 allocs/op
 sqlboiler:     0.59s       294443 ns/op    1591 B/op     35 allocs/op
    reform:     0.60s       302423 ns/op    1807 B/op     51 allocs/op
      sqlc:     0.62s       311126 ns/op    2900 B/op     63 allocs/op
 gorm_prep:     0.65s       322601 ns/op    5277 B/op     66 allocs/op
       dbr:     0.67s       335770 ns/op    2699 B/op     65 allocs/op
        pg:     0.68s       340914 ns/op    1341 B/op     10 allocs/op
       bun:     0.69s       343584 ns/op    5023 B/op     14 allocs/op
      gorm:     0.71s       356073 ns/op    7183 B/op    105 allocs/op
      sqlx:     0.85s       425247 ns/op     856 B/op     19 allocs/op
      zorm:     0.90s       447821 ns/op    3785 B/op     77 allocs/op
       rel:     0.91s       454875 ns/op    2608 B/op     45 allocs/op
      xorm:     0.91s       455744 ns/op    3333 B/op     89 allocs/op
      godb:     0.93s       463339 ns/op    4533 B/op    116 allocs/op
     upper:     0.93s       463400 ns/op    5910 B/op    126 allocs/op
       pop:     1.23s       617219 ns/op    9590 B/op    239 allocs/op

  2000 times - MultiInsert 100 row
       pgx:     1.96s       980007 ns/op  112939 B/op     43 allocs/op
  pgx_pool:     2.03s      1015918 ns/op  112928 B/op     42 allocs/op
     beego:     2.32s      1162158 ns/op  177650 B/op   2745 allocs/op
       raw:     2.39s      1195565 ns/op  183763 B/op    929 allocs/op
    reform:     2.61s      1303542 ns/op  458750 B/op   2745 allocs/op
 gorm_prep:     2.65s      1327046 ns/op  250997 B/op   1890 allocs/op
        pg:     2.95s      1475785 ns/op    3314 B/op    112 allocs/op
       bun:     3.08s      1537734 ns/op   42401 B/op    218 allocs/op
       ent:     3.11s      1555863 ns/op  388987 B/op   4598 allocs/op
      sqlx:     3.52s      1759508 ns/op  170081 B/op   1551 allocs/op
      gorm:     3.54s      1770476 ns/op  291339 B/op   5231 allocs/op
      zorm:     4.01s      2003168 ns/op  199928 B/op   2780 allocs/op
       rel:     4.24s      2120343 ns/op  306895 B/op   3264 allocs/op
      xorm:     4.25s      2124099 ns/op  248150 B/op   5414 allocs/op
      godb:     4.44s      2220986 ns/op  254014 B/op   5894 allocs/op
     upper:     4.47s      2234126 ns/op  322915 B/op   4205 allocs/op
      sqlc:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert

  2000 times - Update
       raw:     0.26s       127973 ns/op     748 B/op     13 allocs/op
       ent:     0.27s       137330 ns/op    4680 B/op     97 allocs/op
     beego:     0.50s       250847 ns/op    1752 B/op     47 allocs/op
      sqlx:     0.53s       263599 ns/op     872 B/op     20 allocs/op
       pgx:     0.55s       274741 ns/op     268 B/op     10 allocs/op
  pgx_pool:     0.56s       279795 ns/op     285 B/op     10 allocs/op
 sqlboiler:     0.56s       281794 ns/op     902 B/op     17 allocs/op
      gorp:     0.58s       291103 ns/op    1205 B/op     32 allocs/op
      sqlc:     0.58s       292220 ns/op     876 B/op     14 allocs/op
       dbr:     0.62s       310718 ns/op    2651 B/op     57 allocs/op
    reform:     0.63s       316862 ns/op    1772 B/op     51 allocs/op
       pop:     0.65s       325799 ns/op    6047 B/op    186 allocs/op
        pg:     0.65s       326132 ns/op     768 B/op      9 allocs/op
 gorm_prep:     0.67s       332652 ns/op    5008 B/op     56 allocs/op
       bun:     0.71s       355071 ns/op    4728 B/op      5 allocs/op
      gorm:     0.74s       369152 ns/op    6752 B/op     99 allocs/op
      zorm:     0.95s       473656 ns/op    3024 B/op     59 allocs/op
      xorm:     0.95s       475018 ns/op    3944 B/op    132 allocs/op
       rel:     0.96s       480808 ns/op    3048 B/op     45 allocs/op
      godb:     0.98s       491155 ns/op    5128 B/op    154 allocs/op
     upper:     2.12s      1059614 ns/op   16657 B/op    390 allocs/op

  2000 times - Read
       pgx:     0.22s       108274 ns/op     892 B/op      8 allocs/op
  pgx_pool:     0.24s       120100 ns/op    1080 B/op      9 allocs/op
      sqlc:     0.24s       121897 ns/op    2172 B/op     52 allocs/op
       raw:     0.26s       130060 ns/op    2065 B/op     50 allocs/op
     beego:     0.26s       130299 ns/op    2097 B/op     76 allocs/op
       pop:     0.27s       133941 ns/op    3185 B/op     67 allocs/op
    reform:     0.27s       136645 ns/op    3197 B/op     86 allocs/op
      gorp:     0.28s       141403 ns/op    3876 B/op    194 allocs/op
 gorm_prep:     0.29s       143435 ns/op    4405 B/op     87 allocs/op
       ent:     0.29s       144231 ns/op    5628 B/op    144 allocs/op
        pg:     0.30s       147890 ns/op     872 B/op     20 allocs/op
 sqlboiler:     0.31s       157082 ns/op     956 B/op     14 allocs/op
       bun:     0.32s       159022 ns/op    5522 B/op     21 allocs/op
       dbr:     0.32s       161055 ns/op    2184 B/op     37 allocs/op
       rel:     0.34s       169903 ns/op    2304 B/op     47 allocs/op
      zorm:     0.34s       171856 ns/op    3064 B/op     65 allocs/op
      gorm:     0.39s       196313 ns/op    4771 B/op     98 allocs/op
      sqlx:     0.57s       283721 ns/op    1976 B/op     43 allocs/op
     upper:     0.58s       290325 ns/op    5054 B/op    110 allocs/op
      godb:     0.59s       294719 ns/op    4080 B/op    102 allocs/op
      xorm:     0.60s       301521 ns/op    4648 B/op    127 allocs/op

  2000 times - MultiRead limit 100
    reform:     0.26s       129680 ns/op    2924 B/op     74 allocs/op
       pgx:     0.48s       237615 ns/op   42949 B/op    504 allocs/op
  pgx_pool:     0.50s       252070 ns/op   43009 B/op    504 allocs/op
     upper:     0.58s       289720 ns/op    4776 B/op     89 allocs/op
       raw:     0.66s       327819 ns/op   38342 B/op   1038 allocs/op
      sqlc:     0.70s       349083 ns/op   73158 B/op   1251 allocs/op
        pg:     0.70s       352196 ns/op   22258 B/op    629 allocs/op
      gorp:     0.81s       403418 ns/op   57372 B/op   1494 allocs/op
      sqlx:     0.83s       416747 ns/op   37480 B/op   1225 allocs/op
       pop:     0.84s       419457 ns/op   68472 B/op   1306 allocs/op
       ent:     0.85s       426475 ns/op   77989 B/op   2035 allocs/op
       dbr:     0.93s       465730 ns/op   30816 B/op   1254 allocs/op
       bun:     0.95s       475073 ns/op   34027 B/op   1124 allocs/op
     beego:     0.97s       484445 ns/op   55198 B/op   3078 allocs/op
 gorm_prep:     0.99s       496064 ns/op   43151 B/op   2081 allocs/op
 sqlboiler:     1.03s       512654 ns/op   66247 B/op   2259 allocs/op
      gorm:     1.14s       569394 ns/op   44296 B/op   2191 allocs/op
      zorm:     1.35s       677404 ns/op  161664 B/op   2950 allocs/op
      godb:     1.41s       705744 ns/op   75235 B/op   3084 allocs/op
       rel:     1.45s       722970 ns/op  100640 B/op   2253 allocs/op
      xorm:     1.58s       789346 ns/op  119381 B/op   4401 allocs/op
```

- orm-benchmark -multi=20

```
  
Reports:

  4000 times - Insert
       pgx:     1.07s       266989 ns/op     270 B/op     10 allocs/op
  pgx_pool:     1.08s       271012 ns/op     286 B/op     10 allocs/op
     beego:     1.16s       290417 ns/op    2384 B/op     57 allocs/op
      gorp:     1.18s       294335 ns/op    1782 B/op     42 allocs/op
       raw:     1.18s       295182 ns/op     719 B/op     13 allocs/op
    reform:     1.20s       300246 ns/op    1790 B/op     51 allocs/op
 sqlboiler:     1.21s       303269 ns/op    1574 B/op     35 allocs/op
       ent:     1.24s       310162 ns/op    4144 B/op     97 allocs/op
      sqlc:     1.27s       317201 ns/op    2884 B/op     63 allocs/op
       dbr:     1.28s       318801 ns/op    2693 B/op     65 allocs/op
 gorm_prep:     1.28s       319827 ns/op    5178 B/op     66 allocs/op
        pg:     1.33s       332161 ns/op    1068 B/op     10 allocs/op
       bun:     1.39s       348268 ns/op    5009 B/op     14 allocs/op
      gorm:     1.50s       373829 ns/op    7174 B/op    106 allocs/op
      sqlx:     1.70s       425541 ns/op     856 B/op     19 allocs/op
      zorm:     1.78s       445606 ns/op    3784 B/op     77 allocs/op
      xorm:     1.79s       447794 ns/op    3326 B/op     89 allocs/op
       rel:     1.81s       451360 ns/op    2608 B/op     45 allocs/op
     upper:     1.89s       473012 ns/op    5895 B/op    126 allocs/op
      godb:     1.90s       474629 ns/op    4526 B/op    116 allocs/op
       pop:     2.53s       632966 ns/op    9585 B/op    239 allocs/op

  4000 times - MultiInsert 100 row
       pgx:     3.81s       953203 ns/op  112916 B/op     43 allocs/op
  pgx_pool:     3.82s       955882 ns/op  112924 B/op     42 allocs/op
       raw:     4.43s      1106833 ns/op  183706 B/op    929 allocs/op
     beego:     4.62s      1155078 ns/op  177599 B/op   2745 allocs/op
 gorm_prep:     5.06s      1264932 ns/op  250982 B/op   1890 allocs/op
    reform:     5.25s      1312490 ns/op  458784 B/op   2745 allocs/op
        pg:     5.70s      1426244 ns/op    3857 B/op    112 allocs/op
       ent:     5.94s      1484563 ns/op  388973 B/op   4598 allocs/op
       bun:     6.17s      1543143 ns/op   42465 B/op    219 allocs/op
      sqlx:     6.94s      1734261 ns/op  169714 B/op   1551 allocs/op
      gorm:     7.05s      1763740 ns/op  291335 B/op   5231 allocs/op
      zorm:     7.58s      1894157 ns/op  199927 B/op   2780 allocs/op
      xorm:     8.14s      2034555 ns/op  248102 B/op   5414 allocs/op
       rel:     8.29s      2071946 ns/op  306896 B/op   3264 allocs/op
      godb:     8.50s      2124676 ns/op  254008 B/op   5894 allocs/op
     upper:     8.81s      2202037 ns/op  322904 B/op   4205 allocs/op
       dbr:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert

  4000 times - Update
       raw:     0.48s       120638 ns/op     750 B/op     13 allocs/op
       ent:     0.57s       141500 ns/op    4678 B/op     97 allocs/op
      sqlx:     1.06s       265594 ns/op     872 B/op     20 allocs/op
       pgx:     1.07s       267763 ns/op     270 B/op     10 allocs/op
  pgx_pool:     1.14s       284423 ns/op     286 B/op     10 allocs/op
      sqlc:     1.14s       286152 ns/op     878 B/op     14 allocs/op
 sqlboiler:     1.16s       288805 ns/op     901 B/op     17 allocs/op
      gorp:     1.16s       290873 ns/op    1204 B/op     32 allocs/op
     beego:     1.17s       293290 ns/op    1752 B/op     47 allocs/op
    reform:     1.20s       300232 ns/op    1775 B/op     51 allocs/op
       pop:     1.28s       321201 ns/op    6047 B/op    186 allocs/op
 gorm_prep:     1.30s       323871 ns/op    5008 B/op     56 allocs/op
       dbr:     1.33s       331948 ns/op    2651 B/op     57 allocs/op
        pg:     1.36s       340797 ns/op     768 B/op      9 allocs/op
       bun:     1.43s       357751 ns/op    4728 B/op      5 allocs/op
      gorm:     1.45s       362857 ns/op    6752 B/op     99 allocs/op
      xorm:     1.83s       458012 ns/op    3944 B/op    132 allocs/op
      zorm:     1.84s       460400 ns/op    3024 B/op     59 allocs/op
      godb:     1.94s       484435 ns/op    5128 B/op    154 allocs/op
       rel:     1.98s       493845 ns/op    3048 B/op     45 allocs/op
     upper:     4.26s      1064031 ns/op   16652 B/op    390 allocs/op

  4000 times - Read
    reform:     0.47s       117544 ns/op    3198 B/op     86 allocs/op
  pgx_pool:     0.48s       120665 ns/op    1079 B/op      9 allocs/op
       pgx:     0.49s       121967 ns/op     894 B/op      8 allocs/op
      sqlc:     0.49s       122101 ns/op    2172 B/op     52 allocs/op
       raw:     0.50s       124489 ns/op    2062 B/op     50 allocs/op
     beego:     0.52s       128807 ns/op    2096 B/op     76 allocs/op
 gorm_prep:     0.57s       142020 ns/op    4403 B/op     87 allocs/op
       pop:     0.58s       144071 ns/op    3172 B/op     67 allocs/op
      gorp:     0.58s       145040 ns/op    3878 B/op    194 allocs/op
        pg:     0.59s       146551 ns/op    1134 B/op     20 allocs/op
       ent:     0.60s       150332 ns/op    5628 B/op    144 allocs/op
       dbr:     0.63s       156860 ns/op    2184 B/op     37 allocs/op
       bun:     0.63s       157647 ns/op    5522 B/op     21 allocs/op
 sqlboiler:     0.65s       161596 ns/op     956 B/op     14 allocs/op
       rel:     0.65s       163512 ns/op    2304 B/op     47 allocs/op
      zorm:     0.69s       171788 ns/op    3064 B/op     65 allocs/op
      gorm:     0.79s       198385 ns/op    4770 B/op     98 allocs/op
      sqlx:     1.12s       280085 ns/op    1976 B/op     43 allocs/op
      godb:     1.17s       292333 ns/op    4080 B/op    102 allocs/op
     upper:     1.20s       299710 ns/op    5055 B/op    110 allocs/op
      xorm:     1.21s       303227 ns/op    4648 B/op    127 allocs/op

  4000 times - MultiRead limit 100
    reform:     0.53s       131992 ns/op    2926 B/op     74 allocs/op
  pgx_pool:     0.98s       244611 ns/op   43007 B/op    504 allocs/op
       pgx:     1.01s       252543 ns/op   42949 B/op    504 allocs/op
     upper:     1.18s       294925 ns/op    4777 B/op     89 allocs/op
       raw:     1.24s       309862 ns/op   38341 B/op   1038 allocs/op
        pg:     1.43s       358053 ns/op   22802 B/op    629 allocs/op
      sqlc:     1.45s       363618 ns/op   73158 B/op   1251 allocs/op
      gorp:     1.61s       401576 ns/op   57370 B/op   1494 allocs/op
       ent:     1.63s       407748 ns/op   77990 B/op   2035 allocs/op
       pop:     1.65s       413467 ns/op   68923 B/op   1306 allocs/op
      sqlx:     1.68s       419739 ns/op   37480 B/op   1225 allocs/op
       dbr:     1.78s       444927 ns/op   30816 B/op   1254 allocs/op
       bun:     1.88s       470339 ns/op   34033 B/op   1124 allocs/op
     beego:     1.90s       474524 ns/op   55195 B/op   3078 allocs/op
 gorm_prep:     2.00s       499107 ns/op   43139 B/op   2081 allocs/op
 sqlboiler:     2.03s       507052 ns/op   66517 B/op   2259 allocs/op
      gorm:     2.31s       577302 ns/op   44292 B/op   2191 allocs/op
      zorm:     2.60s       649175 ns/op  161664 B/op   2950 allocs/op
      godb:     2.79s       697869 ns/op   75232 B/op   3084 allocs/op
       rel:     2.94s       734456 ns/op  100640 B/op   2253 allocs/op
      xorm:     3.17s       791852 ns/op  119375 B/op   4401 allocs/op
```
