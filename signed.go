package assert

import (
	"fmt"
	"testing"

	"golang.org/x/exp/constraints"
)

// SignedAndFloat is a constraint that requires a type to be signed or float.
type SignedAndFloat interface {
	constraints.Signed | constraints.Float
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
