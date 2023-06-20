package benchs

import (
	"database/sql"
	"fmt"
	"sync"
	"testing"

	"github.com/efectn/go-orm-benchmarks/benchs/sqlc/db"
)

type Sqlc struct {
	Instance
	mu         sync.Mutex
	conn       *db.Queries
	db         *sql.DB
	iterations int // Same as b.N, just to customize it
	errors     map[string]string
}

func CreateSqlc(iterations int) Instance {
	sqlc := &Sqlc{
		iterations: iterations,
		errors:     map[string]string{},
	}

	return sqlc
}

func (sqlc *Sqlc) Name() string {
	return "sqlc"
}

func (sqlc *Sqlc) Init() error {
	var err error
	sqlc.db, err = sql.Open("pgx", OrmSource)
	if err != nil {
		return err
	}

	sqlc.conn = db.New(sqlc.db)

	return nil
}

func (sqlc *Sqlc) Close() error {
	return sqlc.db.Close()
}

func (sqlc *Sqlc) GetError(method string) string {
	return sqlc.errors[method]
}

func (sqlc *Sqlc) Insert(b *testing.B) {
	m := NewModel()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Id = 0
		_, err := sqlc.conn.CreateModel(ctx, db.CreateModelParams{
			Name:    m.Name,
			Title:   m.Title,
			Fax:     m.Fax,
			Web:     m.Web,
			Age:     int32(m.Age),
			Right:   m.Right,
			Counter: m.Counter,
		})
		if err != nil {
			sqlc.error(b, "insert", fmt.Sprintf("sqlc: insert: %v", err))
		}
	}
}

func (sqlc *Sqlc) InsertMulti(b *testing.B) {
	sqlc.error(b, "insert_multi", "sqlc: insert_multi: insert multi is not supported on sqlc")
}

func (sqlc *Sqlc) Update(b *testing.B) {
	m := NewModel()

	_, err := sqlc.conn.CreateModel(ctx, db.CreateModelParams{
		Name:    m.Name,
		Title:   m.Title,
		Fax:     m.Fax,
		Web:     m.Web,
		Age:     int32(m.Age),
		Right:   m.Right,
		Counter: m.Counter,
	})
	if err != nil {
		sqlc.error(b, "update", fmt.Sprintf("sqlc: update: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := sqlc.conn.UpdateModel(ctx, db.UpdateModelParams{
			Name:    m.Name,
			Title:   m.Title,
			Fax:     m.Fax,
			Web:     m.Web,
			Age:     int32(m.Age),
			Right:   m.Right,
			Counter: m.Counter,
			ID:      int32(m.Id),
		})
		if err != nil {
			sqlc.error(b, "update", fmt.Sprintf("sqlc: update: %v", err))
		}
	}
}

func (sqlc *Sqlc) Read(b *testing.B) {
	m := NewModel()

	output, err := sqlc.conn.CreateModel(ctx, db.CreateModelParams{
		Name:    m.Name,
		Title:   m.Title,
		Fax:     m.Fax,
		Web:     m.Web,
		Age:     int32(m.Age),
		Right:   m.Right,
		Counter: m.Counter,
	})
	m.Id = int(output.ID)
	if err != nil {
		sqlc.error(b, "read", fmt.Sprintf("sqlc: read: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := sqlc.conn.GetModel(ctx, int32(m.Id))
		if err != nil {
			sqlc.error(b, "read", fmt.Sprintf("sqlc: read: %v", err))
		}
	}
}

func (sqlc *Sqlc) ReadSlice(b *testing.B) {
	m := NewModel()

	for i := 0; i < 100; i++ {
		m.Id = 0

		_, err := sqlc.conn.CreateModel(ctx, db.CreateModelParams{
			Name:    m.Name,
			Title:   m.Title,
			Fax:     m.Fax,
			Web:     m.Web,
			Age:     int32(m.Age),
			Right:   m.Right,
			Counter: m.Counter,
		})
		if err != nil {
			sqlc.error(b, "read_slice", fmt.Sprintf("sqlc: read_slice: %v", err))
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := sqlc.conn.ListModels(ctx, db.ListModelsParams{
			ID:    0,
			Limit: 100,
		})
		if err != nil {
			sqlc.error(b, "read_slice", fmt.Sprintf("sqlc: read_slice: %v", err))
		}
	}
}

func (sqlc *Sqlc) error(b *testing.B, method string, err string) {
	b.Helper()

	sqlc.mu.Lock()
	sqlc.errors[method] = err
	sqlc.mu.Unlock()
	b.Fail()
}

/*
func SqlcReadSlice(b *B) {
	var m *Model

	WrapExecute(b, func() {
		InitDB()
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
			CheckErr(err, b)

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
		CheckErr(err, b)

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
