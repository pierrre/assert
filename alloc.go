package assert

import (
	"fmt"
	"testing"

	"github.com/pierrre/go-libs/raceutil"
)

// AllocsPerRun asserts that a function allocates a certain number of times per run.
//
// If the race detector is enabled, this function does nothing and returns true.
// It prevents the tests from failing due to the increased number of allocations.
//
//nolint:thelper // It's called below.
func AllocsPerRun(tb testing.TB, runs int, f func(), allocs float64, opts ...Option) bool {
	ok := true
	if !raceutil.Enabled {
		a := testing.AllocsPerRun(runs, f)
		ok = a == allocs
		if !ok {
			tb.Helper()
			Fail(
				tb,
				"allocs_per_run",
				fmt.Sprintf("unexpected allocs:\nexpected = %g\nactual = %g", allocs, a),
				1,
				opts...,
			)
		}
	}
	return ok
}
