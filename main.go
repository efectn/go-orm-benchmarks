package main

import (
	"flag"
	"fmt"
	"github.com/efectn/go-orm-benchmarks/helper"
	"math/rand"
	"os"
	"runtime"
	"sort"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/efectn/go-orm-benchmarks/bench"

	_ "github.com/lib/pq"
)

// Version constant
const VERSION = "1.0.2"

var defaultBenchmarkNames = []string{
	"beego", "bun", "dbr", "ent",
	"godb", "gorm", "gorm_prep", "gorp",
	"pg", "pgx", "pgx_pool", "pop",
	"raw", "reform", "rel", "sqlboiler",
	"sqlc", "sqlx", "upper", "xorm",
	"zorm",
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

	flag.IntVar(&helper.OrmMaxIdle, "max_idle", 200, "max idle conns")
	flag.IntVar(&helper.OrmMaxConn, "max_conn", 200, "max open conns")
	flag.StringVar(&helper.OrmSource, "source", "host=localhost user=postgres password=postgres dbname=test sslmode=disable", "postgres dsn source")
	flag.IntVar(&helper.OrmMulti, "multi", 1, "base query nums x multi")
	flag.BoolVar(&helper.DebugMode, "debug", true, "Enable debug mode (also prints not-sorted results of ORMs)")
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

	// Init error map
	helper.Errors = make(map[string]map[string]string, 0)
	for _, name := range defaultBenchmarkNames {
		helper.Errors[name] = make(map[string]string, 0)
	}

	runBenchmarks(orms)
}

func runBenchmarks(orms ListOpts) {
	// Run benchmarks
	benchmarks := map[string]helper.ORMInterface{
		"beego":     bench.CreateBeego(200 * helper.OrmMulti),
		"bun":       bench.CreateBun(200 * helper.OrmMulti),
		"dbr":       bench.CreateDbr(200 * helper.OrmMulti),
		"ent":       bench.CreateEnt(200 * helper.OrmMulti),
		"godb":      bench.CreateGodb(200 * helper.OrmMulti),
		"gorm":      bench.CreateGorm(200 * helper.OrmMulti),
		"gorm_prep": bench.CreateGormPrep(200 * helper.OrmMulti),
		"gorp":      bench.CreateGorp(200 * helper.OrmMulti),
		"pg":        bench.CreatePg(200 * helper.OrmMulti),
		"pgx":       bench.CreatePgx(200 * helper.OrmMulti),
		"pgx_pool":  bench.CreatePgxPool(200 * helper.OrmMulti),
		"pop":       bench.CreatePop(200 * helper.OrmMulti),
		"raw":       bench.CreateRaw(200 * helper.OrmMulti),
		"reform":    bench.CreateReform(200 * helper.OrmMulti),
		"rel":       bench.CreateRel(200 * helper.OrmMulti),
		"sqlboiler": bench.CreateSqlboiler(200 * helper.OrmMulti),
		"sqlc":      bench.CreateSqlc(200 * helper.OrmMulti),
		"sqlx":      bench.CreateSqlx(200 * helper.OrmMulti),
		"upper":     bench.CreateUpper(200 * helper.OrmMulti),
		"xorm":      bench.CreateXorm(200 * helper.OrmMulti),
		"zorm":      bench.CreateZorm(200 * helper.OrmMulti),
	}

	debugTable := new(tabwriter.Writer)
	debugTable.Init(os.Stdout, 0, 8, 2, '\t', tabwriter.AlignRight)

	reports := make(map[string]helper.BenchmarkReport, 0)
	for _, n := range orms {
		orm := benchmarks[n]
		if orm == nil {
			panic(fmt.Sprintf("Unknown ORM: %s", n))
		}

		res, err := helper.RunBenchmarks(orm, reports)
		if err != nil {
			panic(fmt.Sprintf("An error occured while running the benchmarks: %v", err))
		}

		if helper.DebugMode {
			fmt.Printf("\n%s ORM Benchmark Results:\n", n)
			for _, result := range res.Results {
				_, _ = fmt.Fprintf(debugTable, "%s:\t%s\t%d ns/op\t%d B/op\t%d allocs/op\n", result.Method, result.Time, result.NsPerOp, result.MemBytes, result.MemAllocs)
			}
			_ = debugTable.Flush()
		}
	}

	// Sort results
	for _, v := range reports {
		sort.Sort(v)
	}

	// Print final reports
	reportTable := new(tabwriter.Writer)
	reportTable.Init(os.Stdout, 0, 8, 2, '\t', tabwriter.AlignRight)

	_, _ = fmt.Fprintf(reportTable, "Reports:\n\n")
	for method, report := range reports {
		_, _ = fmt.Fprintf(reportTable, "%s - %d Times\n", method, 200*helper.OrmMulti)
		for _, result := range report {
			if result.ErrorMsg == "" {
				_, _ = fmt.Fprintf(reportTable, "%s:\t%s\t%d ns/op\t%d B/op\t%d allocs/op\n", result.Name, result.Time, result.NsPerOp, result.MemBytes, result.MemAllocs)
			} else {
				_, _ = fmt.Fprintf(reportTable, "%s:\t%s\n", result.Name, result.ErrorMsg)
			}
		}
		_, _ = fmt.Fprintf(reportTable, "\n")
	}
	_ = reportTable.Flush()
}
