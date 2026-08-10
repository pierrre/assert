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

func TestChanRecvNil(t *testing.T) {
	var c <-chan int
	ok := ChanRecvNil(t, c)
	True(t, ok)
}

func TestChanRecvNilFail(t *testing.T) {
	c := make(chan int)
	var rc <-chan int = c
	report := asserttest.ReportAuto(t)
	ok := ChanRecvNil(t, rc, report)
	False(t, ok)
}

func TestChanRecvNotNil(t *testing.T) {
	c := make(chan int)
	var rc <-chan int = c
	ok := ChanRecvNotNil(t, rc)
	True(t, ok)
}

func TestChanRecvNotNilFail(t *testing.T) {
	var c <-chan int
	report := asserttest.ReportAuto(t)
	ok := ChanRecvNotNil(t, c, report)
	False(t, ok)
}

func TestChanRecvEmpty(t *testing.T) {
	c := make(chan int)
	var rc <-chan int = c
	ok := ChanRecvEmpty(t, rc)
	True(t, ok)
}

func TestChanRecvEmptyFail(t *testing.T) {
	c := make(chan int, 10)
	c <- 1
	var rc <-chan int = c
	report := asserttest.ReportAuto(t)
	ok := ChanRecvEmpty(t, rc, report)
	False(t, ok)
}

func TestChanRecvNotEmpty(t *testing.T) {
	c := make(chan int, 10)
	c <- 1
	var rc <-chan int = c
	ok := ChanRecvNotEmpty(t, rc)
	True(t, ok)
}

func TestChanRecvNotEmptyFail(t *testing.T) {
	c := make(chan int)
	var rc <-chan int = c
	report := asserttest.ReportAuto(t)
	ok := ChanRecvNotEmpty(t, rc, report)
	False(t, ok)
}

func TestChanRecvLen(t *testing.T) {
	c := make(chan int, 10)
	c <- 1
	var rc <-chan int = c
	ok := ChanRecvLen(t, rc, 1)
	True(t, ok)
}

func TestChanRecvLenFail(t *testing.T) {
	c := make(chan int, 10)
	c <- 1
	var rc <-chan int = c
	report := asserttest.ReportAuto(t)
	ok := ChanRecvLen(t, rc, 2, report)
	False(t, ok)
}

func TestChanSendNil(t *testing.T) {
	var c chan<- int
	ok := ChanSendNil(t, c)
	True(t, ok)
}

func TestChanSendNilFail(t *testing.T) {
	c := make(chan int)
	var sc chan<- int = c
	report := asserttest.ReportAuto(t)
	ok := ChanSendNil(t, sc, report)
	False(t, ok)
}

func TestChanSendNotNil(t *testing.T) {
	c := make(chan int)
	var sc chan<- int = c
	ok := ChanSendNotNil(t, sc)
	True(t, ok)
}

func TestChanSendNotNilFail(t *testing.T) {
	var c chan<- int
	report := asserttest.ReportAuto(t)
	ok := ChanSendNotNil(t, c, report)
	False(t, ok)
}

func TestChanSendEmpty(t *testing.T) {
	c := make(chan int)
	var sc chan<- int = c
	ok := ChanSendEmpty(t, sc)
	True(t, ok)
}

func TestChanSendEmptyFail(t *testing.T) {
	c := make(chan int, 10)
	c <- 1
	var sc chan<- int = c
	report := asserttest.ReportAuto(t)
	ok := ChanSendEmpty(t, sc, report)
	False(t, ok)
}

func TestChanSendNotEmpty(t *testing.T) {
	c := make(chan int, 10)
	c <- 1
	var sc chan<- int = c
	ok := ChanSendNotEmpty(t, sc)
	True(t, ok)
}

func TestChanSendNotEmptyFail(t *testing.T) {
	c := make(chan int)
	var sc chan<- int = c
	report := asserttest.ReportAuto(t)
	ok := ChanSendNotEmpty(t, sc, report)
	False(t, ok)
}

func TestChanSendLen(t *testing.T) {
	c := make(chan int, 10)
	c <- 1
	var sc chan<- int = c
	ok := ChanSendLen(t, sc, 1)
	True(t, ok)
}

func TestChanSendLenFail(t *testing.T) {
	c := make(chan int, 10)
	c <- 1
	var sc chan<- int = c
	report := asserttest.ReportAuto(t)
	ok := ChanSendLen(t, sc, 2, report)
	False(t, ok)
}
