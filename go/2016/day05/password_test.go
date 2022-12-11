package main

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestSequentialPassword(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test.")
	}
	password := NewSequentialPassword("abc")
	should.So(t, password.String(), should.Equal, "18f47a30")
}

func TestPositionalPassword(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test.")
	}

	password := NewPositionalPassword("abc")
	should.So(t, password.String(), should.Equal, "05ace8e3")
}
