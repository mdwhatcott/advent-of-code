package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"
)

func TestRepetitionCode(t *testing.T) {
	code := NewRepetitionCode(bufio.NewScanner(strings.NewReader(input)), 6)
	a := assert.Error(t)
	a.So(code.DecodeFrequent(), should.Equal, "easter")
	a.So(code.DecodeInfrequent(), should.Equal, "advent")

}

const input = `eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar`
