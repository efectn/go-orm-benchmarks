# Results

- orm-benchmark (with no flags)

```
   
Reports:

   200 times - Insert
       pgx:     0.04s       218533 ns/op     244 B/op     10 allocs/op
  pgx_pool:     0.04s       224933 ns/op     265 B/op     10 allocs/op
     beego:     0.05s       242722 ns/op    2375 B/op     56 allocs/op
       ent:     0.05s       248935 ns/op    4172 B/op     99 allocs/op
      gorp:     0.05s       260666 ns/op    2070 B/op     43 allocs/op
    reform:     0.05s       269787 ns/op    2091 B/op     54 allocs/op
 sqlboiler:     0.05s       271063 ns/op    1871 B/op     37 allocs/op
        pg:     0.05s       272491 ns/op    1031 B/op     10 allocs/op
      sqlc:     0.06s       275936 ns/op    3173 B/op     64 allocs/op
       dbr:     0.06s       301618 ns/op    2796 B/op     66 allocs/op
      gorm:     0.06s       304156 ns/op    7132 B/op     89 allocs/op
 gorm_prep:     0.06s       308459 ns/op    5603 B/op     69 allocs/op
       bun:     0.06s       310935 ns/op    5113 B/op     14 allocs/op
       raw:     0.07s       327113 ns/op    1007 B/op     15 allocs/op
      sqlx:     0.08s       375026 ns/op     856 B/op     19 allocs/op
       rel:     0.08s       394976 ns/op    2511 B/op     41 allocs/op
      godb:     0.09s       427806 ns/op    4669 B/op    116 allocs/op
      xorm:     0.09s       428479 ns/op    3473 B/op     90 allocs/op
     upper:     0.10s       479313 ns/op   14203 B/op    685 allocs/op
       pop:     0.12s       620805 ns/op   10434 B/op    250 allocs/op

   200 times - MultiInsert 100 row
  pgx_pool:     0.18s       894597 ns/op  114336 B/op     43 allocs/op
       pgx:     0.19s       939073 ns/op  114299 B/op     42 allocs/op
       raw:     0.20s      1016668 ns/op  192354 B/op    938 allocs/op
     beego:     0.23s      1126527 ns/op  179488 B/op   2746 allocs/op
 gorm_prep:     0.27s      1331997 ns/op  234539 B/op   2182 allocs/op
       ent:     0.28s      1391356 ns/op  412422 B/op   4902 allocs/op
    reform:     0.28s      1395570 ns/op  466426 B/op   2747 allocs/op
       bun:     0.31s      1559872 ns/op   42359 B/op    217 allocs/op
      sqlx:     0.36s      1822046 ns/op  171536 B/op   1552 allocs/op
      gorm:     0.37s      1825034 ns/op  272119 B/op   3728 allocs/op
       rel:     0.40s      2002028 ns/op  303976 B/op   3261 allocs/op
      xorm:     0.40s      2017091 ns/op  254767 B/op   5416 allocs/op
      godb:     0.45s      2242464 ns/op  259948 B/op   5893 allocs/op
     upper:     0.63s      3160448 ns/op  547003 B/op  19610 allocs/op
        pg:     0.69s      3457047 ns/op    3316 B/op    112 allocs/op
      gorp:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert

   200 times - Update
       raw:     0.02s       115212 ns/op     722 B/op     13 allocs/op
       ent:     0.03s       134744 ns/op    4530 B/op     99 allocs/op
      gorp:     0.04s       220502 ns/op    1221 B/op     32 allocs/op
  pgx_pool:     0.04s       223012 ns/op     302 B/op     10 allocs/op
       pgx:     0.04s       224283 ns/op     282 B/op     10 allocs/op
      sqlc:     0.05s       226172 ns/op     891 B/op     14 allocs/op
    reform:     0.05s       227171 ns/op    1744 B/op     51 allocs/op
     beego:     0.05s       233411 ns/op    1755 B/op     47 allocs/op
 sqlboiler:     0.05s       242513 ns/op     925 B/op     17 allocs/op
 gorm_prep:     0.05s       244773 ns/op    5132 B/op     59 allocs/op
      sqlx:     0.05s       266452 ns/op     872 B/op     20 allocs/op
       dbr:     0.06s       279021 ns/op    2651 B/op     57 allocs/op
        pg:     0.06s       280236 ns/op     768 B/op      9 allocs/op
       bun:     0.06s       283303 ns/op    4728 B/op      5 allocs/op
       pop:     0.06s       285669 ns/op    6556 B/op    198 allocs/op
      gorm:     0.07s       329477 ns/op    6625 B/op     81 allocs/op
      xorm:     0.08s       382902 ns/op    3649 B/op    126 allocs/op
       rel:     0.08s       397553 ns/op    2531 B/op     41 allocs/op
      godb:     0.08s       419398 ns/op    5160 B/op    154 allocs/op
     upper:     0.22s      1091279 ns/op   33646 B/op   1510 allocs/op

   200 times - Read
  pgx_pool:     0.02s       109672 ns/op    1078 B/op      9 allocs/op
       pgx:     0.02s       114398 ns/op     886 B/op      8 allocs/op
     beego:     0.02s       119620 ns/op    2094 B/op     75 allocs/op
       raw:     0.02s       123445 ns/op    2098 B/op     50 allocs/op
      sqlc:     0.03s       126667 ns/op    2162 B/op     52 allocs/op
    reform:     0.03s       133584 ns/op    3234 B/op     86 allocs/op
      gorp:     0.03s       137150 ns/op    3869 B/op    194 allocs/op
 gorm_prep:     0.03s       139492 ns/op    4799 B/op     83 allocs/op
       pop:     0.03s       141640 ns/op    3914 B/op     71 allocs/op
        pg:     0.03s       144783 ns/op     873 B/op     20 allocs/op
       ent:     0.03s       151597 ns/op    5422 B/op    149 allocs/op
       bun:     0.03s       151637 ns/op    5490 B/op     21 allocs/op
 sqlboiler:     0.03s       155133 ns/op    1302 B/op     14 allocs/op
       rel:     0.03s       155548 ns/op    1803 B/op     45 allocs/op
       dbr:     0.03s       157170 ns/op    2184 B/op     37 allocs/op
      gorm:     0.04s       187192 ns/op    5154 B/op     94 allocs/op
      sqlx:     0.05s       271936 ns/op    1976 B/op     43 allocs/op
      xorm:     0.06s       291208 ns/op    4615 B/op    125 allocs/op
     upper:     0.06s       296824 ns/op    8126 B/op    315 allocs/op
      godb:     0.06s       321451 ns/op    4112 B/op    102 allocs/op

   200 times - MultiRead limit 100
    reform:     0.03s       125952 ns/op    2955 B/op     74 allocs/op
       pgx:     0.05s       241538 ns/op   42919 B/op    504 allocs/op
  pgx_pool:     0.05s       266012 ns/op   42993 B/op    504 allocs/op
     upper:     0.06s       293725 ns/op    7896 B/op    294 allocs/op
       raw:     0.06s       303647 ns/op   38352 B/op   1038 allocs/op
      sqlc:     0.07s       344999 ns/op   73128 B/op   1251 allocs/op
 sqlboiler:     0.07s       366691 ns/op   58031 B/op   1259 allocs/op
       ent:     0.08s       401110 ns/op   76035 B/op   2039 allocs/op
        pg:     0.08s       405798 ns/op   22258 B/op    629 allocs/op
      gorp:     0.08s       407848 ns/op   57381 B/op   1494 allocs/op
      sqlx:     0.08s       409724 ns/op   39064 B/op   1516 allocs/op
       dbr:     0.09s       438772 ns/op   32416 B/op   1545 allocs/op
     beego:     0.09s       465024 ns/op   55187 B/op   3077 allocs/op
       bun:     0.09s       473119 ns/op   33997 B/op   1124 allocs/op
       pop:     0.10s       483039 ns/op   76048 B/op   1511 allocs/op
 gorm_prep:     0.10s       493507 ns/op   56063 B/op   1977 allocs/op
      gorm:     0.12s       582859 ns/op   57169 B/op   2277 allocs/op
       rel:     0.13s       666417 ns/op   95321 B/op   2248 allocs/op
      godb:     0.14s       690986 ns/op   75263 B/op   3084 allocs/op
      xorm:     0.15s       769534 ns/op  119973 B/op   4687 allocs/op
```

- orm-benchmark -multi=5

```
  
Reports:

  1000 times - Insert
       pgx:     0.22s       215691 ns/op     265 B/op     10 allocs/op
  pgx_pool:     0.22s       218231 ns/op     281 B/op     10 allocs/op
       raw:     0.23s       228389 ns/op     763 B/op     13 allocs/op
     beego:     0.24s       236788 ns/op    2375 B/op     56 allocs/op
 sqlboiler:     0.24s       238594 ns/op    1622 B/op     35 allocs/op
       ent:     0.24s       240297 ns/op    4214 B/op    100 allocs/op
      sqlc:     0.25s       245367 ns/op    2929 B/op     63 allocs/op
    reform:     0.25s       252243 ns/op    1838 B/op     51 allocs/op
 gorm_prep:     0.26s       261892 ns/op    5311 B/op     67 allocs/op
      gorp:     0.26s       264500 ns/op    1827 B/op     42 allocs/op
        pg:     0.27s       265877 ns/op    1889 B/op     10 allocs/op
       dbr:     0.27s       270491 ns/op    2709 B/op     65 allocs/op
       bun:     0.28s       279913 ns/op    5015 B/op     14 allocs/op
      gorm:     0.30s       296692 ns/op    6966 B/op     89 allocs/op
      sqlx:     0.36s       361601 ns/op     856 B/op     19 allocs/op
       rel:     0.39s       391482 ns/op    2504 B/op     42 allocs/op
      godb:     0.39s       394220 ns/op    4576 B/op    116 allocs/op
      xorm:     0.40s       396173 ns/op    3350 B/op     89 allocs/op
     upper:     0.52s       518494 ns/op   13841 B/op    675 allocs/op
       pop:     0.61s       606827 ns/op   10126 B/op    249 allocs/op

  1000 times - MultiInsert 100 row
       pgx:     0.95s       948332 ns/op  114292 B/op     42 allocs/op
  pgx_pool:     0.98s       976417 ns/op  114346 B/op     43 allocs/op
       raw:     1.10s      1103996 ns/op  191429 B/op    932 allocs/op
     beego:     1.14s      1137645 ns/op  179426 B/op   2746 allocs/op
    reform:     1.25s      1252588 ns/op  466329 B/op   2747 allocs/op
 gorm_prep:     1.29s      1290599 ns/op  234456 B/op   2183 allocs/op
       ent:     1.41s      1414321 ns/op  412165 B/op   4900 allocs/op
        pg:     1.55s      1552538 ns/op    3314 B/op    112 allocs/op
       bun:     1.65s      1649430 ns/op   42255 B/op    218 allocs/op
      sqlx:     1.68s      1677371 ns/op  171222 B/op   1552 allocs/op
      gorm:     1.83s      1828698 ns/op  272173 B/op   3729 allocs/op
       rel:     2.05s      2050468 ns/op  303992 B/op   3263 allocs/op
      xorm:     2.05s      2051325 ns/op  254766 B/op   5416 allocs/op
      godb:     2.15s      2149320 ns/op  259953 B/op   5894 allocs/op
     upper:     2.77s      2770828 ns/op  545451 B/op  19495 allocs/op
 sqlboiler:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert

  1000 times - Update
       raw:     0.11s       114575 ns/op     752 B/op     13 allocs/op
       ent:     0.13s       133591 ns/op    4561 B/op     99 allocs/op
  pgx_pool:     0.22s       224194 ns/op     287 B/op     10 allocs/op
      sqlc:     0.23s       225129 ns/op     881 B/op     14 allocs/op
 sqlboiler:     0.23s       225953 ns/op     907 B/op     17 allocs/op
       pgx:     0.23s       226368 ns/op     272 B/op     10 allocs/op
     beego:     0.23s       231537 ns/op    1752 B/op     47 allocs/op
    reform:     0.24s       235573 ns/op    1776 B/op     51 allocs/op
 gorm_prep:     0.25s       252394 ns/op    5113 B/op     59 allocs/op
      gorp:     0.26s       259843 ns/op    1209 B/op     32 allocs/op
      sqlx:     0.26s       260277 ns/op     872 B/op     20 allocs/op
       dbr:     0.26s       260425 ns/op    2651 B/op     57 allocs/op
       bun:     0.28s       281478 ns/op    4728 B/op      5 allocs/op
       pop:     0.28s       281518 ns/op    6588 B/op    198 allocs/op
        pg:     0.29s       286654 ns/op     768 B/op      9 allocs/op
      gorm:     0.35s       352772 ns/op    6600 B/op     81 allocs/op
      xorm:     0.40s       400992 ns/op    3648 B/op    126 allocs/op
       rel:     0.41s       410521 ns/op    2528 B/op     41 allocs/op
      godb:     0.52s       519628 ns/op    5160 B/op    154 allocs/op
     upper:     1.09s      1090421 ns/op   33514 B/op   1503 allocs/op

  1000 times - Read
       pgx:     0.11s       114583 ns/op     889 B/op      8 allocs/op
  pgx_pool:     0.12s       118236 ns/op    1076 B/op      9 allocs/op
      sqlc:     0.12s       119147 ns/op    2176 B/op     52 allocs/op
       raw:     0.12s       120430 ns/op    2058 B/op     50 allocs/op
     beego:     0.13s       125076 ns/op    2089 B/op     75 allocs/op
      gorp:     0.14s       136242 ns/op    3881 B/op    194 allocs/op
 gorm_prep:     0.14s       140762 ns/op    4804 B/op     83 allocs/op
        pg:     0.14s       144230 ns/op     872 B/op     20 allocs/op
       pop:     0.14s       144732 ns/op    3561 B/op     71 allocs/op
       rel:     0.15s       146205 ns/op    1800 B/op     45 allocs/op
    reform:     0.15s       146910 ns/op    3193 B/op     86 allocs/op
       bun:     0.15s       149108 ns/op    5491 B/op     21 allocs/op
       dbr:     0.15s       152112 ns/op    2184 B/op     37 allocs/op
 sqlboiler:     0.16s       155972 ns/op     971 B/op     14 allocs/op
       ent:     0.16s       161514 ns/op    5401 B/op    149 allocs/op
      gorm:     0.19s       186863 ns/op    5155 B/op     94 allocs/op
      sqlx:     0.28s       275282 ns/op    1976 B/op     43 allocs/op
      godb:     0.29s       286770 ns/op    4112 B/op    102 allocs/op
      xorm:     0.29s       288508 ns/op    4616 B/op    125 allocs/op
     upper:     0.31s       309299 ns/op    8137 B/op    315 allocs/op

  1000 times - MultiRead limit 100
    reform:     0.13s       132270 ns/op    2920 B/op     74 allocs/op
       pgx:     0.23s       234005 ns/op   42949 B/op    504 allocs/op
  pgx_pool:     0.23s       234779 ns/op   43009 B/op    504 allocs/op
     upper:     0.30s       303615 ns/op    7849 B/op    294 allocs/op
       raw:     0.31s       309391 ns/op   38342 B/op   1038 allocs/op
      sqlc:     0.35s       346744 ns/op   73158 B/op   1251 allocs/op
        pg:     0.35s       350127 ns/op   22258 B/op    629 allocs/op
 sqlboiler:     0.36s       362392 ns/op   57898 B/op   1259 allocs/op
      gorp:     0.40s       402788 ns/op   57371 B/op   1494 allocs/op
      sqlx:     0.41s       405672 ns/op   39064 B/op   1516 allocs/op
       ent:     0.42s       417560 ns/op   76049 B/op   2039 allocs/op
       dbr:     0.44s       437355 ns/op   32416 B/op   1545 allocs/op
       bun:     0.45s       445493 ns/op   33994 B/op   1124 allocs/op
     beego:     0.46s       457492 ns/op   55183 B/op   3077 allocs/op
       pop:     0.48s       483316 ns/op   75763 B/op   1510 allocs/op
 gorm_prep:     0.50s       495146 ns/op   56043 B/op   1977 allocs/op
      gorm:     0.55s       554714 ns/op   57153 B/op   2277 allocs/op
       rel:     0.66s       656490 ns/op   95321 B/op   2248 allocs/op
      godb:     0.70s       697395 ns/op   75262 B/op   3084 allocs/op
      xorm:     0.76s       756660 ns/op  119957 B/op   4687 allocs/op
```

- orm-benchmark -multi=10

```
  
Reports:

  2000 times - Insert
  pgx_pool:     0.43s       216241 ns/op     285 B/op     10 allocs/op
       raw:     0.43s       216417 ns/op     734 B/op     13 allocs/op
     beego:     0.47s       234975 ns/op    2375 B/op     56 allocs/op
 sqlboiler:     0.48s       240622 ns/op    1592 B/op     35 allocs/op
      sqlc:     0.49s       243949 ns/op    2900 B/op     63 allocs/op
       pgx:     0.50s       248813 ns/op     269 B/op     10 allocs/op
      gorp:     0.51s       254487 ns/op    1798 B/op     42 allocs/op
        pg:     0.53s       264437 ns/op    1342 B/op     10 allocs/op
       ent:     0.54s       268937 ns/op    4215 B/op    100 allocs/op
    reform:     0.56s       279793 ns/op    1807 B/op     51 allocs/op
      gorm:     0.57s       286218 ns/op    6934 B/op     90 allocs/op
       bun:     0.58s       287502 ns/op    5023 B/op     14 allocs/op
       dbr:     0.58s       290459 ns/op    2699 B/op     65 allocs/op
 gorm_prep:     0.62s       311386 ns/op    5291 B/op     67 allocs/op
      sqlx:     0.71s       355447 ns/op     856 B/op     19 allocs/op
       rel:     0.79s       395452 ns/op    2504 B/op     42 allocs/op
      godb:     0.81s       405148 ns/op    4564 B/op    116 allocs/op
      xorm:     0.81s       406552 ns/op    3332 B/op     89 allocs/op
     upper:     0.95s       476759 ns/op   13799 B/op    674 allocs/op
       pop:     1.16s       578073 ns/op   10086 B/op    249 allocs/op

  2000 times - MultiInsert 100 row
  pgx_pool:     1.89s       946469 ns/op  114302 B/op     42 allocs/op
       pgx:     1.92s       959011 ns/op  114299 B/op     43 allocs/op
       raw:     2.09s      1046553 ns/op  191315 B/op    931 allocs/op
     beego:     2.37s      1184271 ns/op  179413 B/op   2746 allocs/op
    reform:     2.44s      1218304 ns/op  466319 B/op   2747 allocs/op
 gorm_prep:     2.54s      1270746 ns/op  234454 B/op   2183 allocs/op
       ent:     2.91s      1453532 ns/op  412140 B/op   4900 allocs/op
        pg:     3.03s      1516890 ns/op    3876 B/op    112 allocs/op
       bun:     3.40s      1697773 ns/op   42245 B/op    218 allocs/op
      sqlx:     3.46s      1732247 ns/op  171228 B/op   1552 allocs/op
      gorm:     3.69s      1842554 ns/op  272240 B/op   3729 allocs/op
      xorm:     4.12s      2061389 ns/op  254764 B/op   5416 allocs/op
       rel:     4.14s      2071836 ns/op  303994 B/op   3263 allocs/op
      godb:     4.33s      2163328 ns/op  259960 B/op   5895 allocs/op
     upper:     5.70s      2851497 ns/op  545255 B/op  19480 allocs/op
      gorp:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert

  2000 times - Update
       raw:     0.23s       115995 ns/op     748 B/op     13 allocs/op
       ent:     0.27s       137096 ns/op    4556 B/op     99 allocs/op
       pgx:     0.44s       217584 ns/op     268 B/op     10 allocs/op
  pgx_pool:     0.45s       222511 ns/op     285 B/op     10 allocs/op
      gorp:     0.46s       229802 ns/op    1205 B/op     32 allocs/op
     beego:     0.47s       234494 ns/op    1752 B/op     47 allocs/op
    reform:     0.48s       237514 ns/op    1772 B/op     51 allocs/op
      sqlc:     0.49s       247311 ns/op     876 B/op     14 allocs/op
 sqlboiler:     0.50s       251777 ns/op     901 B/op     17 allocs/op
 gorm_prep:     0.51s       253345 ns/op    5121 B/op     59 allocs/op
      sqlx:     0.52s       260902 ns/op     872 B/op     20 allocs/op
        pg:     0.56s       278810 ns/op     768 B/op      9 allocs/op
       dbr:     0.57s       282518 ns/op    2651 B/op     57 allocs/op
       pop:     0.57s       287447 ns/op    6582 B/op    198 allocs/op
      gorm:     0.64s       319793 ns/op    6604 B/op     81 allocs/op
       bun:     0.68s       340800 ns/op    4728 B/op      5 allocs/op
       rel:     0.80s       400585 ns/op    2528 B/op     41 allocs/op
      xorm:     0.81s       404093 ns/op    3648 B/op    126 allocs/op
      godb:     0.86s       429592 ns/op    5160 B/op    154 allocs/op
     upper:     2.15s      1072735 ns/op   33507 B/op   1502 allocs/op

  2000 times - Read
       pgx:     0.24s       118492 ns/op     892 B/op      8 allocs/op
  pgx_pool:     0.24s       118818 ns/op    1080 B/op      9 allocs/op
       raw:     0.24s       122280 ns/op    2065 B/op     50 allocs/op
      sqlc:     0.25s       124706 ns/op    2172 B/op     52 allocs/op
     beego:     0.26s       127523 ns/op    2088 B/op     75 allocs/op
    reform:     0.26s       132429 ns/op    3197 B/op     86 allocs/op
      gorp:     0.27s       132709 ns/op    3876 B/op    194 allocs/op
       pop:     0.28s       138041 ns/op    3529 B/op     71 allocs/op
 gorm_prep:     0.28s       139158 ns/op    4800 B/op     83 allocs/op
        pg:     0.28s       141490 ns/op     872 B/op     20 allocs/op
       ent:     0.30s       148844 ns/op    5405 B/op    149 allocs/op
 sqlboiler:     0.30s       149486 ns/op     975 B/op     14 allocs/op
       rel:     0.31s       152810 ns/op    1800 B/op     45 allocs/op
       bun:     0.32s       158530 ns/op    5489 B/op     21 allocs/op
       dbr:     0.33s       163189 ns/op    2184 B/op     37 allocs/op
      gorm:     0.37s       186331 ns/op    5151 B/op     94 allocs/op
      sqlx:     0.55s       274155 ns/op    1976 B/op     43 allocs/op
      godb:     0.58s       290335 ns/op    4112 B/op    102 allocs/op
      xorm:     0.59s       292970 ns/op    4616 B/op    125 allocs/op
     upper:     0.60s       301891 ns/op    8135 B/op    315 allocs/op

  2000 times - MultiRead limit 100
    reform:     0.25s       127005 ns/op    2924 B/op     74 allocs/op
       pgx:     0.46s       231790 ns/op   42949 B/op    504 allocs/op
  pgx_pool:     0.47s       234148 ns/op   43008 B/op    504 allocs/op
     upper:     0.58s       291408 ns/op    7841 B/op    294 allocs/op
       raw:     0.62s       309593 ns/op   38342 B/op   1038 allocs/op
      sqlc:     0.68s       341501 ns/op   73158 B/op   1251 allocs/op
        pg:     0.70s       349068 ns/op   22259 B/op    629 allocs/op
 sqlboiler:     0.75s       374728 ns/op   57881 B/op   1259 allocs/op
      sqlx:     0.79s       394285 ns/op   39064 B/op   1516 allocs/op
       ent:     0.81s       405558 ns/op   76049 B/op   2039 allocs/op
      gorp:     0.83s       413930 ns/op   57372 B/op   1494 allocs/op
       dbr:     0.88s       438300 ns/op   32416 B/op   1545 allocs/op
     beego:     0.92s       460907 ns/op   55182 B/op   3077 allocs/op
       bun:     0.95s       476136 ns/op   33995 B/op   1124 allocs/op
       pop:     0.97s       485463 ns/op   75734 B/op   1510 allocs/op
 gorm_prep:     1.00s       499220 ns/op   56056 B/op   1978 allocs/op
      gorm:     1.14s       569750 ns/op   57158 B/op   2277 allocs/op
       rel:     1.38s       688413 ns/op   95321 B/op   2248 allocs/op
      godb:     1.41s       704437 ns/op   75263 B/op   3084 allocs/op
      xorm:     1.52s       761512 ns/op  119951 B/op   4687 allocs/op
```

- orm-benchmark -multi=20

```
  
Reports:

  4000 times - Insert
  pgx_pool:     0.89s       222688 ns/op     286 B/op     10 allocs/op
       raw:     0.90s       224851 ns/op     719 B/op     13 allocs/op
       pgx:     0.92s       230481 ns/op     271 B/op     10 allocs/op
      gorp:     0.95s       238310 ns/op    1782 B/op     42 allocs/op
 sqlboiler:     0.97s       243744 ns/op    1574 B/op     35 allocs/op
      sqlc:     0.98s       244684 ns/op    2884 B/op     63 allocs/op
     beego:     1.00s       250561 ns/op    2375 B/op     56 allocs/op
       ent:     1.03s       257872 ns/op    4216 B/op    100 allocs/op
 gorm_prep:     1.08s       269052 ns/op    5267 B/op     67 allocs/op
        pg:     1.09s       273017 ns/op     805 B/op     10 allocs/op
       dbr:     1.10s       275689 ns/op    2693 B/op     65 allocs/op
    reform:     1.14s       286242 ns/op    1790 B/op     51 allocs/op
       bun:     1.15s       287451 ns/op    5007 B/op     14 allocs/op
      gorm:     1.23s       308227 ns/op    6930 B/op     89 allocs/op
      sqlx:     1.51s       376436 ns/op     856 B/op     19 allocs/op
       rel:     1.56s       389413 ns/op    2504 B/op     42 allocs/op
      xorm:     1.68s       418873 ns/op    3326 B/op     89 allocs/op
      godb:     1.69s       421799 ns/op    4560 B/op    116 allocs/op
     upper:     1.92s       480286 ns/op   13775 B/op    673 allocs/op
       pop:     2.33s       581746 ns/op   10072 B/op    249 allocs/op

  4000 times - MultiInsert 100 row
       pgx:     3.61s       901597 ns/op  114277 B/op     42 allocs/op
  pgx_pool:     3.61s       901898 ns/op  114312 B/op     43 allocs/op
       raw:     4.00s      1000960 ns/op  191258 B/op    931 allocs/op
     beego:     4.24s      1059653 ns/op  179408 B/op   2746 allocs/op
    reform:     4.54s      1135854 ns/op  466281 B/op   2747 allocs/op
 gorm_prep:     4.81s      1201284 ns/op  234443 B/op   2183 allocs/op
       ent:     5.63s      1407899 ns/op  412125 B/op   4900 allocs/op
        pg:     5.88s      1470050 ns/op    3857 B/op    112 allocs/op
       bun:     6.42s      1604394 ns/op   42241 B/op    219 allocs/op
      sqlx:     6.70s      1675216 ns/op  171203 B/op   1552 allocs/op
      gorm:     7.07s      1768175 ns/op  272257 B/op   3729 allocs/op
       rel:     7.94s      1984215 ns/op  303995 B/op   3263 allocs/op
      xorm:     8.03s      2008628 ns/op  254734 B/op   5416 allocs/op
      godb:     8.23s      2058026 ns/op  259965 B/op   5895 allocs/op
     upper:    10.95s      2738504 ns/op  545161 B/op  19473 allocs/op
 sqlboiler:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert

  4000 times - Update
       raw:     0.48s       118807 ns/op     750 B/op     13 allocs/op
       ent:     0.54s       135127 ns/op    4558 B/op     99 allocs/op
  pgx_pool:     0.92s       228848 ns/op     286 B/op     10 allocs/op
      gorp:     0.93s       232755 ns/op    1204 B/op     32 allocs/op
       pgx:     0.94s       236040 ns/op     270 B/op     10 allocs/op
      sqlc:     0.95s       238700 ns/op     878 B/op     14 allocs/op
     beego:     0.98s       245145 ns/op    1752 B/op     47 allocs/op
    reform:     1.00s       250162 ns/op    1774 B/op     51 allocs/op
 sqlboiler:     1.02s       253826 ns/op     901 B/op     17 allocs/op
 gorm_prep:     1.03s       257080 ns/op    5116 B/op     59 allocs/op
      sqlx:     1.05s       261722 ns/op     872 B/op     20 allocs/op
       pop:     1.09s       271800 ns/op    6582 B/op    198 allocs/op
       dbr:     1.11s       276575 ns/op    2651 B/op     57 allocs/op
        pg:     1.12s       280423 ns/op     768 B/op      9 allocs/op
       bun:     1.17s       293746 ns/op    4728 B/op      5 allocs/op
      gorm:     1.25s       313309 ns/op    6604 B/op     81 allocs/op
       rel:     1.64s       410029 ns/op    2528 B/op     41 allocs/op
      xorm:     1.67s       417757 ns/op    3648 B/op    126 allocs/op
      godb:     1.72s       429225 ns/op    5160 B/op    154 allocs/op
     upper:     4.27s      1068056 ns/op   33484 B/op   1502 allocs/op

  4000 times - Read
       pgx:     0.46s       113961 ns/op     894 B/op      8 allocs/op
  pgx_pool:     0.46s       115924 ns/op    1079 B/op      9 allocs/op
       raw:     0.48s       119726 ns/op    2062 B/op     50 allocs/op
      sqlc:     0.49s       121348 ns/op    2172 B/op     52 allocs/op
     beego:     0.52s       128790 ns/op    2088 B/op     75 allocs/op
    reform:     0.53s       133212 ns/op    3198 B/op     86 allocs/op
      gorp:     0.54s       134824 ns/op    3878 B/op    194 allocs/op
       pop:     0.54s       135684 ns/op    3506 B/op     71 allocs/op
 gorm_prep:     0.57s       141642 ns/op    4801 B/op     83 allocs/op
        pg:     0.58s       145653 ns/op    1134 B/op     20 allocs/op
       ent:     0.59s       147981 ns/op    5404 B/op    149 allocs/op
       bun:     0.60s       150443 ns/op    5489 B/op     21 allocs/op
       dbr:     0.62s       153983 ns/op    2184 B/op     37 allocs/op
       rel:     0.62s       154865 ns/op    1800 B/op     45 allocs/op
 sqlboiler:     0.62s       155469 ns/op     956 B/op     14 allocs/op
      gorm:     0.76s       189176 ns/op    5148 B/op     94 allocs/op
      sqlx:     1.09s       272016 ns/op    1976 B/op     43 allocs/op
      godb:     1.17s       292161 ns/op    4112 B/op    102 allocs/op
      xorm:     1.18s       295009 ns/op    4616 B/op    125 allocs/op
     upper:     1.23s       306743 ns/op    8136 B/op    315 allocs/op

  4000 times - MultiRead limit 100
    reform:     0.51s       128174 ns/op    2926 B/op     74 allocs/op
       pgx:     0.94s       233999 ns/op   42949 B/op    504 allocs/op
  pgx_pool:     0.95s       237544 ns/op   43007 B/op    504 allocs/op
     upper:     1.21s       301665 ns/op    7841 B/op    294 allocs/op
       raw:     1.21s       303330 ns/op   38341 B/op   1038 allocs/op
      sqlc:     1.34s       335415 ns/op   73158 B/op   1251 allocs/op
        pg:     1.42s       353978 ns/op   22259 B/op    629 allocs/op
 sqlboiler:     1.45s       363300 ns/op   57835 B/op   1259 allocs/op
       ent:     1.59s       397582 ns/op   76045 B/op   2039 allocs/op
      sqlx:     1.59s       398497 ns/op   39064 B/op   1516 allocs/op
      gorp:     1.63s       406337 ns/op   57372 B/op   1494 allocs/op
       dbr:     1.75s       437268 ns/op   32416 B/op   1545 allocs/op
       bun:     1.82s       454474 ns/op   33993 B/op   1124 allocs/op
     beego:     1.85s       461376 ns/op   55181 B/op   3077 allocs/op
       pop:     1.92s       479485 ns/op   75715 B/op   1510 allocs/op
 gorm_prep:     2.00s       499800 ns/op   56044 B/op   1977 allocs/op
      gorm:     2.20s       550771 ns/op   57149 B/op   2277 allocs/op
       rel:     2.60s       650333 ns/op   95321 B/op   2248 allocs/op
      godb:     2.78s       694394 ns/op   75264 B/op   3084 allocs/op
      xorm:     3.03s       758109 ns/op  119947 B/op   4687 allocs/op
```
