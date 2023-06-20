package bench

import (
	"github.com/efectn/go-orm-benchmarks/helper"
	sqlxdb "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

type Sqlx struct {
	helper.ORMInterface
	mu         sync.Mutex
	conn       *sqlxdb.DB
	iterations int // Same as b.N, just to customize it
}

func CreateSqlx(iterations int) helper.ORMInterface {
	sqlx := &Sqlx{
		iterations: iterations,
	}

	return sqlx
}

func (sqlx *Sqlx) Name() string {
	return "sqlx"
}

func (sqlx *Sqlx) Init() error {
	var err error
	sqlx.conn, err = sqlxdb.Connect("postgres", helper.OrmSource)
	if err != nil {
		return err
	}

	return nil
}

func (sqlx *Sqlx) Close() error {
	return sqlx.conn.Close()
}

func (sqlx *Sqlx) Insert(b *testing.B) {
	m := NewModel()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := sqlx.conn.Exec(sqlxInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		if err != nil {
			helper.SetError(b, sqlx.Name(), "insert", err.Error())
		}
	}
}

func (sqlx *Sqlx) InsertMulti(b *testing.B) {
	ms := make([]*Model, 0, 100)
	for i := 0; i < 100; i++ {
		ms = append(ms, NewModel())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := sqlx.conn.NamedExec(sqlxInsertMultiSQL, ms)
		if err != nil {
			helper.SetError(b, sqlx.Name(), "insert_multi", err.Error())
		}
	}
}

func (sqlx *Sqlx) Update(b *testing.B) {
	m := NewModel()
	_, err := sqlx.conn.Exec(sqlxInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
	if err != nil {
		helper.SetError(b, sqlx.Name(), "update", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := sqlx.conn.Exec(sqlxUpdateSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter, m.Id)
		if err != nil {
			helper.SetError(b, sqlx.Name(), "update", err.Error())
		}
	}
}

func (sqlx *Sqlx) Read(b *testing.B) {
	m := NewModel()
	_, err := sqlx.conn.Exec(sqlxInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
	if err != nil {
		helper.SetError(b, sqlx.Name(), "read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var m []Model
		err := sqlx.conn.Select(&m, sqlxSelectSQL, 1)
		if err != nil {
			helper.SetError(b, sqlx.Name(), "read", err.Error())
		}
	}
}

func (sqlx *Sqlx) ReadSlice(b *testing.B) {
	m := NewModel()
	for i := 0; i < 100; i++ {
		_, err := sqlx.conn.Exec(sqlxInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		if err != nil {
			helper.SetError(b, sqlx.Name(), "read_slice", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ms := make([]Model, 100)
		err := sqlx.conn.Select(&ms, sqlxSelectMultiSQL)
		if err != nil {
			helper.SetError(b, sqlx.Name(), "read_slice", err.Error())
		}
	}
}
