package assert

import (
	"fmt"
	"testing"
)

// AllocsPerRun asserts that a function allocates a certain number of times per run.
func AllocsPerRun(tb testing.TB, runs int, f func(), allocs float64, opts ...Option) bool {
	tb.Helper()
	a := testing.AllocsPerRun(runs, f)
	ok := a == allocs
	if !ok {
		Fail(
			tb,
			"allocs_per_run",
			fmt.Sprintf("unexpected allocs:\nexpected = %f\nactual = %f", allocs, a),
			opts...,
		)
	}
	return ok
}
