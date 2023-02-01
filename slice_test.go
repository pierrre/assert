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
	report := asserttest.NewReportAuto(t)
	ok := SliceNil(t, []int{}, Report(report))
	False(t, ok)
}

func TestSliceNotNil(t *testing.T) {
	ok := SliceNotNil(t, []int{})
	True(t, ok)
}

func TestSliceNotNilFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := SliceNotNil(t, []int(nil), Report(report))
	False(t, ok)
}

func TestSliceEmpty(t *testing.T) {
	ok := SliceEmpty(t, []int{})
	True(t, ok)
}

func TestSliceEmptyFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := SliceEmpty(t, []int{1}, Report(report))
	False(t, ok)
}

func TestSliceNotEmpty(t *testing.T) {
	ok := SliceNotEmpty(t, []int{1})
	True(t, ok)
}

func TestSliceNotEmptyFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := SliceNotEmpty(t, []int{}, Report(report))
	False(t, ok)
}

func TestSliceLen(t *testing.T) {
	ok := SliceLen(t, []int{1}, 1)
	True(t, ok)
}

func TestSliceLenFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := SliceLen(t, []int{1}, 2, Report(report))
	False(t, ok)
}

func TestSliceEqual(t *testing.T) {
	ok := SliceEqual(t, []int{1}, []int{1})
	True(t, ok)
}

func TestSliceEqualFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := SliceEqual(t, []int{1}, []int{2}, Report(report))
	False(t, ok)
}

func TestSliceNotEqual(t *testing.T) {
	ok := SliceNotEqual(t, []int{1}, []int{2})
	True(t, ok)
}

func TestSliceNotEqualFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := SliceNotEqual(t, []int{1}, []int{1}, Report(report))
	False(t, ok)
}

func TestSliceContains(t *testing.T) {
	ok := SliceContains(t, []int{1, 2}, 1)
	True(t, ok)
}

func TestSliceContainsFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := SliceContains(t, []int{1, 2}, 3, Report(report))
	False(t, ok)
}

func TestSliceNotContains(t *testing.T) {
	ok := SliceNotContains(t, []int{1, 2}, 3)
	True(t, ok)
}

func TestSliceNotContainsFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := SliceNotContains(t, []int{1, 2}, 1, Report(report))
	False(t, ok)
}

func TestSliceContainsAll(t *testing.T) {
	ok := SliceContainsAll(t, []int{1, 2}, []int{1})
	True(t, ok)
}

func TestSliceContainsAllFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := SliceContainsAll(t, []int{1, 2}, []int{3}, Report(report))
	False(t, ok)
}

func TestSliceNotContainsAll(t *testing.T) {
	ok := SliceNotContainsAll(t, []int{1, 2}, []int{3})
	True(t, ok)
}

func TestSliceNotContainsAllFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := SliceNotContainsAll(t, []int{1, 2}, []int{1}, Report(report))
	False(t, ok)
}
