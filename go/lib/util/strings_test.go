package util

import (
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"
)

func TestAnagram(t *testing.T) {
	assert.So(t, Anagram("a", "a"), should.BeTrue)
	assert.So(t, Anagram("a", "b"), should.BeFalse)
	assert.So(t, Anagram("aabb", "bbaa"), should.BeTrue)
}

func TestLevenshtein(t *testing.T) {
	assert.So(t, Levenshtein("", ""), should.Equal, 0)
	assert.So(t, Levenshtein("a", ""), should.Equal, 1)
	assert.So(t, Levenshtein("", "a"), should.Equal, 1)
	assert.So(t, Levenshtein("rosettacode", "raisethysword"), should.Equal, 8)
}
