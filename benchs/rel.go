package benchs

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/go-rel/postgres"
	relware "github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
)

type Rel struct {
	Instance
	mu         sync.Mutex
	conn       relware.Repository
	db         relware.Adapter
	iterations int // Same as b.N, just to customize it
	errors     map[string]string
}

func CreateRel(iterations int) Instance {
	rel := &Rel{
		iterations: iterations,
		errors:     map[string]string{},
	}

	return rel
}

func (rel *Rel) Name() string {
	return "rel"
}

func (rel *Rel) Init() error {
	var err error
	rel.db, err = postgres.Open(OrmSource)
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

func (rel *Rel) GetError(method string) string {
	return rel.errors[method]
}

func (rel *Rel) Insert(b *testing.B) {
	m := NewModel3()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.ID = 0
		err := rel.conn.Insert(ctx, m)
		if err != nil {
			rel.error(b, "insert", fmt.Sprintf("rel: insert: %v", err))
		}
	}
}

func (rel *Rel) InsertMulti(b *testing.B) {
	ms := make([]*Model3, 0, 100)
	for i := 0; i < 100; i++ {
		ms = append(ms, NewModel3())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, m := range ms {
			m.ID = 0
		}
		err := rel.conn.InsertAll(ctx, &ms)
		if err != nil {
			rel.error(b, "insert_multi", fmt.Sprintf("rel: insert_multi: %v", err))
		}
	}
}

func (rel *Rel) Update(b *testing.B) {
	m := NewModel3()
	m.ID = 0
	err := rel.conn.Insert(ctx, m)
	if err != nil {
		rel.error(b, "update", fmt.Sprintf("rel: update: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := rel.conn.Update(ctx, m)
		if err != nil {
			rel.error(b, "update", fmt.Sprintf("rel: update: %v", err))
		}
	}
}

func (rel *Rel) Read(b *testing.B) {
	m := NewModel3()
	m.ID = 0
	err := rel.conn.Insert(ctx, m)
	if err != nil {
		rel.error(b, "read", fmt.Sprintf("rel: read: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := rel.conn.Find(ctx, m)
		if err != nil {
			rel.error(b, "read", fmt.Sprintf("rel: read: %v", err))
		}
	}
}

func (rel *Rel) ReadSlice(b *testing.B) {
	m := NewModel3()
	for i := 0; i < 100; i++ {
		m.ID = 0
		err := rel.conn.Insert(ctx, m)
		if err != nil {
			rel.error(b, "read_slice", fmt.Sprintf("rel: read_slice: %v", err))
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var ms []Model3
		err := rel.conn.FindAll(ctx, &ms, where.Gt("id", 0), relware.Limit(100))
		if err != nil {
			rel.error(b, "read_slice", fmt.Sprintf("rel: read_slice: %v", err))
		}
	}
}

func (rel *Rel) error(b *testing.B, method string, err string) {
	b.Helper()

	rel.mu.Lock()
	rel.errors[method] = err
	rel.mu.Unlock()
	b.Fail()
}
