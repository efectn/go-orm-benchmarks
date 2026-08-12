package bench

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/efectn/go-orm-benchmarks/helper"
	"testing"
)

type Beego struct {
	helper.ORMInterface
	conn orm.Ormer
}

func CreateBeego() helper.ORMInterface {
	return &Beego{}
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
	helper.BenchExec(b, beego.Name(), "Insert", func() error {
		m.Id = 0
		_, err := beego.conn.Insert(m)
		if err != nil {
			return fmt.Errorf("beego: insert: %w", err)
		}
		return nil
	})
}

func (beego *Beego) InsertMulti(b *testing.B) {
	ms := make([]*Model, 0, helper.BatchSize)
	for i := 0; i < helper.BatchSize; i++ {
		ms = append(ms, NewModel())
	}
	helper.BenchExec(b, beego.Name(), "InsertMulti", func() error {
		for _, m := range ms {
			m.Id = 0
		}
		_, err := beego.conn.InsertMulti(helper.BatchSize, ms)
		return err
	})
}

func (beego *Beego) Update(b *testing.B) {
	m := NewModel()

	_, err := beego.conn.Insert(m)
	if err != nil {
		helper.SetError(b, beego.Name(), "Update", err.Error())
		return
	}
	m.Counter = helper.UpdatedCounter

	helper.BenchExec(b, beego.Name(), "Update", func() error {
		_, err := beego.conn.Update(m)
		return err
	})
}

func (beego *Beego) Read(b *testing.B) {
	m := NewModel()

	_, err := beego.conn.Insert(m)
	if err != nil {
		helper.SetError(b, beego.Name(), "Read", err.Error())
		return
	}

	helper.BenchRows(b, beego.Name(), "Read", 1, func() (int, error) {
		if err := beego.conn.Read(m); err != nil {
			return 0, err
		}
		return 1, nil
	})
}

func (beego *Beego) ReadSlice(b *testing.B) {
	m := NewModel()
	for i := 0; i < helper.BatchSize; i++ {
		m.Id = 0
		_, err := beego.conn.Insert(m)
		if err != nil {
			helper.SetError(b, beego.Name(), "ReadSlice", err.Error())
			return
		}
	}

	helper.BenchRows(b, beego.Name(), "ReadSlice", helper.BatchSize, func() (int, error) {
		var models []*Model
		_, err := beego.conn.QueryTable("models").Filter("id__gt", 0).Limit(helper.BatchSize).All(&models)
		return len(models), err
	})
}
