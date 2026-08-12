package helper

import (
	"database/sql"
	"fmt"
	"strings"
)

type BenchmarkControl struct {
	db *sql.DB
}

func NewBenchmarkControl() (*BenchmarkControl, error) {
	db, err := sql.Open("pgx", OrmSource)
	if err != nil {
		return nil, fmt.Errorf("benchmark control: %w", err)
	}
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("benchmark control: %w", err)
	}
	return &BenchmarkControl{db: db}, nil
}

func (control *BenchmarkControl) Close() error {
	return control.db.Close()
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

func (control *BenchmarkControl) CreateTables() error {
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
			);`,
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
			counter bigint NOT NULL,
			CONSTRAINT model5_pkey PRIMARY KEY (id)
			);`,
		},
	}

	for _, query := range queries {
		for _, line := range query {
			_, err := control.db.Exec(line)
			if err != nil {
				return fmt.Errorf("init_tables: %w", err)
			}
		}
	}

	return nil
}

func benchmarkTable(orm string) string {
	if orm == "xorm" {
		return "model5"
	}
	return "models"
}

func (control *BenchmarkControl) Reset(orm, operation string) error {
	if _, err := control.db.Exec(`TRUNCATE TABLE models, model5 RESTART IDENTITY`); err != nil {
		return fmt.Errorf("reset tables: %w", err)
	}

	if operation != "Update" {
		return nil
	}

	table := benchmarkTable(orm)
	query := fmt.Sprintf(`INSERT INTO %s (name, title, fax, web, age, "right", counter) VALUES ($1, $2, $3, $4, $5, $6, $7)`, table)
	if _, err := control.db.Exec(query, "benchmark guard", "benchmark guard", "guard", "guard", -1, false, GuardCounter); err != nil {
		return fmt.Errorf("insert update guard: %w", err)
	}
	return nil
}

func (control *BenchmarkControl) Validate(orm, operation string, iterations int) error {
	table := benchmarkTable(orm)

	if operation == "Update" {
		query := fmt.Sprintf(`SELECT COUNT(*), COUNT(*) FILTER (WHERE counter = $1), COUNT(*) FILTER (WHERE counter = $2) FROM %s`, table)
		var total, updated, guards int64
		if err := control.db.QueryRow(query, UpdatedCounter, GuardCounter).Scan(&total, &updated, &guards); err != nil {
			return fmt.Errorf("validate Update: %w", err)
		}
		if total != 2 || updated != 1 || guards != 1 {
			return fmt.Errorf("validate Update: expected 2 rows (1 updated, 1 guard), got %d rows (%d updated, %d guards)", total, updated, guards)
		}
		return nil
	}

	expected := int64(0)
	switch operation {
	case "Insert":
		expected = int64(iterations)
	case "InsertMulti":
		expected = int64(iterations) * int64(BatchSize)
	case "Read":
		expected = 1
	case "ReadSlice":
		expected = BatchSize
	default:
		return fmt.Errorf("validate: unknown operation %q", operation)
	}

	query := fmt.Sprintf(`SELECT COUNT(*), COUNT(*) FILTER (WHERE counter = $1) FROM %s`, table)
	var rows, initialized int64
	if err := control.db.QueryRow(query, InitialCounter).Scan(&rows, &initialized); err != nil {
		return fmt.Errorf("validate %s: %w", operation, err)
	}
	if rows != expected || initialized != expected {
		return fmt.Errorf("validate %s: expected %d initialized rows, got %d rows (%d initialized)", operation, expected, rows, initialized)
	}
	return nil
}
