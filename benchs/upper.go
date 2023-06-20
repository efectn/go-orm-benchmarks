package benchs

import (
	"fmt"
	db "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"sync"
	"testing"
)

type Upper struct {
	Instance
	mu         sync.Mutex
	conn       db.Session
	iterations int // Same as b.N, just to customize it
	errors     map[string]string
}

func CreateUpper(iterations int) Instance {
	upper := &Upper{
		iterations: iterations,
		errors:     map[string]string{},
	}

	return upper
}

func (upper *Upper) Name() string {
	return "upper"
}

func (upper *Upper) Init() error {
	var err error
	source := SplitSource()

	upper.conn, err = postgresql.Open(postgresql.ConnectionURL{
		Host:     source["host"] + ":5432",
		User:     source["user"],
		Password: source["password"],
		Database: source["dbname"],
	})
	if err != nil {
		return err
	}

	// Disable logger
	db.LC().SetLogger(nil)

	if err := upper.conn.Ping(); err != nil {
		return err
	}

	return nil
}

func (upper *Upper) Close() error {
	return upper.conn.Close()
}

func (upper *Upper) GetError(method string) string {
	return upper.errors[method]
}

func (upper *Upper) Insert(b *testing.B) {
	m := NewModel4()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := upper.conn.Collection("models").Insert(m)
		if err != nil {
			upper.error(b, "insert", fmt.Sprintf("upper: insert: %v", err))
		}
	}
}

func (upper *Upper) InsertMulti(b *testing.B) {
	m := NewModel4()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		batch := upper.conn.SQL().InsertInto("models").Columns("name", "title", "fax", "web", "age", "right", "counter").Batch(100)

		go func() {
			for i := 0; i < 100; i++ {
				batch.Values(m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
			}
			batch.Done()
		}()

		if err := batch.Wait(); err != nil {
			upper.error(b, "insert_multi", fmt.Sprintf("upper: insert_multi: %v", err))
		}
	}
}

func (upper *Upper) Update(b *testing.B) {
	m := NewModel4()

	err := upper.conn.Collection("models").InsertReturning(m)
	if err != nil {
		upper.error(b, "update", fmt.Sprintf("upper: update: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := upper.conn.Collection("models").UpdateReturning(m)
		if err != nil {
			upper.error(b, "update", fmt.Sprintf("upper: update: %v", err))
		}
	}
}

func (upper *Upper) Read(b *testing.B) {
	m := NewModel4()

	err := upper.conn.Collection("models").InsertReturning(m)
	if err != nil {
		upper.error(b, "read", fmt.Sprintf("upper: read: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := upper.conn.SQL().SelectFrom("models").Where("id = ?", m.ID).One(m)
		if err != nil {
			upper.error(b, "read", fmt.Sprintf("upper: read: %v", err))
		}
	}
}

func (upper *Upper) ReadSlice(b *testing.B) {
	m := NewModel4()
	for i := 0; i < 100; i++ {
		m.ID = 0
		err := upper.conn.Collection("models").InsertReturning(m)
		if err != nil {
			upper.error(b, "read_slice", fmt.Sprintf("upper: read_slice: %v", err))
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var ms []*Model4
		err := upper.conn.SQL().SelectFrom("models").Where("id > ?", m.ID).Limit(100).All(&ms)
		if err != nil {
			upper.error(b, "read_slice", fmt.Sprintf("upper: read_slice: %v", err))
		}
	}
}

func (upper *Upper) error(b *testing.B, method string, err string) {
	b.Helper()

	upper.mu.Lock()
	upper.errors[method] = err
	upper.mu.Unlock()
	b.Fail()
}
