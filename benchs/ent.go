package benchs

import (
	"database/sql"
	"fmt"
	"sync"
	"testing"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	entdb "github.com/efectn/go-orm-benchmarks/benchs/ent"
	"github.com/efectn/go-orm-benchmarks/benchs/ent/model"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Ent struct {
	Instance
	mu         sync.Mutex
	conn       *entdb.Client
	dbEnt      *sql.DB
	iterations int // Same as b.N, just to customize it
	errors     map[string]string
}

func CreateEnt(iterations int) Instance {
	ent := &Ent{
		iterations: iterations,
		errors:     map[string]string{},
	}

	return ent
}

func (ent *Ent) Name() string {
	return "ent"
}

func (ent *Ent) Init() error {
	var err error
	ent.dbEnt, err = sql.Open("pgx", OrmSource)
	if err != nil {
		return err
	}

	// Create an ent.Driver from `dbEnt`.
	drv := entsql.OpenDB(dialect.Postgres, ent.dbEnt)

	// Assign to client
	ent.conn = entdb.NewClient(entdb.Driver(drv))

	/*if _, err := dbEnt.Exec("DROP TABLE IF EXISTS models"); err != nil {
		return nil
	}

	if err := client.Schema.Create(ctx); err != nil {
		return nil
	}*/

	return nil
}

func (ent *Ent) Close() error {
	return ent.conn.Close()
}

func (ent *Ent) GetError(method string) string {
	return ent.errors[method]
}

func (ent *Ent) Insert(b *testing.B) {
	m := NewModelAlt()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := ent.conn.Model.Create().
			SetName(m.Name).
			SetTitle(m.Title).
			SetFax(m.Fax).
			SetWeb(m.Web).
			SetAge(m.Age).
			SetRight(m.Right).
			SetCounter(m.Counter).
			Save(ctx)
		if err != nil {
			ent.error(b, "insert", fmt.Sprintf("ent: insert: %v", err))
		}
	}
}

func (ent *Ent) InsertMulti(b *testing.B) {
	ms := make([]Model, 0, 100)
	for i := 0; i < 100; i++ {
		ms = append(ms, NewModelAlt())
	}

	b.ReportAllocs()
	b.ResetTimer()

	bulk := make([]*entdb.ModelCreate, len(ms))
	for i, m := range ms {
		bulk[i] = ent.conn.Model.Create().
			SetName(m.Name).
			SetTitle(m.Title).
			SetFax(m.Fax).
			SetWeb(m.Web).
			SetAge(m.Age).
			SetRight(m.Right).
			SetCounter(m.Counter)
	}

	for i := 0; i < b.N; i++ {
		_, err := ent.conn.Model.CreateBulk(bulk...).Save(ctx)
		if err != nil {
			ent.error(b, "insert_multi", fmt.Sprintf("ent: insert_multi: %v", err))
		}
	}
}

func (ent *Ent) Update(b *testing.B) {
	m := NewModelAlt()

	_, err := ent.conn.Model.Create().
		SetName(m.Name).
		SetTitle(m.Title).
		SetFax(m.Fax).
		SetWeb(m.Web).
		SetAge(m.Age).
		SetRight(m.Right).
		SetCounter(m.Counter).
		Save(ctx)
	if err != nil {
		ent.error(b, "update", fmt.Sprintf("ent: update: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := ent.conn.Model.Update().
			Where(model.IDEQ(m.Id)).
			SetName(m.Name).
			SetTitle(m.Title).
			SetFax(m.Fax).
			SetWeb(m.Web).
			SetAge(m.Age).
			SetRight(m.Right).
			SetCounter(m.Counter).
			Save(ctx)
		if err != nil {
			ent.error(b, "update", fmt.Sprintf("ent: update: %v", err))
		}
	}
}

func (ent *Ent) Read(b *testing.B) {
	m := NewModelAlt()

	_, err := ent.conn.Model.Create().
		SetName(m.Name).
		SetTitle(m.Title).
		SetFax(m.Fax).
		SetWeb(m.Web).
		SetAge(m.Age).
		SetRight(m.Right).
		SetCounter(m.Counter).
		Save(ctx)
	if err != nil {
		ent.error(b, "read", fmt.Sprintf("ent: read: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := ent.conn.Model.Query().Where(model.IDEQ(1)).First(ctx)
		if err != nil {
			ent.error(b, "read", fmt.Sprintf("ent: read: %v", err))
		}
	}
}

func (ent *Ent) ReadSlice(b *testing.B) {
	m := NewModelAlt()
	for i := 0; i < 100; i++ {
		_, err := ent.conn.Model.Create().
			SetName(m.Name).
			SetTitle(m.Title).
			SetFax(m.Fax).
			SetWeb(m.Web).
			SetAge(m.Age).
			SetRight(m.Right).
			SetCounter(m.Counter).
			Save(ctx)
		if err != nil {
			ent.error(b, "read_slice", fmt.Sprintf("ent: read_slice: %v", err))
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := ent.conn.Model.Query().Where(model.IDGT(0)).Unique(false).Limit(100).All(ctx)
		if err != nil {
			ent.error(b, "read_slice", fmt.Sprintf("ent: read_slice: %v", err))
		}
	}
}

func (ent *Ent) error(b *testing.B, method string, err string) {
	b.Helper()

	ent.mu.Lock()
	ent.errors[method] = err
	ent.mu.Unlock()
	b.Fail()
}
