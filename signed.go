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
func Positive[T SignedAndFloat](tb testing.TB, v T, opts ...Option) bool {
	tb.Helper()
	ok := v > 0
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("positive[%s]", TypeString[T]()),
			fmt.Sprintf("not positive:\nv = %s", ValueStringer(v)),
			opts...,
		)
	}
	return ok
}

// Negative asserts that the value is negative.
func Negative[T SignedAndFloat](tb testing.TB, v T, opts ...Option) bool {
	tb.Helper()
	ok := v < 0
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("positive[%s]", TypeString[T]()),
			fmt.Sprintf("not negative:\nv = %s", ValueStringer(v)),
			opts...,
		)
	}
	return ok
}
