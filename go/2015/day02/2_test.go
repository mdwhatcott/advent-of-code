package advent

import (
	"testing"

	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func Test2_ParseDimensions(t *testing.T) {
	l, h, w := parseDimensions("1x2x3")
	should.So(t, []int{l, h, w}, should.Equal, []int{1, 2, 3})
}

func Test2_HowMuchPaper(t *testing.T) {
	should.So(t, howMuchPaper(2, 3, 4), should.Equal, 58)
	should.So(t, howMuchPaper(1, 1, 10), should.Equal, 43)
}

func Test2_HowMuchRibbon(t *testing.T) {
	should.So(t, howMuchRibbon(2, 3, 4), should.Equal, 34)
	should.So(t, howMuchRibbon(1, 1, 10), should.Equal, 14)
}

func Test2_Answer1(t *testing.T) {
	scanner := util.InputScanner()
	paper := 0
	ribbon := 0
	for scanner.Scan() {
		l, h, w := parseDimensions(scanner.Text())
		paper += howMuchPaper(l, h, w)
		ribbon += howMuchRibbon(l, h, w)
	}
	should.So(t, paper, should.Equal, 1606483)
	should.So(t, ribbon, should.Equal, 3842356)
	//t.Logf("How much wrapping paper? %d\n", paper)
	//t.Logf("How much ribbon? %d", ribbon)
}
