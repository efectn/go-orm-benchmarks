# ORM Benchmark

_(forked from https://github.com/frederikhors/orm-benchmark)_

### Environment

- go version go1.16.4 windows/amd64

### PostgreSQL

- PostgreSQL 13 for Linux on x86_64

### ORMs

All package run in no-cache mode.

- [beego/orm](https://github.com/astaxie/beego/tree/master/orm)
- [bun](https://github.com/uptrace/bun)
- [gorm 2](https://github.com/go-gorm/gorm)
- [pg](https://github.com/go-pg/pg)
- [sqlc](https://github.com/kyleconroy/sqlc) (partially)
- [xorm](https://github.com/xormplus/xorm)
- [ent](https://github.com/ent/ent)
- [prisma](https://github.com/prisma/prisma-client-go)
- [upper](https://github.com/upper/db)
- [gorp](https://github.com/go-gorp/gorp)
- [godb](https://github.com/samonzeweb/godb)
- [dbr](https://github.com/gocraft/dbr/)
- [pop](https://github.com/gobuffalo/pop)
- [rel](https://github.com/go-rel/rel)
- [reform](https://github.com/go-reform/reform)

See [`go.mod`](https://github.com/efectn/go-orm-benchmarks/blob/master/go.mod) for their latest versions.

### Run

```shell
go get github.com/efectn/go-orm-benchmarks
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

From [`results.md`](https://github.com/efectn/go-orm-benchmarks/tree/master/results.md):

**orm-benchmark -multi=20**

```
  4000 times - Insert
  raw_stmt:    12.25s      3062956 ns/op     719 B/op     14 allocs/op
 sqlc_prep:    12.43s      3108665 ns/op    2917 B/op     62 allocs/op
      sqlc:    12.44s      3110837 ns/op    2901 B/op     62 allocs/op
       raw:    12.45s      3112034 ns/op     718 B/op     13 allocs/op
       ent:    12.50s      3124994 ns/op    4231 B/op     99 allocs/op
 beego_orm:    12.51s      3127372 ns/op    2409 B/op     56 allocs/op
      gorm:    12.67s      3168597 ns/op    6982 B/op     91 allocs/op
        pg:    12.72s      3179614 ns/op    1368 B/op     10 allocs/op
       bun:    12.87s      3217532 ns/op    5011 B/op     14 allocs/op
      xorm:    13.15s      3286973 ns/op    2837 B/op     78 allocs/op
 gorm_prep:    13.51s      3378548 ns/op    5295 B/op     69 allocs/op
     upper:    14.65s      3663552 ns/op   14417 B/op    682 allocs/op
    prisma:    74.29s     18572790 ns/op   11592 B/op    144 allocs/op

  4000 times - MultiInsert 100 row
  raw_stmt:    16.93s      4233161 ns/op  187179 B/op    931 allocs/op
       raw:    17.07s      4267315 ns/op  191258 B/op    931 allocs/op
 beego_orm:    17.25s      4312606 ns/op  179748 B/op   2746 allocs/op
 gorm_prep:    18.88s      4719702 ns/op  257057 B/op   2488 allocs/op
       ent:    19.28s      4819153 ns/op  412144 B/op   4899 allocs/op
        pg:    19.39s      4848449 ns/op    3857 B/op    112 allocs/op
       bun:    19.68s      4920315 ns/op   42474 B/op    219 allocs/op
      gorm:    20.62s      5153859 ns/op  294869 B/op   4034 allocs/op
      xorm:    21.48s      5369776 ns/op  266294 B/op   5821 allocs/op
      sqlc:     TBD
 sqlc_prep:     TBD
     upper:     TBD
    prisma:     doesn't work

  4000 times - Update
  raw_stmt:     0.56s       141179 ns/op     766 B/op     14 allocs/op
       raw:     0.57s       141288 ns/op     750 B/op     13 allocs/op
       ent:     0.64s       160611 ns/op    4556 B/op     99 allocs/op
 sqlc_prep:    12.29s      3073559 ns/op     894 B/op     15 allocs/op
      sqlc:    12.38s      3095962 ns/op     876 B/op     14 allocs/op
 beego_orm:    12.53s      3132803 ns/op    1801 B/op     47 allocs/op
 gorm_prep:    12.68s      3169936 ns/op    5116 B/op     59 allocs/op
        pg:    12.71s      3178600 ns/op     768 B/op      9 allocs/op
       bun:    12.74s      3184405 ns/op    4730 B/op      5 allocs/op
      gorm:    12.82s      3206087 ns/op    6604 B/op     81 allocs/op
      xorm:    13.53s      3383221 ns/op    2888 B/op    105 allocs/op
     upper:    17.45s      4363193 ns/op   34899 B/op   1524 allocs/op
    prisma:    61.43s     15358152 ns/op   12285 B/op    171 allocs/op

  4000 times - Read
       raw:     0.59s       147384 ns/op    2080 B/op     49 allocs/op
      sqlc:     0.59s       148641 ns/op    2190 B/op     51 allocs/op
  raw_stmt:     0.60s       149383 ns/op    2112 B/op     50 allocs/op
 sqlc_prep:     0.60s       150454 ns/op    2222 B/op     52 allocs/op
 beego_orm:     0.61s       153671 ns/op    2105 B/op     75 allocs/op
        pg:     0.66s       166172 ns/op     872 B/op     20 allocs/op
 gorm_prep:     0.70s       176180 ns/op    4972 B/op     98 allocs/op
       bun:     0.72s       180928 ns/op    5494 B/op     21 allocs/op
       ent:     0.73s       182916 ns/op    5418 B/op    148 allocs/op
      gorm:     0.96s       239778 ns/op    5304 B/op    109 allocs/op
     upper:     1.65s       412404 ns/op    8742 B/op    327 allocs/op
      xorm:     1.74s       434311 ns/op    8779 B/op    216 allocs/op
    prisma:    59.93s     14981387 ns/op    7928 B/op    128 allocs/op

  4000 times - MultiRead limit 100
       raw:     1.43s       356667 ns/op   38355 B/op   1037 allocs/op
  raw_stmt:     1.46s       364608 ns/op   38387 B/op   1038 allocs/op
 sqlc_prep:     1.67s       417137 ns/op   73208 B/op   1251 allocs/op
      sqlc:     1.68s       419150 ns/op   73176 B/op   1250 allocs/op
     upper:     1.71s       426983 ns/op    8415 B/op    304 allocs/op
        pg:     1.81s       452292 ns/op   22544 B/op    629 allocs/op
       ent:     1.99s       498549 ns/op   76065 B/op   2038 allocs/op
       bun:     2.32s       578887 ns/op   32899 B/op   1118 allocs/op
 beego_orm:     2.38s       596191 ns/op   55239 B/op   3077 allocs/op
 gorm_prep:     3.58s       894885 ns/op   71235 B/op   3577 allocs/op
      gorm:     3.95s       988350 ns/op   71628 B/op   3877 allocs/op
    prisma:    67.39s     16846721 ns/op  146641 B/op    853 allocs/op
      xorm:     doesn't work
```

### To-Do List:
- [ ] Remove or add prisma again.
- [ ] Fix TBDs.
- [ ] Add SqlBoiler.
- [x] Add go-reform.
- [ ] Clean and restructure codebase.
- [ ] Add makefile to run benchmarks easily, run benchmarks inside of container.
- [ ] Add final benchmark results with details.
- [ ] Tag release and be ready to compare ORMs performances.