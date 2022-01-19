package benchs

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gormdb *gorm.DB

func init() {
	st := NewSuite("gorm")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, GormInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, GormInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, GormUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, GormRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, GormReadSlice)

		var err error
		gormdb, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  OrmSource,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            false,
			Logger:                 logger.Default.LogMode(logger.Silent),
		})
		CheckErr(err)
	}
}

func GormInsert(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		err := gormdb.Create(m).Error
		CheckErr(err, b)
	}
}

func GormInsertMulti(b *B) {
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
		err := gormdb.Create(&ms).Error
		CheckErr(err, b)
	}
}

func GormUpdate(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		err := gormdb.Create(m).Error
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		err := gormdb.Model(m).Updates(m).Error
		CheckErr(err, b)
	}
}

func GormRead(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		err := gormdb.Create(m).Error
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		err := gormdb.Take(m).Error
		CheckErr(err, b)
	}
}

func GormReadSlice(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.Id = 0
			err := gormdb.Create(m).Error
			CheckErr(err, b)
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model
		err := gormdb.Where("id > ?", 0).Limit(100).Find(&models).Error
		CheckErr(err, b)
	}
}
