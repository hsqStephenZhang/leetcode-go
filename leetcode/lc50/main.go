package main

import (
	"fmt"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func myPow(x float64, n int) float64 {
	power := abs(n)
	sum := 1.0
	base := x
	for power >= 1 {
		if power&1 == 1 {
			sum = sum * base
		}
		base = base * base
		power = power >> 1
	}
	if n < 0 {
		return 1 / sum
	}
	return sum
}

func main() {
	fmt.Println(myPow(2.0, -2))
}
