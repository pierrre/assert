package assertauto

import (
	"testing"
)

func TestEqual(t *testing.T) {
	Equal(t, "aaaaaa")
	Equal(t, 123456)
	Equal(t, 123.456)
	Equal(t, map[string]interface{}{"a": 1, "b": 2, "c": 3})
}

func TestAllocsPerRun(t *testing.T) {
	AllocsPerRun(t, 100, func() {
		_ = make([]byte, 1<<20)
	})
}
