package benchs

import (
	"database/sql"
	"fmt"
	"strconv"
	"sync"
	"testing"
)

const (
	rawInsertBaseSQL   = `INSERT INTO models (name, title, fax, web, age, "right", counter) VALUES `
	rawInsertValuesSQL = `($1, $2, $3, $4, $5, $6, $7)`
	rawInsertSQL       = rawInsertBaseSQL + rawInsertValuesSQL
	rawUpdateSQL       = `UPDATE models SET name = $1, title = $2, fax = $3, web = $4, age = $5, "right" = $6, counter = $7 WHERE id = $8`
	rawSelectSQL       = `SELECT id, name, title, fax, web, age, "right", counter FROM models WHERE id = $1`
	rawSelectMultiSQL  = `SELECT id, name, title, fax, web, age, "right", counter FROM models WHERE id > 0 LIMIT 100`
)

type Raw struct {
	Instance
	mu         sync.Mutex
	conn       *sql.DB
	iterations int // Same as b.N, just to customize it
	errors     map[string]string
}

func CreateRaw(iterations int) Instance {
	raw := &Raw{
		iterations: iterations,
		errors:     map[string]string{},
	}

	return raw
}

func (raw *Raw) Name() string {
	return "raw"
}

func (raw *Raw) Init() error {
	var err error
	raw.conn, err = sql.Open("pgx", OrmSource)
	if err != nil {
		return err
	}

	return nil
}

func (raw *Raw) Close() error {
	return raw.conn.Close()
}

func (raw *Raw) GetError(method string) string {
	return raw.errors[method]
}

func (raw *Raw) Insert(b *testing.B) {
	m := NewModel()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// pq dose not support the LastInsertId method.
		_, err := raw.conn.Exec(rawInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		if err != nil {
			raw.error(b, "insert", fmt.Sprintf("raw: insert: %v", err))
		}
	}
}

func (raw *Raw) InsertMulti(b *testing.B) {
	ms := make([]*Model, 0, 100)
	for i := 0; i < 100; i++ {
		ms = append(ms, NewModel())
	}

	b.ReportAllocs()
	b.ResetTimer()

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
		// pq dose not support the LastInsertId method.
		_, err := raw.conn.Exec(query, args...)
		if err != nil {
			raw.error(b, "insert_multi", fmt.Sprintf("raw: insert_multi: %v", err))
		}
	}
}

func (raw *Raw) Update(b *testing.B) {
	m := NewModel()

	_, err := raw.conn.Exec(rawInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
	if err != nil {
		raw.error(b, "update", fmt.Sprintf("raw: update: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := raw.conn.Exec(rawUpdateSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter, m.Id)
		if err != nil {
			raw.error(b, "update", fmt.Sprintf("raw: update: %v", err))
		}
	}
}

func (raw *Raw) Read(b *testing.B) {
	m := NewModel()

	_, err := raw.conn.Exec(rawInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
	if err != nil {
		raw.error(b, "read", fmt.Sprintf("raw: read: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var mout Model
		err := raw.conn.QueryRow(rawSelectSQL, 1).Scan(
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
			raw.error(b, "read", fmt.Sprintf("raw: read: %v", err))
		}
	}
}

func (raw *Raw) ReadSlice(b *testing.B) {
	m := NewModel()
	for i := 0; i < 100; i++ {
		_, err := raw.conn.Exec(rawInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		if err != nil {
			raw.error(b, "read_slice", fmt.Sprintf("raw: read_slice: %v", err))
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var j int
		models := make([]Model, 100)
		rows, err := raw.conn.Query(rawSelectMultiSQL)
		if err != nil {
			raw.error(b, "read_slice", fmt.Sprintf("raw: read_slice: %v", err))
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
				raw.error(b, "read_slice", fmt.Sprintf("raw: read_slice: %v", err))
			}
		}
		err = rows.Err()
		if err != nil {
			raw.error(b, "read_slice", fmt.Sprintf("raw: read_slice: %v", err))
		}
		err = rows.Close()
		if err != nil {
			raw.error(b, "read_slice", fmt.Sprintf("raw: read_slice: %v", err))
		}
	}
}

func (raw *Raw) error(b *testing.B, method string, err string) {
	b.Helper()

	raw.mu.Lock()
	raw.errors[method] = err
	raw.mu.Unlock()
	b.Fail()
}
