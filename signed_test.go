package assert_test

import (
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestPositive(t *testing.T) {
	ok := Positive(t, 123)
	True(t, ok)
}

func TestPositiveFailNegative(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := Positive(t, -123, Report(report))
	False(t, ok)
}

func TestPositiveFailZero(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := Positive(t, 0, Report(report))
	False(t, ok)
}

func TestNegative(t *testing.T) {
	ok := Negative(t, -123)
	True(t, ok)
}

func TestNegativeFailPositive(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := Negative(t, 123, Report(report))
	False(t, ok)
}

func TestNegativeFailZero(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := Negative(t, 0, Report(report))
	False(t, ok)
}
