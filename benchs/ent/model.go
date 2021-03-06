// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/efectn/go-orm-benchmarks/benchs/ent/model"
)

// Model is the model entity for the Model schema.
type Model struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Fax holds the value of the "fax" field.
	Fax string `json:"fax,omitempty"`
	// Web holds the value of the "web" field.
	Web string `json:"web,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Right holds the value of the "right" field.
	Right bool `json:"right,omitempty"`
	// Counter holds the value of the "counter" field.
	Counter int64 `json:"counter,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Model) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case model.FieldRight:
			values[i] = new(sql.NullBool)
		case model.FieldID, model.FieldAge, model.FieldCounter:
			values[i] = new(sql.NullInt64)
		case model.FieldName, model.FieldTitle, model.FieldFax, model.FieldWeb:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Model", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Model fields.
func (m *Model) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case model.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case model.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case model.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				m.Title = value.String
			}
		case model.FieldFax:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field fax", values[i])
			} else if value.Valid {
				m.Fax = value.String
			}
		case model.FieldWeb:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field web", values[i])
			} else if value.Valid {
				m.Web = value.String
			}
		case model.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				m.Age = int(value.Int64)
			}
		case model.FieldRight:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field right", values[i])
			} else if value.Valid {
				m.Right = value.Bool
			}
		case model.FieldCounter:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field counter", values[i])
			} else if value.Valid {
				m.Counter = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Model.
// Note that you need to call Model.Unwrap() before calling this method if this Model
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Model) Update() *ModelUpdateOne {
	return (&ModelClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Model entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Model) Unwrap() *Model {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Model is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Model) String() string {
	var builder strings.Builder
	builder.WriteString("Model(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", name=")
	builder.WriteString(m.Name)
	builder.WriteString(", title=")
	builder.WriteString(m.Title)
	builder.WriteString(", fax=")
	builder.WriteString(m.Fax)
	builder.WriteString(", web=")
	builder.WriteString(m.Web)
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", m.Age))
	builder.WriteString(", right=")
	builder.WriteString(fmt.Sprintf("%v", m.Right))
	builder.WriteString(", counter=")
	builder.WriteString(fmt.Sprintf("%v", m.Counter))
	builder.WriteByte(')')
	return builder.String()
}

// Models is a parsable slice of Model.
type Models []*Model

func (m Models) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
