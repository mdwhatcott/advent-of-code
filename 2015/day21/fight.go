package main

func Fight(a, b *Character) *Character {
	for {
		a.Attack(b)
		if b.hits <= 0 {
			return a
		}
		a, b = b, a
	}
}
