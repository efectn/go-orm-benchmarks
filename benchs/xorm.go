package benchs

import (
	"xorm.io/xorm"
)

var xo *xorm.Session

func init() {
	st := NewSuite("xorm")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, XormInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, XormInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, XormUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, XormRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, XormReadSlice)

		engine, err := xorm.NewEngine("postgres", OrmSource)
		CheckErr(err)

		xo = engine.NewSession()
	}
}

func XormInsert(b *B) {
	var m *Model5
	WrapExecute(b, func() {
		InitDB()
		m = NewModel5()
	})

	for i := 0; i < b.N; i++ {
		m.ID = 0
		_, err := xo.Insert(m)
		CheckErr(err, b)
	}
}

func XormInsertMulti(b *B) {
	var ms []*Model5
	WrapExecute(b, func() {
		InitDB()
		ms = make([]*Model5, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewModel5())
		}
	})

	for i := 0; i < b.N; i++ {
		_, err := xo.Insert(&ms)
		CheckErr(err, b)
	}
}

func XormUpdate(b *B) {
	var m *Model5
	WrapExecute(b, func() {
		InitDB()
		m = NewModel5()
		_, err := xo.Insert(m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		_, err := xo.Update(m)
		CheckErr(err, b)
	}
}

func XormRead(b *B) {
	var m *Model5
	WrapExecute(b, func() {
		InitDB()
		m = NewModel5()
		_, err := xo.Insert(m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		_, err := xo.Get(m)
		CheckErr(err, b)
	}
}

func XormReadSlice(b *B) {
	var m *Model5
	WrapExecute(b, func() {
		InitDB()
		m = NewModel5()
		for i := 0; i < 100; i++ {
			m.ID = 0
			_, err := xo.Insert(m)
			CheckErr(err, b)
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model5
		err := xo.Table("model5").Where("id > 0").Limit(100).Find(&models)
		CheckErr(err, b)
	}
}
