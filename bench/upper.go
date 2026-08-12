package bench

import (
	"testing"

	"github.com/efectn/go-orm-benchmarks/helper"
	db "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

type Upper struct {
	helper.ORMInterface
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
	helper.BenchExec(b, upper.Name(), "Insert", func() error {
		_, err := upper.conn.Collection("models").Insert(m)
		return err
	})
}

func (upper *Upper) InsertMulti(b *testing.B) {
	m := NewModel4()

	helper.BenchExec(b, upper.Name(), "InsertMulti", func() error {
		batch := upper.conn.SQL().InsertInto("models").Columns("name", "title", "fax", "web", "age", "right", "counter").Batch(helper.BatchSize)

		go func() {
			for i := 0; i < helper.BatchSize; i++ {
				batch.Values(m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
			}
			batch.Done()
		}()

		return batch.Wait()
	})
}

func (upper *Upper) Update(b *testing.B) {
	m := NewModel4()

	err := upper.conn.Collection("models").InsertReturning(m)
	if err != nil {
		helper.SetError(b, upper.Name(), "Update", err.Error())
		return
	}
	m.Counter = helper.UpdatedCounter

	helper.BenchExec(b, upper.Name(), "Update", func() error {
		return upper.conn.Collection("models").Find(m.ID).Update(m)
	})
}

func (upper *Upper) Read(b *testing.B) {
	m := NewModel4()

	err := upper.conn.Collection("models").InsertReturning(m)
	if err != nil {
		helper.SetError(b, upper.Name(), "Read", err.Error())
		return
	}

	helper.BenchRows(b, upper.Name(), "Read", 1, func() (int, error) {
		var result Model4
		err := upper.conn.SQL().SelectFrom("models").Where("id = ?", m.ID).One(&result)
		if err != nil {
			return 0, err
		}
		return 1, nil
	})
}

func (upper *Upper) ReadSlice(b *testing.B) {
	m := NewModel4()
	for i := 0; i < helper.BatchSize; i++ {
		m.ID = 0
		err := upper.conn.Collection("models").InsertReturning(m)
		if err != nil {
			helper.SetError(b, upper.Name(), "ReadSlice", err.Error())
			return
		}
	}

	helper.BenchRows(b, upper.Name(), "ReadSlice", helper.BatchSize, func() (int, error) {
		var ms []*Model4
		err := upper.conn.SQL().SelectFrom("models").Where("id > ?", 0).Limit(helper.BatchSize).All(&ms)
		if err != nil {
			return 0, err
		}
		return len(ms), nil
	})
}
