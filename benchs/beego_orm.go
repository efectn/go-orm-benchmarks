package benchs

import (
	"fmt"

	"database/sql"

	"github.com/astaxie/beego/orm"
)

var bo orm.Ormer

func initDB3() {

	sqls := []string{
		`DROP TABLE IF EXISTS beego_model;`,
		`CREATE TABLE beego_model (
			id SERIAL NOT NULL,
			name text NOT NULL,
			title text NOT NULL,
			fax text NOT NULL,
			web text NOT NULL,
			age integer NOT NULL,
			"right" boolean NOT NULL,
			counter bigint NOT NULL,
			CONSTRAINT beego_model_pkey PRIMARY KEY (id)
			) WITH (OIDS=FALSE);`,
	}

	DB, err := sql.Open("postgres", ORM_SOURCE)
	checkErr(err)
	defer DB.Close()

	err = DB.Ping()
	checkErr(err)

	for _, sql := range sqls {
		_, err = DB.Exec(sql)
		checkErr(err)
	}
}

type BeegoModel struct {
	Id      int `orm:"auto"`
	Name    string
	Title   string
	Fax     string
	Web     string
	Age     int
	Right   bool
	Counter int64
}

func NewBeegoModel() *BeegoModel {
	m := new(BeegoModel)
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
	st := NewSuite("beego_orm")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*ORM_MULTI, BeegoOrmInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*ORM_MULTI, BeegoOrmInsertMulti)
		st.AddBenchmark("Update", 200*ORM_MULTI, BeegoOrmUpdate)
		st.AddBenchmark("Read", 200*ORM_MULTI, BeegoOrmRead)
		st.AddBenchmark("MultiRead limit 100", 200*ORM_MULTI, BeegoOrmReadSlice)

		orm.RegisterDataBase("default", "postgres", ORM_SOURCE, ORM_MAX_IDLE, ORM_MAX_CONN)
		orm.RegisterModel(new(BeegoModel))

		bo = orm.NewOrm()
	}
}

func BeegoOrmInsert(b *B) {
	var m *BeegoModel
	wrapExecute(b, func() {
		initDB3()
		m = NewBeegoModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		if _, err := bo.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func BeegoOrmInsertMulti(b *B) {
	var ms []*BeegoModel
	wrapExecute(b, func() {
		initDB3()
		ms = make([]*BeegoModel, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewBeegoModel())
		}
	})

	for i := 0; i < b.N; i++ {
		if _, err := bo.InsertMulti(100, ms); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func BeegoOrmUpdate(b *B) {
	var m *BeegoModel
	wrapExecute(b, func() {
		initDB3()
		m = NewBeegoModel()
		if _, err := bo.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if _, err := bo.Update(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func BeegoOrmRead(b *B) {
	var m *BeegoModel
	wrapExecute(b, func() {
		initDB3()
		m = NewBeegoModel()
		if _, err := bo.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if err := bo.Read(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func BeegoOrmReadSlice(b *B) {
	var m *BeegoModel
	wrapExecute(b, func() {
		initDB3()
		m = NewBeegoModel()
		for i := 0; i < 100; i++ {
			m.Id = 0
			if _, err := bo.Insert(m); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*BeegoModel
		if _, err := bo.QueryTable("beego_model").Filter("id__gt", 0).Limit(100).All(&models); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
