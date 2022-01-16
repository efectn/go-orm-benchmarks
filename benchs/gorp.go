package benchs

import (
	"log"

	"database/sql"

	"github.com/efectn/go-orm-benchmarks/benchs/ent/model"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	gorpware "gopkg.in/gorp.v1"
)

var gorp *gorpware.DbMap

type ModelGorp struct {
	Id      int64  `db:"id"`
	Name    string `db:"name"`
	Title   string `db:"title"`
	Fax     string `db:"fax"`
	Web     string `db:"web"`
	Age     int    `db:"age"`
	Right   bool   `db:"right"`
	Counter int64  `db:"counter"`
}

func initDB7() {
	db, err := sql.Open("pgx", OrmSource)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	gorp = &gorpware.DbMap{Db: db, Dialect: gorpware.PostgresDialect{}}

	gorp.AddTableWithName(ModelGorp{}, "models_gorp").SetKeys(true, "Id")

	// Run the auto migration tool.
	if _, err = db.Exec("DROP TABLE IF EXISTS models_gorp"); err != nil {
		log.Fatal(err)
	}

	if err := gorp.CreateTablesIfNotExists(); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func NewGorpModel() ModelGorp {
	m := ModelGorp{}
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
	st := NewSuite("gorp")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, GorpInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, EntInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, GorpUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, GorpRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, GorpReadSlice)
	}
}

func GorpInsert(b *B) {
	var m ModelGorp
	wrapExecute(b, func() {
		initDB7()
		m = NewGorpModel()
	})

	for i := 0; i < b.N; i++ {
		if err := gorp.Insert(&m); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}

func GorpInsertMulti(b *B) {
	var ms []ModelGorp
	wrapExecute(b, func() {
		initDB4()
		ms = make([]ModelGorp, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewGorpModel())
		}
	})

	for i := 0; i < b.N; i++ {
		if err := gorp.Insert(&ms); err != nil {
			log.Fatal(err)
			b.FailNow()
		}

	}
}

func GorpUpdate(b *B) {
	var m ModelGorp
	wrapExecute(b, func() {
		initDB7()
		m = NewGorpModel()
		if err := gorp.Insert(&m); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		_, err := gorp.Update(&m)

		if err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}

func GorpRead(b *B) {
	var m ModelGorp
	wrapExecute(b, func() {
		initDB7()
		m = NewGorpModel()
		if err := gorp.Insert(&m); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		//m.Id = 1
		err := gorp.SelectOne(&m, "SELECT * FROM models_gorp")
		if err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}

func GorpReadSlice(b *B) {
	var m ModelGorp
	wrapExecute(b, func() {
		initDB7()
		m = NewGorpModel()
		for i := 0; i < 100; i++ {
			if err := gorp.Insert(&m); err != nil {
				log.Fatal(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var ms []ModelGorp
		gorp.Select(&ms, "select * from models_gorp where id > 0 LIMIT 100")
		_, err := client.Model.Query().Where(model.IDGT(0)).Unique(false).Limit(100).All(ctx)
		if err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}
