package assert

import (
	"fmt"
	"testing"
)

// PointerNil asserts that v is nil.
func PointerNil[T any](tb testing.TB, v *T, opts ...Option) bool {
	tb.Helper()
	ok := v == nil
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("pointer_nil[%s]", TypeString[T]()),
			"not nil",
			opts...,
		)
	}
	return ok
}

// PointerNotNil asserts that v is not nil.
func PointerNotNil[T any](tb testing.TB, v *T, opts ...Option) bool {
	tb.Helper()
	ok := v != nil
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("pointer_not_nil[%s]", TypeString[T]()),
			"nil",
			opts...,
		)
	}
	return ok
}
