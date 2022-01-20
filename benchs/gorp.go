package benchs

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	gorpware "gopkg.in/gorp.v1"
)

var gorp *gorpware.DbMap

func init() {
	st := NewSuite("gorp")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, GorpInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, GorpInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, GorpUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, GorpRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, GorpReadSlice)

		db, err := sql.Open("pgx", OrmSource)
		CheckErr(err)

		gorp = &gorpware.DbMap{Db: db, Dialect: gorpware.PostgresDialect{}}

		gorp.AddTableWithName(Model{}, "models").SetKeys(true, "Id")
	}
}

func GorpInsert(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		err := gorp.Insert(m)
		CheckErr(err, b)
	}
}

func GorpInsertMulti(b *B) {
	panic(fmt.Errorf("doesn't support bulk-insert"))
}

func GorpUpdate(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		err := gorp.Insert(m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		_, err := gorp.Update(m)

		CheckErr(err, b)
	}
}

func GorpRead(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		err := gorp.Insert(m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		err := gorp.SelectOne(m, "SELECT * FROM models")
		CheckErr(err, b)
	}
}

func GorpReadSlice(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			err := gorp.Insert(m)
			CheckErr(err, b)
		}
	})

	for i := 0; i < b.N; i++ {
		var ms []*Model
		_, err := gorp.Select(&ms, "select * from models where id > 0 LIMIT 100")
		CheckErr(err, b)
	}
}
