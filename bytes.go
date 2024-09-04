package assert

import (
	"bytes"
	"fmt"
	"testing"
)

// BytesEqual asserts that b1 and b2 are equal.
// It uses [bytes.Equal] to compare the two byte slices.
//
//nolint:thelper // It's called below.
func BytesEqual(tb testing.TB, b1, b2 []byte, opts ...Option) bool {
	ok := bytes.Equal(b1, b2)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"bytes_equal",
			fmt.Sprintf("not equal:\nb1 = %s\nb2 = %s", ValueStringer(b1), ValueStringer(b2)),
			opts...,
		)
	}
	return ok
}

// BytesNotEqual asserts that b1 and b2 are not equal.
// It uses [bytes.Equal] to compare the two byte slices.
//
//nolint:thelper // It's called below.
func BytesNotEqual(tb testing.TB, b1, b2 []byte, opts ...Option) bool {
	ok := !bytes.Equal(b1, b2)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"bytes_not_equal",
			fmt.Sprintf("equal:\nb1 = %s\nb2 = %s", ValueStringer(b1), ValueStringer(b2)),
			opts...,
		)
	}
	return ok
}
