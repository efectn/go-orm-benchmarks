package benchs

/*
import (
	"fmt"
	"log"

	"github.com/efectn/orm-benchmark/benchs/prisma/db"
)

var prisma *db.PrismaClient

func initDB5() {
	prisma = db.NewClient()
	if err := prisma.Prisma.Connect(); err != nil {
		log.Fatal(err)
	}
}

func closePrisma() {
	if err := prisma.Prisma.Disconnect(); err != nil {
		log.Fatal(err)
	}
}

func NewPrismaModel() db.InnerModelsPrisma {
	m := db.InnerModelsPrisma{}
	m.Name = "Orm Benchmark"
	m.Title = "Just a Benchmark for fun"
	m.Fax = "99909990"
	m.Web = "http://blog.milkpod29.me"
	m.Age = 100
	m.Right = true
	m.Counter = 1000

	return m
}

func init() {
	st := NewSuite("prisma")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, PrismaInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, PrismaInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, PrismaUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, PrismaRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, PrismaReadSlice)
	}
}

func PrismaInsert(b *B) {
	var m db.InnerModelsPrisma
	wrapExecute(b, func() {
		initDB5()
		m = NewPrismaModel()
	})

	for i := 0; i < b.N; i++ {
		_, err := prisma.ModelsPrisma.CreateOne(
			db.ModelsPrisma.Name.Set(m.Name),
			db.ModelsPrisma.Title.Set(m.Title),
			db.ModelsPrisma.Fax.Set(m.Fax),
			db.ModelsPrisma.Web.Set(m.Web),
			db.ModelsPrisma.Age.Set(m.Age),
			db.ModelsPrisma.Right.Set(m.Right),
			db.ModelsPrisma.Counter.Set(m.Counter),
		).Exec(ctx)

		if err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
	closePrisma()
}

func PrismaInsertMulti(b *B) {
	panic(fmt.Errorf("doesn't work"))
}

func PrismaUpdate(b *B) {
	var m db.InnerModelsPrisma
	wrapExecute(b, func() {
		initDB5()
		m = NewPrismaModel()
		_, err := prisma.ModelsPrisma.CreateOne(
			db.ModelsPrisma.Name.Set(m.Name),
			db.ModelsPrisma.Title.Set(m.Title),
			db.ModelsPrisma.Fax.Set(m.Fax),
			db.ModelsPrisma.Web.Set(m.Web),
			db.ModelsPrisma.Age.Set(m.Age),
			db.ModelsPrisma.Right.Set(m.Right),
			db.ModelsPrisma.Counter.Set(m.Counter),
		).Exec(ctx)

		if err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		_, err := prisma.ModelsPrisma.FindMany(
			db.ModelsPrisma.ID.Equals(m.ID),
		).Update(
			db.ModelsPrisma.Name.Set(m.Name),
			db.ModelsPrisma.Title.Set(m.Title),
			db.ModelsPrisma.Fax.Set(m.Fax),
			db.ModelsPrisma.Web.Set(m.Web),
			db.ModelsPrisma.Age.Set(m.Age),
			db.ModelsPrisma.Right.Set(m.Right),
			db.ModelsPrisma.Counter.Set(m.Counter),
		).Exec(ctx)

		if err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
	closePrisma()
}

func PrismaRead(b *B) {
	var m db.InnerModelsPrisma
	wrapExecute(b, func() {
		initDB5()
		m = NewPrismaModel()
		_, err := prisma.ModelsPrisma.CreateOne(
			db.ModelsPrisma.Name.Set(m.Name),
			db.ModelsPrisma.Title.Set(m.Title),
			db.ModelsPrisma.Fax.Set(m.Fax),
			db.ModelsPrisma.Web.Set(m.Web),
			db.ModelsPrisma.Age.Set(m.Age),
			db.ModelsPrisma.Right.Set(m.Right),
			db.ModelsPrisma.Counter.Set(m.Counter),
		).Exec(ctx)

		if err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		m.ID = 1
		_, err := prisma.ModelsPrisma.FindFirst(db.ModelsPrisma.ID.Equals(m.ID)).Exec(ctx)
		if err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
	closePrisma()
}

func PrismaReadSlice(b *B) {
	var m db.InnerModelsPrisma
	wrapExecute(b, func() {
		initDB5()
		m = NewPrismaModel()
		for i := 0; i < 100; i++ {
			_, err := prisma.ModelsPrisma.CreateOne(
				db.ModelsPrisma.Name.Set(m.Name),
				db.ModelsPrisma.Title.Set(m.Title),
				db.ModelsPrisma.Fax.Set(m.Fax),
				db.ModelsPrisma.Web.Set(m.Web),
				db.ModelsPrisma.Age.Set(m.Age),
				db.ModelsPrisma.Right.Set(m.Right),
				db.ModelsPrisma.Counter.Set(m.Counter),
			).Exec(ctx)
			if err != nil {
				log.Fatal(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		_, err := prisma.ModelsPrisma.FindMany().OrderBy(db.ModelsPrisma.ID.Order(db.SortOrderDesc)).Take(100).Exec(ctx)
		if err != nil {
			log.Fatal(err)
			b.FailNow()
		}
	}
	closePrisma()
}
*/
