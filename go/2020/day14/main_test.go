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

func TestPermuteFloatingBits(t *testing.T) {
	values := PermuteFloatingBits(
		"000000000000000000000000000000X1101X",
	)
	assertions.New(t).So(values, should.Resemble, []string{
		"000000000000000000000000000000011010",
		"000000000000000000000000000000011011",
		"000000000000000000000000000000111010",
		"000000000000000000000000000000111011",
	})
}

const example2 = `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`

func TestPart2(t *testing.T) {
	result := part2(bufio.NewScanner(strings.NewReader(example2)))
	assertions.New(t).So(result, should.Equal, 208)
}
