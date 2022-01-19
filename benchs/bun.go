package benchs

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var (
	ctx   = context.Background()
	bundb *bun.DB
)

func init() {
	st := NewSuite("bun")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, BunInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, BunInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, BunUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, BunRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, BunReadSlice)

		sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(ConvertSourceToDSN())))
		sqldb.SetMaxOpenConns(OrmMaxConn)
		sqldb.SetMaxIdleConns(OrmMaxIdle)

		bundb = bun.NewDB(sqldb, pgdialect.New())
	}
}

func BunInsert(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		_, err := bundb.NewInsert().Model(m).Exec(ctx)
		CheckErr(err, b)
	}
}

func BunInsertMulti(b *B) {
	var ms []*Model
	WrapExecute(b, func() {
		InitDB()
		ms = make([]*Model, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewModel())
		}
	})

	for i := 0; i < b.N; i++ {
		for _, m := range ms {
			m.Id = 0
		}

		_, err := bundb.NewInsert().Model(&ms).Exec(ctx)
		CheckErr(err, b)
	}
}

func BunUpdate(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()

		_, err := bundb.NewInsert().Model(m).Exec(ctx)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		_, err := bundb.NewUpdate().Model(m).WherePK().Exec(ctx)
		CheckErr(err, b)
	}
}

func BunRead(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()

		_, err := bundb.NewInsert().Model(m).Exec(ctx)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		err := bundb.NewSelect().Model(m).Scan(ctx)
		CheckErr(err, b)
	}
}

func BunReadSlice(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.Id = 0
			_, err := bundb.NewInsert().Model(m).Exec(ctx)
			CheckErr(err, b)
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model
		err := bundb.NewSelect().
			Model(&models).
			Where("id > ?", 0).
			Limit(100).
			Scan(ctx)
		CheckErr(err, b)
	}
}
