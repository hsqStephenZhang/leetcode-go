package main

import "fmt"

var letterMap = []string{
	"",     // 0
	"",     // 1
	"abc",  // 2
	"def",  // 3
	"ghi",  // 4
	"jkl",  // 5
	"mno",  // 6
	"pqrs", // 7
	"tuv",  // 8
	"wxyz", // 9
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	ans := []string{}
	inner(digits, 0, "", &ans)
	return ans
}

func inner(digits string, index int, path string, ans *[]string) {
	if index == len(digits) {
		*ans = append(*ans, path)
		return
	}

	for _, v := range letterMap[digits[index]-'0'] {
		inner(digits, index+1, path+string(v), ans)
	}
}

func main() {
	fmt.Println(letterCombinations(""))
}
