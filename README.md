# ORM Benchmark

_(forked from https://github.com/yusaer/orm-benchmark)_

### Environment

- go version go1.16.4 windows/amd64

### PostgreSQL

- PostgreSQL 13 for Linux on x86_64

### ORMs

All package run in no-cache mode.

- [gorm 2](https://github.com/go-gorm/gorm)
- [pg](https://github.com/go-pg/pg)
- [bun](https://github.com/uptrace/bun)
- [beego/orm](https://github.com/astaxie/beego/tree/master/orm)
- [xorm](https://github.com/xormplus/xorm)

See [`go.mod`](https://github.com/frederikhors/orm-benchmark/blob/master/go.mod) for their latest versions.

### Run

```shell
go get github.com/frederikhors/orm-benchmark
# build
go install
# all
orm-benchmark -multi=20 -orm=all
# portion
orm-benchmark -multi=20 -orm=gorm
orm-benchmark -multi=20 -orm=pg
orm-benchmark -multi=20 -orm=bun
# ... and so on...
```

### Results

From [`results.md`](https://github.com/frederikhors/orm-benchmark/tree/master/results.md):

**orm-benchmark -multi=20**

```
  4000 times - Insert
  raw_stmt:     7.49s      1872449 ns/op     718 B/op     14 allocs/op
       raw:     7.50s      1875044 ns/op     727 B/op     13 allocs/op
 beego_orm:     7.56s      1890294 ns/op    2410 B/op     56 allocs/op
      sqlc:     7.61s      1902785 ns/op    2902 B/op     62 allocs/op
 sqlc_prep:     7.68s      1921107 ns/op    2917 B/op     62 allocs/op
       bun:     7.69s      1923437 ns/op     918 B/op     13 allocs/op
 gorm_prep:     7.71s      1927455 ns/op    4985 B/op     65 allocs/op
        pg:     7.88s      1968902 ns/op    1778 B/op     12 allocs/op
      gorm:     8.64s      2159774 ns/op    6676 B/op     88 allocs/op
      xorm:    10.36s      2591181 ns/op    3050 B/op     94 allocs/op

  4000 times - MultiInsert 100 row
  raw_stmt:    19.46s      4865525 ns/op  131076 B/op    916 allocs/op
       raw:    20.01s      5002064 ns/op  135155 B/op    916 allocs/op
 gorm_prep:    23.39s      5848246 ns/op  200057 B/op   2168 allocs/op
       bun:    23.85s      5963499 ns/op    4259 B/op    214 allocs/op
        pg:    23.88s      5970434 ns/op    3986 B/op    114 allocs/op
 beego_orm:    23.90s      5975433 ns/op  179815 B/op   2746 allocs/op
      gorm:    28.97s      7242372 ns/op  293955 B/op   3729 allocs/op
      xorm:    35.60s      8900321 ns/op  286030 B/op   7422 allocs/op
 sqlc_prep:     TBD
      sqlc:     TBD

  4000 times - Update
  raw_stmt:     2.62s       656011 ns/op     773 B/op     14 allocs/op
       raw:     2.63s       656770 ns/op     757 B/op     13 allocs/op
 sqlc_prep:     7.46s      1864336 ns/op     894 B/op     15 allocs/op
 beego_orm:     7.50s      1874009 ns/op    1801 B/op     47 allocs/op
      sqlc:     7.52s      1879374 ns/op     876 B/op     14 allocs/op
 gorm_prep:     7.60s      1900330 ns/op    5123 B/op     59 allocs/op
       bun:     7.68s      1920855 ns/op     591 B/op      4 allocs/op
      gorm:     7.82s      1954842 ns/op    6604 B/op     81 allocs/op
        pg:     8.09s      2022080 ns/op     896 B/op     11 allocs/op
      xorm:    10.27s      2566399 ns/op    2994 B/op    119 allocs/op

  4000 times - Read
 beego_orm:     2.64s       658829 ns/op    2105 B/op     75 allocs/op
      sqlc:     2.64s       659501 ns/op    2190 B/op     51 allocs/op
 sqlc_prep:     2.65s       662342 ns/op    2222 B/op     52 allocs/op
       raw:     2.65s       662669 ns/op    2080 B/op     49 allocs/op
  raw_stmt:     2.69s       672677 ns/op    2112 B/op     50 allocs/op
 gorm_prep:     2.75s       686810 ns/op    4908 B/op     97 allocs/op
       bun:     2.76s       689577 ns/op    1313 B/op     18 allocs/op
        pg:     2.84s       710984 ns/op    1000 B/op     22 allocs/op
      gorm:     2.91s       726527 ns/op    5240 B/op    108 allocs/op
      xorm:     5.69s      1423129 ns/op    8322 B/op    237 allocs/op

  4000 times - MultiRead limit 100
       raw:     3.56s       889343 ns/op   38356 B/op   1037 allocs/op
  raw_stmt:     3.57s       891465 ns/op   38388 B/op   1038 allocs/op
      sqlc:     3.72s       930181 ns/op   73176 B/op   1250 allocs/op
 sqlc_prep:     3.73s       933592 ns/op   73208 B/op   1251 allocs/op
        pg:     3.93s       981776 ns/op   24532 B/op    631 allocs/op
       bun:     4.07s      1016640 ns/op   28858 B/op   1116 allocs/op
 beego_orm:     4.19s      1048494 ns/op   55253 B/op   3077 allocs/op
 gorm_prep:     4.96s      1240677 ns/op   71236 B/op   3577 allocs/op
      gorm:     5.19s      1297135 ns/op   71628 B/op   3877 allocs/op
      xorm:     doesn't work
```
