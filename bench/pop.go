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

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := pop.conn.Create(m)
		if err != nil {
			helper.SetError(b, pop.Name(), "Insert", err.Error())
		}
	}
}

func (pop *Pop) InsertMulti(b *testing.B) {
	helper.SetError(b, pop.Name(), "InsertMulti", "bulk-insert is not supported")
}

func (pop *Pop) Update(b *testing.B) {
	m := NewModel3()

	err := pop.conn.Create(m)
	if err != nil {
		helper.SetError(b, pop.Name(), "Update", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := pop.conn.Update(m)
		if err != nil {
			helper.SetError(b, pop.Name(), "Update", err.Error())
		}
	}
}

func (pop *Pop) Read(b *testing.B) {
	m := NewModel3()

	err := pop.conn.Create(m)
	if err != nil {
		helper.SetError(b, pop.Name(), "Read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := pop.conn.First(m)
		if err != nil {
			helper.SetError(b, pop.Name(), "Read", err.Error())
		}
	}
}

func (pop *Pop) ReadSlice(b *testing.B) {
	m := NewModel3()
	for i := 0; i < 100; i++ {
		err := pop.conn.Create(m)
		if err != nil {
			helper.SetError(b, pop.Name(), "ReadSlice", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var ms []Model3
		err := pop.conn.Where("id > 0").Limit(100).All(&ms)
		if err != nil {
			helper.SetError(b, pop.Name(), "ReadSlice", err.Error())
		}
	}
}
