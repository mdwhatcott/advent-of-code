package util

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"advent/lib/parse"
)

func InputCharacters() (all []string) {
	for _, c := range InputString() {
		all = append(all, string(c))
	}
	return all
}

func InputBytes() []byte {
	_, path, _, _ := runtime.Caller(0)
	this := path
	for caller := 1; path == this; caller++ {
		_, path, _, _ = runtime.Caller(caller)
	}
	pattern := regexp.MustCompile(`advent-of-code/go/(\d{4})/day(\d{2})/`)
	matches := pattern.FindStringSubmatch(path)
	year, day := matches[1], matches[2]
	for strings.Contains(path, "advent-of-code") {
		path = filepath.Dir(path)
	}
	path = filepath.Join(path, "advent-of-code-inputs", year, day) + ".txt"
	content, err := os.ReadFile(path)
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
	return parse.Ints(strings.Split(InputString(), sep))
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
