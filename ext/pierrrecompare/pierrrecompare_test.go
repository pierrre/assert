package pierrrecompare

import (
	"testing"

	"github.com/pierrre/assert"
	"github.com/pierrre/assert/asserttest"
)

func init() {
	Configure()
}

func Test(t *testing.T) {
	ok := assert.DeepEqual(t, 1, 1)
	assert.True(t, ok)
}

func TestFail(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := assert.DeepEqual(t, 1, 2, assert.Report(report))
	assert.False(t, ok)
}
