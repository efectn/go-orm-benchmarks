# Results

- orm-benchmark (with no flags)

```
   
Reports:

   200 times - Insert
     beego:     0.06s       311096 ns/op    2372 B/op     56 allocs/op
       pgx:     0.07s       327888 ns/op     244 B/op     10 allocs/op
       raw:     0.07s       339534 ns/op    1011 B/op     15 allocs/op
    reform:     0.07s       340601 ns/op    2085 B/op     53 allocs/op
  pgx_pool:     0.07s       355450 ns/op     265 B/op     10 allocs/op
      sqlc:     0.07s       364352 ns/op    3181 B/op     64 allocs/op
      gorp:     0.08s       376254 ns/op    2067 B/op     43 allocs/op
        pg:     0.08s       396638 ns/op    1031 B/op     11 allocs/op
 sqlboiler:     0.08s       397028 ns/op    1876 B/op     37 allocs/op
      gorm:     0.08s       416779 ns/op    7138 B/op     90 allocs/op
       dbr:     0.09s       425835 ns/op    2802 B/op     66 allocs/op
       ent:     0.09s       435397 ns/op    4172 B/op     99 allocs/op
 gorm_prep:     0.09s       439498 ns/op    5613 B/op     70 allocs/op
      sqlx:     0.10s       494228 ns/op     856 B/op     19 allocs/op
      xorm:     0.10s       506739 ns/op    3440 B/op     90 allocs/op
       rel:     0.11s       535921 ns/op    2511 B/op     41 allocs/op
       bun:     0.11s       553948 ns/op    5297 B/op     14 allocs/op
      godb:     0.11s       573444 ns/op    4673 B/op    116 allocs/op
     upper:     0.13s       625181 ns/op   14206 B/op    685 allocs/op
       pop:     0.17s       857223 ns/op   10402 B/op    250 allocs/op

   200 times - MultiInsert 100 row
  pgx_pool:     0.24s      1201917 ns/op  114372 B/op     43 allocs/op
     beego:     0.28s      1392563 ns/op  179487 B/op   2746 allocs/op
       pgx:     0.28s      1394396 ns/op  114350 B/op     43 allocs/op
       raw:     0.30s      1487875 ns/op  192353 B/op    938 allocs/op
 gorm_prep:     0.31s      1533923 ns/op  235346 B/op   2281 allocs/op
    reform:     0.31s      1550954 ns/op  466460 B/op   2748 allocs/op
       ent:     0.36s      1812657 ns/op  412424 B/op   4902 allocs/op
       bun:     0.38s      1893052 ns/op   42359 B/op    217 allocs/op
        pg:     0.41s      2038036 ns/op    3317 B/op    112 allocs/op
      sqlx:     0.43s      2133174 ns/op  171727 B/op   1552 allocs/op
      gorm:     0.45s      2245907 ns/op  272525 B/op   3827 allocs/op
      xorm:     0.51s      2525090 ns/op  254780 B/op   5416 allocs/op
      godb:     0.51s      2533161 ns/op  259953 B/op   5893 allocs/op
       rel:     0.52s      2580476 ns/op  303975 B/op   3261 allocs/op
     upper:     0.65s      3250919 ns/op  547001 B/op  19610 allocs/op
 sqlboiler:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert

   200 times - Update
       raw:     0.03s       144061 ns/op     721 B/op     13 allocs/op
       ent:     0.03s       146188 ns/op    4529 B/op     99 allocs/op
    reform:     0.06s       278026 ns/op    1744 B/op     51 allocs/op
      sqlc:     0.06s       289449 ns/op     890 B/op     14 allocs/op
      sqlx:     0.06s       313793 ns/op     872 B/op     20 allocs/op
     beego:     0.06s       314684 ns/op    1755 B/op     47 allocs/op
  pgx_pool:     0.06s       323609 ns/op     302 B/op     10 allocs/op
 sqlboiler:     0.07s       327327 ns/op     924 B/op     17 allocs/op
       pgx:     0.07s       338512 ns/op     282 B/op     10 allocs/op
 gorm_prep:     0.07s       340667 ns/op    5132 B/op     59 allocs/op
      gorp:     0.07s       345412 ns/op    1223 B/op     32 allocs/op
       dbr:     0.08s       398622 ns/op    2651 B/op     57 allocs/op
       pop:     0.08s       414942 ns/op    6554 B/op    198 allocs/op
        pg:     0.08s       422774 ns/op     768 B/op      9 allocs/op
      gorm:     0.09s       448213 ns/op    6625 B/op     81 allocs/op
      xorm:     0.10s       505964 ns/op    3649 B/op    126 allocs/op
       rel:     0.11s       547702 ns/op    2531 B/op     41 allocs/op
      godb:     0.11s       558665 ns/op    5160 B/op    154 allocs/op
     upper:     0.29s      1437879 ns/op   33727 B/op   1510 allocs/op
       bun:     0.50s      2512784 ns/op    4728 B/op      5 allocs/op

   200 times - Read
       pgx:     0.03s       126661 ns/op     886 B/op      8 allocs/op
  pgx_pool:     0.03s       134841 ns/op    1078 B/op      9 allocs/op
      sqlc:     0.03s       135212 ns/op    2162 B/op     52 allocs/op
       raw:     0.03s       137244 ns/op    2098 B/op     50 allocs/op
     beego:     0.03s       143180 ns/op    2093 B/op     75 allocs/op
      gorp:     0.03s       143821 ns/op    3869 B/op    194 allocs/op
       pop:     0.03s       154908 ns/op    3914 B/op     71 allocs/op
 gorm_prep:     0.03s       164457 ns/op    4802 B/op     84 allocs/op
    reform:     0.03s       166615 ns/op    3234 B/op     86 allocs/op
       bun:     0.03s       171216 ns/op    5490 B/op     21 allocs/op
       dbr:     0.04s       175851 ns/op    2184 B/op     37 allocs/op
 sqlboiler:     0.04s       184793 ns/op    1303 B/op     14 allocs/op
        pg:     0.04s       186459 ns/op     873 B/op     20 allocs/op
       ent:     0.04s       196487 ns/op    5424 B/op    149 allocs/op
       rel:     0.04s       197333 ns/op    1803 B/op     45 allocs/op
      gorm:     0.05s       240639 ns/op    5155 B/op     95 allocs/op
      sqlx:     0.06s       305352 ns/op    1976 B/op     43 allocs/op
      godb:     0.06s       322727 ns/op    4112 B/op    102 allocs/op
      xorm:     0.07s       337886 ns/op    4616 B/op    125 allocs/op
     upper:     0.07s       352362 ns/op    8125 B/op    315 allocs/op

   200 times - MultiRead limit 100
    reform:     0.03s       140426 ns/op    2955 B/op     74 allocs/op
       pgx:     0.05s       272854 ns/op   42919 B/op    504 allocs/op
  pgx_pool:     0.06s       279023 ns/op   42993 B/op    504 allocs/op
       raw:     0.07s       332792 ns/op   38352 B/op   1038 allocs/op
     upper:     0.07s       363993 ns/op    7896 B/op    294 allocs/op
      sqlc:     0.07s       369456 ns/op   73128 B/op   1251 allocs/op
        pg:     0.08s       398132 ns/op   22263 B/op    629 allocs/op
 sqlboiler:     0.08s       416160 ns/op   58032 B/op   1259 allocs/op
       ent:     0.09s       432455 ns/op   76035 B/op   2039 allocs/op
      gorp:     0.09s       436652 ns/op   57381 B/op   1494 allocs/op
      sqlx:     0.10s       475673 ns/op   39064 B/op   1516 allocs/op
       dbr:     0.10s       477031 ns/op   32416 B/op   1545 allocs/op
       bun:     0.10s       494319 ns/op   33992 B/op   1124 allocs/op
     beego:     0.10s       495389 ns/op   55188 B/op   3077 allocs/op
       pop:     0.11s       555229 ns/op   75953 B/op   1511 allocs/op
 gorm_prep:     0.12s       599259 ns/op   56085 B/op   2079 allocs/op
      gorm:     0.12s       608657 ns/op   57192 B/op   2378 allocs/op
       rel:     0.14s       697483 ns/op   95322 B/op   2248 allocs/op
      godb:     0.15s       752253 ns/op   75263 B/op   3084 allocs/op
      xorm:     0.16s       791186 ns/op  119972 B/op   4687 allocs/op
```

- orm-benchmark -multi=5

```
  
Reports:

  1000 times - Insert
       pgx:     0.28s       278722 ns/op     265 B/op     10 allocs/op
     beego:     0.31s       305531 ns/op    2375 B/op     56 allocs/op
 sqlboiler:     0.31s       305915 ns/op    1623 B/op     35 allocs/op
      gorp:     0.31s       307818 ns/op    1827 B/op     42 allocs/op
      sqlc:     0.32s       319520 ns/op    2929 B/op     63 allocs/op
  pgx_pool:     0.32s       320954 ns/op     281 B/op     10 allocs/op
       raw:     0.32s       324959 ns/op     764 B/op     13 allocs/op
       bun:     0.34s       341566 ns/op    5016 B/op     14 allocs/op
 gorm_prep:     0.34s       341895 ns/op    5312 B/op     68 allocs/op
        pg:     0.35s       346480 ns/op     841 B/op     10 allocs/op
    reform:     0.35s       351447 ns/op    1838 B/op     51 allocs/op
       ent:     0.36s       359737 ns/op    4214 B/op    100 allocs/op
       dbr:     0.38s       378237 ns/op    2709 B/op     65 allocs/op
      gorm:     0.40s       399890 ns/op    6946 B/op     90 allocs/op
      sqlx:     0.44s       443278 ns/op     856 B/op     19 allocs/op
      xorm:     0.49s       487539 ns/op    3343 B/op     89 allocs/op
       rel:     0.53s       527846 ns/op    2504 B/op     42 allocs/op
      godb:     0.59s       585947 ns/op    4575 B/op    116 allocs/op
       pop:     0.72s       720595 ns/op   10144 B/op    249 allocs/op
     upper:     0.73s       729488 ns/op   13840 B/op    675 allocs/op

  1000 times - MultiInsert 100 row
  pgx_pool:     1.22s      1219927 ns/op  114328 B/op     43 allocs/op
       pgx:     1.25s      1253251 ns/op  114308 B/op     43 allocs/op
       raw:     1.34s      1344615 ns/op  191429 B/op    932 allocs/op
     beego:     1.45s      1450914 ns/op  179439 B/op   2746 allocs/op
    reform:     1.51s      1513759 ns/op  466327 B/op   2747 allocs/op
 gorm_prep:     1.61s      1608979 ns/op  235262 B/op   2283 allocs/op
       ent:     1.75s      1746338 ns/op  412165 B/op   4900 allocs/op
        pg:     1.96s      1963408 ns/op    3315 B/op    112 allocs/op
       bun:     2.04s      2042344 ns/op   42272 B/op    218 allocs/op
      sqlx:     2.12s      2118811 ns/op  171259 B/op   1552 allocs/op
      gorm:     2.21s      2211684 ns/op  272908 B/op   3829 allocs/op
       rel:     2.41s      2412611 ns/op  303990 B/op   3263 allocs/op
      xorm:     2.49s      2485593 ns/op  254779 B/op   5416 allocs/op
      godb:     2.50s      2499512 ns/op  259968 B/op   5894 allocs/op
     upper:     3.34s      3335917 ns/op  545452 B/op  19495 allocs/op
       pop:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert

  1000 times - Update
       raw:     0.13s       134945 ns/op     753 B/op     13 allocs/op
       ent:     0.15s       150868 ns/op    4561 B/op     99 allocs/op
       pgx:     0.29s       287152 ns/op     273 B/op     10 allocs/op
 sqlboiler:     0.30s       302721 ns/op     907 B/op     17 allocs/op
      sqlx:     0.30s       304861 ns/op     872 B/op     20 allocs/op
      gorp:     0.31s       305982 ns/op    1209 B/op     32 allocs/op
     beego:     0.31s       309510 ns/op    1752 B/op     47 allocs/op
      sqlc:     0.33s       325367 ns/op     881 B/op     14 allocs/op
  pgx_pool:     0.33s       328407 ns/op     287 B/op     10 allocs/op
       bun:     0.35s       348970 ns/op    4728 B/op      5 allocs/op
       pop:     0.35s       352097 ns/op    6586 B/op    198 allocs/op
    reform:     0.36s       357722 ns/op    1777 B/op     51 allocs/op
       dbr:     0.38s       383689 ns/op    2651 B/op     57 allocs/op
 gorm_prep:     0.39s       385846 ns/op    5113 B/op     59 allocs/op
        pg:     0.41s       413237 ns/op     768 B/op      9 allocs/op
      gorm:     0.42s       416056 ns/op    6600 B/op     81 allocs/op
       rel:     0.54s       542275 ns/op    2528 B/op     41 allocs/op
      godb:     0.54s       542557 ns/op    5160 B/op    154 allocs/op
      xorm:     0.59s       590132 ns/op    3648 B/op    126 allocs/op
     upper:     1.55s      1553969 ns/op   33515 B/op   1503 allocs/op

  1000 times - Read
  pgx_pool:     0.12s       117392 ns/op    1076 B/op      9 allocs/op
       pgx:     0.14s       137836 ns/op     889 B/op      8 allocs/op
      sqlc:     0.14s       143798 ns/op    2176 B/op     52 allocs/op
       raw:     0.14s       143912 ns/op    2057 B/op     50 allocs/op
     beego:     0.15s       145747 ns/op    2089 B/op     75 allocs/op
      gorp:     0.15s       154748 ns/op    3881 B/op    194 allocs/op
       dbr:     0.16s       158865 ns/op    2184 B/op     37 allocs/op
        pg:     0.16s       163674 ns/op     872 B/op     20 allocs/op
 gorm_prep:     0.16s       164206 ns/op    4807 B/op     84 allocs/op
       pop:     0.17s       172672 ns/op    3561 B/op     71 allocs/op
       ent:     0.17s       174061 ns/op    5401 B/op    149 allocs/op
       bun:     0.18s       178513 ns/op    5489 B/op     21 allocs/op
       rel:     0.18s       178691 ns/op    1800 B/op     45 allocs/op
 sqlboiler:     0.19s       185915 ns/op     971 B/op     14 allocs/op
    reform:     0.19s       194086 ns/op    3193 B/op     86 allocs/op
      gorm:     0.21s       211700 ns/op    5158 B/op     95 allocs/op
      sqlx:     0.30s       296749 ns/op    1976 B/op     43 allocs/op
      godb:     0.33s       326356 ns/op    4112 B/op    102 allocs/op
     upper:     0.37s       366732 ns/op    8136 B/op    315 allocs/op
      xorm:     0.40s       397704 ns/op    4616 B/op    125 allocs/op

  1000 times - MultiRead limit 100
    reform:     0.15s       151887 ns/op    2920 B/op     74 allocs/op
  pgx_pool:     0.26s       262673 ns/op   43009 B/op    504 allocs/op
       pgx:     0.28s       279329 ns/op   42950 B/op    504 allocs/op
       raw:     0.33s       333372 ns/op   38342 B/op   1038 allocs/op
      sqlc:     0.36s       357005 ns/op   73158 B/op   1251 allocs/op
     upper:     0.37s       367011 ns/op    7849 B/op    294 allocs/op
      sqlx:     0.43s       429707 ns/op   39064 B/op   1516 allocs/op
 sqlboiler:     0.43s       430889 ns/op   57930 B/op   1259 allocs/op
       ent:     0.43s       430937 ns/op   76049 B/op   2039 allocs/op
      gorp:     0.44s       441764 ns/op   57373 B/op   1494 allocs/op
        pg:     0.44s       442681 ns/op   22260 B/op    629 allocs/op
       dbr:     0.47s       470061 ns/op   32416 B/op   1545 allocs/op
       bun:     0.48s       484180 ns/op   33996 B/op   1124 allocs/op
     beego:     0.50s       500613 ns/op   55184 B/op   3077 allocs/op
       pop:     0.52s       522578 ns/op   75767 B/op   1511 allocs/op
 gorm_prep:     0.58s       582964 ns/op   56067 B/op   2078 allocs/op
      gorm:     0.62s       624198 ns/op   57190 B/op   2378 allocs/op
       rel:     0.71s       710940 ns/op   95321 B/op   2248 allocs/op
      godb:     0.79s       788159 ns/op   75264 B/op   3084 allocs/op
      xorm:     0.93s       928171 ns/op  119953 B/op   4687 allocs/op
```

- orm-benchmark -multi=10

```
  
Reports:

  2000 times - Insert
       pgx:     0.54s       272407 ns/op     269 B/op     10 allocs/op
  pgx_pool:     0.58s       288427 ns/op     285 B/op     10 allocs/op
      gorp:     0.59s       296661 ns/op    1798 B/op     42 allocs/op
 sqlboiler:     0.60s       300715 ns/op    1592 B/op     35 allocs/op
     beego:     0.61s       305918 ns/op    2376 B/op     56 allocs/op
      sqlc:     0.62s       307816 ns/op    2900 B/op     63 allocs/op
       raw:     0.62s       312061 ns/op     734 B/op     13 allocs/op
       ent:     0.63s       313112 ns/op    4215 B/op    100 allocs/op
    reform:     0.63s       315506 ns/op    1807 B/op     51 allocs/op
       dbr:     0.68s       339135 ns/op    2698 B/op     65 allocs/op
 gorm_prep:     0.68s       339969 ns/op    5281 B/op     68 allocs/op
       bun:     0.70s       349586 ns/op    5023 B/op     14 allocs/op
      gorm:     0.77s       386609 ns/op    6961 B/op     90 allocs/op
        pg:     0.81s       403714 ns/op     817 B/op     10 allocs/op
      sqlx:     0.90s       449603 ns/op     856 B/op     19 allocs/op
       rel:     0.98s       488085 ns/op    2504 B/op     42 allocs/op
      godb:     1.00s       500502 ns/op    4564 B/op    116 allocs/op
      xorm:     1.03s       516858 ns/op    3336 B/op     89 allocs/op
     upper:     1.15s       577009 ns/op   13798 B/op    674 allocs/op
       pop:     1.52s       759132 ns/op   10088 B/op    249 allocs/op

  2000 times - MultiInsert 100 row
       pgx:     2.41s      1205430 ns/op  114319 B/op     43 allocs/op
  pgx_pool:     2.48s      1237865 ns/op  114320 B/op     43 allocs/op
       raw:     2.60s      1298020 ns/op  191315 B/op    931 allocs/op
     beego:     2.99s      1493224 ns/op  179433 B/op   2746 allocs/op
    reform:     3.18s      1589592 ns/op  466303 B/op   2747 allocs/op
 gorm_prep:     3.29s      1643758 ns/op  235246 B/op   2283 allocs/op
       ent:     3.58s      1789932 ns/op  412140 B/op   4900 allocs/op
        pg:     3.93s      1967229 ns/op    3314 B/op    112 allocs/op
       bun:     3.99s      1995388 ns/op   42249 B/op    218 allocs/op
      sqlx:     4.30s      2151379 ns/op  171203 B/op   1552 allocs/op
      gorm:     4.34s      2171402 ns/op  272998 B/op   3829 allocs/op
      xorm:     4.74s      2370737 ns/op  254771 B/op   5416 allocs/op
       rel:     4.80s      2401570 ns/op  303993 B/op   3263 allocs/op
      godb:     5.08s      2539940 ns/op  259976 B/op   5895 allocs/op
     upper:     6.54s      3271514 ns/op  545254 B/op  19480 allocs/op
      gorp:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert

  2000 times - Update
       raw:     0.27s       134569 ns/op     748 B/op     13 allocs/op
       ent:     0.28s       141721 ns/op    4556 B/op     99 allocs/op
       pgx:     0.55s       273099 ns/op     268 B/op     10 allocs/op
      sqlc:     0.57s       286365 ns/op     876 B/op     14 allocs/op
 sqlboiler:     0.61s       305263 ns/op     901 B/op     17 allocs/op
     beego:     0.62s       307921 ns/op    1752 B/op     47 allocs/op
      sqlx:     0.62s       310697 ns/op     872 B/op     20 allocs/op
      gorp:     0.63s       314743 ns/op    1205 B/op     32 allocs/op
 gorm_prep:     0.63s       317403 ns/op    5121 B/op     59 allocs/op
    reform:     0.65s       323813 ns/op    1772 B/op     51 allocs/op
  pgx_pool:     0.65s       327247 ns/op     285 B/op     10 allocs/op
       dbr:     0.68s       341315 ns/op    2651 B/op     57 allocs/op
       bun:     0.71s       357050 ns/op    4728 B/op      5 allocs/op
       pop:     0.75s       375233 ns/op    6582 B/op    198 allocs/op
        pg:     0.77s       384981 ns/op     768 B/op      9 allocs/op
      gorm:     0.79s       392703 ns/op    6604 B/op     81 allocs/op
      xorm:     0.99s       494454 ns/op    3648 B/op    126 allocs/op
       rel:     1.00s       501788 ns/op    2528 B/op     41 allocs/op
      godb:     1.08s       541064 ns/op    5160 B/op    154 allocs/op
     upper:     2.91s      1457084 ns/op   33495 B/op   1502 allocs/op

  2000 times - Read
       pgx:     0.26s       128171 ns/op     892 B/op      8 allocs/op
  pgx_pool:     0.27s       133068 ns/op    1080 B/op      9 allocs/op
       raw:     0.28s       139785 ns/op    2065 B/op     50 allocs/op
      sqlc:     0.28s       141067 ns/op    2172 B/op     52 allocs/op
     beego:     0.29s       147175 ns/op    2088 B/op     75 allocs/op
        pg:     0.30s       152404 ns/op     872 B/op     20 allocs/op
       pop:     0.33s       166605 ns/op    3529 B/op     71 allocs/op
    reform:     0.34s       168311 ns/op    3197 B/op     86 allocs/op
 gorm_prep:     0.34s       170623 ns/op    4801 B/op     84 allocs/op
       ent:     0.35s       172857 ns/op    5405 B/op    149 allocs/op
 sqlboiler:     0.35s       173513 ns/op     975 B/op     14 allocs/op
       rel:     0.35s       174789 ns/op    1800 B/op     45 allocs/op
      gorp:     0.35s       175224 ns/op    3876 B/op    194 allocs/op
       bun:     0.35s       175338 ns/op    5489 B/op     21 allocs/op
       dbr:     0.36s       181314 ns/op    2184 B/op     37 allocs/op
      gorm:     0.44s       217657 ns/op    5153 B/op     95 allocs/op
      godb:     0.64s       321959 ns/op    4112 B/op    102 allocs/op
      xorm:     0.68s       341556 ns/op    4616 B/op    125 allocs/op
      sqlx:     0.68s       341908 ns/op    1976 B/op     43 allocs/op
     upper:     0.71s       354071 ns/op    8135 B/op    315 allocs/op

  2000 times - MultiRead limit 100
    reform:     0.34s       170064 ns/op    2924 B/op     74 allocs/op
       pgx:     0.51s       254745 ns/op   42949 B/op    504 allocs/op
  pgx_pool:     0.57s       282877 ns/op   43009 B/op    504 allocs/op
       raw:     0.65s       326658 ns/op   38342 B/op   1038 allocs/op
     upper:     0.67s       337234 ns/op    7841 B/op    294 allocs/op
      sqlc:     0.77s       382684 ns/op   73158 B/op   1251 allocs/op
        pg:     0.78s       390487 ns/op   23308 B/op    629 allocs/op
 sqlboiler:     0.78s       392037 ns/op   57825 B/op   1259 allocs/op
       ent:     0.86s       427809 ns/op   76049 B/op   2039 allocs/op
      sqlx:     0.87s       436788 ns/op   39064 B/op   1516 allocs/op
      gorp:     0.90s       448162 ns/op   57372 B/op   1494 allocs/op
       dbr:     0.95s       474147 ns/op   32416 B/op   1545 allocs/op
       bun:     0.99s       495076 ns/op   33995 B/op   1124 allocs/op
     beego:     1.01s       503308 ns/op   55183 B/op   3077 allocs/op
       pop:     1.13s       563355 ns/op   75737 B/op   1510 allocs/op
 gorm_prep:     1.19s       597403 ns/op   56063 B/op   2078 allocs/op
      gorm:     1.31s       655107 ns/op   57184 B/op   2378 allocs/op
      godb:     1.51s       755685 ns/op   75266 B/op   3084 allocs/op
       rel:     1.52s       760730 ns/op   95321 B/op   2248 allocs/op
      xorm:     1.67s       832962 ns/op  119953 B/op   4687 allocs/op
```

- orm-benchmark -multi=20

```
  
Reports:

  4000 times - Insert
       raw:     1.12s       280677 ns/op     719 B/op     13 allocs/op
  pgx_pool:     1.13s       282185 ns/op     286 B/op     10 allocs/op
       pgx:     1.13s       282228 ns/op     270 B/op     10 allocs/op
 sqlboiler:     1.20s       298863 ns/op    1574 B/op     35 allocs/op
     beego:     1.20s       300376 ns/op    2375 B/op     56 allocs/op
    reform:     1.21s       303309 ns/op    1790 B/op     51 allocs/op
      sqlc:     1.24s       309791 ns/op    2884 B/op     63 allocs/op
      gorp:     1.25s       312610 ns/op    1782 B/op     42 allocs/op
       ent:     1.26s       315935 ns/op    4216 B/op    100 allocs/op
        pg:     1.32s       330917 ns/op     805 B/op     10 allocs/op
       dbr:     1.38s       345130 ns/op    2693 B/op     65 allocs/op
       bun:     1.41s       353739 ns/op    5008 B/op     14 allocs/op
 gorm_prep:     1.47s       366267 ns/op    5263 B/op     68 allocs/op
      gorm:     1.55s       387954 ns/op    6940 B/op     90 allocs/op
       rel:     1.92s       478842 ns/op    2504 B/op     42 allocs/op
      xorm:     1.99s       498569 ns/op    3326 B/op     89 allocs/op
      sqlx:     2.00s       499343 ns/op     856 B/op     19 allocs/op
      godb:     2.10s       524870 ns/op    4558 B/op    116 allocs/op
     upper:     2.43s       607655 ns/op   13775 B/op    673 allocs/op
       pop:     3.02s       754756 ns/op   10072 B/op    249 allocs/op

  4000 times - MultiInsert 100 row
  pgx_pool:     4.69s      1172891 ns/op  114319 B/op     43 allocs/op
       pgx:     4.73s      1182669 ns/op  114309 B/op     43 allocs/op
       raw:     5.11s      1276312 ns/op  191258 B/op    931 allocs/op
     beego:     5.51s      1376602 ns/op  179425 B/op   2746 allocs/op
    reform:     5.90s      1474171 ns/op  466310 B/op   2747 allocs/op
 gorm_prep:     6.25s      1561899 ns/op  235252 B/op   2283 allocs/op
       ent:     6.88s      1721247 ns/op  412125 B/op   4900 allocs/op
        pg:     7.45s      1861363 ns/op    3314 B/op    112 allocs/op
       bun:     7.65s      1912901 ns/op   42244 B/op    219 allocs/op
      sqlx:     8.08s      2021197 ns/op  171199 B/op   1552 allocs/op
      gorm:     8.57s      2143287 ns/op  273050 B/op   3829 allocs/op
       rel:     9.37s      2342020 ns/op  303993 B/op   3263 allocs/op
      xorm:     9.62s      2405331 ns/op  254782 B/op   5416 allocs/op
      godb:     9.97s      2493513 ns/op  259953 B/op   5895 allocs/op
     upper:    12.89s      3223664 ns/op  545162 B/op  19473 allocs/op
 sqlboiler:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert

  4000 times - Update
       raw:     0.54s       135094 ns/op     750 B/op     13 allocs/op
       ent:     0.60s       149399 ns/op    4558 B/op     99 allocs/op
 sqlboiler:     1.19s       296906 ns/op     901 B/op     17 allocs/op
      gorp:     1.19s       298263 ns/op    1204 B/op     32 allocs/op
      sqlx:     1.20s       299071 ns/op     872 B/op     20 allocs/op
       pgx:     1.21s       302504 ns/op     270 B/op     10 allocs/op
  pgx_pool:     1.21s       303719 ns/op     286 B/op     10 allocs/op
    reform:     1.22s       305200 ns/op    1775 B/op     51 allocs/op
     beego:     1.33s       333688 ns/op    1752 B/op     47 allocs/op
       dbr:     1.36s       341014 ns/op    2651 B/op     57 allocs/op
      sqlc:     1.38s       344597 ns/op     878 B/op     14 allocs/op
       pop:     1.40s       350555 ns/op    6582 B/op    198 allocs/op
 gorm_prep:     1.41s       353281 ns/op    5116 B/op     59 allocs/op
       bun:     1.49s       373295 ns/op    4728 B/op      5 allocs/op
      gorm:     1.54s       385691 ns/op    6604 B/op     81 allocs/op
        pg:     1.59s       397484 ns/op     768 B/op      9 allocs/op
      xorm:     2.00s       499739 ns/op    3648 B/op    126 allocs/op
       rel:     2.14s       535980 ns/op    2528 B/op     41 allocs/op
      godb:     2.32s       579566 ns/op    5160 B/op    154 allocs/op
     upper:     5.85s      1461406 ns/op   33493 B/op   1502 allocs/op

  4000 times - Read
       pgx:     0.52s       129544 ns/op     894 B/op      8 allocs/op
  pgx_pool:     0.53s       133516 ns/op    1079 B/op      9 allocs/op
       raw:     0.58s       144277 ns/op    2062 B/op     50 allocs/op
     beego:     0.59s       147388 ns/op    2088 B/op     75 allocs/op
    reform:     0.59s       147548 ns/op    3198 B/op     86 allocs/op
       pop:     0.63s       156263 ns/op    3506 B/op     71 allocs/op
      gorp:     0.64s       160546 ns/op    3878 B/op    194 allocs/op
 gorm_prep:     0.66s       163799 ns/op    4802 B/op     84 allocs/op
       ent:     0.66s       165331 ns/op    5404 B/op    149 allocs/op
      sqlc:     0.67s       166277 ns/op    2172 B/op     52 allocs/op
        pg:     0.68s       170618 ns/op     872 B/op     20 allocs/op
 sqlboiler:     0.71s       177735 ns/op     956 B/op     14 allocs/op
       dbr:     0.72s       180737 ns/op    2184 B/op     37 allocs/op
       bun:     0.72s       180982 ns/op    5489 B/op     21 allocs/op
       rel:     0.85s       213637 ns/op    1800 B/op     45 allocs/op
      gorm:     0.91s       226846 ns/op    5150 B/op     95 allocs/op
      sqlx:     1.25s       311679 ns/op    1976 B/op     43 allocs/op
      godb:     1.30s       326147 ns/op    4112 B/op    102 allocs/op
      xorm:     1.37s       342931 ns/op    4616 B/op    125 allocs/op
     upper:     1.58s       395865 ns/op    8136 B/op    315 allocs/op

  4000 times - MultiRead limit 100
    reform:     0.56s       140210 ns/op    2926 B/op     74 allocs/op
       pgx:     1.07s       266615 ns/op   42949 B/op    504 allocs/op
  pgx_pool:     1.07s       268332 ns/op   43006 B/op    504 allocs/op
       raw:     1.37s       343608 ns/op   38341 B/op   1038 allocs/op
     upper:     1.38s       346050 ns/op    7841 B/op    294 allocs/op
 sqlboiler:     1.60s       400839 ns/op   57844 B/op   1259 allocs/op
      sqlc:     1.62s       403879 ns/op   73158 B/op   1251 allocs/op
        pg:     1.63s       406560 ns/op   23064 B/op    629 allocs/op
      gorp:     1.70s       424885 ns/op   57372 B/op   1494 allocs/op
       ent:     1.71s       426818 ns/op   76044 B/op   2039 allocs/op
      sqlx:     1.72s       430684 ns/op   39064 B/op   1516 allocs/op
       dbr:     1.95s       488398 ns/op   32416 B/op   1545 allocs/op
       bun:     1.97s       491281 ns/op   33992 B/op   1124 allocs/op
     beego:     2.02s       503857 ns/op   55183 B/op   3077 allocs/op
       pop:     2.09s       521265 ns/op   75716 B/op   1510 allocs/op
 gorm_prep:     2.22s       554050 ns/op   56070 B/op   2078 allocs/op
      gorm:     2.59s       647400 ns/op   57185 B/op   2378 allocs/op
      godb:     3.01s       751372 ns/op   75261 B/op   3084 allocs/op
       rel:     3.19s       797195 ns/op   95321 B/op   2248 allocs/op
      xorm:     3.26s       814717 ns/op  119954 B/op   4687 allocs/op
```
