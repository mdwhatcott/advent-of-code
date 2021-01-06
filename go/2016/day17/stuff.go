package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"hash"
)

func Navigate(passcode string) (shortest string, longest int) {
	hasher := md5.New()
	queue := NewLocationQueue()
	origin := NewLocation(1, 1, "")
	queue.Enqueue(origin)

	for queue.Length() > 0 {
		current := queue.Dequeue()
		if current.IsDestination(4, 4) {
			if shortest == "" {
				shortest = current.Directions
			}
			longest = len(current.Directions)
		} else {
			for _, adjacent := range current.AdjacentOpenRooms(passcode, hasher) {
				queue.Enqueue(adjacent)
			}
		}
	}
	return shortest, longest
}

/**************************************************************************/

type LocationQueue struct {
	locations []*Location
}

func NewLocationQueue() *LocationQueue {
	return &LocationQueue{}
}
func (this *LocationQueue) Length() int {
	return len(this.locations)
}

func (this *LocationQueue) Enqueue(location *Location) {
	this.locations = append(this.locations, location)
}

func (this *LocationQueue) Dequeue() *Location {
	l := this.locations[0]
	this.locations[0] = nil
	this.locations = this.locations[1:]
	return l
}

/**************************************************************************/

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

func (this *Location) AdjacentOpenRooms(passcode string, hasher hash.Hash) []*Location {
	hasher.Reset()
	hasher.Write([]byte(passcode + this.Directions))
	sum := hasher.Sum(nil)
	digest := hex.EncodeToString(sum)

	up := digest[0]
	down := digest[1]
	left := digest[2]
	right := digest[3]

	rooms := []*Location{}
	if bytes.IndexByte(open, up) > -1 && this.Y > 1 {
		rooms = append(rooms, NewLocation(this.X, this.Y-1, this.Directions+"U"))
	}
	if bytes.IndexByte(open, down) > -1 && this.Y < 4 {
		rooms = append(rooms, NewLocation(this.X, this.Y+1, this.Directions+"D"))
	}
	if bytes.IndexByte(open, left) > -1 && this.X > 1 {
		rooms = append(rooms, NewLocation(this.X-1, this.Y, this.Directions+"L"))
	}
	if bytes.IndexByte(open, right) > -1 && this.X < 4 {
		rooms = append(rooms, NewLocation(this.X+1, this.Y, this.Directions+"R"))
	}
	return rooms
}

var open = []byte("bcdef")
