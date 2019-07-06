package main

import (
	"bufio"
	"sort"

	"advent/2016/util/pair"
)

type RepetitionCode struct {
	scanner *bufio.Scanner
	length  int
	columns []map[rune]int
}

func NewRepetitionCode(scanner *bufio.Scanner, length int) *RepetitionCode {
	columns := make([]map[rune]int, length)
	for x := 0; x < length; x++ {
		columns[x] = make(map[rune]int)
	}
	return &RepetitionCode{
		scanner: scanner,
		length:  length,
		columns: columns,
	}
}

func (this *RepetitionCode) DecodeFrequent() (answer string) {
	this.gatherCharactersIntoColumns()

	for _, column := range this.columns {
		frequency := pair.RankByFrequency(column)
		answer += string(frequency[0].Key)
	}
	return answer
}

func (this *RepetitionCode) DecodeInfrequent() (answer string) {
	this.gatherCharactersIntoColumns()

	for _, column := range this.columns {
		frequency := pair.RankByFrequency(column)
		sort.Sort(frequency)
		answer += string(frequency[0].Key)
	}
	return answer
}

func (this *RepetitionCode) gatherCharactersIntoColumns() {
	for this.scanner.Scan() {
		for i, c := range this.scanner.Text() {
			this.columns[i][c]++
		}
	}
}
