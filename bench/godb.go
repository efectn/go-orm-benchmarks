package bench

import (
	"github.com/efectn/go-orm-benchmarks/helper"
	_ "github.com/jackc/pgx/v4/stdlib"
	godbware "github.com/samonzeweb/godb"
	"github.com/samonzeweb/godb/adapters/postgresql"
	"sync"
	"testing"
)

type Godb struct {
	helper.ORMInterface
	mu   sync.Mutex
	conn *godbware.DB
}

func CreateGodb() helper.ORMInterface {
	return &Godb{}
}

func (godb *Godb) Name() string {
	return "godb"
}

func (godb *Godb) Init() error {
	var err error
	godb.conn, err = godbware.Open(postgresql.Adapter, helper.OrmSource)
	if err != nil {
		return err
	}

	return nil
}

func (godb *Godb) Close() error {
	return godb.conn.Close()
}

func (godb *Godb) Insert(b *testing.B) {
	m := NewModel2()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := godb.conn.Insert(m).Do()
		if err != nil {
			helper.SetError(b, godb.Name(), "Insert", err.Error())
		}
	}
}

func (godb *Godb) InsertMulti(b *testing.B) {
	ms := make([]*Model2, 0, 100)
	for i := 0; i < 100; i++ {
		ms = append(ms, NewModel2())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := godb.conn.BulkInsert(&ms).Do()
		if err != nil {
			helper.SetError(b, godb.Name(), "InsertMulti", err.Error())
		}
	}
}

func (godb *Godb) Update(b *testing.B) {
	m := NewModel2()

	err := godb.conn.Insert(m).Do()
	if err != nil {
		helper.SetError(b, godb.Name(), "Update", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := godb.conn.Update(m).Do()
		if err != nil {
			helper.SetError(b, godb.Name(), "Update", err.Error())
		}
	}
}

func (godb *Godb) Read(b *testing.B) {
	m := NewModel2()

	err := godb.conn.Insert(m).Do()
	if err != nil {
		helper.SetError(b, godb.Name(), "Read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := godb.conn.Select(m).Do()
		if err != nil {
			helper.SetError(b, godb.Name(), "Read", err.Error())
		}
	}
}

func (godb *Godb) ReadSlice(b *testing.B) {
	m := NewModel2()
	for i := 0; i < 100; i++ {
		err := godb.conn.Insert(m).Do()
		if err != nil {
			helper.SetError(b, godb.Name(), "ReadSlice", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var ms []*Model2
		err := godb.conn.Select(&ms).Where("id > 0").Limit(100).Do()
		if err != nil {
			helper.SetError(b, godb.Name(), "ReadSlice", err.Error())
		}
	}
}
