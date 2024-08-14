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
func Equal[T comparable](tb testing.TB, v T, opts ...Option) bool {
	tb.Helper()
	o := buildOptions(tb, opts)
	if o.update {
		addEntry(tb, entry{
			Equal: jsonEncode(tb, v),
		}, o)
		return true
	}
	e, ok := getEntry(tb, o)
	if !ok {
		return false
	}
	if !assert.SliceNotEmpty(tb, e.Equal, append(o.opts, assert.MessageWrap("assertauto: entry is not \"equal\""))...) {
		return false
	}
	expected := jsonDecode[T](tb, e.Equal)
	return assert.Equal(tb, v, expected, append(o.opts, assert.MessageWrap("assertauto"))...)
}

// DeepEqual asserts that the value is deep equal to the expected value.
func DeepEqual[T any](tb testing.TB, v T, opts ...Option) bool {
	tb.Helper()
	o := buildOptions(tb, opts)
	if o.update {
		addEntry(tb, entry{
			DeepEqual: jsonEncode(tb, v),
		}, o)
		return true
	}
	e, ok := getEntry(tb, o)
	if !ok {
		return false
	}
	if !assert.SliceNotEmpty(tb, e.DeepEqual, append(o.opts, assert.MessageWrap("assertauto: entry is not \"deep equal\""))...) {
		return false
	}
	expected := jsonDecode[T](tb, e.DeepEqual)
	return assert.DeepEqual(tb, v, expected, append(o.opts, assert.MessageWrap("assertauto"))...)
}

// AllocsPerRun asserts that a function allocates a certain number of times per run.
func AllocsPerRun(tb testing.TB, runs int, f func(), opts ...Option) bool {
	tb.Helper()
	o := buildOptions(tb, opts)
	if o.update {
		allocs := testing.AllocsPerRun(runs, f)
		addEntry(tb, entry{
			AllocsPerRun: &allocsPerRunEntry{
				Runs:   runs,
				Allocs: allocs,
			},
		}, o)
		return true
	}
	e, ok := getEntry(tb, o)
	if !ok {
		return false
	}
	if !assert.NotZero(tb, e.AllocsPerRun, append(o.opts, assert.MessageWrap("assertauto: entry is not \"allocs per run\""))...) {
		return false
	}
	if !assert.Equal(tb, e.AllocsPerRun.Runs, runs, append(o.opts, assert.MessageWrap("assertauto: allocs per run: runs"))...) {
		return false
	}
	expected := e.AllocsPerRun.Allocs
	return assert.AllocsPerRun(tb, runs, f, expected, append(o.opts, assert.MessageWrap("assertauto"))...)
}

func addEntry(tb testing.TB, e entry, opts *options) {
	tb.Helper()
	e.Name = opts.name
	tf := getTestFunction(tb, opts)
	tf.addEntry(e)
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
	testFunctionsLock.Lock()
	defer testFunctionsLock.Unlock()
	name := tb.Name()
	tf, ok := testFunctions[name]
	if !ok {
		tf = newTestFunction(tb, opts)
		testFunctions[name] = tf
		tb.Cleanup(func() {
			tb.Helper()
			deleteTestFunction(name)
			tf.cleanup(tb, opts)
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
	entries []entry
}

func newTestFunction(tb testing.TB, opts *options) *testFunction {
	tb.Helper()
	tf := &testFunction{}
	if !opts.update {
		tf.load(tb, opts)
	}
	return tf
}

func (tf *testFunction) load(tb testing.TB, opts *options) {
	tb.Helper()
	fp := getFilePath(opts)
	b, err := os.ReadFile(fp)
	assert.NoError(tb, err)
	f := jsonDecode[file](tb, b)
	tf.entries = f.Entries
}

func (tf *testFunction) save(tb testing.TB, opts *options) {
	tb.Helper()
	f := &file{
		Entries: tf.entries,
	}
	data := jsonEncode(tb, f)
	fp := getFilePath(opts)
	err := os.MkdirAll(opts.directory, 0o755)
	assert.NoError(tb, err)
	err = os.RemoveAll(fp)
	assert.NoError(tb, err)
	err = os.WriteFile(fp, data, 0o644) //nolint:gosec // We want 644.
	assert.NoError(tb, err)
}

func (tf *testFunction) addEntry(entry entry) {
	tf.mu.Lock()
	defer tf.mu.Unlock()
	tf.entries = append(tf.entries, entry)
}

func (tf *testFunction) getEntry(tb testing.TB, opts *options) (entry, bool) {
	tb.Helper()
	tf.mu.Lock()
	defer tf.mu.Unlock()
	ok := assert.SliceNotEmpty(tb, tf.entries, append(opts.opts, assert.MessageWrap("assertauto: no entry remaining"))...)
	if !ok {
		return entry{}, false
	}
	e := tf.entries[0]
	tf.entries = tf.entries[1:]
	return e, true
}

func (tf *testFunction) cleanup(tb testing.TB, opts *options) {
	tb.Helper()
	if opts.update {
		tf.save(tb, opts)
	} else if !tb.Failed() {
		assert.SliceEmpty(tb, tf.entries, append(opts.opts, assert.MessageWrap("assertauto: entries remaining"))...)
	}
}

type file struct {
	Entries []entry `json:"entries"`
}

type entry struct {
	Name         string             `json:"name,omitempty"`
	Equal        json.RawMessage    `json:"equal,omitempty"`
	DeepEqual    json.RawMessage    `json:"deep_equal,omitempty"`
	AllocsPerRun *allocsPerRunEntry `json:"allocs_per_run,omitempty"`
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

func getFilePath(opts *options) string {
	return filepath.Join(opts.directory, opts.fileName+".json")
}

type options struct {
	update    bool
	directory string
	fileName  string
	name      string
	opts      []assert.Option
}

func buildOptions(tb testing.TB, opts []Option) *options {
	tb.Helper()
	o := newOptions(tb)
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func newOptions(tb testing.TB) *options {
	tb.Helper()
	return &options{
		update:    updateGlobal,
		directory: directoryGlobal,
		fileName:  tb.Name(),
	}
}

// Option is an option for assertauto.
type Option func(*options)

func update(u bool) Option {
	return func(o *options) {
		o.update = u
	}
}

func directory(d string) Option {
	return func(o *options) {
		o.directory = d
	}
}

func fileName(n string) Option {
	return func(o *options) {
		o.fileName = n
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
