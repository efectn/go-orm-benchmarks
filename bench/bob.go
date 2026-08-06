package bench

import (
	"testing"

	"github.com/efectn/go-orm-benchmarks/helper"
	_ "github.com/jackc/pgx/v5/stdlib"
	boblib "github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

type Bob struct {
	helper.ORMInterface
	conn boblib.DB
}

var _ helper.ORMInterface = (*Bob)(nil)

func CreateBob() helper.ORMInterface {
	return &Bob{}
}

func (bob *Bob) Name() string {
	return "bob"
}

func (bob *Bob) Init() error {
	var err error
	bob.conn, err = boblib.Open("pgx", helper.OrmSource)
	if err != nil {
		return err
	}

	bob.conn.SetMaxOpenConns(helper.OrmMaxConn)
	bob.conn.SetMaxIdleConns(helper.OrmMaxIdle)
	return nil
}

func (bob *Bob) Close() error {
	return bob.conn.Close()
}

func (bob *Bob) Insert(b *testing.B) {
	setter := newBobModelSetter(NewModel())

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := bobModels.Insert(setter).One(ctx, bob.conn)
		if err != nil {
			helper.SetError(b, bob.Name(), "Insert", err.Error())
		}
	}
}

func (bob *Bob) InsertMulti(b *testing.B) {
	setters := make([]*bobModelSetter, 100)
	for i := range setters {
		setters[i] = newBobModelSetter(NewModel())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := bobModels.Insert(boblib.ToMods(setters...)).All(ctx, bob.conn)
		if err != nil {
			helper.SetError(b, bob.Name(), "InsertMulti", err.Error())
		}
	}
}

func (bob *Bob) Update(b *testing.B) {
	setter := newBobModelSetter(NewModel())
	model, err := bobModels.Insert(setter).One(ctx, bob.conn)
	if err != nil {
		helper.SetError(b, bob.Name(), "Update", err.Error())
		return
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err = model.Update(ctx, bob.conn, setter)
		if err != nil {
			helper.SetError(b, bob.Name(), "Update", err.Error())
		}
	}
}

func (bob *Bob) Read(b *testing.B) {
	model, err := bobModels.Insert(newBobModelSetter(NewModel())).One(ctx, bob.conn)
	if err != nil {
		helper.SetError(b, bob.Name(), "Read", err.Error())
		return
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err = bobModels.Query(
			sm.Where(psql.Quote("models", "id").EQ(psql.Arg(model.ID))),
		).One(ctx, bob.conn)
		if err != nil {
			helper.SetError(b, bob.Name(), "Read", err.Error())
		}
	}
}

func (bob *Bob) ReadSlice(b *testing.B) {
	setter := newBobModelSetter(NewModel())
	for i := 0; i < 100; i++ {
		_, err := bobModels.Insert(setter).One(ctx, bob.conn)
		if err != nil {
			helper.SetError(b, bob.Name(), "ReadSlice", err.Error())
			return
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := bobModels.Query(
			sm.Where(psql.Quote("models", "id").GT(psql.Arg(0))),
			sm.Limit(100),
		).All(ctx, bob.conn)
		if err != nil {
			helper.SetError(b, bob.Name(), "ReadSlice", err.Error())
		}
	}
}
