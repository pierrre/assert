package assert

import (
	"fmt"
	"reflect" //nolint:depguard // Used for DeepEqual.
	"testing"
)

// DeepEqualer is a function that checks if two values are deep equal.
//
// It can be customized to provide a better comparison.
//
// By default it uses reflect.DeepEqual.
var DeepEqualer = func(v1, v2 any) (diff string, equal bool) {
	equal = reflect.DeepEqual(v1, v2)
	if equal {
		return "", true
	}
	return "not implemented", false
}

// DeepEqual asserts that v1 and v2 are deep equal according to DeepEqualer.
func DeepEqual[T any](tb testing.TB, v1, v2 T, opts ...Option) bool {
	tb.Helper()
	diff, equal := DeepEqualer(v1, v2)
	ok := equal
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("deep_equal[%s]", TypeString[T]()),
			fmt.Sprintf("not equal:\ndiff = %s\nv1 = %s\nv2 = %s", diff, ValueStringer(v1), ValueStringer(v2)),
			opts...,
		)
	}
	return ok
}

// NotDeepEqual asserts that v1 and v2 are not deep equal according to DeepEqualer.
func NotDeepEqual[T any](tb testing.TB, v1, v2 T, opts ...Option) bool {
	tb.Helper()
	_, equal := DeepEqualer(v1, v2)
	ok := !equal
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("not_deep_equal[%s]", TypeString[T]()),
			fmt.Sprintf("equal:\nv1 = %s\nv2 = %s", ValueStringer(v1), ValueStringer(v2)),
			opts...,
		)
	}
	return ok
}
