package day05

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/go-collections/queue"
	"github.com/mdwhatcott/go-collections/stack"
	"github.com/mdwhatcott/testing/should"

	"advent/lib/util"
)

var (
	inputLines  = util.InputLines()
	sampleLines = []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
		"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}
)

func sampleStacks() map[rune]*stack.Stack[rune] {
	return map[rune]*stack.Stack[rune]{
		'1': stackFrom('Z', 'N'),
		'2': stackFrom('M', 'C', 'D'),
		'3': stackFrom('P'),
	}
}
func inputStacks() map[rune]*stack.Stack[rune] {
	//goland:noinspection SpellCheckingInspection
	return map[rune]*stack.Stack[rune]{
		'1': stackFrom([]rune("QMGCL")...),
		'2': stackFrom([]rune("RDLCTFHG")...),
		'3': stackFrom([]rune("VJFNMTWR")...),
		'4': stackFrom([]rune("JFDVQP")...),
		'5': stackFrom([]rune("NFMSLBT")...),
		'6': stackFrom([]rune("RNVHCDP")...),
		'7': stackFrom([]rune("HCT")...),
		'8': stackFrom([]rune("GSJVZNHP")...),
		'9': stackFrom([]rune("ZFHG")...),
	}
}

func stackFrom[T comparable](values ...T) *stack.Stack[T] {
	result := stack.New[T](0)
	for _, v := range values {
		result.Push(v)
	}
	return result
}

func TestDay05Part1(t *testing.T) {
	stacks := sampleStacks()
	Execute(sampleLines, NewQueueAsStack[rune](), stacks)
	should.So(t, PeekAll(stacks), should.Equal, []rune("CMZ"))

	stacks = inputStacks()
	Execute(inputLines, NewQueueAsStack[rune](), stacks)
	should.So(t, PeekAll(stacks), should.Equal, []rune("VCTFTJQCG"))
}
func TestDay05Part2(t *testing.T) {
	stacks := sampleStacks()
	Execute(sampleLines, stack.New[rune](0), stacks)
	should.So(t, PeekAll(stacks), should.Equal, []rune("MCD"))

	stacks = inputStacks()
	Execute(inputLines, stack.New[rune](0), stacks)
	should.So(t, PeekAll(stacks), should.Equal, []rune("GCFGLDNJZ"))
}

type StackOrQueue[T comparable] interface {
	Push(T)
	Pop() T
}
type QueueAsStack[T comparable] struct{ *queue.Queue[T] }

func NewQueueAsStack[T comparable]() *QueueAsStack[T] {
	return &QueueAsStack[T]{Queue: queue.New[T](0)}
}
func (q *QueueAsStack[T]) Push(t T) { q.Enqueue(t) }
func (q *QueueAsStack[T]) Pop() T   { return q.Dequeue() }

func Execute(instructions []string, tmp StackOrQueue[rune], stacks map[rune]*stack.Stack[rune]) {
	for _, instruction := range instructions {
		if !strings.HasPrefix(instruction, "move") {
			continue
		}
		fields := strings.Fields(instruction)
		moves := util.ParseInt(fields[1])
		from := stacks[rune(fields[3][0])]
		to := stacks[rune(fields[5][0])]
		for x := 0; x < moves; x++ {
			tmp.Push(from.Pop())
		}
		for x := 0; x < moves; x++ {
			to.Push(tmp.Pop())
		}
	}
}

func PeekAll(stacks map[rune]*stack.Stack[rune]) (result []rune) {
	for _, x := range "123456789"[:len(stacks)] {
		result = append(result, stacks[x].Peek())
	}
	return result
}
