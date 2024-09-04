package assert

import (
	"fmt"
	"testing"
)

// SignedAndFloat is a constraint that requires a type to be signed or float.
type SignedAndFloat interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

// Positive asserts that the value is positive.
//
//nolint:thelper // It's called below.
func Positive[T SignedAndFloat](tb testing.TB, v T, opts ...Option) bool {
	ok := v > 0
	if !ok {
		tb.Helper()
		Fail(
			tb,
			fmt.Sprintf("positive[%s]", typeName[T]()),
			"not positive:\nv = "+ValueStringer(v),
			opts...,
		)
	}
	return ok
}

// Negative asserts that the value is negative.
//
//nolint:thelper // It's called below.
func Negative[T SignedAndFloat](tb testing.TB, v T, opts ...Option) bool {
	ok := v < 0
	if !ok {
		tb.Helper()
		Fail(
			tb,
			fmt.Sprintf("positive[%s]", typeName[T]()),
			"not negative:\nv = "+ValueStringer(v),
			opts...,
		)
	}
	return ok
}
