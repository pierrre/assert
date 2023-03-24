package assert_test

import (
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestAllocsPerRun(t *testing.T) {
	ok := AllocsPerRun(t, 10, func() {}, 0)
	True(t, ok)
}

func TestAllocsPerRunAlloc(t *testing.T) {
	ok := AllocsPerRun(t, 10, func() {
		_ = make([]byte, 1<<20)
	}, 1)
	True(t, ok)
}

func TestAllocsPerRunFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := AllocsPerRun(t, 10, func() {}, 1, Report(report))
	False(t, ok)
}
