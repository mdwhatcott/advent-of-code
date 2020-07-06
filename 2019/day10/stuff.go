package advent

func scanField(rawField []string) (field AsteroidField) {
	for y, line := range rawField {
		for x, char := range line {
			if char == '.' {
				continue
			}
			field = append(field, NewAsteroid(x, y))
		}
	}
	return field
}

type Asteroid struct {
	X, Y float64
}

func NewAsteroid(x, y int) Asteroid {
	return Asteroid{
		X: float64(x),
		Y: float64(y),
	}
}

type AsteroidField []Asteroid

func Slope(a Asteroid, b Asteroid) float64 {
	return (b.X - a.X) / (b.Y - a.Y)
}

func CountVisible(field AsteroidField, asteroid Asteroid) (count int) {
	seen := map[float64]bool{}
	for _, candidate := range field {
		if candidate == asteroid {
			continue
		}
		slope := Slope(asteroid, candidate)
		if !seen[slope] {
			count++
		}
		seen[slope] = true
	}
	return count
}

func BestPlace(field AsteroidField) (MAX Asteroid) {
	var max int
	for _, asteroid := range field {
		visible := CountVisible(field, asteroid)
		if visible > max {
			max = visible
			MAX = asteroid
		}
	}
	return MAX
}
