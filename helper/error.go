package helper

import (
	"sync"
	"testing"
)

var Errors map[string]map[string]string

var mu sync.Mutex

func SetError(b *testing.B, orm, method, err string) {
	b.Helper()

	mu.Lock()
	Errors[orm][method] = err
	mu.Unlock()
	b.Fail()
}

func GetError(orm, method string) string {
	return Errors[orm][method]
}
