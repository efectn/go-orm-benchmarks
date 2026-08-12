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
	helper.BenchExec(b, reform.Name(), "Insert", func() error {
		m.ID = 0
		return reform.conn.Insert(m)
	})
}

func (reform *Reform) InsertMulti(b *testing.B) {
	ms := make([]reformware.Struct, 0, helper.BatchSize)
	for i := 0; i < helper.BatchSize; i++ {
		ms = append(ms, NewReformModel())
	}

	helper.BenchExec(b, reform.Name(), "InsertMulti", func() error {
		for _, model := range ms {
			model.(*r.ReformModels).ID = 0
		}
		return reform.conn.InsertMulti(ms...)
	})
}

func (reform *Reform) Update(b *testing.B) {
	m := NewReformModel()

	err := reform.conn.Insert(m)
	if err != nil {
		helper.SetError(b, reform.Name(), "Update", err.Error())
		return
	}
	m.Counter = helper.UpdatedCounter

	helper.BenchExec(b, reform.Name(), "Update", func() error {
		return reform.conn.Update(m)
	})
}

func (reform *Reform) Read(b *testing.B) {
	m := NewReformModel()

	err := reform.conn.Insert(m)
	if err != nil {
		helper.SetError(b, reform.Name(), "Read", err.Error())
		return
	}

	helper.BenchRows(b, reform.Name(), "Read", 1, func() (int, error) {
		_, err := reform.conn.FindByPrimaryKeyFrom(r.ReformModelsTable, m.ID)
		if err != nil {
			return 0, err
		}
		return 1, nil
	})
}

func (reform *Reform) ReadSlice(b *testing.B) {
	m := NewReformModel()
	for i := 0; i < helper.BatchSize; i++ {
		m.ID = 0
		err := reform.conn.Insert(m)
		if err != nil {
			helper.SetError(b, reform.Name(), "ReadSlice", err.Error())
			return
		}
	}

	helper.BenchRows(b, reform.Name(), "ReadSlice", helper.BatchSize, func() (int, error) {
		models, err := reform.conn.SelectAllFrom(r.ReformModelsTable, "WHERE id > 0 LIMIT $1", helper.BatchSize)
		if err != nil {
			return 0, err
		}
		return len(models), nil
	})
}
