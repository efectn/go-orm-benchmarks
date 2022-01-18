package benchs

import (
	"fmt"
	"log"

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

		if err := pgdb.Ping(ctx); err != nil {
			log.Fatalf("error connecting to postgres: %v", err)
		}
	}
}

func PgInsert(b *B) {
	var m *Model

	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		if _, err := pgdb.Model(m).Insert(); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PgInsertMulti(b *B) {
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

		if _, err := pgdb.Model(&ms).Insert(); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PgUpdate(b *B) {
	var m *Model

	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		if _, err := pgdb.Model(m).Insert(); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if _, err := pgdb.Model(m).WherePK().Update(); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PgRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		if _, err := pgdb.Model(m).Insert(); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if err := pgdb.Model(m).Select(); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PgReadSlice(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.Id = 0
			if _, err := pgdb.Model(m).Insert(); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model
		if err := pgdb.Model(&models).Where("id > ?", 0).Limit(100).Select(); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
