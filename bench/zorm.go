package bench

import (
	"context"
	"testing"

	zormdb "gitee.com/chunanyong/zorm"
	"github.com/efectn/go-orm-benchmarks/helper"
	_ "github.com/lib/pq"
)

var (
	zormCtx         = context.Background()
	readSliceFinder = zormdb.NewFinder().Append("SELECT * FROM models WHERE id > 0")
	page            = &zormdb.Page{PageNo: 1, PageSize: helper.BatchSize}
)

type Zorm struct {
	helper.ORMInterface
}

func CreateZorm() helper.ORMInterface {
	return &Zorm{}
}

func (zorm *Zorm) Name() string {
	return "zorm"
}

func (zorm *Zorm) Init() error {
	readSliceFinder.InjectionCheck = false
	readSliceFinder.SelectTotalCount = false

	dbDaoConfig := zormdb.DataSourceConfig{
		DSN:                helper.OrmSource,
		DriverName:         "postgres",
		Dialect:            "postgresql",
		MaxOpenConns:       helper.OrmMaxConn,
		MaxIdleConns:       helper.OrmMaxIdle,
		SlowSQLMillis:      -1,
		DisableTransaction: true,
	}
	_, err := zormdb.NewDBDao(&dbDaoConfig)
	if err != nil {
		return err
	}

	return nil
}

func (zorm *Zorm) Close() error {
	return nil
}

func (zorm *Zorm) Insert(b *testing.B) {
	m := NewModel7()
	helper.BenchExec(b, zorm.Name(), "Insert", func() error {
		m.ID = 0
		_, err := zormdb.Insert(zormCtx, m)
		return err
	})
}

func (zorm *Zorm) InsertMulti(b *testing.B) {
	ms := make([]zormdb.IEntityStruct, 0, helper.BatchSize)
	for i := 0; i < helper.BatchSize; i++ {
		ms = append(ms, NewModel7())
	}

	helper.BenchExec(b, zorm.Name(), "InsertMulti", func() error {
		for _, m := range ms {
			m.(*Model7).ID = 0
		}
		_, err := zormdb.InsertSlice(zormCtx, ms)
		return err
	})
}

func (zorm *Zorm) Update(b *testing.B) {
	m := NewModel7()

	_, err := zormdb.Insert(zormCtx, m)
	if err != nil {
		helper.SetError(b, zorm.Name(), "Update", err.Error())
		return
	}
	m.Counter = helper.UpdatedCounter

	helper.BenchExec(b, zorm.Name(), "Update", func() error {
		_, err := zormdb.Update(zormCtx, m)
		return err
	})
}

func (zorm *Zorm) Read(b *testing.B) {
	m := NewModel7()

	_, err := zormdb.Insert(zormCtx, m)
	if err != nil {
		helper.SetError(b, zorm.Name(), "Read", err.Error())
		return
	}
	readFinder := zormdb.NewFinder().Append("SELECT * FROM models WHERE id = ?", m.ID)
	readFinder.InjectionCheck = false

	helper.BenchRows(b, zorm.Name(), "Read", 1, func() (int, error) {
		var result Model7
		found, err := zormdb.QueryRow(zormCtx, readFinder, &result)
		if err != nil {
			return 0, err
		}
		if !found {
			return 0, nil
		}
		return 1, nil
	})
}

func (zorm *Zorm) ReadSlice(b *testing.B) {
	m := NewModel7()
	for i := 0; i < helper.BatchSize; i++ {
		m.ID = 0
		_, err := zormdb.Insert(zormCtx, m)
		if err != nil {
			helper.SetError(b, zorm.Name(), "ReadSlice", err.Error())
			return
		}
	}

	helper.BenchRows(b, zorm.Name(), "ReadSlice", helper.BatchSize, func() (int, error) {
		var models []Model7
		err := zormdb.Query(zormCtx, readSliceFinder, &models, page)
		if err != nil {
			return 0, err
		}
		return len(models), nil
	})
}
