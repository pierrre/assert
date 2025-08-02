package assert

import (
	"fmt"
	"runtime/debug"
	"testing"
)

// Panics asserts that the code inside the function f panics.
//
// It returns the recovered value.
func Panics(tb testing.TB, f func(), opts ...Option) (rec any, ok bool) {
	tb.Helper()
	ok = true
	defer func() {
		tb.Helper()
		rec = recover()
		if rec == nil {
			ok = false
			Fail(
				tb,
				"panics",
				"no panic",
				1,
				opts...,
			)
		}
	}()
	f()
	return
}

// NotPanics asserts that the code inside the function f does not panic.
func NotPanics(tb testing.TB, f func(), opts ...Option) (ok bool) {
	tb.Helper()
	ok = true
	defer func() {
		tb.Helper()
		rec := recover()
		if rec != nil {
			ok = false
			st := string(debug.Stack())
			Fail(
				tb,
				"not_panics",
				fmt.Sprintf("panic:\npanic = %s\nstack = %s", ValueStringer(rec), st),
				1,
				opts...,
			)
		}
	}()
	f()
	return
}
