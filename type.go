package assert

import (
	"fmt"
	"testing"

	"github.com/pierrre/go-libs/reflectutil"
)

// Type asserts that v is of type T, and returns it.
//
//nolint:thelper // It's called below.
func Type[T any](tb testing.TB, v any, opts ...Option) (T, bool) {
	vt, ok := v.(T)
	if !ok {
		tb.Helper()
		vs := ValueStringer.Load()
		Fail(
			tb,
			"type",
			fmt.Sprintf("assertion failed:\nsource = %s\ndestination = %s", vs(v), reflectutil.TypeFullNameFor[T]()),
			1,
			opts...,
		)
	}
	return vt, ok
}

// NotType asserts that v is not of type T.
//
//nolint:thelper // It's called below.
func NotType[T any](tb testing.TB, v any, opts ...Option) bool {
	_, ok := v.(T)
	if ok {
		tb.Helper()
		vs := ValueStringer.Load()
		Fail(
			tb,
			"not_type",
			fmt.Sprintf("assertion failed:\nsource = %s\ntype = %s", vs(v), reflectutil.TypeFullNameFor[T]()),
			1,
			opts...,
		)
	}
	return !ok
}
