package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	normal := 0
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			normal += customers[i]
		}
	}

	ans := normal

	for i := 0; i < minutes; i++ {
		if grumpy[i] == 1 {
			normal += customers[i]
		}
	}

	ans = max(ans, normal)

	for i := minutes; i < len(customers); i++ {
		if grumpy[i-minutes] == 1 {
			normal -= customers[i-minutes]
		}
		if grumpy[i] == 1 {
			normal += customers[i]
		}
		ans = max(ans, normal)
	}

	return ans
}

func main() {
	// customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
	// grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
	// minutes := 3
	customers := []int{1}
	grumpy := []int{0}
	minutes := 1
	fmt.Println(maxSatisfied(customers, grumpy, minutes))
}
