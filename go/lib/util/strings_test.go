package util

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestAnagram(t *testing.T) {
	should.So(t, Anagram("a", "a"), should.BeTrue)
	should.So(t, Anagram("a", "b"), should.BeFalse)
	should.So(t, Anagram("aabb", "bbaa"), should.BeTrue)
}

func TestLevenshtein(t *testing.T) {
	should.So(t, Levenshtein("", ""), should.Equal, 0)
	should.So(t, Levenshtein("a", ""), should.Equal, 1)
	should.So(t, Levenshtein("", "a"), should.Equal, 1)
	should.So(t, Levenshtein("rosettacode", "raisethysword"), should.Equal, 8)
}
