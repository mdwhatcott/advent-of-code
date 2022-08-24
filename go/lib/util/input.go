package util

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"runtime"
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
	_, file, _, _ := runtime.Caller(0)
	this := file
	for caller := 1; file == this; caller++ {
		_, file, _, _ = runtime.Caller(caller)
	}
	dir := filepath.Dir(file)
	input := filepath.Join(dir, "input.txt")
	content, err := os.ReadFile(input)
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
