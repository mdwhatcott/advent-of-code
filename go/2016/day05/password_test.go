package main

import (
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"
)

func TestSequentialPassword(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test.")
	}
	a := assert.Error(t)
	password := NewSequentialPassword("abc")
	a.So(password.String(), should.Equal, "18f47a30")
}

func TestPositionalPassword(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test.")
	}

	a := assert.Error(t)
	password := NewPositionalPassword("abc")
	a.So(password.String(), should.Equal, "05ace8e3")
}
