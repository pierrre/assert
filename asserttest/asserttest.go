// Package asserttest provides utilities to test assertions.
package asserttest

import (
	"testing"

	"github.com/pierrre/assert"
	"github.com/pierrre/assert/assertauto"
)

// ReportAuto returns an [assert.Option] that checks the message automatically and that the [assert.ReportFunc] is called.
func ReportAuto(tb testing.TB) assert.Option {
	tb.Helper()
	return assert.Report(newReportFunc(tb, func(msg string) {
		tb.Helper()
		assertauto.Equal(tb, msg, assertauto.AssertOptions(assert.MessageWrap("report message")))
	}))
}

// ReportPrefix returns an [assert.Option] that checks the message prefix and that the [assert.ReportFunc] is called.
func ReportPrefix(tb testing.TB, expectedMsgPrefix string) assert.Option {
	tb.Helper()
	return assert.Report(newReportFunc(tb, func(msg string) {
		tb.Helper()
		assert.StringHasPrefix(tb, msg, expectedMsgPrefix, assert.MessageWrap("report message"))
	}))
}

func newReportFunc(tb testing.TB, checkMsg func(msg string)) assert.ReportFunc {
	tb.Helper()
	called := 0
	report := func(tb testing.TB, actualArgs ...any) {
		tb.Helper()
		called++
		assert.SliceLen(tb, actualArgs, 1)
		msg, _ := assert.Type[string](tb, actualArgs[0])
		checkMsg(msg)
	}
	tb.Cleanup(func() {
		tb.Helper()
		assert.Equal(tb, called, 1, assert.Message("report must be called exactly once"), assert.ReportError())
	})
	return report
}
