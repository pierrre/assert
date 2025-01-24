package assert

import (
	"cmp"
	"fmt"
	"testing"
)

// Greater asserts that v1 > v2.
//
//nolint:thelper // It's called below.
func Greater[T cmp.Ordered](tb testing.TB, v1, v2 T, opts ...Option) bool {
	ok := v1 > v2
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"greater",
			fmt.Sprintf("not greater than:\nv1 = %s\nv2 = %s", ValueStringer(v1), ValueStringer(v2)),
			opts...,
		)
	}
	return ok
}

// GreaterOrEqual asserts that v1 >= v2.
//
//nolint:thelper // It's called below.
func GreaterOrEqual[T cmp.Ordered](tb testing.TB, v1, v2 T, opts ...Option) bool {
	ok := v1 >= v2
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"greater_or_equal",
			fmt.Sprintf("not greater than or equal to:\nv1 = %s\nv2 = %s", ValueStringer(v1), ValueStringer(v2)),
			opts...,
		)
	}
	return ok
}

// Less asserts that v1 < v2.
//
//nolint:thelper // It's called below.
func Less[T cmp.Ordered](tb testing.TB, v1, v2 T, opts ...Option) bool {
	ok := v1 < v2
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"less",
			fmt.Sprintf("not less than:\nv1 = %s\nv2 = %s", ValueStringer(v1), ValueStringer(v2)),
			opts...,
		)
	}
	return ok
}

// LessOrEqual asserts that v1 <= v2.
//
//nolint:thelper // It's called below.
func LessOrEqual[T cmp.Ordered](tb testing.TB, v1, v2 T, opts ...Option) bool {
	ok := v1 <= v2
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"less_or_equal",
			fmt.Sprintf("not less than or equal to:\nv1 = %s\nv2 = %s", ValueStringer(v1), ValueStringer(v2)),
			opts...,
		)
	}
	return ok
}
