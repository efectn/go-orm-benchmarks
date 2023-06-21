package bench

import (
	"github.com/efectn/go-orm-benchmarks/helper"
	"gorm.io/driver/postgres"
	gormdb "gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
	"testing"
)

type GormPrep struct {
	helper.ORMInterface
	mu         sync.Mutex
	conn       *gormdb.DB
	iterations int // Same as b.N, just to customize it
}

func CreateGormPrep(iterations int) helper.ORMInterface {
	gormPrep := &GormPrep{
		iterations: iterations,
	}

	return gormPrep
}

func (gorm *GormPrep) Name() string {
	return "gorm_prep"
}

func (gorm *GormPrep) Init() error {
	var err error
	gorm.conn, err = gormdb.Open(postgres.New(postgres.Config{
		DSN: helper.OrmSource,
	}), &gormdb.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}

	return nil
}

func (gorm *GormPrep) Close() error {
	db, _ := gorm.conn.DB()

	return db.Close()
}

func (gorm *GormPrep) Insert(b *testing.B) {
	m := NewModel()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Id = 0
		err := gorm.conn.Create(m).Error
		if err != nil {
			helper.SetError(b, gorm.Name(), "Insert", err.Error())
		}
	}
}

func (gorm *GormPrep) InsertMulti(b *testing.B) {
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
		err := gorm.conn.Create(&ms).Error
		if err != nil {
			helper.SetError(b, gorm.Name(), "InsertMulti", err.Error())
		}
	}
}

func (gorm *GormPrep) Update(b *testing.B) {
	m := NewModel()

	err := gorm.conn.Create(m).Error
	if err != nil {
		helper.SetError(b, gorm.Name(), "Update", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := gorm.conn.Model(m).Updates(m).Error
		if err != nil {
			helper.SetError(b, gorm.Name(), "Update", err.Error())
		}
	}
}

func (gorm *GormPrep) Read(b *testing.B) {
	m := NewModel()

	err := gorm.conn.Create(m).Error
	if err != nil {
		helper.SetError(b, gorm.Name(), "Read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := gorm.conn.Take(m).Error
		if err != nil {
			helper.SetError(b, gorm.Name(), "Read", err.Error())
		}
	}
}

func (gorm *GormPrep) ReadSlice(b *testing.B) {
	m := NewModel()
	for i := 0; i < 100; i++ {
		m.Id = 0
		err := gorm.conn.Create(m).Error
		if err != nil {
			helper.SetError(b, gorm.Name(), "ReadSlice", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var models []*Model
		err := gorm.conn.Where("id > ?", 0).Limit(100).Find(&models).Error
		if err != nil {
			helper.SetError(b, gorm.Name(), "ReadSlice", err.Error())
		}
	}
}
