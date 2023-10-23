package bench

import (
	"github.com/efectn/go-orm-benchmarks/bench/queryx/db/queryx"
	"github.com/efectn/go-orm-benchmarks/helper"
	"testing"

	"github.com/efectn/go-orm-benchmarks/bench/queryx/db"
	_ "github.com/lib/pq"
)

var (
	c *db.QXClient
)

const (
	queryxSelectMultiSQL = `SELECT * FROM models WHERE id > 0`
)

type Queryx struct {
	helper.ORMInterface
}

func CreateQueryx() helper.ORMInterface {
	return &Queryx{}
}

func (Queryx *Queryx) Name() string {
	return "queryx"
}

func (Queryx *Queryx) Init() error {
	client, err := db.NewClient()
	if err != nil {
		return err
	}
	c = client
	return err
}

func (Queryx *Queryx) Close() error {
	return nil
}

func (Queryx *Queryx) Insert(b *testing.B) {
	m := NewModel8()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.ID = 0
		_, err := c.QueryModel().Create(c.ChangeModel().SetName(m.Name).
			SetTitle(m.Title).SetFax(m.Fax).SetWeb(m.Web).SetAge(int64(m.Age)).SetCounter(int32(m.Counter)).SetRight(true))
		if err != nil {
			helper.SetError(b, Queryx.Name(), "Insert", err.Error())
		}
	}
}

func (Queryx *Queryx) InsertMulti(b *testing.B) {
	m := NewModel8()
	ms := make([]*queryx.ModelChange, 0, 100)
	for i := 0; i < 100; i++ {
		ms = append(ms, c.ChangeModel().SetName(m.Name).
			SetTitle(m.Title).SetFax(m.Fax).SetWeb(m.Web).SetAge(int64(m.Age)).SetCounter(int32(m.Counter)).SetRight(true))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := c.QueryModel().InsertAll(ms)
		if err != nil {
			helper.SetError(b, Queryx.Name(), "InsertMulti", err.Error())
		}
	}
}

func (Queryx *Queryx) Update(b *testing.B) {
	m := NewModel8()
	change := c.ChangeModel().SetName(m.Name).SetRight(true).
		SetTitle(m.Title).SetFax(m.Fax).SetWeb(m.Web).SetAge(int64(m.Age)).SetCounter(int32(m.Counter))
	m8, err := c.QueryModel().Create(change)
	if err != nil {
		helper.SetError(b, Queryx.Name(), "Update", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := c.QueryModel().Where(c.ModelID.EQ(m8.ID)).UpdateAll(change)
		if err != nil {
			helper.SetError(b, Queryx.Name(), "Update", err.Error())
		}
	}
}

func (Queryx *Queryx) Read(b *testing.B) {
	m := NewModel8()
	change := c.ChangeModel().SetName(m.Name).SetRight(true).
		SetTitle(m.Title).SetFax(m.Fax).SetWeb(m.Web).SetAge(int64(m.Age)).SetCounter(int32(m.Counter))
	_, err := c.QueryModel().Create(change)
	if err != nil {
		helper.SetError(b, Queryx.Name(), "Read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := c.QueryModel().FindBy(c.ModelName.EQ(m.Name))
		if err != nil {
			helper.SetError(b, Queryx.Name(), "Read", err.Error())
		}
	}
}

func (Queryx *Queryx) ReadSlice(b *testing.B) {
	m := NewModel8()
	change := c.ChangeModel().SetName(m.Name).SetRight(true).
		SetTitle(m.Title).SetFax(m.Fax).SetWeb(m.Web).SetAge(int64(m.Age)).SetCounter(int32(m.Counter))
	for i := 0; i < 100; i++ {
		_, err := c.QueryModel().Create(change)
		if err != nil {
			helper.SetError(b, Queryx.Name(), "ReadSlice", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := c.QueryModel().FindBySQL(queryxSelectMultiSQL)
		if err != nil {
			helper.SetError(b, Queryx.Name(), "ReadSlice", err.Error())
		}
	}
}
