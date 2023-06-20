package bench

import (
	"context"
	"github.com/efectn/go-orm-benchmarks/helper"
	"sync"
	"testing"

	zormdb "gitee.com/chunanyong/zorm"
	_ "github.com/lib/pq"
)

var (
	zormCtx         = context.Background()
	readFinder      = zormdb.NewFinder().Append("SELECT * FROM models WHERE id = 1")
	readSliceFinder = zormdb.NewFinder().Append("SELECT * FROM models WHERE id > 0")
	page            = &zormdb.Page{PageNo: 1, PageSize: 100}
)

type Zorm struct {
	helper.ORMInterface
	mu         sync.Mutex
	iterations int // Same as b.N, just to customize it
}

func CreateZorm(iterations int) helper.ORMInterface {
	zorm := &Zorm{
		iterations: iterations,
	}

	return zorm
}

func (zorm *Zorm) Name() string {
	return "zorm"
}

func (zorm *Zorm) Init() error {
	readFinder.InjectionCheck = false
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

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.ID = 0
		_, err := zormdb.Insert(zormCtx, m)
		if err != nil {
			helper.SetError(b, zorm.Name(), "insert", err.Error())
		}
	}
}

func (zorm *Zorm) InsertMulti(b *testing.B) {
	ms := make([]zormdb.IEntityStruct, 0, 100)
	for i := 0; i < 100; i++ {
		ms = append(ms, NewModel7())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, m := range ms {
			m7, _ := m.(*Model7)
			m7.ID = 0
		}
		_, err := zormdb.InsertSlice(zormCtx, ms)
		if err != nil {
			helper.SetError(b, zorm.Name(), "insert_multi", err.Error())
		}
	}
}

func (zorm *Zorm) Update(b *testing.B) {
	m := NewModel7()

	_, err := zormdb.Insert(zormCtx, m)
	if err != nil {
		helper.SetError(b, zorm.Name(), "update", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := zormdb.Update(zormCtx, m)
		if err != nil {
			helper.SetError(b, zorm.Name(), "update", err.Error())
		}
	}
}

func (zorm *Zorm) Read(b *testing.B) {
	m := NewModel7()

	_, err := zormdb.Insert(zormCtx, m)
	if err != nil {
		helper.SetError(b, zorm.Name(), "read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := zormdb.QueryRow(zormCtx, readFinder, m)
		if err != nil {
			helper.SetError(b, zorm.Name(), "read", err.Error())
		}
	}
}

func (zorm *Zorm) ReadSlice(b *testing.B) {
	m := NewModel7()
	for i := 0; i < 100; i++ {
		m.ID = 0
		_, err := zormdb.Insert(zormCtx, m)
		if err != nil {
			helper.SetError(b, zorm.Name(), "read_slice", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var models []Model7
		err := zormdb.Query(zormCtx, readSliceFinder, &models, page)
		if err != nil {
			helper.SetError(b, zorm.Name(), "read_slice", err.Error())
		}
	}
}
