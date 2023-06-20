package bench

import (
	"context"
	"gitee.com/chunanyong/zorm"
	r "github.com/efectn/go-orm-benchmarks/bench/reform"
	models "github.com/efectn/go-orm-benchmarks/bench/sqlboiler"
)

var ctx = context.Background()

// Model for GORM, GORP, Beego, Bun, Pg, Raw, Sqlc, Ent
type Model struct {
	Id      int `orm:"auto" gorm:"primary_key" db:"id" bun:",pk,autoincrement"`
	Name    string
	Title   string
	Fax     string
	Web     string
	Age     int
	Right   bool
	Counter int64
}

func (m *Model) TableName() string {
	return "models"
}

func (m *Model) Table() string {
	return "models"
}

func NewModel() *Model {
	m := new(Model)
	m.Name = "Orm Benchmark"
	m.Title = "Just a Benchmark for fun"
	m.Fax = "99909990"
	m.Web = "http://blog.milkpod29.me"
	m.Age = 100
	m.Right = true
	m.Counter = 1000

	return m
}

func NewModelAlt() Model {
	m := Model{}
	m.Name = "Orm Benchmark"
	m.Title = "Just a Benchmark for fun"
	m.Fax = "99909990"
	m.Web = "http://blog.milkpod29.me"
	m.Age = 100
	m.Right = true
	m.Counter = 1000

	return m
}

// Model for Godb, Dbr
type Model2 struct {
	ID      int    `db:"id,key,auto"`
	Name    string `db:"name"`
	Title   string `db:"title"`
	Fax     string `db:"fax"`
	Web     string `db:"web"`
	Age     int    `db:"age"`
	Right   bool   `db:"right"`
	Counter int64  `db:"counter"`
}

func (*Model2) TableName() string {
	return "models"
}

func NewModel2() *Model2 {
	m := new(Model2)
	m.Name = "Orm Benchmark"
	m.Title = "Just a Benchmark for fun"
	m.Fax = "99909990"
	m.Web = "http://blog.milkpod29.me"
	m.Age = 100
	m.Right = true
	m.Counter = 1000

	return m
}

// Model for Pop, Rel
type Model3 struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Title   string `db:"title"`
	Fax     string `db:"fax"`
	Web     string `db:"web"`
	Age     int    `db:"age"`
	Right   bool   `db:"right"`
	Counter int64  `db:"counter"`
}

func (Model3) Table() string {
	return "models"
}

func (Model3) TableName() string {
	return "models"
}

func NewModel3() *Model3 {
	m := new(Model3)
	m.Name = "Orm Benchmark"
	m.Title = "Just a Benchmark for fun"
	m.Fax = "99909990"
	m.Web = "http://blog.milkpod29.me"
	m.Age = 100
	m.Right = true
	m.Counter = 1000

	return m
}

// Model for Upper
type Model4 struct {
	ID      int    `db:"id,omitempty"`
	Name    string `db:"name"`
	Title   string `db:"title"`
	Fax     string `db:"fax"`
	Web     string `db:"web"`
	Age     int    `db:"age"`
	Right   bool   `db:"right"`
	Counter int64  `db:"counter"`
}

func NewModel4() *Model4 {
	m := new(Model4)
	m.Name = "Orm Benchmark"
	m.Title = "Just a Benchmark for fun"
	m.Fax = "99909990"
	m.Web = "http://blog.milkpod29.me"
	m.Age = 100
	m.Right = true
	m.Counter = 1000

	return m
}

// Model for XORM
type Model5 struct {
	ID      int `xorm:"pk autoincr 'id'"`
	Name    string
	Title   string
	Fax     string
	Web     string
	Age     int
	Right   bool
	Counter int64
}

func NewModel5() *Model5 {
	m := new(Model5)
	m.Name = "Orm Benchmark"
	m.Title = "Just a Benchmark for fun"
	m.Fax = "99909990"
	m.Web = "http://blog.milkpod29.me"
	m.Age = 100
	m.Right = true
	m.Counter = 1000

	return m
}

// Model for Sqlboiler
func NewModel6() *models.Model {
	m := new(models.Model)
	m.Name = "Orm Benchmark"
	m.Title = "Just a Benchmark for fun"
	m.Fax = "99909990"
	m.Web = "http://blog.milkpod29.me"
	m.Age = 100
	m.Right = true
	m.Counter = 1000

	return m
}

// Model for zorm
type Model7 struct {
	zorm.EntityStruct
	ID      int    `column:"id"`
	Name    string `column:"name"`
	Title   string `column:"title"`
	Fax     string `column:"fax"`
	Web     string `column:"web"`
	Age     int    `column:"age"`
	Right   bool   `column:"\"right\""`
	Counter int64  `column:"counter"`
}

func (entity *Model7) GetTableName() string {
	return "models"
}

func (entity *Model7) GetPKColumnName() string {
	return "id"
}

func NewModel7() *Model7 {
	m := new(Model7)
	m.Name = "Orm Benchmark"
	m.Title = "Just a Benchmark for fun"
	m.Fax = "99909990"
	m.Web = "http://blog.milkpod29.me"
	m.Age = 100
	m.Right = true
	m.Counter = 1000

	return m
}

func NewReformModel() *r.ReformModels {
	m := new(r.ReformModels)
	m.Name = "Orm Benchmark"
	m.Title = "Just a Benchmark for fun"
	m.Fax = "99909990"
	m.Web = "http://blog.milkpod29.me"
	m.Age = 100
	m.Right = true
	m.Counter = 1000

	return m
}
