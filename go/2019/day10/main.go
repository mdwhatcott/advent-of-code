package advent

import "github.com/mdwhatcott/advent-of-code/go/lib/util"

func Part1() interface{} {
	field := scanField(util.InputLines())
	best := BestPlaceWithCount(field)
	return best.Count
}

func Part2() interface{} {
	field := scanField(util.InputLines())
	best := BestPlaceWithCount(field)
	field = removeOrigin(offsetField(field, -best.Place.X, -best.Place.Y))
	blaster := NewBlaster(field)
	var asteroid Asteroid
	for x := 0; x < 200; x++ {
		aim := blaster.Aim()
		asteroid = blaster.Field()[aim]
		blaster.Fire(aim)
	}
	asteroid.X += best.Place.X
	asteroid.Y += best.Place.Y
	return int(asteroid.X*100 + asteroid.Y)
}
