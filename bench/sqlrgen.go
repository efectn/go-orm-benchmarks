package bench

import (
	"database/sql"
	"testing"

	"github.com/efectn/go-orm-benchmarks/helper"
	sqlrlib "github.com/gandaldf/sqlr"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	sqlrgenInsertSQL          = `INSERT INTO models (name, title, fax, web, age, "right", counter) VALUES (:name, :title, :fax, :web, :age, :right, :counter)`
	sqlrgenInsertReturningSQL = sqlrgenInsertSQL + ` RETURNING id`
	sqlrgenInsertMultiSQL     = `INSERT INTO models (name, title, fax, web, age, "right", counter) VALUES :rows{name,title,fax,web,age,right,counter}`
	sqlrgenUpdateSQL          = `UPDATE models SET name = :name, title = :title, fax = :fax, web = :web, age = :age, "right" = :right, counter = :counter WHERE id = :id`
	sqlrgenSelectSQL          = `SELECT id, name, title, fax, web, age, "right", counter FROM models WHERE id = :id`
	sqlrgenSelectMultiSQL     = `SELECT id, name, title, fax, web, age, "right", counter FROM models WHERE id > 0 LIMIT :limit`
)

//go:generate go run github.com/gandaldf/sqlr/cmd/sqlrgen -type SqlrgenModel -output sqlrgen_model.go

type SqlrgenModel struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Title   string `db:"title"`
	Fax     string `db:"fax"`
	Web     string `db:"web"`
	Age     int    `db:"age"`
	Right   bool   `db:"right"`
	Counter int64  `db:"counter"`
}

type Sqlrgen struct {
	helper.ORMInterface
	conn *sql.DB
	lib  *sqlrlib.SQLR
}

var _ helper.ORMInterface = (*Sqlrgen)(nil)

func CreateSqlrgen() helper.ORMInterface {
	return &Sqlrgen{}
}

func (sqlrgen *Sqlrgen) Name() string {
	return "sqlrgen"
}

func (sqlrgen *Sqlrgen) Init() error {
	var err error
	sqlrgen.conn, err = sql.Open("pgx", helper.OrmSource)
	if err != nil {
		return err
	}

	sqlrgen.conn.SetMaxOpenConns(helper.OrmMaxConn)
	sqlrgen.conn.SetMaxIdleConns(helper.OrmMaxIdle)
	sqlrgen.lib = sqlrlib.New(sqlrlib.Postgres)
	return nil
}

func (sqlrgen *Sqlrgen) Close() error {
	return sqlrgen.conn.Close()
}

func (sqlrgen *Sqlrgen) Insert(b *testing.B) {
	model := newSqlrgenModel()
	helper.BenchExec(b, sqlrgen.Name(), "Insert", func() error {
		_, err := sqlrgen.lib.Write(sqlrgenInsertSQL).
			Bind(model).
			ExecContext(ctx, sqlrgen.conn)
		return err
	})
}

func (sqlrgen *Sqlrgen) InsertMulti(b *testing.B) {
	models := make([]SqlrgenModel, helper.BatchSize)
	for i := range models {
		models[i] = *newSqlrgenModel()
	}

	helper.BenchExec(b, sqlrgen.Name(), "InsertMulti", func() error {
		_, err := sqlrgen.lib.Write(sqlrgenInsertMultiSQL).
			Bind("rows", SQLRRowsOfSqlrgenModel(models)).
			ExecContext(ctx, sqlrgen.conn)
		return err
	})
}

func (sqlrgen *Sqlrgen) Update(b *testing.B) {
	model := newSqlrgenModel()
	if err := sqlrgen.insertModel(model); err != nil {
		helper.SetError(b, sqlrgen.Name(), "Update", err.Error())
		return
	}
	model.Counter = helper.UpdatedCounter

	helper.BenchExec(b, sqlrgen.Name(), "Update", func() error {
		_, err := sqlrgen.lib.Write(sqlrgenUpdateSQL).
			Bind(model).
			ExecContext(ctx, sqlrgen.conn)
		return err
	})
}

func (sqlrgen *Sqlrgen) Read(b *testing.B) {
	model := newSqlrgenModel()
	if err := sqlrgen.insertModel(model); err != nil {
		helper.SetError(b, sqlrgen.Name(), "Read", err.Error())
		return
	}

	helper.BenchRows(b, sqlrgen.Name(), "Read", 1, func() (int, error) {
		var result SqlrgenModel
		err := sqlrgen.lib.Write(sqlrgenSelectSQL).
			Bind(model).
			ScanOneContext(ctx, sqlrgen.conn, &result)
		if err != nil {
			return 0, err
		}
		return 1, nil
	})
}

func (sqlrgen *Sqlrgen) ReadSlice(b *testing.B) {
	for i := 0; i < helper.BatchSize; i++ {
		if err := sqlrgen.insertModel(newSqlrgenModel()); err != nil {
			helper.SetError(b, sqlrgen.Name(), "ReadSlice", err.Error())
			return
		}
	}

	helper.BenchRows(b, sqlrgen.Name(), "ReadSlice", helper.BatchSize, func() (int, error) {
		var models []SqlrgenModel
		err := sqlrgen.lib.Write(sqlrgenSelectMultiSQL).
			Bind("limit", helper.BatchSize).
			ScanAllContext(ctx, sqlrgen.conn, &models)
		return len(models), err
	})
}

func (sqlrgen *Sqlrgen) insertModel(model *SqlrgenModel) error {
	return sqlrgen.lib.Write(sqlrgenInsertReturningSQL).
		Bind(model).
		ScanOneContext(ctx, sqlrgen.conn, &model.ID)
}

func newSqlrgenModel() *SqlrgenModel {
	model := NewModel()
	return &SqlrgenModel{
		ID:      model.Id,
		Name:    model.Name,
		Title:   model.Title,
		Fax:     model.Fax,
		Web:     model.Web,
		Age:     model.Age,
		Right:   model.Right,
		Counter: model.Counter,
	}
}
