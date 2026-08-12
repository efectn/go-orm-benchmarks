package bench

import (
	"testing"

	"github.com/efectn/go-orm-benchmarks/helper"
	_ "github.com/jackc/pgx/v5/stdlib"
	godbware "github.com/samonzeweb/godb"
	"github.com/samonzeweb/godb/adapters/postgresql"
)

type Godb struct {
	helper.ORMInterface
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
	helper.BenchExec(b, godb.Name(), "Insert", func() error {
		m.ID = 0
		return godb.conn.Insert(m).Do()
	})
}

func (godb *Godb) InsertMulti(b *testing.B) {
	ms := make([]*Model2, 0, helper.BatchSize)
	for i := 0; i < helper.BatchSize; i++ {
		ms = append(ms, NewModel2())
	}

	helper.BenchExec(b, godb.Name(), "InsertMulti", func() error {
		for _, m := range ms {
			m.ID = 0
		}
		return godb.conn.BulkInsert(&ms).Do()
	})
}

func (godb *Godb) Update(b *testing.B) {
	m := NewModel2()

	err := godb.conn.Insert(m).Do()
	if err != nil {
		helper.SetError(b, godb.Name(), "Update", err.Error())
		return
	}
	m.Counter = helper.UpdatedCounter

	helper.BenchExec(b, godb.Name(), "Update", func() error {
		return godb.conn.Update(m).Do()
	})
}

func (godb *Godb) Read(b *testing.B) {
	m := NewModel2()

	err := godb.conn.Insert(m).Do()
	if err != nil {
		helper.SetError(b, godb.Name(), "Read", err.Error())
		return
	}

	helper.BenchRows(b, godb.Name(), "Read", 1, func() (int, error) {
		if err := godb.conn.Select(m).Where("id = ?", m.ID).Do(); err != nil {
			return 0, err
		}
		return 1, nil
	})
}

func (godb *Godb) ReadSlice(b *testing.B) {
	m := NewModel2()
	for i := 0; i < helper.BatchSize; i++ {
		m.ID = 0
		err := godb.conn.Insert(m).Do()
		if err != nil {
			helper.SetError(b, godb.Name(), "ReadSlice", err.Error())
			return
		}
	}

	helper.BenchRows(b, godb.Name(), "ReadSlice", helper.BatchSize, func() (int, error) {
		var ms []*Model2
		err := godb.conn.Select(&ms).Where("id > 0").Limit(helper.BatchSize).Do()
		return len(ms), err
	})
}
