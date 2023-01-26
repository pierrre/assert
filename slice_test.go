package assert_test

import (
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestSliceNil(t *testing.T) {
	ok := SliceNil(t, []int(nil))
	True(t, ok)
}

func TestSliceNilFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert slice_nil[int]: not nil:\ns = []")
	ok := SliceNil(t, []int{}, Report(report))
	False(t, ok)
}

func TestSliceNotNil(t *testing.T) {
	ok := SliceNotNil(t, []int{})
	True(t, ok)
}

func TestSliceNotNilFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert slice_not_nil[int]: nil")
	ok := SliceNotNil(t, []int(nil), Report(report))
	False(t, ok)
}

func TestSliceEmpty(t *testing.T) {
	ok := SliceEmpty(t, []int{})
	True(t, ok)
}

func TestSliceEmptyFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert slice_empty[int]: not empty:\nlength = 1\ns = [1]")
	ok := SliceEmpty(t, []int{1}, Report(report))
	False(t, ok)
}

func TestSliceNotEmpty(t *testing.T) {
	ok := SliceNotEmpty(t, []int{1})
	True(t, ok)
}

func TestSliceNotEmptyFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert slice_not_empty[int]: empty")
	ok := SliceNotEmpty(t, []int{}, Report(report))
	False(t, ok)
}

func TestSliceLen(t *testing.T) {
	ok := SliceLen(t, []int{1}, 1)
	True(t, ok)
}

func TestSliceLenFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert slice_len[int]: unexpected length:\nexpected = 2\nactual = 1")
	ok := SliceLen(t, []int{1}, 2, Report(report))
	False(t, ok)
}

func TestSliceEqual(t *testing.T) {
	ok := SliceEqual(t, []int{1}, []int{1})
	True(t, ok)
}

func TestSliceEqualFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert slice_equal[int]: not equal:\ns1 = [1]\ns2 = [2]")
	ok := SliceEqual(t, []int{1}, []int{2}, Report(report))
	False(t, ok)
}

func TestSliceNotEqual(t *testing.T) {
	ok := SliceNotEqual(t, []int{1}, []int{2})
	True(t, ok)
}

func TestSliceNotEqualFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert slice_not_equal[int]: equal:\ns1 = [1]\ns2 = [1]")
	ok := SliceNotEqual(t, []int{1}, []int{1}, Report(report))
	False(t, ok)
}

func TestSliceContains(t *testing.T) {
	ok := SliceContains(t, []int{1, 2}, 1)
	True(t, ok)
}

func TestSliceContainsFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert slice_contains[int]: not contains:\ns = [1 2]\nv = 3")
	ok := SliceContains(t, []int{1, 2}, 3, Report(report))
	False(t, ok)
}

func TestSliceNotContains(t *testing.T) {
	ok := SliceNotContains(t, []int{1, 2}, 3)
	True(t, ok)
}

func TestSliceNotContainsFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert slice_not_contains[int]: contains:\ns = [1 2]\nv = 1")
	ok := SliceNotContains(t, []int{1, 2}, 1, Report(report))
	False(t, ok)
}

func TestSliceContainsAll(t *testing.T) {
	ok := SliceContainsAll(t, []int{1, 2}, []int{1})
	True(t, ok)
}

func TestSliceContainsAllFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert slice_contains_all[int]: not contains all:\ns1 = [1 2]\ns2 = [3]")
	ok := SliceContainsAll(t, []int{1, 2}, []int{3}, Report(report))
	False(t, ok)
}

func TestSliceNotContainsAll(t *testing.T) {
	ok := SliceNotContainsAll(t, []int{1, 2}, []int{3})
	True(t, ok)
}

func TestSliceNotContainsAllFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert slice_not_contains_all[int]: contains all:\ns1 = [1 2]\ns2 = [1]")
	ok := SliceNotContainsAll(t, []int{1, 2}, []int{1}, Report(report))
	False(t, ok)
}
