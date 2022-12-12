package internal_test

import (
	"github.com/danvolchek/AdventOfCode/cmd/internal"
	"testing"
)

type testWriteCloser struct {
	data   string
	closed bool
}

func (t *testWriteCloser) Write(p []byte) (n int, err error) {
	t.data += string(p)

	return len(p), nil
}

func (t *testWriteCloser) Close() error {
	t.closed = true

	return nil
}

func TestMultiWriteCloser(t *testing.T) {
	data1, data2 := []byte("hello"), []byte("world")
	wc1, wc2 := &testWriteCloser{}, &testWriteCloser{}

	mw := &internal.MultiWriteCloser{}

	mw.Add(wc1)

	l, err := mw.Write(data1)
	if l != len(data1) {
		t.Errorf("short write: expected %v, actual %v", len(data1), l)
	}

	if err != nil {
		t.Errorf("unexpected write error: %v", err)
	}

	if wc1.data != string(data1) {
		t.Errorf("incorrect data in wc1 after first write: expected %v, actual %v", string(data1), wc1.data)
	}

	mw.Add(wc2)

	l, err = mw.Write(data2)
	if l != len(data2) {
		t.Errorf("short write: expected %v, actual %v", len(data2), l)
	}

	if err != nil {
		t.Errorf("unexpected write error: %v", err)
	}

	if wc1.data != string(data1)+string(data2) {
		t.Errorf("incorrect data in wc1 after second write: expected %v, actual %v", string(data1)+string(data2), wc1.data)
	}

	if wc2.data != string(data2) {
		t.Errorf("incorrect data in wc2 after second write: expected %v, actual %v", string(data2), wc2.data)
	}

	err = mw.Close()
	if err != nil {
		t.Errorf("unexpected close error: %v", err)
	}

	if !wc1.closed {
		t.Errorf("wc1 not closed")
	}

	if !wc2.closed {
		t.Errorf("wc2 not closed")
	}
}
