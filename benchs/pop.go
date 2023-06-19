package benchs

import (
	"fmt"
	"sync"
	"testing"

	popware "github.com/gobuffalo/pop/v6"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Pop struct {
	Instance
	mu         sync.Mutex
	conn       *popware.Connection
	iterations int // Same as b.N, just to customize it
	errors     map[string]string
}

func CreatePop(iterations int) Instance {
	pop := &Pop{
		iterations: iterations,
		errors:     map[string]string{},
	}

	return pop
}

func (pop *Pop) Name() string {
	return "pop"
}

func (pop *Pop) Init() error {
	var err error
	pop.conn, err = popware.NewConnection(&popware.ConnectionDetails{
		URL: ConvertSourceToDSN(),
	})
	if err != nil {
		return err
	}

	if err = pop.conn.Open(); err != nil {
		return err
	}

	return nil
}

func (pop *Pop) Close() error {
	return pop.conn.Close()
}

func (pop *Pop) GetError(method string) string {
	return pop.errors[method]
}

func (pop *Pop) Insert(b *testing.B) {
	m := NewModel3()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := pop.conn.Create(m)
		if err != nil {
			pop.error(b, "insert", fmt.Sprintf("pop: insert: %v", err))
		}
	}
}

func (pop *Pop) InsertMulti(b *testing.B) {
	pop.error(b, "insert_multi", "pop: insert_multi: insert multi is not supported on pop")
}

func (pop *Pop) Update(b *testing.B) {
	m := NewModel3()

	err := pop.conn.Create(m)
	if err != nil {
		pop.error(b, "update", fmt.Sprintf("pop: update: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := pop.conn.Update(m)
		if err != nil {
			pop.error(b, "update", fmt.Sprintf("pop: update: %v", err))
		}
	}
}

func (pop *Pop) Read(b *testing.B) {
	m := NewModel3()

	err := pop.conn.Create(m)
	if err != nil {
		pop.error(b, "read", fmt.Sprintf("pop: read: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := pop.conn.First(m)
		if err != nil {
			pop.error(b, "read", fmt.Sprintf("pop: read: %v", err))
		}
	}
}

func (pop *Pop) ReadSlice(b *testing.B) {
	m := NewModel3()
	for i := 0; i < 100; i++ {
		err := pop.conn.Create(m)
		if err != nil {
			pop.error(b, "read_slice", fmt.Sprintf("pop: read_slice: %v", err))
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var ms []Model3
		err := pop.conn.Where("id > 0").Limit(100).All(&ms)
		if err != nil {
			pop.error(b, "read_slice", fmt.Sprintf("pop: read_slice: %v", err))
		}
	}
}

func (pop *Pop) error(b *testing.B, method string, err string) {
	b.Helper()

	pop.mu.Lock()
	pop.errors[method] = err
	pop.mu.Unlock()
	b.Fail()
}
