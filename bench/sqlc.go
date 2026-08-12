package bench

import (
	"testing"

	"github.com/efectn/go-orm-benchmarks/helper"

	"github.com/efectn/go-orm-benchmarks/bench/sqlc/db"
	"github.com/jackc/pgx/v5"
)

type Sqlc struct {
	helper.ORMInterface
	conn *db.Queries
	db   *pgx.Conn
}

func CreateSqlc() helper.ORMInterface {
	return &Sqlc{}
}

func (sqlc *Sqlc) Name() string {
	return "sqlc"
}

func (sqlc *Sqlc) Init() error {
	var err error
	sqlc.db, err = pgx.Connect(ctx, helper.OrmSource)
	if err != nil {
		return err
	}

	sqlc.conn = db.New(sqlc.db)

	return nil
}

func (sqlc *Sqlc) Close() error {
	return sqlc.db.Close(ctx)
}

func (sqlc *Sqlc) Insert(b *testing.B) {
	m := NewModel()

	args := db.CreateModelParams{
		Name:    m.Name,
		Title:   m.Title,
		Fax:     m.Fax,
		Web:     m.Web,
		Age:     int32(m.Age),
		Right:   m.Right,
		Counter: m.Counter,
	}

	helper.BenchExec(b, sqlc.Name(), "Insert", func() error {
		return sqlc.conn.CreateModel(ctx, args)
	})
}

func (sqlc *Sqlc) InsertMulti(b *testing.B) {
	ms := make([]db.InsertMultiParams, 0, helper.BatchSize)
	m := NewModel()
	for i := 0; i < helper.BatchSize; i++ {
		ms = append(ms, db.InsertMultiParams{
			Name:    m.Name,
			Title:   m.Title,
			Fax:     m.Fax,
			Web:     m.Web,
			Age:     int32(m.Age),
			Right:   m.Right,
			Counter: m.Counter,
		})
	}

	helper.BenchExec(b, sqlc.Name(), "InsertMulti", func() error {
		_, err := sqlc.conn.InsertMulti(ctx, ms)
		return err
	})
}

func (sqlc *Sqlc) Update(b *testing.B) {
	m := NewModel()
	err := sqlc.createModelReturning(m)
	if err != nil {
		helper.SetError(b, sqlc.Name(), "Update", err.Error())
		return
	}
	m.Counter = helper.UpdatedCounter

	args := db.UpdateModelParams{
		Name:    m.Name,
		Title:   m.Title,
		Fax:     m.Fax,
		Web:     m.Web,
		Age:     int32(m.Age),
		Right:   m.Right,
		Counter: m.Counter,
		ID:      int32(m.Id),
	}
	helper.BenchExec(b, sqlc.Name(), "Update", func() error {
		return sqlc.conn.UpdateModel(ctx, args)
	})
}

func (sqlc *Sqlc) Read(b *testing.B) {
	m := NewModel()

	err := sqlc.createModelReturning(m)
	if err != nil {
		helper.SetError(b, sqlc.Name(), "Read", err.Error())
		return
	}

	helper.BenchRows(b, sqlc.Name(), "Read", 1, func() (int, error) {
		_, err := sqlc.conn.GetModel(ctx, int32(m.Id))
		if err != nil {
			return 0, err
		}
		return 1, nil
	})
}

func (sqlc *Sqlc) ReadSlice(b *testing.B) {
	m := NewModel()

	for i := 0; i < helper.BatchSize; i++ {
		err := sqlc.conn.CreateModel(ctx, db.CreateModelParams{
			Name:    m.Name,
			Title:   m.Title,
			Fax:     m.Fax,
			Web:     m.Web,
			Age:     int32(m.Age),
			Right:   m.Right,
			Counter: m.Counter,
		})
		if err != nil {
			helper.SetError(b, sqlc.Name(), "ReadSlice", err.Error())
			return
		}
	}

	helper.BenchRows(b, sqlc.Name(), "ReadSlice", helper.BatchSize, func() (int, error) {
		models, err := sqlc.conn.ListModels(ctx, int32(helper.BatchSize))
		return len(models), err
	})
}

func (sqlc *Sqlc) createModelReturning(m *Model) error {
	return sqlc.db.QueryRow(
		ctx,
		sqlxInsertReturningSQL,
		m.Name,
		m.Title,
		m.Fax,
		m.Web,
		m.Age,
		m.Right,
		m.Counter,
	).Scan(&m.Id)
}
