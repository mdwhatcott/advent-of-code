package advent

import "strings"

type System struct {
	Satellites []*System
	Label      string
	Magnitude  int
}

func (this *System) Insert(center, satellite string) bool {
	if center == this.Label {
		this.Satellites = append(this.Satellites, &System{
			Label:     satellite,
			Magnitude: this.Magnitude + 1,
		})
		return true
	}
	for _, system := range this.Satellites {
		inserted := system.Insert(center, satellite)
		if inserted {
			return true
		}
	}
	return false
}

func (this *System) OrbitalChecksum() (sum int) {
	for _, system := range this.Satellites {
		sum += system.OrbitalChecksum()
	}
	return sum + this.Magnitude
}

func assembleOrbitalSystem(lines []string) *System {
	system := &System{Label: "COM", Magnitude: 0}

	remaining := map[string]bool{}
	inserted := map[string]bool{}

	for {
		for _, line := range lines {
			if inserted[line] {
				continue
			}
			parts := strings.Split(line, ")")
			center := parts[0]
			satellite := parts[1]
			ok := system.Insert(center, satellite)
			if !ok {
				remaining[line] = true
			} else {
				inserted[line] = true
				delete(remaining, line)
			}
		}
		if len(remaining) == 0 {
			break
		}
	}

	return system
}
