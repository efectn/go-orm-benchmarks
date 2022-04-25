# Results

- orm-benchmark (with no flags)

```
   
Reports:

   200 times - Insert
  pgx_pool:     0.05s       250104 ns/op     265 B/op     10 allocs/op
       pgx:     0.05s       253503 ns/op     244 B/op     10 allocs/op
     beego:     0.05s       273230 ns/op    2375 B/op     56 allocs/op
       raw:     0.06s       275981 ns/op    1015 B/op     15 allocs/op
      gorp:     0.06s       294510 ns/op    2070 B/op     43 allocs/op
        pg:     0.06s       299660 ns/op    1031 B/op     10 allocs/op
    reform:     0.06s       301884 ns/op    2088 B/op     53 allocs/op
 sqlboiler:     0.06s       303986 ns/op    1870 B/op     37 allocs/op
      sqlc:     0.06s       305574 ns/op    3177 B/op     64 allocs/op
 gorm_prep:     0.06s       323437 ns/op    5604 B/op     70 allocs/op
       dbr:     0.07s       327976 ns/op    2796 B/op     66 allocs/op
      gorm:     0.07s       339496 ns/op    7132 B/op     89 allocs/op
       bun:     0.07s       343051 ns/op    5301 B/op     14 allocs/op
      sqlx:     0.08s       416332 ns/op     856 B/op     19 allocs/op
       rel:     0.09s       442845 ns/op    2511 B/op     41 allocs/op
      godb:     0.10s       477746 ns/op    4667 B/op    116 allocs/op
     upper:     0.11s       528646 ns/op   14203 B/op    685 allocs/op
      xorm:     0.11s       540262 ns/op    3440 B/op     90 allocs/op
       pop:     0.14s       678347 ns/op   10450 B/op    250 allocs/op
       ent:     0.19s       959904 ns/op    4172 B/op     99 allocs/op

   200 times - MultiInsert 100 row
       pgx:     0.20s       999654 ns/op  114319 B/op     43 allocs/op
  pgx_pool:     0.21s      1032573 ns/op  114352 B/op     43 allocs/op
       raw:     0.23s      1156051 ns/op  192354 B/op    938 allocs/op
     beego:     0.25s      1225079 ns/op  179491 B/op   2746 allocs/op
    reform:     0.25s      1250609 ns/op  466442 B/op   2747 allocs/op
 gorm_prep:     0.26s      1298326 ns/op  234542 B/op   2182 allocs/op
       ent:     0.30s      1507730 ns/op  412423 B/op   4902 allocs/op
        pg:     0.32s      1584569 ns/op    8929 B/op    112 allocs/op
       bun:     0.33s      1650940 ns/op   42347 B/op    217 allocs/op
      sqlx:     0.39s      1964896 ns/op  171541 B/op   1552 allocs/op
      gorm:     0.40s      2013239 ns/op  272121 B/op   3728 allocs/op
      godb:     0.44s      2197509 ns/op  259948 B/op   5893 allocs/op
      xorm:     0.45s      2265058 ns/op  254800 B/op   5416 allocs/op
       rel:     0.46s      2277662 ns/op  303975 B/op   3261 allocs/op
     upper:     0.59s      2935289 ns/op  547001 B/op  19610 allocs/op
       dbr:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert

   200 times - Update
       raw:     0.03s       129425 ns/op     721 B/op     13 allocs/op
       ent:     0.03s       140752 ns/op    4530 B/op     99 allocs/op
  pgx_pool:     0.05s       251655 ns/op     302 B/op     10 allocs/op
       pgx:     0.05s       251715 ns/op     283 B/op     10 allocs/op
      sqlc:     0.05s       255633 ns/op     890 B/op     14 allocs/op
    reform:     0.05s       257556 ns/op    1745 B/op     51 allocs/op
     beego:     0.05s       258286 ns/op    1755 B/op     47 allocs/op
 sqlboiler:     0.05s       261129 ns/op     926 B/op     17 allocs/op
      gorp:     0.05s       264121 ns/op    1223 B/op     32 allocs/op
 gorm_prep:     0.06s       276982 ns/op    5133 B/op     59 allocs/op
       dbr:     0.06s       295118 ns/op    2651 B/op     57 allocs/op
      sqlx:     0.06s       303751 ns/op     872 B/op     20 allocs/op
        pg:     0.06s       306270 ns/op     768 B/op      9 allocs/op
       pop:     0.06s       314173 ns/op    6555 B/op    198 allocs/op
       bun:     0.06s       318873 ns/op    4728 B/op      5 allocs/op
      gorm:     0.07s       344122 ns/op    6624 B/op     81 allocs/op
      xorm:     0.09s       431793 ns/op    3648 B/op    126 allocs/op
       rel:     0.09s       443894 ns/op    2531 B/op     41 allocs/op
      godb:     0.10s       480432 ns/op    5160 B/op    154 allocs/op
     upper:     0.24s      1194790 ns/op   33648 B/op   1510 allocs/op

   200 times - Read
       pgx:     0.03s       125939 ns/op     886 B/op      8 allocs/op
  pgx_pool:     0.03s       126689 ns/op    1078 B/op      9 allocs/op
       raw:     0.03s       135204 ns/op    2100 B/op     50 allocs/op
      sqlc:     0.03s       136790 ns/op    2162 B/op     52 allocs/op
     beego:     0.03s       142463 ns/op    2094 B/op     75 allocs/op
      gorp:     0.03s       143717 ns/op    3869 B/op    194 allocs/op
    reform:     0.03s       147348 ns/op    3232 B/op     86 allocs/op
       pop:     0.03s       150606 ns/op    3916 B/op     71 allocs/op
 gorm_prep:     0.03s       156324 ns/op    4799 B/op     83 allocs/op
       ent:     0.03s       157926 ns/op    5422 B/op    149 allocs/op
       bun:     0.03s       165386 ns/op    5490 B/op     21 allocs/op
 sqlboiler:     0.03s       166963 ns/op    1116 B/op     14 allocs/op
       dbr:     0.03s       168799 ns/op    2184 B/op     37 allocs/op
        pg:     0.03s       169158 ns/op     873 B/op     20 allocs/op
       rel:     0.04s       187481 ns/op    1803 B/op     45 allocs/op
      gorm:     0.04s       199661 ns/op    5152 B/op     94 allocs/op
      sqlx:     0.06s       293597 ns/op    1976 B/op     43 allocs/op
      godb:     0.06s       311149 ns/op    4112 B/op    102 allocs/op
      xorm:     0.07s       330069 ns/op    4616 B/op    125 allocs/op
     upper:     0.07s       332339 ns/op    8125 B/op    315 allocs/op

   200 times - MultiRead limit 100
    reform:     0.03s       135755 ns/op    2955 B/op     74 allocs/op
       pgx:     0.05s       248958 ns/op   42919 B/op    504 allocs/op
  pgx_pool:     0.06s       290609 ns/op   42993 B/op    504 allocs/op
     upper:     0.06s       322309 ns/op    7896 B/op    294 allocs/op
       raw:     0.07s       330560 ns/op   38352 B/op   1038 allocs/op
      sqlc:     0.08s       389753 ns/op   73128 B/op   1251 allocs/op
        pg:     0.08s       412196 ns/op   22258 B/op    629 allocs/op
      gorp:     0.08s       416565 ns/op   57382 B/op   1494 allocs/op
       ent:     0.09s       428113 ns/op   76035 B/op   2039 allocs/op
 sqlboiler:     0.09s       454182 ns/op   58211 B/op   1259 allocs/op
       dbr:     0.09s       459101 ns/op   32416 B/op   1545 allocs/op
       bun:     0.10s       482099 ns/op   33999 B/op   1124 allocs/op
 gorm_prep:     0.10s       506886 ns/op   56064 B/op   1978 allocs/op
     beego:     0.10s       510110 ns/op   55187 B/op   3077 allocs/op
       pop:     0.10s       516264 ns/op   76027 B/op   1511 allocs/op
      sqlx:     0.11s       528324 ns/op   39064 B/op   1516 allocs/op
      gorm:     0.12s       578925 ns/op   57170 B/op   2277 allocs/op
      godb:     0.15s       725953 ns/op   75263 B/op   3084 allocs/op
       rel:     0.15s       730025 ns/op   95321 B/op   2248 allocs/op
      xorm:     0.17s       835066 ns/op  119975 B/op   4687 allocs/op
```

- orm-benchmark -multi=5

```
  
Reports:

  1000 times - Insert
       pgx:     0.24s       239628 ns/op     265 B/op     10 allocs/op
  pgx_pool:     0.25s       247395 ns/op     281 B/op     10 allocs/op
       raw:     0.25s       253857 ns/op     764 B/op     13 allocs/op
     beego:     0.26s       257726 ns/op    2375 B/op     56 allocs/op
    reform:     0.27s       267358 ns/op    1838 B/op     51 allocs/op
      gorp:     0.27s       268447 ns/op    1826 B/op     42 allocs/op
      sqlc:     0.27s       273987 ns/op    2929 B/op     63 allocs/op
       ent:     0.27s       274983 ns/op    4214 B/op    100 allocs/op
 sqlboiler:     0.28s       275186 ns/op    1621 B/op     35 allocs/op
 gorm_prep:     0.29s       287339 ns/op    5342 B/op     67 allocs/op
        pg:     0.30s       295926 ns/op    1890 B/op     10 allocs/op
       dbr:     0.30s       300880 ns/op    2709 B/op     65 allocs/op
       bun:     0.31s       309015 ns/op    5055 B/op     14 allocs/op
      gorm:     0.35s       348676 ns/op    6938 B/op     90 allocs/op
      sqlx:     0.40s       400766 ns/op     856 B/op     19 allocs/op
       rel:     0.43s       430904 ns/op    2504 B/op     42 allocs/op
      xorm:     0.44s       435724 ns/op    3343 B/op     89 allocs/op
      godb:     0.44s       439933 ns/op    4575 B/op    116 allocs/op
     upper:     0.50s       500247 ns/op   13841 B/op    675 allocs/op
       pop:     0.65s       649471 ns/op   10128 B/op    249 allocs/op

  1000 times - MultiInsert 100 row
  pgx_pool:     1.03s      1026416 ns/op  114335 B/op     43 allocs/op
       pgx:     1.05s      1052134 ns/op  114302 B/op     43 allocs/op
       raw:     1.13s      1128428 ns/op  191429 B/op    932 allocs/op
     beego:     1.21s      1207214 ns/op  179428 B/op   2746 allocs/op
    reform:     1.25s      1247874 ns/op  466340 B/op   2747 allocs/op
 gorm_prep:     1.39s      1389979 ns/op  234463 B/op   2183 allocs/op
       ent:     1.50s      1500558 ns/op  412164 B/op   4900 allocs/op
        pg:     1.65s      1651794 ns/op    3314 B/op    112 allocs/op
       bun:     1.72s      1722469 ns/op   42271 B/op    218 allocs/op
      sqlx:     1.84s      1836438 ns/op  171222 B/op   1552 allocs/op
      gorm:     1.90s      1900847 ns/op  272183 B/op   3729 allocs/op
      xorm:     2.12s      2124734 ns/op  254746 B/op   5416 allocs/op
       rel:     2.14s      2140510 ns/op  303991 B/op   3263 allocs/op
      godb:     2.27s      2272415 ns/op  259957 B/op   5894 allocs/op
     upper:     2.88s      2875130 ns/op  545452 B/op  19495 allocs/op
       dbr:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert

  1000 times - Update
       raw:     0.13s       126646 ns/op     753 B/op     13 allocs/op
       ent:     0.14s       143016 ns/op    4561 B/op     99 allocs/op
  pgx_pool:     0.25s       245408 ns/op     287 B/op     10 allocs/op
       pgx:     0.25s       248717 ns/op     272 B/op     10 allocs/op
      sqlc:     0.25s       252867 ns/op     881 B/op     14 allocs/op
 sqlboiler:     0.26s       256039 ns/op     907 B/op     17 allocs/op
      gorp:     0.26s       256606 ns/op    1209 B/op     32 allocs/op
    reform:     0.26s       259635 ns/op    1776 B/op     51 allocs/op
     beego:     0.26s       259768 ns/op    1752 B/op     47 allocs/op
 gorm_prep:     0.27s       274321 ns/op    5113 B/op     59 allocs/op
      sqlx:     0.28s       280752 ns/op     872 B/op     20 allocs/op
       dbr:     0.29s       288316 ns/op    2651 B/op     57 allocs/op
        pg:     0.30s       302860 ns/op     768 B/op      9 allocs/op
       bun:     0.31s       313464 ns/op    4728 B/op      5 allocs/op
       pop:     0.32s       315220 ns/op    6587 B/op    198 allocs/op
      gorm:     0.34s       340785 ns/op    6600 B/op     81 allocs/op
       rel:     0.43s       430602 ns/op    2528 B/op     41 allocs/op
      xorm:     0.44s       438702 ns/op    3648 B/op    126 allocs/op
      godb:     0.47s       466672 ns/op    5160 B/op    154 allocs/op
     upper:     1.24s      1235439 ns/op   33514 B/op   1503 allocs/op

  1000 times - Read
       pgx:     0.12s       122616 ns/op     889 B/op      8 allocs/op
  pgx_pool:     0.12s       122925 ns/op    1076 B/op      9 allocs/op
      sqlc:     0.13s       128436 ns/op    2176 B/op     52 allocs/op
       raw:     0.13s       132747 ns/op    2057 B/op     50 allocs/op
     beego:     0.14s       137875 ns/op    2089 B/op     75 allocs/op
    reform:     0.14s       143096 ns/op    3194 B/op     86 allocs/op
      gorp:     0.14s       143647 ns/op    3881 B/op    194 allocs/op
       pop:     0.15s       148410 ns/op    3561 B/op     71 allocs/op
       ent:     0.15s       154300 ns/op    5401 B/op    149 allocs/op
 gorm_prep:     0.15s       154783 ns/op    4806 B/op     83 allocs/op
        pg:     0.16s       162859 ns/op     872 B/op     20 allocs/op
       dbr:     0.17s       166520 ns/op    2184 B/op     37 allocs/op
 sqlboiler:     0.17s       166843 ns/op    1008 B/op     14 allocs/op
       rel:     0.17s       167827 ns/op    1800 B/op     45 allocs/op
       bun:     0.17s       170264 ns/op    5489 B/op     21 allocs/op
      gorm:     0.20s       200606 ns/op    5156 B/op     94 allocs/op
      sqlx:     0.29s       294624 ns/op    1976 B/op     43 allocs/op
      godb:     0.31s       313029 ns/op    4112 B/op    102 allocs/op
      xorm:     0.32s       317292 ns/op    4616 B/op    125 allocs/op
     upper:     0.33s       332707 ns/op    8136 B/op    315 allocs/op

  1000 times - MultiRead limit 100
    reform:     0.14s       140003 ns/op    2920 B/op     74 allocs/op
  pgx_pool:     0.26s       259669 ns/op   43009 B/op    504 allocs/op
       pgx:     0.27s       269870 ns/op   42949 B/op    504 allocs/op
       raw:     0.32s       316581 ns/op   38342 B/op   1038 allocs/op
     upper:     0.33s       331085 ns/op    7849 B/op    294 allocs/op
      sqlc:     0.37s       373522 ns/op   73159 B/op   1251 allocs/op
 sqlboiler:     0.38s       381156 ns/op   57893 B/op   1259 allocs/op
        pg:     0.41s       406769 ns/op   23382 B/op    629 allocs/op
      sqlx:     0.42s       420247 ns/op   39064 B/op   1516 allocs/op
      gorp:     0.42s       423195 ns/op   57371 B/op   1494 allocs/op
       ent:     0.43s       429583 ns/op   76049 B/op   2039 allocs/op
       dbr:     0.47s       467564 ns/op   32416 B/op   1545 allocs/op
       bun:     0.48s       476494 ns/op   33994 B/op   1124 allocs/op
     beego:     0.48s       481474 ns/op   55183 B/op   3077 allocs/op
 gorm_prep:     0.49s       492455 ns/op   56051 B/op   1977 allocs/op
       pop:     0.51s       508293 ns/op   75759 B/op   1510 allocs/op
      gorm:     0.57s       566960 ns/op   57164 B/op   2277 allocs/op
       rel:     0.72s       720348 ns/op   95321 B/op   2248 allocs/op
      godb:     0.74s       740986 ns/op   75262 B/op   3084 allocs/op
      xorm:     0.78s       782435 ns/op  119952 B/op   4687 allocs/op
```

- orm-benchmark -multi=10

```
  
Reports:

  2000 times - Insert
       pgx:     0.49s       242936 ns/op     269 B/op     10 allocs/op
  pgx_pool:     0.50s       248395 ns/op     285 B/op     10 allocs/op
       raw:     0.50s       252237 ns/op     734 B/op     13 allocs/op
     beego:     0.53s       263911 ns/op    2375 B/op     56 allocs/op
      gorp:     0.53s       264644 ns/op    1798 B/op     42 allocs/op
    reform:     0.53s       265508 ns/op    1807 B/op     51 allocs/op
 sqlboiler:     0.53s       266960 ns/op    1591 B/op     35 allocs/op
      sqlc:     0.54s       269658 ns/op    2900 B/op     63 allocs/op
       ent:     0.54s       271242 ns/op    4215 B/op    100 allocs/op
 gorm_prep:     0.57s       287278 ns/op    5273 B/op     67 allocs/op
       dbr:     0.60s       300203 ns/op    2698 B/op     65 allocs/op
        pg:     0.61s       304581 ns/op     799 B/op     10 allocs/op
       bun:     0.62s       311967 ns/op    5022 B/op     14 allocs/op
      gorm:     0.71s       355545 ns/op    6934 B/op     90 allocs/op
      sqlx:     0.80s       399087 ns/op     856 B/op     19 allocs/op
       rel:     0.85s       427374 ns/op    2504 B/op     42 allocs/op
      xorm:     0.88s       438582 ns/op    3332 B/op     89 allocs/op
      godb:     0.89s       445068 ns/op    4564 B/op    116 allocs/op
     upper:     1.03s       515993 ns/op   13799 B/op    674 allocs/op
       pop:     1.32s       659699 ns/op   10101 B/op    249 allocs/op

  2000 times - MultiInsert 100 row
       pgx:     2.05s      1024915 ns/op  114301 B/op     43 allocs/op
  pgx_pool:     2.15s      1074554 ns/op  114324 B/op     43 allocs/op
       raw:     2.32s      1162353 ns/op  191315 B/op    931 allocs/op
     beego:     2.34s      1169554 ns/op  179417 B/op   2746 allocs/op
    reform:     2.49s      1243034 ns/op  466330 B/op   2747 allocs/op
 gorm_prep:     2.55s      1272524 ns/op  234442 B/op   2183 allocs/op
       ent:     3.03s      1515863 ns/op  412140 B/op   4900 allocs/op
        pg:     3.37s      1687155 ns/op    3314 B/op    112 allocs/op
       bun:     3.44s      1717708 ns/op   42245 B/op    218 allocs/op
      gorm:     3.76s      1882482 ns/op  272244 B/op   3729 allocs/op
      sqlx:     3.81s      1905680 ns/op  171222 B/op   1552 allocs/op
      xorm:     4.34s      2168260 ns/op  254762 B/op   5416 allocs/op
       rel:     4.40s      2199576 ns/op  303993 B/op   3263 allocs/op
      godb:     4.57s      2286684 ns/op  259958 B/op   5895 allocs/op
     upper:     5.80s      2900105 ns/op  545254 B/op  19480 allocs/op
       dbr:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert

  2000 times - Update
       raw:     0.26s       127706 ns/op     748 B/op     13 allocs/op
       ent:     0.29s       142636 ns/op    4556 B/op     99 allocs/op
      sqlc:     0.49s       243345 ns/op     876 B/op     14 allocs/op
       pgx:     0.50s       249898 ns/op     268 B/op     10 allocs/op
 sqlboiler:     0.51s       255577 ns/op     902 B/op     17 allocs/op
  pgx_pool:     0.51s       256300 ns/op     285 B/op     10 allocs/op
      gorp:     0.52s       258145 ns/op    1205 B/op     32 allocs/op
    reform:     0.52s       260869 ns/op    1773 B/op     51 allocs/op
     beego:     0.53s       267188 ns/op    1752 B/op     47 allocs/op
 gorm_prep:     0.54s       272456 ns/op    5121 B/op     59 allocs/op
      sqlx:     0.58s       289247 ns/op     872 B/op     20 allocs/op
       dbr:     0.59s       296876 ns/op    2651 B/op     57 allocs/op
       bun:     0.63s       312976 ns/op    4728 B/op      5 allocs/op
       pop:     0.63s       313167 ns/op    6583 B/op    198 allocs/op
      gorm:     0.71s       355089 ns/op    6604 B/op     81 allocs/op
        pg:     0.71s       357156 ns/op     768 B/op      9 allocs/op
      xorm:     0.88s       440267 ns/op    3648 B/op    126 allocs/op
       rel:     0.89s       443448 ns/op    2528 B/op     41 allocs/op
      godb:     0.94s       470611 ns/op    5160 B/op    154 allocs/op
     upper:     2.47s      1232813 ns/op   33493 B/op   1502 allocs/op

  2000 times - Read
       pgx:     0.24s       120296 ns/op     892 B/op      8 allocs/op
  pgx_pool:     0.24s       121331 ns/op    1080 B/op      9 allocs/op
      sqlc:     0.27s       135507 ns/op    2172 B/op     52 allocs/op
     beego:     0.28s       137746 ns/op    2088 B/op     75 allocs/op
       raw:     0.28s       137975 ns/op    2065 B/op     50 allocs/op
      gorp:     0.29s       145057 ns/op    3876 B/op    194 allocs/op
    reform:     0.29s       145850 ns/op    3197 B/op     86 allocs/op
       pop:     0.30s       150360 ns/op    3529 B/op     71 allocs/op
 gorm_prep:     0.31s       155266 ns/op    4799 B/op     83 allocs/op
       ent:     0.31s       156130 ns/op    5405 B/op    149 allocs/op
       dbr:     0.33s       162565 ns/op    2184 B/op     37 allocs/op
       rel:     0.33s       163920 ns/op    1800 B/op     45 allocs/op
 sqlboiler:     0.33s       166108 ns/op     974 B/op     14 allocs/op
        pg:     0.34s       171702 ns/op     872 B/op     20 allocs/op
       bun:     0.36s       179330 ns/op    5489 B/op     21 allocs/op
      gorm:     0.40s       201063 ns/op    5151 B/op     94 allocs/op
      sqlx:     0.58s       289282 ns/op    1976 B/op     43 allocs/op
      godb:     0.62s       309069 ns/op    4112 B/op    102 allocs/op
      xorm:     0.64s       321271 ns/op    4616 B/op    125 allocs/op
     upper:     0.67s       332976 ns/op    8135 B/op    315 allocs/op

  2000 times - MultiRead limit 100
    reform:     0.27s       136837 ns/op    2924 B/op     74 allocs/op
       pgx:     0.51s       255109 ns/op   42949 B/op    504 allocs/op
  pgx_pool:     0.52s       260725 ns/op   43009 B/op    504 allocs/op
     upper:     0.65s       323224 ns/op    7841 B/op    294 allocs/op
       raw:     0.65s       324075 ns/op   38342 B/op   1038 allocs/op
      sqlc:     0.73s       365401 ns/op   73158 B/op   1251 allocs/op
 sqlboiler:     0.77s       384577 ns/op   57844 B/op   1259 allocs/op
        pg:     0.82s       408630 ns/op   22259 B/op    629 allocs/op
      gorp:     0.83s       414399 ns/op   57372 B/op   1494 allocs/op
       ent:     0.84s       418788 ns/op   76049 B/op   2039 allocs/op
      sqlx:     0.89s       446809 ns/op   39064 B/op   1516 allocs/op
       dbr:     0.92s       461656 ns/op   32416 B/op   1545 allocs/op
       bun:     0.96s       477679 ns/op   33993 B/op   1124 allocs/op
     beego:     0.96s       478402 ns/op   55183 B/op   3077 allocs/op
 gorm_prep:     1.01s       507027 ns/op   56043 B/op   1977 allocs/op
       pop:     1.04s       520534 ns/op   75740 B/op   1510 allocs/op
      gorm:     1.15s       574259 ns/op   57163 B/op   2277 allocs/op
       rel:     1.41s       706976 ns/op   95321 B/op   2248 allocs/op
      godb:     1.47s       733307 ns/op   75262 B/op   3084 allocs/op
      xorm:     1.56s       781403 ns/op  119950 B/op   4687 allocs/op
```

- orm-benchmark -multi=20

```
  
Reports:

  4000 times - Insert
  pgx_pool:     0.97s       242250 ns/op     286 B/op     10 allocs/op
       pgx:     0.98s       244248 ns/op     270 B/op     10 allocs/op
       raw:     1.01s       251312 ns/op     719 B/op     13 allocs/op
    reform:     1.05s       261629 ns/op    1790 B/op     51 allocs/op
      gorp:     1.07s       266507 ns/op    1782 B/op     42 allocs/op
 sqlboiler:     1.07s       267180 ns/op    1574 B/op     35 allocs/op
     beego:     1.07s       268016 ns/op    2376 B/op     56 allocs/op
      sqlc:     1.10s       273880 ns/op    2884 B/op     63 allocs/op
       ent:     1.11s       277125 ns/op    4216 B/op    100 allocs/op
 gorm_prep:     1.13s       281786 ns/op    5255 B/op     67 allocs/op
        pg:     1.18s       295880 ns/op     805 B/op     10 allocs/op
       dbr:     1.20s       299389 ns/op    2693 B/op     65 allocs/op
       bun:     1.25s       312872 ns/op    5007 B/op     14 allocs/op
      gorm:     1.36s       340830 ns/op    6932 B/op     90 allocs/op
      sqlx:     1.58s       395829 ns/op     856 B/op     19 allocs/op
       rel:     1.72s       429307 ns/op    2504 B/op     42 allocs/op
      xorm:     1.74s       435515 ns/op    3328 B/op     89 allocs/op
      godb:     1.79s       448210 ns/op    4558 B/op    116 allocs/op
     upper:     2.02s       506170 ns/op   13774 B/op    673 allocs/op
       pop:     2.58s       645772 ns/op   10080 B/op    249 allocs/op

  4000 times - MultiInsert 100 row
  pgx_pool:     3.97s       992143 ns/op  114317 B/op     43 allocs/op
       pgx:     4.23s      1057670 ns/op  114295 B/op     43 allocs/op
       raw:     4.27s      1067713 ns/op  191258 B/op    931 allocs/op
     beego:     4.56s      1139690 ns/op  179410 B/op   2746 allocs/op
    reform:     4.81s      1203049 ns/op  466275 B/op   2747 allocs/op
 gorm_prep:     4.85s      1212158 ns/op  234441 B/op   2183 allocs/op
       ent:     5.74s      1434859 ns/op  412125 B/op   4900 allocs/op
        pg:     6.33s      1583631 ns/op    3576 B/op    112 allocs/op
       bun:     6.60s      1650887 ns/op   42241 B/op    218 allocs/op
      sqlx:     7.08s      1769947 ns/op  171188 B/op   1552 allocs/op
      gorm:     7.48s      1870391 ns/op  272272 B/op   3729 allocs/op
       rel:     8.40s      2100572 ns/op  303993 B/op   3263 allocs/op
      xorm:     8.53s      2132449 ns/op  254776 B/op   5416 allocs/op
      godb:     8.93s      2231374 ns/op  259954 B/op   5895 allocs/op
     upper:    11.36s      2840453 ns/op  545161 B/op  19473 allocs/op
       dbr:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert

  4000 times - Update
       raw:     0.51s       127709 ns/op     750 B/op     13 allocs/op
       ent:     0.57s       143357 ns/op    4558 B/op     99 allocs/op
       pgx:     0.96s       241099 ns/op     270 B/op     10 allocs/op
  pgx_pool:     0.99s       246781 ns/op     286 B/op     10 allocs/op
      gorp:     1.01s       251854 ns/op    1204 B/op     32 allocs/op
 sqlboiler:     1.03s       257844 ns/op     901 B/op     17 allocs/op
      sqlc:     1.04s       258809 ns/op     878 B/op     14 allocs/op
    reform:     1.05s       262500 ns/op    1774 B/op     51 allocs/op
     beego:     1.07s       266486 ns/op    1752 B/op     47 allocs/op
 gorm_prep:     1.10s       274047 ns/op    5116 B/op     59 allocs/op
      sqlx:     1.16s       290403 ns/op     872 B/op     20 allocs/op
       dbr:     1.19s       296952 ns/op    2651 B/op     57 allocs/op
        pg:     1.25s       312667 ns/op     768 B/op      9 allocs/op
       pop:     1.26s       314424 ns/op    6582 B/op    198 allocs/op
       bun:     1.27s       318011 ns/op    4728 B/op      5 allocs/op
      gorm:     1.37s       341993 ns/op    6604 B/op     81 allocs/op
      xorm:     1.78s       445827 ns/op    3648 B/op    126 allocs/op
       rel:     1.78s       446022 ns/op    2528 B/op     41 allocs/op
      godb:     1.89s       471861 ns/op    5160 B/op    154 allocs/op
     upper:     4.91s      1227465 ns/op   33483 B/op   1502 allocs/op

  4000 times - Read
  pgx_pool:     0.49s       122268 ns/op    1079 B/op      9 allocs/op
       pgx:     0.49s       123500 ns/op     894 B/op      8 allocs/op
      sqlc:     0.54s       134421 ns/op    2172 B/op     52 allocs/op
       raw:     0.54s       135612 ns/op    2062 B/op     50 allocs/op
     beego:     0.56s       139114 ns/op    2088 B/op     75 allocs/op
    reform:     0.58s       144011 ns/op    3198 B/op     86 allocs/op
      gorp:     0.59s       147064 ns/op    3878 B/op    194 allocs/op
       pop:     0.61s       151662 ns/op    3507 B/op     71 allocs/op
       ent:     0.62s       154280 ns/op    5404 B/op    149 allocs/op
 gorm_prep:     0.62s       154665 ns/op    4801 B/op     83 allocs/op
        pg:     0.64s       158924 ns/op     872 B/op     20 allocs/op
 sqlboiler:     0.65s       161673 ns/op     965 B/op     14 allocs/op
       bun:     0.67s       167912 ns/op    5489 B/op     21 allocs/op
       dbr:     0.68s       169033 ns/op    2184 B/op     37 allocs/op
       rel:     0.68s       169319 ns/op    1800 B/op     45 allocs/op
      gorm:     0.80s       200711 ns/op    5149 B/op     94 allocs/op
      sqlx:     1.19s       298076 ns/op    1976 B/op     43 allocs/op
      godb:     1.25s       312583 ns/op    4112 B/op    102 allocs/op
      xorm:     1.28s       319499 ns/op    4616 B/op    125 allocs/op
     upper:     1.35s       336626 ns/op    8136 B/op    315 allocs/op

  4000 times - MultiRead limit 100
    reform:     0.54s       136140 ns/op    2926 B/op     74 allocs/op
       pgx:     1.04s       260451 ns/op   42949 B/op    504 allocs/op
  pgx_pool:     1.06s       263947 ns/op   43006 B/op    504 allocs/op
       raw:     1.27s       318055 ns/op   38341 B/op   1038 allocs/op
     upper:     1.32s       329366 ns/op    7841 B/op    294 allocs/op
      sqlc:     1.44s       359212 ns/op   73158 B/op   1251 allocs/op
        pg:     1.57s       391951 ns/op   22783 B/op    629 allocs/op
 sqlboiler:     1.60s       398789 ns/op   57837 B/op   1259 allocs/op
      gorp:     1.65s       412130 ns/op   57370 B/op   1494 allocs/op
       ent:     1.68s       420623 ns/op   76044 B/op   2039 allocs/op
      sqlx:     1.69s       423136 ns/op   39064 B/op   1516 allocs/op
       dbr:     1.86s       465616 ns/op   32416 B/op   1545 allocs/op
       bun:     1.91s       477167 ns/op   33991 B/op   1124 allocs/op
     beego:     1.92s       481136 ns/op   55181 B/op   3077 allocs/op
 gorm_prep:     2.02s       504888 ns/op   56043 B/op   1977 allocs/op
       pop:     2.04s       510185 ns/op   75729 B/op   1510 allocs/op
      gorm:     2.30s       575306 ns/op   57160 B/op   2277 allocs/op
       rel:     2.77s       691399 ns/op   95321 B/op   2248 allocs/op
      godb:     2.98s       745226 ns/op   75262 B/op   3084 allocs/op
      xorm:     3.20s       799943 ns/op  119953 B/op   4687 allocs/op
```
