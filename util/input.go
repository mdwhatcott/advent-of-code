package util

import (
	"bytes"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func InputCharacters() (all []string) {
	for _, c := range InputString() {
		all = append(all, string(c))
	}
	return all
}

func InputBytes() []byte {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Panic(err)
	}
	return content
}

func InputInt() int {
	i, _ := strconv.Atoi(InputString())
	return i
}

func InputInts(sep string) []int {
	return ParseInts(strings.Split(InputString(), sep))
}

func InputString() string {
	return strings.TrimSpace(string(InputBytes()))
}

func InputLines() []string {
	return strings.Split(InputString(), "\n")
}

func InputScanner() *Scanner {
	return NewScanner(bytes.NewReader(InputBytes()))
}
