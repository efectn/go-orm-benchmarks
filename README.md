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

See [`go.mod`](https://github.com/efectn/orm-benchmark/blob/master/go.mod) for their latest versions.

### Run

```shell
go get github.com/efectn/orm-benchmark
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

From [`results.md`](https://github.com/efectn/orm-benchmark/tree/master/results.md):

**orm-benchmark -multi=20**

```
  4000 times - Insert
       ent:    12.40s      3100093 ns/op    4230 B/op     99 allocs/op
  raw_stmt:    12.44s      3109874 ns/op     718 B/op     14 allocs/op
      sqlc:    12.46s      3115430 ns/op    2901 B/op     62 allocs/op
 gorm_prep:    12.56s      3139645 ns/op    5305 B/op     69 allocs/op
 sqlc_prep:    12.58s      3144080 ns/op    2917 B/op     62 allocs/op
 beego_orm:    12.64s      3159109 ns/op    2408 B/op     56 allocs/op
        pg:    12.68s      3170189 ns/op    1106 B/op     10 allocs/op
      gorm:    12.75s      3186849 ns/op    6982 B/op     92 allocs/op
      xorm:    13.13s      3281345 ns/op    2838 B/op     78 allocs/op
       raw:    13.93s      3481535 ns/op     718 B/op     13 allocs/op
     upper:    14.05s      3513294 ns/op   14421 B/op    682 allocs/op
       bun:    14.75s      3688283 ns/op    5001 B/op     14 allocs/op
    prisma:    73.75s     18437112 ns/op   11598 B/op    144 allocs/op

  4000 times - MultiInsert 100 row
  raw_stmt:    17.03s      4257882 ns/op  187179 B/op    931 allocs/op
 beego_orm:    17.19s      4296463 ns/op  179827 B/op   2746 allocs/op
       raw:    17.27s      4317134 ns/op  191258 B/op    931 allocs/op
 gorm_prep:    18.25s      4562475 ns/op  257058 B/op   2488 allocs/op
        pg:    19.18s      4794703 ns/op    3858 B/op    112 allocs/op
       bun:    19.46s      4864734 ns/op   42403 B/op    219 allocs/op
       ent:    20.53s      5132518 ns/op  412884 B/op   4913 allocs/op
      gorm:    20.92s      5229956 ns/op  294869 B/op   4034 allocs/op
      xorm:    21.74s      5436193 ns/op  266712 B/op   5822 allocs/op
 sqlc_prep:     TBD
     upper:     TBD
      sqlc:     TBD
    prisma:     doesn't work

  4000 times - Update
  raw_stmt:     0.57s       142535 ns/op     766 B/op     14 allocs/op
       raw:     0.57s       143560 ns/op     750 B/op     13 allocs/op
       ent:     0.67s       167592 ns/op    4556 B/op     99 allocs/op
 sqlc_prep:    12.32s      3079428 ns/op     894 B/op     15 allocs/op
      sqlc:    12.43s      3107383 ns/op     876 B/op     14 allocs/op
 gorm_prep:    12.54s      3136125 ns/op    5116 B/op     59 allocs/op
 beego_orm:    12.63s      3157956 ns/op    1801 B/op     47 allocs/op
       bun:    12.77s      3192386 ns/op    4730 B/op      5 allocs/op
      gorm:    12.97s      3241645 ns/op    6614 B/op     81 allocs/op
      xorm:    13.38s      3345806 ns/op    2889 B/op    105 allocs/op
        pg:    13.83s      3457470 ns/op     768 B/op      9 allocs/op
     upper:    17.81s      4453723 ns/op   34888 B/op   1524 allocs/op
    prisma:    61.35s     15337657 ns/op   12274 B/op    171 allocs/op

  4000 times - Read
  raw_stmt:     0.59s       148074 ns/op    2112 B/op     50 allocs/op
       raw:     0.59s       148321 ns/op    2080 B/op     49 allocs/op
      sqlc:     0.60s       148928 ns/op    2190 B/op     51 allocs/op
 sqlc_prep:     0.60s       150150 ns/op    2222 B/op     52 allocs/op
 beego_orm:     0.62s       155403 ns/op    2105 B/op     75 allocs/op
        pg:     0.69s       171869 ns/op     872 B/op     20 allocs/op
 gorm_prep:     0.72s       179336 ns/op    4972 B/op     98 allocs/op
       bun:     0.73s       181941 ns/op    5492 B/op     21 allocs/op
       ent:     0.76s       189067 ns/op    5418 B/op    148 allocs/op
      gorm:     0.97s       242200 ns/op    5304 B/op    109 allocs/op
     upper:     1.70s       426082 ns/op    8743 B/op    327 allocs/op
      xorm:     1.75s       438664 ns/op    8781 B/op    216 allocs/op
    prisma:    59.97s     14991897 ns/op    7920 B/op    128 allocs/op

  4000 times - MultiRead limit 100
       raw:     1.44s       360034 ns/op   38355 B/op   1037 allocs/op
  raw_stmt:     1.45s       362037 ns/op   38387 B/op   1038 allocs/op
      sqlc:     1.66s       415830 ns/op   73176 B/op   1250 allocs/op
 sqlc_prep:     1.69s       422382 ns/op   73208 B/op   1251 allocs/op
        pg:     1.81s       452676 ns/op   23350 B/op    629 allocs/op
       ent:     2.09s       522594 ns/op   76065 B/op   2038 allocs/op
       bun:     2.27s       566301 ns/op   32890 B/op   1118 allocs/op
 beego_orm:     2.42s       604722 ns/op   55240 B/op   3077 allocs/op
     upper:     3.08s       770824 ns/op    8417 B/op    304 allocs/op
 gorm_prep:     3.64s       910841 ns/op   71236 B/op   3577 allocs/op
      gorm:     4.04s      1009197 ns/op   71628 B/op   3877 allocs/op
    prisma:    67.63s     16908734 ns/op  146860 B/op    853 allocs/op
      xorm:     doesn't work
```
