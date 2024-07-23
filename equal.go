package assert

import (
	"fmt"
	"testing"

	"github.com/pierrre/go-libs/reflectutil"
)

// Equal asserts that v1 == v2.
func Equal[T comparable](tb testing.TB, v1, v2 T, opts ...Option) bool {
	tb.Helper()
	ok := v1 == v2
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("equal[%s]", reflectutil.TypeFullNameFor[T]()),
			fmt.Sprintf("not equal:\nv1 = %s\nv2 = %s", ValueStringer(v1), ValueStringer(v2)),
			opts...,
		)
	}
	return ok
}

// NotEqual asserts that v1 != v2.
func NotEqual[T comparable](tb testing.TB, v1, v2 T, opts ...Option) bool {
	tb.Helper()
	ok := v1 != v2
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("not_equal[%s]", reflectutil.TypeFullNameFor[T]()),
			fmt.Sprintf("equal:\nv1 = %s\nv2 = %s", ValueStringer(v1), ValueStringer(v2)),
			opts...,
		)
	}
	return ok
}
