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
 gorm_prep:    14.10s      3524961 ns/op    5303 B/op     69 allocs/op
 beego_orm:    14.16s      3541129 ns/op    2408 B/op     56 allocs/op
       raw:    14.53s      3632464 ns/op     718 B/op     13 allocs/op
       bun:    14.67s      3666504 ns/op    5011 B/op     14 allocs/op
      sqlc:    14.84s      3711113 ns/op    2901 B/op     62 allocs/op
 sqlc_prep:    14.86s      3714881 ns/op    2917 B/op     62 allocs/op
       ent:    14.93s      3732228 ns/op    4211 B/op     99 allocs/op
        pg:    14.98s      3744975 ns/op    1106 B/op     10 allocs/op
  raw_stmt:    15.30s      3824209 ns/op     718 B/op     14 allocs/op
      gorm:    15.49s      3873033 ns/op    6980 B/op     91 allocs/op
      xorm:    15.81s      3951594 ns/op    2838 B/op     78 allocs/op

  4000 times - MultiInsert 100 row
       ent:     0.01s         2208 ns/op     116 B/op      1 allocs/op
  raw_stmt:    20.08s      5019931 ns/op  187180 B/op    931 allocs/op
       raw:    20.85s      5213338 ns/op  191258 B/op    931 allocs/op
 beego_orm:    22.54s      5635495 ns/op  179810 B/op   2746 allocs/op
        pg:    23.29s      5822309 ns/op    3840 B/op    112 allocs/op
 gorm_prep:    23.35s      5837312 ns/op  257060 B/op   2488 allocs/op
       bun:    24.55s      6136818 ns/op   42452 B/op    219 allocs/op
      gorm:    26.08s      6520473 ns/op  294869 B/op   4034 allocs/op
      xorm:    27.44s      6859227 ns/op  266522 B/op   5822 allocs/op
      sqlc:     TBD
 sqlc_prep:     TBD

  4000 times - Update
       raw:     0.59s       147603 ns/op     750 B/op     13 allocs/op
  raw_stmt:     0.71s       176958 ns/op     766 B/op     14 allocs/op
       ent:     0.83s       208738 ns/op    4556 B/op     99 allocs/op
 sqlc_prep:    14.53s      3631855 ns/op     894 B/op     15 allocs/op
        pg:    14.62s      3655968 ns/op     768 B/op      9 allocs/op
      sqlc:    14.66s      3664894 ns/op     876 B/op     14 allocs/op
 beego_orm:    14.83s      3708351 ns/op    1801 B/op     47 allocs/op
 gorm_prep:    15.05s      3763065 ns/op    5116 B/op     59 allocs/op
      gorm:    15.49s      3873669 ns/op    6604 B/op     81 allocs/op
       bun:    16.28s      4070612 ns/op    4730 B/op      5 allocs/op
      xorm:    16.69s      4173309 ns/op    2889 B/op    105 allocs/op

  4000 times - Read
       raw:     0.68s       170174 ns/op    2080 B/op     49 allocs/op
      sqlc:     0.70s       175305 ns/op    2190 B/op     51 allocs/op
  raw_stmt:     0.71s       177852 ns/op    2112 B/op     50 allocs/op
 sqlc_prep:     0.72s       180413 ns/op    2222 B/op     52 allocs/op
 beego_orm:     0.75s       186608 ns/op    2105 B/op     75 allocs/op
        pg:     0.85s       212069 ns/op     872 B/op     20 allocs/op
 gorm_prep:     0.87s       218067 ns/op    4972 B/op     98 allocs/op
       bun:     0.98s       245944 ns/op    5496 B/op     21 allocs/op
       ent:     0.99s       246346 ns/op    5402 B/op    148 allocs/op
      gorm:     1.31s       327288 ns/op    5303 B/op    109 allocs/op
      xorm:     2.26s       564706 ns/op    8779 B/op    216 allocs/op

  4000 times - MultiRead limit 100
  raw_stmt:     2.03s       508420 ns/op   38388 B/op   1038 allocs/op
       raw:     2.15s       538271 ns/op   38356 B/op   1037 allocs/op
 sqlc_prep:     2.38s       594444 ns/op   73208 B/op   1251 allocs/op
      sqlc:     2.39s       597023 ns/op   73176 B/op   1250 allocs/op
        pg:     2.61s       653459 ns/op   23314 B/op    629 allocs/op
 beego_orm:     3.33s       832417 ns/op   55245 B/op   3077 allocs/op
       bun:     3.54s       884663 ns/op   32906 B/op   1118 allocs/op
 gorm_prep:     4.91s      1226265 ns/op   71236 B/op   3577 allocs/op
      gorm:     5.64s      1410497 ns/op   71628 B/op   3877 allocs/op
       ent:    18.05s      4512000 ns/op   77378 B/op   2137 allocs/op
      xorm:     doesn't work
```
