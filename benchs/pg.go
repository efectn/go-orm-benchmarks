package benchs

import (
	"fmt"
	pgdb "github.com/go-pg/pg/v10"
	"sync"
	"testing"
)

type Pg struct {
	Instance
	mu         sync.Mutex
	conn       *pgdb.DB
	iterations int // Same as b.N, just to customize it
	errors     map[string]string
}

func CreatePg(iterations int) Instance {
	pg := &Pg{
		iterations: iterations,
		errors:     map[string]string{},
	}

	return pg
}

func (pg *Pg) Name() string {
	return "pg"
}

func (pg *Pg) Init() error {
	source := SplitSource()
	pg.conn = pgdb.Connect(&pgdb.Options{
		Addr:     source["host"] + ":5432",
		User:     source["user"],
		Password: source["password"],
		Database: source["dbname"],
	})

	if err := pg.conn.Ping(ctx); err != nil {
		return err
	}

	return nil
}

func (pg *Pg) Close() error {
	return pg.conn.Close()
}

func (pg *Pg) GetError(method string) string {
	return pg.errors[method]
}

func (pg *Pg) Insert(b *testing.B) {
	m := NewModel()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Id = 0
		_, err := pg.conn.Model(m).Insert()
		if err != nil {
			pg.error(b, "insert", fmt.Sprintf("pg: insert: %v", err))
		}
	}
}

func (pg *Pg) InsertMulti(b *testing.B) {
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

		_, err := pg.conn.Model(&ms).Insert()
		if err != nil {
			pg.error(b, "insert_multi", "gorp: insert_multi: insert multi is not supported on gorp")
		}
	}
}

func (pg *Pg) Update(b *testing.B) {
	m := NewModel()

	_, err := pg.conn.Model(m).Insert()
	if err != nil {
		pg.error(b, "update", fmt.Sprintf("pg: update: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := pg.conn.Model(m).WherePK().Update()
		if err != nil {
			pg.error(b, "update", fmt.Sprintf("pg: update: %v", err))
		}
	}
}

func (pg *Pg) Read(b *testing.B) {
	m := NewModel()

	_, err := pg.conn.Model(m).Insert()
	if err != nil {
		pg.error(b, "read", fmt.Sprintf("pg: read: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := pg.conn.Model(m).Limit(1).Select()
		if err != nil {
			pg.error(b, "read", fmt.Sprintf("pg: read: %v", err))
		}
	}
}

func (pg *Pg) ReadSlice(b *testing.B) {
	m := NewModel()
	for i := 0; i < 100; i++ {
		m.Id = 0
		_, err := pg.conn.Model(m).Insert()
		if err != nil {
			pg.error(b, "read_slice", fmt.Sprintf("pg: read_slice: %v", err))
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var models []*Model
		err := pg.conn.Model(&models).Where("id > ?", 0).Limit(100).Select()
		if err != nil {
			pg.error(b, "read_slice", fmt.Sprintf("pg: read_slice: %v", err))
		}
	}
}

func (pg *Pg) error(b *testing.B, method string, err string) {
	b.Helper()

	pg.mu.Lock()
	pg.errors[method] = err
	pg.mu.Unlock()
	b.Fail()
}
