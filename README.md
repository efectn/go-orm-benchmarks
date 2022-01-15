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
  raw_stmt:    14.48s      3621006 ns/op     719 B/op     14 allocs/op
 sqlc_prep:    14.53s      3632453 ns/op    2917 B/op     62 allocs/op
 beego_orm:    14.70s      3674523 ns/op    2408 B/op     56 allocs/op
 gorm_prep:    14.71s      3678378 ns/op    5305 B/op     69 allocs/op
      sqlc:    14.84s      3710283 ns/op    2901 B/op     62 allocs/op
       raw:    14.91s      3726764 ns/op     718 B/op     13 allocs/op
       bun:    15.28s      3819388 ns/op    5010 B/op     14 allocs/op
      gorm:    15.47s      3867065 ns/op    6973 B/op     91 allocs/op
        pg:    15.97s      3993144 ns/op    1106 B/op     10 allocs/op
      xorm:    16.06s      4014147 ns/op    2838 B/op     78 allocs/op
       ent:    17.38s      4346192 ns/op    4892 B/op    115 allocs/op

  4000 times - MultiInsert 100 row
       ent:     0.01s         2137 ns/op     116 B/op      1 allocs/op
  raw_stmt:    20.66s      5165122 ns/op  187179 B/op    931 allocs/op
       raw:    21.08s      5270335 ns/op  191258 B/op    931 allocs/op
 beego_orm:    21.55s      5388559 ns/op  179676 B/op   2746 allocs/op
 gorm_prep:    23.21s      5801483 ns/op  257060 B/op   2488 allocs/op
        pg:    23.21s      5802316 ns/op    4645 B/op    112 allocs/op
       bun:    24.02s      6003825 ns/op   42479 B/op    219 allocs/op
      gorm:    25.72s      6429072 ns/op  294870 B/op   4034 allocs/op
      xorm:    27.87s      6967649 ns/op  266645 B/op   5822 allocs/op
 sqlc_prep:     TBD
      sqlc:     TBD

  4000 times - Update
  raw_stmt:     0.60s       150135 ns/op     766 B/op     14 allocs/op
       raw:     0.61s       151921 ns/op     750 B/op     13 allocs/op
       ent:     2.06s       515320 ns/op    4845 B/op    109 allocs/op
 sqlc_prep:    14.32s      3580655 ns/op     894 B/op     15 allocs/op
 beego_orm:    14.71s      3677928 ns/op    1801 B/op     47 allocs/op
 gorm_prep:    14.82s      3706242 ns/op    5116 B/op     59 allocs/op
        pg:    15.22s      3804301 ns/op    1048 B/op      9 allocs/op
       bun:    15.25s      3811601 ns/op    4731 B/op      5 allocs/op
      sqlc:    15.29s      3822453 ns/op     876 B/op     14 allocs/op
      gorm:    15.32s      3829777 ns/op    6604 B/op     81 allocs/op
      xorm:    16.96s      4239943 ns/op    2889 B/op    105 allocs/op

  4000 times - Read
       raw:     0.68s       169150 ns/op    2080 B/op     49 allocs/op
      sqlc:     0.69s       172446 ns/op    2190 B/op     51 allocs/op
  raw_stmt:     0.70s       174863 ns/op    2112 B/op     50 allocs/op
 sqlc_prep:     0.72s       179857 ns/op    2222 B/op     52 allocs/op
 beego_orm:     0.73s       182812 ns/op    2104 B/op     75 allocs/op
        pg:     0.83s       207801 ns/op     872 B/op     20 allocs/op
 gorm_prep:     0.91s       228216 ns/op    4972 B/op     98 allocs/op
       bun:     0.97s       242026 ns/op    5493 B/op     21 allocs/op
       ent:     0.97s       242071 ns/op    5106 B/op    139 allocs/op
      gorm:     1.27s       316485 ns/op    5304 B/op    109 allocs/op
      xorm:     2.23s       558664 ns/op    8780 B/op    216 allocs/op

  4000 times - MultiRead limit 100
       raw:     2.10s       523769 ns/op   38355 B/op   1037 allocs/op
  raw_stmt:     2.11s       527595 ns/op   38388 B/op   1038 allocs/op
 sqlc_prep:     2.31s       578306 ns/op   73208 B/op   1251 allocs/op
      sqlc:     2.38s       595368 ns/op   73176 B/op   1250 allocs/op
        pg:     2.71s       677566 ns/op   23615 B/op    629 allocs/op
 beego_orm:     3.25s       813020 ns/op   55234 B/op   3077 allocs/op
       bun:     3.52s       880325 ns/op   32908 B/op   1118 allocs/op
 gorm_prep:     4.96s      1239943 ns/op   71236 B/op   3577 allocs/op
      gorm:     5.69s      1421485 ns/op   71628 B/op   3877 allocs/op
       ent:    16.00s      4000532 ns/op   77082 B/op   2128 allocs/op
      xorm:     doesn't work```
