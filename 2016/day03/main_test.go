package main

import (
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func TestValidTriangle(t *testing.T) {
	assert := assertions.New(t)
	assert.So(isValidTriangle("4", "2", "2"), should.BeFalse)
	assert.So(isValidTriangle("2", "4", "2"), should.BeFalse)
	assert.So(isValidTriangle("2", "2", "4"), should.BeFalse)
	assert.So(isValidTriangle("2", "2", "2"), should.BeTrue)
}
