package benchs

import (
	"database/sql"
	"fmt"

	"github.com/frederikhors/orm-benchmark/benchs/sqlc/db"
)

var (
	sqlcQueries *db.Queries
)

func init() {
	st := NewSuite("sqlc")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, SqlcInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, SqlcInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, SqlcUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, SqlcRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, SqlcReadSlice)

		dbConnection, _ := sql.Open("pgx", OrmSource)
		sqlcQueries = db.New(dbConnection)
	}
}

func SqlcInsert(b *B) {
	var m *Model

	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0

		created, err := sqlcQueries.CreateModel(ctx, db.CreateModelParams{
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

func SqlcInsertMulti(b *B) {
	panic(fmt.Errorf("TBD"))
}

func SqlcUpdate(b *B) {
	var m *Model

	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		created, err := sqlcQueries.CreateModel(ctx, db.CreateModelParams{
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
		if err := sqlcQueries.UpdateModel(ctx, db.UpdateModelParams{
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

func SqlcRead(b *B) {
	var m *Model

	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		created, err := sqlcQueries.CreateModel(ctx, db.CreateModelParams{
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
		readed, err := sqlcQueries.GetModel(ctx, int32(m.Id))
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

func SqlcReadSlice(b *B) {
	var m *Model

	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.Id = 0

			created, err := sqlcQueries.CreateModel(ctx, db.CreateModelParams{
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
		sqlcModels, err := sqlcQueries.ListModels(ctx, db.ListModelsParams{
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
