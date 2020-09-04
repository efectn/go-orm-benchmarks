package benchs

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gormdb *gorm.DB

func init() {
	st := NewSuite("gorm")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*ORM_MULTI, GormInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*ORM_MULTI, GormInsertMulti)
		st.AddBenchmark("Update", 200*ORM_MULTI, GormUpdate)
		st.AddBenchmark("Read", 200*ORM_MULTI, GormRead)
		st.AddBenchmark("MultiRead limit 100", 200*ORM_MULTI, GormReadSlice)
		dsn := ORM_SOURCE
		v := postgres.Open(dsn)
		conn, err := gorm.Open(v, &gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
			Logger:                 logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{SlowThreshold: 0}),
		})
		if err != nil {
			fmt.Println(err)
		}
		gormdb = conn
	}
}

func GormInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		if err := gormdb.Create(m).Error; err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GormInsertMulti(b *B) {
	var ms []*Model
	wrapExecute(b, func() {
		initDB()
		ms = make([]*Model, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewModel())
		}
	})

	for i := 0; i < b.N; i++ {
		for _, m := range ms {
			m.Id = 0
		}
		if err := gormdb.Create(&ms).Error; err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GormUpdate(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		if err := gormdb.Create(m).Error; err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if err := gormdb.Model(m).Updates(m).Error; err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GormRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		if err := gormdb.Create(m).Error; err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})
	for i := 0; i < b.N; i++ {
		if err := gormdb.Take(m).Error; err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GormReadSlice(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.Id = 0
			if err := gormdb.Create(m).Error; err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model
		if err := gormdb.Where("id > ?", 0).Limit(100).Find(&models).Error; err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
