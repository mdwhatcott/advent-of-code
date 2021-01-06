package main

import (
	"bytes"
	"strconv"
	"strings"
	"unicode"

	"advent/2016/util/pair"
)

type EncryptedRoom struct {
	EncryptedName string
	SectorID      int
	Valid         bool

	letters        map[rune]int
	parsedChecksum string
	actualChecksum string
}

func (this *EncryptedRoom) Decrypt() string {
	decrypted := make([]byte, len(this.EncryptedName))
	for i, c := range strings.Replace(this.EncryptedName, "-", " ", -1) {
		if c == ' ' {
			decrypted[i] = byte(c)
			continue
		}
		c += rune(this.SectorID % 26)
		if c > 'z' {
			c -= 26
		}
		decrypted[i] = byte(c)
	}
	return string(decrypted)
}

func ParseEncryptedRoom(value string) *EncryptedRoom {
	room := &EncryptedRoom{letters: make(map[rune]int)}
	values := strings.Split(value, "[")
	room.parsedChecksum = strings.TrimRight(values[1], "]")

	for i, c := range value {
		if unicode.IsLetter(c) {
			room.EncryptedName += string(c)
			room.letters[c]++
		} else if c == '-' {
			room.EncryptedName += string(c)
		} else if unicode.IsNumber(c) {
			room.SectorID, _ = strconv.Atoi(values[0][i:])
			break
		}
	}
	room.EncryptedName = strings.TrimRight(room.EncryptedName, "-")
	ranked := pair.RankByFrequency(room.letters)[:maxChecksumLength]
	room.actualChecksum = makeActualChecksum(ranked)
	room.Valid = room.actualChecksum == room.parsedChecksum
	return room
}

func makeActualChecksum(pairs pair.PairList) string {
	buffer := new(bytes.Buffer)
	for _, p := range pairs {
		buffer.WriteRune(p.Key)
	}
	return buffer.String()
}

const maxChecksumLength = 5
