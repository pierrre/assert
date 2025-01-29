package assertauto_test

import (
	"path/filepath"
	"testing"

	"github.com/pierrre/assert"
	. "github.com/pierrre/assert/assertauto"
	"github.com/pierrre/assert/asserttest"
)

func Test(t *testing.T) {
	Equal(t, nil)
	Equal(t, true)
	Equal(t, 123456)
	Equal(t, 123.456)
	Equal(t, "test")
}

func TestEqual(t *testing.T) {
	tmpDir := t.TempDir()
	fp := filepath.Join(tmpDir, "test.txt")
	t.Run("Save", func(t *testing.T) {
		ok := Equal(t, "test", FilePath(fp), Update(true))
		assert.True(t, ok)
	})
	t.Run("Load", func(t *testing.T) {
		ok := Equal(t, "test", FilePath(fp), Update(false))
		assert.True(t, ok)
	})
}

func TestEqualFail(t *testing.T) {
	tmpDir := t.TempDir()
	fp := filepath.Join(tmpDir, "test.txt")
	t.Run("Save", func(t *testing.T) {
		ok := Equal(t, "aaa", FilePath(fp), Update(true))
		assert.True(t, ok)
	})
	report := asserttest.NewReportAuto(t)
	t.Run("Load", func(t *testing.T) {
		ok := Equal(t, "bbb", FilePath(fp), Update(false), AssertOptions(assert.Report(report)))
		assert.False(t, ok)
	})
}

func TestAllocsPerRun(t *testing.T) {
	allocs, ok := AllocsPerRun(t, 100, func() {
		_ = make([]byte, 1<<20)
	})
	assert.True(t, ok)
	Equal(t, allocs)
}

func TestErrorContainsSeparator(t *testing.T) {
	report := asserttest.NewReportAuto(t)
	ok := Equal(t, Separator, AssertOptions(assert.Report(report)), ValueStringer(func(v any) string {
		return "test" + Separator + "test"
	}))
	assert.False(t, ok)
}

func TestErrorNoValuesLeft(t *testing.T) {
	tmpDir := t.TempDir()
	fp := filepath.Join(tmpDir, "test.txt")
	t.Run("Save", func(t *testing.T) {
		ok := Equal(t, "test", FilePath(fp), Update(true))
		assert.True(t, ok)
	})
	report := asserttest.NewReportAuto(t)
	t.Run("Load", func(t *testing.T) {
		ok := Equal(t, "test", FilePath(fp), Update(false))
		assert.True(t, ok)
		ok = Equal(t, "test", FilePath(fp), Update(false), AssertOptions(assert.Report(report)))
		assert.False(t, ok)
	})
}

func TestErrorRemainingValues(t *testing.T) {
	tmpDir := t.TempDir()
	fp := filepath.Join(tmpDir, "test.txt")
	t.Run("Save", func(t *testing.T) {
		Equal(t, "test", FilePath(fp), Update(true))
		Equal(t, "test", FilePath(fp), Update(true))
	})
	report := asserttest.NewReportAuto(t)
	t.Run("Load", func(t *testing.T) {
		Equal(t, "test", FilePath(fp), Update(false), AssertOptions(assert.Report(report)))
	})
}

func TestErrorReadFile(t *testing.T) {
	tmpDir := t.TempDir()
	fp := filepath.Join(tmpDir, "test.txt")
	ok := Equal(t, "test", FilePath(fp), Update(false), AssertOptions(assert.Report(func(args ...any) {})))
	assert.False(t, ok)
}
