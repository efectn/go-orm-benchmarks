package bench

import (
	"testing"

	"github.com/efectn/go-orm-benchmarks/helper"

	dbrware "github.com/gocraft/dbr/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var columns = []string{"name", "title", "fax", "web", "age", "right", "counter"}

type Dbr struct {
	helper.ORMInterface
	conn *dbrware.Session
}

func CreateDbr() helper.ORMInterface {
	return &Dbr{}
}

func (dbr *Dbr) Name() string {
	return "dbr"
}

func (dbr *Dbr) Init() error {
	conn, err := dbrware.Open("postgres", helper.OrmSource, nil)
	if err != nil {
		return err
	}

	dbr.conn = conn.NewSession(nil)
	_, err = dbr.conn.Begin()
	if err != nil {
		return err
	}

	return nil
}

func (dbr *Dbr) Close() error {
	return dbr.conn.Close()
}

func (dbr *Dbr) Insert(b *testing.B) {
	m := NewModel2()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := dbr.conn.InsertInto("models").Columns(columns...).Record(m).Exec()
		if err != nil {
			helper.SetError(b, dbr.Name(), "Insert", err.Error())
		}
	}
}

func (dbr *Dbr) InsertMulti(b *testing.B) {
	helper.SetError(b, dbr.Name(), "InsertMulti", "bulk-insert is not supported")
}

func (dbr *Dbr) Update(b *testing.B) {
	m := NewModel2()

	_, err := dbr.conn.InsertInto("models").Columns(columns...).Record(m).Exec()
	if err != nil {
		helper.SetError(b, dbr.Name(), "Update", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := dbr.conn.Update("models").SetMap(map[string]any{
			"name":    m.Name,
			"title":   m.Title,
			"fax":     m.Fax,
			"web":     m.Web,
			"age":     m.Age,
			"right":   m.Right,
			"counter": m.Counter,
		}).Exec()
		if err != nil {
			helper.SetError(b, dbr.Name(), "Update", err.Error())
		}
	}
}

func (dbr *Dbr) Read(b *testing.B) {
	m := NewModel2()

	_, err := dbr.conn.InsertInto("models").Columns(columns...).Record(m).Exec()
	if err != nil {
		helper.SetError(b, dbr.Name(), "Read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := dbr.conn.Select("*").From("models").Where("id = ?", m.ID).Load(m)
		if err != nil {
			helper.SetError(b, dbr.Name(), "Read", err.Error())
		}
	}
}

func (dbr *Dbr) ReadSlice(b *testing.B) {
	m := NewModel2()
	for i := 0; i < 100; i++ {
		_, err := dbr.conn.InsertInto("models").Columns(columns...).Record(m).Exec()
		if err != nil {
			helper.SetError(b, dbr.Name(), "ReadSlice", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var ms []*Model2
		_, err := dbr.conn.Select("*").From("models").Where("id > 0").Limit(100).Load(&ms)
		if err != nil {
			helper.SetError(b, dbr.Name(), "ReadSlice", err.Error())
		}
	}
}
