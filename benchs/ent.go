package benchs

import (
	"context"
	"log"

	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/efectn/orm-benchmark/benchs/ent"
	"github.com/efectn/orm-benchmark/benchs/ent/model"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

var client *ent.Client

func initDB4() {
	db, err := sql.Open("pgx", OrmSource)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)

	// Assign to client
	client = ent.NewClient(ent.Driver(drv))

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func NewEntModel() ent.Model {
	m := ent.Model{}
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
	st := NewSuite("ent")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, EntInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, EntInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, EntUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, EntRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, EntReadSlice)
	}
}

func EntInsert(b *B) {
	var m ent.Model
	wrapExecute(b, func() {
		initDB4()
		m = NewEntModel()
	})

	for i := 0; i < b.N; i++ {
		_, err := client.Model.Create().
			SetName(m.Name).
			SetTitle(m.Title).
			SetFax(m.Fax).
			SetWeb(m.Web).
			SetAge(m.Age).
			SetRight(m.Right).
			SetCounter(m.Counter).
			Save(ctx)

		if err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}

func EntInsertMulti(b *B) {
	var ms []ent.Model
	wrapExecute(b, func() {
		initDB4()
		ms = make([]ent.Model, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewEntModel())
		}
	})

	bulk := make([]*ent.ModelCreate, len(ms))
	for i, m := range ms {
		bulk[i] = client.Model.Create().
			SetName(m.Name).
			SetTitle(m.Title).
			SetFax(m.Fax).
			SetWeb(m.Web).
			SetAge(m.Age).
			SetRight(m.Right).
			SetCounter(m.Counter)
	}

	if _, err := client.Model.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatal(err)
		b.FailNow()
	}
}

func EntUpdate(b *B) {
	var m ent.Model
	wrapExecute(b, func() {
		initDB4()
		m = NewEntModel()
		_, err := client.Model.Create().
			SetName(m.Name).
			SetTitle(m.Title).
			SetFax(m.Fax).
			SetWeb(m.Web).
			SetAge(m.Age).
			SetRight(m.Right).
			SetCounter(m.Counter).
			Save(ctx)

		if err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		_, err := client.Model.Update().
			Where(model.IDEQ(m.ID)).
			SetName(m.Name).
			SetTitle(m.Title).
			SetFax(m.Fax).
			SetWeb(m.Web).
			SetAge(m.Age).
			SetRight(m.Right).
			SetCounter(m.Counter).
			Save(ctx)

		if err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}

func EntRead(b *B) {
	var m ent.Model
	wrapExecute(b, func() {
		initDB4()
		m = NewEntModel()
		_, err := client.Model.Create().
			SetName(m.Name).
			SetTitle(m.Title).
			SetFax(m.Fax).
			SetWeb(m.Web).
			SetAge(m.Age).
			SetRight(m.Right).
			SetCounter(m.Counter).
			Save(ctx)

		if err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		m.ID = 1
		_, err := client.Model.Query().Where(model.IDEQ(m.ID)).First(ctx)
		if err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}

func EntReadSlice(b *B) {
	var m ent.Model
	wrapExecute(b, func() {
		initDB4()
		m = NewEntModel()
		for i := 0; i < 100; i++ {
			_, err := client.Model.Create().
				SetName(m.Name).
				SetTitle(m.Title).
				SetFax(m.Fax).
				SetWeb(m.Web).
				SetAge(m.Age).
				SetRight(m.Right).
				SetCounter(m.Counter).
				Save(ctx)
			if err != nil {
				log.Fatal(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		_, err := client.Model.Query().Where(model.IDGT(0)).Limit(100).All(ctx)
		if err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}
