package assert_test

import (
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestGreater(t *testing.T) {
	ok := Greater(t, 123, 0)
	True(t, ok)
}

func TestGreaterFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := Greater(t, 0, 123, Report(report))
	False(t, ok)
}

func TestGreaterOrEqual(t *testing.T) {
	ok := GreaterOrEqual(t, 123, 0)
	True(t, ok)
}

func TestGreaterOrEqualFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := GreaterOrEqual(t, 0, 123, Report(report))
	False(t, ok)
}

func TestLess(t *testing.T) {
	ok := Less(t, 0, 123)
	True(t, ok)
}

func TestLessFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := Less(t, 123, 0, Report(report))
	False(t, ok)
}

func TestLessOrEqual(t *testing.T) {
	ok := LessOrEqual(t, 0, 123)
	True(t, ok)
}

func TestLessOrEqualFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := LessOrEqual(t, 123, 0, Report(report))
	False(t, ok)
}
