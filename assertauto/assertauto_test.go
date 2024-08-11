package assertauto

import (
	"testing"

	"github.com/pierrre/assert"
	"github.com/pierrre/pretty"
)

func init() {
	pretty.DefaultCommonValueWriter.ConfigureTest()
}

func TestEqual(t *testing.T) {
	ok := Equal(t, 123, Name("int"))
	assert.True(t, ok)
	ok = Equal(t, "foo", Name("string"))
	assert.True(t, ok)
}

func TestDeepEqual(t *testing.T) {
	ok := DeepEqual(t, &struct {
		Foo string
	}{
		Foo: "bar",
	})
	assert.True(t, ok)
}

func TestAllocsPerRun(t *testing.T) {
	ok := AllocsPerRun(t, 10, func() {
		_ = make([]byte, 1<<20)
	})
	assert.True(t, ok)
}
