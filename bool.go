package assert

import (
	"testing"
)

// True asserts that v == true.
//
//nolint:thelper // It's called below.
func True(tb testing.TB, v bool, opts ...Option) bool {
	ok := v
	if !ok {
		tb.Helper()
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
//
//nolint:thelper // It's called below.
func False(tb testing.TB, v bool, opts ...Option) bool {
	ok := !v
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"false",
			"not false",
			opts...,
		)
	}
	return ok
}
