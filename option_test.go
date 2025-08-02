package assert_test

import (
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestOptionAllocs(t *testing.T) {
	AllocsPerRun(t, 100, func() {
		Equal(t, 123, 123, Message("test"))
	}, 1)
}

func TestLazy(t *testing.T) {
	report := asserttest.ReportAuto(t)
	Fail(t, "test", "message", 0, report, Lazy(func() Option {
		return Message("custom")
	}))
}

func TestLazyAllocs(t *testing.T) {
	AllocsPerRun(t, 100, func() {
		Equal(t, 123, 123, Lazy(func() Option {
			return Message("test")
		}))
	}, 0)
}

func TestOptions(t *testing.T) {
	report := asserttest.ReportAuto(t)
	Fail(t, "test", "message", 0, Options(report, Message("custom")))
}

func TestMessage(t *testing.T) {
	report := asserttest.ReportAuto(t)
	Fail(t, "test", "message", 0, report, Message("custom"))
}

func TestMessagef(t *testing.T) {
	report := asserttest.ReportAuto(t)
	Fail(t, "test", "message", 0, report, Messagef("custom %d", 1))
}

func TestMessageWrap(t *testing.T) {
	report := asserttest.ReportAuto(t)
	Fail(t, "test", "message", 0, report, MessageWrap("custom"))
}

func TestMessageWrapf(t *testing.T) {
	report := asserttest.ReportAuto(t)
	Fail(t, "test", "message", 0, report, MessageWrapf("custom %d", 1))
}

func TestShowStack(t *testing.T) {
	Fail(t, "test", "message", 0, ShowStack(true), ReportLog())
}

func TestReport(t *testing.T) {
	Fail(t, "test", "message", 0, ReportFatal(), ReportError(), ReportSkip(), ReportLog())
}
