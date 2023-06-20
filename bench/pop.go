package bench

import (
	"github.com/efectn/go-orm-benchmarks/helper"
	"sync"
	"testing"

	popware "github.com/gobuffalo/pop/v6"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Pop struct {
	helper.ORMInterface
	mu         sync.Mutex
	conn       *popware.Connection
	iterations int // Same as b.N, just to customize it
}

func CreatePop(iterations int) helper.ORMInterface {
	pop := &Pop{
		iterations: iterations,
	}

	return pop
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
			helper.SetError(b, pop.Name(), "insert", err.Error())
		}
	}
}

func (pop *Pop) InsertMulti(b *testing.B) {
	helper.SetError(b, pop.Name(), "insert_multi", "insert multi is not supported on pop")
}

func (pop *Pop) Update(b *testing.B) {
	m := NewModel3()

	err := pop.conn.Create(m)
	if err != nil {
		helper.SetError(b, pop.Name(), "update", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := pop.conn.Update(m)
		if err != nil {
			helper.SetError(b, pop.Name(), "update", err.Error())
		}
	}
}

func (pop *Pop) Read(b *testing.B) {
	m := NewModel3()

	err := pop.conn.Create(m)
	if err != nil {
		helper.SetError(b, pop.Name(), "read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := pop.conn.First(m)
		if err != nil {
			helper.SetError(b, pop.Name(), "read", err.Error())
		}
	}
}

func (pop *Pop) ReadSlice(b *testing.B) {
	m := NewModel3()
	for i := 0; i < 100; i++ {
		err := pop.conn.Create(m)
		if err != nil {
			helper.SetError(b, pop.Name(), "read_slice", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var ms []Model3
		err := pop.conn.Where("id > 0").Limit(100).All(&ms)
		if err != nil {
			helper.SetError(b, pop.Name(), "read_slice", err.Error())
		}
	}
}
