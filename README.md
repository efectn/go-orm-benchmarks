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


**orm-benchmark -multi=50**

```
 10000 times - Insert
       raw:    18.82s      1882463 ns/op     713 B/op     13 allocs/op
  raw_stmt:    18.87s      1886864 ns/op     720 B/op     14 allocs/op
       bun:    19.60s      1959859 ns/op     907 B/op     13 allocs/op
      gorm:    19.67s      1967267 ns/op    6655 B/op     88 allocs/op
        pg:    19.74s      1974130 ns/op    1055 B/op     12 allocs/op
 beego_orm:    19.85s      1984567 ns/op    2409 B/op     56 allocs/op
      xorm:    25.89s      2589262 ns/op    3035 B/op     94 allocs/op

 10000 times - MultiInsert 100 row
  raw_stmt:    48.24s      4824085 ns/op  131034 B/op    916 allocs/op
       raw:    48.92s      4891580 ns/op  135114 B/op    916 allocs/op
 beego_orm:    56.10s      5610480 ns/op  179821 B/op   2746 allocs/op
       bun:    59.21s      5921031 ns/op    4245 B/op    214 allocs/op
        pg:    60.79s      6078986 ns/op    4396 B/op    114 allocs/op
      gorm:    72.20s      7219802 ns/op  293984 B/op   3729 allocs/op
      xorm:    81.02s      8102211 ns/op  285707 B/op   7421 allocs/op

 10000 times - Update
  raw_stmt:     6.44s       643858 ns/op     768 B/op     14 allocs/op
       raw:     6.56s       656429 ns/op     753 B/op     13 allocs/op
 beego_orm:    19.42s      1942096 ns/op    1801 B/op     47 allocs/op
       bun:    19.54s      1953862 ns/op     585 B/op      4 allocs/op
      gorm:    19.72s      1971955 ns/op    6604 B/op     81 allocs/op
        pg:    19.76s      1975545 ns/op    1121 B/op     11 allocs/op
      xorm:    26.50s      2649894 ns/op    2993 B/op    119 allocs/op

 10000 times - Read
 beego_orm:     6.56s       655910 ns/op    2105 B/op     75 allocs/op
  raw_stmt:     6.57s       657027 ns/op    2110 B/op     50 allocs/op
       raw:     6.61s       661139 ns/op    2079 B/op     49 allocs/op
       bun:     6.87s       687084 ns/op    1307 B/op     18 allocs/op
        pg:     7.09s       708644 ns/op    1225 B/op     22 allocs/op
      gorm:     7.29s       728913 ns/op    5242 B/op    108 allocs/op
      xorm:    14.00s      1400469 ns/op    8318 B/op    237 allocs/op

 10000 times - MultiRead limit 100
       raw:     8.80s       880235 ns/op   38357 B/op   1037 allocs/op
  raw_stmt:     8.80s       880457 ns/op   38389 B/op   1038 allocs/op
        pg:     9.71s       970842 ns/op   24755 B/op    631 allocs/op
       bun:    10.13s      1012692 ns/op   28802 B/op   1116 allocs/op
 beego_orm:    10.52s      1051815 ns/op   55250 B/op   3077 allocs/op
      gorm:    13.09s      1309194 ns/op   71629 B/op   3877 allocs/op
      xorm:     doesn't work
```
