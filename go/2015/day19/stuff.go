package main

import "strings"

type MoleculeMachine struct {
	replacements    []*Replacement
	transformations map[string]struct{}
}

func NewMoleculeMachine() *MoleculeMachine {
	return &MoleculeMachine{transformations: make(map[string]struct{})}
}

func (this *MoleculeMachine) RegisterReplacement(line string) {
	this.replacements = append(this.replacements, NewReplacement(line))
}

func (this *MoleculeMachine) Calibrate(start string) (unique []string) {
	for _, replacement := range this.replacements {
		this.replace(start, replacement)
	}
	return this.gatherResults()
}

func (this *MoleculeMachine) replace(start string, replacement *Replacement) {
	for replacement.SetSubject(start); replacement.Next(); {
		this.transformations[replacement.Transformation()] = struct{}{}
	}
}

func (this *MoleculeMachine) gatherResults() (unique []string) {
	for transformation := range this.transformations {
		unique = append(unique, transformation)
	}
	return unique
}

////////////////////////////////////////////////////////////////////////////////////

type Replacement struct {
	search  string
	replace string
	subject string
	latest  int
}

func NewReplacement(raw string) *Replacement {
	fields := strings.Fields(raw)
	return &Replacement{
		search:  fields[0],
		replace: fields[2],
	}
}

func (this *Replacement) SetSubject(subject string) {
	this.subject = subject
	this.latest = -1
}

func (this *Replacement) Next() bool {
	this.latest++
	next := strings.Index(this.subject[this.latest:], this.search)
	this.latest = this.latest + next
	return next > -1
}

func (this *Replacement) Transformation() string {
	prefix := this.subject[:this.latest]
	suffix := this.subject[this.latest:]
	suffix = strings.Replace(suffix, this.search, this.replace, 1)
	return prefix + suffix
}
