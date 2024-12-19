package testmarker

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

// Marker holds the file name and line number of the test case
type Marker struct {
	tb testing.TB

	caseName string
	File     string
	Line     int
}

var _ fmt.Stringer = (*Marker)(nil)

// Mark marks the file name and line number of the test case
func Mark(tb testing.TB, caseName string) Marker {
	_, file, line, _ := runtime.Caller(1)
	return Marker{
		caseName: caseName,
		File:     filepath.Base(file),
		Line:     line,
	}
}

// String returns the file name and line number of the test case
func (m *Marker) String() string {
	return fmt.Sprintf("%s(%s:%d)", m.caseName, m.File, m.Line)
}
