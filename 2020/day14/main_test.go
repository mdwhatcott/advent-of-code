package advent

import (
	"bufio"
	"strings"
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func TestApplyMask(t *testing.T) {
	assert := assertions.New(t)
	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"
	assert.So(ApplyMask("11", mask), should.Equal, 73)
	assert.So(ApplyMask("101", mask), should.Equal, 101)
	assert.So(ApplyMask("0", mask), should.Equal, 64)
}

func TestPart1(t *testing.T) {
	result := part1(bufio.NewScanner(strings.NewReader(example1)))
	assertions.New(t).So(result, should.Equal, 165)
}

const example1 = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`