package assert

import (
	"fmt"
	"testing"
)

// Type asserts that v is of type T, and returns it.
func Type[T any](tb testing.TB, v any, opts ...Option) (T, bool) {
	tb.Helper()
	vt, ok := v.(T)
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("type[%s]", TypeString[T]()),
			fmt.Sprintf("assertion failed:\nsource = %T\ndestination = %s", v, TypeString[T]()),
			opts...,
		)
	}
	return vt, ok
}
