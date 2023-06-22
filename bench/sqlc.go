package bench

import (
	"database/sql"
	"github.com/efectn/go-orm-benchmarks/helper"
	"testing"

	"github.com/efectn/go-orm-benchmarks/bench/sqlc/db"
)

type Sqlc struct {
	helper.ORMInterface
	conn *db.Queries
	db   *sql.DB
}

func CreateSqlc() helper.ORMInterface {
	return &Sqlc{}
}

func (sqlc *Sqlc) Name() string {
	return "sqlc"
}

func (sqlc *Sqlc) Init() error {
	var err error
	sqlc.db, err = sql.Open("pgx", helper.OrmSource)
	if err != nil {
		return err
	}

	sqlc.conn = db.New(sqlc.db)

	return nil
}

func (sqlc *Sqlc) Close() error {
	return sqlc.db.Close()
}

func (sqlc *Sqlc) Insert(b *testing.B) {
	m := NewModel()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Id = 0
		_, err := sqlc.conn.CreateModel(ctx, db.CreateModelParams{
			Name:    m.Name,
			Title:   m.Title,
			Fax:     m.Fax,
			Web:     m.Web,
			Age:     int32(m.Age),
			Right:   m.Right,
			Counter: m.Counter,
		})
		if err != nil {
			helper.SetError(b, sqlc.Name(), "Insert", err.Error())
		}
	}
}

func (sqlc *Sqlc) InsertMulti(b *testing.B) {
	helper.SetError(b, sqlc.Name(), "InsertMulti", "bulk-insert is not supported")
}

func (sqlc *Sqlc) Update(b *testing.B) {
	m := NewModel()

	_, err := sqlc.conn.CreateModel(ctx, db.CreateModelParams{
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

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := sqlc.conn.UpdateModel(ctx, db.UpdateModelParams{
			Name:    m.Name,
			Title:   m.Title,
			Fax:     m.Fax,
			Web:     m.Web,
			Age:     int32(m.Age),
			Right:   m.Right,
			Counter: m.Counter,
			ID:      int32(m.Id),
		})
		if err != nil {
			helper.SetError(b, sqlc.Name(), "Update", err.Error())
		}
	}
}

func (sqlc *Sqlc) Read(b *testing.B) {
	m := NewModel()

	output, err := sqlc.conn.CreateModel(ctx, db.CreateModelParams{
		Name:    m.Name,
		Title:   m.Title,
		Fax:     m.Fax,
		Web:     m.Web,
		Age:     int32(m.Age),
		Right:   m.Right,
		Counter: m.Counter,
	})
	m.Id = int(output.ID)
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

		_, err := sqlc.conn.CreateModel(ctx, db.CreateModelParams{
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
		_, err := sqlc.conn.ListModels(ctx, db.ListModelsParams{
			ID:    0,
			Limit: 100,
		})
		if err != nil {
			helper.SetError(b, sqlc.Name(), "ReadSlice", err.Error())
		}
	}
}
