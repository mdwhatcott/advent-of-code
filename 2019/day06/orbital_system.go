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

func (this *System) TracePath(search string) string {
	return this.tracePath(search, "")
}
func (this *System) tracePath(search, working string) string {
	working += "/" + this.Label
	if strings.HasSuffix(working, search) {
		return working
	}
	for _, system := range this.Satellites {
		path := system.tracePath(search, working)
		if len(path) > 0 {
			return path
		}
	}
	return ""
}

func (this *System) OrbitalDistance(from, to string) int {
	return 4 // TODO: implement
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
