package benchs

import (
	"fmt"

	popware "github.com/gobuffalo/pop/v6"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var pop *popware.Connection

func init() {
	st := NewSuite("pop")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, PopInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, PopInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, PopUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, PopRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, PopReadSlice)

		var err error
		pop, err = popware.NewConnection(&popware.ConnectionDetails{
			URL: ConvertSourceToDSN(),
		})
		CheckErr(err)

		err = pop.Open()
		CheckErr(err)
	}
}

func PopInsert(b *B) {
	var m *Model3
	WrapExecute(b, func() {
		InitDB()
		m = NewModel3()
	})

	for i := 0; i < b.N; i++ {
		err := pop.Create(m)
		CheckErr(err, b)
	}
}

func PopInsertMulti(b *B) {
	panic(fmt.Errorf("doesn't support bulk-insert"))
}

func PopUpdate(b *B) {
	var m *Model3
	WrapExecute(b, func() {
		InitDB()
		m = NewModel3()
		err := pop.Create(m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		err := pop.Update(m)
		CheckErr(err, b)
	}
}

func PopRead(b *B) {
	var m *Model3
	WrapExecute(b, func() {
		InitDB()
		m = NewModel3()
		err := pop.Create(m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		err := pop.First(m)
		CheckErr(err, b)
	}
}

func PopReadSlice(b *B) {
	var m *Model3
	WrapExecute(b, func() {
		InitDB()
		m = NewModel3()
		for i := 0; i < 100; i++ {
			err := pop.Create(m)
			CheckErr(err, b)
		}
	})

	for i := 0; i < b.N; i++ {
		var ms []Model3
		err := pop.Where("id > 0").Limit(100).All(&ms)
		CheckErr(err, b)
	}
}
