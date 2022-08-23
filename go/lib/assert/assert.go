package assert

import (
	"reflect"
	"testing"
)

// That allows assertions as in: assert.That(t, actual).Equals(expected)

func That(t *testing.T, actual interface{}) *assertion {
	return &assertion{t: t, actual: actual}
}

type assertion struct {
	t      *testing.T
	actual interface{}
}

func (this *assertion) IsNil()   { this.t.Helper(); this.Equals(nil) }
func (this *assertion) IsTrue()  { this.t.Helper(); this.Equals(true) }
func (this *assertion) IsFalse() { this.t.Helper(); this.Equals(false) }
func (this *assertion) Equals(expected interface{}) {
	this.t.Helper()
	if !reflect.DeepEqual(this.actual, expected) {
		this.t.Errorf("\n"+
			"Expected: %#v\n"+
			"Actual:   %#v",
			expected,
			this.actual,
		)
	}
}
