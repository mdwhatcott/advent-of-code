package main

import (
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"
)

func TestValidTriangle(t *testing.T) {
	a := assert.Error(t)
	a.So(isValidTriangle("4", "2", "2"), should.BeFalse)
	a.So(isValidTriangle("2", "4", "2"), should.BeFalse)
	a.So(isValidTriangle("2", "2", "4"), should.BeFalse)
	a.So(isValidTriangle("2", "2", "2"), should.BeTrue)
}
