package assert

import (
	"fmt"
	"testing"

	"github.com/pierrre/compare"
	"github.com/pierrre/go-libs/syncutil/atomicutil"
)

// DeepEqualer is a function that checks if two values are deeply equal.
//
// By default, it uses [compare.DefaultComparator].
var DeepEqualer atomicutil.Value[func(v1, v2 any) (string, bool)]

func init() {
	DeepEqualer.Store(NewDeepEqualerWithComparator(compare.DefaultComparator.Load()))
}

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

// DeepEqual asserts that v1 and v2 are deeply equal according to [DeepEqualer].
//
//nolint:thelper // It's called below.
func DeepEqual[T any](tb testing.TB, v1, v2 T, opts ...Option) bool {
	diff, equal := DeepEqualer.Load()(v1, v2)
	ok := equal
	if !ok {
		tb.Helper()
		vs := ValueStringer.Load()
		Fail(
			tb,
			"deep_equal",
			fmt.Sprintf("not equal:\ndiff = %s\nv1 = %s\nv2 = %s", diff, vs(v1), vs(v2)),
			1,
			opts...,
		)
	}
	return ok
}

// NotDeepEqual asserts that v1 and v2 are not deeply equal according to [DeepEqualer].
//
//nolint:thelper // It's called below.
func NotDeepEqual[T any](tb testing.TB, v1, v2 T, opts ...Option) bool {
	_, equal := DeepEqualer.Load()(v1, v2)
	ok := !equal
	if !ok {
		tb.Helper()
		vs := ValueStringer.Load()
		Fail(
			tb,
			"not_deep_equal",
			fmt.Sprintf("equal:\nv1 = %s\nv2 = %s", vs(v1), vs(v2)),
			1,
			opts...,
		)
	}
	return ok
}
