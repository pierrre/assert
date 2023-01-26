package assert

import (
	"fmt"
	"testing"
)

// Zero asserts that v == zero.
func Zero[T comparable](tb testing.TB, v T, opts ...Option) bool {
	tb.Helper()
	var zero T
	ok := v == zero
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("zero[%s]", TypeString[T]()),
			fmt.Sprintf("not zero:\nv = %s", ValueStringer(v)),
			opts...,
		)
	}
	return ok
}

// NotZero asserts that v != zero.
func NotZero[T comparable](tb testing.TB, v T, opts ...Option) bool {
	tb.Helper()
	var zero T
	ok := v != zero
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("not_zero[%s]", TypeString[T]()),
			fmt.Sprintf("zero:\nv = %s", ValueStringer(v)),
			opts...,
		)
	}
	return ok
}
