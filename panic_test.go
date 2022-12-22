package assert_test

import (
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestPanics(t *testing.T) {
	rec, ok := Panics(t, func() {
		panic("test")
	})
	True(t, ok)
	recStr, _ := Type[string](t, rec)
	Equal(t, recStr, "test")
}

func TestPanicsFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert panics: no panic")
	_, ok := Panics(t, func() {}, Report(report))
	False(t, ok)
}

func TestNotPanics(t *testing.T) {
	ok := NotPanics(t, func() {})
	True(t, ok)
}

func TestNotPanicsFail(t *testing.T) {
	report := asserttest.NewReportPrefix(t, "assert not_panics: panic:\npanic = \"test\"\nstack = ")
	ok := NotPanics(t, func() {
		panic("test")
	}, Report(report))
	False(t, ok)
}
