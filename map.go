package assert

import (
	"fmt"
	"maps"
	"testing"
)

// MapNil asserts that m is nil.
//
//nolint:thelper // It's called below.
func MapNil[M ~map[K]V, K comparable, V any](tb testing.TB, m M, opts ...Option) bool {
	ok := m == nil
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"map_nil",
			"not nil:\nm = "+ValueStringer(m),
			1,
			opts...,
		)
	}
	return ok
}

// MapNotNil asserts that m is not nil.
//
//nolint:thelper // It's called below.
func MapNotNil[M ~map[K]V, K comparable, V any](tb testing.TB, m M, opts ...Option) bool {
	ok := m != nil
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"map_not_nil",
			"nil",
			1,
			opts...,
		)
	}
	return ok
}

// MapEmpty asserts that m is empty.
//
//nolint:thelper // It's called below.
func MapEmpty[M ~map[K]V, K comparable, V any](tb testing.TB, m M, opts ...Option) bool {
	ok := len(m) == 0
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"map_empty",
			fmt.Sprintf("not empty:\nlength = %d\nm = %s", len(m), ValueStringer(m)),
			1,
			opts...,
		)
	}
	return ok
}

// MapNotEmpty asserts that m is not empty.
//
//nolint:thelper // It's called below.
func MapNotEmpty[M ~map[K]V, K comparable, V any](tb testing.TB, m M, opts ...Option) bool {
	ok := len(m) != 0
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"map_not_empty",
			"empty",
			1,
			opts...,
		)
	}
	return ok
}

// MapLen asserts that m has length l.
//
//nolint:thelper // It's called below.
func MapLen[M ~map[K]V, K comparable, V any](tb testing.TB, m M, l int, opts ...Option) bool {
	ok := len(m) == l
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"map_len",
			fmt.Sprintf("unexpected length:\nexpected = %d\nactual = %d", l, len(m)),
			1,
			opts...,
		)
	}
	return ok
}

// MapEqual asserts that m1 and m2 are equal.
//
//nolint:thelper // It's called below.
func MapEqual[M1, M2 ~map[K]V, K, V comparable](tb testing.TB, m1 M1, m2 M2, opts ...Option) bool {
	ok := maps.Equal(m1, m2)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"map_equal",
			fmt.Sprintf("not equal:\nm1 = %s\nm2 = %s", ValueStringer(m1), ValueStringer(m2)),
			1,
			opts...,
		)
	}
	return ok
}

// MapNotEqual asserts that m1 and m2 are not equal.
//
//nolint:thelper // It's called below.
func MapNotEqual[M1, M2 ~map[K]V, K, V comparable](tb testing.TB, m1 M1, m2 M2, opts ...Option) bool {
	ok := !maps.Equal(m1, m2)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"map_not_equal",
			fmt.Sprintf("equal:\nm1 = %s\nm2 = %s", ValueStringer(m1), ValueStringer(m2)),
			1,
			opts...,
		)
	}
	return ok
}
