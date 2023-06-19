package benchs

import (
	"fmt"

	dbrware "github.com/gocraft/dbr/v2"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var dbr *dbrware.Session
var columns = []string{"name", "title", "fax", "web", "age", "right", "counter"}

func init() {
	st := NewSuite("dbr")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, DbrInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, DbrInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, DbrUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, DbrRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, DbrReadSlice)

		conn, err := dbrware.Open("postgres", OrmSource, nil)
		CheckErr(err)

		dbr = conn.NewSession(nil)
		dbr.Begin()
	}
}

func DbrInsert(b *B) {
	var m *Model2
	WrapExecute(b, func() {
		InitDB()
		m = NewModel2()
	})

	for i := 0; i < b.N; i++ {
		_, err := dbr.InsertInto("models").Columns(columns...).Record(m).Exec()
		CheckErr(err, b)
	}
}

func DbrInsertMulti(b *B) {
	panic(fmt.Errorf("doesn't support bulk-insert"))
}

func DbrUpdate(b *B) {
	var m *Model2
	WrapExecute(b, func() {
		InitDB()
		m = NewModel2()
		_, err := dbr.InsertInto("models").Columns(columns...).Record(m).Exec()
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		_, err := dbr.Update("models").SetMap(map[string]interface{}{
			"name":    m.Name,
			"title":   m.Title,
			"fax":     m.Fax,
			"web":     m.Web,
			"age":     m.Age,
			"right":   m.Right,
			"counter": m.Counter,
		}).Exec()
		CheckErr(err, b)
	}
}

func DbrRead(b *B) {
	var m *Model2
	WrapExecute(b, func() {
		InitDB()
		m = NewModel2()
		_, err := dbr.InsertInto("models").Columns(columns...).Record(m).Exec()
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		_, err := dbr.Select("*").From("models").Where("id = ?", m.ID).Load(m)
		CheckErr(err, b)
	}
}

func DbrReadSlice(b *B) {
	var m *Model2
	WrapExecute(b, func() {
		InitDB()
		m = NewModel2()
		for i := 0; i < 100; i++ {
			_, err := dbr.InsertInto("models").Columns(columns...).Record(m).Exec()
			CheckErr(err, b)
		}
	})

	for i := 0; i < b.N; i++ {
		var ms []*Model2
		_, err := dbr.Select("*").From("models").Where("id > 0").Limit(100).Load(&ms)
		CheckErr(err, b)
	}
}
