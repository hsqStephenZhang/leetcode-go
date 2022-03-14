package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func partitionLabels(s string) []int {
	ans := []int{}

	ranges := make([]int, 27)
	for i := 0; i < len(s); i++ {
		ranges[s[i]-'a'] = i
	}
	prePos := 0
	nextPos := ranges[s[0]-'a']
	for i := 1; i < len(s); i++ {
		if i > nextPos {
			ans = append(ans, i-prePos)
			nextPos = ranges[s[i]-'a']
			prePos = i
		} else {
			nextPos = max(nextPos, ranges[s[i]-'a'])
		}
	}

	ans = append(ans, len(s)-prePos)

	return ans
}

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}
