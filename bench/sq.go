package bench

import (
	"database/sql"
	"testing"

	qm "github.com/bokwoon95/sq"
	"github.com/efectn/go-orm-benchmarks/bench/sq/db"
	"github.com/efectn/go-orm-benchmarks/helper"
)

// sq supports both templated queries and a query builder.
// This benchmark only uses the query builder.
type Sq struct {
	helper.ORMInterface
	conn *sql.DB
}

func CreateSq() helper.ORMInterface {
	return &Sq{}
}

func (sq *Sq) Name() string {
	return "sq"
}

func (sq *Sq) Init() error {
	var err error
	sq.conn, err = sql.Open("pgx", helper.OrmSource)
	if err != nil {
		sq.conn = nil
		return err
	}
	return nil
}

func (sq *Sq) Close() error {
	return sq.conn.Close()
}

func (sq *Sq) Insert(b *testing.B) {
	m := NewModelAlt()
	tbl := qm.New[db.MODELS]("")
	helper.BenchExec(b, sq.Name(), "Insert", func() error {
		query := qm.Postgres.InsertInto(tbl).
			Columns(tbl.NAME, tbl.TITLE, tbl.FAX, tbl.WEB, tbl.AGE, tbl.RIGHT, tbl.COUNTER).
			Values(m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		_, err := qm.Exec(sq.conn, query)
		return err
	})
}

func (sq *Sq) InsertMulti(b *testing.B) {
	ms := make([]Model, 0, helper.BatchSize)
	for i := 0; i < helper.BatchSize; i++ {
		ms = append(ms, NewModelAlt())
	}

	helper.BenchExec(b, sq.Name(), "InsertMulti", func() error {
		tbl := qm.New[db.MODELS]("")
		query := qm.Postgres.InsertInto(tbl).Columns(tbl.NAME, tbl.TITLE, tbl.FAX, tbl.WEB, tbl.AGE, tbl.RIGHT, tbl.COUNTER)
		for _, m := range ms {
			query = query.Values(m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		}
		_, err := qm.Exec(sq.conn, query)
		return err
	})
}

func (sq *Sq) Update(b *testing.B) {
	m := NewModelAlt()

	err := sq.insertModel(&m)
	if err != nil {
		helper.SetError(b, sq.Name(), "Update", err.Error())
		return
	}
	m.Counter = helper.UpdatedCounter

	helper.BenchExec(b, sq.Name(), "Update", func() error {
		tbl := qm.New[db.MODELS]("")
		query := qm.Postgres.Update(tbl).Set(
			tbl.NAME.SetString(m.Name),
			tbl.TITLE.SetString(m.Title),
			tbl.FAX.SetString(m.Fax),
			tbl.WEB.SetString(m.Web),
			tbl.AGE.SetInt(m.Age),
			tbl.RIGHT.SetBool(m.Right),
			tbl.COUNTER.SetInt64(m.Counter),
		).Where(tbl.ID.EqInt(m.Id))
		_, err := qm.Exec(sq.conn, query)
		return err
	})
}

func (sq *Sq) Read(b *testing.B) {
	m := NewModelAlt()

	err := sq.insertModel(&m)
	if err != nil {
		helper.SetError(b, sq.Name(), "Read", err.Error())
		return
	}

	helper.BenchRows(b, sq.Name(), "Read", 1, func() (int, error) {
		tbl := qm.New[db.MODELS]("")
		query := qm.Postgres.From(tbl).Where(tbl.ID.EqInt(m.Id))
		_, err := qm.FetchOne(sq.conn, query, sq.modelRowMapper(tbl))
		if err != nil {
			return 0, err
		}
		return 1, nil
	})
}

func (sq *Sq) ReadSlice(b *testing.B) {
	m := NewModelAlt()

	for i := 0; i < helper.BatchSize; i++ {
		m.Id = 0
		err := sq.insertModel(&m)
		if err != nil {
			helper.SetError(b, sq.Name(), "ReadSlice", err.Error())
			return
		}
	}

	helper.BenchRows(b, sq.Name(), "ReadSlice", helper.BatchSize, func() (int, error) {
		tbl := qm.New[db.MODELS]("")
		query := qm.Postgres.From(tbl).Where(tbl.ID.GtInt(0)).Limit(helper.BatchSize)
		models, err := qm.FetchAll(sq.conn, query, sq.modelRowMapper(tbl))
		if err != nil {
			return 0, err
		}
		return len(models), nil
	})
}

func (sq *Sq) insertModel(m *Model) error {
	var err error
	tbl := qm.New[db.MODELS]("")
	query := qm.Postgres.InsertInto(tbl).
		Columns(tbl.NAME, tbl.TITLE, tbl.FAX, tbl.WEB, tbl.AGE, tbl.RIGHT, tbl.COUNTER).
		Values(m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
	m.Id, err = qm.FetchOne(sq.conn, query, func(r *qm.Row) int {
		return r.IntField(tbl.ID)
	})
	return err
}

func (sq *Sq) modelRowMapper(tbl db.MODELS) func(*qm.Row) Model {
	return func(r *qm.Row) Model {
		return Model{
			Id:      r.IntField(tbl.ID),
			Name:    r.StringField(tbl.NAME),
			Title:   r.StringField(tbl.TITLE),
			Fax:     r.StringField(tbl.FAX),
			Web:     r.StringField(tbl.WEB),
			Age:     r.IntField(tbl.AGE),
			Right:   r.BoolField(tbl.RIGHT),
			Counter: r.Int64Field(tbl.COUNTER),
		}
	}
}
