package queryx

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewJSON(t *testing.T) {
	m := map[string]interface{}{"a": 1}
	j1 := NewJSON(m)
	require.Equal(t, m, j1.Val)
	require.True(t, j1.Valid)

	j2 := NewNullableJSON(nil)
	require.False(t, j2.Valid)

	j3 := NewNullableJSON(m)
	require.True(t, j3.Valid)
}

func TestJSONJSON(t *testing.T) {
	type Foo struct {
		X JSON `json:"x"`
		Y JSON `json:"y"`
	}
	x := NewJSON(map[string]interface{}{"a": "b"})
	y := NewNullableJSON(nil)
	s := `{"x":{"a":"b"},"y":null}`

	f1 := Foo{X: x, Y: y}
	b, err := json.Marshal(f1)
	require.NoError(t, err)
	require.Equal(t, s, string(b))

	var f2 Foo
	err = json.Unmarshal([]byte(s), &f2)
	require.NoError(t, err)
	require.Equal(t, x, f2.X)
	require.Equal(t, y, f2.Y)
}
