package bench

import (
	"testing"

	"github.com/efectn/go-orm-benchmarks/helper"

	popware "github.com/gobuffalo/pop/v6"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Pop struct {
	helper.ORMInterface
	conn *popware.Connection
}

func CreatePop() helper.ORMInterface {
	return &Pop{}
}

func (pop *Pop) Name() string {
	return "pop"
}

func (pop *Pop) Init() error {
	var err error
	pop.conn, err = popware.NewConnection(&popware.ConnectionDetails{
		URL: helper.ConvertSourceToDSN(),
	})
	if err != nil {
		return err
	}

	if err = pop.conn.Open(); err != nil {
		return err
	}

	return nil
}

func (pop *Pop) Close() error {
	return pop.conn.Close()
}

func (pop *Pop) Insert(b *testing.B) {
	m := NewModel3()
	helper.BenchExec(b, pop.Name(), "Insert", func() error {
		m.ID = 0
		return pop.conn.Create(m)
	})
}

func (pop *Pop) InsertMulti(b *testing.B) {
	helper.SetError(b, pop.Name(), "InsertMulti", "bulk-insert is not supported")
}

func (pop *Pop) Update(b *testing.B) {
	m := NewModel3()

	err := pop.conn.Create(m)
	if err != nil {
		helper.SetError(b, pop.Name(), "Update", err.Error())
		return
	}
	m.Counter = helper.UpdatedCounter

	helper.BenchExec(b, pop.Name(), "Update", func() error {
		return pop.conn.Update(m)
	})
}

func (pop *Pop) Read(b *testing.B) {
	m := NewModel3()

	err := pop.conn.Create(m)
	if err != nil {
		helper.SetError(b, pop.Name(), "Read", err.Error())
		return
	}

	helper.BenchRows(b, pop.Name(), "Read", 1, func() (int, error) {
		var result Model3
		if err := pop.conn.Find(&result, m.ID); err != nil {
			return 0, err
		}
		return 1, nil
	})
}

func (pop *Pop) ReadSlice(b *testing.B) {
	m := NewModel3()
	for i := 0; i < helper.BatchSize; i++ {
		m.ID = 0
		err := pop.conn.Create(m)
		if err != nil {
			helper.SetError(b, pop.Name(), "ReadSlice", err.Error())
			return
		}
	}

	helper.BenchRows(b, pop.Name(), "ReadSlice", helper.BatchSize, func() (int, error) {
		var ms []Model3
		err := pop.conn.Where("id > 0").Limit(helper.BatchSize).All(&ms)
		if err != nil {
			return 0, err
		}
		return len(ms), nil
	})
}
