package benchs

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gormPrepared *gorm.DB

func init() {
	st := NewSuite("gorm_prep")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, GormPreparedInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, GormPreparedInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, GormPreparedUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, GormPreparedRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, GormPreparedReadSlice)
		var err error
		gormPrepared, err = gorm.Open(postgres.New(postgres.Config{
			DSN: OrmSource,
		}), &gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
			Logger:                 logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			fmt.Println(err)
		}
	}
}

func GormPreparedInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		if err := gormPrepared.Create(m).Error; err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GormPreparedInsertMulti(b *B) {
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
		if err := gormPrepared.Create(&ms).Error; err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GormPreparedUpdate(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		if err := gormPrepared.Create(m).Error; err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if err := gormPrepared.Model(m).Updates(m).Error; err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GormPreparedRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		if err := gormPrepared.Create(m).Error; err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if err := gormPrepared.Take(m).Error; err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GormPreparedReadSlice(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.Id = 0
			if err := gormPrepared.Create(m).Error; err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model
		if err := gormPrepared.Where("id > ?", 0).Limit(100).Find(&models).Error; err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
