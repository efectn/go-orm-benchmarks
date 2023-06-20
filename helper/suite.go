package helper

import (
	"fmt"
	"testing"
	"time"
)

type ORMInterface interface {
	Name() string
	Init() error
	Close() error
	Insert(b *testing.B)
	InsertMulti(b *testing.B)
	Update(b *testing.B)
	Read(b *testing.B)
	ReadSlice(b *testing.B)
}

type BenchmarkResult struct {
	ORM     string   `json:"orm"`
	Results []Result `json:"results"`
}

type Result struct {
	Name     string `json:"name"`
	Method   string
	ErrorMsg string `json:"error_msg"`

	Time      string `json:"time"`
	NsPerOp   int64  `json:"ns_per_op"`
	MemAllocs int64  `json:"mem_allocs"`
	MemBytes  int64  `json:"mem_bytes"`

	t time.Duration
}

type BenchmarkReport []*Result

func (s BenchmarkReport) Len() int { return len(s) }

func (s BenchmarkReport) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s BenchmarkReport) Less(i, j int) bool {
	if s[i].ErrorMsg != "" {
		return false
	}
	if s[j].ErrorMsg != "" {
		return true
	}
	return s[i].t < s[j].t
}

func RunBenchmarks(orm ORMInterface, reports map[string]BenchmarkReport) (BenchmarkResult, error) {
	err := orm.Init()
	if err != nil {
		return BenchmarkResult{}, err
	}

	defer func(orm ORMInterface) {
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
		method := getFuncName(operation)

		gotResult := &Result{
			Name:      orm.Name(),
			Method:    method,
			ErrorMsg:  GetError(orm.Name(), method),
			t:         br.T,
			NsPerOp:   br.NsPerOp(),
			MemAllocs: br.AllocsPerOp(),
			MemBytes:  br.AllocedBytesPerOp(),
		}
		gotResult.Time = fmt.Sprintf("%5.3fs", float64(gotResult.t)/1e9)

		reports[method] = append(reports[method], gotResult)
		result.Results = append(result.Results, *gotResult)
	}

	return result, nil
}
