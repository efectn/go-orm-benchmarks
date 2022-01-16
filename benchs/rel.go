package benchs

import (
	"context"
	"log"

	"github.com/go-rel/postgres"
	relware "github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
)

var rel relware.Repository

func initDB11() {
	conn, err := postgres.Open(OrmSource)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	rel = relware.New(conn)
	if err := rel.Ping(ctx); err != nil {
		log.Fatalf("failed pinging connection: %v", err)
	}

	// Disable debug logging
	rel.Instrumentation(func(ctx context.Context, op string, message string) func(err error) {
		return func(err error) {
			if err != nil {
				log.Fatalf("rel: error: %v", err)
			}
		}
	})

	// Run the auto migration tool.
	/*if _, err = pop.TX.Exec("DROP TABLE IF EXISTS popmodels"); err != nil {
		log.Fatal(err)
	}*/
	initDB()
}

func init() {
	st := NewSuite("rel")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, RelInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, RelInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, RelUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, RelRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, RelReadSlice)
	}
}

func RelInsert(b *B) {
	var m *PopModel
	wrapExecute(b, func() {
		initDB11()
		m = NewPopModel()
	})

	for i := 0; i < b.N; i++ {
		m.ID = 0
		if err := rel.Insert(ctx, m); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}

func RelInsertMulti(b *B) {
	var ms []*PopModel
	wrapExecute(b, func() {
		initDB()
		ms = make([]*PopModel, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewPopModel())
		}
	})

	for i := 0; i < b.N; i++ {
		for _, m := range ms {
			m.ID = 0
		}
		if err := rel.InsertAll(ctx, &ms); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}

func RelUpdate(b *B) {
	var m *PopModel
	wrapExecute(b, func() {
		initDB11()
		m = NewPopModel()
		m.ID = 0
		if err := rel.Insert(ctx, m); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if err := rel.Update(ctx, m); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}

func RelRead(b *B) {
	var m *PopModel
	wrapExecute(b, func() {
		initDB11()
		m = NewPopModel()
		m.ID = 0
		if err := rel.Insert(ctx, m); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if err := rel.Find(ctx, m); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}

func RelReadSlice(b *B) {
	var m *PopModel
	wrapExecute(b, func() {
		initDB11()
		m = NewPopModel()
		for i := 0; i < 100; i++ {
			m.ID = 0
			if err := rel.Insert(ctx, m); err != nil {
				log.Fatal(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var ms []PopModel
		if err := rel.FindAll(ctx, &ms, where.Gt("id", 0), relware.Limit(100)); err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
}
