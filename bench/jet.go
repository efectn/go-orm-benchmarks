package bench

import (
	"database/sql"
	"testing"

	"github.com/efectn/go-orm-benchmarks/helper"

	"github.com/efectn/go-orm-benchmarks/bench/jet/test/public/model"
	. "github.com/efectn/go-orm-benchmarks/bench/jet/test/public/table"
	jetpg "github.com/go-jet/jet/v2/postgres"
)

type Jet struct {
	helper.ORMInterface
	conn *sql.DB
}

func CreateJet() helper.ORMInterface {
	return &Jet{}
}

func (jet *Jet) Name() string {
	return "jet"
}

func (jet *Jet) Init() error {
	var err error
	jet.conn, err = sql.Open("pgx", helper.OrmSource)
	if err != nil {
		return err
	}

	return nil
}

func (jet *Jet) Close() error {
	return jet.conn.Close()
}

func (jet *Jet) Insert(b *testing.B) {
	m := NewModelJet()
	helper.BenchExec(b, jet.Name(), "Insert", func() error {
		_, err := Models.INSERT(Models.MutableColumns).MODEL(m).Exec(jet.conn)
		return err
	})
}

func (jet *Jet) InsertMulti(b *testing.B) {
	ms := make([]*model.Models, 0, helper.BatchSize)
	for i := 0; i < helper.BatchSize; i++ {
		ms = append(ms, NewModelJet())
	}

	helper.BenchExec(b, jet.Name(), "InsertMulti", func() error {
		_, err := Models.INSERT(Models.MutableColumns).MODELS(ms).Exec(jet.conn)
		return err
	})
}

func (jet *Jet) Update(b *testing.B) {
	m, err := jet.create()
	if err != nil {
		helper.SetError(b, jet.Name(), "Update", err.Error())
		return
	}
	m.Counter = helper.UpdatedCounter

	helper.BenchExec(b, jet.Name(), "Update", func() error {
		_, err := Models.UPDATE(Models.MutableColumns).MODEL(m).WHERE(Models.ID.EQ(jetpg.Int32(m.ID))).Exec(jet.conn)
		return err
	})
}

func (jet *Jet) Read(b *testing.B) {
	m, err := jet.create()
	if err != nil {
		helper.SetError(b, jet.Name(), "Read", err.Error())
		return
	}

	helper.BenchRows(b, jet.Name(), "Read", 1, func() (int, error) {
		var mout model.Models
		err := Models.SELECT(Models.AllColumns).
			WHERE(Models.ID.EQ(jetpg.Int32(m.ID))).
			Query(jet.conn, &mout)
		if err != nil {
			return 0, err
		}
		return 1, nil
	})
}

func (jet *Jet) ReadSlice(b *testing.B) {
	for i := 0; i < helper.BatchSize; i++ {
		_, err := jet.create()
		if err != nil {
			helper.SetError(b, jet.Name(), "ReadSlice", err.Error())
			return
		}
	}

	helper.BenchRows(b, jet.Name(), "ReadSlice", helper.BatchSize, func() (int, error) {
		var models []model.Models
		err := Models.SELECT(Models.AllColumns).WHERE(Models.ID.GT(jetpg.Int32(0))).LIMIT(int64(helper.BatchSize)).Query(jet.conn, &models)
		if err != nil {
			return 0, err
		}
		return len(models), nil
	})
}

func (jet *Jet) create() (*model.Models, error) {
	m := NewModelJet()

	var created []model.Models
	err := Models.INSERT(Models.MutableColumns).MODEL(m).RETURNING(Models.ID).Query(jet.conn, &created)
	if err != nil {
		return nil, err
	}
	m.ID = created[0].ID
	return m, nil
}
