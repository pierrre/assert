package assert_test

import (
	"testing"

	. "github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func TestCondition(t *testing.T) {
	ok := Condition(t, func() bool {
		return true
	})
	True(t, ok)
}

func TestConditionFail(t *testing.T) {
	report := asserttest.NewReport(t, "assert condition: failed")
	Condition(t, func() bool {
		return false
	}, Report(report))
}
