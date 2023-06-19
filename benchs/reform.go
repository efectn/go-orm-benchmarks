package benchs

import (
	"database/sql"
	"fmt"
	"sync"
	"testing"

	r "github.com/efectn/go-orm-benchmarks/benchs/reform"
	_ "github.com/jackc/pgx/v4/stdlib"
	"gopkg.in/reform.v1/dialects/postgresql"

	reformware "gopkg.in/reform.v1"
)

type Reform struct {
	Instance
	mu         sync.Mutex
	conn       *reformware.DB
	db         *sql.DB
	iterations int // Same as b.N, just to customize it
	errors     map[string]string
}

func CreateReform(iterations int) Instance {
	reform := &Reform{
		iterations: iterations,
		errors:     map[string]string{},
	}

	return reform
}

func (reform *Reform) Name() string {
	return "reform"
}

func (reform *Reform) Init() error {
	var err error
	reform.db, err = sql.Open("pgx", OrmSource)
	if err != nil {
		return err
	}

	reform.conn = reformware.NewDB(reform.db, postgresql.Dialect, nil)

	return nil
}

func (reform *Reform) Close() error {
	return reform.db.Close()
}

func (reform *Reform) GetError(method string) string {
	return reform.errors[method]
}

func (reform *Reform) Insert(b *testing.B) {
	m := NewReformModel()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := reform.conn.Save(m)
		if err != nil {
			reform.error(b, "insert", fmt.Sprintf("reform: insert: %v", err))
		}
	}
}

func (reform *Reform) InsertMulti(b *testing.B) {
	ms := make([]reformware.Struct, 0, 100)
	for i := 0; i < 100; i++ {
		ms = append(ms, NewReformModel())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := reform.conn.InsertMulti(ms...)
		if err != nil {
			reform.error(b, "insert_multi", fmt.Sprintf("reform: insert_multi: %v", err))
		}
	}
}

func (reform *Reform) Update(b *testing.B) {
	m := NewReformModel()

	err := reform.conn.Save(m)
	if err != nil {
		reform.error(b, "update", fmt.Sprintf("reform: update: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := reform.conn.Update(m)
		if err != nil {
			reform.error(b, "update", fmt.Sprintf("reform: update: %v", err))
		}
	}
}

func (reform *Reform) Read(b *testing.B) {
	m := NewReformModel()

	err := reform.conn.Save(m)
	if err != nil {
		reform.error(b, "read", fmt.Sprintf("reform: read: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := reform.conn.FindByPrimaryKeyFrom(r.ReformModelsTable, m.ID)
		if err != nil {
			reform.error(b, "read", fmt.Sprintf("reform: read: %v", err))
		}
	}
}

func (reform *Reform) ReadSlice(b *testing.B) {
	m := NewReformModel()
	for i := 0; i < 100; i++ {
		err := reform.conn.Save(m)
		if err != nil {
			reform.error(b, "read_slice", fmt.Sprintf("reform: read_slice: %v", err))
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := reform.conn.SelectAllFrom(r.ReformModelsTable, "WHERE id > 0 LIMIT 100")
		if err != nil {
			reform.error(b, "read_slice", fmt.Sprintf("reform: read_slice: %v", err))
		}
	}
}

func (reform *Reform) error(b *testing.B, method string, err string) {
	b.Helper()

	reform.mu.Lock()
	reform.errors[method] = err
	reform.mu.Unlock()
	b.Fail()
}
