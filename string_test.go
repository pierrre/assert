package assert_test

import (
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestStringLen(t *testing.T) {
	ok := StringLen(t, "abc", 3)
	True(t, ok)
}

func TestStringLenFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := StringLen(t, "abc", 4, Report(report))
	False(t, ok)
}

func TestStringContains(t *testing.T) {
	ok := StringContains(t, "abc", "bc")
	True(t, ok)
}

func TestStringContainsFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := StringContains(t, "abc", "bd", Report(report))
	False(t, ok)
}

func TestStringNotContains(t *testing.T) {
	ok := StringNotContains(t, "abc", "bd")
	True(t, ok)
}

func TestStringNotContainsFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := StringNotContains(t, "abc", "bc", Report(report))
	False(t, ok)
}

func TestStringHasPrefix(t *testing.T) {
	ok := StringHasPrefix(t, "abc", "ab")
	True(t, ok)
}

func TestStringHasPrefixFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := StringHasPrefix(t, "abc", "ac", Report(report))
	False(t, ok)
}

func TestStringHasSuffix(t *testing.T) {
	ok := StringHasSuffix(t, "abc", "bc")
	True(t, ok)
}

func TestStringHasSuffixFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := StringHasSuffix(t, "abc", "ac", Report(report))
	False(t, ok)
}

func TestStringEqualFold(t *testing.T) {
	ok := StringEqualFold(t, "abc", "ABC")
	True(t, ok)
}

func TestStringEqualFoldFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := StringEqualFold(t, "abc", "ABD", Report(report))
	False(t, ok)
}
