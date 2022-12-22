package assert_test

import (
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestChanNil(t *testing.T) {
	var c chan int
	ok := ChanNil(t, c)
	True(t, ok)
}

func TestChanNilFail(t *testing.T) {
	c := make(chan int)
	report := asserttest.NewReport(t, "assert chan_nil[int]: not nil")
	ok := ChanNil(t, c, Report(report))
	False(t, ok)
}

func TestChanNotNil(t *testing.T) {
	c := make(chan int)
	ok := ChanNotNil(t, c)
	True(t, ok)
}

func TestChanNotNilFail(t *testing.T) {
	var c chan int
	report := asserttest.NewReport(t, "assert chan_not_nil[int]: nil")
	ok := ChanNotNil(t, c, Report(report))
	False(t, ok)
}

func TestChanEmpty(t *testing.T) {
	c := make(chan int)
	ok := ChanEmpty(t, c)
	True(t, ok)
}

func TestChanEmptyFail(t *testing.T) {
	c := make(chan int, 10)
	c <- 1
	report := asserttest.NewReport(t, "assert chan_empty[int]: not empty:\nlength = 1")
	ok := ChanEmpty(t, c, Report(report))
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
	report := asserttest.NewReport(t, "assert chan_not_empty[int]: empty")
	ok := ChanNotEmpty(t, c, Report(report))
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
	report := asserttest.NewReport(t, "assert chan_len[int]: unexpected length:\nexpected = 2\nactual = 1")
	ok := ChanLen(t, c, 2, Report(report))
	False(t, ok)
}
