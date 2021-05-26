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
       raw:     7.66s      1914203 ns/op     719 B/op     13 allocs/op
       bun:     7.79s      1948428 ns/op     921 B/op     13 allocs/op
      gorm:     7.85s      1962616 ns/op    6662 B/op     87 allocs/op
 beego_orm:     7.95s      1987661 ns/op    2409 B/op     56 allocs/op
      sqlc:     7.97s      1993002 ns/op    2902 B/op     62 allocs/op
        pg:     8.20s      2050734 ns/op    1235 B/op     12 allocs/op
  raw_stmt:     9.02s      2255527 ns/op     718 B/op     14 allocs/op
      xorm:    10.75s      2687473 ns/op    3050 B/op     94 allocs/op

  4000 times - MultiInsert 100 row
  raw_stmt:    21.68s      5420099 ns/op  131075 B/op    916 allocs/op
       raw:    21.99s      5497700 ns/op  135154 B/op    916 allocs/op
 beego_orm:    23.58s      5895580 ns/op  179785 B/op   2746 allocs/op
        pg:    26.07s      6516842 ns/op    4530 B/op    114 allocs/op
       bun:    27.15s      6786342 ns/op    4257 B/op    214 allocs/op
      gorm:    29.97s      7491448 ns/op  293955 B/op   3729 allocs/op
      xorm:    35.68s      8919414 ns/op  285963 B/op   7422 allocs/op
      sqlc:     TBD

  4000 times - Update
       raw:     2.67s       668537 ns/op     757 B/op     13 allocs/op
  raw_stmt:     2.75s       688248 ns/op     773 B/op     14 allocs/op
        pg:     7.82s      1955985 ns/op     896 B/op     11 allocs/op
      sqlc:     7.98s      1994035 ns/op     876 B/op     14 allocs/op
      gorm:     8.10s      2025805 ns/op    6604 B/op     81 allocs/op
       bun:     8.38s      2096225 ns/op     587 B/op      4 allocs/op
 beego_orm:     8.48s      2120507 ns/op    1801 B/op     47 allocs/op
      xorm:    10.69s      2673311 ns/op    2994 B/op    119 allocs/op

  4000 times - Read
       raw:     2.64s       659401 ns/op    2080 B/op     49 allocs/op
  raw_stmt:     2.70s       674592 ns/op    2112 B/op     50 allocs/op
 beego_orm:     2.71s       677034 ns/op    2105 B/op     75 allocs/op
      sqlc:     2.83s       706332 ns/op    2190 B/op     51 allocs/op
        pg:     2.86s       715370 ns/op    1000 B/op     22 allocs/op
       bun:     2.87s       717030 ns/op    1315 B/op     18 allocs/op
      gorm:     2.94s       734625 ns/op    5240 B/op    108 allocs/op
      xorm:     5.87s      1467461 ns/op    8322 B/op    237 allocs/op

  4000 times - MultiRead limit 100
      sqlc:     3.73s       932872 ns/op   74321 B/op   1257 allocs/op
  raw_stmt:     3.75s       936363 ns/op   38387 B/op   1038 allocs/op
       raw:     3.86s       964819 ns/op   38356 B/op   1037 allocs/op
        pg:     3.88s       970220 ns/op   23989 B/op    631 allocs/op
       bun:     4.07s      1018338 ns/op   28833 B/op   1116 allocs/op
 beego_orm:     4.22s      1054243 ns/op   55245 B/op   3077 allocs/op
      gorm:     5.18s      1295536 ns/op   71628 B/op   3877 allocs/op
      xorm:     doesn't work
```
