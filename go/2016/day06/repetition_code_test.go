package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func TestRepetitionCode(t *testing.T) {
	code := NewRepetitionCode(bufio.NewScanner(strings.NewReader(input)), 6)
	assert := assertions.New(t)
	assert.So(code.DecodeFrequent(), should.Equal, "easter")
	assert.So(code.DecodeInfrequent(), should.Equal, "advent")

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
