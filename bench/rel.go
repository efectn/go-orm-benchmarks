package bench

import (
	"context"
	"testing"

	"github.com/efectn/go-orm-benchmarks/helper"
	"github.com/go-rel/postgres"
	relware "github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
)

type Rel struct {
	helper.ORMInterface
	conn relware.Repository
	db   relware.Adapter
}

func CreateRel() helper.ORMInterface {
	return &Rel{}
}

func (rel *Rel) Name() string {
	return "rel"
}

func (rel *Rel) Init() error {
	var err error
	rel.db, err = postgres.Open(helper.OrmSource)
	if err != nil {
		return err
	}

	rel.conn = relware.New(rel.db)
	if err := rel.conn.Ping(ctx); err != nil {
		return err
	}

	// Disable debug logging
	rel.conn.Instrumentation(func(ctx context.Context, op string, message string, args ...any) func(err error) {
		return func(err error) {
			if err != nil {
				panic(err)
			}
		}
	})

	return nil
}

func (rel *Rel) Close() error {
	return rel.db.Close()
}

func (rel *Rel) Insert(b *testing.B) {
	m := NewModel3()
	helper.BenchExec(b, rel.Name(), "Insert", func() error {
		m.ID = 0
		return rel.conn.Insert(ctx, m)
	})
}

func (rel *Rel) InsertMulti(b *testing.B) {
	ms := make([]*Model3, 0, helper.BatchSize)
	for i := 0; i < helper.BatchSize; i++ {
		ms = append(ms, NewModel3())
	}

	helper.BenchExec(b, rel.Name(), "InsertMulti", func() error {
		for _, m := range ms {
			m.ID = 0
		}
		return rel.conn.InsertAll(ctx, &ms)
	})
}

func (rel *Rel) Update(b *testing.B) {
	m := NewModel3()
	m.ID = 0
	err := rel.conn.Insert(ctx, m)
	if err != nil {
		helper.SetError(b, rel.Name(), "Update", err.Error())
		return
	}
	m.Counter = helper.UpdatedCounter

	helper.BenchExec(b, rel.Name(), "Update", func() error {
		return rel.conn.Update(ctx, m)
	})
}

func (rel *Rel) Read(b *testing.B) {
	m := NewModel3()
	m.ID = 0
	err := rel.conn.Insert(ctx, m)
	if err != nil {
		helper.SetError(b, rel.Name(), "Read", err.Error())
		return
	}

	helper.BenchRows(b, rel.Name(), "Read", 1, func() (int, error) {
		var result Model3
		if err := rel.conn.Find(ctx, &result, where.Eq("id", m.ID)); err != nil {
			return 0, err
		}
		return 1, nil
	})
}

func (rel *Rel) ReadSlice(b *testing.B) {
	m := NewModel3()
	for i := 0; i < helper.BatchSize; i++ {
		m.ID = 0
		err := rel.conn.Insert(ctx, m)
		if err != nil {
			helper.SetError(b, rel.Name(), "ReadSlice", err.Error())
			return
		}
	}

	helper.BenchRows(b, rel.Name(), "ReadSlice", helper.BatchSize, func() (int, error) {
		var ms []Model3
		err := rel.conn.FindAll(ctx, &ms, where.Gt("id", 0), relware.Limit(helper.BatchSize))
		if err != nil {
			return 0, err
		}
		return len(ms), nil
	})
}
