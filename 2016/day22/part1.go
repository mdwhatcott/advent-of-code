package main

func (this Disk) countViableSectorPairs() int {
	viable := 0
	for i, node1 := range this {
		for j, node2 := range this {
			if i == j {
				continue
			}
			if node1.Used == 0 {
				continue
			}
			if node2.Avail < node1.Used {
				continue
			}
			viable++
		}
	}
	return viable
}
