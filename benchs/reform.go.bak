package benchs

import (
	"database/sql"

	r "github.com/efectn/go-orm-benchmarks/benchs/reform"
	_ "github.com/jackc/pgx/v4/stdlib"
	"gopkg.in/reform.v1/dialects/postgresql"

	reformware "gopkg.in/reform.v1"
)

var reform *reformware.DB

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

		db, err := sql.Open("pgx", OrmSource)
		CheckErr(err)

		reform = reformware.NewDB(db, postgresql.Dialect, nil)
	}
}

func ReformInsert(b *B) {
	var m *r.ReformModels
	WrapExecute(b, func() {
		InitDB()
		m = NewReformModel()
	})

	for i := 0; i < b.N; i++ {
		err := reform.Save(m)
		CheckErr(err, b)
	}
}

func ReformInsertMulti(b *B) {
	var ms []reformware.Struct
	WrapExecute(b, func() {
		InitDB()
		ms = make([]reformware.Struct, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewReformModel())
		}
	})

	for i := 0; i < b.N; i++ {
		err := reform.InsertMulti(ms...)
		CheckErr(err, b)
	}
}

func ReformUpdate(b *B) {
	var m *r.ReformModels
	WrapExecute(b, func() {
		InitDB()
		m = NewReformModel()
		err := reform.Save(m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		err := reform.Update(m)
		CheckErr(err, b)
	}
}

func ReformRead(b *B) {
	var m *r.ReformModels
	WrapExecute(b, func() {
		InitDB()
		m = NewReformModel()
		err := reform.Save(m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		_, err := reform.FindByPrimaryKeyFrom(r.ReformModelsTable, m.ID)
		CheckErr(err, b)
	}
}

func ReformReadSlice(b *B) {
	var m *r.ReformModels
	WrapExecute(b, func() {
		InitDB()
		m = NewReformModel()
		for i := 0; i < 100; i++ {
			err := reform.Save(m)
			CheckErr(err, b)
		}
	})

	for i := 0; i < b.N; i++ {
		_, err := reform.SelectAllFrom(r.ReformModelsTable, "WHERE id > 0 LIMIT 100")
		CheckErr(err, b)
	}
}
