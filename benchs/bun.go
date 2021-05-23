package benchs

import (
	"context"
	"database/sql"
	"fmt"

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

		dsn := "postgres://postgres:@localhost:5432/test?sslmode=disable"
		sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
		sqldb.SetMaxOpenConns(OrmMaxConn)
		sqldb.SetMaxIdleConns(OrmMaxIdle)
		bundb = bun.NewDB(sqldb, pgdialect.New())
	}
}

func BunInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		if _, err := bundb.NewInsert().Model(m).Exec(ctx); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func BunInsertMulti(b *B) {
	var ms []*Model
	wrapExecute(b, func() {
		initDB()
		ms = make([]*Model, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewModel())
		}
	})

	for i := 0; i < b.N; i++ {
		for _, m := range ms {
			m.Id = 0
		}
		if _, err := bundb.NewInsert().Model(&ms).Exec(ctx); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func BunUpdate(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		if _, err := bundb.NewInsert().Model(m).Exec(ctx); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if _, err := bundb.NewUpdate().Model(m).WherePK().Exec(ctx); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func BunRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		if _, err := bundb.NewInsert().Model(m).Exec(ctx); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if err := bundb.NewSelect().Model(m).Scan(ctx); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func BunReadSlice(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.Id = 0
			if _, err := bundb.NewInsert().Model(m).Exec(ctx); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model
		if err := bundb.NewSelect().
			Model(&models).
			Where("id > ?", 0).
			Limit(100).
			Scan(ctx); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
