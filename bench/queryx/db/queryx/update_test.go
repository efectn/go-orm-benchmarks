package queryx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewUpdate(t *testing.T) {
	s := NewUpdate().Table("users").Columns("name", "email").Values("test", "test@example.com")
	sql, args := s.ToSQL()
	require.Equal(t, "UPDATE users SET name = ?, email = ?", sql)
	require.Equal(t, []interface{}{"test", "test@example.com"}, args)
}
