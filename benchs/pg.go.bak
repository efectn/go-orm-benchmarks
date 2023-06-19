package benchs

import (
	"github.com/go-pg/pg/v10"
)

var pgdb *pg.DB

func init() {
	st := NewSuite("pg")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, PgInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, PgInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, PgUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, PgRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, PgReadSlice)

		source := SplitSource()
		pgdb = pg.Connect(&pg.Options{
			Addr:     source["host"] + ":5432",
			User:     source["user"],
			Password: source["password"],
			Database: source["dbname"],
		})

		err := pgdb.Ping(ctx)
		CheckErr(err)
	}
}

func PgInsert(b *B) {
	var m *Model

	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		_, err := pgdb.Model(m).Insert()
		CheckErr(err, b)
	}
}

func PgInsertMulti(b *B) {
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

		_, err := pgdb.Model(&ms).Insert()
		CheckErr(err, b)
	}
}

func PgUpdate(b *B) {
	var m *Model

	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		_, err := pgdb.Model(m).Insert()
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		_, err := pgdb.Model(m).WherePK().Update()
		CheckErr(err, b)
	}
}

func PgRead(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		_, err := pgdb.Model(m).Insert()
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		err := pgdb.Model(m).Select()
		CheckErr(err, b)
	}
}

func PgReadSlice(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.Id = 0
			_, err := pgdb.Model(m).Insert()
			CheckErr(err, b)
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model
		err := pgdb.Model(&models).Where("id > ?", 0).Limit(100).Select()
		CheckErr(err, b)
	}
}
