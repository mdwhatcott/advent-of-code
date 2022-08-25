package util

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestCompareSuite(t *testing.T) {
	suite.Run(&CompareSuite{T: suite.New(t)}, suite.Options.UnitTests())
}

type CompareSuite struct {
	*suite.T
}

func (this *CompareSuite) TestSum() {
	this.So(Sum(3, 2, 1), should.Equal, 6)
}

func (this *CompareSuite) TestMin() {
	this.So(Min(3, 2, 1), should.Equal, 1)
	this.So(Min(3, -1, 2, 1), should.Equal, -1)
}

func (this *CompareSuite) TestMax() {
	this.So(Max(3, 2, 1), should.Equal, 3)
	this.So(Max(-3, -1, -2), should.Equal, -1)
}
