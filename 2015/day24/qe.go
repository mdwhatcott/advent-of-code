package main

func QuantumEntanglement(weights ...int) (product int) {
	if len(weights) == 0 {
		return 0
	}
	product = weights[0]
	for _, w := range weights[1:] {
		product *= w
	}
	return product
}
