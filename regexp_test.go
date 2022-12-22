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
	report := asserttest.NewReport(t, "assert regexp_match: no match:\nrs = \"z\"\ns = \"abc\"")
	ok := RegexpMatch(t, "z", "abc", Report(report))
	False(t, ok)
}

func TestRegexpFailCompile(t *testing.T) {
	report := asserttest.NewReport(t, "assert regexp_compile: compilation failed:\nexpr = \"\\\\\"\nerr = error parsing regexp: trailing backslash at end of expression: ``")
	ok := RegexpMatch(t, "\\", "abc", Report(report))
	False(t, ok)
}

func TestNotRegexp(t *testing.T) {
	ok := RegexpNotMatch(t, regexp.MustCompile("z"), "abc")
	True(t, ok)
}

func TestNotRegexpFailMatch(t *testing.T) {
	report := asserttest.NewReport(t, "assert regexp_not_match: match:\nrs = \"bc\"\ns = \"abc\"")
	ok := RegexpNotMatch(t, "bc", "abc", Report(report))
	False(t, ok)
}

func TestNotRegexpFailCompile(t *testing.T) {
	report := asserttest.NewReport(t, "assert regexp_compile: compilation failed:\nexpr = \"\\\\\"\nerr = error parsing regexp: trailing backslash at end of expression: ``")
	ok := RegexpNotMatch(t, "\\", "abc", Report(report))
	False(t, ok)
}
