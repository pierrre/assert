package assert

import (
	"fmt"
	"strings"
	"testing"
)

// StringEmpty asserts that s is empty.
func StringEmpty(tb testing.TB, s string, opts ...Option) bool {
	tb.Helper()
	ok := s == ""
	if !ok {
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
func StringNotEmpty(tb testing.TB, s string, opts ...Option) bool {
	tb.Helper()
	ok := s != ""
	if !ok {
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
func StringLen(tb testing.TB, s string, l int, opts ...Option) bool {
	tb.Helper()
	ok := len(s) == l
	if !ok {
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
func StringContains(tb testing.TB, s, substr string, opts ...Option) bool {
	tb.Helper()
	ok := strings.Contains(s, substr)
	if !ok {
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
func StringNotContains(tb testing.TB, s, substr string, opts ...Option) bool {
	tb.Helper()
	ok := !strings.Contains(s, substr)
	if !ok {
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
func StringHasPrefix(tb testing.TB, s, prefix string, opts ...Option) bool {
	tb.Helper()
	ok := strings.HasPrefix(s, prefix)
	if !ok {
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
func StringHasSuffix(tb testing.TB, s, suffix string, opts ...Option) bool {
	tb.Helper()
	ok := strings.HasSuffix(s, suffix)
	if !ok {
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
func StringEqualFold(tb testing.TB, s1, s2 string, opts ...Option) bool {
	tb.Helper()
	ok := strings.EqualFold(s1, s2)
	if !ok {
		Fail(
			tb,
			"string_equal_fold",
			fmt.Sprintf("not equal fold:\ns1 = %q\ns2 = %q", s1, s2),
			opts...,
		)
	}
	return ok
}
