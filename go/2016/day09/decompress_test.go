package main

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func Test(t *testing.T) {
	should.So(t, GetUncompressedSize("ADVENT", false), should.Equal, len("ADVENT"))
	should.So(t, GetUncompressedSize("A(1x5)BC", false), should.Equal, len("ABBBBBC"))
	should.So(t, GetUncompressedSize("(3x3)XYZ", false), should.Equal, len("XYZXYZXYZ"))
	should.So(t, GetUncompressedSize("A(2x2)BCD(2x2)EFG", false), should.Equal, len("ABCBCDEFEFG"))
	should.So(t, GetUncompressedSize("(6x1)(1x3)A", false), should.Equal, len("(1x3)A"))
}

func Test2(t *testing.T) {
	should.So(t, GetUncompressedSize("(3x3)XYZ", true), should.Equal, len("XYZXYZXYZ"))
	should.So(t, GetUncompressedSize("X(8x2)(3x3)ABCY", true), should.Equal, len("XABCABCABCABCABCABCY"))
	should.So(t, GetUncompressedSize("(27x12)(20x12)(13x14)(7x10)(1x12)A", true), should.Equal, 241920)
	should.So(t, GetUncompressedSize("(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", true), should.Equal, 445)
}
