# Go ORM Benchmarks

[![Go Reference](https://pkg.go.dev/badge/github.com/efectn/go-orm-benchmarks.svg)](https://pkg.go.dev/github.com/efectn/go-orm-benchmarks)

Advanced benchmarks for +10 Go ORMs. Originally forked from [orm-benchmark](https://github.com/frederikhors/orm-benchmark).

### ORMs

All package run in no-cache mode.

- [beego/orm](https://github.com/astaxie/beego/tree/master/orm)
- [bun](https://github.com/uptrace/bun)
- [gorm 2](https://github.com/go-gorm/gorm)
- [pg](https://github.com/go-pg/pg)
- [sqlc](https://github.com/kyleconroy/sqlc)
- [xorm](https://github.com/xormplus/xorm)
- [ent](https://github.com/ent/ent)
- [upper](https://github.com/upper/db)
- [gorp](https://github.com/go-gorp/gorp)
- [godb](https://github.com/samonzeweb/godb)
- [dbr](https://github.com/gocraft/dbr/)
- [pop](https://github.com/gobuffalo/pop)
- [rel](https://github.com/go-rel/rel)
- [reform](https://github.com/go-reform/reform)
- [sqlboiler](https://github.com/volatiletech/sqlboiler)
- [sqlx](https://github.com/jmoiron/sqlx)

See [`go.mod`](go.mod) for their latest versions.

### Run

```shell
# install
go install github.com/efectn/go-orm-benchmarks@latest
# all
go-orm-benchmarks -multi=20 -orm=all
# portion
go-orm-benchmarks -multi=20 -orm=gorm
go-orm-benchmarks -multi=10 -orm=pg
go-orm-benchmarks -multi=15 -orm=bun
# ... and so on...
```

**_Note: Also, you can run `./run_benchmarks.sh` and you can get output like results.md format._**

### Results

#### Benchmark Enviroment:

- go version go1.18.1 linux/amd64
- Intel(R) Core(TM) i5-5200U CPU @ 2.20GHz
- 16GB 1600MHz memory
- PostgreSQL 14.2

From [`results.md`](results.md):

**orm-benchmark -multi=20**

```
  4000 times - Insert
       raw:    13.02s      3255212 ns/op     720 B/op     13 allocs/op
      gorp:    13.16s      3288827 ns/op    1783 B/op     42 allocs/op
      sqlc:    13.18s      3294672 ns/op    2885 B/op     63 allocs/op
 sqlboiler:    13.19s      3298647 ns/op    1576 B/op     35 allocs/op
       dbr:    13.47s      3367938 ns/op    2693 B/op     65 allocs/op
        pg:    13.58s      3393871 ns/op     806 B/op     10 allocs/op
       bun:    14.03s      3506665 ns/op    5014 B/op     14 allocs/op
      gorm:    14.56s      3639400 ns/op    6933 B/op     90 allocs/op
       ent:    14.77s      3693396 ns/op    4216 B/op    100 allocs/op
     beego:    14.91s      3726279 ns/op    2376 B/op     56 allocs/op
    reform:    15.27s      3816701 ns/op    1792 B/op     51 allocs/op
       rel:    15.30s      3824917 ns/op    2504 B/op     42 allocs/op
      xorm:    15.32s      3829436 ns/op    3327 B/op     89 allocs/op
      godb:    15.85s      3962599 ns/op    4559 B/op    115 allocs/op
     upper:    16.78s      4194212 ns/op   13775 B/op    673 allocs/op
       pop:    18.49s      4622682 ns/op   10084 B/op    249 allocs/op

  4000 times - MultiInsert 100 row
       raw:    25.01s      6253405 ns/op  191258 B/op    931 allocs/op
     beego:    25.12s      6279130 ns/op  179717 B/op   2746 allocs/op
        pg:    26.51s      6626510 ns/op    4120 B/op    112 allocs/op
    reform:    27.32s      6831205 ns/op  466461 B/op   2748 allocs/op
       bun:    28.48s      7121107 ns/op   42484 B/op    219 allocs/op
       ent:    29.63s      7406418 ns/op  412129 B/op   4900 allocs/op
      gorm:    32.42s      8105585 ns/op  272300 B/op   3729 allocs/op
       rel:    33.48s      8370838 ns/op  303994 B/op   3263 allocs/op
      godb:    35.75s      8938434 ns/op  260078 B/op   5895 allocs/op
      xorm:    40.19s     10046385 ns/op  255168 B/op   5417 allocs/op
     upper:    46.22s     11555653 ns/op  545167 B/op  19473 allocs/op
      gorp:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert

  4000 times - Update
       raw:     0.42s       106196 ns/op     750 B/op     13 allocs/op
       ent:     0.56s       140521 ns/op    4558 B/op     99 allocs/op
    reform:    13.42s      3355531 ns/op    1775 B/op     51 allocs/op
      gorp:    13.45s      3362383 ns/op    1204 B/op     32 allocs/op
      sqlc:    13.47s      3368367 ns/op     876 B/op     14 allocs/op
     beego:    13.60s      3399569 ns/op    1753 B/op     47 allocs/op
        pg:    14.26s      3564773 ns/op     768 B/op      9 allocs/op
       bun:    14.33s      3581336 ns/op    4732 B/op      5 allocs/op
      gorm:    14.49s      3621870 ns/op    6604 B/op     81 allocs/op
 sqlboiler:    14.61s      3652173 ns/op     901 B/op     17 allocs/op
       dbr:    14.95s      3737299 ns/op    2651 B/op     57 allocs/op
      xorm:    15.11s      3777118 ns/op    3650 B/op    126 allocs/op
       pop:    15.57s      3892037 ns/op    6586 B/op    198 allocs/op
      godb:    16.16s      4040401 ns/op    5162 B/op    154 allocs/op
       rel:    16.84s      4209090 ns/op    2528 B/op     41 allocs/op
     upper:    25.06s      6263904 ns/op   33503 B/op   1502 allocs/op

  4000 times - Read
      sqlc:     0.49s       122556 ns/op    2172 B/op     52 allocs/op
       raw:     0.49s       123397 ns/op    2062 B/op     50 allocs/op
    reform:     0.51s       128650 ns/op    3199 B/op     86 allocs/op
       pop:     0.56s       139620 ns/op    3555 B/op     71 allocs/op
       bun:     0.57s       142709 ns/op    5495 B/op     21 allocs/op
       rel:     0.59s       146392 ns/op    1800 B/op     45 allocs/op
 sqlboiler:     0.63s       157321 ns/op     965 B/op     14 allocs/op
     beego:     0.63s       158172 ns/op    2088 B/op     75 allocs/op
        pg:     0.64s       160565 ns/op     872 B/op     20 allocs/op
       ent:     0.67s       168133 ns/op    5404 B/op    149 allocs/op
       dbr:     0.70s       174635 ns/op    2184 B/op     37 allocs/op
      gorp:     0.72s       180514 ns/op    3878 B/op    194 allocs/op
      gorm:     0.91s       227725 ns/op    5152 B/op     94 allocs/op
      godb:     1.22s       306069 ns/op    4113 B/op    102 allocs/op
      xorm:     1.31s       326767 ns/op    4618 B/op    125 allocs/op
     upper:     2.18s       545960 ns/op    8136 B/op    315 allocs/op

  4000 times - MultiRead limit 100
    reform:     0.72s       180644 ns/op    2927 B/op     74 allocs/op
       raw:     1.48s       370397 ns/op   38342 B/op   1038 allocs/op
     upper:     1.53s       382464 ns/op    7841 B/op    294 allocs/op
        pg:     1.87s       466911 ns/op   23856 B/op    629 allocs/op
      sqlc:     1.88s       468887 ns/op   73158 B/op   1251 allocs/op
 sqlboiler:     2.06s       516048 ns/op   58594 B/op   1260 allocs/op
       ent:     2.35s       587629 ns/op   76045 B/op   2039 allocs/op
      gorp:     2.37s       593522 ns/op   57389 B/op   1494 allocs/op
       pop:     2.74s       684419 ns/op   75921 B/op   1511 allocs/op
       bun:     2.80s       699047 ns/op   34025 B/op   1124 allocs/op
       dbr:     2.85s       711905 ns/op   32416 B/op   1545 allocs/op
     beego:     2.87s       717417 ns/op   55201 B/op   3077 allocs/op
      gorm:     3.66s       913801 ns/op   57181 B/op   2278 allocs/op
       rel:     6.59s      1647903 ns/op   95321 B/op   2248 allocs/op
      godb:     7.41s      1851727 ns/op   75289 B/op   3084 allocs/op
      xorm:     7.54s      1884039 ns/op  120000 B/op   4687 allocs/op
```

### To-Do List:

- [ ] Add [sqlh](github.com/nofeaturesonlybugs/sqlh).
- [ ] Use benchmark tools of `testing` as benchmark suite.

### License

go-orm-benchmarks is licensed under the terms of the **MIT License** (see [LICENSE](LICENSE)).
