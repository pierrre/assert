package assert

import (
	"math"
	"testing"
)

// Float represents all floating-point numeric types.
type Float interface {
	~float32 | ~float64
}

// FloatInf asserts that f is an infinity according to sign.
// If sign > 0, asserts that f is positive infinity.
// If sign < 0, asserts that f is negative infinity.
// If sign == 0, asserts that f is either positive or negative infinity.
//
//nolint:thelper // It's called below.
func FloatInf[T Float](tb testing.TB, f T, sign int, opts ...Option) bool {
	ok := math.IsInf(float64(f), sign)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"float_inf",
			"not infinite:\nf = "+ValueStringer(f),
			1,
			opts...,
		)
	}
	return ok
}

// FloatNotInf asserts that f is not an infinity according to sign.
// If sign > 0, asserts that f is not positive infinity.
// If sign < 0, asserts that f is not negative infinity.
// If sign == 0, asserts that f is not any infinity (finite value).
//
//nolint:thelper // It's called below.
func FloatNotInf[T Float](tb testing.TB, f T, sign int, opts ...Option) bool {
	ok := !math.IsInf(float64(f), sign)
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"float_not_inf",
			"infinite:\nf = "+ValueStringer(f),
			1,
			opts...,
		)
	}
	return ok
}

// FloatNaN asserts that f is NaN (not-a-number).
//
//nolint:thelper // It's called below.
func FloatNaN[T Float](tb testing.TB, f T, opts ...Option) bool {
	ok := math.IsNaN(float64(f))
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"float_nan",
			"not NaN:\nf = "+ValueStringer(f),
			1,
			opts...,
		)
	}
	return ok
}

// FloatNotNaN asserts that f is not NaN (not-a-number).
//
//nolint:thelper // It's called below.
func FloatNotNaN[T Float](tb testing.TB, f T, opts ...Option) bool {
	ok := !math.IsNaN(float64(f))
	if !ok {
		tb.Helper()
		Fail(
			tb,
			"float_not_nan",
			"NaN:\nf = "+ValueStringer(f),
			1,
			opts...,
		)
	}
	return ok
}
