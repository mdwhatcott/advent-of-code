package day09

import (
	"unicode/utf8"
)

type Item int

const (
	OpenGroup    Item = iota
	CloseGroup
	OpenGarbage
	Garbage
	CloseGarbage
	IgnoreNext
)

const eof rune = -1

type Lexer struct {
	out        chan Item
	input      string
	checkpoint int
	horizon    int
	width      int
}

func LexItems(input string) chan Item {
	lexer := &Lexer{input: input, out: make(chan Item)}
	go lexer.Run()
	return lexer.out
}

func (this *Lexer) Backup() {
	this.horizon -= this.width
}

func (this *Lexer) Next() (r rune) {
	if this.horizon >= len(this.input) {
		return eof
	}
	r, this.width = utf8.DecodeRuneInString(this.input[this.horizon:])
	this.horizon += this.width
	return r
}

func (this *Lexer) Peek() rune {
	r := this.Next()
	this.Backup()
	return r
}

func (this *Lexer) Emit(lexed Item) {
	this.out <- lexed
	this.checkpoint = this.horizon
}

func (this *Lexer) Run() {
	for state := lexStream; state != nil; {
		state = state(this)
	}
	close(this.out)
}

type StateFunction func(*Lexer) StateFunction

func lexStream(lexer *Lexer) StateFunction {
	for {
		switch lexer.Peek() {
		case eof:
			return nil
		case '!':
			return lexIgnoreNext(lexer)
		case '{':
			return lexOpenGroup(lexer)
		case '<':
			return lexOpenGarbage(lexer)
		case '}':
			return lexCloseGroup(lexer)
		default:
			lexer.Next()
		}
	}
}
func lexIgnoreNext(lexer *Lexer) StateFunction {
	lexer.horizon += 2
	lexer.Emit(IgnoreNext)
	return lexStream(lexer)
}
func lexOpenGroup(lexer *Lexer) StateFunction {
	lexer.horizon += len("{")
	lexer.Emit(OpenGroup)
	return lexStream(lexer)
}
func lexCloseGroup(lexer *Lexer) StateFunction {
	lexer.horizon += len("}")
	lexer.Emit(CloseGroup)
	return lexStream(lexer)
}
func lexOpenGarbage(lexer *Lexer) StateFunction {
	lexer.horizon += len("<")
	lexer.Emit(OpenGarbage)
	return lexGarbageStream(lexer)
}
func lexGarbageStream(lexer *Lexer) StateFunction {
	for {
		switch lexer.Peek() {
		case eof:
			return nil
		case '!':
			return lexIgnoreNextInGarbage(lexer)
		case '>':
			return lexCloseGarbage(lexer)
		default:
			return lexGarbage(lexer)
		}
	}
}
func lexIgnoreNextInGarbage(lexer *Lexer) StateFunction {
	lexer.horizon += 2
	lexer.Emit(IgnoreNext)
	return lexGarbageStream(lexer)
}
func lexGarbage(lexer *Lexer) StateFunction {
	lexer.horizon++
	lexer.Emit(Garbage)
	return lexGarbageStream(lexer)
}
func lexCloseGarbage(lexer *Lexer) StateFunction {
	lexer.horizon += len(">")
	lexer.Emit(CloseGarbage)
	return lexStream(lexer)
}
