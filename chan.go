package assert

import (
	"fmt"
	"testing"
)

// TODO: find a way to support receive/send only chans. (more functions or use type parameters ?)

// ChanEmpty asserts that c is empty.
//
//nolint:thelper // It's called below.
func ChanEmpty[T any](tb testing.TB, c chan T, opts ...Option) bool {
	ok := len(c) == 0
	if !ok {
		tb.Helper()
		Fail(
			tb,
			fmt.Sprintf("chan_empty[%s]", typeName[T]()),
			fmt.Sprintf("not empty:\nlength = %d", len(c)),
			opts...,
		)
	}
	return ok
}

// ChanNotEmpty asserts that c is not empty.
//
//nolint:thelper // It's called below.
func ChanNotEmpty[T any](tb testing.TB, c chan T, opts ...Option) bool {
	ok := len(c) != 0
	if !ok {
		tb.Helper()
		Fail(
			tb,
			fmt.Sprintf("chan_not_empty[%s]", typeName[T]()),
			"empty",
			opts...,
		)
	}
	return ok
}

// ChanLen asserts that c has length l.
//
//nolint:thelper // It's called below.
func ChanLen[T any](tb testing.TB, c chan T, l int, opts ...Option) bool {
	ok := len(c) == l
	if !ok {
		tb.Helper()
		Fail(
			tb,
			fmt.Sprintf("chan_len[%s]", typeName[T]()),
			fmt.Sprintf("unexpected length:\nexpected = %d\nactual = %d", l, len(c)),
			opts...,
		)
	}
	return ok
}
