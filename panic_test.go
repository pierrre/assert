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
	report := asserttest.ReportAuto(t)
	_, ok := Panics(t, func() {}, report)
	False(t, ok)
}

func TestNotPanics(t *testing.T) {
	ok := NotPanics(t, func() {})
	True(t, ok)
}

func TestNotPanicsFail(t *testing.T) {
	report := asserttest.ReportPrefix(t, "assert not_panics: panic:\npanic = [string] (len=4) \"test\"\nstack = ")
	ok := NotPanics(t, func() {
		panic("test")
	}, report)
	False(t, ok)
}
