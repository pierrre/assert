package assert

import (
	"fmt"
	"slices"
	"testing"
)

// SliceNil asserts that s is nil.
//
//nolint:thelper // It's called below.
func SliceNil[S ~[]E, E any](tb testing.TB, s S, opts ...Option) bool {
	ok := s == nil
	if !ok {
		tb.Helper()
		Fail(
			tb,
			fmt.Sprintf("slice_nil[%s]", typeName[E]()),
			"not nil:\ns = "+ValueStringer(s),
			opts...,
		)
	}
	return ok
}

// SliceNotNil asserts that s is not nil.
//
//nolint:thelper // It's called below.
func SliceNotNil[S ~[]E, E any](tb testing.TB, s S, opts ...Option) bool {
	ok := s != nil
	if !ok {
		tb.Helper()
		Fail(
			tb,
			fmt.Sprintf("slice_not_nil[%s]", typeName[E]()),
			"nil",
			opts...,
		)
	}
	return ok
}

// SliceEmpty asserts that s is empty.
//
//nolint:thelper // It's called below.
func SliceEmpty[S ~[]E, E any](tb testing.TB, s S, opts ...Option) bool {
	ok := len(s) == 0
	if !ok {
		tb.Helper()
		Fail(
			tb,
			fmt.Sprintf("slice_empty[%s]", typeName[E]()),
			fmt.Sprintf("not empty:\nlength = %d\ns = %s", len(s), ValueStringer(s)),
			opts...,
		)
	}
	return ok
}

// SliceNotEmpty asserts that s is not empty.
//
//nolint:thelper // It's called below.
func SliceNotEmpty[S ~[]E, E any](tb testing.TB, s S, opts ...Option) bool {
	ok := len(s) != 0
	if !ok {
		tb.Helper()
		Fail(
			tb,
			fmt.Sprintf("slice_not_empty[%s]", typeName[E]()),
			"empty",
			opts...,
		)
	}
	return ok
}

// SliceLen asserts that s has length l.
//
//nolint:thelper // It's called below.
func SliceLen[S ~[]E, E any](tb testing.TB, s S, l int, opts ...Option) bool {
	ok := len(s) == l
	if !ok {
		tb.Helper()
		Fail(
			tb,
			fmt.Sprintf("slice_len[%s]", typeName[E]()),
			fmt.Sprintf("unexpected length:\nexpected = %d\nactual = %d", l, len(s)),
			opts...,
		)
	}
	return ok
}

// SliceEqual asserts that s1 and s2 are equal.
//
//nolint:thelper // It's called below.
func SliceEqual[S ~[]E, E comparable](tb testing.TB, s1, s2 S, opts ...Option) bool {
	ok := slices.Equal(s1, s2)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			fmt.Sprintf("slice_equal[%s]", typeName[E]()),
			fmt.Sprintf("not equal:\ns1 = %s\ns2 = %s", ValueStringer(s1), ValueStringer(s2)),
			opts...,
		)
	}
	return ok
}

// SliceNotEqual asserts that s1 and s2 are not equal.
//
//nolint:thelper // It's called below.
func SliceNotEqual[S ~[]E, E comparable](tb testing.TB, s1, s2 S, opts ...Option) bool {
	ok := !slices.Equal(s1, s2)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			fmt.Sprintf("slice_not_equal[%s]", typeName[E]()),
			fmt.Sprintf("equal:\ns1 = %s\ns2 = %s", ValueStringer(s1), ValueStringer(s2)),
			opts...,
		)
	}
	return ok
}

// SliceContains asserts that s contains v.
//
//nolint:thelper // It's called below.
func SliceContains[S ~[]E, E comparable](tb testing.TB, s S, v E, opts ...Option) bool {
	ok := slices.Contains(s, v)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			fmt.Sprintf("slice_contains[%s]", typeName[E]()),
			fmt.Sprintf("not contains:\ns = %s\nv = %s", ValueStringer(s), ValueStringer(v)),
			opts...,
		)
	}
	return ok
}

// SliceNotContains asserts that s does not contain v.
//
//nolint:thelper // It's called below.
func SliceNotContains[S ~[]E, E comparable](tb testing.TB, s S, v E, opts ...Option) bool {
	ok := !slices.Contains(s, v)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			fmt.Sprintf("slice_not_contains[%s]", typeName[E]()),
			fmt.Sprintf("contains:\ns = %s\nv = %s", ValueStringer(s), ValueStringer(v)),
			opts...,
		)
	}
	return ok
}

// SliceContainsAll asserts that s1 contains all elements in s2.
//
//nolint:thelper // It's called below.
func SliceContainsAll[S ~[]E, E comparable](tb testing.TB, s1, s2 S, opts ...Option) bool {
	ok := sliceContainsAll(s1, s2)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			fmt.Sprintf("slice_contains_all[%s]", typeName[E]()),
			fmt.Sprintf("not contains all:\ns1 = %s\ns2 = %s", ValueStringer(s1), ValueStringer(s2)),
			opts...,
		)
	}
	return ok
}

// SliceNotContainsAll asserts that s1 does not contain all elements in s2.
//
//nolint:thelper // It's called below.
func SliceNotContainsAll[S ~[]E, E comparable](tb testing.TB, s1, s2 S, opts ...Option) bool {
	ok := !sliceContainsAll(s1, s2)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			fmt.Sprintf("slice_not_contains_all[%s]", typeName[E]()),
			fmt.Sprintf("contains all:\ns1 = %s\ns2 = %s", ValueStringer(s1), ValueStringer(s2)),
			opts...,
		)
	}
	return ok
}

func sliceContainsAll[S ~[]E, E comparable](s1, s2 S) bool {
	for _, v := range s2 {
		if !slices.Contains(s1, v) {
			return false
		}
	}
	return true
}
