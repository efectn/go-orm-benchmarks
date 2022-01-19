package benchs

import (
	"github.com/astaxie/beego/orm"
)

var bo orm.Ormer

func init() {
	st := NewSuite("beego")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, BeegoOrmInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, BeegoOrmInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, BeegoOrmUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, BeegoOrmRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, BeegoOrmReadSlice)

		err := orm.RegisterDataBase("default", "postgres", OrmSource, OrmMaxIdle, OrmMaxConn)
		CheckErr(err)

		orm.RegisterModel(new(Model))
		bo = orm.NewOrm()
	}
}

func BeegoOrmInsert(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0

		_, err := bo.Insert(m)
		CheckErr(err, b)
	}
}

func BeegoOrmInsertMulti(b *B) {
	var ms []*Model
	WrapExecute(b, func() {
		InitDB()
		ms = make([]*Model, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewModel())
		}
	})

	for i := 0; i < b.N; i++ {
		_, err := bo.InsertMulti(100, ms)
		CheckErr(err, b)
	}
}

func BeegoOrmUpdate(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()

		_, err := bo.Insert(m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		_, err := bo.Update(m)
		CheckErr(err, b)
	}
}

func BeegoOrmRead(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		_, err := bo.Insert(m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		err := bo.Read(m)
		CheckErr(err, b)
	}
}

func BeegoOrmReadSlice(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.Id = 0
			_, err := bo.Insert(m)
			CheckErr(err, b)
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model
		_, err := bo.QueryTable("models").Filter("id__gt", 0).Limit(100).All(&models)
		CheckErr(err, b)
	}
}
