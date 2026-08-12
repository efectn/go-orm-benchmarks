package bench

import (
	"github.com/efectn/go-orm-benchmarks/helper"
	"gorm.io/driver/postgres"
	gormdb "gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

type Gorm struct {
	helper.ORMInterface
	conn *gormdb.DB
}

func CreateGorm() helper.ORMInterface {
	return &Gorm{}
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

	helper.BenchExec(b, gorm.Name(), "Insert", func() error {
		m.Id = 0
		return gorm.conn.Create(m).Error
	})
}

func (gorm *Gorm) InsertMulti(b *testing.B) {
	ms := make([]*Model, 0, helper.BatchSize)
	for i := 0; i < helper.BatchSize; i++ {
		ms = append(ms, NewModel())
	}

	helper.BenchExec(b, gorm.Name(), "InsertMulti", func() error {
		for _, m := range ms {
			m.Id = 0
		}
		return gorm.conn.Create(&ms).Error
	})
}

func (gorm *Gorm) Update(b *testing.B) {
	m := NewModel()

	err := gorm.conn.Create(m).Error
	if err != nil {
		helper.SetError(b, gorm.Name(), "Update", err.Error())
		return
	}
	m.Counter = helper.UpdatedCounter

	helper.BenchExec(b, gorm.Name(), "Update", func() error {
		return gorm.conn.Model(m).Updates(m).Error
	})
}

func (gorm *Gorm) Read(b *testing.B) {
	m := NewModel()

	err := gorm.conn.Create(m).Error
	if err != nil {
		helper.SetError(b, gorm.Name(), "Read", err.Error())
		return
	}

	helper.BenchRows(b, gorm.Name(), "Read", 1, func() (int, error) {
		if err := gorm.conn.Take(m).Error; err != nil {
			return 0, err
		}
		return 1, nil
	})
}

func (gorm *Gorm) ReadSlice(b *testing.B) {
	m := NewModel()
	for i := 0; i < helper.BatchSize; i++ {
		m.Id = 0
		err := gorm.conn.Create(m).Error
		if err != nil {
			helper.SetError(b, gorm.Name(), "ReadSlice", err.Error())
			return
		}
	}

	helper.BenchRows(b, gorm.Name(), "ReadSlice", helper.BatchSize, func() (int, error) {
		var models []*Model
		err := gorm.conn.Where("id > ?", 0).Limit(helper.BatchSize).Find(&models).Error
		return len(models), err
	})
}
