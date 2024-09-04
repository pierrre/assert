package assert

import (
	"fmt"
	"testing"
)

// Equal asserts that v1 == v2.
//
//nolint:thelper // It's called below.
func Equal[T comparable](tb testing.TB, v1, v2 T, opts ...Option) bool {
	ok := v1 == v2
	if !ok {
		tb.Helper()
		Fail(
			tb,
			fmt.Sprintf("equal[%s]", typeName[T]()),
			fmt.Sprintf("not equal:\nv1 = %s\nv2 = %s", ValueStringer(v1), ValueStringer(v2)),
			opts...,
		)
	}
	return ok
}

// NotEqual asserts that v1 != v2.
//
//nolint:thelper // It's called below.
func NotEqual[T comparable](tb testing.TB, v1, v2 T, opts ...Option) bool {
	ok := v1 != v2
	if !ok {
		tb.Helper()
		Fail(
			tb,
			fmt.Sprintf("not_equal[%s]", typeName[T]()),
			fmt.Sprintf("equal:\nv1 = %s\nv2 = %s", ValueStringer(v1), ValueStringer(v2)),
			opts...,
		)
	}
	return ok
}
