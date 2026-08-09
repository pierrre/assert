package assert_test

import (
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
	"github.com/pierrre/go-libs/raceutil"
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
	if raceutil.Enabled {
		t.Skip("allocs are not measured under -race")
	}
	report := asserttest.ReportAuto(t)
	ok := AllocsPerRun(t, 10, func() {}, 1, report)
	False(t, ok)
}

func TestAllocsPerRunRace(t *testing.T) {
	if !raceutil.Enabled {
		t.Skip("allocs check is only skipped under -race")
	}
	lct := &logCaptureTB{T: t}
	ok := AllocsPerRun(lct, 10, func() {}, 1)
	True(t, ok)
	True(t, lct.logged)
}

type logCaptureTB struct {
	*testing.T
	logged bool
}

func (t *logCaptureTB) Log(args ...any) {
	t.logged = true
	t.T.Log(args...)
}

func (t *logCaptureTB) Logf(format string, args ...any) {
	t.logged = true
	t.T.Logf(format, args...)
}
