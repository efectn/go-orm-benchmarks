package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"time"

	"github.com/efectn/go-orm-benchmarks/benchs"

	_ "github.com/lib/pq"
)

// Version constant
const VERSION = "1.0.2"

var defaultBenchmarkNames = []string{
	"beego", "bun", "dbr", "ent",
	"godb", "gorm", "gorm_prep", "gorp",
	"pg", "pgx", "pgx_pool", "pop",
	"raw", "reform",
}

type ListOpts []string

func (opts *ListOpts) String() string {
	return fmt.Sprint(*opts)
}

func (opts *ListOpts) Set(value string) error {
	if value != "all" && !strings.Contains(" "+strings.Join(defaultBenchmarkNames, " ")+" ", " "+value+" ") {
		return fmt.Errorf("wrong run name %s", value)
	}

	*opts = append(*opts, value)
	return nil
}

func (opts ListOpts) Shuffle() {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(opts); i++ {
		a := rd.Intn(len(opts))
		b := rd.Intn(len(opts))
		opts[a], opts[b] = opts[b], opts[a]
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var orms ListOpts
	var all bool

	flag.IntVar(&benchs.OrmMaxIdle, "max_idle", 200, "max idle conns")
	flag.IntVar(&benchs.OrmMaxConn, "max_conn", 200, "max open conns")
	flag.StringVar(&benchs.OrmSource, "source", "host=localhost user=postgres password=postgres dbname=test sslmode=disable", "postgres dsn source")
	flag.IntVar(&benchs.OrmMulti, "multi", 1, "base query nums x multi")
	flag.BoolVar(&benchs.DebugMode, "debug", true, "Enable debug mode (print not-sorted results of ORMs)")
	flag.Var(&orms, "orm", "orm name: all, "+strings.Join(defaultBenchmarkNames, ", "))
	flag.Parse()

	// Check it is all
	if len(orms) == 0 {
		all = true
	} else {
		for _, n := range orms {
			if n == "all" {
				all = true
			}
		}
	}

	if all {
		orms = defaultBenchmarkNames
	}
	orms.Shuffle()

	// Run benchmarks
	benchmarks := map[string]benchs.Instance{
		"beego":     benchs.CreateBeego(200 * benchs.OrmMulti),
		"bun":       benchs.CreateBun(200 * benchs.OrmMulti),
		"dbr":       benchs.CreateDbr(200 * benchs.OrmMulti),
		"ent":       benchs.CreateEnt(200 * benchs.OrmMulti),
		"godb":      benchs.CreateGodb(200 * benchs.OrmMulti),
		"gorm":      benchs.CreateGorm(200 * benchs.OrmMulti),
		"gorm_prep": benchs.CreateGormPrep(200 * benchs.OrmMulti),
		"gorp":      benchs.CreateGorp(200 * benchs.OrmMulti),
		"pg":        benchs.CreatePg(200 * benchs.OrmMulti),
		"pgx":       benchs.CreatePgx(200 * benchs.OrmMulti),
		"pgx_pool":  benchs.CreatePgxPool(200 * benchs.OrmMulti),
		"pop":       benchs.CreatePop(200 * benchs.OrmMulti),
		"raw":       benchs.CreateRaw(200 * benchs.OrmMulti),
		"reform":    benchs.CreateReform(200 * benchs.OrmMulti),
	}

	for _, n := range orms {
		if benchs.DebugMode {
			fmt.Printf("ORM: %s\n", n)
		}

		bench := benchmarks[n]
		if bench == nil {
			panic(fmt.Sprintf("Unknown ORM: %s", n))
		}

		res, err := benchs.RunBenchmarks(bench)
		if err != nil {
			panic(fmt.Sprintf("An error occured while running the benchmarks: %v", err))
		}

		fmt.Println(res)
	}
}
