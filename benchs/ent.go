package benchs

import (
	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/efectn/go-orm-benchmarks/benchs/ent"
	"github.com/efectn/go-orm-benchmarks/benchs/ent/model"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var client *ent.Client
var dbEnt *sql.DB

func initDBEnt() {
	// Run the auto migration.
	_, err := dbEnt.Exec("DROP TABLE IF EXISTS models")
	CheckErr(err)

	err = client.Schema.Create(ctx)
	CheckErr(err)
}

func init() {
	st := NewSuite("ent")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, EntInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, EntInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, EntUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, EntRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, EntReadSlice)

		var err error
		dbEnt, err = sql.Open("pgx", OrmSource)
		CheckErr(err)

		// Create an ent.Driver from `dbEnt`.
		drv := entsql.OpenDB(dialect.Postgres, dbEnt)

		// Assign to client
		client = ent.NewClient(ent.Driver(drv))

	}
}

func EntInsert(b *B) {
	var m Model
	WrapExecute(b, func() {
		initDBEnt()
		m = NewModelAlt()
	})

	for i := 0; i < b.N; i++ {
		created, err := client.Model.Create().
			SetName(m.Name).
			SetTitle(m.Title).
			SetFax(m.Fax).
			SetWeb(m.Web).
			SetAge(m.Age).
			SetRight(m.Right).
			SetCounter(m.Counter).
			Save(ctx)
		CheckErr(err)

		m.Id = created.ID
	}
}

func EntInsertMulti(b *B) {
	var ms []Model
	WrapExecute(b, func() {
		initDBEnt()
		ms = make([]Model, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewModelAlt())
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

	for i := 0; i < b.N; i++ {
		_, err := client.Model.CreateBulk(bulk...).Save(ctx)
		CheckErr(err)

	}
}

func EntUpdate(b *B) {
	var m Model
	WrapExecute(b, func() {
		initDBEnt()
		m = NewModelAlt()
		created, err := client.Model.Create().
			SetName(m.Name).
			SetTitle(m.Title).
			SetFax(m.Fax).
			SetWeb(m.Web).
			SetAge(m.Age).
			SetRight(m.Right).
			SetCounter(m.Counter).
			Save(ctx)
		CheckErr(err)

		m.Id = created.ID
	})

	for i := 0; i < b.N; i++ {
		_, err := client.Model.Update().
			Where(model.IDEQ(m.Id)).
			SetName(m.Name).
			SetTitle(m.Title).
			SetFax(m.Fax).
			SetWeb(m.Web).
			SetAge(m.Age).
			SetRight(m.Right).
			SetCounter(m.Counter).
			Save(ctx)

		CheckErr(err)
	}
}

func EntRead(b *B) {
	var m Model
	WrapExecute(b, func() {
		initDBEnt()
		m = NewModelAlt()
		_, err := client.Model.Create().
			SetName(m.Name).
			SetTitle(m.Title).
			SetFax(m.Fax).
			SetWeb(m.Web).
			SetAge(m.Age).
			SetRight(m.Right).
			SetCounter(m.Counter).
			Save(ctx)
		CheckErr(err)
	})

	for i := 0; i < b.N; i++ {
		_, err := client.Model.Query().Where(model.IDEQ(1)).First(ctx)
		CheckErr(err)
	}
}

func EntReadSlice(b *B) {
	var m Model
	WrapExecute(b, func() {
		initDBEnt()
		m = NewModelAlt()
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
			CheckErr(err)
		}
	})

	for i := 0; i < b.N; i++ {
		_, err := client.Model.Query().Where(model.IDGT(0)).Unique(false).Limit(100).All(ctx)
		CheckErr(err)
	}
}
