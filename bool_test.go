package assert_test

import (
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestTrue(t *testing.T) {
	ok := True(t, true)
	True(t, ok)
}

func TestTrueFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := True(t, false, Report(report))
	False(t, ok)
}

func TestFalse(t *testing.T) {
	ok := False(t, false)
	True(t, ok)
}

func TestFalseFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := False(t, true, Report(report))
	False(t, ok)
}
