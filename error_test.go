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
	report := asserttest.NewReport(t, "assert error: no error")
	ok := Error(t, nil, Report(report))
	False(t, ok)
}

func TestNoError(t *testing.T) {
	ok := NoError(t, nil)
	True(t, ok)
}

func TestNoErrorFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert no_error: error: error")
	ok := NoError(t, errors.New("error"), Report(report))
	False(t, ok)
}

func TestErrorIs(t *testing.T) {
	ok := ErrorIs(t, io.EOF, io.EOF)
	True(t, ok)
}

func TestErrorIsFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert error_is: no match:\nerr = error\ntarget = EOF")
	ok := ErrorIs(t, errors.New("error"), io.EOF, Report(report))
	False(t, ok)
}

func TestErrorNotIs(t *testing.T) {
	ok := ErrorNotIs(t, errors.New("error"), io.EOF)
	True(t, ok)
}

func TestErrorNotIsFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert error_not_is: match:\nerr = EOF\ntarget = EOF")
	ok := ErrorNotIs(t, io.EOF, io.EOF, Report(report))
	False(t, ok)
}

func TestErrorAs(t *testing.T) {
	var err error
	ok := ErrorAs(t, errors.New("error"), &err)
	True(t, ok)
}

func TestErrorAsFail(t *testing.T) {
	var timeParseError *time.ParseError
	report := asserttest.NewReport(t, "assert error_as: no match:\nerr = error\ntarget = **time.ParseError")
	ok := ErrorAs(t, errors.New("error"), &timeParseError, Report(report))
	False(t, ok)
}

func TestErrorEqual(t *testing.T) {
	ok := ErrorEqual(t, errors.New("error"), "error")
	True(t, ok)
}

func TestErrorEqualFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert error_equal: not equal:\nerr = error\nmessage = \"zzz\"")
	ok := ErrorEqual(t, errors.New("error"), "zzz", Report(report))
	False(t, ok)
}

func TestErrorContains(t *testing.T) {
	ok := ErrorContains(t, errors.New("aaa error bbb"), "error")
	True(t, ok)
}

func TestErrorContainsFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert error_contains: not contains:\nerr = error\nsubstr = \"zzz\"")
	ok := ErrorContains(t, errors.New("error"), "zzz", Report(report))
	False(t, ok)
}
