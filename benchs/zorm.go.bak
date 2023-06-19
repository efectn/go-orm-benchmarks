package benchs

import (
	"context"

	"gitee.com/chunanyong/zorm"
	_ "github.com/lib/pq"
)

var (
	zormCtx         = context.Background()
	readFinder      = zorm.NewFinder().Append("SELECT * FROM models WHERE id = 1")
	readSliceFinder = zorm.NewFinder().Append("SELECT * FROM models WHERE id > 0")
	page            = &zorm.Page{PageNo: 1, PageSize: 100}
)

func init() {
	st := NewSuite("zorm")
	readFinder.InjectionCheck = false
	readSliceFinder.InjectionCheck = false
	readSliceFinder.SelectTotalCount = false
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, ZormInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, ZormInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, ZormUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, ZormRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, ZormReadSlice)

		dbDaoConfig := zorm.DataSourceConfig{
			DSN:                OrmSource,
			DriverName:         "postgres",
			Dialect:            "postgresql",
			MaxOpenConns:       OrmMaxConn,
			MaxIdleConns:       OrmMaxIdle,
			SlowSQLMillis:      -1,
			DisableTransaction: true,
		}
		_, err := zorm.NewDBDao(&dbDaoConfig)
		CheckErr(err)
	}
}

func ZormInsert(b *B) {
	var m *Model7
	WrapExecute(b, func() {
		InitDB()
		m = NewModel7()
	})

	for i := 0; i < b.N; i++ {
		m.ID = 0
		_, err := zorm.Insert(zormCtx, m)
		CheckErr(err, b)
	}
}

func ZormInsertMulti(b *B) {
	var ms []zorm.IEntityStruct
	WrapExecute(b, func() {
		InitDB()
		ms = make([]zorm.IEntityStruct, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewModel7())
		}
	})

	for i := 0; i < b.N; i++ {
		for _, m := range ms {
			m7, _ := m.(*Model7)
			m7.ID = 0
		}
		_, err := zorm.InsertSlice(zormCtx, ms)
		CheckErr(err, b)
	}
}

func ZormUpdate(b *B) {
	var m *Model7
	WrapExecute(b, func() {
		InitDB()
		m = NewModel7()
		_, err := zorm.Insert(zormCtx, m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		_, err := zorm.Update(zormCtx, m)
		CheckErr(err, b)
	}
}

func ZormRead(b *B) {
	var m *Model7
	WrapExecute(b, func() {
		InitDB()
		m = NewModel7()
		_, err := zorm.Insert(zormCtx, m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		_, err := zorm.QueryRow(zormCtx, readFinder, m)
		CheckErr(err, b)
	}
}

func ZormReadSlice(b *B) {
	var m *Model7
	WrapExecute(b, func() {
		InitDB()
		m = NewModel7()
		for i := 0; i < 100; i++ {
			m.ID = 0
			_, err := zorm.Insert(zormCtx, m)
			CheckErr(err, b)
		}
	})

	for i := 0; i < b.N; i++ {
		var models []Model7
		err := zorm.Query(zormCtx, readSliceFinder, &models, page)
		CheckErr(err, b)
	}
}
