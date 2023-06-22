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

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Id = 0
		_, err := pg.conn.Model(m).Insert()
		if err != nil {
			helper.SetError(b, pg.Name(), "Insert", err.Error())
		}
	}
}

func (pg *Pg) InsertMulti(b *testing.B) {
	ms := make([]*Model, 0, 100)
	for i := 0; i < 100; i++ {
		ms = append(ms, NewModel())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, m := range ms {
			m.Id = 0
		}

		_, err := pg.conn.Model(&ms).Insert()
		if err != nil {
			helper.SetError(b, pg.Name(), "InsertMulti", err.Error())
		}
	}
}

func (pg *Pg) Update(b *testing.B) {
	m := NewModel()

	_, err := pg.conn.Model(m).Insert()
	if err != nil {
		helper.SetError(b, pg.Name(), "Update", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := pg.conn.Model(m).WherePK().Update()
		if err != nil {
			helper.SetError(b, pg.Name(), "Update", err.Error())
		}
	}
}

func (pg *Pg) Read(b *testing.B) {
	m := NewModel()

	_, err := pg.conn.Model(m).Insert()
	if err != nil {
		helper.SetError(b, pg.Name(), "Read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := pg.conn.Model(m).Limit(1).Select()
		if err != nil {
			helper.SetError(b, pg.Name(), "Read", err.Error())
		}
	}
}

func (pg *Pg) ReadSlice(b *testing.B) {
	m := NewModel()
	for i := 0; i < 100; i++ {
		m.Id = 0
		_, err := pg.conn.Model(m).Insert()
		if err != nil {
			helper.SetError(b, pg.Name(), "ReadSlice", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var models []*Model
		err := pg.conn.Model(&models).Where("id > ?", 0).Limit(100).Select()
		if err != nil {
			helper.SetError(b, pg.Name(), "ReadSlice", err.Error())
		}
	}
}
