package assert

import (
	"fmt"
	"testing"

	"golang.org/x/exp/constraints"
)

// Greater asserts that v1 > v2.
func Greater[T constraints.Ordered](tb testing.TB, v1, v2 T, opts ...Option) bool {
	tb.Helper()
	ok := v1 > v2
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("greater[%s]", TypeString[T]()),
			fmt.Sprintf("not greater than:\nv1 = %s\nv2 = %s", ValueStringer(v1), ValueStringer(v2)),
			opts...,
		)
	}
	return ok
}

// GreaterOrEqual asserts that v1 >= v2.
func GreaterOrEqual[T constraints.Ordered](tb testing.TB, v1, v2 T, opts ...Option) bool {
	tb.Helper()
	ok := v1 >= v2
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("greater_or_equal[%s]", TypeString[T]()),
			fmt.Sprintf("not greater than or equal to:\nv1 = %s\nv2 = %s", ValueStringer(v1), ValueStringer(v2)),
			opts...,
		)
	}
	return ok
}

// Less asserts that v1 < v2.
func Less[T constraints.Ordered](tb testing.TB, v1, v2 T, opts ...Option) bool {
	tb.Helper()
	ok := v1 < v2
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("less[%s]", TypeString[T]()),
			fmt.Sprintf("not less than:\nv1 = %s\nv2 = %s", ValueStringer(v1), ValueStringer(v2)),
			opts...,
		)
	}
	return ok
}

// LessOrEqual asserts that v1 <= v2.
func LessOrEqual[T constraints.Ordered](tb testing.TB, v1, v2 T, opts ...Option) bool {
	tb.Helper()
	ok := v1 <= v2
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("less_or_equal[%s]", TypeString[T]()),
			fmt.Sprintf("not less than or equal to:\nv1 = %s\nv2 = %s", ValueStringer(v1), ValueStringer(v2)),
			opts...,
		)
	}
	return ok
}
