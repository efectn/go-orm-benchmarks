package benchs

import (
	"database/sql"
	"fmt"
	"sync"
	"testing"

	models "github.com/efectn/go-orm-benchmarks/benchs/sqlboiler"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Sqlboiler struct {
	Instance
	mu         sync.Mutex
	conn       *sql.DB
	iterations int // Same as b.N, just to customize it
	errors     map[string]string
}

func CreateSqlboiler(iterations int) Instance {
	sqlboiler := &Sqlboiler{
		iterations: iterations,
		errors:     map[string]string{},
	}

	return sqlboiler
}

func (sqlboiler *Sqlboiler) Name() string {
	return "sqlboiler"
}

func (sqlboiler *Sqlboiler) Init() error {
	var err error
	sqlboiler.conn, err = sql.Open("pgx", OrmSource)
	if err != nil {
		return err
	}

	boil.SetDB(sqlboiler.conn)

	return nil
}

func (sqlboiler *Sqlboiler) Close() error {
	return sqlboiler.conn.Close()
}

func (sqlboiler *Sqlboiler) GetError(method string) string {
	return sqlboiler.errors[method]
}

func (sqlboiler *Sqlboiler) Insert(b *testing.B) {
	m := NewModel6()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.ID = 0
		err := m.Insert(ctx, sqlboiler.conn, boil.Infer())
		if err != nil {
			sqlboiler.error(b, "insert", fmt.Sprintf("sqlboiler: insert: %v", err))
		}
	}
}

func (sqlboiler *Sqlboiler) InsertMulti(b *testing.B) {
	sqlboiler.error(b, "insert_multi", "sqlboiler: insert_multi: insert multi is not supported on sqlboiler")
}

func (sqlboiler *Sqlboiler) Update(b *testing.B) {
	m := NewModel6()
	m.ID = 0
	err := m.Insert(ctx, sqlboiler.conn, boil.Infer())
	if err != nil {
		sqlboiler.error(b, "update", fmt.Sprintf("sqlboiler: update: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := m.Update(ctx, sqlboiler.conn, boil.Infer())
		if err != nil {
			sqlboiler.error(b, "update", fmt.Sprintf("sqlboiler: update: %v", err))
		}
	}
}

func (sqlboiler *Sqlboiler) Read(b *testing.B) {
	m := NewModel6()
	m.ID = 0
	err := m.Insert(ctx, sqlboiler.conn, boil.Infer())
	if err != nil {
		sqlboiler.error(b, "read", fmt.Sprintf("sqlboiler: read: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := models.Models(qm.Where("id = 0")).Exec(sqlboiler.conn)
		if err != nil {
			sqlboiler.error(b, "read", fmt.Sprintf("sqlboiler: read: %v", err))
		}
	}
}

func (sqlboiler *Sqlboiler) ReadSlice(b *testing.B) {
	m := NewModel6()
	for i := 0; i < 100; i++ {
		m.ID = 0
		err := m.Insert(ctx, sqlboiler.conn, boil.Infer())
		if err != nil {
			sqlboiler.error(b, "read_slice", fmt.Sprintf("sqlboiler: read_slice: %v", err))
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := models.Models(qm.Where("id > 0"), qm.Limit(100)).All(ctx, sqlboiler.conn)
		if err != nil {
			sqlboiler.error(b, "read_slice", fmt.Sprintf("sqlboiler: read_slice: %v", err))
		}
	}
}

func (sqlboiler *Sqlboiler) error(b *testing.B, method string, err string) {
	b.Helper()

	sqlboiler.mu.Lock()
	sqlboiler.errors[method] = err
	sqlboiler.mu.Unlock()
	b.Fail()
}
