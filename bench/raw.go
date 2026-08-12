package bench

import (
	"database/sql"
	"github.com/efectn/go-orm-benchmarks/helper"
	"strconv"
	"testing"
)

const (
	rawInsertBaseSQL      = `INSERT INTO models (name, title, fax, web, age, "right", counter) VALUES `
	rawInsertValuesSQL    = `($1, $2, $3, $4, $5, $6, $7)`
	rawInsertSQL          = rawInsertBaseSQL + rawInsertValuesSQL
	rawInsertReturningSQL = rawInsertSQL + ` RETURNING id`
	rawUpdateSQL          = `UPDATE models SET name = $1, title = $2, fax = $3, web = $4, age = $5, "right" = $6, counter = $7 WHERE id = $8`
	rawSelectSQL          = `SELECT id, name, title, fax, web, age, "right", counter FROM models WHERE id = $1`
	rawSelectMultiSQL     = `SELECT id, name, title, fax, web, age, "right", counter FROM models WHERE id > 0 LIMIT $1`
)

type Raw struct {
	helper.ORMInterface
	conn *sql.DB
}

func CreateRaw() helper.ORMInterface {
	return &Raw{}
}

func (raw *Raw) Name() string {
	return "raw"
}

func (raw *Raw) Init() error {
	var err error
	raw.conn, err = sql.Open("pgx", helper.OrmSource)
	if err != nil {
		return err
	}

	return nil
}

func (raw *Raw) Close() error {
	return raw.conn.Close()
}

func (raw *Raw) Insert(b *testing.B) {
	m := NewModel()
	helper.BenchExec(b, raw.Name(), "Insert", func() error {
		// pq dose not support the LastInsertId method.
		_, err := raw.conn.Exec(rawInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		return err
	})
}

func (raw *Raw) InsertMulti(b *testing.B) {
	ms := make([]*Model, 0, helper.BatchSize)
	for i := 0; i < helper.BatchSize; i++ {
		ms = append(ms, NewModel())
	}

	var valuesSQL string
	counter := 1
	for i := 0; i < helper.BatchSize; i++ {
		hoge := ""
		for j := 0; j < 7; j++ {
			if j != 6 {
				hoge += "$" + strconv.Itoa(counter) + ","
			} else {
				hoge += "$" + strconv.Itoa(counter)
			}
			counter++

		}
		if i != helper.BatchSize-1 {
			valuesSQL += "(" + hoge + "),"
		} else {
			valuesSQL += "(" + hoge + ")"
		}
	}

	helper.BenchExec(b, raw.Name(), "InsertMulti", func() error {
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
		// pq dose not support the LastInsertId method.
		_, err := raw.conn.Exec(query, args...)
		return err
	})
}

func (raw *Raw) Update(b *testing.B) {
	m := NewModel()

	err := raw.conn.QueryRow(rawInsertReturningSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter).Scan(&m.Id)
	if err != nil {
		helper.SetError(b, raw.Name(), "Update", err.Error())
		return
	}
	m.Counter = helper.UpdatedCounter

	helper.BenchExec(b, raw.Name(), "Update", func() error {
		_, err := raw.conn.Exec(rawUpdateSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter, m.Id)
		return err
	})
}

func (raw *Raw) Read(b *testing.B) {
	m := NewModel()

	err := raw.conn.QueryRow(rawInsertReturningSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter).Scan(&m.Id)
	if err != nil {
		helper.SetError(b, raw.Name(), "Read", err.Error())
		return
	}

	helper.BenchRows(b, raw.Name(), "Read", 1, func() (int, error) {
		var mout Model
		err := raw.conn.QueryRow(rawSelectSQL, m.Id).Scan(
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
			return 0, err
		}
		return 1, nil
	})
}

func (raw *Raw) ReadSlice(b *testing.B) {
	m := NewModel()
	for i := 0; i < helper.BatchSize; i++ {
		_, err := raw.conn.Exec(rawInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		if err != nil {
			helper.SetError(b, raw.Name(), "ReadSlice", err.Error())
			return
		}
	}

	helper.BenchRows(b, raw.Name(), "ReadSlice", helper.BatchSize, func() (int, error) {
		models := make([]Model, helper.BatchSize)
		rows, err := raw.conn.Query(rawSelectMultiSQL, helper.BatchSize)
		if err != nil {
			return 0, err
		}
		defer rows.Close()

		j := 0
		for ; rows.Next() && j < len(models); j++ {
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
				return j, err
			}
		}
		return j, rows.Err()
	})
}
