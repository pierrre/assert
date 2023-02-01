package assert_test

import (
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestMessage(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	Fail(t, "test", "message", Report(report), Message("custom"))
}

func TestMessagef(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	Fail(t, "test", "message", Report(report), Messagef("custom %d", 1))
}

func TestMessageWrap(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	Fail(t, "test", "message", Report(report), MessageWrap("custom"))
}

func TestMessageWrapf(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	Fail(t, "test", "message", Report(report), MessageWrapf("custom %d", 1))
}
