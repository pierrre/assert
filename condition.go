package assert

import (
	"testing"
)

// Condition asserts that f returns true.
//
//nolint:thelper // It's called below.
func Condition(tb testing.TB, f func() bool, opts ...Option) bool {
	ok := f()
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"condition",
			"failed",
			1,
			opts...,
		)
	}
	return ok
}
