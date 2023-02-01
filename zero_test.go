package assert_test

import (
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestZero(t *testing.T) {
	ok := Zero(t, 0)
	True(t, ok)
}

func TestZeroFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := Zero(t, 123, Report(report))
	False(t, ok)
}

func TestNotZero(t *testing.T) {
	ok := NotZero(t, 123)
	True(t, ok)
}

func TestNotZeroFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := NotZero(t, 0, Report(report))
	False(t, ok)
}
