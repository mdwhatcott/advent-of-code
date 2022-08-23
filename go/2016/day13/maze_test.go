package main

import (
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"
)

func TestExample(t *testing.T) {
	queue := NewLocationQueue()
	assert.Error(t).So(BreadthFirstSearch(queue, 10, 7, 4), should.Equal, 11)
}
