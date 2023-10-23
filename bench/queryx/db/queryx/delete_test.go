package queryx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDelete(t *testing.T) {
	s := NewDelete().From("users")

	sql, args := s.ToSQL()
	require.Equal(t, "DELETE FROM users", sql)
	require.Equal(t, []interface{}{}, args)
}
