package main

type IPv7Chunk struct {
	chunk    string
	hypernet bool

	aba []string
}

func makeChunks(ip string) (chunks []IPv7Chunk) {
	current := IPv7Chunk{hypernet: false}
	for _, c := range ip {
		if c == '[' {
			chunks = append(chunks, current)
			current = IPv7Chunk{hypernet: true}
		} else if c == ']' {
			chunks = append(chunks, current)
			current = IPv7Chunk{hypernet: false}
		} else {
			current.chunk += string(c)
		}
	}
	return append(chunks, current)
}

func IsTLSCompliant(ip string) bool {
	yes := false
	for _, chunk := range makeChunks(ip) {
		if chunk.hypernet && HasABBA(chunk.chunk) {
			return false
		} else if !chunk.hypernet && HasABBA(chunk.chunk) {
			yes = true
		}
	}
	return yes
}

func IsSSLCompliant(ip string) bool {
	chunks := makeChunks(ip)
	for _, super := range chunks {
		if !super.hypernet {
			for _, aba := range ExtractABA(super.chunk) {
				for _, hyper := range chunks {
					if hyper.hypernet && HasCorrespondingBAB(hyper.chunk, aba) {
						return true
					}
				}
			}
		}
	}
	return false
}

func HasABBA(s string) bool {
	if len(s) == 4 {
		return IsABBA(s)
	}
	for x := 0; x < len(s)-3; x++ {
		candidate := s[x : x+4]
		if IsABBA(candidate) {
			return true
		}
	}
	return false
}

func IsABBA(s string) bool {
	if s[0] == s[1] {
		return false
	}
	return s[0] == s[3] && s[1] == s[2]
}

func ExtractABA(s string) (all []string) {
	for x := 0; x < len(s)-2; x++ {
		candidate := s[x : x+3]
		if IsABA(candidate) {
			all = append(all, candidate)
		}
	}
	return all
}
func IsABA(s string) bool {
	return s[0] == s[2] && s[0] != s[1]
}

func HasCorrespondingBAB(i string, aba string) bool {
	bab := aba[1:] + aba[1:2]
	extracted := ExtractABA(i)
	for _, extract := range extracted {
		if extract == bab {
			return true
		}
	}
	return false
}
