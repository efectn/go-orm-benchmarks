package bench

import (
	"database/sql"
	"github.com/efectn/go-orm-benchmarks/helper"
	"sync"
	"testing"

	bundb "github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Bun struct {
	helper.ORMInterface
	mu   sync.Mutex
	conn *bundb.DB
}

func CreateBun() helper.ORMInterface {
	return &Bun{}
}

func (bun *Bun) Name() string {
	return "bun"
}

func (bun *Bun) Init() error {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(helper.ConvertSourceToDSN())))
	sqldb.SetMaxOpenConns(helper.OrmMaxConn)
	sqldb.SetMaxIdleConns(helper.OrmMaxIdle)

	bun.conn = bundb.NewDB(sqldb, pgdialect.New())
	return nil
}

func (bun *Bun) Close() error {
	return bun.conn.Close()
}

func (bun *Bun) Insert(b *testing.B) {
	m := NewModel()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Id = 0

		_, err := bun.conn.NewInsert().Model(m).Exec(ctx)
		if err != nil {
			helper.SetError(b, bun.Name(), "Insert", err.Error())
		}
	}
}

func (bun *Bun) InsertMulti(b *testing.B) {
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

		_, err := bun.conn.NewInsert().Model(&ms).Exec(ctx)
		if err != nil {
			helper.SetError(b, bun.Name(), "InsertMulti", err.Error())
		}
	}
}

func (bun *Bun) Update(b *testing.B) {
	m := NewModel()

	_, err := bun.conn.NewInsert().Model(m).Exec(ctx)
	if err != nil {
		helper.SetError(b, bun.Name(), "Update", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := bun.conn.NewUpdate().Model(m).WherePK().Exec(ctx)
		if err != nil {
			helper.SetError(b, bun.Name(), "Update", err.Error())
		}
	}
}

func (bun *Bun) Read(b *testing.B) {
	m := NewModel()

	_, err := bun.conn.NewInsert().Model(m).Exec(ctx)
	if err != nil {
		helper.SetError(b, bun.Name(), "Read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := bun.conn.NewSelect().Model(m).Scan(ctx)
		if err != nil {
			helper.SetError(b, bun.Name(), "Read", err.Error())
		}
	}
}

func (bun *Bun) ReadSlice(b *testing.B) {
	m := NewModel()
	for i := 0; i < 100; i++ {
		m.Id = 0
		_, err := bun.conn.NewInsert().Model(m).Exec(ctx)
		if err != nil {
			helper.SetError(b, bun.Name(), "ReadSlice", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var models []*Model
		err := bun.conn.NewSelect().
			Model(&models).
			Where("id > ?", 0).
			Limit(100).
			Scan(ctx)
		if err != nil {
			helper.SetError(b, bun.Name(), "ReadSlice", err.Error())
		}
	}
}
