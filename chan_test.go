package assert_test

import (
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestChanEmpty(t *testing.T) {
	c := make(chan int)
	ok := ChanEmpty(t, c)
	True(t, ok)
}

func TestChanEmptyFail(t *testing.T) {
	c := make(chan int, 10)
	c <- 1
	report := asserttest.ReportAuto(t)
	ok := ChanEmpty(t, c, report)
	False(t, ok)
}

func TestChanNotEmpty(t *testing.T) {
	c := make(chan int, 10)
	c <- 1
	ok := ChanNotEmpty(t, c)
	True(t, ok)
}

func TestChanNotEmptyFail(t *testing.T) {
	c := make(chan int)
	report := asserttest.ReportAuto(t)
	ok := ChanNotEmpty(t, c, report)
	False(t, ok)
}

func TestChanLen(t *testing.T) {
	c := make(chan int, 10)
	c <- 1
	ok := ChanLen(t, c, 1)
	True(t, ok)
}

func TestChanLenFail(t *testing.T) {
	c := make(chan int, 10)
	c <- 1
	report := asserttest.ReportAuto(t)
	ok := ChanLen(t, c, 2, report)
	False(t, ok)
}

func TestChanNil(t *testing.T) {
	var c chan int
	ok := ChanNil(t, c)
	True(t, ok)
}

func TestChanNilFail(t *testing.T) {
	c := make(chan int)
	report := asserttest.ReportAuto(t)
	ok := ChanNil(t, c, report)
	False(t, ok)
}

func TestChanNotNil(t *testing.T) {
	c := make(chan int)
	ok := ChanNotNil(t, c)
	True(t, ok)
}

func TestChanNotNilFail(t *testing.T) {
	var c chan int
	report := asserttest.ReportAuto(t)
	ok := ChanNotNil(t, c, report)
	False(t, ok)
}
