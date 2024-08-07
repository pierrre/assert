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
	directory    = "_assertauto"
	updateEnvVar = "ASSERTAUTO_UPDATE"
)

var update = false

func init() {
	s, ok := os.LookupEnv(updateEnvVar)
	if !ok {
		return
	}
	b, err := strconv.ParseBool(s)
	if err != nil {
		err = fmt.Errorf("parse %s environment variable: %w", updateEnvVar, err)
		panic(err)
	}
	update = b
	if update {
		err = os.RemoveAll(directory)
		if err != nil {
			panic(err)
		}
	}
}

// Equal asserts that the value is equal to the expected value.
func Equal[T comparable](tb testing.TB, v T, opts ...Option) bool {
	tb.Helper()
	o := buildOptions(opts)
	if update {
		addEntry(tb, entry{
			Equal: jsonEncode(tb, v),
		}, o)
		return true
	}
	e := getEntry(tb, o)
	if !assert.SliceNotEmpty(tb, e.Equal, append(o.opts, assert.MessageWrap("assertauto: entry is not equal"))...) {
		return false
	}
	expected := jsonDecode[T](tb, e.Equal)
	return assert.Equal(tb, v, expected, append(o.opts, assert.MessageWrap("assertauto"))...)
}

// DeepEqual asserts that the value is deep equal to the expected value.
func DeepEqual[T any](tb testing.TB, v T, opts ...Option) bool {
	tb.Helper()
	o := buildOptions(opts)
	if update {
		addEntry(tb, entry{
			DeepEqual: jsonEncode(tb, v),
		}, o)
		return true
	}
	e := getEntry(tb, o)
	if !assert.SliceNotEmpty(tb, e.DeepEqual, assert.MessageWrap("assertauto: entry is not deep equal")) {
		return false
	}
	expected := jsonDecode[T](tb, e.DeepEqual)
	return assert.DeepEqual(tb, v, expected, append(o.opts, assert.MessageWrap("assertauto"))...)
}

func addEntry(tb testing.TB, e entry, opts *options) {
	tb.Helper()
	e.Name = opts.name
	tf := getTestFunction(tb)
	tf.addEntry(e)
}

func getEntry(tb testing.TB, opts *options) entry {
	tb.Helper()
	tf := getTestFunction(tb)
	e := tf.getEntry(tb)
	assert.Equal(tb, e.Name, opts.name, assert.MessageWrap("assertauto: entry name"))
	return e
}

var (
	testFunctionsMutex sync.Mutex
	testFunctions      = make(map[string]*testFunction)
)

func getTestFunction(tb testing.TB) *testFunction {
	tb.Helper()
	testFunctionsMutex.Lock()
	defer testFunctionsMutex.Unlock()
	name := tb.Name()
	tf, ok := testFunctions[name]
	if !ok {
		tf = &testFunction{}
		testFunctions[name] = tf
		tb.Cleanup(func() {
			tb.Helper()
			deleteTestFunction(name)
			tf.cleanup(tb)
		})
	}
	return tf
}

func deleteTestFunction(name string) {
	testFunctionsMutex.Lock()
	defer testFunctionsMutex.Unlock()
	delete(testFunctions, name)
}

type testFunction struct {
	mu          sync.Mutex
	initialized bool
	entries     []entry
}

func (tf *testFunction) load(tb testing.TB) {
	tb.Helper()
	fp := getFilePath(tb)
	b, err := os.ReadFile(fp)
	assert.NoError(tb, err)
	f := jsonDecode[file](tb, b)
	tf.entries = f.Entries
}

func (tf *testFunction) save(tb testing.TB) {
	tb.Helper()
	f := &file{
		Entries: tf.entries,
	}
	data := jsonEncode(tb, f)
	fp := getFilePath(tb)
	dir := filepath.Dir(fp)
	err := os.MkdirAll(dir, 0o755)
	assert.NoError(tb, err)
	err = os.WriteFile(fp, data, 0o644) //nolint:gosec // We want 644.
	assert.NoError(tb, err)
}

func (tf *testFunction) addEntry(entry entry) {
	tf.mu.Lock()
	defer tf.mu.Unlock()
	tf.entries = append(tf.entries, entry)
}

func (tf *testFunction) getEntry(tb testing.TB) entry {
	tb.Helper()
	tf.mu.Lock()
	defer tf.mu.Unlock()
	if !tf.initialized && !update {
		tf.load(tb)
		tf.initialized = true
	}
	assert.SliceNotEmpty(tb, tf.entries, assert.MessageWrap("assertauto: no entry remaining"))
	e := tf.entries[0]
	tf.entries = tf.entries[1:]
	return e
}

func (tf *testFunction) cleanup(tb testing.TB) {
	tb.Helper()
	if update {
		tf.save(tb)
	} else if !tb.Failed() {
		assert.SliceEmpty(tb, tf.entries, assert.MessageWrap("assertauto: entries remaining"))
	}
}

type file struct {
	Entries []entry `json:"entries"`
}

type entry struct {
	Name      string          `json:"name,omitempty"`
	Equal     json.RawMessage `json:"equal,omitempty"`
	DeepEqual json.RawMessage `json:"deep_equal,omitempty"`
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

func getFilePath(tb testing.TB) string {
	tb.Helper()
	return filepath.Join(directory, tb.Name()+".json")
}

type options struct {
	name string
	opts []assert.Option
}

func buildOptions(opts []Option) *options {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// Option is an option for assertauto.
type Option func(*options)

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
