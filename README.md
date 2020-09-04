# ORM Benchmark

### Environment

* go version go1.6 linux/amd64

### PostgreSQL

* PostgreSQL 9.5 for Linux on x86_64

### ORMs

All package run in no-cache mode.

* [dbr](https://github.com/gocraft/dbr) (in preparation)
* [genmai](https://github.com/naoina/genmai) (in preparation)
* [gorm](https://github.com/jinzhu/gorm)
* [gorp](https://github.com/go-gorp/gorp) (in preparation)
* [pg](https://github.com/go-pg/pg)
* [beego/orm](https://github.com/astaxie/beego/tree/master/orm)
* [sqlx](https://github.com/jmoiron/sqlx) (in preparation)
* [xorm](https://github.com/xormplus/xorm)
	
### Run

```go
go get github.com/milkpod29/orm-benchmark
# build
go install
# all
orm-benchmark -multi=20 -orm=all
# portion
orm-benchmark -multi=20 -orm=xorm
```

### Reports

左から

* ORM名
* 総実行時間（少ないほど良い）
* １回の実行にかかった時間（少ないほど良い）
* 実行ごとに割り当てられたメモリのサイズ（少ないほど良い）
* １回の実行でメモリアロケーション（メモリ割り当て/配分）が行われた回数（少ないほど良い）

#### First time

```
 40000 times - Insert
        pg:    28.11s       702699 ns/op     680 B/op     17 allocs/op
       raw:    28.19s       704850 ns/op     568 B/op     25 allocs/op
      xorm:    52.28s      1307026 ns/op    4050 B/op    116 allocs/op
 beego_orm:    63.99s      1599813 ns/op    2713 B/op     78 allocs/op
      gorm:    78.97s      1974368 ns/op    8189 B/op    189 allocs/op

 10000 times - MultiInsert 100 row
        pg:    27.27s      2726560 ns/op   15164 B/op    317 allocs/op
 beego_orm:    33.12s      3312423 ns/op  176680 B/op   3264 allocs/op
       raw:    33.84s      3383613 ns/op  115808 B/op   1529 allocs/op
      xorm:    50.42s      5042414 ns/op 2373289 B/op   8095 allocs/op
      gorm:     Not support multi insert

 40000 times - Update
       raw:    16.10s       402576 ns/op     568 B/op     28 allocs/op
        pg:    33.24s       830982 ns/op     608 B/op     11 allocs/op
      xorm:    58.68s      1466943 ns/op    4177 B/op    138 allocs/op
 beego_orm:    68.03s      1700812 ns/op    2281 B/op     71 allocs/op
      gorm:    98.61s      2465242 ns/op    8253 B/op    209 allocs/op

 80000 times - Read
       raw:    41.77s       522184 ns/op     896 B/op     29 allocs/op
      gorm:    59.69s       746157 ns/op    8195 B/op    194 allocs/op
      xorm:    67.30s       841274 ns/op    7747 B/op    240 allocs/op
        pg:    98.98s      1237223 ns/op     960 B/op     38 allocs/op
 beego_orm:    151.66s      1895795 ns/op    3081 B/op    108 allocs/op

 40000 times - MultiRead limit 100
       raw:    35.98s       899513 ns/op   36561 B/op   1509 allocs/op
        pg:    38.77s       969259 ns/op   25393 B/op    942 allocs/op
      xorm:    55.26s      1381464 ns/op    3169 B/op     80 allocs/op
 beego_orm:    75.09s      1877161 ns/op   91153 B/op   4601 allocs/op
      gorm:    197.95s      4948706 ns/op  628190 B/op  17018 allocs/op
```
 
#### Second time
 
```
 40000 times - Insert
       raw:    26.54s       663378 ns/op     568 B/op     25 allocs/op
        pg:    27.96s       699040 ns/op     745 B/op     17 allocs/op
      xorm:    52.40s      1310069 ns/op    4050 B/op    116 allocs/op
 beego_orm:    64.94s      1623420 ns/op    2713 B/op     78 allocs/op
      gorm:    96.02s      2400406 ns/op    8124 B/op    189 allocs/op

 10000 times - MultiInsert 100 row
        pg:    25.57s      2556601 ns/op   15164 B/op    317 allocs/op
       raw:    26.32s      2632441 ns/op  115808 B/op   1529 allocs/op
 beego_orm:    30.97s      3096774 ns/op  176681 B/op   3264 allocs/op
      xorm:    47.30s      4730103 ns/op 2373300 B/op   8095 allocs/op
      gorm:     Not support multi insert

 40000 times - Update
       raw:    24.04s       601121 ns/op     568 B/op     28 allocs/op
        pg:    26.71s       667631 ns/op     608 B/op     11 allocs/op
      xorm:    45.58s      1139446 ns/op    4177 B/op    138 allocs/op
 beego_orm:    48.22s      1205619 ns/op    2281 B/op     71 allocs/op
      gorm:    96.84s      2420984 ns/op    8252 B/op    209 allocs/op

 80000 times - Read
       raw:    38.05s       475586 ns/op     896 B/op     29 allocs/op
      gorm:    61.08s       763527 ns/op    8195 B/op    194 allocs/op
      xorm:    62.42s       780222 ns/op    7748 B/op    240 allocs/op
        pg:    81.14s      1014239 ns/op     960 B/op     38 allocs/op
 beego_orm:    97.41s      1217656 ns/op    3081 B/op    108 allocs/op

 40000 times - MultiRead limit 100
       raw:    31.86s       796471 ns/op   36561 B/op   1509 allocs/op
        pg:    33.30s       832391 ns/op   25393 B/op    942 allocs/op
      xorm:    57.09s      1427295 ns/op    3169 B/op     80 allocs/op
 beego_orm:    63.73s      1593203 ns/op   91154 B/op   4601 allocs/op
      gorm:    209.57s      5239374 ns/op  628185 B/op  17018 allocs/op
```

### Another environment.

* Mac OS X Yosemite 10.10.5
* 2.5 GHz Intel Core i5
* 16GB 1600MHz DDR3

#### First time

```
 40000 times - Insert
       raw:    31.12s       778024 ns/op     568 B/op     25 allocs/op
        pg:    35.09s       877176 ns/op     746 B/op     17 allocs/op
      xorm:    47.64s      1190920 ns/op    4050 B/op    116 allocs/op
 beego_orm:    50.80s      1270017 ns/op    2713 B/op     78 allocs/op
      gorm:    84.36s      2108979 ns/op    8125 B/op    189 allocs/op

 10000 times - MultiInsert 100 row
        pg:    41.00s      4099767 ns/op   15164 B/op    317 allocs/op
       raw:    43.40s      4339534 ns/op  115809 B/op   1529 allocs/op
 beego_orm:    50.50s      5049734 ns/op  176684 B/op   3264 allocs/op
      xorm:    65.46s      6545607 ns/op 2373346 B/op   8095 allocs/op
      gorm:     Not support multi insert

 40000 times - Update
       raw:    11.56s       289081 ns/op     568 B/op     28 allocs/op
        pg:    48.30s      1207478 ns/op     608 B/op     11 allocs/op
      xorm:    49.56s      1238995 ns/op    4178 B/op    138 allocs/op
 beego_orm:    66.21s      1655274 ns/op    2281 B/op     71 allocs/op
      gorm:    84.11s      2102763 ns/op    8253 B/op    209 allocs/op

 80000 times - Read
       raw:    25.65s       320640 ns/op     896 B/op     29 allocs/op
        pg:    30.32s       378952 ns/op     960 B/op     38 allocs/op
 beego_orm:    65.18s       814704 ns/op    3081 B/op    108 allocs/op
      xorm:    75.87s       948316 ns/op    7813 B/op    241 allocs/op
      gorm:    75.89s       948676 ns/op    8196 B/op    194 allocs/op

 40000 times - MultiRead limit 100
      xorm:    32.75s       818868 ns/op    3265 B/op     82 allocs/op
       raw:    55.66s      1391485 ns/op   36561 B/op   1509 allocs/op
        pg:    64.12s      1603056 ns/op   25394 B/op    942 allocs/op
 beego_orm:    85.69s      2142218 ns/op   91159 B/op   4601 allocs/op
      gorm:    262.57s      6564279 ns/op  628239 B/op  17019 allocs/op
```

#### Second time

```
 40000 times - Insert
       raw:    28.57s       714132 ns/op     568 B/op     25 allocs/op
        pg:    34.95s       873690 ns/op     680 B/op     17 allocs/op
      xorm:    45.54s      1138420 ns/op    4050 B/op    116 allocs/op
 beego_orm:    60.91s      1522682 ns/op    2713 B/op     78 allocs/op
      gorm:    77.55s      1938706 ns/op    8190 B/op    189 allocs/op

      gorm:     4.54s      2271262 ns/op    5561 B/op     87 allocs/op

 10000 times - MultiInsert 100 row
       raw:    40.89s      4088782 ns/op  115809 B/op   1529 allocs/op
        pg:    42.50s      4249919 ns/op   15164 B/op    317 allocs/op
 beego_orm:    48.39s      4839005 ns/op  176683 B/op   3264 allocs/op
      xorm:    64.71s      6470947 ns/op 2373252 B/op   8094 allocs/op
      gorm:     Not support multi insert

      gorm:     9.25s      4627112 ns/op  208763 B/op   3571 allocs/op

 40000 times - Update
       raw:    10.63s       265667 ns/op     568 B/op     28 allocs/op
        pg:    51.84s      1295906 ns/op     608 B/op     11 allocs/op
      xorm:    64.63s      1615842 ns/op    4177 B/op    138 allocs/op
 beego_orm:    64.95s      1623816 ns/op    2281 B/op     71 allocs/op
      gorm:    86.03s      2150721 ns/op    8253 B/op    209 allocs/op

 80000 times - Read
       raw:    22.54s       281776 ns/op     896 B/op     29 allocs/op
        pg:    37.49s       468683 ns/op     960 B/op     38 allocs/op
      xorm:    84.79s      1059842 ns/op    7812 B/op    241 allocs/op
 beego_orm:    86.27s      1078420 ns/op    3081 B/op    108 allocs/op
      gorm:    87.23s      1090387 ns/op    8196 B/op    194 allocs/op

      gorm:     2.95s       736650 ns/op    3776 B/op     58 allocs/op

 40000 times - MultiRead limit 100
      xorm:    38.86s       971585 ns/op    3265 B/op     82 allocs/op
       raw:    47.20s      1179997 ns/op   36561 B/op   1509 allocs/op
        pg:    68.46s      1711417 ns/op   25394 B/op    942 allocs/op
 beego_orm:    76.58s      1914539 ns/op   91157 B/op   4601 allocs/op
      gorm:    251.53s     6288155 ns/op  628245 B/op  17019 allocs/op

      gorm:     9.67s      2416851 ns/op   37757 B/op   1481 allocs/op
```

```2020-03-01
 10000 times - Insert
        pg:    14.05s      1404613 ns/op     637 B/op     11 allocs/op
       raw:    14.24s      1423653 ns/op     720 B/op     21 allocs/op
 beego_orm:    14.57s      1457007 ns/op    2440 B/op     58 allocs/op
      xorm:    19.86s      1985720 ns/op    3379 B/op    109 allocs/op
      gorm:    25.06s      2505715 ns/op    5534 B/op     87 allocs/op

 10000 times - MultiInsert 100 row
 beego_orm:    28.97s      2897276 ns/op  186786 B/op   2845 allocs/op
        pg:    34.97s      3496896 ns/op   14381 B/op    212 allocs/op
       raw:    40.28s      4028068 ns/op  140093 B/op   1421 allocs/op
      gorm:    42.49s      4249166 ns/op  208761 B/op   3571 allocs/op
      xorm:    47.07s      4707072 ns/op  303853 B/op   7795 allocs/op

 20000 times - Update
      gorm:     0.00s          175 ns/op     336 B/op      2 allocs/op
       raw:    15.47s       773537 ns/op     728 B/op     21 allocs/op
 beego_orm:    28.40s      1420158 ns/op    1824 B/op     49 allocs/op
        pg:    28.88s      1444039 ns/op     592 B/op      9 allocs/op
      xorm:    45.56s      2278068 ns/op    3409 B/op    135 allocs/op

 20000 times - Read
 beego_orm:    15.75s       787694 ns/op    2160 B/op     80 allocs/op
       raw:    15.99s       799425 ns/op     920 B/op     28 allocs/op
      gorm:    16.57s       828602 ns/op    3776 B/op     58 allocs/op
        pg:    17.12s       856036 ns/op     728 B/op     17 allocs/op
      xorm:    40.93s      2046390 ns/op    9114 B/op    265 allocs/op

 20000 times - MultiRead limit 100
       raw:    23.43s      1171517 ns/op   28656 B/op   1212 allocs/op
        pg:    25.75s      1287599 ns/op   24380 B/op    622 allocs/op
 beego_orm:    30.93s      1546322 ns/op   60872 B/op   3478 allocs/op
      xorm:    33.18s      1659021 ns/op    3232 B/op     91 allocs/op
      gorm:    44.18s      2209081 ns/op   37800 B/op   1482 allocs/op
```

```2020-03-09
 10000 times - Insert
        pg:    13.94s      1393880 ns/op     637 B/op     11 allocs/op
 beego_orm:    13.94s      1393982 ns/op    2440 B/op     58 allocs/op
       raw:    21.00s      2100118 ns/op     720 B/op     21 allocs/op
      gorm:    21.82s      2182228 ns/op    5455 B/op     78 allocs/op
      xorm:    22.60s      2260292 ns/op    3378 B/op    109 allocs/op

 10000 times - MultiInsert 100 row
 beego_orm:    28.87s      2887480 ns/op  186770 B/op   2845 allocs/op
        pg:    32.38s      3237629 ns/op    3883 B/op    111 allocs/op
       raw:    47.45s      4744620 ns/op  140093 B/op   1421 allocs/op
      gorm:    47.70s      4770326 ns/op  208681 B/op   3562 allocs/op
      xorm:    73.37s      7336828 ns/op  303868 B/op   7796 allocs/op

 20000 times - Update
       raw:    21.65s      1082684 ns/op     728 B/op     21 allocs/op
        pg:    28.64s      1431990 ns/op     592 B/op      9 allocs/op
 beego_orm:    33.33s      1666395 ns/op    1824 B/op     49 allocs/op
      gorm:    48.12s      2405972 ns/op    5424 B/op     92 allocs/op
      xorm:    58.31s      2915564 ns/op    3409 B/op    135 allocs/op

 20000 times - Read
        pg:    15.08s       754172 ns/op     728 B/op     17 allocs/op
 beego_orm:    15.19s       759446 ns/op    2160 B/op     80 allocs/op
       raw:    16.24s       812130 ns/op     920 B/op     28 allocs/op
      gorm:    19.49s       974361 ns/op    3976 B/op     63 allocs/op
      xorm:    47.02s      2351104 ns/op    9114 B/op    265 allocs/op

 20000 times - MultiRead limit 100
       raw:    25.62s      1281093 ns/op   28656 B/op   1212 allocs/op
        pg:    28.29s      1414309 ns/op   24380 B/op    622 allocs/op
 beego_orm:    29.39s      1469610 ns/op   60872 B/op   3478 allocs/op
      gorm:    44.04s      2202086 ns/op   38008 B/op   1488 allocs/op
      xorm:     doesn't work


#### sync Pool
Reports:

  2000 times - Insert
       raw:     2.63s      1315944 ns/op     720 B/op     21 allocs/op
 beego_orm:     2.84s      1418683 ns/op    2441 B/op     58 allocs/op
        pg:     2.90s      1448360 ns/op     661 B/op     11 allocs/op
      gorm:     4.24s      2118711 ns/op    3805 B/op     75 allocs/op
      xorm:     5.96s      2979356 ns/op    3387 B/op    109 allocs/op

  2000 times - MultiInsert 100 row
 beego_orm:     5.74s      2868270 ns/op  186777 B/op   2845 allocs/op
        pg:     6.86s      3430628 ns/op    3893 B/op    111 allocs/op
      gorm:     9.16s      4579216 ns/op  207038 B/op   3559 allocs/op
       raw:     9.58s      4790311 ns/op  140176 B/op   1421 allocs/op
      xorm:    11.34s      5668071 ns/op  303849 B/op   7795 allocs/op

  4000 times - Update
       raw:     3.09s       772244 ns/op     728 B/op     21 allocs/op
 beego_orm:     5.40s      1349573 ns/op    1824 B/op     49 allocs/op
        pg:     6.38s      1593807 ns/op     592 B/op      9 allocs/op
      xorm:     9.18s      2295833 ns/op    3409 B/op    135 allocs/op
      gorm:     9.47s      2366993 ns/op    3744 B/op     89 allocs/op

  4000 times - Read
 beego_orm:     2.87s       717837 ns/op    2160 B/op     80 allocs/op
       raw:     3.14s       784612 ns/op     920 B/op     28 allocs/op
        pg:     3.56s       889688 ns/op     728 B/op     17 allocs/op
      gorm:     3.87s       967444 ns/op    2298 B/op     60 allocs/op
      xorm:     6.90s      1724661 ns/op    9115 B/op    265 allocs/op

  4000 times - MultiRead limit 100
        pg:     5.49s      1371451 ns/op   24380 B/op    622 allocs/op
       raw:     5.75s      1436428 ns/op   28656 B/op   1212 allocs/op
 beego_orm:     5.94s      1483969 ns/op   60873 B/op   3478 allocs/op
      gorm:     9.19s      2297250 ns/op   36334 B/op   1485 allocs/op
      xorm:     doesn't work

### strigns.Builder

gorm.DB

  1000 times - Insert
       raw:     1.36s      1357432 ns/op     720 B/op     21 allocs/op
        pg:     1.39s      1386543 ns/op     690 B/op     11 allocs/op
 beego_orm:     1.42s      1420223 ns/op    2442 B/op     58 allocs/op
      xorm:     2.20s      2197230 ns/op    3397 B/op    109 allocs/op
      gorm:     2.27s      2267578 ns/op    3696 B/op     61 allocs/op

  1000 times - MultiInsert 100 row
 beego_orm:     2.61s      2610007 ns/op  186825 B/op   2845 allocs/op
        pg:     3.38s      3384084 ns/op    3905 B/op    111 allocs/op
       raw:     4.11s      4111140 ns/op  140280 B/op   1422 allocs/op
      xorm:     4.78s      4783824 ns/op  303857 B/op   7795 allocs/op
      gorm:     5.73s      5728429 ns/op  190556 B/op   2371 allocs/op

  2000 times - Update
       raw:     1.52s       757735 ns/op     728 B/op     21 allocs/op
        pg:     2.86s      1432344 ns/op     592 B/op      9 allocs/op
 beego_orm:     2.92s      1458986 ns/op    1825 B/op     49 allocs/op
      xorm:     4.46s      2228364 ns/op    3409 B/op    135 allocs/op
      gorm:     5.34s      2671460 ns/op    3488 B/op     66 allocs/op

  2000 times - Read
       raw:     1.51s       753105 ns/op     920 B/op     28 allocs/op
 beego_orm:     1.88s       941952 ns/op    2160 B/op     80 allocs/op
        pg:     1.90s       949536 ns/op     728 B/op     17 allocs/op
      gorm:     2.08s      1042103 ns/op    2232 B/op     56 allocs/op
      xorm:     3.67s      1836125 ns/op    9114 B/op    265 allocs/op

  2000 times - MultiRead limit 100
       raw:     2.45s      1226566 ns/op   28656 B/op   1212 allocs/op
 beego_orm:     3.69s      1847376 ns/op   60871 B/op   3478 allocs/op
        pg:     3.99s      1993293 ns/op   24380 B/op    622 allocs/op
      gorm:     4.68s      2339868 ns/op   36211 B/op   1475 allocs/op
      xorm:     doesn't work

*gorm.DB
  2000 times - Insert
       raw:     2.62s      1309265 ns/op     720 B/op     21 allocs/op
        pg:     2.86s      1428407 ns/op     661 B/op     11 allocs/op
 beego_orm:     2.95s      1475661 ns/op    2440 B/op     58 allocs/op
      xorm:     4.58s      2289102 ns/op    3387 B/op    109 allocs/op
      gorm:     4.98s      2488317 ns/op    3660 B/op     61 allocs/op

  2000 times - MultiInsert 100 row
 beego_orm:     5.90s      2948159 ns/op  186790 B/op   2845 allocs/op
        pg:     7.02s      3507731 ns/op    3893 B/op    111 allocs/op
      gorm:     8.07s      4034748 ns/op  190558 B/op   2371 allocs/op
       raw:     8.85s      4424962 ns/op  140176 B/op   1421 allocs/op
      xorm:    11.00s      5499982 ns/op  303842 B/op   7795 allocs/op

  4000 times - Update
       raw:     2.85s       712253 ns/op     728 B/op     21 allocs/op
 beego_orm:     5.43s      1357020 ns/op    1824 B/op     49 allocs/op
        pg:     5.93s      1481451 ns/op     592 B/op      9 allocs/op
      xorm:     9.21s      2301517 ns/op    3409 B/op    135 allocs/op
      gorm:     9.42s      2354277 ns/op    3488 B/op     66 allocs/op

  4000 times - Read
       raw:     3.04s       760552 ns/op     920 B/op     28 allocs/op
 beego_orm:     3.24s       809453 ns/op    2160 B/op     80 allocs/op
        pg:     4.14s      1035779 ns/op     728 B/op     17 allocs/op
      gorm:     4.66s      1166078 ns/op    2232 B/op     56 allocs/op
      xorm:     6.51s      1626400 ns/op    9114 B/op    265 allocs/op

  4000 times - MultiRead limit 100
       raw:     4.91s      1227053 ns/op   28656 B/op   1212 allocs/op
 beego_orm:     6.09s      1522783 ns/op   60871 B/op   3478 allocs/op
        pg:     7.56s      1889477 ns/op   24380 B/op    622 allocs/op
      gorm:    12.86s      3215044 ns/op   36206 B/op   1475 allocs/op
      xorm:     doesn't work

### Raw

  2000 times - Insert
  raw_stmt:     2.69s      1346612 ns/op     720 B/op     21 allocs/op
 beego_orm:     3.01s      1504618 ns/op    2441 B/op     58 allocs/op
        pg:     3.60s      1801486 ns/op     661 B/op     11 allocs/op
       raw:     4.61s      2305232 ns/op     903 B/op     22 allocs/op
      xorm:     5.07s      2536236 ns/op    3387 B/op    109 allocs/op
      gorm:     5.07s      2536595 ns/op    3665 B/op     61 allocs/op

  2000 times - MultiInsert 100 row
 beego_orm:     5.88s      2941263 ns/op  186769 B/op   2845 allocs/op
  raw_stmt:     8.09s      4044649 ns/op  140176 B/op   1421 allocs/op
       raw:     8.88s      4440730 ns/op  140176 B/op   1421 allocs/op
        pg:     8.97s      4485576 ns/op    3892 B/op    111 allocs/op
      xorm:     9.81s      4903324 ns/op  303860 B/op   7795 allocs/op
      gorm:    10.64s      5318804 ns/op  190556 B/op   2371 allocs/op

  4000 times - Update
  raw_stmt:     3.59s       896615 ns/op     728 B/op     21 allocs/op
 beego_orm:     6.02s      1504257 ns/op    1824 B/op     49 allocs/op
       raw:     6.41s      1601855 ns/op     904 B/op     22 allocs/op
        pg:     6.97s      1742098 ns/op     592 B/op      9 allocs/op
      gorm:     9.45s      2361780 ns/op    3488 B/op     66 allocs/op
      xorm:    10.31s      2576940 ns/op    3409 B/op    135 allocs/op

  4000 times - Read
  raw_stmt:     2.92s       730469 ns/op     920 B/op     28 allocs/op
 beego_orm:     3.09s       771647 ns/op    2160 B/op     80 allocs/op
        pg:     3.41s       853146 ns/op     728 B/op     17 allocs/op
      gorm:     4.04s      1009306 ns/op    2232 B/op     56 allocs/op
       raw:     6.61s      1653018 ns/op    1504 B/op     41 allocs/op
      xorm:     9.36s      2340635 ns/op    9114 B/op    265 allocs/op

  4000 times - MultiRead limit 100
        pg:     5.73s      1431535 ns/op   24380 B/op    622 allocs/op
 beego_orm:     5.99s      1498542 ns/op   60871 B/op   3478 allocs/op
  raw_stmt:     6.18s      1544481 ns/op   28656 B/op   1212 allocs/op
       raw:     6.28s      1569931 ns/op   30656 B/op   1513 allocs/op
      gorm:     9.38s      2345166 ns/op   36206 B/op   1475 allocs/op
      xorm:     doesn't work
```

```2020-03-10
 10000 times - Insert
  raw_stmt:    14.05s      1404809 ns/op     720 B/op     21 allocs/op
        pg:    14.10s      1409742 ns/op     637 B/op     11 allocs/op
 beego_orm:    14.31s      1431135 ns/op    2440 B/op     58 allocs/op
      gorm:    21.04s      2103939 ns/op    3631 B/op     61 allocs/op
      xorm:    22.24s      2224413 ns/op    3378 B/op    109 allocs/op
       raw:    23.78s      2377810 ns/op     897 B/op     22 allocs/op

 10000 times - MultiInsert 100 row
 beego_orm:    26.73s      2673442 ns/op  186776 B/op   2845 allocs/op
  raw_stmt:    29.71s      2970811 ns/op  125615 B/op   1417 allocs/op
        pg:    35.97s      3597264 ns/op    3883 B/op    111 allocs/op
      gorm:    42.29s      4228831 ns/op  190556 B/op   2371 allocs/op
       raw:    47.71s      4770942 ns/op  140093 B/op   1421 allocs/op
      xorm:    50.84s      5083784 ns/op  303866 B/op   7796 allocs/op

 20000 times - Update
  raw_stmt:    15.06s       752798 ns/op     728 B/op     21 allocs/op
 beego_orm:    27.10s      1354750 ns/op    1824 B/op     49 allocs/op
       raw:    32.74s      1636959 ns/op     904 B/op     22 allocs/op
      gorm:    43.17s      2158402 ns/op    3488 B/op     66 allocs/op
      xorm:    44.89s      2244256 ns/op    3408 B/op    135 allocs/op
        pg:    45.82s      2291038 ns/op     592 B/op      9 allocs/op

 20000 times - Read
 beego_orm:    13.23s       661715 ns/op    2160 B/op     80 allocs/op
  raw_stmt:    16.35s       817663 ns/op     920 B/op     28 allocs/op
        pg:    16.50s       825100 ns/op     728 B/op     17 allocs/op
      gorm:    16.60s       830109 ns/op    2232 B/op     56 allocs/op
       raw:    31.21s      1560652 ns/op    1504 B/op     41 allocs/op
      xorm:    34.44s      1721864 ns/op    9114 B/op    265 allocs/op

 20000 times - MultiRead limit 100
  raw_stmt:    23.58s      1179083 ns/op   28656 B/op   1212 allocs/op
       raw:    26.31s      1315371 ns/op   30656 B/op   1513 allocs/op
 beego_orm:    26.69s      1334363 ns/op   60873 B/op   3478 allocs/op
        pg:    38.04s      1902064 ns/op   24380 B/op    622 allocs/op
      gorm:    42.45s      2122355 ns/op   36206 B/op   1475 allocs/op
      xorm:     doesn't work
```


```2020-03-14
  2000 times - Insert
  raw_stmt:     2.80s      1398198 ns/op     720 B/op     21 allocs/op
 beego_orm:     3.07s      1534607 ns/op    2441 B/op     58 allocs/op
        pg:     3.14s      1570768 ns/op     660 B/op     11 allocs/op
      xorm:     4.55s      2276666 ns/op    3387 B/op    109 allocs/op
       raw:     4.61s      2303380 ns/op     904 B/op     22 allocs/op
      gorm:     5.00s      2498648 ns/op    4061 B/op     73 allocs/op

  2000 times - MultiInsert 100 row
 beego_orm:     5.62s      2807532 ns/op  186801 B/op   2845 allocs/op
  raw_stmt:     6.43s      3213397 ns/op  125704 B/op   1417 allocs/op
        pg:     7.05s      3523873 ns/op    3892 B/op    111 allocs/op
       raw:     9.74s      4872435 ns/op  140177 B/op   1421 allocs/op
      xorm:     9.98s      4988471 ns/op  303824 B/op   7795 allocs/op
      gorm:    10.20s      5099222 ns/op  192564 B/op   2582 allocs/op

  4000 times - Update
  raw_stmt:     3.70s       924885 ns/op     728 B/op     21 allocs/op
 beego_orm:     5.84s      1459925 ns/op    1824 B/op     49 allocs/op
       raw:     6.41s      1603621 ns/op     904 B/op     22 allocs/op
        pg:     6.84s      1709939 ns/op     592 B/op      9 allocs/op
      gorm:     8.68s      2170005 ns/op    3920 B/op     78 allocs/op
      xorm:     9.07s      2268356 ns/op    3409 B/op    135 allocs/op

  4000 times - Read
      gorm:     3.28s       819406 ns/op    2234 B/op     56 allocs/op
 beego_orm:     3.32s       829673 ns/op    2160 B/op     80 allocs/op
  raw_stmt:     3.65s       911790 ns/op     920 B/op     28 allocs/op
        pg:     3.90s       974933 ns/op     728 B/op     17 allocs/op
       raw:     6.37s      1591971 ns/op    1504 B/op     41 allocs/op
      xorm:     8.12s      2031231 ns/op    9114 B/op    265 allocs/op

  4000 times - MultiRead limit 100
        pg:     5.27s      1317796 ns/op   24380 B/op    622 allocs/op
       raw:     5.48s      1369175 ns/op   30656 B/op   1513 allocs/op
  raw_stmt:     5.63s      1408448 ns/op   28656 B/op   1212 allocs/op
 beego_orm:     5.70s      1423965 ns/op   60872 B/op   3478 allocs/op
      gorm:     9.55s      2387885 ns/op   36206 B/op   1475 allocs/op
      xorm:     doesn't work
```
