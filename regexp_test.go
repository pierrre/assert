package assert_test

import (
	"regexp"
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestRegexp(t *testing.T) {
	ok := RegexpMatch(t, regexp.MustCompile("bc"), "abc")
	True(t, ok)
}

func TestRegexpFailMatch(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := RegexpMatch(t, "z", "abc", report)
	False(t, ok)
}

func TestRegexpFailCompile(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := RegexpMatch(t, "\\", "abc", report)
	False(t, ok)
}

func TestNotRegexp(t *testing.T) {
	ok := RegexpNotMatch(t, regexp.MustCompile("z"), "abc")
	True(t, ok)
}

func TestNotRegexpFailMatch(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := RegexpNotMatch(t, "bc", "abc", report)
	False(t, ok)
}

func TestNotRegexpFailCompile(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := RegexpNotMatch(t, "\\", "abc", report)
	False(t, ok)
}
