// Code generated by ent, DO NOT EDIT.

package model

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the model type in the database.
	Label = "model"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldFax holds the string denoting the fax field in the database.
	FieldFax = "fax"
	// FieldWeb holds the string denoting the web field in the database.
	FieldWeb = "web"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldRight holds the string denoting the right field in the database.
	FieldRight = "right"
	// FieldCounter holds the string denoting the counter field in the database.
	FieldCounter = "counter"
	// Table holds the table name of the model in the database.
	Table = "models"
)

// Columns holds all SQL columns for model fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldTitle,
	FieldFax,
	FieldWeb,
	FieldAge,
	FieldRight,
	FieldCounter,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Model queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByFax orders the results by the fax field.
func ByFax(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFax, opts...).ToFunc()
}

// ByWeb orders the results by the web field.
func ByWeb(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWeb, opts...).ToFunc()
}

// ByAge orders the results by the age field.
func ByAge(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAge, opts...).ToFunc()
}

// ByRight orders the results by the right field.
func ByRight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRight, opts...).ToFunc()
}

// ByCounter orders the results by the counter field.
func ByCounter(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCounter, opts...).ToFunc()
}
