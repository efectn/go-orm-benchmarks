package benchs

import (
	"database/sql"
	"fmt"
	"strconv"
)

func init() {
	st := NewSuite("raw_stmt")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*ORM_MULTI, RawStmtInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*ORM_MULTI, RawStmtInsertMulti)
		st.AddBenchmark("Update", 200*ORM_MULTI, RawStmtUpdate)
		st.AddBenchmark("Read", 200*ORM_MULTI, RawStmtRead)
		st.AddBenchmark("MultiRead limit 100", 200*ORM_MULTI, RawStmtReadSlice)

		raw, _ = sql.Open("postgres", ORM_SOURCE)
	}
}

func RawStmtInsert(b *B) {
	var m *Model
	var stmt *sql.Stmt
	wrapExecute(b, func() {
		var err error
		initDB()
		m = NewModel()
		stmt, err = raw.Prepare(rawInsertSQL)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})
	defer stmt.Close()

	for i := 0; i < b.N; i++ {
		// pq dose not support the LastInsertId method.
		_, err := stmt.Exec(m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func RawStmtInsertMulti(b *B) {
	var ms []*Model
	var stmt *sql.Stmt
	wrapExecute(b, func() {
		initDB()

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

	stmt, err := raw.Prepare(rawInsertBaseSQL + valuesSQL)
	if err != nil {
		fmt.Println(err)
		b.FailNow()
	}
	defer stmt.Close()

	for i := 0; i < b.N; i++ {
		nFields := 7
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
		// pq dose not support the LastInsertId method.
		_, err := stmt.Exec(args...)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func RawStmtUpdate(b *B) {
	var m *Model
	var stmt *sql.Stmt
	wrapExecute(b, func() {
		var err error
		initDB()
		m = NewModel()
		rawInsert(m)
		stmt, err = raw.Prepare(rawUpdateSQL)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})
	defer stmt.Close()

	for i := 0; i < b.N; i++ {
		_, err := stmt.Exec(m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter, m.Id)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func RawStmtRead(b *B) {
	var m *Model
	var stmt *sql.Stmt
	wrapExecute(b, func() {
		var err error
		initDB()
		m = NewModel()
		rawInsert(m)
		stmt, err = raw.Prepare(rawSelectSQL)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})
	defer stmt.Close()

	for i := 0; i < b.N; i++ {
		var mout Model
		err := stmt.QueryRow(1).Scan(
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
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func RawStmtReadSlice(b *B) {
	var m *Model
	var stmt *sql.Stmt
	wrapExecute(b, func() {
		var err error
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			err = rawInsert(m)
			if err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
		stmt, err = raw.Prepare(rawSelectMultiSQL)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})
	defer stmt.Close()

	for i := 0; i < b.N; i++ {
		var j int
		models := make([]Model, 100)
		rows, err := stmt.Query()
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
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
			if err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
		models = models[:j]
		if err = rows.Err(); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
		if err = rows.Close(); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
