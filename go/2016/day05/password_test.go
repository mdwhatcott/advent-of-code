package main

import (
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func TestSequentialPassword(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test.")
	}
	assert := assertions.New(t)
	password := NewSequentialPassword("abc")
	assert.So(password.String(), should.Equal, "18f47a30")
}

func TestPositionalPassword(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test.")
	}

	assert := assertions.New(t)
	password := NewPositionalPassword("abc")
	assert.So(password.String(), should.Equal, "05ace8e3")
}
