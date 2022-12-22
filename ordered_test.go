package assert_test

import (
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestGreater(t *testing.T) {
	ok := Greater(t, 123, 0)
	True(t, ok)
}

func TestGreaterFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert greater[int]: not greater than:\nv1 = 0\nv2 = 123")
	ok := Greater(t, 0, 123, Report(report))
	False(t, ok)
}

func TestGreaterOrEqual(t *testing.T) {
	ok := GreaterOrEqual(t, 123, 0)
	True(t, ok)
}

func TestGreaterOrEqualFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert greater_or_equal[int]: not greater than or equal to:\nv1 = 0\nv2 = 123")
	ok := GreaterOrEqual(t, 0, 123, Report(report))
	False(t, ok)
}

func TestLess(t *testing.T) {
	ok := Less(t, 0, 123)
	True(t, ok)
}

func TestLessFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert less[int]: not less than:\nv1 = 123\nv2 = 0")
	ok := Less(t, 123, 0, Report(report))
	False(t, ok)
}

func TestLessOrEqual(t *testing.T) {
	ok := LessOrEqual(t, 0, 123)
	True(t, ok)
}

func TestLessOrEqualFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert less_or_equal[int]: not less than or equal to:\nv1 = 123\nv2 = 0")
	ok := LessOrEqual(t, 123, 0, Report(report))
	False(t, ok)
}
