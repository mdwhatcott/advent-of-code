package advent

import (
	"encoding/hex"
	"strings"

	"github.com/mdwhatcott/advent-of-code-go-lib/parse"
	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func Part1() (valid int) {
	for _, entry := range strings.Split(util.InputString(), "\n\n") {
		if parsePassport(entry).HasRequiredFields() {
			valid++
		}
	}
	return valid
}
func Part2() (valid int) {
	for _, entry := range strings.Split(util.InputString(), "\n\n") {
		if parsePassport(entry).HasAllValidFields() {
			valid++
		}
	}
	return valid
}

type Passport map[string]string

func parsePassport(entry string) Passport {
	passport := make(Passport)
	fields := strings.Fields(entry)
	for _, field := range fields {
		parts := strings.Split(field, ":")
		passport[parts[0]] = parts[1]
	}
	return passport
}

func (this Passport) HasRequiredFields() bool {
	return len(this["byr"]) > 0 &&
		len(this["iyr"]) > 0 &&
		len(this["eyr"]) > 0 &&
		len(this["hgt"]) > 0 &&
		len(this["hcl"]) > 0 &&
		len(this["ecl"]) > 0 &&
		len(this["pid"]) > 0
}
func (this Passport) HasAllValidFields() bool {
	return this.hasValidBYR() &&
		this.hasValidECL() &&
		this.hasValidEYR() &&
		this.hasValidHCL() &&
		this.hasValidHGT() &&
		this.hasValidIYR() &&
		this.hasValidPID()
}
func (this Passport) hasValidBYR() bool {
	value := parse.Int(this["byr"])
	return 1920 <= value && value <= 2002
}
func (this Passport) hasValidIYR() bool {
	value := parse.Int(this["iyr"])
	return 2010 <= value && value <= 2020
}
func (this Passport) hasValidEYR() bool {
	value := parse.Int(this["eyr"])
	return 2020 <= value && value <= 2030
}
func (this Passport) hasValidHGT() bool {
	value := this["hgt"]
	if strings.HasSuffix(value, "cm") {
		parsed := parse.Int(strings.TrimSuffix(value, "cm"))
		return 150 <= parsed && parsed <= 193
	} else if strings.HasSuffix(value, "in") {
		parsed := parse.Int(strings.TrimSuffix(value, "in"))
		return 59 <= parsed && parsed <= 76
	} else {
		return false
	}
}
func (this Passport) hasValidHCL() bool {
	value := this["hcl"]
	if !strings.HasPrefix(value, "#") {
		return false
	}
	value = value[1:]
	if len(value) != 6 {
		return false
	}
	d, err := hex.DecodeString(value)
	return err == nil && len(d) == 3
}
func (this Passport) hasValidECL() bool {
	switch this["ecl"] {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	}
	return false
}
func (this Passport) hasValidPID() bool {
	value := this["pid"]

	if len(value) != 9 {
		return false
	}
	for _, c := range value {
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return true
		}
		return false
	}
	return true
}
