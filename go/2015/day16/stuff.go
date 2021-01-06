package main

type AuntSue map[string]int

func (this AuntSue) Matches(evidence AuntSue) bool {
	for key := range evidence {
		if value, found := this[key]; found {
			if value != evidence[key] {
				return false
			}
		}
	}
	return true
}

func (this AuntSue) MatchesRange(evidence AuntSue) bool {
	for key := range evidence {
		if value, found := this[key]; found {
			switch key {
			case "cats", "trees":
				if value <= evidence[key] {
					return false
				}
			case "pomeranians", "goldfish":
				if value >= evidence[key] {
					return false
				}
			default:
				if value != evidence[key] {
					return false
				}
			}
		}
	}
	return true
}
