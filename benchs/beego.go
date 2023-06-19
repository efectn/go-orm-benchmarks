package benchs

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"sync"
	"testing"
)

type Beego struct {
	Instance
	mu         sync.Mutex
	conn       orm.Ormer
	iterations int // Same as b.N, just to customize it
	errors     map[string]string
}

func CreateBeego(iterations int) Instance {
	beego := &Beego{
		iterations: iterations,
		errors:     map[string]string{},
	}

	return beego
}

func (beego *Beego) Name() string {
	return "beego"
}

func (beego *Beego) Init() error {
	err := orm.RegisterDataBase("default", "postgres", OrmSource, OrmMaxIdle, OrmMaxConn)
	if err != nil {
		return err
	}

	orm.RegisterModel(new(Model))
	beego.conn = orm.NewOrm()

	return nil
}

func (beego *Beego) Close() {
	//
}

func (beego *Beego) GetError(method string) string {
	return beego.errors[method]
}

func (beego *Beego) Insert(b *testing.B) {
	m := NewModel()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Id = 0

		_, err := beego.conn.Insert(m)
		if err != nil {
			beego.error(b, "insert", fmt.Sprintf("beego: insert: %v", err))
		}
	}
}

func (beego *Beego) InsertMulti(b *testing.B) {
	ms := make([]*Model, 0, 100)
	for i := 0; i < 100; i++ {
		ms = append(ms, NewModel())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := beego.conn.InsertMulti(100, ms)
		if err != nil {
			beego.error(b, "insert_multi", fmt.Sprintf("beego: insert_multi: %v", err))
		}
	}
}

func (beego *Beego) Update(b *testing.B) {
	m := NewModel()

	_, err := beego.conn.Insert(m)
	if err != nil {
		b.Errorf("beego: update: %v", err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := beego.conn.Update(m)
		if err != nil {
			b.Errorf("beego: update: %v", err)
		}
	}
}

func (beego *Beego) Read(b *testing.B) {
	m := NewModel()

	_, err := beego.conn.Insert(m)
	if err != nil {
		b.Errorf("beego: read: %v", err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := beego.conn.Read(m)
		if err != nil {
			b.Errorf("beego: read: %v", err)
		}
	}
}

func (beego *Beego) ReadSlice(b *testing.B) {
	m := NewModel()
	for i := 0; i < 100; i++ {
		m.Id = 0
		_, err := beego.conn.Insert(m)
		if err != nil {
			b.Errorf("beego: read_slice: %v", err)
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var models []*Model
		_, err := beego.conn.QueryTable("models").Filter("id__gt", 0).Limit(100).All(&models)
		if err != nil {
			b.Errorf("beego: read_slice: %v", err)
		}
	}
}

func (beego *Beego) error(b *testing.B, method string, err string) {
	b.Helper()

	beego.mu.Lock()
	beego.errors[method] = err
	beego.mu.Unlock()
	b.Fail()
}
