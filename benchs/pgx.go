package benchs

import (
	"fmt"
	pgxdb "github.com/jackc/pgx/v4"
	"sync"
	"testing"
)

const (
	sqlxInsertBaseSQL   = `INSERT INTO models (name, title, fax, web, age, "right", counter) VALUES `
	sqlxInsertValuesSQL = `($1, $2, $3, $4, $5, $6, $7)`
	sqlxInsertSQL       = sqlxInsertBaseSQL + sqlxInsertValuesSQL
	sqlxInsertNamesSQL  = `(:name, :title, :fax, :web, :age, :right, :counter)`
	sqlxInsertMultiSQL  = sqlxInsertBaseSQL + sqlxInsertNamesSQL
	sqlxUpdateSQL       = `UPDATE models SET name = $1, title = $2, fax = $3, web = $4, age = $5, "right" = $6, counter = $7 WHERE id = $8`
	sqlxSelectSQL       = `SELECT * FROM models WHERE id = $1`
	sqlxSelectMultiSQL  = `SELECT * FROM models WHERE id > 0 LIMIT 100`
)

type Pgx struct {
	Instance
	mu         sync.Mutex
	conn       *pgxdb.Conn
	iterations int // Same as b.N, just to customize it
	errors     map[string]string
}

func CreatePgx(iterations int) Instance {
	pgx := &Pgx{
		iterations: iterations,
		errors:     map[string]string{},
	}

	return pgx
}

func (pgx *Pgx) Name() string {
	return "pgx"
}

func (pgx *Pgx) Init() error {
	var err error
	pgx.conn, err = pgxdb.Connect(ctx, OrmSource)
	if err != nil {
		return err
	}

	return nil
}

func (pgx *Pgx) Close() error {
	return pgx.conn.Close(ctx)
}

func (pgx *Pgx) GetError(method string) string {
	return pgx.errors[method]
}

func (pgx *Pgx) Insert(b *testing.B) {
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

func (pgx *Pgx) InsertMulti(b *testing.B) {
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

func (pgx *Pgx) Update(b *testing.B) {
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

func (pgx *Pgx) Read(b *testing.B) {
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

func (pgx *Pgx) ReadSlice(b *testing.B) {
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

func (pgx *Pgx) error(b *testing.B, method string, err string) {
	b.Helper()

	pgx.mu.Lock()
	pgx.errors[method] = err
	pgx.mu.Unlock()
	b.Fail()
}
