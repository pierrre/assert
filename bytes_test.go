package assert_test

import (
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestBytesEqual(t *testing.T) {
	ok := BytesEqual(t, []byte("abc"), []byte("abc"))
	True(t, ok)
}

func TestBytesEqualFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := BytesEqual(t, []byte("abc"), []byte("abd"), report)
	False(t, ok)
}

func TestBytesNotEqual(t *testing.T) {
	ok := BytesNotEqual(t, []byte("abc"), []byte("abd"))
	True(t, ok)
}

func TestBytesNotEqualFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := BytesNotEqual(t, []byte("abc"), []byte("abc"), report)
	False(t, ok)
}

func TestBytesContains(t *testing.T) {
	ok := BytesContains(t, []byte("abc"), []byte("b"))
	True(t, ok)
}

func TestBytesContainsFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := BytesContains(t, []byte("abc"), []byte("d"), report)
	False(t, ok)
}

func TestBytesNotContains(t *testing.T) {
	ok := BytesNotContains(t, []byte("abc"), []byte("d"))
	True(t, ok)
}

func TestBytesNotContainsFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := BytesNotContains(t, []byte("abc"), []byte("b"), report)
	False(t, ok)
}

func TestBytesHasPrefix(t *testing.T) {
	ok := BytesHasPrefix(t, []byte("abc"), []byte("ab"))
	True(t, ok)
}

func TestBytesHasPrefixFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := BytesHasPrefix(t, []byte("abc"), []byte("ac"), report)
	False(t, ok)
}

func TestBytesNotHasPrefix(t *testing.T) {
	ok := BytesNotHasPrefix(t, []byte("abc"), []byte("ac"))
	True(t, ok)
}

func TestBytesNotHasPrefixFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := BytesNotHasPrefix(t, []byte("abc"), []byte("ab"), report)
	False(t, ok)
}

func TestBytesHasSuffix(t *testing.T) {
	ok := BytesHasSuffix(t, []byte("abc"), []byte("bc"))
	True(t, ok)
}

func TestBytesHasSuffixFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := BytesHasSuffix(t, []byte("abc"), []byte("ac"), report)
	False(t, ok)
}

func TestBytesNotHasSuffix(t *testing.T) {
	ok := BytesNotHasSuffix(t, []byte("abc"), []byte("ac"))
	True(t, ok)
}

func TestBytesNotHasSuffixFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := BytesNotHasSuffix(t, []byte("abc"), []byte("bc"), report)
	False(t, ok)
}
