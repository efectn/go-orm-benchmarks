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
- [pgx](https://github.com/jackc/pgx)
- [zorm](https://gitee.com/chunanyong/zorm)
- [gen](https://gorm.io/gen/index.html)
- [queryx](https://github.com/swiftcarrot/queryx)

See [`go.mod`](go.mod) for their latest versions.

### Run

```shell
# install
go install github.com/efectn/go-orm-benchmarks@latest
# all
go-orm-benchmarks -orm=all
# portion
go-orm-benchmarks -orm=gorm
go-orm-benchmarks -orm=pg
go-orm-benchmarks -orm=bun
# ... and so on...
```

**_Note: Also, you can run `./run_benchmarks.sh` and you can get output like results.md format._**

### Results
Look at [`results.md`](results.md) to see detailed benchmark results.

**Note:** All results are automatically generated by [Github Actions](https://github.com/features/actions). Benchmark results may sometimes be wrong.

### License

go-orm-benchmarks is licensed under the terms of the **MIT License** (see [LICENSE](LICENSE)).
