package benchs

import (
	"fmt"
	"log"

	dbrware "github.com/gocraft/dbr/v2"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var dbr *dbrware.Session
var columns = []string{"name", "title", "fax", "web", "age", "right", "counter"}

func initDB9() {
	conn, err := dbrware.Open("postgres", OrmSource, nil)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	dbr = conn.NewSession(nil)
	dbr.Begin()

	// Run the auto migration tool.
	if _, err = dbr.Exec("DROP TABLE IF EXISTS model_godb"); err != nil {
		log.Fatal(err)
	}
	initDB()
}

func init() {
	st := NewSuite("dbr")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, DbrInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, DbrInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, DbrUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, DbrRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, DbrReadSlice)
	}
}

func DbrInsert(b *B) {
	var m ModelGodb
	wrapExecute(b, func() {
		initDB9()
		m = NewGodbModel()
	})

	for i := 0; i < b.N; i++ {
		if _, err := dbr.InsertInto("models_godb").Columns(columns...).Record(&m).Exec(); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}

func DbrInsertMulti(b *B) {
	panic(fmt.Errorf("doesn't support bulk-insert"))
}

func DbrUpdate(b *B) {
	var m ModelGodb
	wrapExecute(b, func() {
		initDB9()
		m = NewGodbModel()
		if _, err := dbr.InsertInto("models_godb").Columns(columns...).Record(&m).Exec(); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if _, err := dbr.Update("models_godb").SetMap(map[string]interface{}{
			"name":    m.Name,
			"title":   m.Title,
			"fax":     m.Fax,
			"web":     m.Web,
			"age":     m.Age,
			"right":   m.Right,
			"counter": m.Counter,
		}).Exec(); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}

func DbrRead(b *B) {
	var m ModelGodb
	wrapExecute(b, func() {
		initDB9()
		m = NewGodbModel()
		if _, err := dbr.InsertInto("models_godb").Columns(columns...).Record(&m).Exec(); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if _, err := dbr.Select("*").From("models_godb").Where("id = ?", m.ID).Load(&m); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}

func DbrReadSlice(b *B) {
	var m ModelGodb
	wrapExecute(b, func() {
		initDB9()
		m = NewGodbModel()
		for i := 0; i < 100; i++ {
			if _, err := dbr.InsertInto("models_godb").Columns(columns...).Record(&m).Exec(); err != nil {
				log.Fatal(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var ms []ModelGodb
		if _, err := dbr.Select("*").From("models_godb").Where("id > 0").Limit(100).Load(&ms); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}
