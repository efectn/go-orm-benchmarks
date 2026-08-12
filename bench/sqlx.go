package bench

import (
	"github.com/efectn/go-orm-benchmarks/helper"
	sqlxdb "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"testing"
)

const (
	sqlxInsertBaseSQL      = `INSERT INTO models (name, title, fax, web, age, "right", counter) VALUES `
	sqlxInsertValuesSQL    = `($1, $2, $3, $4, $5, $6, $7)`
	sqlxInsertSQL          = sqlxInsertBaseSQL + sqlxInsertValuesSQL
	sqlxInsertReturningSQL = sqlxInsertSQL + ` RETURNING id`
	sqlxInsertNamesSQL     = `(:name, :title, :fax, :web, :age, :right, :counter)`
	sqlxInsertMultiSQL     = sqlxInsertBaseSQL + sqlxInsertNamesSQL
	sqlxUpdateSQL          = `UPDATE models SET name = $1, title = $2, fax = $3, web = $4, age = $5, "right" = $6, counter = $7 WHERE id = $8`
	sqlxSelectSQL          = `SELECT * FROM models WHERE id = $1`
	sqlxSelectMultiSQL     = `SELECT * FROM models WHERE id > 0 LIMIT $1`
)

type Sqlx struct {
	helper.ORMInterface
	conn *sqlxdb.DB
}

func CreateSqlx() helper.ORMInterface {
	return &Sqlx{}
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
	helper.BenchExec(b, sqlx.Name(), "Insert", func() error {
		_, err := sqlx.conn.Exec(sqlxInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		return err
	})
}

func (sqlx *Sqlx) InsertMulti(b *testing.B) {
	ms := make([]*Model, 0, helper.BatchSize)
	for i := 0; i < helper.BatchSize; i++ {
		ms = append(ms, NewModel())
	}
	helper.BenchExec(b, sqlx.Name(), "InsertMulti", func() error {
		_, err := sqlx.conn.NamedExec(sqlxInsertMultiSQL, ms)
		return err
	})
}

func (sqlx *Sqlx) Update(b *testing.B) {
	m := NewModel()
	err := sqlx.conn.QueryRow(sqlxInsertReturningSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter).Scan(&m.Id)
	if err != nil {
		helper.SetError(b, sqlx.Name(), "Update", err.Error())
		return
	}
	m.Counter = helper.UpdatedCounter

	helper.BenchExec(b, sqlx.Name(), "Update", func() error {
		_, err := sqlx.conn.Exec(sqlxUpdateSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter, m.Id)
		return err
	})
}

func (sqlx *Sqlx) Read(b *testing.B) {
	m := NewModel()
	err := sqlx.conn.QueryRow(sqlxInsertReturningSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter).Scan(&m.Id)
	if err != nil {
		helper.SetError(b, sqlx.Name(), "Read", err.Error())
		return
	}

	helper.BenchRows(b, sqlx.Name(), "Read", 1, func() (int, error) {
		var mout Model
		err := sqlx.conn.Get(&mout, sqlxSelectSQL, m.Id)
		if err != nil {
			return 0, err
		}
		return 1, nil
	})
}

func (sqlx *Sqlx) ReadSlice(b *testing.B) {
	m := NewModel()
	for i := 0; i < helper.BatchSize; i++ {
		_, err := sqlx.conn.Exec(sqlxInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		if err != nil {
			helper.SetError(b, sqlx.Name(), "ReadSlice", err.Error())
			return
		}
	}

	helper.BenchRows(b, sqlx.Name(), "ReadSlice", helper.BatchSize, func() (int, error) {
		var ms []Model
		err := sqlx.conn.Select(&ms, sqlxSelectMultiSQL, helper.BatchSize)
		return len(ms), err
	})
}
