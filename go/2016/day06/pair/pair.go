package pair

import "sort"

func RankByFrequency(characters map[rune]int) PairList {
	pairs := make(PairList, len(characters))
	i := 0
	for k, v := range characters {
		pairs[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pairs))
	return pairs
}

type Pair struct {
	Key   rune
	Value int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool {
	if p[i].Value < p[j].Value {
		return true
	}
	if p[i].Value == p[j].Value {
		return p[i].Key > p[j].Key
	}
	return false
}
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
