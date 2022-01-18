package benchs

/*
import (
	"database/sql"
	"fmt"

	"github.com/efectn/go-orm-benchmarks/benchs/sqlc_prepared/db"
)

var (
	sqlcPreparedQueries *db.Queries
)

func init() {
	st := NewSuite("sqlc_prep")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, SqlcPreparedInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, SqlcPreparedInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, SqlcPreparedUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, SqlcPreparedRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, SqlcPreparedReadSlice)

		dbConnection, _ := sql.Open("pgx", OrmSource)

		var err error

		sqlcPreparedQueries, err = db.Prepare(ctx, dbConnection)
		if err != nil {
			panic(err)
		}
	}
}

func SqlcPreparedInsert(b *B) {
	var m *Model

	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0

		created, err := sqlcPreparedQueries.CreateModel(ctx, db.CreateModelParams{
			Name:    m.Name,
			Title:   m.Title,
			Fax:     m.Fax,
			Web:     m.Web,
			Age:     int32(m.Age),
			Right:   m.Right,
			Counter: m.Counter,
		})
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}

		m = &Model{
			Id:      int(created.ID),
			Name:    created.Name,
			Title:   created.Title,
			Fax:     created.Fax,
			Web:     created.Web,
			Age:     int(created.Age),
			Right:   created.Right,
			Counter: created.Counter,
		}
	}
}

func SqlcPreparedInsertMulti(b *B) {
	panic(fmt.Errorf("TBD"))
}

func SqlcPreparedUpdate(b *B) {
	var m *Model

	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		created, err := sqlcPreparedQueries.CreateModel(ctx, db.CreateModelParams{
			Name:    m.Name,
			Title:   m.Title,
			Fax:     m.Fax,
			Web:     m.Web,
			Age:     int32(m.Age),
			Right:   m.Right,
			Counter: m.Counter,
		})
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}

		m = &Model{
			Id:      int(created.ID),
			Name:    created.Name,
			Title:   created.Title,
			Fax:     created.Fax,
			Web:     created.Web,
			Age:     int(created.Age),
			Right:   created.Right,
			Counter: created.Counter,
		}
	})

	for i := 0; i < b.N; i++ {
		if err := sqlcPreparedQueries.UpdateModel(ctx, db.UpdateModelParams{
			Name:    m.Name,
			Title:   m.Title,
			Fax:     m.Fax,
			Web:     m.Web,
			Age:     int32(m.Age),
			Right:   m.Right,
			Counter: m.Counter,
			ID:      int32(m.Id),
		}); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func SqlcPreparedRead(b *B) {
	var m *Model

	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		created, err := sqlcPreparedQueries.CreateModel(ctx, db.CreateModelParams{
			Name:    m.Name,
			Title:   m.Title,
			Fax:     m.Fax,
			Web:     m.Web,
			Age:     int32(m.Age),
			Right:   m.Right,
			Counter: m.Counter,
		})
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}

		m = &Model{
			Id:      int(created.ID),
			Name:    created.Name,
			Title:   created.Title,
			Fax:     created.Fax,
			Web:     created.Web,
			Age:     int(created.Age),
			Right:   created.Right,
			Counter: created.Counter,
		}
	})

	for i := 0; i < b.N; i++ {
		readed, err := sqlcPreparedQueries.GetModel(ctx, int32(m.Id))
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}

		m = &Model{
			Id:      int(readed.ID),
			Name:    readed.Name,
			Title:   readed.Title,
			Fax:     readed.Fax,
			Web:     readed.Web,
			Age:     int(readed.Age),
			Right:   readed.Right,
			Counter: readed.Counter,
		}
	}
}

func SqlcPreparedReadSlice(b *B) {
	var m *Model

	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.Id = 0

			created, err := sqlcPreparedQueries.CreateModel(ctx, db.CreateModelParams{
				Name:    m.Name,
				Title:   m.Title,
				Fax:     m.Fax,
				Web:     m.Web,
				Age:     int32(m.Age),
				Right:   m.Right,
				Counter: m.Counter,
			})
			if err != nil {
				fmt.Println(err)
				b.FailNow()
			}

			m = &Model{
				Id:      int(created.ID),
				Name:    created.Name,
				Title:   created.Title,
				Fax:     created.Fax,
				Web:     created.Web,
				Age:     int(created.Age),
				Right:   created.Right,
				Counter: created.Counter,
			}
		}
	})

	for i := 0; i < b.N; i++ {
		sqlcModels, err := sqlcPreparedQueries.ListModels(ctx, db.ListModelsParams{
			ID:    0,
			Limit: 100,
		})
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}

		models := make([]*Model, len(sqlcModels))

		for i2 := range sqlcModels {
			models[i2] = &Model{
				Name:    sqlcModels[i2].Name,
				Title:   sqlcModels[i2].Title,
				Fax:     sqlcModels[i2].Fax,
				Web:     sqlcModels[i2].Web,
				Age:     int(sqlcModels[i2].Age),
				Right:   sqlcModels[i2].Right,
				Counter: sqlcModels[i2].Counter,
			}
		}
	}
}
*/
