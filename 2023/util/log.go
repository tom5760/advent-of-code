package util

import (
	"log"
	"testing"
)

func NewTestLogger(t *testing.T) *log.Logger {
	t.Helper()

	return log.New(&testLogWriter{t: t}, "", 0)
}

type testLogWriter struct {
	t *testing.T
}

func (w *testLogWriter) Write(p []byte) (int, error) {
	w.t.Log(string(p))
	return len(p), nil
}
