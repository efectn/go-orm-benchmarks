package bench

import (
	"github.com/efectn/go-orm-benchmarks/helper"
	pgxdb "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"testing"
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
	pgx.conn, err = pgxpool.Connect(ctx, helper.OrmSource)
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

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := pgx.conn.Exec(ctx, sqlxInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		if err != nil {
			helper.SetError(b, pgx.Name(), "Insert", err.Error())
		}
	}
}

func (pgx *PgxPool) InsertMulti(b *testing.B) {
	var rows = make([][]interface{}, 0)

	m := NewModel()
	for i := 0; i < 100; i++ {
		rows = append(rows, []interface{}{m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter})
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := pgx.conn.CopyFrom(ctx, pgxdb.Identifier{"models"}, columns, pgxdb.CopyFromRows(rows))
		if err != nil {
			helper.SetError(b, pgx.Name(), "InsertMulti", err.Error())
		}
	}
}

func (pgx *PgxPool) Update(b *testing.B) {
	m := NewModel()
	m.Id = 1

	_, err := pgx.conn.Exec(ctx, sqlxInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
	if err != nil {
		helper.SetError(b, pgx.Name(), "Update", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := pgx.conn.Exec(ctx, sqlxUpdateSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter, m.Id)
		if err != nil {
			helper.SetError(b, pgx.Name(), "Update", err.Error())
		}
	}
}

func (pgx *PgxPool) Read(b *testing.B) {
	m := NewModel()
	m.Id = 1

	_, err := pgx.conn.Exec(ctx, sqlxInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
	if err != nil {
		helper.SetError(b, pgx.Name(), "Read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var m Model
		err := pgx.conn.QueryRow(ctx, sqlxSelectSQL, 1).Scan(
			&m.Id,
			&m.Name,
			&m.Title,
			&m.Fax,
			&m.Web,
			&m.Age,
			&m.Right,
			&m.Counter,
		)
		if err != nil {
			helper.SetError(b, pgx.Name(), "Read", err.Error())
		}
	}
}

func (pgx *PgxPool) ReadSlice(b *testing.B) {
	m := NewModel()
	for i := 0; i < 100; i++ {
		_, err := pgx.conn.Exec(ctx, sqlxInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		if err != nil {
			helper.SetError(b, pgx.Name(), "ReadSlice", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ms := make([]Model, 100)
		rows, err := pgx.conn.Query(ctx, sqlxSelectMultiSQL)
		if err != nil {
			helper.SetError(b, pgx.Name(), "ReadSlice", err.Error())
		}

		for j := 0; rows.Next() && j < len(ms); j++ {
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
				helper.SetError(b, pgx.Name(), "ReadSlice", err.Error())
			}
		}
		err = rows.Err()
		if err != nil {
			helper.SetError(b, pgx.Name(), "ReadSlice", err.Error())
		}

		rows.Close()
	}
}
