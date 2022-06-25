package benchs

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
)

var raw *sql.DB

const (
	rawInsertBaseSQL   = `INSERT INTO models (name, title, fax, web, age, "right", counter) VALUES `
	rawInsertValuesSQL = `($1, $2, $3, $4, $5, $6, $7)`
	rawInsertSQL       = rawInsertBaseSQL + rawInsertValuesSQL + " RETURNING id"
	rawUpdateSQL       = `UPDATE models SET name = $1, title = $2, fax = $3, web = $4, age = $5, "right" = $6, counter = $7 WHERE id = $8`
	rawSelectSQL       = `SELECT id, name, title, fax, web, age, "right", counter FROM models WHERE id = $1`
	rawSelectMultiSQL  = `SELECT id, name, title, fax, web, age, "right", counter FROM models WHERE id > 0 LIMIT 100`
)

func init() {
	st := NewSuite("raw")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, RawInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, RawInsertMulti)
		st.AddBenchmark("Update", 200*OrmMulti, RawUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, RawRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, RawReadSlice)

		raw, _ = sql.Open("pgx", OrmSource)
	}
}

func RawInsert(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		rows, err := raw.Query(rawInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		CheckErr(err, b)

		if !rows.Next() {
			log.Printf("[go-orm-benchmarks] ERR: No rows inserted")
			b.FailNow()
		}

		err = rows.Scan(&m.Id)
		CheckErr(err, b)
	}
}

func rawInsert(m *Model) error {
	rows, err := raw.Query(rawInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
	if err != nil {
		return err
	}

	if !rows.Next() {
		return errors.New("no rows inserted")
	}

	err = rows.Scan(&m.Id)
	if err != nil {
		return err
	}

	return nil
}

func RawInsertMulti(b *B) {
	var ms []*Model
	WrapExecute(b, func() {
		InitDB()

		ms = make([]*Model, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewModel())
		}
	})

	var valuesSQL string
	counter := 1
	for i := 0; i < 100; i++ {
		hoge := ""
		for j := 0; j < 7; j++ {
			if j != 6 {
				hoge += "$" + strconv.Itoa(counter) + ","
			} else {
				hoge += "$" + strconv.Itoa(counter)
			}
			counter++

		}
		if i != 99 {
			valuesSQL += "(" + hoge + "),"
		} else {
			valuesSQL += "(" + hoge + ")"
		}
	}

	for i := 0; i < b.N; i++ {
		nFields := 7
		query := rawInsertBaseSQL + valuesSQL
		args := make([]interface{}, len(ms)*nFields)
		for j := range ms {
			offset := j * nFields
			args[offset+0] = ms[j].Name
			args[offset+1] = ms[j].Title
			args[offset+2] = ms[j].Fax
			args[offset+3] = ms[j].Web
			args[offset+4] = ms[j].Age
			args[offset+5] = ms[j].Right
			args[offset+6] = ms[j].Counter
		}
		rows, err := raw.Query(query, args...)
		CheckErr(err, b)

		for j := range ms {
			if !rows.Next() {
				log.Printf("[go-orm-benchmarks] ERR: Not enough rows inserted")
				b.FailNow()
			}

			err = rows.Scan(&ms[j].Id)
			CheckErr(err, b)
		}
	}
}

func RawUpdate(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		err := rawInsert(m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		_, err := raw.Exec(rawUpdateSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter, m.Id)
		CheckErr(err, b)
	}
}

func RawRead(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		err := rawInsert(m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		var mout Model
		err := raw.QueryRow(rawSelectSQL, 1).Scan(
			//err := stmt.QueryRow(m.Id).Scan(
			&mout.Id,
			&mout.Name,
			&mout.Title,
			&mout.Fax,
			&mout.Web,
			&mout.Age,
			&mout.Right,
			&mout.Counter,
		)
		CheckErr(err, b)
	}
}

func RawReadSlice(b *B) {
	var m *Model
	WrapExecute(b, func() {
		var err error
		InitDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			err = rawInsert(m)
			CheckErr(err, b)
		}
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		var j int
		models := make([]Model, 100)
		rows, err := raw.Query(rawSelectMultiSQL)
		CheckErr(err, b)
		for j = 0; rows.Next() && j < len(models); j++ {
			err = rows.Scan(
				&models[j].Id,
				&models[j].Name,
				&models[j].Title,
				&models[j].Fax,
				&models[j].Web,
				&models[j].Age,
				&models[j].Right,
				&models[j].Counter,
			)
			CheckErr(err, b)
		}
		err = rows.Err()
		CheckErr(err, b)

		err = rows.Close()
		CheckErr(err, b)
	}
}
