package assert

import (
	"fmt"
	"testing"
)

// TODO: find a way to support receive/send only chans. (more functions or use type parameters ?)

// ChanNil asserts that c is nil.
func ChanNil[T any](tb testing.TB, c chan T, opts ...Option) bool {
	tb.Helper()
	ok := c == nil
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("chan_nil[%s]", TypeString[T]()),
			"not nil",
			opts...,
		)
	}
	return ok
}

// ChanNotNil asserts that c is not nil.
func ChanNotNil[T any](tb testing.TB, c chan T, opts ...Option) bool {
	tb.Helper()
	ok := c != nil
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("chan_not_nil[%s]", TypeString[T]()),
			"nil",
			opts...,
		)
	}
	return ok
}

// ChanEmpty asserts that c is empty.
func ChanEmpty[T any](tb testing.TB, c chan T, opts ...Option) bool {
	tb.Helper()
	ok := len(c) == 0
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("chan_empty[%s]", TypeString[T]()),
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
			fmt.Sprintf("chan_not_empty[%s]", TypeString[T]()),
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
			fmt.Sprintf("chan_len[%s]", TypeString[T]()),
			fmt.Sprintf("unexpected length:\nexpected = %d\nactual = %d", l, len(c)),
			opts...,
		)
	}
	return ok
}
