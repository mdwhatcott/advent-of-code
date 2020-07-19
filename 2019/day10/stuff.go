package advent

import "fmt"

func offsetField(original AsteroidField, x, y float64) (offset AsteroidField) {
	for _, asteroid := range original {
		offset = append(offset, Asteroid{
			X: asteroid.X + x,
			Y: asteroid.Y + y,
		})
	}
	return offset
}

func removeOrigin(original AsteroidField) (filtered AsteroidField) {
	for _, asteroid := range original {
		if asteroid.X != 0 || asteroid.Y != 0 {
			filtered = append(filtered, asteroid)
		}
	}
	return filtered

}

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

func (this Asteroid) String() string {
	return fmt.Sprintf("(%d,%d)", int(this.X), int(this.Y))
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
		if candidate.Y > asteroid.Y {
			break
		}
		slope := Slope(asteroid, candidate)
		fmt.Println(asteroid, candidate, slope)
		if !seen[slope] {
			count++
		}
		seen[slope] = true
	}

	for key := range seen {
		delete(seen, key)
	}

	for _, candidate := range field {
		if candidate == asteroid {
			continue
		}
		if candidate.Y <= asteroid.Y {
			continue
		}
		slope := Slope(asteroid, candidate)
		fmt.Println(asteroid, candidate, slope)
		if !seen[slope] {
			count++
		}
		seen[slope] = true
	}
	fmt.Println(seen)
	return count
}

func BestPlaceWithCount(field AsteroidField) (result PlaceCount) {
	for _, asteroid := range field {
		visible := CountVisible(field, asteroid)
		if visible > result.Count {
			result.Count = visible
			result.Place = asteroid
		}
	}
	return result
}

type PlaceCount struct {
	Place Asteroid
	Count int
}
