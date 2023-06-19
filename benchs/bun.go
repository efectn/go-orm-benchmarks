package benchs

import (
	"database/sql"
	"fmt"
	"sync"
	"testing"

	bundb "github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Bun struct {
	Instance
	mu         sync.Mutex
	conn       *bundb.DB
	iterations int // Same as b.N, just to customize it
	errors     map[string]string
}

func CreateBun(iterations int) Instance {
	bun := &Bun{
		iterations: iterations,
		errors:     map[string]string{},
	}

	return bun
}

func (bun *Bun) Name() string {
	return "bun"
}

func (bun *Bun) Init() error {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(ConvertSourceToDSN())))
	sqldb.SetMaxOpenConns(OrmMaxConn)
	sqldb.SetMaxIdleConns(OrmMaxIdle)

	bun.conn = bundb.NewDB(sqldb, pgdialect.New())
	return nil
}

func (bun *Bun) Close() error {
	return bun.conn.Close()
}

func (bun *Bun) GetError(method string) string {
	return bun.errors[method]
}

func (bun *Bun) Insert(b *testing.B) {
	m := NewModel()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Id = 0

		_, err := bun.conn.NewInsert().Model(m).Exec(ctx)
		if err != nil {
			bun.error(b, "insert", fmt.Sprintf("bun: insert: %v", err))
		}
	}
}

func (bun *Bun) InsertMulti(b *testing.B) {
	ms := make([]*Model, 0, 100)
	for i := 0; i < 100; i++ {
		ms = append(ms, NewModel())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, m := range ms {
			m.Id = 0
		}

		_, err := bun.conn.NewInsert().Model(&ms).Exec(ctx)
		if err != nil {
			bun.error(b, "insert_multi", fmt.Sprintf("bun: insert_multi: %v", err))
		}
	}
}

func (bun *Bun) Update(b *testing.B) {
	m := NewModel()

	_, err := bun.conn.NewInsert().Model(m).Exec(ctx)
	if err != nil {
		bun.error(b, "update", fmt.Sprintf("bun: update: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := bun.conn.NewUpdate().Model(m).WherePK().Exec(ctx)
		if err != nil {
			bun.error(b, "update", fmt.Sprintf("bun: update: %v", err))
		}
	}
}

func (bun *Bun) Read(b *testing.B) {
	m := NewModel()

	_, err := bun.conn.NewInsert().Model(m).Exec(ctx)
	if err != nil {
		bun.error(b, "read", fmt.Sprintf("bun: read: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := bun.conn.NewSelect().Model(m).Scan(ctx)
		if err != nil {
			bun.error(b, "read", fmt.Sprintf("bun: read: %v", err))
		}
	}
}

func (bun *Bun) ReadSlice(b *testing.B) {
	m := NewModel()
	for i := 0; i < 100; i++ {
		m.Id = 0
		_, err := bun.conn.NewInsert().Model(m).Exec(ctx)
		if err != nil {
			bun.error(b, "read_slice", fmt.Sprintf("bun: read_slice: %v", err))
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var models []*Model
		err := bun.conn.NewSelect().
			Model(&models).
			Where("id > ?", 0).
			Limit(100).
			Scan(ctx)
		if err != nil {
			bun.error(b, "read_slice", fmt.Sprintf("bun: read_slice: %v", err))
		}
	}
}

func (bun *Bun) error(b *testing.B, method string, err string) {
	b.Helper()

	bun.mu.Lock()
	bun.errors[method] = err
	bun.mu.Unlock()
	b.Fail()
}
