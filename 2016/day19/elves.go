package main

func WinningElf2(count int) int {
	q := new(Queue)

	for x := 0; x < count; x++ {
		q.Enqueue(&Elf{Position: x + 1, Presents: 1})
	}

	for q.Len() > 1 {
		b := q.DequeueAcross()
		a := q.Dequeue()
		a.Presents += b.Presents
		//fmt.Println(a.Position, "takes" , b.Presents, "presents from", b.Position)
		q.Enqueue(a)
	}

	return q.Dequeue().Position
}
func WinningElf(count int) int {
	q := new(Queue)

	for x := 0; x < count; x++ {
		q.Enqueue(&Elf{Position: x + 1, Presents: 1})
	}

	for q.Len() > 1 {
		a := q.Dequeue()
		b := q.Dequeue()
		a.Presents += b.Presents
		q.Enqueue(a)
	}

	return q.Dequeue().Position
}

type Elf struct {
	Position int
	Presents int
}

type Queue struct {
	items []*Elf
}

func (this *Queue) Len() int {
	return len(this.items)
}

func (this *Queue) Enqueue(item *Elf) {
	this.items = append(this.items, item)
}

func (this *Queue) DequeueAcross() *Elf {
	length := len(this.items)
	i := length / 2
	e := this.items[i]
	this.items[i] = nil
	this.items = append(this.items[:i], this.items[i+1:]...)
	return e
}

func (this *Queue) Dequeue() *Elf {
	if len(this.items) == 0 {
		return nil
	}
	e := this.items[0]
	this.items[0] = nil
	this.items = this.items[1:]
	return e
}
