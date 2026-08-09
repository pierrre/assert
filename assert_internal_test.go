package assert

import (
	"testing"
)

func TestBytesWriterPoolMaxCapBounded(t *testing.T) {
	if bytesWriterPool.MaxCap <= 0 {
		t.Fatalf("bytesWriterPool.MaxCap must be bounded (positive)")
	}
}

func TestBytesWriterPoolDiscardsOversized(t *testing.T) {
	maxCap := bytesWriterPool.MaxCap
	w1 := bytesWriterPool.Get()
	w1.Grow(maxCap + 1)
	if w1.Cap() <= maxCap {
		t.Fatalf("writer was not grown beyond MaxCap")
	}
	bytesWriterPool.Put(w1)
	w2 := bytesWriterPool.Get()
	if w2.Cap() > maxCap {
		t.Fatalf("pool retained an oversized writer")
	}
}
