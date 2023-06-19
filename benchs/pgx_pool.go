package benchs

import (
	"fmt"
	pgxdb "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"sync"
	"testing"
)

type PgxPool struct {
	Instance
	mu         sync.Mutex
	conn       *pgxpool.Pool
	iterations int // Same as b.N, just to customize it
	errors     map[string]string
}

func CreatePgxPool(iterations int) Instance {
	pgx := &PgxPool{
		iterations: iterations,
		errors:     map[string]string{},
	}

	return pgx
}

func (pgx *PgxPool) Name() string {
	return "pgx_pool"
}

func (pgx *PgxPool) Init() error {
	var err error
	pgx.conn, err = pgxpool.Connect(ctx, OrmSource)
	if err != nil {
		return err
	}

	return nil
}

func (pgx *PgxPool) Close() error {
	pgx.conn.Close()

	return nil
}

func (pgx *PgxPool) GetError(method string) string {
	return pgx.errors[method]
}

func (pgx *PgxPool) Insert(b *testing.B) {
	m := NewModel()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := pgx.conn.Exec(ctx, sqlxInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		if err != nil {
			pgx.error(b, "insert", fmt.Sprintf("pgx: insert: %v", err))
		}
	}
}

func (pgx *PgxPool) InsertMulti(b *testing.B) {
	var rows = make([][]interface{}, 0)

	m := NewModel()
	for i := 0; i < 100; i++ {
		rows = append(rows, []interface{}{m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter})
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := pgx.conn.CopyFrom(ctx, pgxdb.Identifier{"models"}, columns, pgxdb.CopyFromRows(rows))
		if err != nil {
			pgx.error(b, "insert_multi", fmt.Sprintf("pgx: insert_multi: %v", err))
		}
	}
}

func (pgx *PgxPool) Update(b *testing.B) {
	m := NewModel()
	m.Id = 1

	_, err := pgx.conn.Exec(ctx, sqlxInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
	if err != nil {
		pgx.error(b, "update", fmt.Sprintf("pgx: update: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := pgx.conn.Exec(ctx, sqlxUpdateSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter, m.Id)
		if err != nil {
			pgx.error(b, "update", fmt.Sprintf("pgx: update: %v", err))
		}
	}
}

func (pgx *PgxPool) Read(b *testing.B) {
	m := NewModel()
	m.Id = 1

	_, err := pgx.conn.Exec(ctx, sqlxInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
	if err != nil {
		pgx.error(b, "update", fmt.Sprintf("pgx: update: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var m Model
		err := pgx.conn.QueryRow(ctx, sqlxSelectSQL, 1).Scan(
			&m.Id,
			&m.Name,
			&m.Title,
			&m.Fax,
			&m.Web,
			&m.Age,
			&m.Right,
			&m.Counter,
		)
		if err != nil {
			pgx.error(b, "pgx", fmt.Sprintf("pgx: update: %v", err))
		}
	}
}

func (pgx *PgxPool) ReadSlice(b *testing.B) {
	m := NewModel()
	for i := 0; i < 100; i++ {
		_, err := pgx.conn.Exec(ctx, sqlxInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		if err != nil {
			pgx.error(b, "read_slice", fmt.Sprintf("pgx: read_slice: %v", err))
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ms := make([]Model, 100)
		rows, err := pgx.conn.Query(ctx, sqlxSelectMultiSQL)
		if err != nil {
			pgx.error(b, "read_slice", fmt.Sprintf("pgx: read_slice: %v", err))
		}

		for j := 0; rows.Next() && j < len(ms); j++ {
			err = rows.Scan(
				&ms[j].Id,
				&ms[j].Name,
				&ms[j].Title,
				&ms[j].Fax,
				&ms[j].Web,
				&ms[j].Age,
				&ms[j].Right,
				&ms[j].Counter,
			)
			if err != nil {
				pgx.error(b, "read_slice", fmt.Sprintf("pgx: read_slice: %v", err))
			}
		}
		err = rows.Err()
		if err != nil {
			pgx.error(b, "read_slice", fmt.Sprintf("pgx: read_slice: %v", err))
		}

		rows.Close()
	}
}

func (pgx *PgxPool) error(b *testing.B, method string, err string) {
	b.Helper()

	pgx.mu.Lock()
	pgx.errors[method] = err
	pgx.mu.Unlock()
	b.Fail()
}
