package main

type Character struct {
	name   string
	hits   int
	armor  int
	damage int
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

func NewCharacter(name string, hits, damage, armor int) *Character {
	return &Character{
		name:   name,
		hits:   hits,
		armor:  armor,
		damage: damage,
	}
}
