package assert

import (
	"fmt"
	"testing"
)

// AllocsPerRun asserts that a function allocates a certain number of times per run.
//
//nolint:thelper // aaaa
func AllocsPerRun(tb testing.TB, runs int, f func(), allocs float64, opts ...Option) bool {
	a := testing.AllocsPerRun(runs, f)
	ok := a == allocs
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"allocs_per_run",
			fmt.Sprintf("unexpected allocs:\nexpected = %g\nactual = %g", allocs, a),
			opts...,
		)
	}
	return ok
}
