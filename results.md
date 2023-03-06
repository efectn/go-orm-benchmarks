# Results

- orm-benchmark (with no flags)

```
   
Reports:

   200 times - Insert
  pgx_pool:     0.05s       272472 ns/op     265 B/op     10 allocs/op
       pgx:     0.06s       283890 ns/op     244 B/op     10 allocs/op
     beego:     0.06s       300738 ns/op    2372 B/op     56 allocs/op
    reform:     0.07s       351180 ns/op    2085 B/op     53 allocs/op
       raw:     0.07s       357407 ns/op    1008 B/op     15 allocs/op
       dbr:     0.07s       360904 ns/op    2800 B/op     66 allocs/op
 gorm_prep:     0.07s       362619 ns/op    5803 B/op     70 allocs/op
      gorp:     0.07s       366083 ns/op    2068 B/op     43 allocs/op
        pg:     0.07s       367937 ns/op    6275 B/op     11 allocs/op
       ent:     0.07s       368552 ns/op    4116 B/op     97 allocs/op
 sqlboiler:     0.08s       383749 ns/op    1867 B/op     36 allocs/op
      sqlc:     0.08s       386453 ns/op    3172 B/op     64 allocs/op
       bun:     0.08s       393508 ns/op    5284 B/op     14 allocs/op
      gorm:     0.08s       394725 ns/op    6952 B/op     90 allocs/op
       rel:     0.10s       484160 ns/op    2590 B/op     43 allocs/op
     upper:     0.10s       485793 ns/op    6210 B/op    130 allocs/op
      xorm:     0.10s       501884 ns/op    3444 B/op     90 allocs/op
      godb:     0.10s       503588 ns/op    4670 B/op    116 allocs/op
      zorm:     0.10s       508019 ns/op    3801 B/op     76 allocs/op
      sqlx:     0.11s       564941 ns/op     856 B/op     19 allocs/op
       pop:     0.15s       739509 ns/op    9960 B/op    238 allocs/op

   200 times - MultiInsert 100 row
  pgx_pool:     0.21s      1069001 ns/op  112996 B/op     43 allocs/op
       raw:     0.24s      1200415 ns/op  184802 B/op    936 allocs/op
       pgx:     0.24s      1224278 ns/op  112961 B/op     43 allocs/op
 gorm_prep:     0.29s      1447026 ns/op  227792 B/op   2279 allocs/op
    reform:     0.29s      1472451 ns/op  458861 B/op   2745 allocs/op
     beego:     0.30s      1482628 ns/op  177650 B/op   2745 allocs/op
       bun:     0.33s      1643274 ns/op   42487 B/op    217 allocs/op
       ent:     0.33s      1655447 ns/op  387613 B/op   4600 allocs/op
        pg:     0.34s      1707059 ns/op    8929 B/op    112 allocs/op
      gorm:     0.38s      1902264 ns/op  271117 B/op   3827 allocs/op
      zorm:     0.40s      1987462 ns/op  199928 B/op   2780 allocs/op
     upper:     0.45s      2228512 ns/op  323194 B/op   4218 allocs/op
      godb:     0.46s      2307552 ns/op  255222 B/op   5892 allocs/op
      xorm:     0.50s      2499385 ns/op  248052 B/op   5414 allocs/op
       rel:     0.51s      2545526 ns/op  306852 B/op   3261 allocs/op
      sqlx:     0.65s      3229398 ns/op  171478 B/op   1551 allocs/op
 sqlboiler:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert

   200 times - Update
       raw:     0.03s       134517 ns/op     721 B/op     13 allocs/op
       ent:     0.03s       155828 ns/op    4618 B/op     98 allocs/op
       pgx:     0.05s       274140 ns/op     282 B/op     10 allocs/op
      sqlx:     0.06s       291976 ns/op     872 B/op     20 allocs/op
     beego:     0.06s       301782 ns/op    1758 B/op     47 allocs/op
      gorp:     0.06s       307746 ns/op    1221 B/op     32 allocs/op
 sqlboiler:     0.06s       322066 ns/op     926 B/op     17 allocs/op
      sqlc:     0.06s       322483 ns/op     891 B/op     14 allocs/op
  pgx_pool:     0.07s       329179 ns/op     302 B/op     10 allocs/op
    reform:     0.07s       332265 ns/op    1747 B/op     51 allocs/op
       bun:     0.07s       334371 ns/op    4728 B/op      5 allocs/op
       dbr:     0.07s       342667 ns/op    2651 B/op     57 allocs/op
       pop:     0.07s       347063 ns/op    6022 B/op    186 allocs/op
 gorm_prep:     0.07s       349937 ns/op    5133 B/op     59 allocs/op
        pg:     0.07s       370877 ns/op     768 B/op      9 allocs/op
      gorm:     0.08s       397517 ns/op    6624 B/op     81 allocs/op
      godb:     0.09s       470843 ns/op    5161 B/op    154 allocs/op
      xorm:     0.10s       483265 ns/op    3944 B/op    132 allocs/op
       rel:     0.10s       484505 ns/op    3027 B/op     44 allocs/op
      zorm:     0.10s       524154 ns/op    3024 B/op     59 allocs/op
     upper:     0.23s      1155221 ns/op   16686 B/op    391 allocs/op

   200 times - Read
  pgx_pool:     0.03s       129519 ns/op    1078 B/op      9 allocs/op
       pgx:     0.03s       130711 ns/op     886 B/op      8 allocs/op
      sqlc:     0.03s       143053 ns/op    2162 B/op     52 allocs/op
       raw:     0.03s       143317 ns/op    2098 B/op     50 allocs/op
     beego:     0.03s       149764 ns/op    2093 B/op     75 allocs/op
      gorp:     0.03s       151266 ns/op    3869 B/op    194 allocs/op
       pop:     0.03s       153889 ns/op    3384 B/op     67 allocs/op
    reform:     0.03s       154879 ns/op    3232 B/op     86 allocs/op
        pg:     0.03s       159724 ns/op     873 B/op     20 allocs/op
       ent:     0.03s       161329 ns/op    5426 B/op    150 allocs/op
       bun:     0.03s       164512 ns/op    5522 B/op     21 allocs/op
 gorm_prep:     0.03s       164724 ns/op    4816 B/op     85 allocs/op
       dbr:     0.03s       173131 ns/op    2184 B/op     37 allocs/op
       rel:     0.04s       175335 ns/op    2315 B/op     49 allocs/op
 sqlboiler:     0.04s       176497 ns/op    1116 B/op     14 allocs/op
      zorm:     0.04s       192541 ns/op    3072 B/op     67 allocs/op
      gorm:     0.04s       208571 ns/op    5153 B/op     96 allocs/op
     upper:     0.06s       321030 ns/op    5041 B/op    110 allocs/op
      xorm:     0.06s       321064 ns/op    4648 B/op    127 allocs/op
      sqlx:     0.06s       324312 ns/op    1976 B/op     43 allocs/op
      godb:     0.07s       328424 ns/op    4112 B/op    102 allocs/op

   200 times - MultiRead limit 100
    reform:     0.03s       146936 ns/op    2955 B/op     74 allocs/op
  pgx_pool:     0.06s       278454 ns/op   42994 B/op    504 allocs/op
     upper:     0.06s       310946 ns/op    4819 B/op     89 allocs/op
       pgx:     0.07s       326949 ns/op   42919 B/op    504 allocs/op
       raw:     0.07s       341500 ns/op   38353 B/op   1038 allocs/op
      sqlc:     0.08s       404070 ns/op   73128 B/op   1251 allocs/op
        pg:     0.09s       438270 ns/op   22263 B/op    629 allocs/op
       pop:     0.10s       478363 ns/op   68764 B/op   1307 allocs/op
       dbr:     0.10s       479346 ns/op   32416 B/op   1545 allocs/op
      sqlx:     0.10s       484435 ns/op   39064 B/op   1516 allocs/op
       ent:     0.10s       507517 ns/op   76126 B/op   2040 allocs/op
       bun:     0.10s       507830 ns/op   34026 B/op   1124 allocs/op
      gorp:     0.10s       514463 ns/op   57386 B/op   1494 allocs/op
     beego:     0.11s       527712 ns/op   55187 B/op   3077 allocs/op
 sqlboiler:     0.11s       566483 ns/op   66212 B/op   2259 allocs/op
 gorm_prep:     0.11s       574246 ns/op   56095 B/op   2080 allocs/op
      gorm:     0.13s       659895 ns/op   57190 B/op   2379 allocs/op
      zorm:     0.16s       794733 ns/op  163233 B/op   3241 allocs/op
      godb:     0.16s       814819 ns/op   75263 B/op   3084 allocs/op
       rel:     0.16s       824930 ns/op  100616 B/op   2252 allocs/op
      xorm:     0.17s       866291 ns/op  120201 B/op   4692 allocs/op
```

- orm-benchmark -multi=5

```
  
Reports:

  1000 times - Insert
  pgx_pool:     0.28s       278566 ns/op     281 B/op     10 allocs/op
       pgx:     0.28s       283174 ns/op     265 B/op     10 allocs/op
 sqlboiler:     0.32s       315871 ns/op    1621 B/op     35 allocs/op
      gorp:     0.32s       316357 ns/op    1826 B/op     42 allocs/op
     beego:     0.32s       316990 ns/op    2375 B/op     56 allocs/op
       raw:     0.32s       318280 ns/op     763 B/op     13 allocs/op
       ent:     0.33s       325062 ns/op    4118 B/op     98 allocs/op
       dbr:     0.34s       335262 ns/op    2710 B/op     65 allocs/op
 gorm_prep:     0.34s       339230 ns/op    5348 B/op     68 allocs/op
        pg:     0.35s       351278 ns/op    1889 B/op     10 allocs/op
      sqlc:     0.35s       352451 ns/op    2928 B/op     63 allocs/op
    reform:     0.36s       361342 ns/op    1837 B/op     51 allocs/op
       bun:     0.37s       371995 ns/op    5054 B/op     14 allocs/op
      gorm:     0.41s       413620 ns/op    6983 B/op     90 allocs/op
      sqlx:     0.46s       464280 ns/op     856 B/op     19 allocs/op
      xorm:     0.47s       471112 ns/op    3344 B/op     89 allocs/op
       rel:     0.48s       480896 ns/op    2584 B/op     44 allocs/op
      zorm:     0.50s       500711 ns/op    3786 B/op     77 allocs/op
     upper:     0.51s       514888 ns/op    5946 B/op    126 allocs/op
      godb:     0.53s       528114 ns/op    4576 B/op    116 allocs/op
       pop:     0.71s       712186 ns/op    9592 B/op    237 allocs/op

  1000 times - MultiInsert 100 row
  pgx_pool:     1.08s      1076238 ns/op  112945 B/op     43 allocs/op
       pgx:     1.11s      1108585 ns/op  112934 B/op     43 allocs/op
       raw:     1.19s      1191798 ns/op  183877 B/op    930 allocs/op
     beego:     1.32s      1319107 ns/op  177608 B/op   2745 allocs/op
 gorm_prep:     1.51s      1512248 ns/op  227702 B/op   2281 allocs/op
    reform:     1.53s      1525255 ns/op  458800 B/op   2745 allocs/op
        pg:     1.59s      1588637 ns/op    4437 B/op    112 allocs/op
       bun:     1.73s      1727634 ns/op   42417 B/op    218 allocs/op
       ent:     1.74s      1737467 ns/op  387395 B/op   4598 allocs/op
      gorm:     1.95s      1952759 ns/op  271496 B/op   3829 allocs/op
      sqlx:     1.96s      1959778 ns/op  170416 B/op   1551 allocs/op
      xorm:     2.25s      2250517 ns/op  248065 B/op   5414 allocs/op
      zorm:     2.30s      2300391 ns/op  199928 B/op   2780 allocs/op
       rel:     2.33s      2329495 ns/op  306869 B/op   3263 allocs/op
      godb:     2.40s      2400716 ns/op  255245 B/op   5894 allocs/op
     upper:     2.50s      2496332 ns/op  322950 B/op   4207 allocs/op
       pop:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert

  1000 times - Update
       raw:     0.14s       135112 ns/op     752 B/op     13 allocs/op
       ent:     0.15s       154166 ns/op    4609 B/op     98 allocs/op
  pgx_pool:     0.29s       285949 ns/op     287 B/op     10 allocs/op
       pgx:     0.29s       292741 ns/op     272 B/op     10 allocs/op
      gorp:     0.30s       295046 ns/op    1209 B/op     32 allocs/op
      sqlx:     0.30s       302182 ns/op     872 B/op     20 allocs/op
    reform:     0.30s       304071 ns/op    1777 B/op     51 allocs/op
      sqlc:     0.31s       305977 ns/op     880 B/op     14 allocs/op
     beego:     0.32s       320685 ns/op    1752 B/op     47 allocs/op
 sqlboiler:     0.32s       323442 ns/op     906 B/op     17 allocs/op
 gorm_prep:     0.33s       332709 ns/op    5113 B/op     59 allocs/op
       pop:     0.35s       350921 ns/op    6051 B/op    186 allocs/op
        pg:     0.37s       370911 ns/op     768 B/op      9 allocs/op
       dbr:     0.37s       371546 ns/op    2651 B/op     57 allocs/op
       bun:     0.38s       382962 ns/op    4728 B/op      5 allocs/op
      gorm:     0.40s       399944 ns/op    6600 B/op     81 allocs/op
      xorm:     0.46s       456299 ns/op    3944 B/op    132 allocs/op
      zorm:     0.50s       502456 ns/op    3024 B/op     59 allocs/op
       rel:     0.52s       520359 ns/op    3024 B/op     44 allocs/op
      godb:     0.54s       536424 ns/op    5160 B/op    154 allocs/op
     upper:     1.16s      1160302 ns/op   16662 B/op    390 allocs/op

  1000 times - Read
       pgx:     0.13s       131530 ns/op     889 B/op      8 allocs/op
  pgx_pool:     0.13s       132253 ns/op    1076 B/op      9 allocs/op
       raw:     0.14s       136018 ns/op    2057 B/op     50 allocs/op
      sqlc:     0.15s       147731 ns/op    2176 B/op     52 allocs/op
     beego:     0.15s       150396 ns/op    2089 B/op     75 allocs/op
    reform:     0.15s       153381 ns/op    3193 B/op     86 allocs/op
 gorm_prep:     0.15s       154496 ns/op    4821 B/op     85 allocs/op
      gorp:     0.16s       159532 ns/op    3881 B/op    194 allocs/op
       pop:     0.17s       165075 ns/op    3254 B/op     67 allocs/op
        pg:     0.17s       168266 ns/op     872 B/op     20 allocs/op
       ent:     0.17s       168968 ns/op    5448 B/op    150 allocs/op
 sqlboiler:     0.17s       170175 ns/op     971 B/op     14 allocs/op
       bun:     0.18s       178298 ns/op    5522 B/op     21 allocs/op
       rel:     0.18s       184154 ns/op    2312 B/op     49 allocs/op
       dbr:     0.19s       185289 ns/op    2184 B/op     37 allocs/op
      zorm:     0.19s       191733 ns/op    3072 B/op     67 allocs/op
      gorm:     0.21s       214178 ns/op    5157 B/op     96 allocs/op
      sqlx:     0.29s       292301 ns/op    1976 B/op     43 allocs/op
      godb:     0.32s       319634 ns/op    4112 B/op    102 allocs/op
     upper:     0.32s       320774 ns/op    5055 B/op    110 allocs/op
      xorm:     0.33s       331563 ns/op    4648 B/op    127 allocs/op

  1000 times - MultiRead limit 100
    reform:     0.15s       146401 ns/op    2921 B/op     74 allocs/op
       pgx:     0.27s       265769 ns/op   42949 B/op    504 allocs/op
  pgx_pool:     0.27s       268003 ns/op   43009 B/op    504 allocs/op
     upper:     0.32s       324291 ns/op    4783 B/op     89 allocs/op
       raw:     0.36s       357824 ns/op   38342 B/op   1038 allocs/op
        pg:     0.39s       392868 ns/op   22260 B/op    629 allocs/op
      sqlc:     0.39s       394635 ns/op   73158 B/op   1251 allocs/op
      gorp:     0.43s       428793 ns/op   57370 B/op   1494 allocs/op
       ent:     0.45s       449129 ns/op   76097 B/op   2040 allocs/op
      sqlx:     0.46s       455456 ns/op   39064 B/op   1516 allocs/op
       pop:     0.46s       456897 ns/op   68371 B/op   1306 allocs/op
     beego:     0.50s       495004 ns/op   55195 B/op   3077 allocs/op
       dbr:     0.50s       503114 ns/op   32416 B/op   1545 allocs/op
       bun:     0.52s       522219 ns/op   34032 B/op   1124 allocs/op
 sqlboiler:     0.55s       547825 ns/op   66426 B/op   2259 allocs/op
 gorm_prep:     0.58s       575395 ns/op   56062 B/op   2079 allocs/op
      gorm:     0.65s       650946 ns/op   57186 B/op   2379 allocs/op
      zorm:     0.75s       748574 ns/op  163233 B/op   3241 allocs/op
       rel:     0.78s       778741 ns/op  100616 B/op   2252 allocs/op
      godb:     0.80s       796787 ns/op   75271 B/op   3084 allocs/op
      xorm:     0.85s       847418 ns/op  120183 B/op   4692 allocs/op
```

- orm-benchmark -multi=10

```
  
Reports:

  2000 times - Insert
       pgx:     0.59s       294845 ns/op     269 B/op     10 allocs/op
       raw:     0.61s       304968 ns/op     734 B/op     13 allocs/op
       ent:     0.62s       312205 ns/op    4119 B/op     98 allocs/op
 sqlboiler:     0.63s       315380 ns/op    1592 B/op     35 allocs/op
  pgx_pool:     0.63s       316589 ns/op     285 B/op     10 allocs/op
     beego:     0.64s       318933 ns/op    2375 B/op     56 allocs/op
      gorp:     0.65s       327332 ns/op    1798 B/op     42 allocs/op
      sqlc:     0.66s       327739 ns/op    2900 B/op     63 allocs/op
    reform:     0.67s       333558 ns/op    1807 B/op     51 allocs/op
 gorm_prep:     0.67s       334432 ns/op    5285 B/op     68 allocs/op
       dbr:     0.73s       365672 ns/op    2699 B/op     65 allocs/op
        pg:     0.74s       370062 ns/op    1342 B/op     10 allocs/op
      gorm:     0.75s       376983 ns/op    6937 B/op     90 allocs/op
       bun:     0.79s       394346 ns/op    5022 B/op     14 allocs/op
      sqlx:     0.92s       462284 ns/op     856 B/op     19 allocs/op
       rel:     1.01s       503972 ns/op    2584 B/op     44 allocs/op
      xorm:     1.03s       515197 ns/op    3332 B/op     89 allocs/op
      godb:     1.04s       522165 ns/op    4564 B/op    116 allocs/op
     upper:     1.08s       538942 ns/op    5913 B/op    126 allocs/op
      zorm:     1.08s       541540 ns/op    3785 B/op     77 allocs/op
       pop:     1.45s       724198 ns/op    9556 B/op    237 allocs/op

  2000 times - MultiInsert 100 row
  pgx_pool:     2.17s      1086049 ns/op  112936 B/op     43 allocs/op
       pgx:     2.21s      1103726 ns/op  112947 B/op     43 allocs/op
       raw:     2.40s      1200569 ns/op  183763 B/op    929 allocs/op
     beego:     2.58s      1289266 ns/op  177567 B/op   2745 allocs/op
    reform:     2.99s      1496796 ns/op  458754 B/op   2745 allocs/op
 gorm_prep:     3.01s      1503772 ns/op  227696 B/op   2281 allocs/op
        pg:     3.22s      1610897 ns/op    4400 B/op    112 allocs/op
       bun:     3.41s      1705318 ns/op   42437 B/op    218 allocs/op
       ent:     3.54s      1767982 ns/op  387370 B/op   4598 allocs/op
      sqlx:     3.80s      1899187 ns/op  170355 B/op   1551 allocs/op
      gorm:     3.86s      1928858 ns/op  271584 B/op   3829 allocs/op
      zorm:     4.44s      2222364 ns/op  199928 B/op   2780 allocs/op
      xorm:     4.63s      2313168 ns/op  248050 B/op   5414 allocs/op
       rel:     4.71s      2355820 ns/op  306871 B/op   3263 allocs/op
      godb:     4.89s      2443281 ns/op  255251 B/op   5894 allocs/op
     upper:     5.07s      2536485 ns/op  322916 B/op   4205 allocs/op
      sqlc:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert

  2000 times - Update
       raw:     0.27s       133846 ns/op     748 B/op     13 allocs/op
       ent:     0.30s       147663 ns/op    4608 B/op     98 allocs/op
      sqlx:     0.60s       300395 ns/op     872 B/op     20 allocs/op
       pgx:     0.60s       302329 ns/op     268 B/op     10 allocs/op
      sqlc:     0.61s       307427 ns/op     876 B/op     14 allocs/op
      gorp:     0.64s       318402 ns/op    1205 B/op     32 allocs/op
 sqlboiler:     0.65s       323305 ns/op     901 B/op     17 allocs/op
  pgx_pool:     0.65s       326454 ns/op     285 B/op     10 allocs/op
    reform:     0.69s       342556 ns/op    1773 B/op     51 allocs/op
 gorm_prep:     0.70s       349369 ns/op    5121 B/op     59 allocs/op
     beego:     0.70s       352293 ns/op    1752 B/op     47 allocs/op
       pop:     0.74s       369221 ns/op    6047 B/op    186 allocs/op
       dbr:     0.74s       370302 ns/op    2651 B/op     57 allocs/op
        pg:     0.75s       373958 ns/op     768 B/op      9 allocs/op
       bun:     0.75s       376480 ns/op    4728 B/op      5 allocs/op
      gorm:     0.81s       405756 ns/op    6604 B/op     81 allocs/op
      xorm:     1.05s       526504 ns/op    3944 B/op    132 allocs/op
       rel:     1.06s       529309 ns/op    3024 B/op     44 allocs/op
      godb:     1.06s       529866 ns/op    5161 B/op    154 allocs/op
      zorm:     1.10s       548834 ns/op    3024 B/op     59 allocs/op
     upper:     2.39s      1194973 ns/op   16625 B/op    390 allocs/op

  2000 times - Read
       pgx:     0.24s       121297 ns/op     892 B/op      8 allocs/op
  pgx_pool:     0.27s       133236 ns/op    1080 B/op      9 allocs/op
      sqlc:     0.29s       142770 ns/op    2172 B/op     52 allocs/op
       raw:     0.29s       143185 ns/op    2065 B/op     50 allocs/op
    reform:     0.30s       150515 ns/op    3197 B/op     86 allocs/op
     beego:     0.31s       152586 ns/op    2088 B/op     75 allocs/op
       pop:     0.32s       157611 ns/op    3185 B/op     67 allocs/op
      gorp:     0.32s       158568 ns/op    3876 B/op    194 allocs/op
        pg:     0.32s       160615 ns/op     872 B/op     20 allocs/op
 gorm_prep:     0.33s       164809 ns/op    4816 B/op     85 allocs/op
       ent:     0.34s       168007 ns/op    5452 B/op    150 allocs/op
       bun:     0.35s       173878 ns/op    5522 B/op     21 allocs/op
 sqlboiler:     0.35s       173944 ns/op     956 B/op     14 allocs/op
       dbr:     0.35s       175417 ns/op    2184 B/op     37 allocs/op
       rel:     0.36s       180877 ns/op    2312 B/op     49 allocs/op
      zorm:     0.39s       194053 ns/op    3072 B/op     67 allocs/op
      gorm:     0.44s       218435 ns/op    5151 B/op     96 allocs/op
      sqlx:     0.63s       315316 ns/op    1976 B/op     43 allocs/op
      godb:     0.64s       320326 ns/op    4112 B/op    102 allocs/op
      xorm:     0.69s       343132 ns/op    4648 B/op    127 allocs/op
     upper:     0.69s       347203 ns/op    5054 B/op    110 allocs/op

  2000 times - MultiRead limit 100
    reform:     0.30s       147553 ns/op    2924 B/op     74 allocs/op
       pgx:     0.55s       276551 ns/op   42949 B/op    504 allocs/op
  pgx_pool:     0.57s       283223 ns/op   43008 B/op    504 allocs/op
     upper:     0.65s       323937 ns/op    4776 B/op     89 allocs/op
       raw:     0.70s       348480 ns/op   38342 B/op   1038 allocs/op
      sqlc:     0.77s       385293 ns/op   73158 B/op   1251 allocs/op
        pg:     0.80s       400522 ns/op   23382 B/op    629 allocs/op
      gorp:     0.86s       427968 ns/op   57374 B/op   1494 allocs/op
      sqlx:     0.87s       435495 ns/op   39063 B/op   1516 allocs/op
       ent:     0.93s       462604 ns/op   76093 B/op   2040 allocs/op
       pop:     0.93s       466650 ns/op   68463 B/op   1306 allocs/op
       dbr:     0.97s       485156 ns/op   32416 B/op   1545 allocs/op
     beego:     1.02s       507837 ns/op   55183 B/op   3077 allocs/op
       bun:     1.03s       516111 ns/op   34026 B/op   1124 allocs/op
 sqlboiler:     1.11s       553283 ns/op   66575 B/op   2260 allocs/op
 gorm_prep:     1.16s       578261 ns/op   56070 B/op   2079 allocs/op
      gorm:     1.28s       639399 ns/op   57171 B/op   2379 allocs/op
      zorm:     1.51s       753981 ns/op  163233 B/op   3241 allocs/op
       rel:     1.57s       784525 ns/op  100616 B/op   2252 allocs/op
      godb:     1.60s       802285 ns/op   75269 B/op   3084 allocs/op
      xorm:     1.70s       850227 ns/op  120179 B/op   4692 allocs/op
```

- orm-benchmark -multi=20

```
  
Reports:

  4000 times - Insert
  pgx_pool:     1.17s       291718 ns/op     286 B/op     10 allocs/op
       pgx:     1.19s       298020 ns/op     270 B/op     10 allocs/op
       raw:     1.21s       303208 ns/op     719 B/op     13 allocs/op
      gorp:     1.24s       309374 ns/op    1782 B/op     42 allocs/op
    reform:     1.28s       320532 ns/op    1790 B/op     51 allocs/op
     beego:     1.30s       325515 ns/op    2375 B/op     56 allocs/op
       bun:     1.34s       335418 ns/op    5008 B/op     14 allocs/op
       ent:     1.35s       337201 ns/op    4120 B/op     98 allocs/op
       dbr:     1.36s       340398 ns/op    2693 B/op     65 allocs/op
 gorm_prep:     1.38s       344045 ns/op    5275 B/op     68 allocs/op
      sqlc:     1.40s       349658 ns/op    2884 B/op     63 allocs/op
 sqlboiler:     1.41s       352320 ns/op    1574 B/op     35 allocs/op
        pg:     1.43s       357487 ns/op    1068 B/op     10 allocs/op
      gorm:     1.54s       385407 ns/op    6938 B/op     90 allocs/op
      sqlx:     1.89s       473667 ns/op     856 B/op     19 allocs/op
      godb:     1.95s       487965 ns/op    4558 B/op    116 allocs/op
      xorm:     1.95s       488128 ns/op    3326 B/op     89 allocs/op
     upper:     1.96s       490927 ns/op    5896 B/op    126 allocs/op
       rel:     1.99s       498455 ns/op    2584 B/op     44 allocs/op
      zorm:     2.03s       506506 ns/op    3784 B/op     77 allocs/op
       pop:     2.92s       729156 ns/op    9537 B/op    237 allocs/op

  4000 times - MultiInsert 100 row
       pgx:     4.09s      1023365 ns/op  112924 B/op     43 allocs/op
  pgx_pool:     4.17s      1042489 ns/op  112933 B/op     43 allocs/op
       raw:     4.57s      1143629 ns/op  183706 B/op    929 allocs/op
     beego:     4.78s      1195512 ns/op  177562 B/op   2745 allocs/op
    reform:     5.60s      1400888 ns/op  458748 B/op   2745 allocs/op
 gorm_prep:     5.68s      1419652 ns/op  227698 B/op   2281 allocs/op
        pg:     6.17s      1541495 ns/op    3857 B/op    112 allocs/op
       bun:     6.47s      1616455 ns/op   42420 B/op    219 allocs/op
       ent:     6.74s      1685671 ns/op  387357 B/op   4598 allocs/op
      gorm:     7.63s      1906353 ns/op  271632 B/op   3829 allocs/op
      sqlx:     7.63s      1906565 ns/op  170777 B/op   1551 allocs/op
      zorm:     8.93s      2232039 ns/op  199928 B/op   2780 allocs/op
       rel:     9.13s      2283477 ns/op  306872 B/op   3263 allocs/op
      xorm:     9.17s      2291943 ns/op  248120 B/op   5414 allocs/op
      godb:     9.36s      2339300 ns/op  255246 B/op   5894 allocs/op
     upper:    10.27s      2566906 ns/op  322904 B/op   4205 allocs/op
       pop:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert

  4000 times - Update
       ent:     0.63s       157741 ns/op    4606 B/op     98 allocs/op
       raw:     0.63s       158070 ns/op     750 B/op     13 allocs/op
    reform:     1.14s       285358 ns/op    1774 B/op     51 allocs/op
      sqlx:     1.20s       299426 ns/op     872 B/op     20 allocs/op
  pgx_pool:     1.20s       299780 ns/op     286 B/op     10 allocs/op
      gorp:     1.20s       300619 ns/op    1204 B/op     32 allocs/op
       pgx:     1.22s       305685 ns/op     270 B/op     10 allocs/op
     beego:     1.27s       318517 ns/op    1752 B/op     47 allocs/op
      sqlc:     1.32s       329260 ns/op     878 B/op     14 allocs/op
       dbr:     1.32s       330021 ns/op    2651 B/op     57 allocs/op
 gorm_prep:     1.39s       348029 ns/op    5116 B/op     59 allocs/op
 sqlboiler:     1.40s       349687 ns/op     901 B/op     17 allocs/op
       bun:     1.40s       350046 ns/op    4728 B/op      5 allocs/op
       pop:     1.42s       354990 ns/op    6046 B/op    186 allocs/op
        pg:     1.47s       368462 ns/op     768 B/op      9 allocs/op
      gorm:     1.58s       395061 ns/op    6604 B/op     81 allocs/op
      zorm:     2.06s       514354 ns/op    3024 B/op     59 allocs/op
      xorm:     2.07s       516497 ns/op    3944 B/op    132 allocs/op
       rel:     2.10s       525817 ns/op    3024 B/op     44 allocs/op
      godb:     2.11s       526529 ns/op    5160 B/op    154 allocs/op
     upper:     4.56s      1140248 ns/op   16619 B/op    390 allocs/op

  4000 times - Read
       pgx:     0.52s       129814 ns/op     894 B/op      8 allocs/op
  pgx_pool:     0.53s       132272 ns/op    1079 B/op      9 allocs/op
       raw:     0.57s       141526 ns/op    2062 B/op     50 allocs/op
      sqlc:     0.58s       144321 ns/op    2172 B/op     52 allocs/op
     beego:     0.61s       151680 ns/op    2088 B/op     75 allocs/op
    reform:     0.63s       156424 ns/op    3199 B/op     86 allocs/op
       pop:     0.63s       156790 ns/op    3172 B/op     67 allocs/op
      gorp:     0.64s       159179 ns/op    3878 B/op    194 allocs/op
        pg:     0.65s       161486 ns/op     872 B/op     20 allocs/op
 gorm_prep:     0.66s       165907 ns/op    4818 B/op     85 allocs/op
 sqlboiler:     0.68s       169720 ns/op     956 B/op     14 allocs/op
       rel:     0.68s       170934 ns/op    2312 B/op     49 allocs/op
       ent:     0.70s       174638 ns/op    5452 B/op    150 allocs/op
       bun:     0.72s       178943 ns/op    5521 B/op     21 allocs/op
       dbr:     0.72s       180002 ns/op    2184 B/op     37 allocs/op
      zorm:     0.80s       198875 ns/op    3072 B/op     67 allocs/op
      gorm:     0.85s       213233 ns/op    5149 B/op     96 allocs/op
      sqlx:     1.25s       312210 ns/op    1976 B/op     43 allocs/op
      godb:     1.32s       328903 ns/op    4112 B/op    102 allocs/op
      xorm:     1.35s       336683 ns/op    4648 B/op    127 allocs/op
     upper:     1.36s       340018 ns/op    5055 B/op    110 allocs/op

  4000 times - MultiRead limit 100
    reform:     0.58s       144887 ns/op    2926 B/op     74 allocs/op
  pgx_pool:     1.11s       278261 ns/op   43007 B/op    504 allocs/op
       pgx:     1.12s       279273 ns/op   42949 B/op    504 allocs/op
     upper:     1.32s       330749 ns/op    4777 B/op     89 allocs/op
       raw:     1.37s       343515 ns/op   38341 B/op   1038 allocs/op
        pg:     1.57s       392642 ns/op   23064 B/op    629 allocs/op
      sqlc:     1.59s       397631 ns/op   73158 B/op   1251 allocs/op
      gorp:     1.72s       431230 ns/op   57373 B/op   1494 allocs/op
      sqlx:     1.80s       449611 ns/op   39064 B/op   1516 allocs/op
       pop:     1.87s       467179 ns/op   68359 B/op   1306 allocs/op
       ent:     1.89s       472013 ns/op   76095 B/op   2040 allocs/op
       dbr:     1.95s       487923 ns/op   32416 B/op   1545 allocs/op
     beego:     2.01s       501270 ns/op   55186 B/op   3077 allocs/op
       bun:     2.07s       517973 ns/op   34027 B/op   1124 allocs/op
 sqlboiler:     2.20s       549603 ns/op   66353 B/op   2259 allocs/op
 gorm_prep:     2.32s       580567 ns/op   56073 B/op   2079 allocs/op
      gorm:     2.61s       651632 ns/op   57171 B/op   2379 allocs/op
      zorm:     3.10s       774994 ns/op  163233 B/op   3241 allocs/op
       rel:     3.11s       778513 ns/op  100616 B/op   2252 allocs/op
      godb:     3.18s       795995 ns/op   75268 B/op   3084 allocs/op
      xorm:     3.48s       870889 ns/op  120184 B/op   4692 allocs/op
```
