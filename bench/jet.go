package bench

import (
	"database/sql"
	"github.com/efectn/go-orm-benchmarks/helper"

	"github.com/efectn/go-orm-benchmarks/bench/jet/test/public/model"
	. "github.com/efectn/go-orm-benchmarks/bench/jet/test/public/table"
	jetpg "github.com/go-jet/jet/v2/postgres"

	"testing"
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

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := Models.INSERT(Models.MutableColumns).MODEL(m).Exec(jet.conn)
		if err != nil {
			helper.SetError(b, jet.Name(), "Insert", err.Error())
		}
	}
}

func (jet *Jet) InsertMulti(b *testing.B) {
	ms := make([]*model.Models, 0, 100)
	for i := 0; i < 100; i++ {
		ms = append(ms, NewModelJet())
	}

	b.ReportAllocs()
	b.ResetTimer()

	bulk := make([]*model.Models, len(ms))
	for i, m := range ms {
		bulk[i] = m
	}

	for i := 0; i < b.N; i++ {
		_, err := Models.INSERT(Models.MutableColumns).MODELS(ms).Exec(jet.conn)
		if err != nil {
			helper.SetError(b, jet.Name(), "InsertMulti", err.Error())
		}
	}
}

func (jet *Jet) Update(b *testing.B) {
	m, err := jet.create()
	if err != nil {
		helper.SetError(b, jet.Name(), "Update", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := Models.UPDATE(Models.MutableColumns).MODEL(m).WHERE(Models.ID.EQ(jetpg.Int32(m.ID))).Exec(jet.conn)
		if err != nil {
			helper.SetError(b, jet.Name(), "Update", err.Error())
		}
	}
}

func (jet *Jet) Read(b *testing.B) {
	m, err := jet.create()
	if err != nil {
		helper.SetError(b, jet.Name(), "Read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var mout model.Models
		err := Models.SELECT(Models.AllColumns).
			WHERE(Models.ID.EQ(jetpg.Int32(m.ID))).
			Query(jet.conn, &mout)
		if err != nil {
			helper.SetError(b, jet.Name(), "Read", err.Error())
		}
	}
}

func (jet *Jet) ReadSlice(b *testing.B) {
	for i := 0; i < 100; i++ {
		_, err := jet.create()
		if err != nil {
			helper.SetError(b, jet.Name(), "ReadSlice", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		models := make([]model.Models, 100)
		err := Models.SELECT(Models.AllColumns).WHERE(Models.ID.GT(jetpg.Int32(0))).LIMIT(100).Query(jet.conn, &models)
		if err != nil {
			helper.SetError(b, jet.Name(), "ReadSlice", err.Error())
		}
	}
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
