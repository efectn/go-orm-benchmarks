package benchs

import (
	"fmt"
	"sync"
	"testing"
	xormdb "xorm.io/xorm"
)

type Xorm struct {
	Instance
	mu         sync.Mutex
	conn       *xormdb.Session
	iterations int // Same as b.N, just to customize it
	errors     map[string]string
}

func CreateXorm(iterations int) Instance {
	xorm := &Xorm{
		iterations: iterations,
		errors:     map[string]string{},
	}

	return xorm
}

func (xorm *Xorm) Name() string {
	return "xorm"
}

func (xorm *Xorm) Init() error {
	engine, err := xormdb.NewEngine("postgres", OrmSource)
	if err != nil {
		return err
	}

	xorm.conn = engine.NewSession()

	return nil
}

func (xorm *Xorm) Close() error {
	return xorm.conn.Close()
}

func (xorm *Xorm) GetError(method string) string {
	return xorm.errors[method]
}

func (xorm *Xorm) Insert(b *testing.B) {
	m := NewModel5()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.ID = 0
		_, err := xorm.conn.Insert(m)
		if err != nil {
			xorm.error(b, "insert", fmt.Sprintf("xorm: insert: %v", err))
		}
	}
}

func (xorm *Xorm) InsertMulti(b *testing.B) {
	ms := make([]*Model5, 0, 100)
	for i := 0; i < 100; i++ {
		ms = append(ms, NewModel5())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := xorm.conn.Insert(&ms)
		if err != nil {
			xorm.error(b, "insert_multi", fmt.Sprintf("xorm: insert_multi: %v", err))
		}
	}
}

func (xorm *Xorm) Update(b *testing.B) {
	m := NewModel5()

	_, err := xorm.conn.Insert(m)
	if err != nil {
		xorm.error(b, "update", fmt.Sprintf("xorm: update: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := xorm.conn.ID(m.ID).Update(m)
		if err != nil {
			xorm.error(b, "update", fmt.Sprintf("xorm: update: %v", err))
		}
	}
}

func (xorm *Xorm) Read(b *testing.B) {
	m := NewModel5()

	_, err := xorm.conn.Insert(m)
	if err != nil {
		xorm.error(b, "read", fmt.Sprintf("xorm: read: %v", err))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := xorm.conn.ID(1).NoAutoCondition().Get(m)
		if err != nil {
			xorm.error(b, "read", fmt.Sprintf("xorm: read: %v", err))
		}
	}
}

func (xorm *Xorm) ReadSlice(b *testing.B) {
	m := NewModel5()

	for i := 0; i < 100; i++ {
		m.ID = 0
		_, err := xorm.conn.Insert(m)
		if err != nil {
			xorm.error(b, "read_slice", fmt.Sprintf("xorm: read_slice: %v", err))
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var models []*Model5
		err := xorm.conn.Where("id > 0").Limit(100).Find(&models)
		if err != nil {
			xorm.error(b, "read_slice", fmt.Sprintf("xorm: read_slice: %v", err))
		}
	}
}

func (xorm *Xorm) error(b *testing.B, method string, err string) {
	b.Helper()

	xorm.mu.Lock()
	xorm.errors[method] = err
	xorm.mu.Unlock()
	b.Fail()
}
