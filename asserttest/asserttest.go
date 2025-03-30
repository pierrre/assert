// Package asserttest provides utilities to test assertions.
package asserttest

import (
	"testing"

	"github.com/pierrre/assert"
	"github.com/pierrre/assert/assertauto"
)

// NewReport returns a report function that checks the message and that the report function is called.
func NewReport(tb testing.TB, expectedMsg string) assert.ReportFunc {
	tb.Helper()
	return newReport(tb, func(msg string) {
		tb.Helper()
		assert.Equal(tb, msg, expectedMsg, assert.MessageWrap("report message"))
	})
}

// NewReportAuto returns a report function that checks the message and that the report function is called.
func NewReportAuto(tb testing.TB) assert.ReportFunc {
	tb.Helper()
	return newReport(tb, func(msg string) {
		tb.Helper()
		assertauto.Equal(tb, msg, assertauto.AssertOptions(assert.MessageWrap("report message")))
	})
}

// NewReportPrefix returns a report function that checks the message prefix and that the report function is called.
func NewReportPrefix(tb testing.TB, expectedMsgPrefix string) assert.ReportFunc {
	tb.Helper()
	return newReport(tb, func(msg string) {
		tb.Helper()
		assert.StringHasPrefix(tb, msg, expectedMsgPrefix, assert.MessageWrap("report message"))
	})
}

func newReport(tb testing.TB, checkMsg func(msg string)) assert.ReportFunc {
	tb.Helper()
	reportCalled := false
	report := func(tb testing.TB, actualArgs ...any) {
		tb.Helper()
		reportCalled = true
		assert.SliceLen(tb, actualArgs, 1)
		msg, _ := assert.Type[string](tb, actualArgs[0])
		checkMsg(msg)
	}
	tb.Cleanup(func() {
		tb.Helper()
		assert.True(tb, reportCalled, assert.Message("report not called"), assert.Report(testing.TB.Error))
	})
	return report
}
