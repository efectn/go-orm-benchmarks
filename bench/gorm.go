package bench

import (
	"github.com/efectn/go-orm-benchmarks/helper"
	"gorm.io/driver/postgres"
	gormdb "gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
	"testing"
)

type Gorm struct {
	helper.ORMInterface
	mu         sync.Mutex
	conn       *gormdb.DB
	iterations int // Same as b.N, just to customize it
}

func CreateGorm(iterations int) helper.ORMInterface {
	gorm := &Gorm{
		iterations: iterations,
	}

	return gorm
}

func (gorm *Gorm) Name() string {
	return "gorm"
}

func (gorm *Gorm) Init() error {
	var err error
	gorm.conn, err = gormdb.Open(postgres.New(postgres.Config{
		DSN:                  helper.OrmSource,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gormdb.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            false,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}

	return nil
}

func (gorm *Gorm) Close() error {
	db, _ := gorm.conn.DB()

	return db.Close()
}

func (gorm *Gorm) Insert(b *testing.B) {
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

func (gorm *Gorm) InsertMulti(b *testing.B) {
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

func (gorm *Gorm) Update(b *testing.B) {
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

func (gorm *Gorm) Read(b *testing.B) {
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

func (gorm *Gorm) ReadSlice(b *testing.B) {
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
