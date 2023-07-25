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

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Id = 0
		if err := query.Use(gen.conn).Model.WithContext(ctx).Create(&m); err != nil {
			helper.SetError(b, gen.Name(), "Insert", err.Error())
		}
	}
}

func (gen Gen) InsertMulti(b *testing.B) {
	ms := make([]*models.Model, 0, 100)
	for i := 0; i < 100; i++ {
		m := models.NewModelAlt()
		ms = append(ms, &m)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ms := make([]*models.Model, 0, 100)
		for i := 0; i < 100; i++ {
			m := models.NewModelAlt()
			ms = append(ms, &m)
		}
		if err := query.Use(gen.conn).Model.WithContext(ctx).Create(ms...); err != nil {
			helper.SetError(b, gen.Name(), "InsertMulti", err.Error())
		}
	}
}

func (gen *Gen) Update(b *testing.B) {
	m := models.NewModelAlt()
	if err := query.Use(gen.conn).Model.WithContext(ctx).Create(&m); err != nil {
		helper.SetError(b, gen.Name(), "Insert", err.Error())
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := query.Use(gen.conn).Model.WithContext(ctx).
			Where(query.Use(gen.conn).Model.Id.Eq(m.Id)).
			Updates(m)
		if err != nil {
			helper.SetError(b, gen.Name(), "Update", err.Error())
		}
	}
}

func (gen *Gen) Read(b *testing.B) {
	m := models.NewModelAlt()
	if err := query.Use(gen.conn).Model.WithContext(ctx).Create(&m); err != nil {
		helper.SetError(b, gen.Name(), "Insert", err.Error())
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := query.Use(gen.conn).Model.WithContext(ctx).
			Where(query.Use(gen.conn).Model.Id.Eq(m.Id)).
			First()
		if err != nil {
			helper.SetError(b, gen.Name(), "Read", err.Error())
		}
	}

}

func (gen *Gen) ReadSlice(b *testing.B) {
	m := models.NewModelAlt()
	for i := 0; i < 100; i++ {
		m.Id = 0
		if err := query.Use(gen.conn).Model.WithContext(ctx).Create(&m); err != nil {
			helper.SetError(b, gen.Name(), "ReadSlice", err.Error())
		}
	}

	for i := 0; i < b.N; i++ {
		_, err := query.Use(gen.conn).Model.WithContext(ctx).
			Where(query.Use(gen.conn).Model.Id.Gt(0)).
			Limit(100).
			Find()
		if err != nil {
			helper.SetError(b, gen.Name(), "ReadSlice", err.Error())
		}
	}
}
