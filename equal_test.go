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
	report := asserttest.ReportAuto(t)
	ok := Equal(t, 123, 456, report)
	False(t, ok)
}

func TestNotEqual(t *testing.T) {
	ok := NotEqual(t, 123, 456)
	True(t, ok)
}

func TestNotEqualFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := NotEqual(t, 123, 123, report)
	False(t, ok)
}
