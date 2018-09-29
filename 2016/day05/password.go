package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"hash"
	"strconv"
	"strings"
)

type SequentialPassword struct {
	id       string
	counter  int
	hasher   hash.Hash
	password string
}

func (this *SequentialPassword) String() string {
	return this.password
}

func NewSequentialPassword(id string) *SequentialPassword {
	password := &SequentialPassword{id: id, hasher: md5.New()}
	password.compute()
	return password
}

func (this *SequentialPassword) compute() {
	for x := 0; x < 8; x++ {
		this.password += this.next()
	}
}

func (this *SequentialPassword) next() string {
	for {
		this.counter++
		this.hasher.Write([]byte(this.id + strconv.Itoa(this.counter)))
		sum := this.hasher.Sum(nil)
		this.hasher.Reset()

		if digest := hex.EncodeToString(sum); strings.HasPrefix(digest, "00000") {
			return digest[5:6]
		}
	}
}

/**************************************************************************/

type PositionalPassword struct {
	id       string
	counter  int
	hasher   hash.Hash
	password []byte
}

func NewPositionalPassword(id string) *PositionalPassword {
	password := &PositionalPassword{
		id:       id,
		password: bytes.Repeat([]byte(" "), 8),
		hasher:   md5.New(),
	}
	password.compute()
	return password
}

func (this *PositionalPassword) compute() {
	for bytes.Contains(this.password, []byte(" ")) {
		this.counter++
		this.hasher.Write([]byte(this.id + strconv.Itoa(this.counter)))
		sum := this.hasher.Sum(nil)
		this.hasher.Reset()

		digest := hex.EncodeToString(sum)
		if !bytes.HasPrefix([]byte(digest), []byte("00000")) {
			continue
		}
		position, err := strconv.Atoi(digest[5:6])
		if err != nil || position > 7 || this.password[position] != ' ' {
			continue
		}
		this.password[position] = digest[6]
	}
}

func (this *PositionalPassword) String() string {
	return string(this.password)
}
