package benchs

import (
	"fmt"
	"gorm.io/driver/postgres"
	gormdb "gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
	"testing"
)

type Gorm struct {
	Instance
	mu         sync.Mutex
	conn       *gormdb.DB
	iterations int // Same as b.N, just to customize it
	errors     map[string]string
}

func CreateGorm(iterations int) Instance {
	gorm := &Gorm{
		iterations: iterations,
		errors:     map[string]string{},
	}

	return gorm
}

func (gorm *Gorm) Name() string {
	return "gorm"
}

func (gorm *Gorm) Init() error {
	var err error
	gorm.conn, err = gormdb.Open(postgres.New(postgres.Config{
		DSN:                  OrmSource,
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

func (gorm *Gorm) GetError(method string) string {
	return gorm.errors[method]
}

func (gorm *Gorm) Insert(b *testing.B) {
	m := NewModel()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Id = 0
		err := gorm.conn.Create(m).Error
		if err != nil {
			gorm.error(b, "insert", fmt.Sprintf("gorm: insert: %v", err))
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
			gorm.error(b, "insert_multi", fmt.Sprintf("gorm: insert_multi: %v", err))
		}
	}
}

func (gorm *Gorm) Update(b *testing.B) {
	m := NewModel()

	err := gorm.conn.Create(m).Error
	if err != nil {
		gorm.error(b, "update", fmt.Sprintf("gorm: update: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := gorm.conn.Model(m).Updates(m).Error
		if err != nil {
			gorm.error(b, "update", fmt.Sprintf("gorm: update: %v", err))
		}
	}
}

func (gorm *Gorm) Read(b *testing.B) {
	m := NewModel()

	err := gorm.conn.Create(m).Error
	if err != nil {
		gorm.error(b, "read", fmt.Sprintf("gorm: read: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := gorm.conn.Take(m).Error
		if err != nil {
			gorm.error(b, "read", fmt.Sprintf("gorm: read: %v", err))
		}
	}
}

func (gorm *Gorm) ReadSlice(b *testing.B) {
	m := NewModel()
	for i := 0; i < 100; i++ {
		m.Id = 0
		err := gorm.conn.Create(m).Error
		if err != nil {
			gorm.error(b, "read_slice", fmt.Sprintf("gorm: read_slice: %v", err))
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var models []*Model
		err := gorm.conn.Where("id > ?", 0).Limit(100).Find(&models).Error
		if err != nil {
			gorm.error(b, "read_slice", fmt.Sprintf("gorm: read_slice: %v", err))
		}
	}
}

func (gorm *Gorm) error(b *testing.B, method string, err string) {
	b.Helper()

	gorm.mu.Lock()
	gorm.errors[method] = err
	gorm.mu.Unlock()
	b.Fail()
}
