package day12

import (
	"strings"

	"github.com/mdwhatcott/advent-of-code/go/lib/parse"
	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

var pipes = make(map[int]*Pipe)

func Answers() (part1, part2 int) {
	for scanner := util.InputScanner(); scanner.Scan(); {
		text := scanner.Text()
		text = strings.Replace(text, "<->", "", 1)
		text = strings.Replace(text, ",", "", -1)
		fields := strings.Fields(text)
		from := getPipe(fields[0])
		tos := fields[1:]
		for _, to := range tos {
			from.connect(getPipe(to))
		}
	}

	return len(group(zero())), countGroups()
}

func countGroups() (groups int) {
	seen := map[int]bool{}
	for _, pipe := range pipes {
		g := group(pipe)
		if !match(seen, g) {
			groups++
			for key := range g {
				seen[key] = true
			}
		}
	}
	return groups
}

func zero() *Pipe {
	for _, pipe := range pipes {
		if pipe.id == 0 {
			return pipe
		}
	}
	panic("Not found!")
}

func match(seen, group map[int]bool) bool {
	for key := range group {
		return seen[key]
	}
	return false
}

func group(start *Pipe) map[int]bool {
	queue := make(chan *Pipe, 1024)
	queue <- start
	seen := map[int]bool{}
	for len(queue) > 0 {
		popped := <-queue
		if seen[popped.id] {
			continue
		}
		seen[popped.id] = true
		for _, other := range popped.others {
			queue <- other
		}
	}
	return seen
}

func getPipe(name string) *Pipe {
	id := parse.Int(name)
	pipe := pipes[id]
	if pipe == nil {
		pipe = &Pipe{id: id}
		pipes[id] = pipe
	}
	return pipe
}

type Pipe struct {
	id     int
	others []*Pipe
}

func (this *Pipe) connect(other *Pipe) {
	this.others = append(this.others, other)
	other.others = append(other.others, this)
}
