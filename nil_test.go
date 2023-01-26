package assert_test

import (
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestNil(t *testing.T) {
	ok := Nil(t, nil)
	Equal(t, true, ok)
}

func TestNilSlice(t *testing.T) {
	ok := Nil(t, []int(nil))
	Equal(t, true, ok)
}

func TestNilFailSlice(t *testing.T) {
	report := asserttest.NewReport(t, "assert nil: not nil:\nv = []")
	ok := Nil(t, []int{}, Report(report))
	Equal(t, false, ok)
}

func TestNilFailString(t *testing.T) {
	report := asserttest.NewReport(t, "assert nil: not nil:\nv = \"foo\"")
	ok := Nil(t, "foo", Report(report))
	Equal(t, false, ok)
}

func TestNotNilSlice(t *testing.T) {
	ok := NotNil(t, []int{})
	Equal(t, true, ok)
}

func TestNotNilString(t *testing.T) {
	ok := NotNil(t, "foo")
	Equal(t, true, ok)
}

func TestNotNilFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert not_nil: nil")
	ok := NotNil(t, nil, Report(report))
	Equal(t, false, ok)
}

func TestNotNilFailSlice(t *testing.T) {
	report := asserttest.NewReport(t, "assert not_nil: nil")
	ok := NotNil(t, []int(nil), Report(report))
	Equal(t, false, ok)
}
