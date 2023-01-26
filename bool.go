package assert

import (
	"testing"
)

// True asserts that v == true.
func True(tb testing.TB, v bool, opts ...Option) bool {
	tb.Helper()
	ok := v
	if !ok {
		Fail(
			tb,
			"true",
			"not true",
			opts...,
		)
	}
	return ok
}

// False asserts that v == false.
func False(tb testing.TB, v bool, opts ...Option) bool {
	tb.Helper()
	ok := !v
	if !ok {
		Fail(
			tb,
			"false",
			"not false",
			opts...,
		)
	}
	return ok
}
