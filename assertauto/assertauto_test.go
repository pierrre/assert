package assertauto_test

import (
	"testing"

	"github.com/pierrre/assert"
	. "github.com/pierrre/assert/assertauto"
	"github.com/pierrre/assert/asserttest"
	"github.com/pierrre/pretty"
)

func init() {
	pretty.DefaultCommonValueWriter.ConfigureTest()
}

func TestEqual(t *testing.T) {
	test(t, func(t *testing.T, opts ...Option) { //nolint:thelper // This is not a helper.
		ok := Equal(t, 123, append(opts, Name("int"))...)
		assert.True(t, ok)
		ok = Equal(t, "foo", append(opts, Name("string"))...)
		assert.True(t, ok)
	})
}

func TestEqualFail(t *testing.T) {
	rootTest := t
	tmpDir := t.TempDir()
	t.Run("Write", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(true))
		ok := Equal(t, 123, Update(true))
		assert.True(t, ok)
	})
	t.Run("Read", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(false))
		ok := Equal(t, 456, Update(false), AssertOptions(assert.Report(asserttest.NewReportAuto(rootTest))))
		assert.False(t, ok)
	})
}

func TestEqualFailName(t *testing.T) {
	testFailName(t, func(t *testing.T, opts ...Option) bool { //nolint:thelper // This is not a helper.
		return Equal(t, 123, opts...)
	})
}

func TestEqualFailEntryType(t *testing.T) {
	rootTest := t
	tmpDir := t.TempDir()
	t.Run("Write", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(true))
		ok := DeepEqual(t, &testStruct{
			Foo: "bar",
		}, Update(true))
		assert.True(t, ok)
	})
	t.Run("Read", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(false))
		ok := Equal(t, 123, Update(false), AssertOptions(assert.Report(asserttest.NewReportAuto(rootTest))))
		assert.False(t, ok)
	})
}

func TestEqualFailType(t *testing.T) {
	rootTest := t
	tmpDir := t.TempDir()
	t.Run("Write", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(true))
		ok := Equal(t, "foo", Update(true))
		assert.True(t, ok)
	})
	t.Run("Read", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(false))
		ok := Equal(t, 123, Update(false), AssertOptions(assert.Report(asserttest.NewReportAuto(rootTest))))
		assert.False(t, ok)
	})
}

func TestDeepEqual(t *testing.T) {
	test(t, func(t *testing.T, opts ...Option) { //nolint:thelper // This is not a helper.
		ok := DeepEqual(t, &testStruct{
			Foo: "bar",
		}, opts...)
		assert.True(t, ok)
	})
}

func TestDeepEqualFail(t *testing.T) {
	rootTest := t
	tmpDir := t.TempDir()
	t.Run("Write", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(true))
		ok := DeepEqual(t, &testStruct{
			Foo: "bar",
		}, Update(true))
		assert.True(t, ok)
	})
	t.Run("Read", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(false))
		ok := DeepEqual(t, &testStruct{
			Foo: "baz",
		}, Update(false), AssertOptions(assert.Report(asserttest.NewReportAuto(rootTest))))
		assert.False(t, ok)
	})
}

func TestDeepEqualFailName(t *testing.T) {
	testFailName(t, func(t *testing.T, opts ...Option) bool { //nolint:thelper // This is not a helper.
		return DeepEqual(t, &testStruct{
			Foo: "bar",
		}, opts...)
	})
}

func TestDeepEqualFailEntryType(t *testing.T) {
	rootTest := t
	tmpDir := t.TempDir()
	t.Run("Write", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(true))
		ok := Equal(t, 123, Update(true))
		assert.True(t, ok)
	})
	t.Run("Read", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(false))
		ok := DeepEqual(t, &testStruct{
			Foo: "bar",
		}, Update(false), AssertOptions(assert.Report(asserttest.NewReportAuto(rootTest))))
		assert.False(t, ok)
	})
}

func TestDeepEqualFailType(t *testing.T) {
	rootTest := t
	tmpDir := t.TempDir()
	t.Run("Write", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(true))
		ok := DeepEqual(t, &testStruct{
			Foo: "bar",
		}, Update(true))
		assert.True(t, ok)
	})
	t.Run("Read", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(false))
		ok := DeepEqual(t, 123, Update(false), AssertOptions(assert.Report(asserttest.NewReportAuto(rootTest))))
		assert.False(t, ok)
	})
}

func TestAllocsPerRun(t *testing.T) {
	test(t, func(t *testing.T, opts ...Option) { //nolint:thelper // This is not a helper.
		allocs, ok := AllocsPerRun(t, 10, func() {
			_ = make([]byte, 1<<20)
		}, opts...)
		assert.True(t, ok)
		assert.Equal(t, 1, allocs)
	})
}

func TestAllocsPerRunFail(t *testing.T) {
	rootTest := t
	tmpDir := t.TempDir()
	t.Run("Write", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(true))
		_, ok := AllocsPerRun(t, 10, func() {}, Update(true))
		assert.True(t, ok)
	})
	t.Run("Read", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(false))
		_, ok := AllocsPerRun(t, 10, func() {
			_ = make([]byte, 1<<20)
		}, Update(false), AssertOptions(assert.Report(asserttest.NewReportAuto(rootTest))))
		assert.False(t, ok)
	})
}

func TestAllocsPerRunFailName(t *testing.T) {
	testFailName(t, func(t *testing.T, opts ...Option) bool { //nolint:thelper // This is not a helper.
		_, ok := AllocsPerRun(t, 10, func() {
			_ = make([]byte, 1<<20)
		}, opts...)
		return ok
	})
}

func TestAllocsPerRunFailEntryType(t *testing.T) {
	rootTest := t
	tmpDir := t.TempDir()
	t.Run("Write", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(true))
		ok := Equal(t, 123, Update(true))
		assert.True(t, ok)
	})
	t.Run("Read", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(false))
		_, ok := AllocsPerRun(t, 10, func() {
			_ = make([]byte, 1<<20)
		}, Update(false), AssertOptions(assert.Report(asserttest.NewReportAuto(rootTest))))
		assert.False(t, ok)
	})
}

func TestAllocsPerRunFailRuns(t *testing.T) {
	rootTest := t
	tmpDir := t.TempDir()
	t.Run("Write", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(true))
		_, ok := AllocsPerRun(t, 10, func() {
			_ = make([]byte, 1<<20)
		}, Update(true))
		assert.True(t, ok)
	})
	t.Run("Read", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(false))
		_, ok := AllocsPerRun(t, 20, func() {
			_ = make([]byte, 1<<20)
		}, Update(false), AssertOptions(assert.Report(asserttest.NewReportAuto(rootTest))))
		assert.False(t, ok)
	})
}

func TestFailEntryNoRemaining(t *testing.T) {
	rootTest := t
	tmpDir := t.TempDir()
	t.Run("Write", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(true))
		ok := Equal(t, 123, Update(true))
		assert.True(t, ok)
	})
	t.Run("Read", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(false))
		ok := Equal(t, 123, Update(false))
		assert.True(t, ok)
		ok = Equal(t, 123, Update(false), AssertOptions(assert.Report(asserttest.NewReportAuto(rootTest))))
		assert.False(t, ok)
	})
}

func TestFailEntriesRemaining(t *testing.T) {
	rootTest := t
	tmpDir := t.TempDir()
	t.Run("Write", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(true))
		ok := Equal(t, 123, Update(true))
		assert.True(t, ok)
		ok = Equal(t, 123, Update(true))
		assert.True(t, ok)
	})
	t.Run("Read", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(false), AssertOptions(assert.Report(asserttest.NewReportAuto(rootTest))))
		ok := Equal(t, 123, Update(false))
		assert.True(t, ok)
	})
}

func test(t *testing.T, f func(t *testing.T, opts ...Option)) {
	t.Helper()
	rootTest := t
	tmpDir := t.TempDir()
	t.Run("Write", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(true))
		f(t, Update(true))
	})
	t.Run("Read", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(false))
		f(t, Update(false))
	})
	t.Run("Normal", func(t *testing.T) {
		InitTestFunction(t, DirectoryGlobal, rootTest.Name())
		f(t)
	})
}

func testFailName(t *testing.T, f func(t *testing.T, opts ...Option) bool) {
	t.Helper()
	rootTest := t
	tmpDir := t.TempDir()
	t.Run("Write", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(true))
		ok := f(t, Update(true), Name("foo"))
		assert.True(t, ok)
	})
	t.Run("Read", func(t *testing.T) {
		InitTestFunction(t, tmpDir, rootTest.Name(), Update(false))
		ok := f(t, Update(false), Name("bar"), AssertOptions(assert.Report(asserttest.NewReportAuto(rootTest))))
		assert.False(t, ok)
	})
}

type testStruct struct {
	Foo string
}
