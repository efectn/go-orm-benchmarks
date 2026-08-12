package bench

import (
	"testing"

	"github.com/efectn/go-orm-benchmarks/helper"
	_ "github.com/jackc/pgx/v5/stdlib"
	boblib "github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/sm"
	"github.com/stephenafamo/bob/dialect/psql/um"
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
	helper.BenchExec(b, bob.Name(), "Insert", func() error {
		_, err := bobModels.Insert(setter).Exec(ctx, bob.conn)
		return err
	})
}

func (bob *Bob) InsertMulti(b *testing.B) {
	setters := make([]*bobModelSetter, helper.BatchSize)
	for i := range setters {
		setters[i] = newBobModelSetter(NewModel())
	}

	helper.BenchExec(b, bob.Name(), "InsertMulti", func() error {
		_, err := bobModels.Insert(boblib.ToMods(setters...)).Exec(ctx, bob.conn)
		return err
	})
}

func (bob *Bob) Update(b *testing.B) {
	setter := newBobModelSetter(NewModel())
	model, err := bobModels.Insert(setter).One(ctx, bob.conn)
	if err != nil {
		helper.SetError(b, bob.Name(), "Update", err.Error())
		return
	}
	updated := NewModel()
	updated.Counter = helper.UpdatedCounter
	setter = newBobModelSetter(updated)

	helper.BenchExec(b, bob.Name(), "Update", func() error {
		_, err = bobModels.Update(
			setter.UpdateMod(),
			um.Where(model.primaryKeyEquals()),
		).Exec(ctx, bob.conn)
		return err
	})
}

func (bob *Bob) Read(b *testing.B) {
	model, err := bobModels.Insert(newBobModelSetter(NewModel())).One(ctx, bob.conn)
	if err != nil {
		helper.SetError(b, bob.Name(), "Read", err.Error())
		return
	}

	helper.BenchRows(b, bob.Name(), "Read", 1, func() (int, error) {
		_, err = bobModels.Query(
			sm.Where(psql.Quote("models", "id").EQ(psql.Arg(model.ID))),
		).One(ctx, bob.conn)
		if err != nil {
			return 0, err
		}
		return 1, nil
	})
}

func (bob *Bob) ReadSlice(b *testing.B) {
	setter := newBobModelSetter(NewModel())
	for i := 0; i < helper.BatchSize; i++ {
		_, err := bobModels.Insert(setter).One(ctx, bob.conn)
		if err != nil {
			helper.SetError(b, bob.Name(), "ReadSlice", err.Error())
			return
		}
	}

	helper.BenchRows(b, bob.Name(), "ReadSlice", helper.BatchSize, func() (int, error) {
		models, err := bobModels.Query(
			sm.Where(psql.Quote("models", "id").GT(psql.Arg(0))),
			sm.Limit(helper.BatchSize),
		).All(ctx, bob.conn)
		return len(models), err
	})
}
