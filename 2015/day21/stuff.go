package main

func Fight(a, b *Character) *Character {
	alternate := map[*Character]*Character{a: b, b: a}
	current := a
	for {
		other := alternate[current]
		current.Attack(other)
		if other.hits <= 0 {
			return current
		}
		current = other
	}
}

type Character struct {
	name   string
	hits   int
	damage int
	armor  int
}

func (this *Character) Attack(that *Character) {
	damage := this.damage - that.armor
	if damage <= 0 {
		damage = 1
	}
	that.hits -= damage
	//fmt.Printf("The %s deals %d-%d = %d damage; the %s goes down to %d hit points.\n",
	//	this.name, this.damage, that.armor, damage, that.name, that.hits)
}

func NewCharacter(name string, hits int, damage int, armor int) *Character {
	return &Character{
		name: name,
		hits:   hits,
		damage: damage,
		armor:  armor,
	}
}
