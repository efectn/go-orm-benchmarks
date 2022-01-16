package benchs

import (
	"fmt"
	"log"

	popware "github.com/gobuffalo/pop/v6"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var pop *popware.Connection

type PopModel struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Title   string `db:"title"`
	Fax     string `db:"fax"`
	Web     string `db:"web"`
	Age     int    `db:"age"`
	Right   bool   `db:"right"`
	Counter int64  `db:"counter"`
}

func NewPopModel() *PopModel {
	m := new(PopModel)
	m.Name = "Orm Benchmark"
	m.Title = "Just a Benchmark for fun"
	m.Fax = "99909990"
	m.Web = "http://blog.milkpod29.me"
	m.Age = 100
	m.Right = true
	m.Counter = 1000

	return m
}

func initDB10() {
	var err error
	pop, err = popware.NewConnection(&popware.ConnectionDetails{
		URL: "postgres://postgres:postgres@localhost:5432/test?sslmode=disable",
	})
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	if err = pop.Open(); err != nil {
		log.Fatal(err)
	}

	// Run the auto migration tool.
	/*if _, err = pop.TX.Exec("DROP TABLE IF EXISTS popmodels"); err != nil {
		log.Fatal(err)
	}*/
	initDB()
}

func init() {
	st := NewSuite("pop")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, PopInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, PopInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, PopUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, PopRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, PopReadSlice)
	}
}

func PopInsert(b *B) {
	var m *PopModel
	wrapExecute(b, func() {
		initDB10()
		m = NewPopModel()
	})

	for i := 0; i < b.N; i++ {
		//m.Id = 0
		if err := pop.Create(m); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}

func PopInsertMulti(b *B) {
	panic(fmt.Errorf("doesn't support bulk-insert"))
}

func PopUpdate(b *B) {
	var m *PopModel
	wrapExecute(b, func() {
		initDB10()
		m = NewPopModel()
		if err := pop.Create(m); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if err := pop.Update(m); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}

func PopRead(b *B) {
	var m *PopModel
	wrapExecute(b, func() {
		initDB10()
		m = NewPopModel()
		if err := pop.Create(m); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if err := pop.First(m); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}

func PopReadSlice(b *B) {
	var m *PopModel
	wrapExecute(b, func() {
		initDB10()
		m = NewPopModel()
		for i := 0; i < 100; i++ {
			if err := pop.Create(m); err != nil {
				log.Fatal(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var ms []PopModel
		if err := pop.Where("id > 0").Limit(100).All(&ms); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}
