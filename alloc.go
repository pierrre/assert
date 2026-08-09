package assert

import (
	"fmt"
	"testing"

	"github.com/pierrre/go-libs/raceutil"
)

// AllocsPerRun asserts that a function allocates a certain number of times per run.
//
// If the race detector is enabled, this function does nothing and returns true.
// This prevents tests from failing due to the increased number of allocations.
// It logs a message to [testing.TB.Log] indicating that the check was skipped.
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
	} else {
		tb.Log("assert allocs_per_run: skipped because the race detector is enabled")
	}
	return ok
}
