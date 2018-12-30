package advent

import (
	"runtime"
	"testing"
)

func AssertEqual(t *testing.T, actual, expected interface{}) {
	if actual != expected {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf(assertFormat, file, line, actual, expected)
	}
}

const assertFormat = `
At:   %s:%d
Got:  [%v]
Want: [%v]
`

