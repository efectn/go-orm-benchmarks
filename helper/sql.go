package helper

import (
	"database/sql"
	"fmt"
	"strings"
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
