package bench

import (
	"database/sql"
	"testing"

	"github.com/efectn/go-orm-benchmarks/helper"
	sqlrlib "github.com/gandaldf/sqlr"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	sqlrInsertSQL          = `INSERT INTO models (name, title, fax, web, age, "right", counter) VALUES (:name, :title, :fax, :web, :age, :right, :counter)`
	sqlrInsertReturningSQL = sqlrInsertSQL + ` RETURNING id`
	sqlrInsertMultiSQL     = `INSERT INTO models (name, title, fax, web, age, "right", counter) VALUES :rows{name,title,fax,web,age,right,counter}`
	sqlrUpdateSQL          = `UPDATE models SET name = :name, title = :title, fax = :fax, web = :web, age = :age, "right" = :right, counter = :counter WHERE id = :id`
	sqlrSelectSQL          = `SELECT id, name, title, fax, web, age, "right", counter FROM models WHERE id = :id`
	sqlrSelectMultiSQL     = `SELECT id, name, title, fax, web, age, "right", counter FROM models WHERE id > 0 LIMIT 100`
)

type Sqlr struct {
	helper.ORMInterface
	conn *sql.DB
	lib  *sqlrlib.SQLR
}

var _ helper.ORMInterface = (*Sqlr)(nil)

func CreateSqlr() helper.ORMInterface {
	return &Sqlr{}
}

func (sqlr *Sqlr) Name() string {
	return "sqlr"
}

func (sqlr *Sqlr) Init() error {
	var err error
	sqlr.conn, err = sql.Open("pgx", helper.OrmSource)
	if err != nil {
		return err
	}

	sqlr.conn.SetMaxOpenConns(helper.OrmMaxConn)
	sqlr.conn.SetMaxIdleConns(helper.OrmMaxIdle)
	sqlr.lib = sqlrlib.New(sqlrlib.Postgres)
	return nil
}

func (sqlr *Sqlr) Close() error {
	return sqlr.conn.Close()
}

func (sqlr *Sqlr) Insert(b *testing.B) {
	params := newSqlrParams()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := sqlr.lib.Write(sqlrInsertSQL).
			Bind(params).
			ExecContext(ctx, sqlr.conn)
		if err != nil {
			helper.SetError(b, sqlr.Name(), "Insert", err.Error())
		}
	}
}

func (sqlr *Sqlr) InsertMulti(b *testing.B) {
	rows := make([]sqlrlib.P, 100)
	for i := range rows {
		rows[i] = newSqlrParams()
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := sqlr.lib.Write(sqlrInsertMultiSQL).
			Bind("rows", rows).
			ExecContext(ctx, sqlr.conn)
		if err != nil {
			helper.SetError(b, sqlr.Name(), "InsertMulti", err.Error())
		}
	}
}

func (sqlr *Sqlr) Update(b *testing.B) {
	params := newSqlrParams()
	id, err := sqlr.insertRow(params)
	if err != nil {
		helper.SetError(b, sqlr.Name(), "Update", err.Error())
		return
	}
	params["id"] = id

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err = sqlr.lib.Write(sqlrUpdateSQL).
			Bind(params).
			ExecContext(ctx, sqlr.conn)
		if err != nil {
			helper.SetError(b, sqlr.Name(), "Update", err.Error())
		}
	}
}

func (sqlr *Sqlr) Read(b *testing.B) {
	params := newSqlrParams()
	id, err := sqlr.insertRow(params)
	if err != nil {
		helper.SetError(b, sqlr.Name(), "Read", err.Error())
		return
	}
	params["id"] = id

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var result struct {
			ID      int    `db:"id"`
			Name    string `db:"name"`
			Title   string `db:"title"`
			Fax     string `db:"fax"`
			Web     string `db:"web"`
			Age     int    `db:"age"`
			Right   bool   `db:"right"`
			Counter int64  `db:"counter"`
		}
		err = sqlr.lib.Write(sqlrSelectSQL).
			Bind(params).
			ScanOneContext(ctx, sqlr.conn, &result)
		if err != nil {
			helper.SetError(b, sqlr.Name(), "Read", err.Error())
		}
	}
}

func (sqlr *Sqlr) ReadSlice(b *testing.B) {
	params := newSqlrParams()
	for i := 0; i < 100; i++ {
		if _, err := sqlr.insertRow(params); err != nil {
			helper.SetError(b, sqlr.Name(), "ReadSlice", err.Error())
			return
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var rows []struct {
			ID      int    `db:"id"`
			Name    string `db:"name"`
			Title   string `db:"title"`
			Fax     string `db:"fax"`
			Web     string `db:"web"`
			Age     int    `db:"age"`
			Right   bool   `db:"right"`
			Counter int64  `db:"counter"`
		}
		err := sqlr.lib.Write(sqlrSelectMultiSQL).
			ScanAllContext(ctx, sqlr.conn, &rows)
		if err != nil {
			helper.SetError(b, sqlr.Name(), "ReadSlice", err.Error())
		}
	}
}

func (sqlr *Sqlr) insertRow(params sqlrlib.P) (int, error) {
	var id int
	err := sqlr.lib.Write(sqlrInsertReturningSQL).
		Bind(params).
		ScanOneContext(ctx, sqlr.conn, &id)
	return id, err
}

func newSqlrParams() sqlrlib.P {
	model := NewModel()
	return sqlrlib.P{
		"name":    model.Name,
		"title":   model.Title,
		"fax":     model.Fax,
		"web":     model.Web,
		"age":     model.Age,
		"right":   model.Right,
		"counter": model.Counter,
	}
}
