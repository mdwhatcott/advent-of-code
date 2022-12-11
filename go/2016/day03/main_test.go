package main

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestValidTriangle(t *testing.T) {
	should.So(t, isValidTriangle("4", "2", "2"), should.BeFalse)
	should.So(t, isValidTriangle("2", "4", "2"), should.BeFalse)
	should.So(t, isValidTriangle("2", "2", "4"), should.BeFalse)
	should.So(t, isValidTriangle("2", "2", "2"), should.BeTrue)
}
