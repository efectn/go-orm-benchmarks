package benchs

import (
	"fmt"
	"sync"
	"testing"

	dbrware "github.com/gocraft/dbr/v2"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var columns = []string{"name", "title", "fax", "web", "age", "right", "counter"}

type Dbr struct {
	Instance
	mu         sync.Mutex
	conn       *dbrware.Session
	iterations int // Same as b.N, just to customize it
	errors     map[string]string
}

func CreateDbr(iterations int) Instance {
	dbr := &Dbr{
		iterations: iterations,
		errors:     map[string]string{},
	}

	return dbr
}

func (dbr *Dbr) Name() string {
	return "dbr"
}

func (dbr *Dbr) Init() error {
	conn, err := dbrware.Open("postgres", OrmSource, nil)
	if err != nil {
		return err
	}

	dbr.conn = conn.NewSession(nil)
	_, err = dbr.conn.Begin()
	if err != nil {
		return err
	}

	return nil
}

func (dbr *Dbr) Close() error {
	return dbr.conn.Close()
}

func (dbr *Dbr) GetError(method string) string {
	return dbr.errors[method]
}

func (dbr *Dbr) Insert(b *testing.B) {
	m := NewModel2()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := dbr.conn.InsertInto("models").Columns(columns...).Record(m).Exec()
		if err != nil {
			dbr.error(b, "insert", fmt.Sprintf("dbr: insert: %v", err))
		}
	}
}

func (dbr *Dbr) InsertMulti(b *testing.B) {
	dbr.error(b, "insert_multi", "dbr: insert_multi: insert multi is not supported on dbr")
}

func (dbr *Dbr) Update(b *testing.B) {
	m := NewModel2()

	_, err := dbr.conn.InsertInto("models").Columns(columns...).Record(m).Exec()
	if err != nil {
		dbr.error(b, "update", fmt.Sprintf("dbr: update: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := dbr.conn.Update("models").SetMap(map[string]any{
			"name":    m.Name,
			"title":   m.Title,
			"fax":     m.Fax,
			"web":     m.Web,
			"age":     m.Age,
			"right":   m.Right,
			"counter": m.Counter,
		}).Exec()
		if err != nil {
			dbr.error(b, "update", fmt.Sprintf("dbr: update: %v", err))
		}
	}
}

func (dbr *Dbr) Read(b *testing.B) {
	m := NewModel2()

	_, err := dbr.conn.InsertInto("models").Columns(columns...).Record(m).Exec()
	if err != nil {
		dbr.error(b, "read", fmt.Sprintf("dbr: read: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := dbr.conn.Select("*").From("models").Where("id = ?", m.ID).Load(m)
		if err != nil {
			dbr.error(b, "read", fmt.Sprintf("dbr: read: %v", err))
		}
	}
}

func (dbr *Dbr) ReadSlice(b *testing.B) {
	m := NewModel2()
	for i := 0; i < 100; i++ {
		_, err := dbr.conn.InsertInto("models").Columns(columns...).Record(m).Exec()
		if err != nil {
			dbr.error(b, "read_slice", fmt.Sprintf("dbr: read_slice: %v", err))
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var ms []*Model2
		_, err := dbr.conn.Select("*").From("models").Where("id > 0").Limit(100).Load(&ms)
		if err != nil {
			dbr.error(b, "read_slice", fmt.Sprintf("dbr: read_slice: %v", err))
		}
	}
}

func (dbr *Dbr) error(b *testing.B, method string, err string) {
	b.Helper()

	dbr.mu.Lock()
	dbr.errors[method] = err
	dbr.mu.Unlock()
	b.Fail()
}
