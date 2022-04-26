# Results

- orm-benchmark (with no flags)

```
   
Reports:

   200 times - Insert
       pgx:     0.05s       227659 ns/op     244 B/op     10 allocs/op
     beego:     0.05s       233119 ns/op    2373 B/op     56 allocs/op
  pgx_pool:     0.05s       233865 ns/op     265 B/op     10 allocs/op
       ent:     0.05s       253687 ns/op    4172 B/op     99 allocs/op
       raw:     0.05s       259695 ns/op    1007 B/op     15 allocs/op
        pg:     0.05s       262804 ns/op    6272 B/op     10 allocs/op
      gorp:     0.05s       266944 ns/op    2077 B/op     43 allocs/op
      sqlc:     0.05s       271219 ns/op    3176 B/op     64 allocs/op
    reform:     0.06s       275200 ns/op    2091 B/op     53 allocs/op
 gorm_prep:     0.06s       278577 ns/op    5605 B/op     69 allocs/op
 sqlboiler:     0.06s       281793 ns/op    1873 B/op     37 allocs/op
      gorm:     0.06s       301040 ns/op    7132 B/op     89 allocs/op
       bun:     0.06s       301429 ns/op    5293 B/op     14 allocs/op
       dbr:     0.07s       361344 ns/op    2797 B/op     66 allocs/op
      sqlx:     0.07s       372841 ns/op     856 B/op     19 allocs/op
       rel:     0.08s       399056 ns/op    2512 B/op     41 allocs/op
      xorm:     0.08s       415945 ns/op    3442 B/op     90 allocs/op
      godb:     0.09s       429720 ns/op    4668 B/op    116 allocs/op
     upper:     0.10s       477106 ns/op   14199 B/op    685 allocs/op
       pop:     0.12s       602909 ns/op   10407 B/op    250 allocs/op

   200 times - MultiInsert 100 row
  pgx_pool:     0.18s       906329 ns/op  114321 B/op     42 allocs/op
       pgx:     0.18s       918736 ns/op  114301 B/op     42 allocs/op
     beego:     0.22s      1096581 ns/op  179509 B/op   2746 allocs/op
       raw:     0.24s      1181797 ns/op  192353 B/op    938 allocs/op
 gorm_prep:     0.25s      1257600 ns/op  234549 B/op   2182 allocs/op
    reform:     0.27s      1325980 ns/op  466410 B/op   2747 allocs/op
       ent:     0.29s      1427453 ns/op  412422 B/op   4902 allocs/op
       bun:     0.32s      1622840 ns/op   42358 B/op    217 allocs/op
        pg:     0.32s      1624433 ns/op    3316 B/op    112 allocs/op
      sqlx:     0.34s      1683967 ns/op  171536 B/op   1552 allocs/op
      gorm:     0.36s      1776964 ns/op  272129 B/op   3728 allocs/op
      xorm:     0.40s      2007348 ns/op  254772 B/op   5416 allocs/op
       rel:     0.43s      2152288 ns/op  303975 B/op   3261 allocs/op
     upper:     0.55s      2768126 ns/op  547000 B/op  19610 allocs/op
      godb:     0.90s      4501430 ns/op  259957 B/op   5893 allocs/op
      gorp:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert

   200 times - Update
       raw:     0.02s       115042 ns/op     722 B/op     13 allocs/op
       ent:     0.03s       131612 ns/op    4529 B/op     99 allocs/op
       pgx:     0.04s       224131 ns/op     282 B/op     10 allocs/op
      sqlc:     0.05s       226022 ns/op     890 B/op     14 allocs/op
    reform:     0.05s       231072 ns/op    1744 B/op     51 allocs/op
      gorp:     0.05s       231741 ns/op    1221 B/op     32 allocs/op
     beego:     0.05s       233325 ns/op    1755 B/op     47 allocs/op
 sqlboiler:     0.05s       234181 ns/op     925 B/op     17 allocs/op
 gorm_prep:     0.05s       241294 ns/op    5132 B/op     59 allocs/op
  pgx_pool:     0.05s       242452 ns/op     302 B/op     10 allocs/op
       pop:     0.05s       260647 ns/op    6558 B/op    198 allocs/op
      sqlx:     0.05s       262964 ns/op     872 B/op     20 allocs/op
       dbr:     0.05s       268681 ns/op    2650 B/op     57 allocs/op
        pg:     0.05s       274966 ns/op     768 B/op      9 allocs/op
       bun:     0.06s       278942 ns/op    4728 B/op      5 allocs/op
      gorm:     0.06s       293094 ns/op    6625 B/op     81 allocs/op
      xorm:     0.08s       394444 ns/op    3648 B/op    126 allocs/op
       rel:     0.08s       410771 ns/op    2531 B/op     41 allocs/op
      godb:     0.08s       424823 ns/op    5160 B/op    154 allocs/op
     upper:     0.21s      1066349 ns/op   33647 B/op   1510 allocs/op

   200 times - Read
       pgx:     0.02s       112913 ns/op     886 B/op      8 allocs/op
  pgx_pool:     0.02s       114234 ns/op    1079 B/op      9 allocs/op
      sqlc:     0.02s       123738 ns/op    2162 B/op     52 allocs/op
       raw:     0.02s       124607 ns/op    2098 B/op     50 allocs/op
     beego:     0.03s       127367 ns/op    2094 B/op     75 allocs/op
    reform:     0.03s       131583 ns/op    3233 B/op     86 allocs/op
      gorp:     0.03s       137122 ns/op    3870 B/op    194 allocs/op
       pop:     0.03s       137914 ns/op    3914 B/op     71 allocs/op
 gorm_prep:     0.03s       141372 ns/op    4800 B/op     83 allocs/op
        pg:     0.03s       145149 ns/op     873 B/op     20 allocs/op
       ent:     0.03s       149792 ns/op    5422 B/op    149 allocs/op
       bun:     0.03s       151453 ns/op    5490 B/op     21 allocs/op
 sqlboiler:     0.03s       154175 ns/op    1303 B/op     14 allocs/op
       rel:     0.03s       154445 ns/op    1803 B/op     45 allocs/op
       dbr:     0.03s       154750 ns/op    2184 B/op     37 allocs/op
      gorm:     0.04s       179429 ns/op    5154 B/op     94 allocs/op
      sqlx:     0.05s       269331 ns/op    1976 B/op     43 allocs/op
      godb:     0.06s       286652 ns/op    4112 B/op    102 allocs/op
      xorm:     0.06s       288803 ns/op    4616 B/op    125 allocs/op
     upper:     0.07s       329554 ns/op    8127 B/op    315 allocs/op

   200 times - MultiRead limit 100
    reform:     0.02s       124165 ns/op    2955 B/op     74 allocs/op
       pgx:     0.05s       231341 ns/op   42919 B/op    504 allocs/op
  pgx_pool:     0.05s       242914 ns/op   42993 B/op    504 allocs/op
     upper:     0.06s       291840 ns/op    7896 B/op    294 allocs/op
       raw:     0.06s       300628 ns/op   38352 B/op   1038 allocs/op
      sqlc:     0.07s       352281 ns/op   73128 B/op   1251 allocs/op
 sqlboiler:     0.07s       370133 ns/op   57840 B/op   1259 allocs/op
        pg:     0.07s       374806 ns/op   22258 B/op    629 allocs/op
       ent:     0.08s       408010 ns/op   76036 B/op   2039 allocs/op
      gorp:     0.08s       409467 ns/op   57386 B/op   1494 allocs/op
       dbr:     0.09s       441958 ns/op   32416 B/op   1545 allocs/op
      sqlx:     0.09s       445937 ns/op   39064 B/op   1516 allocs/op
       bun:     0.09s       446426 ns/op   33997 B/op   1124 allocs/op
     beego:     0.10s       483329 ns/op   55187 B/op   3077 allocs/op
       pop:     0.10s       491834 ns/op   76055 B/op   1511 allocs/op
 gorm_prep:     0.10s       496336 ns/op   56071 B/op   1978 allocs/op
      gorm:     0.11s       558407 ns/op   57185 B/op   2278 allocs/op
       rel:     0.13s       660669 ns/op   95321 B/op   2248 allocs/op
      godb:     0.14s       707378 ns/op   75265 B/op   3084 allocs/op
      xorm:     0.15s       772369 ns/op  119976 B/op   4687 allocs/op
```

- orm-benchmark -multi=5

```
  
Reports:

  1000 times - Insert
  pgx_pool:     0.21s       213147 ns/op     281 B/op     10 allocs/op
       pgx:     0.21s       214230 ns/op     265 B/op     10 allocs/op
       raw:     0.23s       233896 ns/op     764 B/op     13 allocs/op
    reform:     0.24s       243278 ns/op    1837 B/op     51 allocs/op
 sqlboiler:     0.24s       243895 ns/op    1621 B/op     35 allocs/op
     beego:     0.25s       248434 ns/op    2375 B/op     56 allocs/op
       ent:     0.25s       249987 ns/op    4214 B/op    100 allocs/op
      gorp:     0.25s       254728 ns/op    1827 B/op     42 allocs/op
      sqlc:     0.26s       255774 ns/op    2929 B/op     63 allocs/op
 gorm_prep:     0.26s       260731 ns/op    5341 B/op     67 allocs/op
        pg:     0.27s       266853 ns/op    1889 B/op     10 allocs/op
       dbr:     0.29s       290591 ns/op    2709 B/op     65 allocs/op
       bun:     0.30s       295891 ns/op    5053 B/op     14 allocs/op
      gorm:     0.30s       301646 ns/op    6938 B/op     90 allocs/op
      sqlx:     0.36s       356676 ns/op     856 B/op     19 allocs/op
       rel:     0.38s       380244 ns/op    2504 B/op     42 allocs/op
      xorm:     0.40s       399327 ns/op    3344 B/op     89 allocs/op
      godb:     0.42s       421242 ns/op    4576 B/op    116 allocs/op
     upper:     0.46s       458549 ns/op   13842 B/op    675 allocs/op
       pop:     0.60s       601792 ns/op   10124 B/op    249 allocs/op

  1000 times - MultiInsert 100 row
  pgx_pool:     0.94s       937005 ns/op  114331 B/op     43 allocs/op
       pgx:     0.95s       950949 ns/op  114287 B/op     42 allocs/op
       raw:     1.06s      1058302 ns/op  191429 B/op    932 allocs/op
     beego:     1.12s      1122218 ns/op  179427 B/op   2746 allocs/op
    reform:     1.19s      1185983 ns/op  466297 B/op   2747 allocs/op
 gorm_prep:     1.28s      1284534 ns/op  234469 B/op   2183 allocs/op
       ent:     1.50s      1498473 ns/op  412166 B/op   4900 allocs/op
        pg:     1.51s      1506310 ns/op    3314 B/op    112 allocs/op
       bun:     1.66s      1657991 ns/op   42256 B/op    218 allocs/op
      sqlx:     1.74s      1741525 ns/op  171372 B/op   1552 allocs/op
      gorm:     1.90s      1904765 ns/op  272187 B/op   3729 allocs/op
      xorm:     2.02s      2019010 ns/op  254758 B/op   5416 allocs/op
       rel:     2.09s      2090682 ns/op  303992 B/op   3263 allocs/op
      godb:     2.15s      2148731 ns/op  259963 B/op   5894 allocs/op
     upper:     2.80s      2799364 ns/op  545452 B/op  19495 allocs/op
       dbr:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert

  1000 times - Update
       raw:     0.12s       117789 ns/op     753 B/op     13 allocs/op
       ent:     0.13s       129969 ns/op    4561 B/op     99 allocs/op
       pgx:     0.22s       220910 ns/op     272 B/op     10 allocs/op
  pgx_pool:     0.23s       227238 ns/op     287 B/op     10 allocs/op
      sqlc:     0.23s       231636 ns/op     880 B/op     14 allocs/op
 sqlboiler:     0.24s       238943 ns/op     907 B/op     17 allocs/op
    reform:     0.24s       241285 ns/op    1777 B/op     51 allocs/op
 gorm_prep:     0.26s       258963 ns/op    5113 B/op     59 allocs/op
      sqlx:     0.26s       260803 ns/op     872 B/op     20 allocs/op
       pop:     0.27s       270599 ns/op    6586 B/op    198 allocs/op
     beego:     0.27s       271000 ns/op    1753 B/op     47 allocs/op
      gorp:     0.27s       271703 ns/op    1209 B/op     32 allocs/op
       dbr:     0.28s       275231 ns/op    2651 B/op     57 allocs/op
       bun:     0.29s       289020 ns/op    4728 B/op      5 allocs/op
        pg:     0.29s       294181 ns/op     768 B/op      9 allocs/op
      gorm:     0.32s       317154 ns/op    6600 B/op     81 allocs/op
       rel:     0.39s       390330 ns/op    2528 B/op     41 allocs/op
      xorm:     0.40s       396787 ns/op    3648 B/op    126 allocs/op
      godb:     0.43s       426843 ns/op    5160 B/op    154 allocs/op
     upper:     1.05s      1046664 ns/op   33530 B/op   1503 allocs/op

  1000 times - Read
  pgx_pool:     0.11s       113070 ns/op    1076 B/op      9 allocs/op
       pgx:     0.11s       113278 ns/op     889 B/op      8 allocs/op
       raw:     0.12s       119880 ns/op    2057 B/op     50 allocs/op
     beego:     0.13s       125629 ns/op    2089 B/op     75 allocs/op
    reform:     0.13s       131600 ns/op    3193 B/op     86 allocs/op
      sqlc:     0.13s       132532 ns/op    2176 B/op     52 allocs/op
       pop:     0.14s       137127 ns/op    3561 B/op     71 allocs/op
 gorm_prep:     0.14s       139860 ns/op    4806 B/op     83 allocs/op
      gorp:     0.14s       140472 ns/op    3881 B/op    194 allocs/op
        pg:     0.14s       141906 ns/op     872 B/op     20 allocs/op
 sqlboiler:     0.15s       145262 ns/op     971 B/op     14 allocs/op
       ent:     0.15s       147921 ns/op    5401 B/op    149 allocs/op
       bun:     0.15s       149042 ns/op    5489 B/op     21 allocs/op
       rel:     0.15s       153026 ns/op    1800 B/op     45 allocs/op
       dbr:     0.16s       155114 ns/op    2184 B/op     37 allocs/op
      gorm:     0.19s       187377 ns/op    5157 B/op     94 allocs/op
      sqlx:     0.27s       265711 ns/op    1976 B/op     43 allocs/op
      xorm:     0.29s       285904 ns/op    4616 B/op    125 allocs/op
      godb:     0.29s       288490 ns/op    4112 B/op    102 allocs/op
     upper:     0.30s       304964 ns/op    8137 B/op    315 allocs/op

  1000 times - MultiRead limit 100
    reform:     0.13s       127340 ns/op    2920 B/op     74 allocs/op
  pgx_pool:     0.23s       233691 ns/op   43009 B/op    504 allocs/op
       pgx:     0.24s       243298 ns/op   42949 B/op    504 allocs/op
     upper:     0.30s       300343 ns/op    7849 B/op    294 allocs/op
       raw:     0.30s       301182 ns/op   38342 B/op   1038 allocs/op
      sqlc:     0.34s       340692 ns/op   73158 B/op   1251 allocs/op
        pg:     0.37s       367760 ns/op   22259 B/op    629 allocs/op
 sqlboiler:     0.38s       376155 ns/op   57933 B/op   1259 allocs/op
      gorp:     0.41s       407976 ns/op   57373 B/op   1494 allocs/op
      sqlx:     0.41s       408359 ns/op   39064 B/op   1516 allocs/op
       ent:     0.41s       410656 ns/op   76050 B/op   2039 allocs/op
       dbr:     0.44s       437467 ns/op   32416 B/op   1545 allocs/op
       bun:     0.46s       458272 ns/op   33992 B/op   1124 allocs/op
     beego:     0.47s       466516 ns/op   55183 B/op   3077 allocs/op
       pop:     0.48s       477759 ns/op   75753 B/op   1510 allocs/op
 gorm_prep:     0.49s       494724 ns/op   56057 B/op   1977 allocs/op
      gorm:     0.56s       555923 ns/op   57164 B/op   2277 allocs/op
       rel:     0.66s       659139 ns/op   95321 B/op   2248 allocs/op
      godb:     0.70s       703385 ns/op   75264 B/op   3084 allocs/op
      xorm:     0.76s       755017 ns/op  119953 B/op   4687 allocs/op
```

- orm-benchmark -multi=10

```
  
Reports:

  2000 times - Insert
       pgx:     0.42s       211974 ns/op     269 B/op     10 allocs/op
  pgx_pool:     0.43s       214503 ns/op     285 B/op     10 allocs/op
       raw:     0.45s       226986 ns/op     734 B/op     13 allocs/op
      gorp:     0.48s       238847 ns/op    1799 B/op     42 allocs/op
 sqlboiler:     0.49s       244374 ns/op    1591 B/op     35 allocs/op
     beego:     0.49s       244453 ns/op    2375 B/op     56 allocs/op
    reform:     0.49s       246274 ns/op    1807 B/op     51 allocs/op
       ent:     0.49s       246755 ns/op    4215 B/op    100 allocs/op
 gorm_prep:     0.52s       259752 ns/op    5295 B/op     67 allocs/op
      sqlc:     0.53s       266253 ns/op    2900 B/op     63 allocs/op
       dbr:     0.56s       279049 ns/op    2698 B/op     65 allocs/op
        pg:     0.56s       280153 ns/op    1341 B/op     10 allocs/op
       bun:     0.57s       283575 ns/op    5022 B/op     14 allocs/op
      gorm:     0.63s       315153 ns/op    6930 B/op     89 allocs/op
      sqlx:     0.76s       378484 ns/op     856 B/op     19 allocs/op
       rel:     0.78s       389284 ns/op    2504 B/op     42 allocs/op
      xorm:     0.80s       400069 ns/op    3335 B/op     89 allocs/op
      godb:     0.81s       402792 ns/op    4564 B/op    116 allocs/op
     upper:     0.93s       462709 ns/op   13800 B/op    674 allocs/op
       pop:     1.25s       625976 ns/op   10091 B/op    249 allocs/op

  2000 times - MultiInsert 100 row
       pgx:     1.86s       931565 ns/op  114290 B/op     43 allocs/op
  pgx_pool:     1.89s       943341 ns/op  114305 B/op     42 allocs/op
       raw:     2.07s      1035705 ns/op  191315 B/op    931 allocs/op
     beego:     2.29s      1144813 ns/op  179425 B/op   2746 allocs/op
    reform:     2.45s      1223166 ns/op  466292 B/op   2747 allocs/op
 gorm_prep:     2.53s      1262714 ns/op  234449 B/op   2183 allocs/op
       ent:     3.00s      1499753 ns/op  412140 B/op   4900 allocs/op
        pg:     3.04s      1521097 ns/op    3314 B/op    112 allocs/op
       bun:     3.26s      1628274 ns/op   42244 B/op    218 allocs/op
      sqlx:     3.43s      1714201 ns/op  171184 B/op   1552 allocs/op
      gorm:     3.61s      1803786 ns/op  272232 B/op   3729 allocs/op
      xorm:     4.07s      2036594 ns/op  254773 B/op   5416 allocs/op
       rel:     4.14s      2069462 ns/op  303994 B/op   3263 allocs/op
      godb:     4.31s      2154977 ns/op  259960 B/op   5895 allocs/op
     upper:     5.59s      2793942 ns/op  545255 B/op  19480 allocs/op
 sqlboiler:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert

  2000 times - Update
       raw:     0.24s       117793 ns/op     748 B/op     13 allocs/op
       ent:     0.26s       131933 ns/op    4556 B/op     99 allocs/op
       pgx:     0.43s       214510 ns/op     268 B/op     10 allocs/op
      sqlc:     0.45s       223540 ns/op     876 B/op     14 allocs/op
  pgx_pool:     0.45s       225987 ns/op     285 B/op     10 allocs/op
      gorp:     0.46s       231673 ns/op    1205 B/op     32 allocs/op
 sqlboiler:     0.47s       237129 ns/op     901 B/op     17 allocs/op
     beego:     0.47s       237258 ns/op    1752 B/op     47 allocs/op
    reform:     0.48s       241969 ns/op    1773 B/op     51 allocs/op
      sqlx:     0.52s       262024 ns/op     872 B/op     20 allocs/op
       dbr:     0.54s       267742 ns/op    2651 B/op     57 allocs/op
       pop:     0.54s       267796 ns/op    6582 B/op    198 allocs/op
 gorm_prep:     0.55s       275853 ns/op    5121 B/op     59 allocs/op
        pg:     0.56s       279798 ns/op     768 B/op      9 allocs/op
       bun:     0.57s       287305 ns/op    4728 B/op      5 allocs/op
      gorm:     0.68s       342437 ns/op    6604 B/op     81 allocs/op
       rel:     0.81s       405405 ns/op    2528 B/op     41 allocs/op
      xorm:     0.81s       405510 ns/op    3648 B/op    126 allocs/op
      godb:     0.85s       425699 ns/op    5160 B/op    154 allocs/op
     upper:     2.12s      1059031 ns/op   33503 B/op   1502 allocs/op

  2000 times - Read
  pgx_pool:     0.23s       115458 ns/op    1080 B/op      9 allocs/op
       pgx:     0.24s       118775 ns/op     892 B/op      8 allocs/op
       raw:     0.24s       120674 ns/op    2065 B/op     50 allocs/op
      sqlc:     0.24s       120807 ns/op    2172 B/op     52 allocs/op
     beego:     0.25s       125455 ns/op    2088 B/op     75 allocs/op
      gorp:     0.27s       132854 ns/op    3876 B/op    194 allocs/op
    reform:     0.27s       135226 ns/op    3197 B/op     86 allocs/op
 gorm_prep:     0.28s       140332 ns/op    4800 B/op     83 allocs/op
       pop:     0.29s       144513 ns/op    3529 B/op     71 allocs/op
        pg:     0.29s       145484 ns/op     872 B/op     20 allocs/op
       ent:     0.30s       148251 ns/op    5405 B/op    149 allocs/op
       rel:     0.30s       149040 ns/op    1800 B/op     45 allocs/op
 sqlboiler:     0.30s       149477 ns/op     974 B/op     14 allocs/op
       bun:     0.30s       150110 ns/op    5489 B/op     21 allocs/op
       dbr:     0.30s       151946 ns/op    2184 B/op     37 allocs/op
      gorm:     0.37s       183055 ns/op    5151 B/op     94 allocs/op
      sqlx:     0.54s       268799 ns/op    1976 B/op     43 allocs/op
      godb:     0.58s       287847 ns/op    4112 B/op    102 allocs/op
      xorm:     0.59s       294857 ns/op    4616 B/op    125 allocs/op
     upper:     0.61s       303345 ns/op    8135 B/op    315 allocs/op

  2000 times - MultiRead limit 100
    reform:     0.25s       125341 ns/op    2924 B/op     74 allocs/op
  pgx_pool:     0.47s       233691 ns/op   43009 B/op    504 allocs/op
       pgx:     0.48s       239535 ns/op   42949 B/op    504 allocs/op
     upper:     0.59s       294760 ns/op    7841 B/op    294 allocs/op
       raw:     0.61s       303145 ns/op   38342 B/op   1038 allocs/op
      sqlc:     0.69s       346043 ns/op   73158 B/op   1251 allocs/op
        pg:     0.72s       360845 ns/op   22821 B/op    629 allocs/op
 sqlboiler:     0.75s       373051 ns/op   57836 B/op   1259 allocs/op
      sqlx:     0.81s       402690 ns/op   39064 B/op   1516 allocs/op
       ent:     0.81s       404547 ns/op   76049 B/op   2039 allocs/op
      gorp:     0.83s       417013 ns/op   57372 B/op   1494 allocs/op
       dbr:     0.90s       449592 ns/op   32416 B/op   1545 allocs/op
       bun:     0.91s       452726 ns/op   33992 B/op   1124 allocs/op
     beego:     0.94s       468948 ns/op   55184 B/op   3077 allocs/op
       pop:     0.97s       487404 ns/op   75745 B/op   1510 allocs/op
 gorm_prep:     0.99s       496074 ns/op   56052 B/op   1977 allocs/op
      gorm:     1.10s       548895 ns/op   57153 B/op   2277 allocs/op
       rel:     1.34s       668031 ns/op   95321 B/op   2248 allocs/op
      godb:     1.37s       684509 ns/op   75264 B/op   3084 allocs/op
      xorm:     1.53s       763945 ns/op  119955 B/op   4687 allocs/op
```

- orm-benchmark -multi=20

```
  
Reports:

  4000 times - Insert
  pgx_pool:     0.86s       213805 ns/op     286 B/op     10 allocs/op
       pgx:     0.86s       213847 ns/op     270 B/op     10 allocs/op
       raw:     0.90s       223869 ns/op     719 B/op     13 allocs/op
 sqlboiler:     0.97s       243313 ns/op    1574 B/op     35 allocs/op
     beego:     0.97s       243408 ns/op    2376 B/op     56 allocs/op
       ent:     0.98s       243952 ns/op    4216 B/op    100 allocs/op
    reform:     0.98s       244010 ns/op    1790 B/op     51 allocs/op
      sqlc:     0.99s       247987 ns/op    2884 B/op     63 allocs/op
      gorp:     1.02s       255649 ns/op    1782 B/op     42 allocs/op
        pg:     1.06s       265901 ns/op     806 B/op     10 allocs/op
 gorm_prep:     1.06s       266181 ns/op    5257 B/op     67 allocs/op
       dbr:     1.10s       275919 ns/op    2693 B/op     65 allocs/op
       bun:     1.15s       286672 ns/op    5008 B/op     14 allocs/op
      gorm:     1.21s       302438 ns/op    6930 B/op     89 allocs/op
      sqlx:     1.46s       364646 ns/op     856 B/op     19 allocs/op
      godb:     1.58s       395900 ns/op    4560 B/op    116 allocs/op
       rel:     1.61s       402689 ns/op    2504 B/op     42 allocs/op
      xorm:     1.69s       423269 ns/op    3326 B/op     89 allocs/op
     upper:     1.90s       474438 ns/op   13775 B/op    673 allocs/op
       pop:     2.40s       598972 ns/op   10072 B/op    249 allocs/op

  4000 times - MultiInsert 100 row
  pgx_pool:     3.75s       936659 ns/op  114308 B/op     43 allocs/op
       pgx:     3.76s       940823 ns/op  114289 B/op     42 allocs/op
       raw:     4.01s      1003174 ns/op  191258 B/op    931 allocs/op
     beego:     4.45s      1113290 ns/op  179430 B/op   2746 allocs/op
    reform:     4.84s      1210254 ns/op  466307 B/op   2747 allocs/op
 gorm_prep:     4.89s      1222283 ns/op  234453 B/op   2183 allocs/op
       ent:     5.59s      1396607 ns/op  412127 B/op   4900 allocs/op
        pg:     5.88s      1469090 ns/op    3314 B/op    112 allocs/op
       bun:     6.50s      1625545 ns/op   42241 B/op    219 allocs/op
      sqlx:     6.68s      1669944 ns/op  171160 B/op   1552 allocs/op
      gorm:     7.32s      1829291 ns/op  272251 B/op   3729 allocs/op
       rel:     8.03s      2008130 ns/op  303994 B/op   3263 allocs/op
      xorm:     8.19s      2047216 ns/op  254789 B/op   5416 allocs/op
      godb:     8.43s      2106831 ns/op  259973 B/op   5895 allocs/op
     upper:    11.00s      2749454 ns/op  545162 B/op  19473 allocs/op
       pop:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert

  4000 times - Update
       raw:     0.47s       118723 ns/op     750 B/op     13 allocs/op
       ent:     0.53s       132360 ns/op    4558 B/op     99 allocs/op
       pgx:     0.90s       225363 ns/op     270 B/op     10 allocs/op
      sqlc:     0.92s       230925 ns/op     878 B/op     14 allocs/op
  pgx_pool:     0.94s       234537 ns/op     286 B/op     10 allocs/op
    reform:     0.96s       240249 ns/op    1774 B/op     51 allocs/op
      gorp:     0.97s       242332 ns/op    1204 B/op     32 allocs/op
 sqlboiler:     0.99s       246359 ns/op     901 B/op     17 allocs/op
     beego:     0.99s       246944 ns/op    1752 B/op     47 allocs/op
 gorm_prep:     1.01s       253530 ns/op    5116 B/op     59 allocs/op
      sqlx:     1.06s       264848 ns/op     872 B/op     20 allocs/op
        pg:     1.10s       275449 ns/op     768 B/op      9 allocs/op
       dbr:     1.12s       281180 ns/op    2651 B/op     57 allocs/op
       pop:     1.15s       288051 ns/op    6582 B/op    198 allocs/op
       bun:     1.17s       291689 ns/op    4728 B/op      5 allocs/op
      gorm:     1.26s       315009 ns/op    6604 B/op     81 allocs/op
      xorm:     1.61s       403578 ns/op    3648 B/op    126 allocs/op
       rel:     1.64s       409590 ns/op    2528 B/op     41 allocs/op
      godb:     1.71s       427625 ns/op    5160 B/op    154 allocs/op
     upper:     4.26s      1064508 ns/op   33492 B/op   1502 allocs/op

  4000 times - Read
       pgx:     0.46s       115624 ns/op     894 B/op      8 allocs/op
  pgx_pool:     0.47s       118368 ns/op    1079 B/op      9 allocs/op
       raw:     0.49s       121255 ns/op    2062 B/op     50 allocs/op
      sqlc:     0.49s       121809 ns/op    2172 B/op     52 allocs/op
    reform:     0.53s       131346 ns/op    3198 B/op     86 allocs/op
     beego:     0.53s       133047 ns/op    2088 B/op     75 allocs/op
      gorp:     0.53s       133170 ns/op    3878 B/op    194 allocs/op
       pop:     0.55s       137935 ns/op    3507 B/op     71 allocs/op
 gorm_prep:     0.56s       140093 ns/op    4801 B/op     83 allocs/op
        pg:     0.58s       144409 ns/op     872 B/op     20 allocs/op
       ent:     0.59s       147517 ns/op    5404 B/op    149 allocs/op
 sqlboiler:     0.60s       149544 ns/op     956 B/op     14 allocs/op
       bun:     0.60s       150471 ns/op    5489 B/op     21 allocs/op
       dbr:     0.62s       154359 ns/op    2184 B/op     37 allocs/op
       rel:     0.62s       155049 ns/op    1800 B/op     45 allocs/op
      gorm:     0.75s       188690 ns/op    5148 B/op     94 allocs/op
      sqlx:     1.10s       274103 ns/op    1976 B/op     43 allocs/op
      godb:     1.16s       289814 ns/op    4112 B/op    102 allocs/op
      xorm:     1.18s       294857 ns/op    4616 B/op    125 allocs/op
     upper:     1.21s       302740 ns/op    8136 B/op    315 allocs/op

  4000 times - MultiRead limit 100
    reform:     0.50s       125520 ns/op    2926 B/op     74 allocs/op
       pgx:     0.94s       234146 ns/op   42949 B/op    504 allocs/op
  pgx_pool:     0.94s       235656 ns/op   43006 B/op    504 allocs/op
     upper:     1.19s       297053 ns/op    7841 B/op    294 allocs/op
       raw:     1.23s       308142 ns/op   38341 B/op   1038 allocs/op
      sqlc:     1.35s       336517 ns/op   73158 B/op   1251 allocs/op
        pg:     1.43s       356783 ns/op   22260 B/op    629 allocs/op
 sqlboiler:     1.49s       373184 ns/op   57814 B/op   1259 allocs/op
       ent:     1.59s       397462 ns/op   76045 B/op   2039 allocs/op
      gorp:     1.62s       405264 ns/op   57370 B/op   1494 allocs/op
      sqlx:     1.63s       406548 ns/op   39064 B/op   1516 allocs/op
       bun:     1.80s       451186 ns/op   33997 B/op   1124 allocs/op
       dbr:     1.84s       461176 ns/op   32416 B/op   1545 allocs/op
       pop:     1.92s       480882 ns/op   75719 B/op   1510 allocs/op
     beego:     1.93s       482122 ns/op   55184 B/op   3077 allocs/op
 gorm_prep:     1.95s       488301 ns/op   56053 B/op   1977 allocs/op
      gorm:     2.23s       556772 ns/op   57142 B/op   2277 allocs/op
       rel:     2.63s       658713 ns/op   95321 B/op   2248 allocs/op
      godb:     2.80s       699485 ns/op   75266 B/op   3084 allocs/op
      xorm:     3.08s       771169 ns/op  119954 B/op   4687 allocs/op
```
