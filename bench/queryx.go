package bench

import (
	"github.com/efectn/go-orm-benchmarks/bench/queryx/db/queryx"
	"github.com/efectn/go-orm-benchmarks/helper"
	"testing"

	"github.com/efectn/go-orm-benchmarks/bench/queryx/db"
	_ "github.com/lib/pq"
)

const (
	queryxSelectMultiSQL = `SELECT * FROM models WHERE id > 0`
)

type Queryx struct {
	helper.ORMInterface
	client *db.QXClient
}

func CreateQueryx() helper.ORMInterface {
	return &Queryx{}
}

func (q *Queryx) Name() string {
	return "queryx"
}

func (q *Queryx) Init() error {
	client, err := db.NewClient()
	if err != nil {
		return err
	}
	q.client = client
	return err
}

func (q *Queryx) Close() error {
	return nil
}

func (q *Queryx) Insert(b *testing.B) {
	m := NewModel8()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.ID = 0
		_, err := q.client.QueryModel().Create(q.client.ChangeModel().SetName(m.Name).
			SetTitle(m.Title).SetFax(m.Fax).SetWeb(m.Web).SetAge(int64(m.Age)).SetCounter(int32(m.Counter)).SetRight(true))
		if err != nil {
			helper.SetError(b, q.Name(), "Insert", err.Error())
		}
	}
}

func (q *Queryx) InsertMulti(b *testing.B) {
	m := NewModel8()
	ms := make([]*queryx.ModelChange, 0, 100)
	for i := 0; i < 100; i++ {
		ms = append(ms, q.client.ChangeModel().SetName(m.Name).
			SetTitle(m.Title).SetFax(m.Fax).SetWeb(m.Web).SetAge(int64(m.Age)).SetCounter(int32(m.Counter)).SetRight(true))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := q.client.QueryModel().InsertAll(ms)
		if err != nil {
			helper.SetError(b, q.Name(), "InsertMulti", err.Error())
		}
	}
}

func (q *Queryx) Update(b *testing.B) {
	m := NewModel8()
	change := q.client.ChangeModel().SetName(m.Name).SetRight(true).
		SetTitle(m.Title).SetFax(m.Fax).SetWeb(m.Web).SetAge(int64(m.Age)).SetCounter(int32(m.Counter))
	m8, err := q.client.QueryModel().Create(change)
	if err != nil {
		helper.SetError(b, q.Name(), "Update", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := q.client.QueryModel().Where(q.client.ModelID.EQ(m8.ID)).UpdateAll(change)
		if err != nil {
			helper.SetError(b, q.Name(), "Update", err.Error())
		}
	}
}

func (q *Queryx) Read(b *testing.B) {
	m := NewModel8()
	change := q.client.ChangeModel().SetName(m.Name).SetRight(true).
		SetTitle(m.Title).SetFax(m.Fax).SetWeb(m.Web).SetAge(int64(m.Age)).SetCounter(int32(m.Counter))
	_, err := q.client.QueryModel().Create(change)
	if err != nil {
		helper.SetError(b, q.Name(), "Read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := q.client.QueryModel().FindBy(q.client.ModelName.EQ(m.Name))
		if err != nil {
			helper.SetError(b, q.Name(), "Read", err.Error())
		}
	}
}

func (q *Queryx) ReadSlice(b *testing.B) {
	m := NewModel8()
	change := q.client.ChangeModel().SetName(m.Name).SetRight(true).
		SetTitle(m.Title).SetFax(m.Fax).SetWeb(m.Web).SetAge(int64(m.Age)).SetCounter(int32(m.Counter))
	for i := 0; i < 100; i++ {
		_, err := q.client.QueryModel().Create(change)
		if err != nil {
			helper.SetError(b, q.Name(), "ReadSlice", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := q.client.QueryModel().FindBySQL(queryxSelectMultiSQL)
		if err != nil {
			helper.SetError(b, q.Name(), "ReadSlice", err.Error())
		}
	}
}
