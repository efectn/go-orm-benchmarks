package bench

import (
	"github.com/efectn/go-orm-benchmarks/helper"
	pgdb "github.com/go-pg/pg/v10"
	"testing"
)

type Pg struct {
	helper.ORMInterface
	conn *pgdb.DB
}

func CreatePg() helper.ORMInterface {
	return &Pg{}
}

func (pg *Pg) Name() string {
	return "pg"
}

func (pg *Pg) Init() error {
	source := helper.SplitSource()
	pg.conn = pgdb.Connect(&pgdb.Options{
		Addr:     source["host"] + ":5432",
		User:     source["user"],
		Password: source["password"],
		Database: source["dbname"],
	})

	if err := pg.conn.Ping(ctx); err != nil {
		return err
	}

	return nil
}

func (pg *Pg) Close() error {
	return pg.conn.Close()
}

func (pg *Pg) Insert(b *testing.B) {
	m := NewModel()
	helper.BenchExec(b, pg.Name(), "Insert", func() error {
		m.Id = 0
		_, err := pg.conn.Model(m).Insert()
		return err
	})
}

func (pg *Pg) InsertMulti(b *testing.B) {
	ms := make([]*Model, 0, helper.BatchSize)
	for i := 0; i < helper.BatchSize; i++ {
		ms = append(ms, NewModel())
	}

	helper.BenchExec(b, pg.Name(), "InsertMulti", func() error {
		for _, m := range ms {
			m.Id = 0
		}

		_, err := pg.conn.Model(&ms).Insert()
		return err
	})
}

func (pg *Pg) Update(b *testing.B) {
	m := NewModel()

	_, err := pg.conn.Model(m).Insert()
	if err != nil {
		helper.SetError(b, pg.Name(), "Update", err.Error())
		return
	}
	m.Counter = helper.UpdatedCounter

	helper.BenchExec(b, pg.Name(), "Update", func() error {
		_, err := pg.conn.Model(m).WherePK().Update()
		return err
	})
}

func (pg *Pg) Read(b *testing.B) {
	m := NewModel()

	_, err := pg.conn.Model(m).Insert()
	if err != nil {
		helper.SetError(b, pg.Name(), "Read", err.Error())
		return
	}

	helper.BenchRows(b, pg.Name(), "Read", 1, func() (int, error) {
		if err := pg.conn.Model(m).WherePK().Select(); err != nil {
			return 0, err
		}
		return 1, nil
	})
}

func (pg *Pg) ReadSlice(b *testing.B) {
	m := NewModel()
	for i := 0; i < helper.BatchSize; i++ {
		m.Id = 0
		_, err := pg.conn.Model(m).Insert()
		if err != nil {
			helper.SetError(b, pg.Name(), "ReadSlice", err.Error())
			return
		}
	}

	helper.BenchRows(b, pg.Name(), "ReadSlice", helper.BatchSize, func() (int, error) {
		var models []*Model
		err := pg.conn.Model(&models).Where("id > ?", 0).Limit(helper.BatchSize).Select()
		return len(models), err
	})
}
