package assert

import (
	"testing"
)

// SignedNumber represents all signed numeric types (signed integers and floats).
type SignedNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

// Positive asserts that the value is positive.
//
//nolint:thelper // It's called below.
func Positive[T SignedNumber](tb testing.TB, v T, opts ...Option) bool {
	ok := v > 0
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"positive",
			"not positive:\nv = "+ValueStringer.Load()(v),
			1,
			opts...,
		)
	}
	return ok
}

// Negative asserts that the value is negative.
//
//nolint:thelper // It's called below.
func Negative[T SignedNumber](tb testing.TB, v T, opts ...Option) bool {
	ok := v < 0
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"negative",
			"not negative:\nv = "+ValueStringer.Load()(v),
			1,
			opts...,
		)
	}
	return ok
}
