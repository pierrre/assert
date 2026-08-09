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
		vs := ValueStringer.Load()
		Fail(
			tb,
			"no_error",
			"error: "+vs(err),
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
		vs := ValueStringer.Load()
		Fail(
			tb,
			"error_is",
			fmt.Sprintf("no match:\nerr = %s\ntarget = %s", vs(err), vs(target)),
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
		vs := ValueStringer.Load()
		Fail(
			tb,
			"error_not_is",
			fmt.Sprintf("match:\nerr = %s\ntarget = %s", vs(err), vs(target)),
			1,
			opts...,
		)
	}
	return ok
}

// ErrorAs asserts that [errors.As] returns true.
//
// target must be a non-nil pointer to a type that implements error, or to any interface type.
//
//nolint:thelper // It's called below.
func ErrorAs(tb testing.TB, err error, target any, opts ...Option) bool {
	ok := errors.As(err, target)
	if !ok {
		tb.Helper()
		vs := ValueStringer.Load()
		Fail(
			tb,
			"error_as",
			fmt.Sprintf("no match:\nerr = %s\ntarget = %s", vs(err), vs(target)),
			1,
			opts...,
		)
	}
	return ok
}

// ErrorAsType asserts that [errors.AsType] returns true and returns the error of type E.
//
//nolint:thelper // It's called below.
func ErrorAsType[E error](tb testing.TB, err error, opts ...Option) (E, bool) {
	e, ok := errors.AsType[E](err)
	if !ok {
		tb.Helper()
		vs := ValueStringer.Load()
		Fail(
			tb,
			"error_as_type",
			fmt.Sprintf("no match:\nerr = %s\ntype = %T", vs(err), *new(E)),
			1,
			opts...,
		)
	}
	return e, ok
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
		vs := ValueStringer.Load()
		Fail(
			tb,
			"error_equal",
			fmt.Sprintf("not equal:\nerr = %s\nmessage = %q", vs(err), message),
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
		vs := ValueStringer.Load()
		Fail(
			tb,
			"error_contains",
			fmt.Sprintf("not contains:\nerr = %s\nsubstr = %q", vs(err), substr),
			1,
			opts...,
		)
	}
	return ok
}
