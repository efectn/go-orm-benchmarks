package benchs

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	godbware "github.com/samonzeweb/godb"
	"github.com/samonzeweb/godb/adapters/postgresql"
)

var godb *godbware.DB

func init() {
	st := NewSuite("godb")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, GodbInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, GodbInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, GodbUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, GodbRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, GodbReadSlice)

		var err error
		godb, err = godbware.Open(postgresql.Adapter, OrmSource)
		CheckErr(err)
	}
}

func GodbInsert(b *B) {
	var m *Model2
	WrapExecute(b, func() {
		InitDB()
		m = NewModel2()
	})

	for i := 0; i < b.N; i++ {
		err := godb.Insert(m).Do()
		CheckErr(err, b)
	}
}

func GodbInsertMulti(b *B) {
	var ms []*Model2
	WrapExecute(b, func() {
		InitDB()
		ms = make([]*Model2, 0, 100)
		for i := 1; i < 101; i++ {
			ms = append(ms, NewModel2())
		}
	})

	for i := 0; i < b.N; i++ {
		err := godb.BulkInsert(&ms).Do()
		CheckErr(err, b)

	}
}

func GodbUpdate(b *B) {
	var m *Model2
	WrapExecute(b, func() {
		InitDB()
		m = NewModel2()
		err := godb.Insert(m).Do()
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		err := godb.Update(m).Do()
		CheckErr(err, b)
	}
}

func GodbRead(b *B) {
	var m *Model2
	WrapExecute(b, func() {
		InitDB()
		m = NewModel2()
		err := godb.Insert(m).Do()
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		err := godb.Select(m).Do()
		CheckErr(err, b)
	}
}

func GodbReadSlice(b *B) {
	var m *Model2
	WrapExecute(b, func() {
		InitDB()
		m = NewModel2()
		for i := 0; i < 100; i++ {
			err := godb.Insert(m).Do()
			CheckErr(err, b)
		}
	})

	for i := 0; i < b.N; i++ {
		var ms []*Model2
		err := godb.Select(&ms).Where("id > 0").Limit(100).Do()
		CheckErr(err, b)
	}
}
