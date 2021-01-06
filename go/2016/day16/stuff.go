package main

import "bytes"

func DragonChecksum(a string, length int) string {
	return Checksum(ExpandDragon(a, length))
}

func Checksum(a string) string {
	var sum bytes.Buffer
	for x := 0; x < len(a)-1; x += 2 {
		if pair := a[x : x+2]; pair == "00" || pair == "11" {
			sum.WriteString("1")
		} else {
			sum.WriteString("0")
		}
	}
	if sum.Len()%2 == 0 {
		return Checksum(sum.String())
	}
	return sum.String()
}

func ExpandDragon(a string, length int) string {
	for len(a) < length {
		a = Dragon(a)
	}
	return a[:length]
}

func Dragon(a string) string {
	return a + "0" + Invert(Reverse(a))
}

func Reverse(a string) string {
	var b bytes.Buffer
	for x := len(a) - 1; x >= 0; x-- {
		b.WriteString(a[x : x+1])
	}
	return b.String()
}

func Invert(a string) string {
	var b bytes.Buffer
	for x := 0; x < len(a); x++ {
		i := a[x : x+1]
		if i == "1" {
			b.WriteString("0")
		} else {
			b.WriteString("1")
		}
	}
	return b.String()
}
