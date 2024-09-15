package bench

import (
	"database/sql"
	"testing"

	"github.com/efectn/go-orm-benchmarks/helper"

	r "github.com/efectn/go-orm-benchmarks/bench/reform"
	_ "github.com/jackc/pgx/v5/stdlib"
	"gopkg.in/reform.v1/dialects/postgresql"

	reformware "gopkg.in/reform.v1"
)

type Reform struct {
	helper.ORMInterface
	conn *reformware.DB
	db   *sql.DB
}

func CreateReform() helper.ORMInterface {
	return &Reform{}
}

func (reform *Reform) Name() string {
	return "reform"
}

func (reform *Reform) Init() error {
	var err error
	reform.db, err = sql.Open("pgx", helper.OrmSource)
	if err != nil {
		return err
	}

	reform.conn = reformware.NewDB(reform.db, postgresql.Dialect, nil)

	return nil
}

func (reform *Reform) Close() error {
	return reform.db.Close()
}

func (reform *Reform) Insert(b *testing.B) {
	m := NewReformModel()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := reform.conn.Save(m)
		if err != nil {
			helper.SetError(b, reform.Name(), "Insert", err.Error())
		}
	}
}

func (reform *Reform) InsertMulti(b *testing.B) {
	ms := make([]reformware.Struct, 0, 100)
	for i := 0; i < 100; i++ {
		ms = append(ms, NewReformModel())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := reform.conn.InsertMulti(ms...)
		if err != nil {
			helper.SetError(b, reform.Name(), "InsertMulti", err.Error())
		}
	}
}

func (reform *Reform) Update(b *testing.B) {
	m := NewReformModel()

	err := reform.conn.Save(m)
	if err != nil {
		helper.SetError(b, reform.Name(), "Update", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := reform.conn.Update(m)
		if err != nil {
			helper.SetError(b, reform.Name(), "Update", err.Error())
		}
	}
}

func (reform *Reform) Read(b *testing.B) {
	m := NewReformModel()

	err := reform.conn.Save(m)
	if err != nil {
		helper.SetError(b, reform.Name(), "Read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := reform.conn.FindByPrimaryKeyFrom(r.ReformModelsTable, m.ID)
		if err != nil {
			helper.SetError(b, reform.Name(), "Read", err.Error())
		}
	}
}

func (reform *Reform) ReadSlice(b *testing.B) {
	m := NewReformModel()
	for i := 0; i < 100; i++ {
		err := reform.conn.Save(m)
		if err != nil {
			helper.SetError(b, reform.Name(), "ReadSlice", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := reform.conn.SelectAllFrom(r.ReformModelsTable, "WHERE id > 0 LIMIT 100")
		if err != nil {
			helper.SetError(b, reform.Name(), "ReadSlice", err.Error())
		}
	}
}
