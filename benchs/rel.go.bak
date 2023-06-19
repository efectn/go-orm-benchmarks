package benchs

import (
	"context"
	"log"

	"github.com/go-rel/postgres"
	relware "github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
)

var rel relware.Repository

func init() {
	st := NewSuite("rel")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, RelInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, RelInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, RelUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, RelRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, RelReadSlice)

		conn, err := postgres.Open(OrmSource)
		if err != nil {
			log.Fatalf("failed opening connection to postgres: %v", err)
		}

		rel = relware.New(conn)
		err = rel.Ping(ctx)
		CheckErr(err)

		// Disable debug logging
		rel.Instrumentation(func(ctx context.Context, op string, message string, args ...any) func(err error) {
			return func(err error) {
				CheckErr(err)
			}
		})
	}
}

func RelInsert(b *B) {
	var m *Model3
	WrapExecute(b, func() {
		InitDB()
		m = NewModel3()
	})

	for i := 0; i < b.N; i++ {
		m.ID = 0
		err := rel.Insert(ctx, m)
		CheckErr(err, b)
	}
}

func RelInsertMulti(b *B) {
	var ms []*Model3
	WrapExecute(b, func() {
		InitDB()
		ms = make([]*Model3, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewModel3())
		}
	})

	for i := 0; i < b.N; i++ {
		for _, m := range ms {
			m.ID = 0
		}
		err := rel.InsertAll(ctx, &ms)
		CheckErr(err, b)
	}
}

func RelUpdate(b *B) {
	var m *Model3
	WrapExecute(b, func() {
		InitDB()
		m = NewModel3()
		m.ID = 0
		err := rel.Insert(ctx, m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		err := rel.Update(ctx, m)
		CheckErr(err, b)
	}
}

func RelRead(b *B) {
	var m *Model3
	WrapExecute(b, func() {
		InitDB()
		m = NewModel3()
		m.ID = 0
		err := rel.Insert(ctx, m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		err := rel.Find(ctx, m)
		CheckErr(err, b)
	}
}

func RelReadSlice(b *B) {
	var m *Model3
	WrapExecute(b, func() {
		InitDB()
		m = NewModel3()
		for i := 0; i < 100; i++ {
			m.ID = 0
			err := rel.Insert(ctx, m)
			CheckErr(err, b)
		}
	})

	for i := 0; i < b.N; i++ {
		var ms []Model3
		err := rel.FindAll(ctx, &ms, where.Gt("id", 0), relware.Limit(100))
		CheckErr(err, b)
	}
}
