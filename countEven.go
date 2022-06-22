package main

func countEven(num int) int {
	sum := 0
	cpy := num / 10

	for cpy > 0 {
		sum += cpy % 10
		cpy /= 10
	}

	cpy = sum % 2

	return (num - cpy) / 2
}
