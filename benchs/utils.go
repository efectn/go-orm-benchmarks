package benchs

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	OrmMulti   int
	OrmMaxIdle int
	OrmMaxConn int
	OrmSource  string
	DebugMode  bool
)

// Convert ORMSource to DSN (dburl)
func ConvertSourceToDSN() string {
	template := "postgres://$(user):$(password)@$(host):5432/$(dbname)"

	// Parse one-by-one instead of using REGEX because of performance issues
	for _, option := range strings.Split(OrmSource, " ") {
		k := strings.Split(option, "=")[0]
		v := strings.Split(option, "=")[1]

		if strings.Contains(template, "$("+k+")") {
			template = strings.ReplaceAll(template, "$("+k+")", v)
		} else {
			template += "?" + option
		}
	}

	return template
}

func SplitSource() map[string]string {
	options := make(map[string]string)
	// Split one-by-one instead of using REGEX because of performance issues
	for _, option := range strings.Split(OrmSource, " ") {
		k := strings.Split(option, "=")[0]
		v := strings.Split(option, "=")[1]

		options[k] = v
	}

	return options
}

func CheckErr(err error, b ...*B) {
	if err != nil {
		log.Fatalf("[go-orm-benchmarks] ERR: %v", err)

		if len(b) > 0 {
			b[0].FailNow()
		}
	}
}

func WrapExecute(b *B, cbk func()) {
	b.StopTimer()
	defer b.StartTimer()
	cbk()
	b.ResetTimer()
}

func InitDB() {
	sqls := [][]string{
		{
			`DROP TABLE IF EXISTS models;`,
			`CREATE TABLE models (
			id SERIAL NOT NULL,
			name text NOT NULL,
			title text NOT NULL,
			fax text NOT NULL,
			web text NOT NULL,
			age integer NOT NULL,
			"right" boolean NOT NULL,
			counter bigint NOT NULL,
			CONSTRAINT models_pkey PRIMARY KEY (id)
			) WITH (OIDS=FALSE);`,
		},
		{
			`DROP TABLE IF EXISTS model5;`,
			`CREATE TABLE model5 (
			id SERIAL NOT NULL,
			name text NOT NULL,
			title text NOT NULL,
			fax text NOT NULL,
			web text NOT NULL,
			age integer NOT NULL,
			"right" boolean NOT NULL,
			counter bigint NOT NULL
			) WITH (OIDS=FALSE);`,
		},
	}

	DB, err := sql.Open("pgx", OrmSource)
	CheckErr(err)
	defer func() {
		err := DB.Close()
		CheckErr(err)
	}()

	err = DB.Ping()
	CheckErr(err)

	for _, sql := range sqls {
		for _, line := range sql {
			_, err = DB.Exec(line)
			CheckErr(err)
		}
	}

}
