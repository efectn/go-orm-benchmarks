package bench

import (
	"database/sql"
	"testing"

	"github.com/efectn/go-orm-benchmarks/helper"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/aarondl/sqlboiler/v4/queries/qm"
	models "github.com/efectn/go-orm-benchmarks/bench/sqlboiler"
	_ "github.com/jackc/pgx/v5/stdlib"
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
	helper.BenchExec(b, sqlboiler.Name(), "Insert", func() error {
		m.ID = 0
		return m.Insert(ctx, sqlboiler.conn, boil.Infer())
	})
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
		return
	}
	m.Counter = helper.UpdatedCounter

	helper.BenchExec(b, sqlboiler.Name(), "Update", func() error {
		_, err := m.Update(ctx, sqlboiler.conn, boil.Infer())
		return err
	})
}

func (sqlboiler *Sqlboiler) Read(b *testing.B) {
	m := NewModel6()
	m.ID = 0
	err := m.Insert(ctx, sqlboiler.conn, boil.Infer())
	if err != nil {
		helper.SetError(b, sqlboiler.Name(), "Read", err.Error())
		return
	}

	helper.BenchRows(b, sqlboiler.Name(), "Read", 1, func() (int, error) {
		_, err := models.FindModel(ctx, sqlboiler.conn, m.ID)
		if err != nil {
			return 0, err
		}
		return 1, nil
	})
}

func (sqlboiler *Sqlboiler) ReadSlice(b *testing.B) {
	m := NewModel6()
	for i := 0; i < helper.BatchSize; i++ {
		m.ID = 0
		err := m.Insert(ctx, sqlboiler.conn, boil.Infer())
		if err != nil {
			helper.SetError(b, sqlboiler.Name(), "ReadSlice", err.Error())
			return
		}
	}

	helper.BenchRows(b, sqlboiler.Name(), "ReadSlice", helper.BatchSize, func() (int, error) {
		rows, err := models.Models(qm.Where("id > 0"), qm.Limit(helper.BatchSize)).All(ctx, sqlboiler.conn)
		if err != nil {
			return 0, err
		}
		return len(rows), nil
	})
}
