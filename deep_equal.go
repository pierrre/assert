package assert

import (
	"fmt"
	"testing"

	"github.com/pierrre/compare"
)

// DeepEqualer is a function that checks if two values are deep equal.
//
// It can be customized to provide a better comparison.
//
// By default it uses [compare.DefaultComparator].
var DeepEqualer = NewDeepEqualerWithComparator(compare.DefaultComparator)

// NewDeepEqualerWithComparator creates a new [DeepEqualer] with a custom [compare.Comparator].
func NewDeepEqualerWithComparator(cr *compare.Comparator) func(v1, v2 any) (string, bool) {
	return func(v1, v2 any) (string, bool) {
		res := cr.Compare(v1, v2)
		if len(res) == 0 {
			return "", true
		}
		diff := fmt.Sprintf("%+v", res)
		return diff, false
	}
}

// DeepEqual asserts that v1 and v2 are deep equal according to [DeepEqualer].
//
//nolint:thelper // It's called below.
func DeepEqual[T any](tb testing.TB, v1, v2 T, opts ...Option) bool {
	diff, equal := DeepEqualer(v1, v2)
	ok := equal
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"deep_equal",
			fmt.Sprintf("not equal:\ndiff = %s\nv1 = %s\nv2 = %s", diff, ValueStringer(v1), ValueStringer(v2)),
			opts...,
		)
	}
	return ok
}

// NotDeepEqual asserts that v1 and v2 are not deep equal according to [DeepEqualer].
//
//nolint:thelper // It's called below.
func NotDeepEqual[T any](tb testing.TB, v1, v2 T, opts ...Option) bool {
	_, equal := DeepEqualer(v1, v2)
	ok := !equal
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"not_deep_equal",
			fmt.Sprintf("equal:\nv1 = %s\nv2 = %s", ValueStringer(v1), ValueStringer(v2)),
			opts...,
		)
	}
	return ok
}
