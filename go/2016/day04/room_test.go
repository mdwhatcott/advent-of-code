package main

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestRoomValidationFixture(t *testing.T) {
	suite.Run(&RoomValidationFixture{T: suite.New(t)}, suite.Options.UnitTests())
}

type RoomValidationFixture struct {
	*suite.T
}

func (this *RoomValidationFixture) TestParsing() {
	this.So(ParseEncryptedRoom("aaaaa-bbb-z-y-x-123[abxyz]"), should.Equal, &EncryptedRoom{
		EncryptedName:  "aaaaa-bbb-z-y-x",
		SectorID:       123,
		Valid:          true,
		letters:        map[rune]int{'a': 5, 'b': 3, 'x': 1, 'y': 1, 'z': 1},
		parsedChecksum: "abxyz",
		actualChecksum: "abxyz",
	})

	this.So(ParseEncryptedRoom("a-b-c-d-e-f-g-h-987[abcde]"), should.Equal, &EncryptedRoom{
		EncryptedName:  "a-b-c-d-e-f-g-h",
		SectorID:       987,
		Valid:          true,
		letters:        map[rune]int{'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 1, 'f': 1, 'g': 1, 'h': 1},
		parsedChecksum: "abcde",
		actualChecksum: "abcde",
	})

	this.So(ParseEncryptedRoom("not-a-real-room-404[oarel]"), should.Equal, &EncryptedRoom{
		EncryptedName:  "not-a-real-room",
		SectorID:       404,
		Valid:          true,
		letters:        map[rune]int{'a': 2, 'n': 1, 'o': 3, 't': 1, 'r': 2, 'e': 1, 'l': 1, 'm': 1},
		parsedChecksum: "oarel",
		actualChecksum: "oarel",
	})

	this.So(ParseEncryptedRoom("totally-real-room-200[decoy]"), should.Equal, &EncryptedRoom{
		EncryptedName:  "totally-real-room",
		SectorID:       200,
		Valid:          false,
		letters:        map[rune]int{'t': 2, 'o': 3, 'a': 2, 'l': 3, 'y': 1, 'r': 2, 'e': 1, 'm': 1},
		parsedChecksum: "decoy",
		actualChecksum: "loart",
	})
}

func (this *RoomValidationFixture) TestDecryption() {
	this.So(ParseEncryptedRoom("aaaaa-bbb-z-y-x-1[abxyz]").Decrypt(), should.Equal, "bbbbb ccc a z y")
	this.So(ParseEncryptedRoom("aaaaa-bbb-z-y-x-27[abxyz]").Decrypt(), should.Equal, "bbbbb ccc a z y")
	this.So(ParseEncryptedRoom("aaaaa-bbb-z-y-x-2[abxyz]").Decrypt(), should.Equal, "ccccc ddd b a z")
}
