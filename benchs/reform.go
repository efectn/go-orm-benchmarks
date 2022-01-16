package benchs

import (
	"database/sql"
	"log"

	r "github.com/efectn/go-orm-benchmarks/benchs/reform"
	_ "github.com/jackc/pgx/v4/stdlib"
	"gopkg.in/reform.v1/dialects/postgresql"

	reformware "gopkg.in/reform.v1"
)

var reform *reformware.DB

func initDB12() {
	db, err := sql.Open("pgx", OrmSource)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	reform = reformware.NewDB(db, postgresql.Dialect, nil)

	// Run the auto migration tool.
	/*if _, err = pop.TX.Exec("DROP TABLE IF EXISTS popmodels"); err != nil {
		log.Fatal(err)
	}*/
	initDB()
}

func NewReformModel() *r.ReformModels {
	m := new(r.ReformModels)
	m.Name = "Orm Benchmark"
	m.Title = "Just a Benchmark for fun"
	m.Fax = "99909990"
	m.Web = "http://blog.milkpod29.me"
	m.Age = 100
	m.Right = true
	m.Counter = 1000

	return m
}

func init() {
	st := NewSuite("reform")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, ReformInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, ReformInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, ReformUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, ReformRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, ReformReadSlice)
	}
}

func ReformInsert(b *B) {
	var m *r.ReformModels
	wrapExecute(b, func() {
		initDB12()
		m = NewReformModel()
	})

	for i := 0; i < b.N; i++ {
		m.ID = 0
		if err := reform.Save(m); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}

func ReformInsertMulti(b *B) {
	var ms []reformware.Struct
	wrapExecute(b, func() {
		initDB()
		ms = make([]reformware.Struct, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewReformModel())
		}
	})

	for i := 0; i < b.N; i++ {
		if err := reform.InsertMulti(ms...); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}

func ReformUpdate(b *B) {
	var m *r.ReformModels
	wrapExecute(b, func() {
		initDB12()
		m = NewReformModel()
		if err := reform.Save(m); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if err := reform.Update(m); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}

func ReformRead(b *B) {
	var m *r.ReformModels
	wrapExecute(b, func() {
		initDB12()
		m = NewReformModel()
		if err := reform.Save(m); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if _, err := reform.FindByPrimaryKeyFrom(r.ReformModelsTable, m.ID); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}

func ReformReadSlice(b *B) {
	var m *r.ReformModels
	wrapExecute(b, func() {
		initDB12()
		m = NewReformModel()
		for i := 0; i < 100; i++ {
			m.ID = 0
			if err := reform.Save(m); err != nil {
				log.Fatal(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		if _, err := reform.SelectAllFrom(r.ReformModelsTable, "WHERE id > 0 LIMIT 100"); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}
