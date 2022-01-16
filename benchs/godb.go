package benchs

import (
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
	godbware "github.com/samonzeweb/godb"
	"github.com/samonzeweb/godb/adapters/postgresql"
)

var godb *godbware.DB

type ModelGodb struct {
	ID      int    `db:"id,key,auto"`
	Name    string `db:"name"`
	Title   string `db:"title"`
	Fax     string `db:"fax"`
	Web     string `db:"web"`
	Age     int    `db:"age"`
	Right   bool   `db:"right"`
	Counter int64  `db:"counter"`
}

func (*ModelGodb) TableName() string {
	return "models_godb"
}

func initDB8() {
	var err error
	godb, err = godbware.Open(postgresql.Adapter, OrmSource)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Run the auto migration tool.
	if _, err = godb.CurrentDB().Exec("DROP TABLE IF EXISTS model_godb"); err != nil {
		log.Fatal(err)
	}
	initDB()
}

func NewGodbModel() ModelGodb {
	m := ModelGodb{}
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
	st := NewSuite("godb")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, GodbInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, GodbInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, GodbUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, GodbRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, GodbReadSlice)
	}
}

func GodbInsert(b *B) {
	var m ModelGodb
	wrapExecute(b, func() {
		initDB8()
		m = NewGodbModel()
	})

	for i := 0; i < b.N; i++ {
		if err := godb.Insert(&m).Do(); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}

func GodbInsertMulti(b *B) {
	var ms []ModelGodb
	wrapExecute(b, func() {
		initDB8()
		ms = make([]ModelGodb, 0, 100)
		for i := 1; i < 101; i++ {
			ms = append(ms, NewGodbModel())
		}
	})

	for i := 0; i < b.N; i++ {
		if err := godb.BulkInsert(&ms).Do(); err != nil {
			log.Fatal(err)
			b.FailNow()
		}

	}
}

func GodbUpdate(b *B) {
	var m ModelGodb
	wrapExecute(b, func() {
		initDB8()
		m = NewGodbModel()
		if err := godb.Insert(&m).Do(); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if err := godb.Update(&m).Do(); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}

func GodbRead(b *B) {
	var m ModelGodb
	wrapExecute(b, func() {
		initDB8()
		m = NewGodbModel()
		if err := godb.Insert(&m).Do(); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if err := godb.Select(&m).Do(); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}

func GodbReadSlice(b *B) {
	var m ModelGodb
	wrapExecute(b, func() {
		initDB8()
		m = NewGodbModel()
		for i := 0; i < 100; i++ {
			if err := godb.Insert(&m).Do(); err != nil {
				log.Fatal(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var ms []ModelGodb
		if err := godb.Select(&ms).Where("id > 0").Limit(100).Do(); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}
