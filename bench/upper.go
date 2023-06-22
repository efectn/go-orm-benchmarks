package bench

import (
	"github.com/efectn/go-orm-benchmarks/helper"
	db "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"sync"
	"testing"
)

type Upper struct {
	helper.ORMInterface
	mu   sync.Mutex
	conn db.Session
}

func CreateUpper() helper.ORMInterface {
	return &Upper{}
}

func (upper *Upper) Name() string {
	return "upper"
}

func (upper *Upper) Init() error {
	var err error
	source := helper.SplitSource()

	upper.conn, err = postgresql.Open(postgresql.ConnectionURL{
		Host:     source["host"] + ":5432",
		User:     source["user"],
		Password: source["password"],
		Database: source["dbname"],
	})
	if err != nil {
		return err
	}

	// Disable logger
	db.LC().SetLogger(nil)

	if err := upper.conn.Ping(); err != nil {
		return err
	}

	return nil
}

func (upper *Upper) Close() error {
	return upper.conn.Close()
}

func (upper *Upper) Insert(b *testing.B) {
	m := NewModel4()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := upper.conn.Collection("models").Insert(m)
		if err != nil {
			helper.SetError(b, upper.Name(), "Insert", err.Error())
		}
	}
}

func (upper *Upper) InsertMulti(b *testing.B) {
	m := NewModel4()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		batch := upper.conn.SQL().InsertInto("models").Columns("name", "title", "fax", "web", "age", "right", "counter").Batch(100)

		go func() {
			for i := 0; i < 100; i++ {
				batch.Values(m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
			}
			batch.Done()
		}()

		if err := batch.Wait(); err != nil {
			helper.SetError(b, upper.Name(), "InsertMulti", err.Error())
		}
	}
}

func (upper *Upper) Update(b *testing.B) {
	m := NewModel4()

	err := upper.conn.Collection("models").InsertReturning(m)
	if err != nil {
		helper.SetError(b, upper.Name(), "Update", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := upper.conn.Collection("models").UpdateReturning(m)
		if err != nil {
			helper.SetError(b, upper.Name(), "Update", err.Error())
		}
	}
}

func (upper *Upper) Read(b *testing.B) {
	m := NewModel4()

	err := upper.conn.Collection("models").InsertReturning(m)
	if err != nil {
		helper.SetError(b, upper.Name(), "Read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := upper.conn.SQL().SelectFrom("models").Where("id = ?", m.ID).One(m)
		if err != nil {
			helper.SetError(b, upper.Name(), "Read", err.Error())
		}
	}
}

func (upper *Upper) ReadSlice(b *testing.B) {
	m := NewModel4()
	for i := 0; i < 100; i++ {
		m.ID = 0
		err := upper.conn.Collection("models").InsertReturning(m)
		if err != nil {
			helper.SetError(b, upper.Name(), "ReadSlice", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var ms []*Model4
		err := upper.conn.SQL().SelectFrom("models").Where("id > ?", m.ID).Limit(100).All(&ms)
		if err != nil {
			helper.SetError(b, upper.Name(), "ReadSlice", err.Error())
		}
	}
}
