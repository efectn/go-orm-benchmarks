package benchs

import (
	db "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

var upper db.Session

func init() {
	st := NewSuite("upper")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, UpperInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, UpperInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, UpperUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, UpperRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, UpperReadSlice)

		var err error

		source := SplitSource()
		upper, err = postgresql.Open(postgresql.ConnectionURL{
			Host:     source["host"] + ":5432",
			User:     source["user"],
			Password: source["password"],
			Database: source["dbname"],
		})
		CheckErr(err)

		// Disable logger
		db.LC().SetLogger(nil)

		err = upper.Ping()
		CheckErr(err)
	}
}

func UpperInsert(b *B) {
	var m *Model4
	WrapExecute(b, func() {
		InitDB()
		m = NewModel4()
	})

	for i := 1; i < b.N+1; i++ {
		_, err := upper.Collection("models").Insert(m)
		CheckErr(err, b)
	}
}

func UpperInsertMulti(b *B) {
	var m *Model4
	WrapExecute(b, func() {
		InitDB()
		m = NewModel4()
	})

	for i := 1; i < b.N+1; i++ {
		batch := upper.SQL().InsertInto("models").Columns("name", "title", "fax", "web", "age", "right", "counter").Batch(100)

		go func() {
			for i := 0; i < 100; i++ {
				batch.Values(m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
			}
			batch.Done()
		}()

		err := batch.Wait()
		CheckErr(err, b)
	}
}

func UpperUpdate(b *B) {
	var m *Model4
	WrapExecute(b, func() {
		InitDB()
		m = NewModel4()

		err := upper.Collection("models").InsertReturning(m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		err := upper.Collection("models").UpdateReturning(m)
		CheckErr(err, b)
	}
}

func UpperRead(b *B) {
	var m *Model4
	WrapExecute(b, func() {
		InitDB()
		m = NewModel4()

		err := upper.Collection("models").InsertReturning(m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		err := upper.SQL().SelectFrom("models").Where("id = ?", m.ID).One(m)
		CheckErr(err, b)
	}
}

func UpperReadSlice(b *B) {
	var m *Model4
	WrapExecute(b, func() {
		InitDB()
		m = NewModel4()

		err := upper.Collection("models").InsertReturning(m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		var ms []*Model4
		err := upper.SQL().SelectFrom("models").Where("id > ?", m.ID).Limit(100).All(&ms)
		CheckErr(err, b)
	}
}
