package benchs

import (
	"database/sql"
	"fmt"
	"sync"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	gorpware "gopkg.in/gorp.v1"
)

type Gorp struct {
	Instance
	mu         sync.Mutex
	conn       *gorpware.DbMap
	iterations int // Same as b.N, just to customize it
	errors     map[string]string
}

func CreateGorp(iterations int) Instance {
	gorp := &Gorp{
		iterations: iterations,
		errors:     map[string]string{},
	}

	return gorp
}

func (gorp *Gorp) Name() string {
	return "gorp"
}

func (gorp *Gorp) Init() error {
	db, err := sql.Open("pgx", OrmSource)
	if err != nil {
		return err
	}

	gorp.conn = &gorpware.DbMap{Db: db, Dialect: gorpware.PostgresDialect{}}
	gorp.conn.AddTableWithName(Model{}, "models").SetKeys(true, "Id")

	return nil
}

func (gorp *Gorp) Close() error {
	return gorp.conn.Db.Close()
}

func (gorp *Gorp) GetError(method string) string {
	return gorp.errors[method]
}

func (gorp *Gorp) Insert(b *testing.B) {
	m := NewModel()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := gorp.conn.Insert(m)
		if err != nil {
			gorp.error(b, "insert", fmt.Sprintf("gorp: insert: %v", err))
		}
	}
}

func (gorp *Gorp) InsertMulti(b *testing.B) {
	gorp.error(b, "insert_multi", "gorp: insert_multi: insert multi is not supported on gorp")
}

func (gorp *Gorp) Update(b *testing.B) {
	m := NewModel()

	err := gorp.conn.Insert(m)
	if err != nil {
		gorp.error(b, "update", fmt.Sprintf("gorp: update: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := gorp.conn.Update(m)
		if err != nil {
			gorp.error(b, "update", fmt.Sprintf("gorp: update: %v", err))
		}
	}
}

func (gorp *Gorp) Read(b *testing.B) {
	m := NewModel()

	err := gorp.conn.Insert(m)
	if err != nil {
		gorp.error(b, "read", fmt.Sprintf("gorp: read: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := gorp.conn.SelectOne(m, "SELECT * FROM models LIMIT 1")
		if err != nil {
			gorp.error(b, "read", fmt.Sprintf("gorp: read: %v", err))
		}
	}
}

// TODO: not working
func (gorp *Gorp) ReadSlice(b *testing.B) {
	m := NewModel()
	for i := 0; i < 100; i++ {
		err := gorp.conn.Insert(m)
		gorp.error(b, "read_slice", fmt.Sprintf("gorp: read_slice: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var ms []*Model
		_, err := gorp.conn.Select(&ms, "select * from models where id > 0 LIMIT 100")
		if err != nil {
			gorp.error(b, "read_slice", fmt.Sprintf("gorp: read_slice: %v", err))
		}
	}
}

func (gorp *Gorp) error(b *testing.B, method string, err string) {
	b.Helper()

	gorp.mu.Lock()
	gorp.errors[method] = err
	gorp.mu.Unlock()
	b.Fail()
}
