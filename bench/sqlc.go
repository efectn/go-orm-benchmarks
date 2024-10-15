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

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Id = 0
		err := sqlc.conn.CreateModel(ctx, args)
		if err != nil {
			helper.SetError(b, sqlc.Name(), "Insert", err.Error())
		}
	}
}

func (sqlc *Sqlc) InsertMulti(b *testing.B) {
	ms := make([]db.InsertMultiParams, 0, 100)
	m := NewModel()
	for i := 0; i < 100; i++ {
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

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := sqlc.conn.InsertMulti(ctx, ms)
		if err != nil {
			helper.SetError(b, sqlc.Name(), "InsertMulti", err.Error())
		}
	}
}

func (sqlc *Sqlc) Update(b *testing.B) {
	m := NewModel()
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
		helper.SetError(b, sqlc.Name(), "Update", err.Error())
	}

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
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := sqlc.conn.UpdateModel(ctx, args)
		if err != nil {
			helper.SetError(b, sqlc.Name(), "Update", err.Error())
		}
	}
}

func (sqlc *Sqlc) Read(b *testing.B) {
	m := NewModel()

	err := sqlc.conn.CreateModel(ctx, db.CreateModelParams{
		Name:    m.Name,
		Title:   m.Title,
		Fax:     m.Fax,
		Web:     m.Web,
		Age:     int32(m.Age),
		Right:   m.Right,
		Counter: m.Counter,
	})
	m.Id = 1
	if err != nil {
		helper.SetError(b, sqlc.Name(), "Read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := sqlc.conn.GetModel(ctx, int32(m.Id))
		if err != nil {
			helper.SetError(b, sqlc.Name(), "Read", err.Error())
		}
	}
}

func (sqlc *Sqlc) ReadSlice(b *testing.B) {
	m := NewModel()

	for i := 0; i < 100; i++ {
		m.Id = 0

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
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := sqlc.conn.ListModels(ctx)
		if err != nil {
			helper.SetError(b, sqlc.Name(), "ReadSlice", err.Error())
		}
	}
}
