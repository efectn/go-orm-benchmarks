package benchs

import (
	"fmt"
	"log"

	db "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

var upper db.Session

func initDB6() {
	var err error

	source := SplitSource()
	upper, err = postgresql.Open(postgresql.ConnectionURL{
		Host:     source["host"] + ":5432",
		User:     source["user"],
		Password: source["password"],
		Database: source["dbname"],
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := upper.Ping(); err != nil {
		log.Fatalf("error connecting to postgres: %v", err)
	}

	initDB()
}

func closeUpper() {
	if err := upper.Close(); err != nil {
		log.Fatal(err)
	}
}

func NewUpperModel() ModelUpper {
	m := ModelUpper{}
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
	st := NewSuite("upper")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, UpperInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, UpperInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, UpperUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, UpperRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, UpperReadSlice)
	}
}

func UpperInsert(b *B) {
	var m ModelUpper
	wrapExecute(b, func() {
		initDB6()
		m = NewUpperModel()
	})

	for i := 1; i < b.N+1; i++ {
		_, err := upper.Collection("models_upper").Insert(&m)

		if err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
	closeUpper()
}

func UpperInsertMulti(b *B) {
	panic(fmt.Errorf("TBD"))
}

func UpperUpdate(b *B) {
	var m ModelUpper
	wrapExecute(b, func() {
		initDB6()
		m = NewUpperModel()

		if err := upper.Collection("models_upper").InsertReturning(&m); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		err := upper.Collection("models_upper").UpdateReturning(&m)

		if err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
	closeUpper()
}

func UpperRead(b *B) {
	var m ModelUpper
	wrapExecute(b, func() {
		initDB6()
		m = NewUpperModel()

		if err := upper.Collection("models_upper").InsertReturning(&m); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		err := upper.SQL().SelectFrom("models_upper").Where("id = ?", m.Id).One(&m)
		if err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
	closeUpper()
}

func UpperReadSlice(b *B) {
	var m ModelUpper
	wrapExecute(b, func() {
		initDB6()
		m = NewUpperModel()

		if err := upper.Collection("models_upper").InsertReturning(&m); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		var ms []ModelUpper
		err := upper.SQL().SelectFrom("models_upper").Where("id > ?", m.Id).Limit(100).All(&ms)
		if err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
	closeUpper()
}
