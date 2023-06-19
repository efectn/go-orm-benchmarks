package benchs

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gormdbp *gorm.DB

func init() {
	st := NewSuite("gorm_prep")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, GormPrepInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, GormPrepInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, GormPrepUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, GormPrepRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, GormPrepReadSlice)

		var err error
		gormdbp, err = gorm.Open(postgres.New(postgres.Config{
			DSN: OrmSource,
		}), &gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
			Logger:                 logger.Default.LogMode(logger.Silent),
		})
		CheckErr(err)
	}
}

func GormPrepInsert(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		err := gormdbp.Create(m).Error
		CheckErr(err, b)
	}
}

func GormPrepInsertMulti(b *B) {
	var ms []*Model
	WrapExecute(b, func() {
		InitDB()
		ms = make([]*Model, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewModel())
		}
	})

	for i := 0; i < b.N; i++ {
		for _, m := range ms {
			m.Id = 0
		}
		err := gormdbp.Create(&ms).Error
		CheckErr(err, b)
	}
}

func GormPrepUpdate(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		err := gormdbp.Create(m).Error
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		err := gormdbp.Model(m).Updates(m).Error
		CheckErr(err, b)
	}
}

func GormPrepRead(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		err := gormdbp.Create(m).Error
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		err := gormdbp.Take(m).Error
		CheckErr(err, b)
	}
}

func GormPrepReadSlice(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.Id = 0
			err := gormdbp.Create(m).Error
			CheckErr(err, b)
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model
		err := gormdbp.Where("id > ?", 0).Limit(100).Find(&models).Error
		CheckErr(err, b)
	}
}
