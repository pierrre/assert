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
	report := asserttest.NewReportAuto(t)
	ok := BytesEqual(t, []byte("abc"), []byte("abd"), Report(report))
	False(t, ok)
}

func TestBytesNotEqual(t *testing.T) {
	ok := BytesNotEqual(t, []byte("abc"), []byte("abd"))
	True(t, ok)
}

func TestBytesNotEqualFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := BytesNotEqual(t, []byte("abc"), []byte("abc"), Report(report))
	False(t, ok)
}
