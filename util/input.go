package util

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

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

func InputString() string {
	return strings.TrimSpace(string(InputBytes()))
}

func InputLines() []string {
	return strings.Split(InputString(), "\n")
}

func InputScanner() *bufio.Scanner {
	return bufio.NewScanner(bytes.NewReader(InputBytes()))
}
