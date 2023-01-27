package assert

import (
	"fmt"
	"reflect" //nolint:depguard // Used for type checking.
	"testing"
)

// TODO: remove once Go 1.20 is supported.

// Nil asserts that v is nil.
//
// This function exists because it's not possible to define a type parameter for interface/function.
// Use SliceNil/MapNil/ChanNil/PointerNil if possible.
func Nil(tb testing.TB, v any, opts ...Option) bool {
	tb.Helper()
	ok := isNil(v)
	if !ok {
		Fail(
			tb,
			"nil",
			fmt.Sprintf("not nil:\nv = %s", ValueStringer(v)),
			opts...,
		)
	}
	return ok
}

// NotNil asserts that v is not nil.
//
// This function exists because it's not possible to define a type parameter for interface/function.
// Use SliceNotNil/MapNotNil/ChanNotNil/PointerNotNil if possible.
func NotNil(tb testing.TB, v any, opts ...Option) bool {
	tb.Helper()
	ok := !isNil(v)
	if !ok {
		Fail(
			tb,
			"not_nil",
			"nil",
			opts...,
		)
	}
	return ok
}

func isNil(v any) bool {
	if v == nil {
		return true
	}
	value := reflect.ValueOf(v)
	kind := value.Kind()
	switch kind { //nolint:exhaustive // We only care about the kinds that can be nil.
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return value.IsNil()
	}
	return false
}
