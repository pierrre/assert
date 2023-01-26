package assert

import (
	"testing"
)

// Condition asserts that f returns true.
func Condition(tb testing.TB, f func() bool, opts ...Option) bool {
	tb.Helper()
	ok := f()
	if !ok {
		Fail(
			tb,
			"condition",
			"failed",
			opts...,
		)
	}
	return ok
}
