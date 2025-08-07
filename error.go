package assert

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

// Error asserts that err is not nil.
//
//nolint:thelper // It's called below.
func Error(tb testing.TB, err error, opts ...Option) bool {
	ok := err != nil
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"error",
			"no error",
			1,
			opts...,
		)
	}
	return ok
}

// NoError asserts that err is nil.
//
//nolint:thelper // It's called below.
func NoError(tb testing.TB, err error, opts ...Option) bool {
	ok := err == nil
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"no_error",
			"error: "+ValueStringer(err),
			1,
			opts...,
		)
	}
	return ok
}

// ErrorIs asserts that [errors.Is] returns true.
//
//nolint:thelper // It's called below.
func ErrorIs(tb testing.TB, err, target error, opts ...Option) bool {
	ok := errors.Is(err, target)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"error_is",
			fmt.Sprintf("no match:\nerr = %s\ntarget = %s", ValueStringer(err), ValueStringer(target)),
			1,
			opts...,
		)
	}
	return ok
}

// ErrorNotIs asserts that [errors.Is] returns false.
//
//nolint:thelper // It's called below.
func ErrorNotIs(tb testing.TB, err, target error, opts ...Option) bool {
	ok := !errors.Is(err, target)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"error_not_is",
			fmt.Sprintf("match:\nerr = %s\ntarget = %s", ValueStringer(err), ValueStringer(target)),
			1,
			opts...,
		)
	}
	return ok
}

// ErrorAs asserts that [errors.As] returns true.
//
//nolint:thelper // It's called below.
func ErrorAs(tb testing.TB, err error, target any, opts ...Option) bool {
	ok := errors.As(err, target)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"error_as",
			fmt.Sprintf("no match:\nerr = %s\ntarget = %T", ValueStringer(err), target),
			1,
			opts...,
		)
	}
	return ok
}

// ErrorEqual asserts that the result of [error.Error] is equal to message.
func ErrorEqual(tb testing.TB, err error, message string, opts ...Option) bool {
	tb.Helper()
	ok := Error(tb, err, opts...)
	if !ok {
		return false
	}
	ok = err.Error() == message
	if !ok {
		Fail(
			tb,
			"error_equal",
			fmt.Sprintf("not equal:\nerr = %s\nmessage = %q", ValueStringer(err), message),
			1,
			opts...,
		)
	}
	return ok
}

// ErrorContains asserts that the result of [error.Error] contains substr.
func ErrorContains(tb testing.TB, err error, substr string, opts ...Option) bool {
	tb.Helper()
	ok := Error(tb, err, opts...)
	if !ok {
		return false
	}
	ok = strings.Contains(err.Error(), substr)
	if !ok {
		Fail(
			tb,
			"error_contains",
			fmt.Sprintf("not contains:\nerr = %s\nsubstr = %q", ValueStringer(err), substr),
			1,
			opts...,
		)
	}
	return ok
}
