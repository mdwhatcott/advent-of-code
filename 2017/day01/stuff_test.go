package day01

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

func (this *Stuff) TestNeighborSums() {
	this.So(sumOfIdenticalNeighborDigits(chars("1122")), should.Equal, 3)
	this.So(sumOfIdenticalNeighborDigits(chars("1111")), should.Equal, 4)
	this.So(sumOfIdenticalNeighborDigits(chars("1234")), should.Equal, 0)
	this.So(sumOfIdenticalNeighborDigits(chars("91212129")), should.Equal, 9)
}

func (this *Stuff) TestOppositeSums() {
	this.So(sumOfIdenticalOppositeDigits(chars("1212")), should.Equal, 6)
	this.So(sumOfIdenticalOppositeDigits(chars("1221")), should.Equal, 0)
	this.So(sumOfIdenticalOppositeDigits(chars("123425")), should.Equal, 4)
	this.So(sumOfIdenticalOppositeDigits(chars("123123")), should.Equal, 12)
	this.So(sumOfIdenticalOppositeDigits(chars("12131415")), should.Equal, 4)
}

func chars(s string) (all []string) {
	for _, c := range s {
		all = append(all, string(c))
	}
	return all
}
