package assert

import (
	"fmt"
	"regexp"
	"testing"
)

// RegexpString is a type that can be either a *regexp.Regexp or a string.
//
// If it's a string, it's automatically compiled to a *regexp.Regexp.
type RegexpString interface {
	*regexp.Regexp | string
}

// RegexpMatch asserts that rs matches s.
func RegexpMatch[RS RegexpString](tb testing.TB, rs RS, s string, opts ...Option) bool {
	tb.Helper()
	r := getRegexp(tb, rs, opts...)
	if r == nil {
		return false
	}
	ok := r.MatchString(s)
	if !ok {
		Fail(
			tb,
			"regexp_match",
			fmt.Sprintf("no match:\nrs = %q\ns = %q", r, s),
			opts...,
		)
	}
	return ok
}

// RegexpNotMatch asserts that rs doesn't match s.
func RegexpNotMatch[RS RegexpString](tb testing.TB, rs RS, s string, opts ...Option) bool {
	tb.Helper()
	r := getRegexp(tb, rs, opts...)
	if r == nil {
		return false
	}
	ok := !r.MatchString(s)
	if !ok {
		Fail(
			tb,
			"regexp_not_match",
			fmt.Sprintf("match:\nrs = %q\ns = %q", r, s),
			opts...,
		)
	}
	return ok
}

func getRegexp[RS RegexpString](tb testing.TB, rs RS, opts ...Option) *regexp.Regexp {
	tb.Helper()
	r, ok := any(rs).(*regexp.Regexp)
	if ok {
		return r
	}
	s := any(rs).(string) //nolint:forcetypeassert // We know it's a string.
	r, err := regexp.Compile(s)
	if err != nil {
		Fail(
			tb,
			"regexp_compile",
			fmt.Sprintf("compilation failed:\nexpr = %q\nerr = %s", s, ErrorStringer(err)),
			opts...,
		)
		return nil
	}
	return r
}
