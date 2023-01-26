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
	report := asserttest.NewReport(t, "assert deep_equal[int]: not equal:\ndiff = .: int not equal\nv1 = 1\nv2 = 2")
	ok := assert.DeepEqual(t, 1, 2, assert.Report(report))
	assert.False(t, ok)
}
