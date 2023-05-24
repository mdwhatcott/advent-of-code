package day18

import (
	"testing"

	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func TestParseTree(t *testing.T) {
	raw := "[[[[1,2],[3,4]],[[5,6],[7,8]]],9]"
	tree := ParseTree(raw)
	should.So(t, tree.Left.Left.Left.Right.Value, should.Equal, 2)
	should.So(t, tree.Left.Left.Left.Right.Next.Value, should.Equal, 3)
	should.So(t, tree.String(), should.Equal, raw)
}
func TestExplode(t *testing.T) {
	raw := "[[[[[9,8],1],2],3],4]"
	tree := ParseTree(raw)
	Process(tree)
	should.So(t, tree.String(), should.Equal, "[[[[0,9],2],3],4]")
}
func TestSplit(t *testing.T) {
	raw := "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]"
	tree := ParseTree(raw)
	Process(tree)
	should.So(t, tree.String(), should.Equal, "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]")
}
func TestAdd(t *testing.T) {
	a := ParseTree("[[[[4,3],4],4],[7,[[8,4],9]]]")
	b := ParseTree("[1,1]")
	c := Add(a, b)
	should.So(t, c.String(), should.Equal, "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]")
}
func TestReduce(t *testing.T) {
	a := ParseTree("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]")
	b := Reduce(a)
	should.So(t, b.String(), should.Equal, "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")
}
func TestSum1(t *testing.T) {
	node := Sum([]string{"[1,1]", "[2,2]", "[3,3]", "[4,4]"})
	should.So(t, node.String(), should.Equal, "[[[[1,1],[2,2]],[3,3]],[4,4]]")
}
func TestSum2(t *testing.T) {
	node := Sum([]string{"[1,1]", "[2,2]", "[3,3]", "[4,4]", "[5,5]"})
	should.So(t, node.String(), should.Equal, "[[[[3,0],[5,3]],[4,4]],[5,5]]")
}
func TestSum3(t *testing.T) {
	node := Sum([]string{"[1,1]", "[2,2]", "[3,3]", "[4,4]", "[5,5]", "[6,6]"})
	should.So(t, node.String(), should.Equal, "[[[[5,0],[7,4]],[5,5]],[6,6]]")
}
func TestSum4(t *testing.T) {
	node := Sum([]string{
		"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
		"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
		"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
		"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
		"[7,[5,[[3,8],[1,4]]]]",
		"[[2,[2,2]],[8,[8,1]]]",
		"[2,9]",
		"[1,[[[9,3],9],[[9,0],[0,7]]]]",
		"[[[5,[7,4]],7],1]",
		"[[[[4,2],2],6],[8,7]]",
	})
	should.So(t, node.String(), should.Equal, "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]")
}
func TestMagnitude(t *testing.T) {
	should.So(t, ParseTree("[[1,2],[[3,4],5]]").Magnitude(), should.Equal, 143)
	should.So(t, ParseTree("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]").Magnitude(), should.Equal, 1384)
	should.So(t, ParseTree("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]").Magnitude(), should.Equal, 3488)
}
func TestSolveSum(t *testing.T) {
	sum := Sum([]string{
		"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
		"[[[5,[2,8]],4],[5,[[9,9],0]]]",
		"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
		"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
		"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
		"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
		"[[[[5,4],[7,7]],8],[[8,3],8]]",
		"[[9,3],[[9,9],[6,[4,9]]]]",
		"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
		"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
	})
	should.So(t, sum.Magnitude(), should.Equal, 4140)
}
func TestPart1(t *testing.T) {
	should.So(t, Sum(util.InputLines()).Magnitude(), should.Equal, 3691)
}
func TestMaxSumPair(t *testing.T) {
	should.So(t, MaxSumPair([]string{
		"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
		"[[[5,[2,8]],4],[5,[[9,9],0]]]",
		"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
		"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
		"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
		"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
		"[[[[5,4],[7,7]],8],[[8,3],8]]",
		"[[9,3],[[9,9],[6,[4,9]]]]",
		"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
		"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
	}), should.Equal, 3993)
}
func TestPart2(t *testing.T) {
	should.So(t, MaxSumPair(util.InputLines()), should.Equal, 4756)
}
