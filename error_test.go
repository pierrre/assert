package assert_test

import (
	"errors"
	"io"
	"testing"
	"time"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestError(t *testing.T) {
	err := errors.New("error")
	ok := Error(t, err)
	True(t, ok)
}

func TestErrorFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := Error(t, nil, report)
	False(t, ok)
}

func TestNoError(t *testing.T) {
	ok := NoError(t, nil)
	True(t, ok)
}

func TestNoErrorFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := NoError(t, errors.New("error"), report)
	False(t, ok)
}

func TestErrorIs(t *testing.T) {
	ok := ErrorIs(t, io.EOF, io.EOF)
	True(t, ok)
}

func TestErrorIsFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := ErrorIs(t, errors.New("error"), io.EOF, report)
	False(t, ok)
}

func TestErrorIsFailNil(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := ErrorIs(t, nil, io.EOF, report)
	False(t, ok)
}

func TestErrorNotIs(t *testing.T) {
	ok := ErrorNotIs(t, errors.New("error"), io.EOF)
	True(t, ok)
}

func TestErrorNotIsFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := ErrorNotIs(t, io.EOF, io.EOF, report)
	False(t, ok)
}

func TestErrorAs(t *testing.T) {
	var err error
	ok := ErrorAs(t, errors.New("error"), &err)
	True(t, ok)
}

func TestErrorAsFail(t *testing.T) {
	var timeParseError *time.ParseError
	report := asserttest.ReportAuto(t)
	ok := ErrorAs(t, errors.New("error"), &timeParseError, report)
	False(t, ok)
}

func TestErrorEqual(t *testing.T) {
	ok := ErrorEqual(t, errors.New("error"), "error")
	True(t, ok)
}

func TestErrorEqualFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := ErrorEqual(t, errors.New("error"), "zzz", report)
	False(t, ok)
}

func TestErrorEqualFailNil(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := ErrorEqual(t, nil, "zzz", report)
	False(t, ok)
}

func TestErrorContains(t *testing.T) {
	ok := ErrorContains(t, errors.New("aaa error bbb"), "error")
	True(t, ok)
}

func TestErrorContainsFail(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := ErrorContains(t, errors.New("error"), "zzz", report)
	False(t, ok)
}

func TestErrorContainsFailNil(t *testing.T) {
	report := asserttest.ReportAuto(t)
	ok := ErrorContains(t, nil, "zzz", report)
	False(t, ok)
}
