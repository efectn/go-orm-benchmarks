# ORM Benchmark

_(forked from https://github.com/yusaer/orm-benchmark)_

### Environment

- go version go1.15.1 linux/amd64

### PostgreSQL

- PostgreSQL 13 for Linux on x86_64

### ORMs

All package run in no-cache mode.

- [gorm 2](https://github.com/go-gorm/gorm)
- [pg](https://github.com/go-pg/pg)
- [bun](https://github.com/uptrace/bun)
- [beego/orm](https://github.com/astaxie/beego/tree/master/orm)
- [xorm](https://github.com/xormplus/xorm)

### Run

```shell
go get github.com/frederikhors/orm-benchmark
# build
go install
# all
orm-benchmark -multi=20 -orm=all
# portion
orm-benchmark -multi=20 -orm=gorm
```

### Reports

- **orm-benchmark**

```
   200 times - Insert
      gorm:     0.83s      4151199 ns/op    5739 B/op     70 allocs/op
 beego_orm:     0.84s      4215179 ns/op    2446 B/op     56 allocs/op
  raw_stmt:     0.87s      4331976 ns/op     696 B/op     18 allocs/op
        pg:     0.89s      4453546 ns/op    2726 B/op     11 allocs/op
       raw:     1.00s      5006239 ns/op    1331 B/op     20 allocs/op
      xorm:     1.03s      5174285 ns/op    3601 B/op     96 allocs/op

   200 times - MultiInsert 100 row
        pg:    10.49s     52425163 ns/op    4157 B/op    112 allocs/op
      gorm:    10.49s     52462248 ns/op  205225 B/op   2170 allocs/op
  raw_stmt:    10.58s     52893563 ns/op  125099 B/op   1223 allocs/op
 beego_orm:    10.59s     52930793 ns/op  186210 B/op   2745 allocs/op
       raw:    10.67s     53328383 ns/op  139504 B/op   1227 allocs/op
      xorm:    20.81s    104035686 ns/op  294867 B/op   7422 allocs/op

   200 times - Update
  raw_stmt:     0.21s      1073649 ns/op     712 B/op     19 allocs/op
       raw:     0.35s      1759354 ns/op     888 B/op     20 allocs/op
      gorm:     0.85s      4247077 ns/op    5339 B/op     62 allocs/op
 beego_orm:     0.85s      4260391 ns/op    1824 B/op     47 allocs/op
        pg:     0.88s      4388607 ns/op     808 B/op     11 allocs/op
      xorm:     1.08s      5381148 ns/op    3100 B/op    119 allocs/op

   200 times - Read
 beego_orm:     0.18s       892952 ns/op    2141 B/op     75 allocs/op
      gorm:     0.18s       910013 ns/op    4816 B/op     97 allocs/op
        pg:     0.19s       935001 ns/op    1913 B/op     16 allocs/op
  raw_stmt:     0.23s      1171310 ns/op     888 B/op     24 allocs/op
       raw:     0.36s      1794155 ns/op    1472 B/op     37 allocs/op
      xorm:     0.38s      1920296 ns/op    8581 B/op    237 allocs/op

   200 times - MultiRead limit 100
       raw:     0.25s      1265075 ns/op   28240 B/op   1312 allocs/op
        pg:     0.26s      1289987 ns/op   25652 B/op    622 allocs/op
  raw_stmt:     0.28s      1387518 ns/op   26256 B/op   1011 allocs/op
 beego_orm:     0.29s      1445014 ns/op   57690 B/op   3077 allocs/op
      gorm:     0.35s      1740135 ns/op   73186 B/op   3581 allocs/op
      xorm:     doesn't work
```

- **orm-benchmark -multi=20**

```
  4000 times - Insert
  raw_stmt:    16.95s      4237944 ns/op     696 B/op     18 allocs/op
 beego_orm:    17.14s      4286242 ns/op    2441 B/op     56 allocs/op
      gorm:    17.15s      4288553 ns/op    5072 B/op     68 allocs/op
        pg:    17.47s      4367692 ns/op    1001 B/op     10 allocs/op
      xorm:    20.84s      5208774 ns/op    3137 B/op     94 allocs/op
       raw:    20.96s      5240486 ns/op     895 B/op     19 allocs/op

  4000 times - MultiInsert 100 row
 beego_orm:    209.08s     52270107 ns/op  186185 B/op   2745 allocs/op
      gorm:    209.83s     52456385 ns/op  204792 B/op   2172 allocs/op
        pg:    210.04s     52510175 ns/op    4432 B/op    112 allocs/op
  raw_stmt:    210.55s     52638061 ns/op  124040 B/op   1216 allocs/op
       raw:    211.72s     52929722 ns/op  138516 B/op   1220 allocs/op
      xorm:    410.86s    102714839 ns/op  294723 B/op   7422 allocs/op

  4000 times - Update
  raw_stmt:     3.51s       878663 ns/op     712 B/op     19 allocs/op
       raw:     7.07s      1767725 ns/op     888 B/op     20 allocs/op
      gorm:    17.09s      4272306 ns/op    5259 B/op     62 allocs/op
 beego_orm:    17.18s      4295715 ns/op    1817 B/op     47 allocs/op
        pg:    17.33s      4331658 ns/op     923 B/op     11 allocs/op
      xorm:    21.06s      5264025 ns/op    3097 B/op    119 allocs/op

  4000 times - Read
 beego_orm:     3.53s       882455 ns/op    2128 B/op     75 allocs/op
      gorm:     3.73s       931874 ns/op    4772 B/op     97 allocs/op
        pg:     3.74s       935846 ns/op     944 B/op     16 allocs/op
  raw_stmt:     4.02s      1004141 ns/op     888 B/op     24 allocs/op
      xorm:     7.59s      1897352 ns/op    8575 B/op    237 allocs/op
       raw:    10.36s      2589251 ns/op    1472 B/op     37 allocs/op

  4000 times - MultiRead limit 100
  raw_stmt:     4.74s      1185603 ns/op   26256 B/op   1011 allocs/op
       raw:     5.52s      1380345 ns/op   28240 B/op   1312 allocs/op
        pg:     5.69s      1422820 ns/op   25839 B/op    622 allocs/op
 beego_orm:     5.81s      1452369 ns/op   57692 B/op   3077 allocs/op
      gorm:     6.94s      1736142 ns/op   73108 B/op   3581 allocs/op
      xorm:     doesn't work
```
