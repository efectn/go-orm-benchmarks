package benchs

import (
	"testing"
)

var (
	OrmMulti   int
	OrmMaxIdle int
	OrmMaxConn int
	OrmSource  string
	DebugMode  bool
)

type Instance interface {
	Name() string
	GetError(method string) string
	Init() error
	Close() error
	Insert(b *testing.B)
	InsertMulti(b *testing.B)
	Update(b *testing.B)
	Read(b *testing.B)
	ReadSlice(b *testing.B)
}

type BenchmarkResult struct {
	ORM     string
	Results []Result
}

type Result struct {
	Name     string
	ErrorMsg string

	T         int64
	MemAllocs int64
	MemBytes  int64
}

func RunBenchmarks(orm Instance) (BenchmarkResult, error) {
	err := orm.Init()
	if err != nil {
		return BenchmarkResult{}, err
	}

	defer func(orm Instance) {
		_ = orm.Close()
	}(orm)

	var result BenchmarkResult
	operations := []func(b *testing.B){orm.Insert, orm.InsertMulti, orm.Update, orm.Read, orm.ReadSlice}

	result.ORM = orm.Name()
	for _, operation := range operations {
		// Clean tables for each run
		err := CreateTables()
		if err != nil {
			panic(err)
		}

		br := testing.Benchmark(operation)
		name := getFuncName(operation)

		result.Results = append(result.Results, Result{
			Name:      name,
			ErrorMsg:  orm.GetError(name),
			T:         br.NsPerOp(),
			MemAllocs: br.AllocsPerOp(),
			MemBytes:  br.AllocedBytesPerOp(),
		})
	}

	return result, nil
}
