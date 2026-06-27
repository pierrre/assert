package assert_test

import (
	"math"
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestFloatInf(t *testing.T) {
	ok := FloatInf(t, math.Inf(1), 1)
	True(t, ok)
}

func TestFloatInfNegative(t *testing.T) {
	ok := FloatInf(t, math.Inf(-1), -1)
	True(t, ok)
}

func TestFloatInfAny(t *testing.T) {
	ok := FloatInf(t, math.Inf(1), 0)
	True(t, ok)
}

func TestFloatInfAnyNegative(t *testing.T) {
	ok := FloatInf(t, math.Inf(-1), 0)
	True(t, ok)
}

func TestFloatInfFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := FloatInf(t, 123.456, 0, report)
	False(t, ok)
}

func TestFloatInfFailNegative(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := FloatInf(t, math.Inf(1), -1, report)
	False(t, ok)
}

func TestFloatInfFailPositive(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := FloatInf(t, math.Inf(-1), 1, report)
	False(t, ok)
}

func TestFloatNotInf(t *testing.T) {
	ok := FloatNotInf(t, 123.456, 0)
	True(t, ok)
}

func TestFloatNotInfPositive(t *testing.T) {
	ok := FloatNotInf(t, 123.456, 1)
	True(t, ok)
}

func TestFloatNotInfNegative(t *testing.T) {
	ok := FloatNotInf(t, 123.456, -1)
	True(t, ok)
}

func TestFloatNotInfFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := FloatNotInf(t, math.Inf(1), 0, report)
	False(t, ok)
}

func TestFloatNotInfFailPositive(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := FloatNotInf(t, math.Inf(1), 1, report)
	False(t, ok)
}

func TestFloatNotInfFailNegative(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := FloatNotInf(t, math.Inf(-1), -1, report)
	False(t, ok)
}

func TestFloatNaN(t *testing.T) {
	ok := FloatNaN(t, math.NaN())
	True(t, ok)
}

func TestFloatNaNFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := FloatNaN(t, 123.456, report)
	False(t, ok)
}

func TestFloatNotNaN(t *testing.T) {
	ok := FloatNotNaN(t, 123.456)
	True(t, ok)
}

func TestFloatNotNaNFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := FloatNotNaN(t, math.NaN(), report)
	False(t, ok)
}

func TestFloatInf32(t *testing.T) {
	ok := FloatInf(t, float32(math.Inf(1)), 1)
	True(t, ok)
}

func TestFloatNotInf32(t *testing.T) {
	ok := FloatNotInf(t, float32(123.456), 0)
	True(t, ok)
}

func TestFloatNaN32(t *testing.T) {
	ok := FloatNaN(t, float32(math.NaN()))
	True(t, ok)
}

func TestFloatNotNaN32(t *testing.T) {
	ok := FloatNotNaN(t, float32(123.456))
	True(t, ok)
}
