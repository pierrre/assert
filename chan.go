package assert

import (
	"fmt"
	"testing"
)

// ChanNil asserts that c is nil.
//
//nolint:thelper // It's called below.
func ChanNil[T any](tb testing.TB, c chan T, opts ...Option) bool {
	ok := c == nil
	if !ok {
		tb.Helper()
		vs := ValueStringer.Load()
		Fail(
			tb,
			"chan_nil",
			"not nil:\nc = "+vs(c),
			1,
			opts...,
		)
	}
	return ok
}

// ChanNotNil asserts that c is not nil.
//
//nolint:thelper // It's called below.
func ChanNotNil[T any](tb testing.TB, c chan T, opts ...Option) bool {
	ok := c != nil
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"chan_not_nil",
			"nil",
			1,
			opts...,
		)
	}
	return ok
}

// ChanEmpty asserts that c is empty.
//
//nolint:thelper // It's called below.
func ChanEmpty[T any](tb testing.TB, c chan T, opts ...Option) bool {
	ok := len(c) == 0
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"chan_empty",
			fmt.Sprintf("not empty:\nlength = %d", len(c)),
			1,
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
			"chan_not_empty",
			"empty",
			1,
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
			"chan_len",
			fmt.Sprintf("unexpected length:\nexpected = %d\nactual = %d", l, len(c)),
			1,
			opts...,
		)
	}
	return ok
}

// ChanRecvNil asserts that c is nil.
//
//nolint:thelper // It's called below.
func ChanRecvNil[T any](tb testing.TB, c <-chan T, opts ...Option) bool {
	ok := c == nil
	if !ok {
		tb.Helper()
		vs := ValueStringer.Load()
		Fail(
			tb,
			"chan_recv_nil",
			"not nil:\nc = "+vs(c),
			1,
			opts...,
		)
	}
	return ok
}

// ChanRecvNotNil asserts that c is not nil.
//
//nolint:thelper // It's called below.
func ChanRecvNotNil[T any](tb testing.TB, c <-chan T, opts ...Option) bool {
	ok := c != nil
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"chan_recv_not_nil",
			"nil",
			1,
			opts...,
		)
	}
	return ok
}

// ChanRecvEmpty asserts that c is empty.
//
//nolint:thelper // It's called below.
func ChanRecvEmpty[T any](tb testing.TB, c <-chan T, opts ...Option) bool {
	ok := len(c) == 0
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"chan_recv_empty",
			fmt.Sprintf("not empty:\nlength = %d", len(c)),
			1,
			opts...,
		)
	}
	return ok
}

// ChanRecvNotEmpty asserts that c is not empty.
//
//nolint:thelper // It's called below.
func ChanRecvNotEmpty[T any](tb testing.TB, c <-chan T, opts ...Option) bool {
	ok := len(c) != 0
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"chan_recv_not_empty",
			"empty",
			1,
			opts...,
		)
	}
	return ok
}

// ChanRecvLen asserts that c has length l.
//
//nolint:thelper // It's called below.
func ChanRecvLen[T any](tb testing.TB, c <-chan T, l int, opts ...Option) bool {
	ok := len(c) == l
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"chan_recv_len",
			fmt.Sprintf("unexpected length:\nexpected = %d\nactual = %d", l, len(c)),
			1,
			opts...,
		)
	}
	return ok
}

// ChanSendNil asserts that c is nil.
//
//nolint:thelper // It's called below.
func ChanSendNil[T any](tb testing.TB, c chan<- T, opts ...Option) bool {
	ok := c == nil
	if !ok {
		tb.Helper()
		vs := ValueStringer.Load()
		Fail(
			tb,
			"chan_send_nil",
			"not nil:\nc = "+vs(c),
			1,
			opts...,
		)
	}
	return ok
}

// ChanSendNotNil asserts that c is not nil.
//
//nolint:thelper // It's called below.
func ChanSendNotNil[T any](tb testing.TB, c chan<- T, opts ...Option) bool {
	ok := c != nil
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"chan_send_not_nil",
			"nil",
			1,
			opts...,
		)
	}
	return ok
}

// ChanSendEmpty asserts that c is empty.
//
//nolint:thelper // It's called below.
func ChanSendEmpty[T any](tb testing.TB, c chan<- T, opts ...Option) bool {
	ok := len(c) == 0
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"chan_send_empty",
			fmt.Sprintf("not empty:\nlength = %d", len(c)),
			1,
			opts...,
		)
	}
	return ok
}

// ChanSendNotEmpty asserts that c is not empty.
//
//nolint:thelper // It's called below.
func ChanSendNotEmpty[T any](tb testing.TB, c chan<- T, opts ...Option) bool {
	ok := len(c) != 0
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"chan_send_not_empty",
			"empty",
			1,
			opts...,
		)
	}
	return ok
}

// ChanSendLen asserts that c has length l.
//
//nolint:thelper // It's called below.
func ChanSendLen[T any](tb testing.TB, c chan<- T, l int, opts ...Option) bool {
	ok := len(c) == l
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"chan_send_len",
			fmt.Sprintf("unexpected length:\nexpected = %d\nactual = %d", l, len(c)),
			1,
			opts...,
		)
	}
	return ok
}
