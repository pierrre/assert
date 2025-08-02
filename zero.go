package assert

import (
	"testing"
)

// Zero asserts that v == zero.
//
//nolint:thelper // It's called below.
func Zero[T comparable](tb testing.TB, v T, opts ...Option) bool {
	var zero T
	ok := v == zero
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"zero",
			"not zero:\nv = "+ValueStringer(v),
			1,
			opts...,
		)
	}
	return ok
}

// NotZero asserts that v != zero.
//
//nolint:thelper // It's called below.
func NotZero[T comparable](tb testing.TB, v T, opts ...Option) bool {
	var zero T
	ok := v != zero
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"not_zero",
			"zero:\nv = "+ValueStringer(v),
			1,
			opts...,
		)
	}
	return ok
}
