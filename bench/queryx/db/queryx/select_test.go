package queryx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSelect(t *testing.T) {
	s := NewSelect().Select("users.*").From("users")
	sql, args := s.ToSQL()
	require.Equal(t, `SELECT users.* FROM users`, sql)
	require.Equal(t, []interface{}{}, args)

	sql, args = s.Update().Columns("name", "email").Values("test", "test@example.com").ToSQL()
	require.Equal(t, `UPDATE users SET name = ?, email = ?`, sql)
	require.Equal(t, []interface{}{"test", "test@example.com"}, args)

	s1 := s.Limit(1)
	sql, args = s1.ToSQL()
	require.Equal(t, `SELECT users.* FROM users LIMIT ?`, sql)
	require.Equal(t, []interface{}{1}, args)
}

func TestSelectWhere(t *testing.T) {
	s1 := NewSelect().Select("users.*").From("users").Where(NewClause("id = ?", []interface{}{1}))
	sql, args := s1.ToSQL()
	require.Equal(t, `SELECT users.* FROM users WHERE id = ?`, sql)
	require.Equal(t, []interface{}{1}, args)

	s1.Where(NewClause("name = ?", []interface{}{"test"}))
	sql, args = s1.ToSQL()
	require.Equal(t, `SELECT users.* FROM users WHERE (id = ?) AND (name = ?)`, sql)
	require.Equal(t, []interface{}{1, "test"}, args)
}
