package helper

import (
	"fmt"
	"testing"
)

// BenchExec runs an operation for the calibrated number of iterations. Setup
// must be completed before calling BenchExec.
func BenchExec(b *testing.B, orm, operation string, fn func() error) {
	b.Helper()
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		if err := fn(); err != nil {
			b.StopTimer()
			SetError(b, orm, operation, err.Error())
			return
		}
	}

	b.StopTimer()
}

// BenchRows is BenchExec for read operations. It validates the cardinality of
// the last successful iteration after stopping the timer, so the check is not
// included in the benchmark result.
func BenchRows(b *testing.B, orm, operation string, expected int, fn func() (int, error)) {
	b.Helper()
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()

	rows := 0
	for i := 0; i < b.N; i++ {
		var err error
		rows, err = fn()
		if err != nil {
			b.StopTimer()
			SetError(b, orm, operation, err.Error())
			return
		}
	}

	b.StopTimer()
	if rows != expected {
		SetError(b, orm, operation, fmt.Sprintf("expected %d materialized rows, got %d", expected, rows))
	}
}
