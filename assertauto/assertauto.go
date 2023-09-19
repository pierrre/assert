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
func Equal[T comparable](tb testing.TB, v T, opts ...assert.Option) bool {
	tb.Helper()
	tf := getTestFunction(tb)
	if update {
		e := entry{
			Equal: jsonEncode(tb, v),
		}
		tf.addEntry(e)
		return true
	}
	e := tf.getEntry(tb)
	assert.SliceNotEmpty(tb, e.Equal, assert.MessageWrap("assertauto: entry is not equal"))
	expected := jsonDecode[T](tb, e.Equal)
	return assert.Equal(tb, v, expected, opts...)
}

// DeepEqual asserts that the value is deep equal to the expected value.
func DeepEqual[T any](tb testing.TB, v T, opts ...assert.Option) bool {
	tb.Helper()
	tf := getTestFunction(tb)
	if update {
		e := entry{
			DeepEqual: jsonEncode(tb, v),
		}
		tf.addEntry(e)
		return true
	}
	e := tf.getEntry(tb)
	assert.SliceNotEmpty(tb, e.DeepEqual, assert.MessageWrap("assertauto: entry is not deep equal"))
	expected := jsonDecode[T](tb, e.DeepEqual)
	return assert.DeepEqual(tb, v, expected, opts...)
}

var (
	testFunctionsMutex sync.Mutex
	testFunctions      = make(map[string]*testFunction)
)

func getTestFunction(tb testing.TB) *testFunction {
	tb.Helper()
	testFunctionsMutex.Lock()
	defer testFunctionsMutex.Unlock()
	tf, ok := testFunctions[tb.Name()]
	if !ok {
		tf = &testFunction{}
		testFunctions[tb.Name()] = tf
		tb.Cleanup(func() {
			tb.Helper()
			tf.cleanup(tb)
		})
	}
	return tf
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
	} else {
		assert.SliceEmpty(tb, tf.entries, assert.MessageWrap("assertauto: entries remaining"))
	}
}

type file struct {
	Entries []entry `json:"entries"`
}

type entry struct {
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
