package starter

import (
	"sort"
	"strings"

	"github.com/mdwhatcott/funcy"
	"github.com/mdwhatcott/must/strconvmust"
)

func Part1(lines []string) (result int) {
	var hands []Hand
	for _, line := range lines {
		hands = append(hands, ParsePart1Hand(line))
	}
	sort.Slice(hands, func(i, j int) bool {
		iType := hands[i].Part1Type()
		jType := hands[j].Part1Type()
		if iType == jType {
			return string(hands[i]) > string(hands[j])
		}
		return iType > jType
	})
	for rank, hand := range hands {
		result += (rank + 1) * hand.Bid()
	}
	return result
}
func Part2(lines []string) (result int) {
	var hands []Hand
	for _, line := range lines {
		hands = append(hands, ParsePart2Hand(line))
	}
	sort.Slice(hands, func(i, j int) bool {
		iType := hands[i].Part2Type()
		jType := hands[j].Part2Type()
		if iType == jType {
			return string(hands[i]) > string(hands[j])
		}
		return iType > jType
	})
	for rank, hand := range hands {
		result += (rank + 1) * hand.Bid()
	}
	return result
}

// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2

var part1Replacer = strings.NewReplacer(
	"A", "A",
	"K", "B",
	"Q", "C",
	"J", "D",
	"T", "E",
	"9", "F",
	"8", "G",
	"7", "H",
	"6", "I",
	"5", "J",
	"4", "K",
	"3", "L",
	"2", "M",
	"1", "N",
)
var part1Reverser = strings.NewReplacer(
	"A", "A",
	"B", "K",
	"C", "Q",
	"D", "J",
	"E", "T",
	"F", "9",
	"G", "8",
	"H", "7",
	"I", "6",
	"J", "5",
	"K", "4",
	"L", "3",
	"M", "2",
	"N", "1",
)

var part2Replacer = strings.NewReplacer(
	"A", "A",
	"K", "B",
	"Q", "C",
	"J", "Z", // now the lowest
	"T", "E",
	"9", "F",
	"8", "G",
	"7", "H",
	"6", "I",
	"5", "J",
	"4", "K",
	"3", "L",
	"2", "M",
	"1", "N",
)
var part2Reverser = strings.NewReplacer(
	"A", "A",
	"B", "K",
	"C", "Q",
	"Z", "J", // from the lowest
	"E", "T",
	"F", "9",
	"G", "8",
	"H", "7",
	"I", "6",
	"J", "5",
	"K", "4",
	"L", "3",
	"M", "2",
	"N", "1",
)

func ParsePart1Hand(s string) Hand {
	hand, bid, _ := strings.Cut(s, " ")
	return Hand(part1Replacer.Replace(hand) + " " + bid)
}
func ParsePart2Hand(s string) Hand {
	hand, bid, _ := strings.Cut(s, " ")
	return Hand(part2Replacer.Replace(hand) + " " + bid)
}

type Hand []rune

func (this Hand) Part1Type() HandType {
	freq := funcy.Frequencies(funcy.MapValues(funcy.Frequencies(this[:5])))
	switch {
	case freq[5] == 1:
		return FiveOfAKind
	case freq[4] == 1:
		return FourOfAKind
	case freq[3] == 1 && freq[2] == 1:
		return FullHouse
	case freq[3] == 1:
		return ThreeOfAKind
	case freq[2] == 2:
		return TwoPair
	case freq[2] == 1:
		return OnePair
	default:
		return HighCard
	}
}
func (this Hand) Part2Type() HandType {
	/*
		5 J:                           5 of a kind
		4 J:                           5 of a kind
		3 J and 2 X:                   5 of a kind
		3 J and an X and a Y:          4 of a kind
		2 J and 3 of something:        5 of a kind
		2 J and 2 of something else:   4 of a kind
		2 J and 3 unique other cards:  3 of a kind
		1 J and 4 of something else:   5 of a kind
		1 J and 3 of something else:   4 of a kind
		1 J with 2 other pairs:        full house
		1 J with 1 other pair:         three of a kind
		1 J and 4 other cards:         one pair
	*/
	cardCounts := funcy.Frequencies(this[:5])
	switch cardCounts['Z'] {
	case 5:
		return FiveOfAKind
	case 4:
		return FiveOfAKind
	case 3:
		if len(cardCounts) == 2 {
			return FiveOfAKind
		}
		return FourOfAKind
	case 2:
		if len(cardCounts) == 2 {
			return FiveOfAKind
		}
		if len(cardCounts) == 3 {
			return FourOfAKind
		}
		return ThreeOfAKind
	case 1:
		pairCount := 0
		for _, value := range cardCounts {
			if value == 2 {
				pairCount++
			}
		}
		if pairCount == 2 {
			return FullHouse
		}
		if len(cardCounts) == 2 {
			return FiveOfAKind
		}
		if len(cardCounts) == 3 {
			return FourOfAKind
		}
		if len(cardCounts) == 4 {
			return ThreeOfAKind
		}
		return OnePair
	}
	return this.Part1Type()
}

func (this Hand) Bid() int {
	return strconvmust.Atoi(string(this[6:]))
}

func (this Hand) Part1String() string {
	hand, bid, _ := strings.Cut(string(this), " ")
	return part1Reverser.Replace(hand) + " " + bid
}
func (this Hand) Part2String() string {
	hand, bid, _ := strings.Cut(string(this), " ")
	return part2Reverser.Replace(hand) + " " + bid
}

type HandType int

func (this HandType) String() string {
	switch this {
	case FiveOfAKind:
		return "FiveOfAKind"
	case FourOfAKind:
		return "FourOfAKind"
	case FullHouse:
		return "FullHouse"
	case ThreeOfAKind:
		return "ThreeOfAKind"
	case TwoPair:
		return "TwoPair"
	case OnePair:
		return "OnePair"
	case HighCard:
		return "HighCard"
	}
	panic("invalid hand type")
}

const (
	FiveOfAKind HandType = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)
