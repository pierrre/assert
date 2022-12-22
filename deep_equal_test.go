package assert_test

import (
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestDeepEqual(t *testing.T) {
	ok := DeepEqual(t, 1, 1)
	True(t, ok)
}

func TestDeepEqualFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert deep_equal[int]: not equal:\ndiff = not implemented\nv1 = 1\nv2 = 2")
	ok := DeepEqual(t, 1, 2, Report(report))
	False(t, ok)
}

func TestNotDeepEqual(t *testing.T) {
	ok := NotDeepEqual(t, 1, 2)
	True(t, ok)
}

func TestNotDeepEqualFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert not_deep_equal[int]: equal:\nv1 = 1\nv2 = 1")
	ok := NotDeepEqual(t, 1, 1, Report(report))
	False(t, ok)
}
