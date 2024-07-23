package assert

import (
	"fmt"
	"testing"

	"github.com/pierrre/go-libs/reflectutil"
)

// Type asserts that v is of type T, and returns it.
func Type[T any](tb testing.TB, v any, opts ...Option) (T, bool) {
	tb.Helper()
	vt, ok := v.(T)
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("type[%s]", reflectutil.TypeFullNameFor[T]()),
			fmt.Sprintf("assertion failed:\nsource = %T\ndestination = %s", v, reflectutil.TypeFullNameFor[T]()),
			opts...,
		)
	}
	return vt, ok
}
