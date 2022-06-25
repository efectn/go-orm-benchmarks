# Results

- orm-benchmark (with no flags)

```
   
Reports:

   200 times - Insert
  pgx_pool:     0.09s       455016 ns/op     266 B/op     10 allocs/op
       pgx:     0.10s       497652 ns/op     244 B/op     10 allocs/op
     beego:     0.10s       516980 ns/op    2374 B/op     56 allocs/op
      sqlc:     0.10s       520920 ns/op    3179 B/op     64 allocs/op
       raw:     0.11s       534650 ns/op    1006 B/op     15 allocs/op
    reform:     0.11s       538550 ns/op    2093 B/op     53 allocs/op
       dbr:     0.11s       557703 ns/op    2798 B/op     66 allocs/op
 sqlboiler:     0.11s       563488 ns/op    1873 B/op     37 allocs/op
      gorp:     0.11s       567749 ns/op    2076 B/op     43 allocs/op
 gorm_prep:     0.12s       578364 ns/op    5566 B/op     70 allocs/op
       ent:     0.12s       582350 ns/op    4172 B/op     99 allocs/op
       bun:     0.12s       591500 ns/op    5294 B/op     14 allocs/op
        pg:     0.12s       595724 ns/op    6275 B/op     11 allocs/op
      gorm:     0.12s       597391 ns/op    7182 B/op     90 allocs/op
      sqlx:     0.14s       703030 ns/op     856 B/op     19 allocs/op
       rel:     0.15s       769571 ns/op    2511 B/op     41 allocs/op
      xorm:     0.16s       783396 ns/op    3442 B/op     90 allocs/op
      godb:     0.17s       850648 ns/op    4671 B/op    116 allocs/op
     upper:     0.18s       893935 ns/op   14205 B/op    685 allocs/op
       pop:     0.21s      1054372 ns/op   10436 B/op    250 allocs/op

   200 times - MultiInsert 100 row
  pgx_pool:     0.30s      1478320 ns/op  114405 B/op     43 allocs/op
       pgx:     0.31s      1549298 ns/op  114333 B/op     43 allocs/op
       raw:     0.32s      1601169 ns/op  192354 B/op    938 allocs/op
     beego:     0.32s      1614287 ns/op  179519 B/op   2746 allocs/op
    reform:     0.36s      1783033 ns/op  466438 B/op   2747 allocs/op
       ent:     0.39s      1965623 ns/op  412422 B/op   4902 allocs/op
 gorm_prep:     0.44s      2206292 ns/op  235336 B/op   2281 allocs/op
       bun:     0.48s      2384672 ns/op   42356 B/op    217 allocs/op
        pg:     0.48s      2403116 ns/op    8929 B/op    112 allocs/op
      gorm:     0.48s      2422925 ns/op  272533 B/op   3827 allocs/op
      sqlx:     0.52s      2598596 ns/op  171536 B/op   1552 allocs/op
      xorm:     0.53s      2665495 ns/op  254806 B/op   5416 allocs/op
       rel:     0.55s      2737861 ns/op  303974 B/op   3261 allocs/op
      godb:     0.60s      2988488 ns/op  259944 B/op   5893 allocs/op
     upper:     0.91s      4549805 ns/op  547001 B/op  19610 allocs/op
      sqlc:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert

   200 times - Update
       raw:     0.04s       192461 ns/op     721 B/op     13 allocs/op
       ent:     0.04s       216585 ns/op    4529 B/op     99 allocs/op
      sqlx:     0.09s       431393 ns/op     872 B/op     20 allocs/op
       pgx:     0.09s       463765 ns/op     282 B/op     10 allocs/op
    reform:     0.09s       468075 ns/op    1744 B/op     51 allocs/op
 sqlboiler:     0.10s       492893 ns/op     924 B/op     17 allocs/op
  pgx_pool:     0.10s       498383 ns/op     302 B/op     10 allocs/op
       pop:     0.10s       503204 ns/op    6559 B/op    198 allocs/op
      sqlc:     0.10s       506504 ns/op     891 B/op     14 allocs/op
      gorp:     0.10s       507412 ns/op    1221 B/op     32 allocs/op
       dbr:     0.10s       523120 ns/op    2651 B/op     57 allocs/op
     beego:     0.11s       527907 ns/op    1755 B/op     47 allocs/op
 gorm_prep:     0.11s       551450 ns/op    5132 B/op     59 allocs/op
        pg:     0.12s       584310 ns/op     768 B/op      9 allocs/op
      gorm:     0.12s       588737 ns/op    6624 B/op     81 allocs/op
       bun:     0.12s       595931 ns/op    4728 B/op      5 allocs/op
      xorm:     0.15s       771985 ns/op    3648 B/op    126 allocs/op
       rel:     0.16s       804915 ns/op    2531 B/op     41 allocs/op
      godb:     0.17s       838846 ns/op    5160 B/op    154 allocs/op
     upper:     0.37s      1858703 ns/op   33648 B/op   1510 allocs/op

   200 times - Read
       pgx:     0.04s       191837 ns/op     886 B/op      8 allocs/op
  pgx_pool:     0.04s       195595 ns/op    1078 B/op      9 allocs/op
      sqlc:     0.04s       198668 ns/op    2162 B/op     52 allocs/op
     beego:     0.04s       199401 ns/op    2094 B/op     75 allocs/op
       raw:     0.04s       224784 ns/op    2100 B/op     50 allocs/op
       dbr:     0.05s       226916 ns/op    2184 B/op     37 allocs/op
    reform:     0.05s       229498 ns/op    3232 B/op     86 allocs/op
        pg:     0.05s       231312 ns/op     873 B/op     20 allocs/op
 sqlboiler:     0.05s       239986 ns/op    1303 B/op     14 allocs/op
       pop:     0.05s       245367 ns/op    3914 B/op     71 allocs/op
      gorp:     0.05s       247669 ns/op    3869 B/op    194 allocs/op
 gorm_prep:     0.05s       247760 ns/op    4800 B/op     84 allocs/op
       rel:     0.05s       251045 ns/op    1803 B/op     45 allocs/op
       ent:     0.06s       279769 ns/op    5424 B/op    149 allocs/op
       bun:     0.06s       290865 ns/op    5490 B/op     21 allocs/op
      gorm:     0.06s       294879 ns/op    5157 B/op     95 allocs/op
      sqlx:     0.08s       416614 ns/op    1976 B/op     43 allocs/op
      godb:     0.09s       442079 ns/op    4112 B/op    102 allocs/op
     upper:     0.10s       478927 ns/op    8126 B/op    315 allocs/op
      xorm:     0.11s       572655 ns/op    4616 B/op    125 allocs/op

   200 times - MultiRead limit 100
    reform:     0.04s       216676 ns/op    2956 B/op     74 allocs/op
  pgx_pool:     0.08s       394447 ns/op   42993 B/op    504 allocs/op
       raw:     0.10s       487462 ns/op   38352 B/op   1038 allocs/op
     upper:     0.10s       498045 ns/op    7896 B/op    294 allocs/op
 sqlboiler:     0.10s       498234 ns/op   57841 B/op   1259 allocs/op
        pg:     0.10s       517104 ns/op   22258 B/op    629 allocs/op
      sqlx:     0.11s       543746 ns/op   39064 B/op   1516 allocs/op
       ent:     0.11s       544689 ns/op   76035 B/op   2039 allocs/op
      gorp:     0.11s       553690 ns/op   57386 B/op   1494 allocs/op
       pgx:     0.12s       575743 ns/op   42919 B/op    504 allocs/op
     beego:     0.13s       662287 ns/op   55195 B/op   3077 allocs/op
      sqlc:     0.13s       665221 ns/op   73128 B/op   1251 allocs/op
       bun:     0.13s       671196 ns/op   33994 B/op   1124 allocs/op
       pop:     0.14s       675096 ns/op   75980 B/op   1511 allocs/op
       dbr:     0.14s       699685 ns/op   32416 B/op   1545 allocs/op
 gorm_prep:     0.16s       811506 ns/op   56077 B/op   2078 allocs/op
      gorm:     0.17s       844117 ns/op   57201 B/op   2378 allocs/op
      godb:     0.19s       947518 ns/op   75263 B/op   3084 allocs/op
       rel:     0.20s      1024612 ns/op   95321 B/op   2248 allocs/op
      xorm:     0.24s      1216746 ns/op  119974 B/op   4687 allocs/op
```

- orm-benchmark -multi=5

```
  
Reports:

  1000 times - Insert
       pgx:     0.48s       477980 ns/op     265 B/op     10 allocs/op
       raw:     0.49s       485519 ns/op     764 B/op     13 allocs/op
  pgx_pool:     0.49s       486861 ns/op     281 B/op     10 allocs/op
      gorp:     0.50s       500670 ns/op    1827 B/op     42 allocs/op
     beego:     0.52s       518605 ns/op    2375 B/op     56 allocs/op
      sqlc:     0.53s       531917 ns/op    2929 B/op     63 allocs/op
       dbr:     0.53s       533002 ns/op    2710 B/op     65 allocs/op
       ent:     0.55s       549147 ns/op    4214 B/op    100 allocs/op
 sqlboiler:     0.56s       558613 ns/op    1620 B/op     35 allocs/op
        pg:     0.56s       558744 ns/op    1890 B/op     10 allocs/op
 gorm_prep:     0.56s       559330 ns/op    5321 B/op     68 allocs/op
    reform:     0.57s       565211 ns/op    1839 B/op     51 allocs/op
       bun:     0.61s       606268 ns/op    5053 B/op     14 allocs/op
      gorm:     0.62s       623248 ns/op    6937 B/op     90 allocs/op
      sqlx:     0.70s       695930 ns/op     856 B/op     19 allocs/op
      xorm:     0.76s       761621 ns/op    3351 B/op     89 allocs/op
      godb:     0.76s       764358 ns/op    4577 B/op    116 allocs/op
       rel:     0.79s       786221 ns/op    2504 B/op     42 allocs/op
     upper:     0.87s       867042 ns/op   13841 B/op    675 allocs/op
       pop:     1.03s      1030258 ns/op   10121 B/op    249 allocs/op

  1000 times - MultiInsert 100 row
       pgx:     1.51s      1509487 ns/op  114335 B/op     43 allocs/op
  pgx_pool:     1.54s      1540298 ns/op  114370 B/op     43 allocs/op
       raw:     1.58s      1581121 ns/op  191429 B/op    932 allocs/op
     beego:     1.63s      1633548 ns/op  179435 B/op   2746 allocs/op
    reform:     1.70s      1701732 ns/op  466351 B/op   2747 allocs/op
 gorm_prep:     1.72s      1715955 ns/op  235265 B/op   2283 allocs/op
       ent:     2.01s      2014722 ns/op  412167 B/op   4900 allocs/op
        pg:     2.02s      2018604 ns/op    4364 B/op    112 allocs/op
      gorm:     2.33s      2331972 ns/op  272894 B/op   3829 allocs/op
       bun:     2.35s      2348216 ns/op   42269 B/op    218 allocs/op
      sqlx:     2.36s      2358129 ns/op  171259 B/op   1552 allocs/op
      xorm:     2.62s      2621344 ns/op  254757 B/op   5416 allocs/op
       rel:     2.72s      2724592 ns/op  303991 B/op   3263 allocs/op
      godb:     2.79s      2792383 ns/op  259970 B/op   5894 allocs/op
     upper:     3.62s      3620745 ns/op  545452 B/op  19495 allocs/op
      gorp:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert

  1000 times - Update
       raw:     0.18s       177471 ns/op     752 B/op     13 allocs/op
       ent:     0.22s       220526 ns/op    4561 B/op     99 allocs/op
      sqlx:     0.41s       409202 ns/op     872 B/op     20 allocs/op
  pgx_pool:     0.48s       476851 ns/op     287 B/op     10 allocs/op
      sqlc:     0.49s       486617 ns/op     880 B/op     14 allocs/op
       pgx:     0.49s       488426 ns/op     272 B/op     10 allocs/op
     beego:     0.49s       491010 ns/op    1753 B/op     47 allocs/op
      gorp:     0.50s       500261 ns/op    1209 B/op     32 allocs/op
    reform:     0.51s       507634 ns/op    1777 B/op     51 allocs/op
 sqlboiler:     0.52s       516117 ns/op     907 B/op     17 allocs/op
 gorm_prep:     0.55s       551332 ns/op    5113 B/op     59 allocs/op
       bun:     0.56s       560118 ns/op    4728 B/op      5 allocs/op
       dbr:     0.57s       566273 ns/op    2651 B/op     57 allocs/op
       pop:     0.59s       590671 ns/op    6587 B/op    198 allocs/op
      gorm:     0.63s       628642 ns/op    6600 B/op     81 allocs/op
        pg:     0.66s       658536 ns/op     768 B/op      9 allocs/op
      xorm:     0.78s       780272 ns/op    3648 B/op    126 allocs/op
      godb:     0.80s       802017 ns/op    5160 B/op    154 allocs/op
       rel:     0.84s       842742 ns/op    2528 B/op     41 allocs/op
     upper:     1.77s      1770174 ns/op   33518 B/op   1503 allocs/op

  1000 times - Read
       raw:     0.19s       187383 ns/op    2057 B/op     50 allocs/op
    reform:     0.19s       192739 ns/op    3193 B/op     86 allocs/op
     beego:     0.19s       194843 ns/op    2089 B/op     75 allocs/op
      sqlc:     0.20s       198480 ns/op    2176 B/op     52 allocs/op
       pgx:     0.20s       204291 ns/op     889 B/op      8 allocs/op
      gorp:     0.22s       219710 ns/op    3881 B/op    194 allocs/op
  pgx_pool:     0.22s       223166 ns/op    1076 B/op      9 allocs/op
       dbr:     0.22s       224556 ns/op    2184 B/op     37 allocs/op
       pop:     0.23s       226025 ns/op    3561 B/op     71 allocs/op
       ent:     0.23s       232831 ns/op    5401 B/op    149 allocs/op
       rel:     0.24s       236573 ns/op    1800 B/op     45 allocs/op
       bun:     0.24s       237772 ns/op    5489 B/op     21 allocs/op
 sqlboiler:     0.24s       244055 ns/op    1008 B/op     14 allocs/op
        pg:     0.25s       253588 ns/op     872 B/op     20 allocs/op
 gorm_prep:     0.27s       273435 ns/op    4807 B/op     84 allocs/op
      gorm:     0.30s       299989 ns/op    5156 B/op     95 allocs/op
      sqlx:     0.41s       406178 ns/op    1976 B/op     43 allocs/op
      godb:     0.43s       430804 ns/op    4112 B/op    102 allocs/op
      xorm:     0.48s       475540 ns/op    4616 B/op    125 allocs/op
     upper:     0.50s       499812 ns/op    8137 B/op    315 allocs/op

  1000 times - MultiRead limit 100
    reform:     0.19s       191517 ns/op    2920 B/op     74 allocs/op
       pgx:     0.39s       389510 ns/op   42949 B/op    504 allocs/op
  pgx_pool:     0.39s       393637 ns/op   43009 B/op    504 allocs/op
       raw:     0.44s       443016 ns/op   38342 B/op   1038 allocs/op
     upper:     0.46s       459684 ns/op    7849 B/op    294 allocs/op
 sqlboiler:     0.50s       504771 ns/op   57868 B/op   1259 allocs/op
      sqlc:     0.51s       505651 ns/op   73158 B/op   1251 allocs/op
      gorp:     0.52s       524104 ns/op   57373 B/op   1494 allocs/op
       ent:     0.54s       537856 ns/op   76049 B/op   2039 allocs/op
      sqlx:     0.54s       538735 ns/op   39064 B/op   1516 allocs/op
        pg:     0.55s       548573 ns/op   23457 B/op    629 allocs/op
       dbr:     0.58s       576076 ns/op   32416 B/op   1545 allocs/op
     beego:     0.59s       589107 ns/op   55185 B/op   3077 allocs/op
       bun:     0.59s       590483 ns/op   33994 B/op   1124 allocs/op
       pop:     0.67s       667911 ns/op   75754 B/op   1510 allocs/op
 gorm_prep:     0.74s       744623 ns/op   56071 B/op   2079 allocs/op
      gorm:     0.79s       791774 ns/op   57173 B/op   2378 allocs/op
       rel:     0.90s       903076 ns/op   95322 B/op   2248 allocs/op
      xorm:     0.94s       939495 ns/op  119956 B/op   4687 allocs/op
      godb:     1.02s      1016323 ns/op   75265 B/op   3084 allocs/op
```

- orm-benchmark -multi=10

```
  
Reports:

  2000 times - Insert
       pgx:     0.97s       483168 ns/op     269 B/op     10 allocs/op
       raw:     0.97s       484960 ns/op     734 B/op     13 allocs/op
  pgx_pool:     0.97s       486798 ns/op     285 B/op     10 allocs/op
    reform:     1.01s       505437 ns/op    1807 B/op     51 allocs/op
      gorp:     1.01s       507063 ns/op    1798 B/op     42 allocs/op
     beego:     1.02s       508635 ns/op    2376 B/op     56 allocs/op
 sqlboiler:     1.06s       532215 ns/op    1592 B/op     35 allocs/op
      sqlc:     1.09s       544902 ns/op    2900 B/op     63 allocs/op
       ent:     1.10s       548487 ns/op    4215 B/op    100 allocs/op
 gorm_prep:     1.12s       558009 ns/op    5285 B/op     68 allocs/op
       dbr:     1.12s       559111 ns/op    2699 B/op     65 allocs/op
       bun:     1.13s       565797 ns/op    5023 B/op     14 allocs/op
        pg:     1.17s       583325 ns/op    1341 B/op     10 allocs/op
      gorm:     1.24s       620206 ns/op    6956 B/op     90 allocs/op
      sqlx:     1.47s       734542 ns/op     856 B/op     19 allocs/op
       rel:     1.52s       758334 ns/op    2504 B/op     42 allocs/op
      xorm:     1.55s       773863 ns/op    3332 B/op     89 allocs/op
      godb:     1.55s       776638 ns/op    4564 B/op    116 allocs/op
     upper:     1.73s       863152 ns/op   13801 B/op    674 allocs/op
       pop:     2.18s      1088208 ns/op   10091 B/op    249 allocs/op

  2000 times - MultiInsert 100 row
  pgx_pool:     2.99s      1493826 ns/op  114331 B/op     43 allocs/op
       raw:     3.07s      1533627 ns/op  191315 B/op    931 allocs/op
       pgx:     3.14s      1571276 ns/op  114310 B/op     43 allocs/op
     beego:     3.20s      1599077 ns/op  179421 B/op   2746 allocs/op
    reform:     3.37s      1686366 ns/op  466312 B/op   2747 allocs/op
 gorm_prep:     3.51s      1756092 ns/op  235253 B/op   2283 allocs/op
        pg:     4.14s      2069355 ns/op    3876 B/op    112 allocs/op
       ent:     4.23s      2113017 ns/op  412139 B/op   4900 allocs/op
       bun:     4.25s      2122778 ns/op   42245 B/op    218 allocs/op
      sqlx:     4.71s      2357418 ns/op  171296 B/op   1552 allocs/op
      gorm:     4.76s      2381479 ns/op  272993 B/op   3829 allocs/op
      xorm:     5.19s      2593526 ns/op  254748 B/op   5416 allocs/op
       rel:     5.46s      2728224 ns/op  303994 B/op   3263 allocs/op
      godb:     5.96s      2979905 ns/op  259962 B/op   5895 allocs/op
     upper:     8.22s      4110594 ns/op  545255 B/op  19480 allocs/op
       dbr:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert

  2000 times - Update
       raw:     0.40s       199403 ns/op     748 B/op     13 allocs/op
       ent:     0.44s       221364 ns/op    4556 B/op     99 allocs/op
      sqlx:     0.83s       412583 ns/op     872 B/op     20 allocs/op
      sqlc:     1.00s       498317 ns/op     876 B/op     14 allocs/op
      gorp:     1.00s       501259 ns/op    1205 B/op     32 allocs/op
     beego:     1.01s       503890 ns/op    1752 B/op     47 allocs/op
  pgx_pool:     1.01s       505072 ns/op     285 B/op     10 allocs/op
       pgx:     1.04s       518822 ns/op     268 B/op     10 allocs/op
 sqlboiler:     1.04s       521096 ns/op     901 B/op     17 allocs/op
    reform:     1.07s       533196 ns/op    1772 B/op     51 allocs/op
       dbr:     1.10s       549420 ns/op    2651 B/op     57 allocs/op
 gorm_prep:     1.12s       558055 ns/op    5121 B/op     59 allocs/op
       pop:     1.16s       580923 ns/op    6583 B/op    198 allocs/op
       bun:     1.18s       591538 ns/op    4728 B/op      5 allocs/op
        pg:     1.19s       593182 ns/op     768 B/op      9 allocs/op
      gorm:     1.31s       657367 ns/op    6604 B/op     81 allocs/op
       rel:     1.55s       776084 ns/op    2528 B/op     41 allocs/op
      xorm:     1.59s       793288 ns/op    3648 B/op    126 allocs/op
      godb:     1.67s       835363 ns/op    5160 B/op    154 allocs/op
     upper:     3.61s      1802665 ns/op   33506 B/op   1502 allocs/op

  2000 times - Read
       pgx:     0.38s       189056 ns/op     892 B/op      8 allocs/op
  pgx_pool:     0.38s       192134 ns/op    1080 B/op      9 allocs/op
      sqlc:     0.39s       196650 ns/op    2172 B/op     52 allocs/op
       raw:     0.41s       206953 ns/op    2065 B/op     50 allocs/op
        pg:     0.43s       216762 ns/op     872 B/op     20 allocs/op
     beego:     0.43s       217393 ns/op    2088 B/op     75 allocs/op
    reform:     0.45s       223504 ns/op    3197 B/op     86 allocs/op
      gorp:     0.45s       226929 ns/op    3876 B/op    194 allocs/op
 sqlboiler:     0.46s       229718 ns/op     974 B/op     14 allocs/op
       pop:     0.47s       232617 ns/op    3529 B/op     71 allocs/op
       rel:     0.47s       233970 ns/op    1800 B/op     45 allocs/op
 gorm_prep:     0.48s       240765 ns/op    4802 B/op     84 allocs/op
       bun:     0.49s       243113 ns/op    5489 B/op     21 allocs/op
       dbr:     0.50s       249554 ns/op    2184 B/op     37 allocs/op
       ent:     0.51s       256562 ns/op    5405 B/op    149 allocs/op
      gorm:     0.63s       313092 ns/op    5152 B/op     95 allocs/op
      sqlx:     0.80s       400164 ns/op    1976 B/op     43 allocs/op
      godb:     0.94s       469499 ns/op    4112 B/op    102 allocs/op
     upper:     0.96s       480212 ns/op    8135 B/op    315 allocs/op
      xorm:     1.05s       526083 ns/op    4616 B/op    125 allocs/op

  2000 times - MultiRead limit 100
    reform:     0.40s       199238 ns/op    2924 B/op     74 allocs/op
  pgx_pool:     0.68s       337768 ns/op   43008 B/op    504 allocs/op
       pgx:     0.73s       363277 ns/op   42949 B/op    504 allocs/op
     upper:     0.94s       470029 ns/op    7841 B/op    294 allocs/op
       raw:     0.98s       487680 ns/op   38342 B/op   1038 allocs/op
      sqlc:     0.98s       491805 ns/op   73158 B/op   1251 allocs/op
        pg:     1.03s       513662 ns/op   23308 B/op    629 allocs/op
 sqlboiler:     1.03s       515087 ns/op   57864 B/op   1259 allocs/op
      gorp:     1.05s       526133 ns/op   57372 B/op   1494 allocs/op
      sqlx:     1.08s       541866 ns/op   39064 B/op   1516 allocs/op
       ent:     1.15s       577245 ns/op   76049 B/op   2039 allocs/op
       dbr:     1.22s       611461 ns/op   32416 B/op   1545 allocs/op
       bun:     1.30s       651370 ns/op   33992 B/op   1124 allocs/op
       pop:     1.34s       670409 ns/op   75743 B/op   1510 allocs/op
     beego:     1.36s       678361 ns/op   55183 B/op   3077 allocs/op
 gorm_prep:     1.43s       717223 ns/op   56070 B/op   2079 allocs/op
      gorm:     1.63s       815156 ns/op   57177 B/op   2378 allocs/op
       rel:     1.75s       872978 ns/op   95321 B/op   2248 allocs/op
      godb:     1.96s       979659 ns/op   75263 B/op   3084 allocs/op
      xorm:     2.04s      1022297 ns/op  119950 B/op   4687 allocs/op
```

- orm-benchmark -multi=20

```
  
Reports:

  4000 times - Insert
       pgx:     1.94s       484887 ns/op     270 B/op     10 allocs/op
  pgx_pool:     1.95s       486289 ns/op     286 B/op     10 allocs/op
       raw:     1.97s       493270 ns/op     719 B/op     13 allocs/op
    reform:     2.01s       502362 ns/op    1790 B/op     51 allocs/op
      gorp:     2.04s       509224 ns/op    1782 B/op     42 allocs/op
 sqlboiler:     2.06s       515997 ns/op    1574 B/op     35 allocs/op
     beego:     2.12s       528797 ns/op    2375 B/op     56 allocs/op
      sqlc:     2.18s       546154 ns/op    2884 B/op     63 allocs/op
       ent:     2.19s       546859 ns/op    4216 B/op    100 allocs/op
        pg:     2.21s       551752 ns/op     805 B/op     10 allocs/op
       dbr:     2.22s       555562 ns/op    2693 B/op     65 allocs/op
 gorm_prep:     2.23s       556508 ns/op    5263 B/op     68 allocs/op
       bun:     2.35s       587903 ns/op    5007 B/op     14 allocs/op
      gorm:     2.55s       638517 ns/op    6949 B/op     90 allocs/op
      sqlx:     2.94s       735400 ns/op     856 B/op     19 allocs/op
       rel:     3.04s       758790 ns/op    2504 B/op     42 allocs/op
      xorm:     3.10s       775444 ns/op    3326 B/op     89 allocs/op
      godb:     3.20s       798888 ns/op    4559 B/op    116 allocs/op
     upper:     3.47s       866731 ns/op   13775 B/op    673 allocs/op
       pop:     4.23s      1058409 ns/op   10074 B/op    249 allocs/op

  4000 times - MultiInsert 100 row
       raw:     5.63s      1407303 ns/op  191258 B/op    931 allocs/op
       pgx:     5.91s      1478677 ns/op  114301 B/op     43 allocs/op
  pgx_pool:     6.05s      1513618 ns/op  114316 B/op     43 allocs/op
    reform:     6.33s      1583119 ns/op  466279 B/op   2747 allocs/op
     beego:     6.39s      1597640 ns/op  179422 B/op   2746 allocs/op
 gorm_prep:     6.71s      1677490 ns/op  235237 B/op   2283 allocs/op
       ent:     7.51s      1876295 ns/op  412126 B/op   4900 allocs/op
        pg:     8.17s      2042350 ns/op    3576 B/op    112 allocs/op
       bun:     8.21s      2052788 ns/op   42241 B/op    218 allocs/op
      sqlx:     9.49s      2371302 ns/op  171287 B/op   1552 allocs/op
      gorm:     9.66s      2414539 ns/op  273037 B/op   3829 allocs/op
      xorm:    10.53s      2632170 ns/op  254745 B/op   5416 allocs/op
       rel:    10.54s      2635492 ns/op  303993 B/op   3263 allocs/op
      godb:    11.38s      2845574 ns/op  259959 B/op   5895 allocs/op
     upper:    15.19s      3796340 ns/op  545162 B/op  19473 allocs/op
       pop:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert

  4000 times - Update
       raw:     0.78s       194482 ns/op     750 B/op     13 allocs/op
       ent:     0.87s       216761 ns/op    4558 B/op     99 allocs/op
      sqlx:     1.68s       420810 ns/op     872 B/op     20 allocs/op
      sqlc:     2.04s       508914 ns/op     878 B/op     14 allocs/op
       pgx:     2.06s       514874 ns/op     270 B/op     10 allocs/op
    reform:     2.07s       517736 ns/op    1774 B/op     51 allocs/op
  pgx_pool:     2.08s       519679 ns/op     286 B/op     10 allocs/op
 sqlboiler:     2.08s       520130 ns/op     901 B/op     17 allocs/op
      gorp:     2.08s       520296 ns/op    1204 B/op     32 allocs/op
     beego:     2.13s       533340 ns/op    1752 B/op     47 allocs/op
 gorm_prep:     2.15s       537928 ns/op    5116 B/op     59 allocs/op
       dbr:     2.28s       570677 ns/op    2651 B/op     57 allocs/op
       bun:     2.39s       597576 ns/op    4728 B/op      5 allocs/op
       pop:     2.40s       600643 ns/op    6582 B/op    198 allocs/op
        pg:     2.49s       621997 ns/op     768 B/op      9 allocs/op
      gorm:     2.50s       624973 ns/op    6604 B/op     81 allocs/op
      xorm:     3.16s       789416 ns/op    3648 B/op    126 allocs/op
       rel:     3.17s       793382 ns/op    2528 B/op     41 allocs/op
      godb:     3.39s       846618 ns/op    5160 B/op    154 allocs/op
     upper:     7.26s      1814782 ns/op   33485 B/op   1502 allocs/op

  4000 times - Read
  pgx_pool:     0.74s       185435 ns/op    1079 B/op      9 allocs/op
       pgx:     0.77s       191331 ns/op     894 B/op      8 allocs/op
      sqlc:     0.82s       206103 ns/op    2172 B/op     52 allocs/op
       raw:     0.83s       206281 ns/op    2062 B/op     50 allocs/op
    reform:     0.85s       212031 ns/op    3198 B/op     86 allocs/op
     beego:     0.89s       221259 ns/op    2088 B/op     75 allocs/op
      gorp:     0.89s       222971 ns/op    3878 B/op    194 allocs/op
       rel:     0.92s       230867 ns/op    1800 B/op     45 allocs/op
        pg:     0.94s       235636 ns/op     872 B/op     20 allocs/op
       dbr:     0.96s       240492 ns/op    2184 B/op     37 allocs/op
       pop:     0.96s       240630 ns/op    3506 B/op     71 allocs/op
 sqlboiler:     0.99s       246676 ns/op     956 B/op     14 allocs/op
       bun:     0.99s       247371 ns/op    5489 B/op     21 allocs/op
       ent:     1.01s       253451 ns/op    5404 B/op    149 allocs/op
 gorm_prep:     1.02s       255602 ns/op    4801 B/op     84 allocs/op
      gorm:     1.19s       296872 ns/op    5149 B/op     95 allocs/op
      sqlx:     1.73s       431608 ns/op    1976 B/op     43 allocs/op
      godb:     1.87s       467924 ns/op    4112 B/op    102 allocs/op
      xorm:     1.93s       482138 ns/op    4616 B/op    125 allocs/op
     upper:     2.10s       524349 ns/op    8136 B/op    315 allocs/op

  4000 times - MultiRead limit 100
    reform:     0.81s       202942 ns/op    2926 B/op     74 allocs/op
       pgx:     1.50s       374236 ns/op   42949 B/op    504 allocs/op
  pgx_pool:     1.50s       375515 ns/op   43006 B/op    504 allocs/op
       raw:     1.70s       425345 ns/op   38341 B/op   1038 allocs/op
      sqlc:     1.93s       482891 ns/op   73158 B/op   1251 allocs/op
     upper:     2.00s       500493 ns/op    7841 B/op    294 allocs/op
        pg:     2.01s       503264 ns/op   22258 B/op    629 allocs/op
      gorp:     2.08s       521211 ns/op   57370 B/op   1494 allocs/op
       ent:     2.17s       542495 ns/op   76045 B/op   2039 allocs/op
 sqlboiler:     2.24s       561107 ns/op   57832 B/op   1259 allocs/op
      sqlx:     2.27s       568689 ns/op   39064 B/op   1516 allocs/op
       dbr:     2.40s       599724 ns/op   32416 B/op   1545 allocs/op
       bun:     2.59s       646290 ns/op   33990 B/op   1124 allocs/op
     beego:     2.63s       657631 ns/op   55183 B/op   3077 allocs/op
       pop:     2.77s       692130 ns/op   75717 B/op   1510 allocs/op
 gorm_prep:     2.99s       748236 ns/op   56050 B/op   2078 allocs/op
      gorm:     3.09s       771292 ns/op   57170 B/op   2378 allocs/op
       rel:     3.77s       942016 ns/op   95321 B/op   2248 allocs/op
      xorm:     3.84s       960855 ns/op  119947 B/op   4687 allocs/op
      godb:     4.08s      1020626 ns/op   75263 B/op   3084 allocs/op
```
