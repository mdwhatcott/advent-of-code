package main

import "strings"

type Keypad struct {
	current    string
	directions map[string]string
}

func NewKeypad(directions map[string]string) *Keypad {
	return &Keypad{current: "5", directions: directions}
}

func (this *Keypad) DeriveCode(lines string) (code string) {
	for _, instruction := range strings.Fields(lines) {
		code += this.FindDigit(instruction)
	}
	return code
}
func (this *Keypad) FindDigit(instructions string) string {
	for _, c := range instructions {
		this.current = this.directions[this.current+string(c)]
	}
	return this.current
}

/*
  1
 234
56789
 ABC
  D
*/
var thirteenKeyDirections = map[string]string{
	"1U": "1",
	"1L": "1",
	"1D": "3",
	"1R": "1",

	"2U": "2",
	"2L": "2",
	"2D": "6",
	"2R": "3",

	"3U": "1",
	"3L": "2",
	"3D": "7",
	"3R": "4",

	"4U": "4",
	"4L": "3",
	"4D": "8",
	"4R": "4",

	"5U": "5",
	"5L": "5",
	"5D": "5",
	"5R": "6",

	"6U": "2",
	"6L": "5",
	"6D": "A",
	"6R": "7",

	"7U": "3",
	"7L": "6",
	"7D": "B",
	"7R": "8",

	"8U": "4",
	"8L": "7",
	"8D": "C",
	"8R": "9",

	"9U": "9",
	"9L": "8",
	"9D": "9",
	"9R": "9",

	"AU": "6",
	"AL": "A",
	"AD": "A",
	"AR": "B",

	"BU": "7",
	"BL": "A",
	"BD": "D",
	"BR": "C",

	"CU": "8",
	"CL": "B",
	"CD": "C",
	"CR": "C",

	"DU": "B",
	"DL": "D",
	"DD": "D",
	"DR": "D",
}

/*
123
456
789
*/
var nineKeyDirections = map[string]string{
	"1U": "1",
	"1L": "1",
	"1D": "4",
	"1R": "2",

	"2U": "2",
	"2L": "1",
	"2D": "5",
	"2R": "3",

	"3U": "3",
	"3L": "2",
	"3D": "6",
	"3R": "3",

	"4U": "1",
	"4L": "4",
	"4D": "7",
	"4R": "5",

	"5U": "2",
	"5L": "4",
	"5D": "8",
	"5R": "6",

	"6U": "3",
	"6L": "5",
	"6D": "9",
	"6R": "6",

	"7U": "4",
	"7L": "7",
	"7D": "7",
	"7R": "8",

	"8U": "5",
	"8L": "7",
	"8D": "8",
	"8R": "9",

	"9U": "6",
	"9L": "8",
	"9D": "9",
	"9R": "9",
}
