package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash"
	"strings"
)

type Generator struct {
	salt    string
	hasher  hash.Hash
	cache   map[int]string
	stretch bool
}

func NewGenerator(salt string) *Generator {
	return &Generator{
		salt:   salt,
		hasher: md5.New(),
		cache:  make(map[int]string),
	}
}

func (this *Generator) IndexOfKey(target int) int {
	x := 0
	keys := []int{}

	for {
		if this.IsKey(x) {
			keys = append(keys, x)
			if len(keys) == target {
				return x
			}
		}
		x++
	}
}

func (this *Generator) IsKey(i int) bool {
	digest := this.MD5(i)

	tripleChar := this.firstRunOfThree(digest)
	if tripleChar == "" {
		return false
	}

	start := i + 1
	stop := start + 1000

	for x := start; x < stop; x++ {
		if this.hasRunOfFive(this.MD5(x), tripleChar) {
			return true
		}
	}

	return false
}

func (this *Generator) firstRunOfThree(digest string) string {
	for x := 0; x < len(digest)-2; x++ {
		if digest[x] == digest[x+1] && digest[x] == digest[x+2] {
			return digest[x : x+1]
		}
	}
	return ""
}

func (this *Generator) hasRunOfFive(digest string, char string) bool {
	return strings.Contains(digest, strings.Repeat(char, 5))
}

func (this *Generator) MD5(i int) string {
	if cached, found := this.cache[i]; found {
		return cached
	}
	base := this.md5Index(i)

	if this.stretch {
		for x := 0; x < 2016; x++ {
			base = this.md5(base)
		}
	}
	this.cache[i] = base
	return base
}
func (this *Generator) md5Index(i int) string {
	return this.md5(fmt.Sprintf("%s%d", this.salt, i))
}
func (this *Generator) md5(i string) string {
	this.hasher.Reset()
	this.hasher.Write([]byte(i))
	sum := this.hasher.Sum(nil)
	return hex.EncodeToString(sum)
}
