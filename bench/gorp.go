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

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := gorp.conn.Insert(m)
		if err != nil {
			helper.SetError(b, gorp.Name(), "Insert", err.Error())
		}
	}
}

func (gorp *Gorp) InsertMulti(b *testing.B) {
	helper.SetError(b, gorp.Name(), "InsertMulti", "bulk-insert is not supported")
}

func (gorp *Gorp) Update(b *testing.B) {
	m := NewModel()

	err := gorp.conn.Insert(m)
	if err != nil {
		helper.SetError(b, gorp.Name(), "Update", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := gorp.conn.Update(m)
		if err != nil {
			helper.SetError(b, gorp.Name(), "Update", err.Error())
		}
	}
}

func (gorp *Gorp) Read(b *testing.B) {
	m := NewModel()

	err := gorp.conn.Insert(m)
	if err != nil {
		helper.SetError(b, gorp.Name(), "Read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := gorp.conn.SelectOne(m, "SELECT * FROM models LIMIT 1")
		if err != nil {
			helper.SetError(b, gorp.Name(), "Read", err.Error())
		}
	}
}

func (gorp *Gorp) ReadSlice(b *testing.B) {
	m := NewModel()
	for i := 0; i < 100; i++ {
		err := gorp.conn.Insert(m)
		if err != nil {
			helper.SetError(b, gorp.Name(), "ReadSlice", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var ms []*Model
		_, err := gorp.conn.Select(&ms, "select * from models where id > 0 LIMIT 100")
		if err != nil {
			helper.SetError(b, gorp.Name(), "ReadSlice", err.Error())
		}
	}
}
