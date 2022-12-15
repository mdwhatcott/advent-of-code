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
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}
)

func sampleStacks() map[rune]*stack.Stack[rune] {
	return map[rune]*stack.Stack[rune]{
		'1': stackFrom("ZN"),
		'2': stackFrom("MCD"),
		'3': stackFrom("P"),
	}
}
func inputStacks() map[rune]*stack.Stack[rune] {
	//goland:noinspection SpellCheckingInspection
	return map[rune]*stack.Stack[rune]{
		'1': stackFrom("QMGCL"),
		'2': stackFrom("RDLCTFHG"),
		'3': stackFrom("VJFNMTWR"),
		'4': stackFrom("JFDVQP"),
		'5': stackFrom("NFMSLBT"),
		'6': stackFrom("RNVHCDP"),
		'7': stackFrom("HCT"),
		'8': stackFrom("GSJVZNHP"),
		'9': stackFrom("ZFHG"),
	}
}
func stackFrom(s string) *stack.Stack[rune] {
	result := stack.New[rune](0)
	for _, v := range s {
		result.Push(v)
	}
	return result
}

func TestDay05(t *testing.T) {
	should.So(t, Execute(sampleLines, NewQueueAsStack[rune](), sampleStacks()), should.Equal, "CMZ")
	should.So(t, Execute(inputLines, NewQueueAsStack[rune](), inputStacks()), should.Equal, "VCTFTJQCG")

	should.So(t, Execute(sampleLines, stack.New[rune](0), sampleStacks()), should.Equal, "MCD")
	should.So(t, Execute(inputLines, stack.New[rune](0), inputStacks()), should.Equal, "GCFGLDNJZ")
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

func Execute(instructions []string, tmp StackOrQueue[rune], stacks map[rune]*stack.Stack[rune]) string {
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
	return peekAll(stacks)
}
func peekAll(stacks map[rune]*stack.Stack[rune]) string {
	var buffer []rune
	for _, x := range "123456789"[:len(stacks)] {
		buffer = append(buffer, stacks[x].Peek())
	}
	return string(buffer)
}
