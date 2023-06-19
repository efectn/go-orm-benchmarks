package benchs

import (
	"github.com/jackc/pgx/v4"
)

var pgxdb *pgx.Conn

func init() {
	st := NewSuite("pgx")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, PgxInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, PgxInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, PgxUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, PgxRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, PgxReadSlice)

		db, err := pgx.Connect(ctx, OrmSource)
		CheckErr(err)

		pgxdb = db
	}
}

func PgxInsert(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		_, err := pgxdb.Exec(ctx, sqlxInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		CheckErr(err, b)
	}
}

func pgxInsert(m *Model) error {
	_, err := pgxdb.Exec(ctx, sqlxInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
	CheckErr(err)

	return nil
}

func PgxInsertMulti(b *B) {
	var rows = make([][]interface{}, 0)
	WrapExecute(b, func() {
		InitDB()
		m := NewModel()
		for i := 0; i < 100; i++ {
			rows = append(rows, []interface{}{m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter})
		}
	})

	for i := 0; i < b.N; i++ {
		_, err := pgxdb.CopyFrom(ctx, pgx.Identifier{"models"}, columns, pgx.CopyFromRows(rows))
		CheckErr(err, b)
	}
}

func PgxUpdate(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		m.Id = 1
		err := pgxInsert(m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		_, err := pgxdb.Exec(ctx, sqlxUpdateSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter, m.Id)
		CheckErr(err, b)
	}
}

func PgxRead(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		err := pgxInsert(m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		var m Model
		err := pgxdb.QueryRow(ctx, sqlxSelectSQL, 1).Scan(
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

func PgxReadSlice(b *B) {
	var m *Model
	WrapExecute(b, func() {
		var err error
		InitDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			err = pgxInsert(m)
			CheckErr(err, b)
		}
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		ms := make([]Model, 100)
		rows, err := pgxdb.Query(ctx, sqlxSelectMultiSQL)
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
