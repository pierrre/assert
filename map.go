package assert

import (
	"fmt"
	"maps"
	"testing"
)

// MapNil asserts that m is nil.
func MapNil[M ~map[K]V, K comparable, V any](tb testing.TB, m M, opts ...Option) bool {
	tb.Helper()
	ok := m == nil
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("map_nil[%s, %s]", typeName[K](), typeName[V]()),
			"not nil:\nm = "+ValueStringer(m),
			opts...,
		)
	}
	return ok
}

// MapNotNil asserts that m is not nil.
func MapNotNil[M ~map[K]V, K comparable, V any](tb testing.TB, m M, opts ...Option) bool {
	tb.Helper()
	ok := m != nil
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("map_not_nil[%s, %s]", typeName[K](), typeName[V]()),
			"nil",
			opts...,
		)
	}
	return ok
}

// MapEmpty asserts that m is empty.
func MapEmpty[M ~map[K]V, K comparable, V any](tb testing.TB, m M, opts ...Option) bool {
	tb.Helper()
	ok := len(m) == 0
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("map_empty[%s, %s]", typeName[K](), typeName[V]()),
			fmt.Sprintf("not empty:\nlength = %d\nm = %s", len(m), ValueStringer(m)),
			opts...,
		)
	}
	return ok
}

// MapNotEmpty asserts that m is not empty.
func MapNotEmpty[M ~map[K]V, K comparable, V any](tb testing.TB, m M, opts ...Option) bool {
	tb.Helper()
	ok := len(m) != 0
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("map_not_empty[%s, %s]", typeName[K](), typeName[V]()),
			"empty",
			opts...,
		)
	}
	return ok
}

// MapLen asserts that m has length l.
func MapLen[M ~map[K]V, K comparable, V any](tb testing.TB, m M, l int, opts ...Option) bool {
	tb.Helper()
	ok := len(m) == l
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("map_len[%s, %s]", typeName[K](), typeName[V]()),
			fmt.Sprintf("unexpected length:\nexpected = %d\nactual = %d", l, len(m)),
			opts...,
		)
	}
	return ok
}

// MapEqual asserts that m1 and m2 are equal.
func MapEqual[M1, M2 ~map[K]V, K, V comparable](tb testing.TB, m1 M1, m2 M2, opts ...Option) bool {
	tb.Helper()
	ok := maps.Equal(m1, m2)
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("map_equal[%s, %s]", typeName[K](), typeName[V]()),
			fmt.Sprintf("not equal:\nm1 = %s\nm2 = %s", ValueStringer(m1), ValueStringer(m2)),
			opts...,
		)
	}
	return ok
}

// MapNotEqual asserts that m1 and m2 are not equal.
func MapNotEqual[M1, M2 ~map[K]V, K, V comparable](tb testing.TB, m1 M1, m2 M2, opts ...Option) bool {
	tb.Helper()
	ok := !maps.Equal(m1, m2)
	if !ok {
		Fail(
			tb,
			fmt.Sprintf("map_not_equal[%s, %s]", typeName[K](), typeName[V]()),
			fmt.Sprintf("equal:\nm1 = %s\nm2 = %s", ValueStringer(m1), ValueStringer(m2)),
			opts...,
		)
	}
	return ok
}
