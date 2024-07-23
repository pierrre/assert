package assert

import (
	"fmt"
	"testing"

	"github.com/pierrre/go-libs/reflectutil"
)

// TODO: find a way to support receive/send only chans. (more functions or use type parameters ?)

// ChanEmpty asserts that c is empty.
func ChanEmpty[T any](tb testing.TB, c chan T, opts ...Option) bool {
	tb.Helper()
	ok := len(c) == 0
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("chan_empty[%s]", reflectutil.TypeFullNameFor[T]()),
			fmt.Sprintf("not empty:\nlength = %d", len(c)),
			opts...,
		)
	}
	return ok
}

// ChanNotEmpty asserts that c is not empty.
func ChanNotEmpty[T any](tb testing.TB, c chan T, opts ...Option) bool {
	tb.Helper()
	ok := len(c) != 0
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("chan_not_empty[%s]", reflectutil.TypeFullNameFor[T]()),
			"empty",
			opts...,
		)
	}
	return ok
}

// ChanLen asserts that c has length l.
func ChanLen[T any](tb testing.TB, c chan T, l int, opts ...Option) bool {
	tb.Helper()
	ok := len(c) == l
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("chan_len[%s]", reflectutil.TypeFullNameFor[T]()),
			fmt.Sprintf("unexpected length:\nexpected = %d\nactual = %d", l, len(c)),
			opts...,
		)
	}
	return ok
}
