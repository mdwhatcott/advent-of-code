package day21

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestStuff(t *testing.T) {
	gunit.Run(new(Stuff), t)
}

type Stuff struct {
	*gunit.Fixture
}

func (this *Stuff) Test2x2PatternRotation() {
	this.So(Transformations("../.."), should.Resemble, []string{
		"../..",
	})
	this.So(Transformations("#./.#"), should.Resemble, []string{
		"#./.#",
		".#/#.",
	})
	this.So(Transformations("##/##"), should.Resemble, []string{"##/##"})
	this.So(Transformations("12/34"), should.Resemble, []string{
		"12/34",
		"31/42",
		"43/21",
		"24/13",
		"34/12",
		"13/24",
		"21/43",
		"42/31",
	})
}

func (this *Stuff) Test3x3PatternRotation() {
	this.So(Transformations(".../.../..."), should.Resemble, []string{
		".../.../...",
	})
	this.So(Transformations("123/456/789"), should.Resemble, []string{
		"123/456/789",
		"741/852/963",
		"987/654/321",
		"369/258/147",
		"789/456/123",
		"147/258/369",
		"321/654/987",
		"963/852/741",
	})
}

func (this *Stuff) TestPatternTransformations() {
	transformer := RegisterEnhancementRules(".#./..#/### => #..#/..../..../#..#")

	this.So(func() { transformer.Enhance("not found") }, should.Panic)
	this.So(transformer.Enhance(".#./..#/###"), should.Resemble, "#..#/..../..../#..#")
	this.So(transformer.Enhance("#../#.#/##."), should.Resemble, "#..#/..../..../#..#")
	this.So(transformer.Enhance("###/#../.#."), should.Resemble, "#..#/..../..../#..#")
	this.So(transformer.Enhance(".##/#.#/..#"), should.Resemble, "#..#/..../..../#..#")
	this.So(transformer.Enhance("###/..#/.#."), should.Resemble, "#..#/..../..../#..#")
	this.So(transformer.Enhance("..#/#.#/.##"), should.Resemble, "#..#/..../..../#..#")
	this.So(transformer.Enhance(".#./#../###"), should.Resemble, "#..#/..../..../#..#")
	this.So(transformer.Enhance("##./#.#/#.."), should.Resemble, "#..#/..../..../#..#")
}

func (this *Stuff) TestSplittingOfGridIntoPatterns() {
	this.So(SplitPatterns("../.."), should.Resemble, []string{"../.."})
	this.So(SplitPatterns(".../.../..."), should.Resemble, []string{".../.../..."})
	this.So(SplitPatterns("##../##../..##/..##"), should.Resemble,
		[]string{"##/##", "../..", "../..", "##/##"})
	this.So(SplitPatterns(""+
		"....#./"+
		".#..#./"+
		"....#./"+
		".....#/"+
		"###.#./"+
		"...#..",
	), should.Resemble, []string{
		"../.#",
		"../..",
		"#./#.",
		"../..",
		"../..",
		"#./.#",
		"##/..",
		"#./.#",
		"#./..",
	})
}

func (this *Stuff) TestReassemblePatternsIntoGrid() {
	this.So(ReassembleGrid("ab/ef", "cd/gh", "ij/mn", "kl/op"), should.Equal, "abcd\nefgh\nijkl\nmnop\n")
	this.So(ReassembleGrid("abc/ghi/mno", "def/jkl/pqr", "stu/yz0/456", "vwx/123/789"), should.Equal,
		"abcdef\nghijkl\nmnopqr\nstuvwx\nyz0123\n456789\n")
	this.So(ReassembleGrid(
		"1111/1111/1111/1111",
		"2222/2222/2222/2222",
		"3333/3333/3333/3333",
		"4444/4444/4444/4444",
	), should.Resemble, "11112222\n11112222\n11112222\n11112222\n33334444\n33334444\n33334444\n33334444\n")
}

func (this *Stuff) TestExampleRules() {
	exampleRules := RegisterEnhancementRules(
		"../.# => ##./#../...",
		".#./..#/### => #..#/..../..../#..#",
	)
	this.So(CountFractalPixels(exampleRules, 2), should.Equal, 12)
}
func (this *Stuff) TestPart1() {
	this.So(CountFractalPixels(realRules, 5), should.Equal, 110)
}
func (this *Stuff) TestPart2() {
	this.So(CountFractalPixels(realRules, 18), should.Equal, 1277716)
}
