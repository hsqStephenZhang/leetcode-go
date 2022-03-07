package main

import "fmt"

func minWindow(s string, t string) string {
	target := make([]int, 30*2)
	count := make([]int, 30*2)
	total := 0
	targetNum := 0
	for _, ch := range t {
		target[ch-'A']++
		if target[ch-'A'] == 1 {
			targetNum++
		}
	}

	startIdx := -1
	minLength := len(s) + 1

	right := 0
	left := 0
	for right < len(s) {
		count[s[right]-'A']++
		if count[s[right]-'A'] == target[s[right]-'A'] {
			total++
		}
		if total == targetNum {
			if right-left+1 < minLength {
				minLength = right - left + 1
				startIdx = left
			}
			for left < right {
				count[s[left]-'A']--
				if count[s[left]-'A'] == target[s[left]-'A']-1 {
					total--
					left += 1
					break
				}
				if right-left < minLength {
					minLength = right - left
					startIdx = left + 1
				}
				left += 1
			}
		}
		right += 1
	}
	if startIdx != -1 {
		return string(s[startIdx : startIdx+minLength])
	} else {
		return ""
	}
}

func main() {
	s := "cwaefgcf"
	t := "cae"
	fmt.Println(minWindow(s, t))
}
