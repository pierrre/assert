package assert

import (
	"fmt"
	"testing"

	"github.com/pierrre/go-libs/reflectutil"
)

// Zero asserts that v == zero.
func Zero[T comparable](tb testing.TB, v T, opts ...Option) bool {
	tb.Helper()
	var zero T
	ok := v == zero
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("zero[%s]", reflectutil.TypeFullNameFor[T]()),
			"not zero:\nv = "+ValueStringer(v),
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
			fmt.Sprintf("not_zero[%s]", reflectutil.TypeFullNameFor[T]()),
			"zero:\nv = "+ValueStringer(v),
			opts...,
		)
	}
	return ok
}
