package assert_test

import (
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestEqual(t *testing.T) {
	ok := Equal(t, 123, 123)
	True(t, ok)
}

func TestEqualFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert equal[int]: not equal:\nv1 = 123\nv2 = 456")
	ok := Equal(t, 123, 456, Report(report))
	False(t, ok)
}

func TestNotEqual(t *testing.T) {
	ok := NotEqual(t, 123, 456)
	True(t, ok)
}

func TestNotEqualFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert not_equal[int]: equal:\nv1 = 123\nv2 = 123")
	ok := NotEqual(t, 123, 123, Report(report))
	False(t, ok)
}
