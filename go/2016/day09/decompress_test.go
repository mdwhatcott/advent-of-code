package main

import (
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"
)

func Test(t *testing.T) {
	a := assert.Error(t)

	a.So(GetUncompressedSize("ADVENT", false), should.Equal, len("ADVENT"))
	a.So(GetUncompressedSize("A(1x5)BC", false), should.Equal, len("ABBBBBC"))
	a.So(GetUncompressedSize("(3x3)XYZ", false), should.Equal, len("XYZXYZXYZ"))
	a.So(GetUncompressedSize("A(2x2)BCD(2x2)EFG", false), should.Equal, len("ABCBCDEFEFG"))
	a.So(GetUncompressedSize("(6x1)(1x3)A", false), should.Equal, len("(1x3)A"))
}

func Test2(t *testing.T) {
	a := assert.Error(t)

	a.So(GetUncompressedSize("(3x3)XYZ", true), should.Equal, len("XYZXYZXYZ"))
	a.So(GetUncompressedSize("X(8x2)(3x3)ABCY", true), should.Equal, len("XABCABCABCABCABCABCY"))
	a.So(GetUncompressedSize("(27x12)(20x12)(13x14)(7x10)(1x12)A", true), should.Equal, 241920)
	a.So(GetUncompressedSize("(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", true), should.Equal, 445)
}
