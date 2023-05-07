package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func TestRepetitionCode(t *testing.T) {
	code := NewRepetitionCode(bufio.NewScanner(strings.NewReader(util.InputString())), 6)
	should.So(t, code.DecodeFrequent(), should.Equal, "easter")
	should.So(t, code.DecodeInfrequent(), should.Equal, "advent")

}
