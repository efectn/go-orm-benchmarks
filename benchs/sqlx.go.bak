package benchs

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var sqlxdb *sqlx.DB

const (
	sqlxInsertBaseSQL   = `INSERT INTO models (name, title, fax, web, age, "right", counter) VALUES `
	sqlxInsertValuesSQL = `($1, $2, $3, $4, $5, $6, $7)`
	sqlxInsertSQL       = sqlxInsertBaseSQL + sqlxInsertValuesSQL
	sqlxInsertNamesSQL  = `(:name, :title, :fax, :web, :age, :right, :counter)`
	sqlxInsertMultiSQL  = sqlxInsertBaseSQL + sqlxInsertNamesSQL
	sqlxUpdateSQL       = `UPDATE models SET name = $1, title = $2, fax = $3, web = $4, age = $5, "right" = $6, counter = $7 WHERE id = $8`
	sqlxSelectSQL       = `SELECT * FROM models WHERE id = $1`
	sqlxSelectMultiSQL  = `SELECT * FROM models WHERE id > 0 LIMIT 100`
)

func init() {
	st := NewSuite("sqlx")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, SqlxInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, SqlxInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, SqlxUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, SqlxRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, SqlxReadSlice)

		db, err := sqlx.Connect("postgres", OrmSource)
		CheckErr(err)

		sqlxdb = db
	}
}

func SqlxInsert(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		_, err := sqlxdb.Exec(sqlxInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		CheckErr(err, b)
	}
}

func sqlxInsert(m *Model) error {
	_, err := sqlxdb.Exec(sqlxInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
	CheckErr(err)

	return nil
}

func SqlxInsertMulti(b *B) {
	var ms []*Model
	WrapExecute(b, func() {
		InitDB()
		ms = make([]*Model, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewModel())
		}
	})

	for i := 0; i < b.N; i++ {
		_, err := sqlxdb.NamedExec(sqlxInsertMultiSQL, ms)
		CheckErr(err, b)
	}
}

func SqlxUpdate(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		err := sqlxInsert(m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		_, err := sqlxdb.Exec(sqlxUpdateSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter, m.Id)
		CheckErr(err, b)
	}
}

func SqlxRead(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		err := sqlxInsert(m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		m := []Model{}
		err := sqlxdb.Select(&m, sqlxSelectSQL, 1)
		CheckErr(err, b)
	}
}

func SqlxReadSlice(b *B) {
	var m *Model
	WrapExecute(b, func() {
		var err error
		InitDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			err = sqlxInsert(m)
			CheckErr(err, b)
		}
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		ms := make([]Model, 100)
		err := sqlxdb.Select(&ms, sqlxSelectMultiSQL)
		CheckErr(err, b)
	}
}
