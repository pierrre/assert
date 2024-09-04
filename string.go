package assert

import (
	"fmt"
	"strings"
	"testing"
)

// StringEmpty asserts that s is empty.
//
//nolint:thelper // It's called below.
func StringEmpty(tb testing.TB, s string, opts ...Option) bool {
	ok := s == ""
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"string_empty",
			fmt.Sprintf("not empty:\ns = %q", s),
			opts...,
		)
	}
	return ok
}

// StringNotEmpty asserts that s is not empty.
//
//nolint:thelper // It's called below.
func StringNotEmpty(tb testing.TB, s string, opts ...Option) bool {
	ok := s != ""
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"string_not_empty",
			"empty",
			opts...,
		)
	}
	return ok
}

// StringLen asserts that s has length l.
//
//nolint:thelper // It's called below.
func StringLen(tb testing.TB, s string, l int, opts ...Option) bool {
	ok := len(s) == l
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"string_len",
			fmt.Sprintf("unexpected length:\nexpected = %d\nactual = %d", l, len(s)),
			opts...,
		)
	}
	return ok
}

// StringContains asserts that s contains substr.
//
//nolint:thelper // It's called below.
func StringContains(tb testing.TB, s, substr string, opts ...Option) bool {
	ok := strings.Contains(s, substr)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"string_contains",
			fmt.Sprintf("not contains:\ns = %q\nsubstr = %q", s, substr),
			opts...,
		)
	}
	return ok
}

// StringNotContains asserts that s does not contain substr.
//
//nolint:thelper // It's called below.
func StringNotContains(tb testing.TB, s, substr string, opts ...Option) bool {
	ok := !strings.Contains(s, substr)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"string_not_contains",
			fmt.Sprintf("contains:\ns = %q\nsubstr = %q", s, substr),
			opts...,
		)
	}
	return ok
}

// StringHasPrefix asserts that s begins with prefix.
//
//nolint:thelper // It's called below.
func StringHasPrefix(tb testing.TB, s, prefix string, opts ...Option) bool {
	ok := strings.HasPrefix(s, prefix)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"string_has_prefix",
			fmt.Sprintf("no prefix:\ns = %q\nprefix = %q", s, prefix),
			opts...,
		)
	}
	return ok
}

// StringHasSuffix asserts that s ends with suffix.
//
//nolint:thelper // It's called below.
func StringHasSuffix(tb testing.TB, s, suffix string, opts ...Option) bool {
	ok := strings.HasSuffix(s, suffix)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"string_has_suffix",
			fmt.Sprintf("no suffix:\ns = %q\nsuffix = %q", s, suffix),
			opts...,
		)
	}
	return ok
}

// StringEqualFold asserts that s1 and s2 are equal, ignoring case.
//
//nolint:thelper // It's called below.
func StringEqualFold(tb testing.TB, s1, s2 string, opts ...Option) bool {
	ok := strings.EqualFold(s1, s2)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"string_equal_fold",
			fmt.Sprintf("not equal fold:\ns1 = %q\ns2 = %q", s1, s2),
			opts...,
		)
	}
	return ok
}
