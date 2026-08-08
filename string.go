package assert

import (
	"fmt"
	"strings"
	"testing"
)

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
			fmt.Sprintf("unexpected length:\nexpected = %d\nactual = %d\ns = %s", l, len(s), ValueStringer.Load()(s)),
			1,
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
			fmt.Sprintf("not contains:\ns = %s\nsubstr = %q", ValueStringer.Load()(s), substr),
			1,
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
			fmt.Sprintf("contains:\ns = %s\nsubstr = %q", ValueStringer.Load()(s), substr),
			1,
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
			fmt.Sprintf("not has prefix:\ns = %s\nprefix = %q", ValueStringer.Load()(s), prefix),
			1,
			opts...,
		)
	}
	return ok
}

// StringNotHasPrefix asserts that s does not begin with prefix.
//
//nolint:thelper // It's called below.
func StringNotHasPrefix(tb testing.TB, s, prefix string, opts ...Option) bool {
	ok := !strings.HasPrefix(s, prefix)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"string_not_has_prefix",
			fmt.Sprintf("has prefix:\ns = %s\nprefix = %q", ValueStringer.Load()(s), prefix),
			1,
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
			fmt.Sprintf("not has suffix:\ns = %s\nsuffix = %q", ValueStringer.Load()(s), suffix),
			1,
			opts...,
		)
	}
	return ok
}

// StringNotHasSuffix asserts that s does not end with suffix.
//
//nolint:thelper // It's called below.
func StringNotHasSuffix(tb testing.TB, s, suffix string, opts ...Option) bool {
	ok := !strings.HasSuffix(s, suffix)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"string_not_has_suffix",
			fmt.Sprintf("has suffix:\ns = %s\nsuffix = %q", ValueStringer.Load()(s), suffix),
			1,
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
			fmt.Sprintf("not equal fold:\ns1 = %s\ns2 = %s", ValueStringer.Load()(s1), ValueStringer.Load()(s2)),
			1,
			opts...,
		)
	}
	return ok
}
