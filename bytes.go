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
			1,
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
			1,
			opts...,
		)
	}
	return ok
}

// BytesContains asserts that b contains subslice.
// It uses [bytes.Contains] to check if subslice is contained in b.
//
//nolint:thelper // It's called below.
func BytesContains(tb testing.TB, b, subslice []byte, opts ...Option) bool {
	ok := bytes.Contains(b, subslice)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"bytes_contains",
			fmt.Sprintf("not contains:\nb = %s\nsubslice = %s", ValueStringer(b), ValueStringer(subslice)),
			1,
			opts...,
		)
	}
	return ok
}

// BytesNotContains asserts that b does not contain subslice.
// It uses [bytes.Contains] to check if subslice is contained in b.
//
//nolint:thelper // It's called below.
func BytesNotContains(tb testing.TB, b, subslice []byte, opts ...Option) bool {
	ok := !bytes.Contains(b, subslice)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"bytes_not_contains",
			fmt.Sprintf("contains:\nb = %s\nsubslice = %s", ValueStringer(b), ValueStringer(subslice)),
			1,
			opts...,
		)
	}
	return ok
}

// BytesHasPrefix asserts that b begins with prefix.
// It uses [bytes.HasPrefix] to check if b begins with prefix.
//
//nolint:thelper // It's called below.
func BytesHasPrefix(tb testing.TB, b, prefix []byte, opts ...Option) bool {
	ok := bytes.HasPrefix(b, prefix)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"bytes_has_prefix",
			fmt.Sprintf("no prefix:\nb = %s\nprefix = %s", ValueStringer(b), ValueStringer(prefix)),
			1,
			opts...,
		)
	}
	return ok
}

// BytesNotHasPrefix asserts that b does not begin with prefix.
// It uses [bytes.HasPrefix] to check if b begins with prefix.
//
//nolint:thelper // It's called below.
func BytesNotHasPrefix(tb testing.TB, b, prefix []byte, opts ...Option) bool {
	ok := !bytes.HasPrefix(b, prefix)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"bytes_not_has_prefix",
			fmt.Sprintf("has prefix:\nb = %s\nprefix = %s", ValueStringer(b), ValueStringer(prefix)),
			1,
			opts...,
		)
	}
	return ok
}

// BytesHasSuffix asserts that b ends with suffix.
// It uses [bytes.HasSuffix] to check if b ends with suffix.
//
//nolint:thelper // It's called below.
func BytesHasSuffix(tb testing.TB, b, suffix []byte, opts ...Option) bool {
	ok := bytes.HasSuffix(b, suffix)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"bytes_has_suffix",
			fmt.Sprintf("no suffix:\nb = %s\nsuffix = %s", ValueStringer(b), ValueStringer(suffix)),
			1,
			opts...,
		)
	}
	return ok
}

// BytesNotHasSuffix asserts that b does not end with suffix.
// It uses [bytes.HasSuffix] to check if b ends with suffix.
//
//nolint:thelper // It's called below.
func BytesNotHasSuffix(tb testing.TB, b, suffix []byte, opts ...Option) bool {
	ok := !bytes.HasSuffix(b, suffix)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"bytes_not_has_suffix",
			fmt.Sprintf("has suffix:\nb = %s\nsuffix = %s", ValueStringer(b), ValueStringer(suffix)),
			1,
			opts...,
		)
	}
	return ok
}
