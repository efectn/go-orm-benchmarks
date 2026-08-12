package bench

import (
	"database/sql"
	"testing"

	"github.com/efectn/go-orm-benchmarks/helper"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
	gorpware "gopkg.in/gorp.v1"
)

type Gorp struct {
	helper.ORMInterface
	conn *gorpware.DbMap
}

func CreateGorp() helper.ORMInterface {
	return &Gorp{}
}

func (gorp *Gorp) Name() string {
	return "gorp"
}

func (gorp *Gorp) Init() error {
	db, err := sql.Open("pgx", helper.OrmSource)
	if err != nil {
		return err
	}

	gorp.conn = &gorpware.DbMap{Db: db, Dialect: gorpware.PostgresDialect{}}
	gorp.conn.AddTableWithName(Model{}, "models").SetKeys(true, "Id")

	return nil
}

func (gorp *Gorp) Close() error {
	return gorp.conn.Db.Close()
}

func (gorp *Gorp) Insert(b *testing.B) {
	m := NewModel()
	helper.BenchExec(b, gorp.Name(), "Insert", func() error {
		m.Id = 0
		return gorp.conn.Insert(m)
	})
}

func (gorp *Gorp) InsertMulti(b *testing.B) {
	helper.SetError(b, gorp.Name(), "InsertMulti", "bulk-insert is not supported")
}

func (gorp *Gorp) Update(b *testing.B) {
	m := NewModel()

	err := gorp.conn.Insert(m)
	if err != nil {
		helper.SetError(b, gorp.Name(), "Update", err.Error())
		return
	}
	m.Counter = helper.UpdatedCounter

	helper.BenchExec(b, gorp.Name(), "Update", func() error {
		_, err := gorp.conn.Update(m)
		return err
	})
}

func (gorp *Gorp) Read(b *testing.B) {
	m := NewModel()

	err := gorp.conn.Insert(m)
	if err != nil {
		helper.SetError(b, gorp.Name(), "Read", err.Error())
		return
	}

	helper.BenchRows(b, gorp.Name(), "Read", 1, func() (int, error) {
		if err := gorp.conn.SelectOne(m, "SELECT * FROM models WHERE id = $1", m.Id); err != nil {
			return 0, err
		}
		return 1, nil
	})
}

func (gorp *Gorp) ReadSlice(b *testing.B) {
	m := NewModel()
	for i := 0; i < helper.BatchSize; i++ {
		m.Id = 0
		err := gorp.conn.Insert(m)
		if err != nil {
			helper.SetError(b, gorp.Name(), "ReadSlice", err.Error())
			return
		}
	}

	helper.BenchRows(b, gorp.Name(), "ReadSlice", helper.BatchSize, func() (int, error) {
		var ms []*Model
		_, err := gorp.conn.Select(&ms, "select * from models where id > 0 LIMIT $1", helper.BatchSize)
		return len(ms), err
	})
}
