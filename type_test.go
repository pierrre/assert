package assert_test

import (
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestType(t *testing.T) {
	v, ok := Type[int](t, 123)
	True(t, ok)
	Equal(t, v, 123)
}

func TestTypeFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	_, ok := Type[int](t, "123", Report(report))
	False(t, ok)
}
