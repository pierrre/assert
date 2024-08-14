// Package assertauto provides helpers to automatically update the expected values of assertions.
//
// This is highly experimental an not yet ready for public usage.
package assertauto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"testing"

	"github.com/pierrre/assert"
	"github.com/pierrre/go-libs/reflectutil"
)

const (
	directoryEnvVar  = "ASSERTAUTO_DIRECTORY"
	directoryDefault = "_assertauto"
	updateEnvVar     = "ASSERTAUTO_UPDATE"
	updateDefault    = false
)

var (
	directoryGlobal = initDirectoryGlobal()
	updateGlobal    = initUpdateGlobal()
)

func init() {
	if updateGlobal {
		err := os.RemoveAll(directoryGlobal)
		if err != nil {
			panic(err)
		}
	}
}

func initDirectoryGlobal() string {
	s, ok := os.LookupEnv(directoryEnvVar)
	if !ok {
		return directoryDefault
	}
	if s == "" {
		return directoryDefault
	}
	return s
}

func initUpdateGlobal() bool {
	s, ok := os.LookupEnv(updateEnvVar)
	if !ok {
		return updateDefault
	}
	b, err := strconv.ParseBool(s)
	if err != nil {
		err = fmt.Errorf("parse %s environment variable: %w", updateEnvVar, err)
		panic(err)
	}
	return b
}

// Equal asserts that the value is equal to the expected value.
func Equal[T comparable](tb testing.TB, v T, optfs ...Option) bool { //nolint:dupl // TODO deduplicate.
	tb.Helper()
	opts := buildOptions(optfs)
	typeName := getTypeName[T]()
	if opts.update {
		addEntry(tb, entry{
			Equal: &equalEntry{
				Type:  typeName,
				Value: jsonEncode(tb, v),
			},
		}, opts)
		return true
	}
	e, ok := getEntry(tb, opts)
	if !ok {
		return false
	}
	if !assert.NotZero(tb, e.Equal, append(opts.opts, assert.Message("assertauto: equal: wrong entry type"))...) {
		return false
	}
	if !assert.Equal(tb, e.Equal.Type, typeName, append(opts.opts, assert.MessageWrap("assertauto: equal: type"))...) {
		return false
	}
	expected := jsonDecode[T](tb, e.Equal.Value)
	return assert.Equal(tb, v, expected, append(opts.opts, assert.MessageWrap("assertauto"))...)
}

// DeepEqual asserts that the value is deep equal to the expected value.
func DeepEqual[T any](tb testing.TB, v T, optfs ...Option) bool { //nolint:dupl // TODO deduplicate.
	tb.Helper()
	opts := buildOptions(optfs)
	typeName := getTypeName[T]()
	if opts.update {
		addEntry(tb, entry{
			DeepEqual: &deepEqualEntry{
				Type:  typeName,
				Value: jsonEncode(tb, v),
			},
		}, opts)
		return true
	}
	e, ok := getEntry(tb, opts)
	if !ok {
		return false
	}
	if !assert.NotZero(tb, e.DeepEqual, append(opts.opts, assert.Message("assertauto: deep equal: wrong entry type"))...) {
		return false
	}
	if !assert.Equal(tb, e.DeepEqual.Type, typeName, append(opts.opts, assert.MessageWrap("assertauto: deep equal: type"))...) {
		return false
	}
	expected := jsonDecode[T](tb, e.DeepEqual.Value)
	return assert.DeepEqual(tb, v, expected, append(opts.opts, assert.MessageWrap("assertauto"))...)
}

// AllocsPerRun asserts that a function allocates a certain number of times per run.
func AllocsPerRun(tb testing.TB, runs int, f func(), optfs ...Option) bool {
	tb.Helper()
	opts := buildOptions(optfs)
	if opts.update {
		allocs := testing.AllocsPerRun(runs, f)
		addEntry(tb, entry{
			AllocsPerRun: &allocsPerRunEntry{
				Runs:   runs,
				Allocs: allocs,
			},
		}, opts)
		return true
	}
	e, ok := getEntry(tb, opts)
	if !ok {
		return false
	}
	if !assert.NotZero(tb, e.AllocsPerRun, append(opts.opts, assert.Message("assertauto: allocs per run: wrong entry type"))...) {
		return false
	}
	if !assert.Equal(tb, e.AllocsPerRun.Runs, runs, append(opts.opts, assert.MessageWrap("assertauto: allocs per run: runs"))...) {
		return false
	}
	expected := e.AllocsPerRun.Allocs
	return assert.AllocsPerRun(tb, runs, f, expected, append(opts.opts, assert.MessageWrap("assertauto"))...)
}

func addEntry(tb testing.TB, e entry, opts *options) {
	tb.Helper()
	e.Name = opts.name
	tf := getTestFunction(tb, opts)
	tf.addEntry(tb, e)
}

func getEntry(tb testing.TB, opts *options) (entry, bool) {
	tb.Helper()
	tf := getTestFunction(tb, opts)
	e, ok := tf.getEntry(tb, opts)
	if !ok {
		return entry{}, false
	}
	if !assert.Equal(tb, e.Name, opts.name, append(opts.opts, assert.MessageWrap("assertauto: entry name"))...) {
		return e, false
	}
	return e, true
}

var (
	testFunctionsLock sync.Mutex
	testFunctions     = make(map[string]*testFunction)
)

func getTestFunction(tb testing.TB, opts *options) *testFunction {
	tb.Helper()
	return getTestFunctionWithFile(tb, "", opts)
}

func getTestFunctionWithFile(tb testing.TB, fp string, opts *options) *testFunction {
	tb.Helper()
	testFunctionsLock.Lock()
	defer testFunctionsLock.Unlock()
	name := tb.Name()
	tf, ok := testFunctions[name]
	if !ok {
		tf = newTestFunction(tb, fp, opts)
		testFunctions[name] = tf
		tb.Cleanup(func() {
			tb.Helper()
			deleteTestFunction(name)
		})
	}
	return tf
}

func deleteTestFunction(name string) {
	testFunctionsLock.Lock()
	defer testFunctionsLock.Unlock()
	delete(testFunctions, name)
}

type testFunction struct {
	mu      sync.Mutex
	update  bool
	entries []entry
}

func newTestFunction(tb testing.TB, fp string, opts *options) *testFunction {
	tb.Helper()
	if fp == "" {
		fp = getFilePathGlobal(tb)
	}
	tf := &testFunction{
		update: opts.update,
	}
	if !tf.update {
		tf.load(tb, fp)
	}
	tb.Cleanup(func() {
		tb.Helper()
		tf.cleanup(tb, fp, opts)
	})
	return tf
}

func (tf *testFunction) load(tb testing.TB, fp string) {
	tb.Helper()
	b, err := os.ReadFile(fp)
	assert.NoError(tb, err)
	f := jsonDecode[file](tb, b)
	tf.entries = f.Entries
}

func (tf *testFunction) save(tb testing.TB, fp string) {
	tb.Helper()
	f := &file{
		Entries: tf.entries,
	}
	data := jsonEncode(tb, f)
	dir := filepath.Dir(fp)
	err := os.MkdirAll(dir, 0o755)
	assert.NoError(tb, err)
	err = os.RemoveAll(fp)
	assert.NoError(tb, err)
	err = os.WriteFile(fp, data, 0o644) //nolint:gosec // We want 644.
	assert.NoError(tb, err)
}

func (tf *testFunction) addEntry(tb testing.TB, entry entry) {
	tb.Helper()
	tf.mu.Lock()
	defer tf.mu.Unlock()
	assert.True(tb, tf.update, assert.MessageWrap("assertauto: cannot add entry if update is false"))
	tf.entries = append(tf.entries, entry)
}

func (tf *testFunction) getEntry(tb testing.TB, opts *options) (entry, bool) {
	tb.Helper()
	tf.mu.Lock()
	defer tf.mu.Unlock()
	assert.False(tb, tf.update, assert.MessageWrap("assertauto: cannot get entry if update is true"))
	ok := assert.SliceNotEmpty(tb, tf.entries, append(opts.opts, assert.MessageWrap("assertauto: no entry remaining"))...)
	if !ok {
		return entry{}, false
	}
	e := tf.entries[0]
	tf.entries = tf.entries[1:]
	return e, true
}

func (tf *testFunction) cleanup(tb testing.TB, fp string, opts *options) {
	tb.Helper()
	if opts.update {
		tf.save(tb, fp)
	} else if !tb.Failed() {
		assert.SliceEmpty(tb, tf.entries, append(opts.opts, assert.MessageWrap("assertauto: entries remaining"))...)
	}
}

type file struct {
	Entries []entry `json:"entries"`
}

type entry struct {
	Name         string             `json:"name,omitempty"`
	Equal        *equalEntry        `json:"equal,omitempty"`
	DeepEqual    *deepEqualEntry    `json:"deep_equal,omitempty"`
	AllocsPerRun *allocsPerRunEntry `json:"allocs_per_run,omitempty"`
}

type equalEntry struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}

type deepEqualEntry struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}

type allocsPerRunEntry struct {
	Runs   int     `json:"runs"`
	Allocs float64 `json:"allocs"`
}

func jsonEncode(tb testing.TB, v any) []byte {
	tb.Helper()
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetIndent("", "\t")
	enc.SetEscapeHTML(false)
	err := enc.Encode(v)
	assert.NoError(tb, err)
	return buf.Bytes()
}

func jsonDecode[T any](tb testing.TB, data []byte) T {
	tb.Helper()
	var v T
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	err := dec.Decode(&v)
	assert.NoError(tb, err)
	return v
}

func getFilePathGlobal(tb testing.TB) string {
	tb.Helper()
	return getFilePath(directoryGlobal, tb.Name())
}

func getFilePath(dir string, name string) string {
	return filepath.Join(dir, name+".json")
}

type options struct {
	update bool
	name   string
	opts   []assert.Option
}

func buildOptions(opts []Option) *options {
	o := newOptions()
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func newOptions() *options {
	return &options{
		update: updateGlobal,
	}
}

// Option is an option for assertauto.
type Option func(*options)

func update(u bool) Option {
	return func(o *options) {
		o.update = u
	}
}

// Name returns an [Option] that sets the name of the entry.
func Name(name string) Option {
	return func(o *options) {
		o.name = name
	}
}

// AssertOptions returns an [Option] that sets the assert options.
func AssertOptions(opts ...assert.Option) Option {
	return func(o *options) {
		o.opts = opts
	}
}

func getTypeName[T any]() string {
	return reflectutil.TypeFullNameFor[T]()
}
