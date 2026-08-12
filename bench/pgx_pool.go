package bench

import (
	"testing"

	"github.com/efectn/go-orm-benchmarks/helper"
	pgxdb "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxPool struct {
	helper.ORMInterface
	conn *pgxpool.Pool
}

func CreatePgxPool() helper.ORMInterface {
	return &PgxPool{}
}

func (pgx *PgxPool) Name() string {
	return "pgx_pool"
}

func (pgx *PgxPool) Init() error {
	var err error
	pgx.conn, err = pgxpool.New(ctx, helper.OrmSource)
	if err != nil {
		return err
	}

	return nil
}

func (pgx *PgxPool) Close() error {
	pgx.conn.Close()

	return nil
}

func (pgx *PgxPool) Insert(b *testing.B) {
	m := NewModel()
	helper.BenchExec(b, pgx.Name(), "Insert", func() error {
		_, err := pgx.conn.Exec(ctx, sqlxInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		return err
	})
}

func (pgx *PgxPool) InsertMulti(b *testing.B) {
	var rows = make([][]interface{}, 0)

	m := NewModel()
	for i := 0; i < helper.BatchSize; i++ {
		rows = append(rows, []interface{}{m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter})
	}

	helper.BenchExec(b, pgx.Name(), "InsertMulti", func() error {
		_, err := pgx.conn.CopyFrom(ctx, pgxdb.Identifier{"models"}, columns, pgxdb.CopyFromRows(rows))
		return err
	})
}

func (pgx *PgxPool) Update(b *testing.B) {
	m := NewModel()
	err := pgx.conn.QueryRow(ctx, sqlxInsertReturningSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter).Scan(&m.Id)
	if err != nil {
		helper.SetError(b, pgx.Name(), "Update", err.Error())
		return
	}
	m.Counter = helper.UpdatedCounter

	helper.BenchExec(b, pgx.Name(), "Update", func() error {
		_, err := pgx.conn.Exec(ctx, sqlxUpdateSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter, m.Id)
		return err
	})
}

func (pgx *PgxPool) Read(b *testing.B) {
	m := NewModel()
	err := pgx.conn.QueryRow(ctx, sqlxInsertReturningSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter).Scan(&m.Id)
	if err != nil {
		helper.SetError(b, pgx.Name(), "Read", err.Error())
		return
	}

	helper.BenchRows(b, pgx.Name(), "Read", 1, func() (int, error) {
		var result Model
		err := pgx.conn.QueryRow(ctx, sqlxSelectSQL, m.Id).Scan(
			&result.Id,
			&result.Name,
			&result.Title,
			&result.Fax,
			&result.Web,
			&result.Age,
			&result.Right,
			&result.Counter,
		)
		if err != nil {
			return 0, err
		}
		return 1, nil
	})
}

func (pgx *PgxPool) ReadSlice(b *testing.B) {
	m := NewModel()
	for i := 0; i < helper.BatchSize; i++ {
		_, err := pgx.conn.Exec(ctx, sqlxInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		if err != nil {
			helper.SetError(b, pgx.Name(), "ReadSlice", err.Error())
			return
		}
	}

	helper.BenchRows(b, pgx.Name(), "ReadSlice", helper.BatchSize, func() (int, error) {
		ms := make([]Model, helper.BatchSize)
		rows, err := pgx.conn.Query(ctx, sqlxSelectMultiSQL, helper.BatchSize)
		if err != nil {
			return 0, err
		}
		defer rows.Close()

		j := 0
		for ; rows.Next() && j < len(ms); j++ {
			err = rows.Scan(
				&ms[j].Id,
				&ms[j].Name,
				&ms[j].Title,
				&ms[j].Fax,
				&ms[j].Web,
				&ms[j].Age,
				&ms[j].Right,
				&ms[j].Counter,
			)
			if err != nil {
				return j, err
			}
		}
		return j, rows.Err()
	})
}
