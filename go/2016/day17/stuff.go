package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"hash"

	"github.com/mdwhatcott/go-collections/queue"
)

func Navigate(passcode string) (shortest string, longest int) {
	hasher := md5.New()
	q := queue.New[*Location](0)
	q.Enqueue(NewLocation(1, 1, ""))

	for !q.Empty() {
		current := q.Dequeue()
		if current.IsDestination(4, 4) {
			if shortest == "" {
				shortest = current.Directions
			}
			longest = len(current.Directions)
		} else {
			for _, adjacent := range current.AdjacentOpenRooms(passcode, hasher) {
				q.Enqueue(adjacent)
			}
		}
	}
	return shortest, longest
}

type Location struct {
	X, Y       int
	Directions string
}

func NewLocation(x, y int, directions string) *Location {
	return &Location{X: x, Y: y, Directions: directions}
}

func (this *Location) IsDestination(x, y int) bool {
	return this.X == x && this.Y == y
}

func (this *Location) AdjacentOpenRooms(passcode string, hasher hash.Hash) (results []*Location) {
	hasher.Reset()
	hasher.Write([]byte(passcode + this.Directions))
	sum := hasher.Sum(nil)
	digest := hex.EncodeToString(sum)

	up := digest[0]
	down := digest[1]
	left := digest[2]
	right := digest[3]

	if bytes.IndexByte(open, up) > -1 && this.Y > 1 {
		results = append(results, NewLocation(this.X, this.Y-1, this.Directions+"U"))
	}
	if bytes.IndexByte(open, down) > -1 && this.Y < 4 {
		results = append(results, NewLocation(this.X, this.Y+1, this.Directions+"D"))
	}
	if bytes.IndexByte(open, left) > -1 && this.X > 1 {
		results = append(results, NewLocation(this.X-1, this.Y, this.Directions+"L"))
	}
	if bytes.IndexByte(open, right) > -1 && this.X < 4 {
		results = append(results, NewLocation(this.X+1, this.Y, this.Directions+"R"))
	}
	return results
}

var open = []byte("bcdef")
