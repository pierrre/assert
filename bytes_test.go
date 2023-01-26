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
	report := asserttest.NewReport(t, "assert bytes_equal: not equal:\nb1 = [97 98 99]\nb2 = [97 98 100]")
	ok := BytesEqual(t, []byte("abc"), []byte("abd"), Report(report))
	False(t, ok)
}

func TestBytesNotEqual(t *testing.T) {
	ok := BytesNotEqual(t, []byte("abc"), []byte("abd"))
	True(t, ok)
}

func TestBytesNotEqualFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert bytes_not_equal: equal:\nb1 = [97 98 99]\nb2 = [97 98 99]")
	ok := BytesNotEqual(t, []byte("abc"), []byte("abc"), Report(report))
	False(t, ok)
}
