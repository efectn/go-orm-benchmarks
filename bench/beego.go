package bench

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/efectn/go-orm-benchmarks/helper"
	"sync"
	"testing"
)

type Beego struct {
	helper.ORMInterface
	mu         sync.Mutex
	conn       orm.Ormer
	iterations int // Same as b.N, just to customize it
}

func CreateBeego(iterations int) helper.ORMInterface {
	beego := &Beego{
		iterations: iterations,
	}

	return beego
}

func (beego *Beego) Name() string {
	return "beego"
}

func (beego *Beego) Init() error {
	err := orm.RegisterDataBase("default", "postgres", helper.OrmSource, helper.OrmMaxIdle, helper.OrmMaxConn)
	if err != nil {
		return err
	}

	orm.RegisterModel(new(Model))
	beego.conn = orm.NewOrm()

	return nil
}

func (beego *Beego) Close() error {
	return nil
}

func (beego *Beego) Insert(b *testing.B) {
	m := NewModel()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Id = 0

		_, err := beego.conn.Insert(m)
		if err != nil {
			helper.SetError(b, beego.Name(), "insert", fmt.Sprintf("beego: insert: %v", err))
		}
	}
}

func (beego *Beego) InsertMulti(b *testing.B) {
	ms := make([]*Model, 0, 100)
	for i := 0; i < 100; i++ {
		ms = append(ms, NewModel())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := beego.conn.InsertMulti(100, ms)
		if err != nil {
			helper.SetError(b, beego.Name(), "insert_multi", err.Error())
		}
	}
}

func (beego *Beego) Update(b *testing.B) {
	m := NewModel()

	_, err := beego.conn.Insert(m)
	if err != nil {
		helper.SetError(b, beego.Name(), "update", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := beego.conn.Update(m)
		if err != nil {
			helper.SetError(b, beego.Name(), "update", err.Error())
		}
	}
}

func (beego *Beego) Read(b *testing.B) {
	m := NewModel()

	_, err := beego.conn.Insert(m)
	if err != nil {
		helper.SetError(b, beego.Name(), "read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := beego.conn.Read(m)
		if err != nil {
			helper.SetError(b, beego.Name(), "read", err.Error())
		}
	}
}

func (beego *Beego) ReadSlice(b *testing.B) {
	m := NewModel()
	for i := 0; i < 100; i++ {
		m.Id = 0
		_, err := beego.conn.Insert(m)
		if err != nil {
			helper.SetError(b, beego.Name(), "read_slice", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var models []*Model
		_, err := beego.conn.QueryTable("models").Filter("id__gt", 0).Limit(100).All(&models)
		if err != nil {
			helper.SetError(b, beego.Name(), "read_slice", err.Error())
		}
	}
}
