package bench

import (
	"testing"

	"github.com/efectn/go-orm-benchmarks/helper"
	xormdb "xorm.io/xorm"
)

type Xorm struct {
	helper.ORMInterface
	conn *xormdb.Session
}

func CreateXorm() helper.ORMInterface {
	return &Xorm{}
}

func (xorm *Xorm) Name() string {
	return "xorm"
}

func (xorm *Xorm) Init() error {
	engine, err := xormdb.NewEngine("postgres", helper.OrmSource)
	if err != nil {
		return err
	}

	xorm.conn = engine.NewSession()

	return nil
}

func (xorm *Xorm) Close() error {
	return xorm.conn.Close()
}

func (xorm *Xorm) Insert(b *testing.B) {
	m := NewModel5()
	helper.BenchExec(b, xorm.Name(), "Insert", func() error {
		m.ID = 0
		_, err := xorm.conn.Insert(m)
		return err
	})
}

func (xorm *Xorm) InsertMulti(b *testing.B) {
	ms := make([]*Model5, 0, helper.BatchSize)
	for i := 0; i < helper.BatchSize; i++ {
		ms = append(ms, NewModel5())
	}

	helper.BenchExec(b, xorm.Name(), "InsertMulti", func() error {
		for _, m := range ms {
			m.ID = 0
		}
		_, err := xorm.conn.Insert(&ms)
		return err
	})
}

func (xorm *Xorm) Update(b *testing.B) {
	m := NewModel5()

	_, err := xorm.conn.Insert(m)
	if err != nil {
		helper.SetError(b, xorm.Name(), "Update", err.Error())
		return
	}
	m.Counter = helper.UpdatedCounter

	helper.BenchExec(b, xorm.Name(), "Update", func() error {
		_, err := xorm.conn.ID(m.ID).Update(m)
		return err
	})
}

func (xorm *Xorm) Read(b *testing.B) {
	m := NewModel5()

	_, err := xorm.conn.Insert(m)
	if err != nil {
		helper.SetError(b, xorm.Name(), "Read", err.Error())
		return
	}

	helper.BenchRows(b, xorm.Name(), "Read", 1, func() (int, error) {
		var result Model5
		found, err := xorm.conn.ID(m.ID).NoAutoCondition().Get(&result)
		if err != nil {
			return 0, err
		}
		if !found {
			return 0, nil
		}
		return 1, nil
	})
}

func (xorm *Xorm) ReadSlice(b *testing.B) {
	m := NewModel5()

	for i := 0; i < helper.BatchSize; i++ {
		m.ID = 0
		_, err := xorm.conn.Insert(m)
		if err != nil {
			helper.SetError(b, xorm.Name(), "ReadSlice", err.Error())
			return
		}
	}

	helper.BenchRows(b, xorm.Name(), "ReadSlice", helper.BatchSize, func() (int, error) {
		var models []*Model5
		err := xorm.conn.Where("id > 0").Limit(helper.BatchSize).Find(&models)
		if err != nil {
			return 0, err
		}
		return len(models), nil
	})
}
