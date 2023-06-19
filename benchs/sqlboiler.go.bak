package benchs

import (
	"database/sql"
	"fmt"

	models "github.com/efectn/go-orm-benchmarks/benchs/sqlboiler"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var sqlboiler *sql.DB

func init() {
	st := NewSuite("sqlboiler")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, SqlboilerInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, SqlboilerInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, SqlboilerUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, SqlboilerRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, SqlboilerReadSlice)

		var err error
		sqlboiler, err = sql.Open("pgx", OrmSource)
		CheckErr(err)

		boil.SetDB(sqlboiler)
	}
}

func SqlboilerInsert(b *B) {
	var m *models.Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel6()
	})

	for i := 0; i < b.N; i++ {
		m.ID = 0
		err := m.Insert(ctx, sqlboiler, boil.Infer())
		CheckErr(err, b)
	}
}

func SqlboilerInsertMulti(b *B) {
	panic(fmt.Errorf("doesn't support bulk-insert"))
}

func SqlboilerUpdate(b *B) {
	var m *models.Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel6()
		m.ID = 0
		err := m.Insert(ctx, sqlboiler, boil.Infer())
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		_, err := m.Update(ctx, sqlboiler, boil.Infer())
		CheckErr(err, b)
	}
}

func SqlboilerRead(b *B) {
	var m *models.Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel6()
		m.ID = 0
		err := m.Insert(ctx, sqlboiler, boil.Infer())
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		_, err := models.Models(qm.Where("id = 0")).Exec(sqlboiler)
		CheckErr(err, b)
	}
}

func SqlboilerReadSlice(b *B) {
	var m *models.Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel6()
		for i := 0; i < 100; i++ {
			m.ID = 0
			err := m.Insert(ctx, sqlboiler, boil.Infer())
			CheckErr(err, b)
		}
	})

	for i := 0; i < b.N; i++ {
		_, err := models.Models(qm.Where("id > 0"), qm.Limit(100)).All(ctx, sqlboiler)
		CheckErr(err, b)
	}
}
