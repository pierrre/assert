package assert

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

// ErrorStringer is a function that returns a string representation of an error.
//
// It can be customized to provide a better error message.
var ErrorStringer func(error) string = func(err error) string {
	return fmt.Sprintf("%q", err)
}

// Error asserts that err is not nil.
func Error(tb testing.TB, err error, opts ...Option) bool {
	tb.Helper()
	ok := err != nil
	if !ok {
		Fail(
			tb,
			"error",
			"no error",
			opts...,
		)
	}
	return ok
}

// NoError asserts that err is nil.
func NoError(tb testing.TB, err error, opts ...Option) bool {
	tb.Helper()
	ok := err == nil
	if !ok {
		Fail(
			tb,
			"no_error",
			"error: "+ErrorStringer(err),
			opts...,
		)
	}
	return ok
}

// ErrorIs asserts that [errors.Is] returns true.
func ErrorIs(tb testing.TB, err, target error, opts ...Option) bool {
	tb.Helper()
	ok := errors.Is(err, target)
	if !ok {
		Fail(
			tb,
			"error_is",
			fmt.Sprintf("no match:\nerr = %s\ntarget = %s", ErrorStringer(err), ErrorStringer(target)),
			opts...,
		)
	}
	return ok
}

// ErrorNotIs asserts that [errors.Is] returns false.
func ErrorNotIs(tb testing.TB, err, target error, opts ...Option) bool {
	tb.Helper()
	ok := !errors.Is(err, target)
	if !ok {
		Fail(
			tb,
			"error_not_is",
			fmt.Sprintf("match:\nerr = %s\ntarget = %s", ErrorStringer(err), ErrorStringer(target)),
			opts...,
		)
	}
	return ok
}

// ErrorAs asserts that [errors.As] returns true.
func ErrorAs(tb testing.TB, err error, target any, opts ...Option) bool {
	tb.Helper()
	ok := errors.As(err, target)
	if !ok {
		Fail(
			tb,
			"error_as",
			fmt.Sprintf("no match:\nerr = %s\ntarget = %T", ErrorStringer(err), target),
			opts...,
		)
	}
	return ok
}

// ErrorEqual asserts that the result of [error.Error] is equal to message.
func ErrorEqual(tb testing.TB, err error, message string, opts ...Option) bool {
	tb.Helper()
	ok := err.Error() == message
	if !ok {
		Fail(
			tb,
			"error_equal",
			fmt.Sprintf("not equal:\nerr = %s\nmessage = %q", ErrorStringer(err), message),
			opts...,
		)
	}
	return ok
}

// ErrorContains asserts that the result of [error.Error] contains substr.
func ErrorContains(tb testing.TB, err error, substr string, opts ...Option) bool {
	tb.Helper()
	ok := strings.Contains(err.Error(), substr)
	if !ok {
		Fail(
			tb,
			"error_contains",
			fmt.Sprintf("not contains:\nerr = %s\nsubstr = %q", ErrorStringer(err), substr),
			opts...,
		)
	}
	return ok
}
