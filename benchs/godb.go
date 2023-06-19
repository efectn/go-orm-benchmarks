package benchs

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	godbware "github.com/samonzeweb/godb"
	"github.com/samonzeweb/godb/adapters/postgresql"
	"sync"
	"testing"
)

type Godb struct {
	Instance
	mu         sync.Mutex
	conn       *godbware.DB
	iterations int // Same as b.N, just to customize it
	errors     map[string]string
}

func CreateGodb(iterations int) Instance {
	godb := &Godb{
		iterations: iterations,
		errors:     map[string]string{},
	}

	return godb
}

func (godb *Godb) Name() string {
	return "godb"
}

func (godb *Godb) Init() error {
	var err error
	godb.conn, err = godbware.Open(postgresql.Adapter, OrmSource)
	if err != nil {
		return err
	}

	return nil
}

func (godb *Godb) Close() error {
	return godb.conn.Close()
}

func (godb *Godb) GetError(method string) string {
	return godb.errors[method]
}

func (godb *Godb) Insert(b *testing.B) {
	m := NewModel2()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := godb.conn.Insert(m).Do()
		if err != nil {
			godb.error(b, "insert", fmt.Sprintf("godb: insert: %v", err))
		}
	}
}

func (godb *Godb) InsertMulti(b *testing.B) {
	ms := make([]*Model2, 0, 100)
	for i := 0; i < 100; i++ {
		ms = append(ms, NewModel2())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := godb.conn.BulkInsert(&ms).Do()
		if err != nil {
			godb.error(b, "insert_multi", fmt.Sprintf("godb: insert_multi: %v", err))
		}
	}
}

func (godb *Godb) Update(b *testing.B) {
	m := NewModel2()

	err := godb.conn.Insert(m).Do()
	if err != nil {
		godb.error(b, "update", fmt.Sprintf("godb: update: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := godb.conn.Update(m).Do()
		if err != nil {
			godb.error(b, "update", fmt.Sprintf("godb: update: %v", err))
		}
	}
}

func (godb *Godb) Read(b *testing.B) {
	m := NewModel2()

	err := godb.conn.Insert(m).Do()
	if err != nil {
		godb.error(b, "read", fmt.Sprintf("godb: read: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := godb.conn.Select(m).Do()
		if err != nil {
			godb.error(b, "read", fmt.Sprintf("godb: read: %v", err))
		}
	}
}

func (godb *Godb) ReadSlice(b *testing.B) {
	m := NewModel2()
	for i := 0; i < 100; i++ {
		err := godb.conn.Insert(m).Do()
		if err != nil {
			godb.error(b, "read_slice", fmt.Sprintf("godb: read_slice: %v", err))
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var ms []*Model2
		err := godb.conn.Select(&ms).Where("id > 0").Limit(100).Do()
		if err != nil {
			godb.error(b, "read_slice", fmt.Sprintf("godb: read_slice: %v", err))
		}
	}
}

func (godb *Godb) error(b *testing.B, method string, err string) {
	b.Helper()

	godb.mu.Lock()
	godb.errors[method] = err
	godb.mu.Unlock()
	b.Fail()
}
