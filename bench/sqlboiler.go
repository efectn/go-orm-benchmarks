package bench

import (
	"database/sql"
	"testing"

	"github.com/efectn/go-orm-benchmarks/helper"

	models "github.com/efectn/go-orm-benchmarks/bench/sqlboiler"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Sqlboiler struct {
	helper.ORMInterface
	conn *sql.DB
}

func CreateSqlboiler() helper.ORMInterface {
	return &Sqlboiler{}
}

func (sqlboiler *Sqlboiler) Name() string {
	return "sqlboiler"
}

func (sqlboiler *Sqlboiler) Init() error {
	var err error
	sqlboiler.conn, err = sql.Open("pgx", helper.OrmSource)
	if err != nil {
		return err
	}

	boil.SetDB(sqlboiler.conn)

	return nil
}

func (sqlboiler *Sqlboiler) Close() error {
	return sqlboiler.conn.Close()
}

func (sqlboiler *Sqlboiler) Insert(b *testing.B) {
	m := NewModel6()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.ID = 0
		err := m.Insert(ctx, sqlboiler.conn, boil.Infer())
		if err != nil {
			helper.SetError(b, sqlboiler.Name(), "Insert", err.Error())
		}
	}
}

func (sqlboiler *Sqlboiler) InsertMulti(b *testing.B) {
	helper.SetError(b, sqlboiler.Name(), "InsertMulti", "bulk-insert is not supported")
}

func (sqlboiler *Sqlboiler) Update(b *testing.B) {
	m := NewModel6()
	m.ID = 0
	err := m.Insert(ctx, sqlboiler.conn, boil.Infer())
	if err != nil {
		helper.SetError(b, sqlboiler.Name(), "Update", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := m.Update(ctx, sqlboiler.conn, boil.Infer())
		if err != nil {
			helper.SetError(b, sqlboiler.Name(), "Update", err.Error())
		}
	}
}

func (sqlboiler *Sqlboiler) Read(b *testing.B) {
	m := NewModel6()
	m.ID = 0
	err := m.Insert(ctx, sqlboiler.conn, boil.Infer())
	if err != nil {
		helper.SetError(b, sqlboiler.Name(), "Read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := models.Models(qm.Where("id = 0")).Exec(sqlboiler.conn)
		if err != nil {
			helper.SetError(b, sqlboiler.Name(), "Read", err.Error())
		}
	}
}

func (sqlboiler *Sqlboiler) ReadSlice(b *testing.B) {
	m := NewModel6()
	for i := 0; i < 100; i++ {
		m.ID = 0
		err := m.Insert(ctx, sqlboiler.conn, boil.Infer())
		if err != nil {
			helper.SetError(b, sqlboiler.Name(), "ReadSlice", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := models.Models(qm.Where("id > 0"), qm.Limit(100)).All(ctx, sqlboiler.conn)
		if err != nil {
			helper.SetError(b, sqlboiler.Name(), "ReadSlice", err.Error())
		}
	}
}
