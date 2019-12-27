package set

import (
	"sort"
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestSetFixture(t *testing.T) {
    gunit.Run(new(SetFixture), t)
}

type SetFixture struct {
    *gunit.Fixture
}

func (this *SetFixture) Setup() {
	this.So(len(NewSet()), should.Equal, 0)
	this.So(NewSet().Len(), should.Equal, 0)
}

func (this *SetFixture) TestContains() {
	this.So(NewSet(1).Contains(1), should.BeTrue)
	this.So(NewSet().Contains(1), should.BeFalse)
}

func (this *SetFixture) TestAdd() {
	set := NewSet()
	set.Add(1, 2, 3)
	this.So(set.Contains(1), should.BeTrue)
	this.So(set.Contains(2), should.BeTrue)
	this.So(set.Contains(3), should.BeTrue)
	this.So(set.Len(), should.Equal, 3)
}

func (this *SetFixture) TestItems() {
	set := NewSet(1, 2, 3, 4, 5)
	items := set.Items()
	sort.Slice(items, func(i, j int) bool {
		return items[i].(int) < items[j].(int)
	})
	this.So(items, should.Resemble, []interface{}{1, 2, 3, 4, 5})
}

func (this *SetFixture) TestIsSubset() {
	this.So(NewSet(1, 2, 3).IsSubset(NewSet(1, 2, 3, 4, 5)), should.BeTrue)
	this.So(NewSet(4, 5, 6).IsSubset(NewSet(1, 2, 3, 4, 5)), should.BeFalse)
}

func (this *SetFixture) TestIsSuperset() {
	this.So(NewSet(1, 2, 3, 4, 5).IsSuperset(NewSet(1, 2, 3)), should.BeTrue)
	this.So(NewSet(1, 2, 3, 4, 5).IsSuperset(NewSet(4, 5, 6)), should.BeFalse)
}

func (this *SetFixture) TestUnion() {
	this.So(NewSet(1, 2, 3).Union(NewSet(1, 2, 3)), should.Resemble, NewSet(1, 2, 3))
	this.So(NewSet(1, 2, 3).Union(NewSet(2, 3, 4)), should.Resemble, NewSet(1, 2, 3, 4))
	this.So(NewSet(1, 2, 3).Union(NewSet(4, 5, 6)), should.Resemble, NewSet(1, 2, 3, 4, 5, 6))
}

func (this *SetFixture) TestIntersection() {
	this.So(NewSet(1, 2, 3).Intersection(NewSet(4, 5, 6)), should.Resemble, NewSet())
	this.So(NewSet(1, 2, 3).Intersection(NewSet(2, 3, 4)), should.Resemble, NewSet(2, 3))
}

func (this *SetFixture) TestDifference() {
	this.So(NewSet(1, 2, 3).Difference(NewSet(4, 5, 6)), should.Resemble, NewSet(1, 2, 3))
	this.So(NewSet(1, 2, 3).Difference(NewSet(2, 3)), should.Resemble, NewSet(1))
}

func (this *SetFixture) TestSymmetricDifference() {
	this.So(NewSet(1, 2, 3).SymmetricDifference(NewSet(4, 5, 6)), should.Resemble, NewSet(1, 2, 3, 4, 5, 6))
	this.So(NewSet(1, 2, 3).SymmetricDifference(NewSet(2, 3, 4)), should.Resemble, NewSet(1, 4))
}
