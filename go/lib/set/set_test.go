package set

import (
	"sort"
	"testing"

	"advent/lib/set/internal/assert"
)

func TestCreation(t *testing.T) {
	assert.That(t, len(New[int](0))).Equals(0)
	assert.That(t, len(From[int]())).Equals(0)
	assert.That(t, From[int]().Len()).Equals(0)
}
func TestContains(t *testing.T) {
	assert.That(t, From[int](1).Contains(1)).IsTrue()
	assert.That(t, From[int]().Contains(1)).IsFalse()
}
func TestAdd(t *testing.T) {
	set := New[int](0)
	set.Add(1, 2, 3)
	assert.That(t, set.Contains(1)).IsTrue()
	assert.That(t, set.Contains(2)).IsTrue()
	assert.That(t, set.Contains(3)).IsTrue()
	assert.That(t, set.Len()).Equals(3)
}
func TestRemove(t *testing.T) {
	set := From[int](1, 2, 3)
	set.Remove(2)
	assert.That(t, set.Contains(1)).IsTrue()
	assert.That(t, set.Contains(2)).IsFalse()
	assert.That(t, set.Contains(3)).IsTrue()
	assert.That(t, set.Len()).Equals(2)
}
func TestClear(t *testing.T) {
	set := From[int](1, 2, 3)
	set.Clear()
	assert.That(t, set.Contains(1)).IsFalse()
	assert.That(t, set.Contains(2)).IsFalse()
	assert.That(t, set.Contains(3)).IsFalse()
	assert.That(t, set.Len()).Equals(0)
}
func TestSlice(t *testing.T) {
	set := From[int](1, 2, 3, 4, 5)
	items := set.Slice()
	sort.Slice(items, func(i, j int) bool {
		return items[i] < items[j]
	})
	assert.That(t, set.Len()).Equals(5)
	assert.That(t, len(items)).Equals(5)
	assert.That(t, items).Equals([]int{1, 2, 3, 4, 5})
}
func TestEqual(t *testing.T) {
	assert.That(t, From[int](1, 2, 3).Equal(From[int](3, 2, 1))).IsTrue()
	assert.That(t, From[int](1, 2).Equal(From[int](3, 2, 1))).IsFalse()
	assert.That(t, From[int](1, 2, 2).Equal(From[int](1, 2, 3))).IsFalse()
}
func TestIsSubset(t *testing.T) {
	assert.That(t, From[int](1, 2, 3).IsSubset(From[int](1, 2, 3, 4, 5))).IsTrue()
	assert.That(t, From[int](4, 5, 6).IsSubset(From[int](1, 2, 3, 4, 5))).IsFalse()
}
func TestIsSuperset(t *testing.T) {
	assert.That(t, From[int](1, 2, 3, 4, 5).IsSuperset(From[int](1, 2, 3))).IsTrue()
	assert.That(t, From[int](1, 2, 3, 4, 5).IsSuperset(From[int](4, 5, 6))).IsFalse()
}
func TestUnion(t *testing.T) {
	assert.That(t, From[int](1, 2, 3).Union(From[int](1, 2, 3))).Equals(From[int](1, 2, 3))
	assert.That(t, From[int](1, 2, 3).Union(From[int](2, 3, 4))).Equals(From[int](1, 2, 3, 4))
	assert.That(t, From[int](1, 2, 3).Union(From[int](4, 5, 6))).Equals(From[int](1, 2, 3, 4, 5, 6))
}
func TestIntersection(t *testing.T) {
	assert.That(t, From[int](1, 2, 3).Intersection(From[int](4, 5, 6))).Equals(From[int]())
	assert.That(t, From[int](1, 2, 3).Intersection(From[int](2, 3, 4))).Equals(From[int](2, 3))
}
func TestDifference(t *testing.T) {
	assert.That(t, From[int](1, 2, 3).Difference(From[int](4, 5, 6))).Equals(From[int](1, 2, 3))
	assert.That(t, From[int](1, 2, 3).Difference(From[int](2, 3))).Equals(From[int](1))
}
func TestSymmetricDifference(t *testing.T) {
	assert.That(t, From[int](1, 2, 3).SymmetricDifference(From[int](4, 5, 6))).Equals(From[int](1, 2, 3, 4, 5, 6))
	assert.That(t, From[int](1, 2, 3).SymmetricDifference(From[int](2, 3, 4))).Equals(From[int](1, 4))
}
