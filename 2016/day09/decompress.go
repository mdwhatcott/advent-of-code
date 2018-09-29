package main

import (
	"fmt"
	"strings"
)

// Credits: https://github.com/letsdrinktea/AoC-2016/blob/master/Day9.java
// Credits: http://pastebin.com/wqkTGAvg
func GetUncompressedSize(s string, v2 bool) int {
	return getSize(s, 0, len(s), v2)
}
func getSize(input string, start, length int, v2 bool) (size int) {
	for c := start; c < start+length; {
		if input[c] == '(' {
			end := c + strings.Index(input[c:], ")")
			var length, repeat int
			fmt.Sscanf(input[c:end], "(%dx%d)", &length, &repeat)
			c = end + 1
			if v2 {
				size += repeat * getSize(input, c, length, v2)
			} else {
				size += repeat * length
			}
			c += length
		} else {
			size++
			c++
		}
	}
	return size
}
