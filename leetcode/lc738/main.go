package main

import (
	"fmt"
	"strconv"
)

func monotoneIncreasingDigits(n int) int {
	bits := []byte(strconv.Itoa(n))
	length := len(bits)
	for i := length - 1; i >= 1; i-- {
		if bits[i] < bits[i-1] {
			bits[i-1]--
			for j := i; j < length; j++ {
				bits[j] = '9'
			}
		}
	}
	res, _ := strconv.Atoi(string(bits))
	return res
}

func main() {
	fmt.Println(monotoneIncreasingDigits(10))
}
