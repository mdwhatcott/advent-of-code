package main

import "bytes"

type Row struct {
	tiles []bool
}

func ParseRow(raw string) *Row {
	r := &Row{}
	for _, c := range raw {
		r.tiles = append(r.tiles, c == '^')
	}
	return r
}

func (this *Row) String() string {
	var b bytes.Buffer
	for _, t := range this.tiles {
		if t {
			b.WriteString("^")
		} else {
			b.WriteString(".")
		}
	}
	return b.String()
}

func (this *Row) Safe() int {
	safe := 0
	for _, t := range this.tiles {
		if !t {
			safe++
		}
	}
	return safe
}

func (this *Row) Next() *Row {
	next := &Row{}
	for i := range this.tiles {
		next.tiles = append(next.tiles, this.NextIsTrap(i))
	}
	return next
}

func (this *Row) NextIsTrap(i int) bool {
	left := this.at(i - 1)
	center := this.at(i)
	right := this.at(i + 1)

	if left && center && !right {
		return true
	}
	if !left && center && right {
		return true
	}
	if left && !center && !right {
		return true
	}
	if !left && !center && right {
		return true
	}
	return false
}

func (this *Row) at(i int) bool {
	if i < 0 || i >= len(this.tiles) {
		return false
	}
	return this.tiles[i]
}
