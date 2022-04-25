package benchs

import (
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var pgxpdb *pgxpool.Pool

func init() {
	st := NewSuite("pgx_pool")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, PgxPoolInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, PgxPoolInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, PgxPoolUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, PgxPoolRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, PgxPoolReadSlice)

		db, err := pgxpool.Connect(ctx, OrmSource)
		CheckErr(err)

		pgxpdb = db
	}
}

func PgxPoolInsert(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		_, err := pgxpdb.Exec(ctx, sqlxInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		CheckErr(err, b)
	}
}

func pgxPoolInsert(m *Model) error {
	_, err := pgxpdb.Exec(ctx, sqlxInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
	CheckErr(err)

	return nil
}

func PgxPoolInsertMulti(b *B) {
	var rows = make([][]interface{}, 0)
	WrapExecute(b, func() {
		InitDB()
		m := NewModel()
		for i := 0; i < 100; i++ {
			rows = append(rows, []interface{}{m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter})
		}
	})

	for i := 0; i < b.N; i++ {
		_, err := pgxpdb.CopyFrom(ctx, pgx.Identifier{"models"}, columns, pgx.CopyFromRows(rows))
		CheckErr(err, b)
	}
}

func PgxPoolUpdate(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		m.Id = 1
		err := pgxPoolInsert(m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		_, err := pgxpdb.Exec(ctx, sqlxUpdateSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter, m.Id)
		CheckErr(err, b)
	}
}

func PgxPoolRead(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		err := pgxPoolInsert(m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		var m Model
		err := pgxpdb.QueryRow(ctx, sqlxSelectSQL, 1).Scan(
			&m.Id,
			&m.Name,
			&m.Title,
			&m.Fax,
			&m.Web,
			&m.Age,
			&m.Right,
			&m.Counter,
		)
		CheckErr(err, b)
	}
}

func PgxPoolReadSlice(b *B) {
	var m *Model
	WrapExecute(b, func() {
		var err error
		InitDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			err = pgxPoolInsert(m)
			CheckErr(err, b)
		}
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		ms := make([]Model, 100)
		rows, err := pgxpdb.Query(ctx, sqlxSelectMultiSQL)
		CheckErr(err, b)

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
			CheckErr(err, b)
		}
		err = rows.Err()
		CheckErr(err, b)
		rows.Close()
	}
}
