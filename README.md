# Go ORM Benchmarks

[![Go Reference](https://pkg.go.dev/badge/github.com/efectn/go-orm-benchmarks.svg)](https://pkg.go.dev/github.com/efectn/go-orm-benchmarks)

Advanced benchmarks for +10 Go ORMs. Originally forked from [orm-benchmark](https://github.com/frederikhors/orm-benchmark).

### ORMs

All package run in no-cache mode.

- [beego/orm](https://github.com/astaxie/beego/tree/master/orm)
- [bun](https://github.com/uptrace/bun)
- [gorm 2](https://github.com/go-gorm/gorm)
- [pg](https://github.com/go-pg/pg)
- [sqlc](https://github.com/kyleconroy/sqlc) (partially)
- [xorm](https://github.com/xormplus/xorm)
- [ent](https://github.com/ent/ent)
- [upper](https://github.com/upper/db)
- [gorp](https://github.com/go-gorp/gorp)
- [godb](https://github.com/samonzeweb/godb)
- [dbr](https://github.com/gocraft/dbr/)
- [pop](https://github.com/gobuffalo/pop)
- [rel](https://github.com/go-rel/rel)
- [reform](https://github.com/go-reform/reform)

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

- go version go1.17.6 linux/amd64
- Intel(R) Core(TM) i5-5200U CPU @ 2.20GHz
- 16GB 1600MHz memory
- PostgreSQL 14.1

From [`results.md`](results.md):

**orm-benchmark -multi=20**

```
  4000 times - Insert
       raw:    11.60s      2899684 ns/op     720 B/op     13 allocs/op
    reform:    11.77s      2941454 ns/op    1796 B/op     51 allocs/op
      gorp:    11.79s      2947211 ns/op    1802 B/op     41 allocs/op
       ent:    11.80s      2950143 ns/op    4230 B/op     99 allocs/op
       dbr:    12.06s      3014671 ns/op    2707 B/op     65 allocs/op
      sqlc:    12.10s      3024474 ns/op    2903 B/op     62 allocs/op
 sqlboiler:    12.11s      3028471 ns/op    1594 B/op     34 allocs/op
       bun:    12.17s      3041652 ns/op    5012 B/op     14 allocs/op
      xorm:    12.37s      3092313 ns/op    2648 B/op     77 allocs/op
        pg:    12.43s      3107421 ns/op     806 B/op     10 allocs/op
      gorm:    12.48s      3120269 ns/op    6982 B/op     91 allocs/op
     beego:    12.54s      3134933 ns/op    2376 B/op     56 allocs/op
      godb:    12.56s      3139276 ns/op    4573 B/op    115 allocs/op
     upper:    13.24s      3309224 ns/op   14330 B/op    682 allocs/op
       pop:    13.45s      3363355 ns/op   10112 B/op    248 allocs/op
       rel:    14.23s      3558389 ns/op    2424 B/op     42 allocs/op

  4000 times - MultiInsert 100 row
       raw:    16.60s      4148822 ns/op  191259 B/op    931 allocs/op
     beego:    16.94s      4235940 ns/op  179617 B/op   2746 allocs/op
    reform:    17.97s      4493121 ns/op  466352 B/op   2747 allocs/op
       ent:    19.29s      4821323 ns/op  412143 B/op   4899 allocs/op
       bun:    19.44s      4859012 ns/op   42409 B/op    219 allocs/op
        pg:    19.97s      4991967 ns/op    3595 B/op    112 allocs/op
      gorm:    21.48s      5370545 ns/op  294868 B/op   4034 allocs/op
       rel:    22.32s      5580700 ns/op  289914 B/op   3363 allocs/op
      godb:    22.55s      5637113 ns/op  259997 B/op   5895 allocs/op
     upper:    26.79s      6697677 ns/op  545938 B/op  19485 allocs/op
       pop:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
      xorm:     doesn't support bulk-insert

  4000 times - Update
       raw:     0.42s       105054 ns/op     750 B/op     13 allocs/op
       ent:     0.52s       128906 ns/op    4558 B/op     99 allocs/op
      sqlc:    11.60s      2899772 ns/op     878 B/op     14 allocs/op
 sqlboiler:    11.61s      2902704 ns/op     901 B/op     17 allocs/op
    reform:    11.71s      2926893 ns/op    1775 B/op     51 allocs/op
     beego:    11.71s      2927845 ns/op    1752 B/op     47 allocs/op
      gorp:    11.80s      2950816 ns/op    1204 B/op     32 allocs/op
       dbr:    11.84s      2959964 ns/op    2651 B/op     57 allocs/op
       pop:    12.01s      3002106 ns/op    6583 B/op    198 allocs/op
      gorm:    12.32s      3079383 ns/op    6604 B/op     81 allocs/op
      xorm:    12.45s      3113117 ns/op    2872 B/op    104 allocs/op
        pg:    12.54s      3134497 ns/op     768 B/op      9 allocs/op
       bun:    12.55s      3137002 ns/op    4730 B/op      5 allocs/op
       rel:    12.61s      3151848 ns/op    2384 B/op     40 allocs/op
      godb:    14.74s      3685113 ns/op    5161 B/op    154 allocs/op
     upper:    16.13s      4032826 ns/op   34681 B/op   1522 allocs/op

  4000 times - Read
       raw:     0.46s       114290 ns/op    2080 B/op     49 allocs/op
      sqlc:     0.46s       115783 ns/op    2190 B/op     51 allocs/op
     beego:     0.48s       120233 ns/op    2088 B/op     75 allocs/op
    reform:     0.50s       125172 ns/op    3213 B/op     85 allocs/op
       pop:     0.53s       133705 ns/op    3553 B/op     70 allocs/op
      gorp:     0.54s       134596 ns/op    3896 B/op    193 allocs/op
        pg:     0.56s       138831 ns/op     872 B/op     20 allocs/op
       bun:     0.62s       155269 ns/op    5493 B/op     21 allocs/op
       dbr:     0.63s       158484 ns/op    2184 B/op     37 allocs/op
       ent:     0.64s       159790 ns/op    5418 B/op    148 allocs/op
 sqlboiler:     0.65s       161999 ns/op     965 B/op     14 allocs/op
       rel:     0.65s       162498 ns/op    1824 B/op     47 allocs/op
      gorm:     0.94s       233844 ns/op    5304 B/op    109 allocs/op
      godb:     1.11s       278508 ns/op    4113 B/op    102 allocs/op
     upper:     1.51s       377915 ns/op    8697 B/op    324 allocs/op
      xorm:     1.63s       407517 ns/op    8755 B/op    215 allocs/op

  4000 times - MultiRead limit 100
    reform:     0.47s       117124 ns/op    2944 B/op     73 allocs/op
      xorm:     0.78s       195170 ns/op    3274 B/op     77 allocs/op
     upper:     1.42s       355861 ns/op    8390 B/op    303 allocs/op
       raw:     1.45s       362607 ns/op   38355 B/op   1037 allocs/op
        pg:     1.57s       392794 ns/op   23834 B/op    629 allocs/op
      sqlc:     1.72s       429824 ns/op   73176 B/op   1250 allocs/op
 sqlboiler:     1.84s       460480 ns/op   58731 B/op   1259 allocs/op
      gorp:     2.02s       504338 ns/op   57403 B/op   1493 allocs/op
       ent:     2.04s       510775 ns/op   76059 B/op   2038 allocs/op
       dbr:     2.42s       604583 ns/op   32416 B/op   1545 allocs/op
     beego:     2.48s       619675 ns/op   55191 B/op   3077 allocs/op
       bun:     2.49s       622581 ns/op   32891 B/op   1118 allocs/op
       pop:     2.56s       640857 ns/op   75816 B/op   1509 allocs/op
      godb:     3.26s       815760 ns/op   75272 B/op   3084 allocs/op
       rel:     3.52s       880888 ns/op   95352 B/op   2250 allocs/op
      gorm:     4.39s      1096908 ns/op   71628 B/op   3877 allocs/op
```

### To-Do List:

- [x] Remove or add prisma again.
- [x] Fix TBDs.
- [x] Add SqlBoiler.
- [x] Add go-reform.
- [x] Clean and restructure codebase.
- [x] Add ~~makefile~~ script to run benchmarks easily, run benchmarks inside of container.
- [x] Add final benchmark results with details.
- [x] Tag release and be ready to compare ORMs performances.

### License

go-orm-benchmarks is licensed under the terms of the **MIT License** (see [LICENSE](LICENSE)).