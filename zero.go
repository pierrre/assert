package assert

import (
	"fmt"
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
			fmt.Sprintf("zero[%s]", typeName[T]()),
			"not zero:\nv = "+ValueStringer(v),
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
			fmt.Sprintf("not_zero[%s]", typeName[T]()),
			"zero:\nv = "+ValueStringer(v),
			opts...,
		)
	}
	return ok
}
