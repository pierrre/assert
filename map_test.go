package assert_test

import (
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestMapNil(t *testing.T) {
	ok := MapNil(t, map[string]string(nil))
	True(t, ok)
}

func TestMapNilFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := MapNil(t, map[string]string{}, report)
	False(t, ok)
}

func TestMapNotNil(t *testing.T) {
	ok := MapNotNil(t, map[string]string{})
	True(t, ok)
}

func TestMapNotNilFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := MapNotNil(t, map[string]string(nil), report)
	False(t, ok)
}

func TestMapEmpty(t *testing.T) {
	ok := MapEmpty(t, map[string]string{})
	True(t, ok)
}

func TestMapEmptyFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := MapEmpty(t, map[string]string{"foo": "bar"}, report)
	False(t, ok)
}

func TestMapNotEmpty(t *testing.T) {
	ok := MapNotEmpty(t, map[string]string{"foo": "bar"})
	True(t, ok)
}

func TestMapNotEmptyFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := MapNotEmpty(t, map[string]string{}, report)
	False(t, ok)
}

func TestMapLen(t *testing.T) {
	ok := MapLen(t, map[string]string{"foo": "bar"}, 1)
	True(t, ok)
}

func TestMapLenFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := MapLen(t, map[string]string{"foo": "bar"}, 2, report)
	False(t, ok)
}

func TestMapEqual(t *testing.T) {
	ok := MapEqual(t, map[string]string{"foo": "bar"}, map[string]string{"foo": "bar"})
	True(t, ok)
}

func TestMapEqualFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := MapEqual(t, map[string]string{"foo": "bar"}, map[string]string{}, report)
	False(t, ok)
}

func TestMapNotEqual(t *testing.T) {
	ok := MapNotEqual(t, map[string]string{"foo": "bar"}, map[string]string{})
	True(t, ok)
}

func TestMapNotEqualFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := MapNotEqual(t, map[string]string{"foo": "bar"}, map[string]string{"foo": "bar"}, report)
	False(t, ok)
}
