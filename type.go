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
		Fail(
			tb,
			"type",
			fmt.Sprintf("assertion failed:\nsource = %T\ndestination = %s", v, reflectutil.TypeFullNameFor[T]()),
			opts...,
		)
	}
	return vt, ok
}
