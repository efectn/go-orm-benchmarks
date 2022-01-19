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
  raw_stmt:    12.01s      3002522 ns/op     719 B/op     14 allocs/op
 beego_orm:    12.09s      3022997 ns/op    2408 B/op     56 allocs/op
    reform:    12.23s      3056810 ns/op    2602 B/op     58 allocs/op
       ent:    12.24s      3059942 ns/op    4230 B/op     99 allocs/op
      gorp:    12.25s      3063563 ns/op    1788 B/op     39 allocs/op
        pg:    12.35s      3087888 ns/op    1068 B/op     10 allocs/op
 gorm_prep:    12.35s      3088644 ns/op    5300 B/op     69 allocs/op
      sqlc:    12.38s      3094607 ns/op    2903 B/op     62 allocs/op
       dbr:    12.41s      3101695 ns/op    2688 B/op     65 allocs/op
       bun:    12.46s      3115317 ns/op    5010 B/op     14 allocs/op
      gorm:    12.51s      3126918 ns/op    6980 B/op     91 allocs/op
      godb:    12.95s      3237834 ns/op    4586 B/op    115 allocs/op
      xorm:    13.01s      3253057 ns/op    2854 B/op     78 allocs/op
       rel:    13.32s      3330119 ns/op    2424 B/op     42 allocs/op
     upper:    13.52s      3380797 ns/op   14414 B/op    682 allocs/op
       raw:    13.60s      3399530 ns/op     720 B/op     13 allocs/op
       pop:    13.98s      3494398 ns/op   15136 B/op    620 allocs/op

  4000 times - MultiInsert 100 row
  raw_stmt:    16.63s      4157382 ns/op  187179 B/op    931 allocs/op
 beego_orm:    17.52s      4379926 ns/op  179569 B/op   2746 allocs/op
       raw:    17.67s      4417005 ns/op  191258 B/op    931 allocs/op
    reform:    17.67s      4417369 ns/op  467370 B/op   2747 allocs/op
 gorm_prep:    18.75s      4686265 ns/op  257058 B/op   2488 allocs/op
      gorp:    19.30s      4825682 ns/op  412140 B/op   4899 allocs/op
       bun:    19.83s      4957073 ns/op   42414 B/op    219 allocs/op
       ent:    20.16s      5039215 ns/op  412142 B/op   4899 allocs/op
        pg:    20.21s      5051404 ns/op    3839 B/op    112 allocs/op
      gorm:    21.09s      5273030 ns/op  294868 B/op   4034 allocs/op
      xorm:    21.63s      5408257 ns/op  266568 B/op   5822 allocs/op
       rel:    21.76s      5440309 ns/op  289913 B/op   3363 allocs/op
      godb:    22.84s      5710840 ns/op  260025 B/op   5895 allocs/op
     upper:     TBD
       dbr:     doesn't support bulk-insert
      sqlc:     TBD
       pop:     doesn't support bulk-insert

  4000 times - Update
  raw_stmt:     0.42s       105004 ns/op     766 B/op     14 allocs/op
       raw:     0.42s       105708 ns/op     750 B/op     13 allocs/op
       ent:     0.51s       128295 ns/op    4556 B/op     99 allocs/op
    reform:    12.09s      3021676 ns/op    1805 B/op     51 allocs/op
 beego_orm:    12.12s      3030522 ns/op    1800 B/op     47 allocs/op
 gorm_prep:    12.27s      3068128 ns/op    5117 B/op     59 allocs/op
      sqlc:    12.29s      3072214 ns/op     878 B/op     14 allocs/op
       dbr:    12.36s      3089163 ns/op    2635 B/op     56 allocs/op
      gorp:    12.36s      3089352 ns/op    1204 B/op     32 allocs/op
        pg:    12.43s      3106954 ns/op     768 B/op      9 allocs/op
       pop:    12.58s      3144366 ns/op   12917 B/op    663 allocs/op
       bun:    12.74s      3184031 ns/op    4730 B/op      5 allocs/op
      gorm:    12.74s      3184901 ns/op    6604 B/op     81 allocs/op
      xorm:    13.04s      3260125 ns/op    2889 B/op    105 allocs/op
       rel:    13.16s      3290694 ns/op    2504 B/op     40 allocs/op
      godb:    13.17s      3291765 ns/op    5161 B/op    154 allocs/op
     upper:    16.35s      4086557 ns/op   34839 B/op   1522 allocs/op

  4000 times - Read
      sqlc:     0.44s       110184 ns/op    2190 B/op     51 allocs/op
       raw:     0.45s       111642 ns/op    2080 B/op     49 allocs/op
  raw_stmt:     0.46s       114036 ns/op    2112 B/op     50 allocs/op
 beego_orm:     0.48s       120056 ns/op    2104 B/op     75 allocs/op
    reform:     0.50s       123865 ns/op    3467 B/op     85 allocs/op
      gorp:     0.54s       133938 ns/op    3450 B/op    137 allocs/op
 gorm_prep:     0.56s       140534 ns/op    4972 B/op     98 allocs/op
        pg:     0.57s       143308 ns/op     872 B/op     20 allocs/op
       bun:     0.60s       150881 ns/op    5493 B/op     21 allocs/op
       ent:     0.61s       151551 ns/op    5418 B/op    148 allocs/op
       dbr:     0.63s       156274 ns/op    2184 B/op     37 allocs/op
       rel:     0.64s       161060 ns/op    1824 B/op     47 allocs/op
       pop:     0.73s       182434 ns/op    8711 B/op    442 allocs/op
      gorm:     0.83s       208299 ns/op    5303 B/op    109 allocs/op
      godb:     1.16s       289771 ns/op    4161 B/op    102 allocs/op
     upper:     1.39s       347123 ns/op    8712 B/op    324 allocs/op
      xorm:     1.46s       364538 ns/op    8780 B/op    216 allocs/op

  4000 times - MultiRead limit 100
       raw:     1.26s       316072 ns/op   38355 B/op   1037 allocs/op
  raw_stmt:     1.28s       319219 ns/op   38387 B/op   1038 allocs/op
     upper:     1.29s       322579 ns/op    8401 B/op    303 allocs/op
    reform:     1.41s       353159 ns/op   56157 B/op   1270 allocs/op
      sqlc:     1.48s       370139 ns/op   73176 B/op   1250 allocs/op
        pg:     1.60s       401090 ns/op   23308 B/op    629 allocs/op
       ent:     1.78s       443760 ns/op   76063 B/op   2038 allocs/op
 beego_orm:     2.09s       522492 ns/op   55228 B/op   3077 allocs/op
       bun:     2.12s       529040 ns/op   32885 B/op   1118 allocs/op
       dbr:     2.16s       540184 ns/op   54872 B/op   1545 allocs/op
       pop:     2.39s       596267 ns/op   80110 B/op   1850 allocs/op
       rel:     3.15s       787805 ns/op   95352 B/op   2250 allocs/op
 gorm_prep:     3.25s       812247 ns/op   71236 B/op   3577 allocs/op
      godb:     3.46s       865254 ns/op   97722 B/op   3084 allocs/op
      gorp:     3.70s       924067 ns/op  165467 B/op   3575 allocs/op
      gorm:     3.81s       951788 ns/op   71627 B/op   3877 allocs/op
      xorm:     doesn't work
```

### To-Do List:

- [x] Remove or add prisma again.
- [ ] Fix TBDs.
- [ ] Add SqlBoiler.
- [x] Add go-reform.
- [x] Clean and restructure codebase.
- [x] Add ~~makefile~~ script to run benchmarks easily, run benchmarks inside of container.
- [ ] Add final benchmark results with details.
- [ ] Tag release and be ready to compare ORMs performances.

### License

go-orm-benchmarks is licensed under the terms of the **MIT License** (see [LICENSE](LICENSE)).