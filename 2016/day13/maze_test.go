package main

import (
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func TestExample(t *testing.T) {
	queue := NewLocationQueue()
	assertions.New(t).So(BreadthFirstSearch(queue, 10, 7, 4), should.Equal, 11)
}
