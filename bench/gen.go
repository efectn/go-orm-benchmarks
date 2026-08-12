package bench

import (
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/efectn/go-orm-benchmarks/bench/gen/models"
	"github.com/efectn/go-orm-benchmarks/bench/gen/query"
	"github.com/efectn/go-orm-benchmarks/helper"
)

type Gen struct {
	helper.ORMInterface
	conn *gorm.DB
}

func CreateGen() helper.ORMInterface {
	return &Gen{}
}

func (gen *Gen) Name() string {
	return "gen"
}

func (gen *Gen) Init() error {
	var err error
	gen.conn, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  helper.OrmSource,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            false,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}

	return nil
}

func (gen *Gen) Close() error {
	db, _ := gen.conn.DB()

	return db.Close()
}

func (gen *Gen) Insert(b *testing.B) {
	m := models.NewModelAlt()
	helper.BenchExec(b, gen.Name(), "Insert", func() error {
		m.Id = 0
		return query.Use(gen.conn).Model.WithContext(ctx).Create(&m)
	})
}

func (gen *Gen) InsertMulti(b *testing.B) {
	ms := make([]*models.Model, 0, helper.BatchSize)
	for i := 0; i < helper.BatchSize; i++ {
		m := models.NewModelAlt()
		ms = append(ms, &m)
	}

	helper.BenchExec(b, gen.Name(), "InsertMulti", func() error {
		for _, m := range ms {
			m.Id = 0
		}
		return query.Use(gen.conn).Model.WithContext(ctx).Create(ms...)
	})
}

func (gen *Gen) Update(b *testing.B) {
	m := models.NewModelAlt()
	if err := query.Use(gen.conn).Model.WithContext(ctx).Create(&m); err != nil {
		helper.SetError(b, gen.Name(), "Update", err.Error())
		return
	}
	m.Counter = helper.UpdatedCounter

	helper.BenchExec(b, gen.Name(), "Update", func() error {
		_, err := query.Use(gen.conn).Model.WithContext(ctx).
			Where(query.Use(gen.conn).Model.Id.Eq(m.Id)).
			Updates(m)
		return err
	})
}

func (gen *Gen) Read(b *testing.B) {
	m := models.NewModelAlt()
	if err := query.Use(gen.conn).Model.WithContext(ctx).Create(&m); err != nil {
		helper.SetError(b, gen.Name(), "Read", err.Error())
		return
	}

	helper.BenchRows(b, gen.Name(), "Read", 1, func() (int, error) {
		_, err := query.Use(gen.conn).Model.WithContext(ctx).
			Where(query.Use(gen.conn).Model.Id.Eq(m.Id)).
			First()
		if err != nil {
			return 0, err
		}
		return 1, nil
	})
}

func (gen *Gen) ReadSlice(b *testing.B) {
	m := models.NewModelAlt()
	for i := 0; i < helper.BatchSize; i++ {
		m.Id = 0
		if err := query.Use(gen.conn).Model.WithContext(ctx).Create(&m); err != nil {
			helper.SetError(b, gen.Name(), "ReadSlice", err.Error())
			return
		}
	}

	helper.BenchRows(b, gen.Name(), "ReadSlice", helper.BatchSize, func() (int, error) {
		found, err := query.Use(gen.conn).Model.WithContext(ctx).
			Where(query.Use(gen.conn).Model.Id.Gt(0)).
			Limit(helper.BatchSize).
			Find()
		return len(found), err
	})
}
