package bench

import (
	"database/sql"
	"github.com/efectn/go-orm-benchmarks/helper"
	"testing"

	bundb "github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Bun struct {
	helper.ORMInterface
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
	helper.BenchExec(b, bun.Name(), "Insert", func() error {
		m.Id = 0
		_, err := bun.conn.NewInsert().Model(m).Exec(ctx)
		return err
	})
}

func (bun *Bun) InsertMulti(b *testing.B) {
	ms := make([]*Model, 0, helper.BatchSize)
	for i := 0; i < helper.BatchSize; i++ {
		ms = append(ms, NewModel())
	}

	helper.BenchExec(b, bun.Name(), "InsertMulti", func() error {
		for _, m := range ms {
			m.Id = 0
		}

		_, err := bun.conn.NewInsert().Model(&ms).Exec(ctx)
		return err
	})
}

func (bun *Bun) Update(b *testing.B) {
	m := NewModel()

	_, err := bun.conn.NewInsert().Model(m).Exec(ctx)
	if err != nil {
		helper.SetError(b, bun.Name(), "Update", err.Error())
		return
	}
	m.Counter = helper.UpdatedCounter

	helper.BenchExec(b, bun.Name(), "Update", func() error {
		_, err := bun.conn.NewUpdate().Model(m).WherePK().Exec(ctx)
		return err
	})
}

func (bun *Bun) Read(b *testing.B) {
	m := NewModel()

	_, err := bun.conn.NewInsert().Model(m).Exec(ctx)
	if err != nil {
		helper.SetError(b, bun.Name(), "Read", err.Error())
		return
	}

	helper.BenchRows(b, bun.Name(), "Read", 1, func() (int, error) {
		if err := bun.conn.NewSelect().Model(m).WherePK().Scan(ctx); err != nil {
			return 0, err
		}
		return 1, nil
	})
}

func (bun *Bun) ReadSlice(b *testing.B) {
	m := NewModel()
	for i := 0; i < helper.BatchSize; i++ {
		m.Id = 0
		_, err := bun.conn.NewInsert().Model(m).Exec(ctx)
		if err != nil {
			helper.SetError(b, bun.Name(), "ReadSlice", err.Error())
			return
		}
	}

	helper.BenchRows(b, bun.Name(), "ReadSlice", helper.BatchSize, func() (int, error) {
		var models []*Model
		err := bun.conn.NewSelect().
			Model(&models).
			Where("id > ?", 0).
			Limit(helper.BatchSize).
			Scan(ctx)
		return len(models), err
	})
}
