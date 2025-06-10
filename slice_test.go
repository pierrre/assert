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
	report := asserttest.ReportAuto(t)
	ok := SliceNil(t, []int{}, report)
	False(t, ok)
}

func TestSliceNotNil(t *testing.T) {
	ok := SliceNotNil(t, []int{})
	True(t, ok)
}

func TestSliceNotNilFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := SliceNotNil(t, []int(nil), report)
	False(t, ok)
}

func TestSliceEmpty(t *testing.T) {
	ok := SliceEmpty(t, []int{})
	True(t, ok)
}

func TestSliceEmptyFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := SliceEmpty(t, []int{1}, report)
	False(t, ok)
}

func TestSliceNotEmpty(t *testing.T) {
	ok := SliceNotEmpty(t, []int{1})
	True(t, ok)
}

func TestSliceNotEmptyFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := SliceNotEmpty(t, []int{}, report)
	False(t, ok)
}

func TestSliceLen(t *testing.T) {
	ok := SliceLen(t, []int{1}, 1)
	True(t, ok)
}

func TestSliceLenFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := SliceLen(t, []int{1}, 2, report)
	False(t, ok)
}

func TestSliceEqual(t *testing.T) {
	ok := SliceEqual(t, []int{1}, []int{1})
	True(t, ok)
}

func TestSliceEqualFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := SliceEqual(t, []int{1}, []int{2}, report)
	False(t, ok)
}

func TestSliceNotEqual(t *testing.T) {
	ok := SliceNotEqual(t, []int{1}, []int{2})
	True(t, ok)
}

func TestSliceNotEqualFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := SliceNotEqual(t, []int{1}, []int{1}, report)
	False(t, ok)
}

func TestSliceContains(t *testing.T) {
	ok := SliceContains(t, []int{1, 2}, 1)
	True(t, ok)
}

func TestSliceContainsFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := SliceContains(t, []int{1, 2}, 3, report)
	False(t, ok)
}

func TestSliceNotContains(t *testing.T) {
	ok := SliceNotContains(t, []int{1, 2}, 3)
	True(t, ok)
}

func TestSliceNotContainsFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := SliceNotContains(t, []int{1, 2}, 1, report)
	False(t, ok)
}
