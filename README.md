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

- **orm-benchmark -multi=20**

```
  4000 times - Insert
  raw_stmt:     0.38s        94280 ns/op     718 B/op     14 allocs/op
       raw:     0.39s        96719 ns/op     718 B/op     13 allocs/op
 beego_orm:     0.48s       118994 ns/op    2411 B/op     56 allocs/op
       bun:     0.57s       142285 ns/op     918 B/op     12 allocs/op
        pg:     0.58s       145496 ns/op    1235 B/op     12 allocs/op
      gorm:     0.70s       175294 ns/op    6665 B/op     88 allocs/op
      xorm:     0.76s       189533 ns/op    3032 B/op     94 allocs/op

  4000 times - MultiInsert 100 row
       raw:     4.59s      1147385 ns/op  135155 B/op    916 allocs/op
  raw_stmt:     4.59s      1148137 ns/op  131076 B/op    916 allocs/op
 beego_orm:     5.50s      1375637 ns/op  179962 B/op   2747 allocs/op
       bun:     6.18s      1544648 ns/op    4265 B/op    214 allocs/op
        pg:     7.01s      1753495 ns/op    5039 B/op    114 allocs/op
      gorm:     9.52s      2379219 ns/op  293956 B/op   3729 allocs/op
      xorm:    11.66s      2915478 ns/op  286140 B/op   7422 allocs/op

  4000 times - Update
  raw_stmt:     0.26s        65781 ns/op     773 B/op     14 allocs/op
       raw:     0.31s        77209 ns/op     757 B/op     13 allocs/op
 beego_orm:     0.43s       107064 ns/op    1802 B/op     47 allocs/op
       bun:     0.56s       139839 ns/op     589 B/op      4 allocs/op
        pg:     0.60s       149608 ns/op     896 B/op     11 allocs/op
      gorm:     0.74s       185970 ns/op    6604 B/op     81 allocs/op
      xorm:     0.81s       203240 ns/op    2994 B/op    119 allocs/op

  4000 times - Read
       raw:     0.33s        81671 ns/op    2081 B/op     49 allocs/op
  raw_stmt:     0.34s        85847 ns/op    2112 B/op     50 allocs/op
 beego_orm:     0.38s        94777 ns/op    2106 B/op     75 allocs/op
        pg:     0.42s       106148 ns/op    1526 B/op     22 allocs/op
       bun:     0.43s       106904 ns/op    1319 B/op     18 allocs/op
      gorm:     0.65s       162221 ns/op    5240 B/op    108 allocs/op
      xorm:     1.13s       281738 ns/op    8326 B/op    237 allocs/op

  4000 times - MultiRead limit 100
       raw:     1.52s       380351 ns/op   38356 B/op   1037 allocs/op
  raw_stmt:     1.54s       385541 ns/op   38388 B/op   1038 allocs/op
        pg:     1.86s       465468 ns/op   24045 B/op    631 allocs/op
       bun:     2.58s       645354 ns/op   30009 B/op   1122 allocs/op
 beego_orm:     2.93s       732028 ns/op   55280 B/op   3077 allocs/op
      gorm:     4.97s      1241831 ns/op   71628 B/op   3877 allocs/op
      xorm:     doesn't work
```
