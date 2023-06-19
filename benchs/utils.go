package benchs

import (
	"database/sql"
	"fmt"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// from https://stackoverflow.com/questions/56616196/how-to-convert-camel-case-string-to-snake-case, needs cleaner solution
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func getFuncName(function any) string {
	name := strings.Split(runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name(), ".")
	straightName := strings.Split(name[len(name)-1], "-")[0]

	return ToSnakeCase(straightName)
}

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

func CreateTables() error {
	queries := [][]string{
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

	db, err := sql.Open("pgx", OrmSource)
	if err != nil {
		return fmt.Errorf("init_tables: %w", err)
	}

	defer func() {
		_ = db.Close()
	}()

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("init_tables: %w", err)
	}

	for _, query := range queries {
		for _, line := range query {
			_, err = db.Exec(line)
			if err != nil {
				return fmt.Errorf("init_tables: %w", err)
			}
		}
	}

	return nil
}
