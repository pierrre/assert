package assert_test

import (
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestStringEmpty(t *testing.T) {
	ok := StringEmpty(t, "")
	True(t, ok)
}

func TestStringEmptyFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert string_empty: not empty:\ns = \"abc\"")
	ok := StringEmpty(t, "abc", Report(report))
	False(t, ok)
}

func TestStringNotEmpty(t *testing.T) {
	ok := StringNotEmpty(t, "abc")
	True(t, ok)
}

func TestStringNotEmptyFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert string_not_empty: empty")
	ok := StringNotEmpty(t, "", Report(report))
	False(t, ok)
}

func TestStringLen(t *testing.T) {
	ok := StringLen(t, "abc", 3)
	True(t, ok)
}

func TestStringLenFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert string_len: unexpected length:\nexpected = 4\nactual = 3")
	ok := StringLen(t, "abc", 4, Report(report))
	False(t, ok)
}

func TestStringContains(t *testing.T) {
	ok := StringContains(t, "abc", "bc")
	True(t, ok)
}

func TestStringContainsFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert string_contains: not contains:\ns = \"abc\"\nsubstr = \"bd\"")
	ok := StringContains(t, "abc", "bd", Report(report))
	False(t, ok)
}

func TestStringNotContains(t *testing.T) {
	ok := StringNotContains(t, "abc", "bd")
	True(t, ok)
}

func TestStringNotContainsFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert string_not_contains: contains:\ns = \"abc\"\nsubstr = \"bc\"")
	ok := StringNotContains(t, "abc", "bc", Report(report))
	False(t, ok)
}

func TestStringHasPrefix(t *testing.T) {
	ok := StringHasPrefix(t, "abc", "ab")
	True(t, ok)
}

func TestStringHasPrefixFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert string_has_prefix: no prefix:\ns = \"abc\"\nprefix = \"ac\"")
	ok := StringHasPrefix(t, "abc", "ac", Report(report))
	False(t, ok)
}

func TestStringHasSuffix(t *testing.T) {
	ok := StringHasSuffix(t, "abc", "bc")
	True(t, ok)
}

func TestStringHasSuffixFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert string_has_suffix: no suffix:\ns = \"abc\"\nsuffix = \"ac\"")
	ok := StringHasSuffix(t, "abc", "ac", Report(report))
	False(t, ok)
}

func TestStringEqualFold(t *testing.T) {
	ok := StringEqualFold(t, "abc", "ABC")
	True(t, ok)
}

func TestStringEqualFoldFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert string_equal_fold: not equal fold:\ns1 = \"abc\"\ns2 = \"ABD\"")
	ok := StringEqualFold(t, "abc", "ABD", Report(report))
	False(t, ok)
}
