package queryx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewInsert(t *testing.T) {
	s1 := NewInsert().Into("users")
	sql, args := s1.ToSQL()
	require.Equal(t, "INSERT INTO users DEFAULT VALUES", sql)
	require.Equal(t, []interface{}{}, args)

	var columns []string
	var values []interface{}
	s2 := NewInsert().Into("users").Columns(columns...).Values(values...).Returning("id", "name")
	sql, args = s2.ToSQL()
	require.Equal(t, "INSERT INTO users DEFAULT VALUES RETURNING id, name", sql)
	require.Equal(t, []interface{}{}, args)

	s3 := NewInsert().Into("users").Columns("name", "email").Values("test", "test@example.com")
	sql, args = s3.ToSQL()
	require.Equal(t, "INSERT INTO users(name, email) VALUES (?, ?)", sql)
	require.Equal(t, []interface{}{"test", "test@example.com"}, args)

	s4 := NewInsert().Into("users").Columns("name", "email").
		Values("test1", "test1@example.com").
		Values("test2", "test2@example.com")
	sql, args = s4.ToSQL()
	require.Equal(t, "INSERT INTO users(name, email) VALUES (?, ?), (?, ?)", sql)
	require.Equal(t, []interface{}{"test1", "test1@example.com", "test2", "test2@example.com"}, args)
}
