package main

import "fmt"

func findAnagrams(s string, p string) []int {
	length1 := len(s)
	length2 := len(p)
	if length1 < length2 {
		return []int{}
	}
	target := make([]int, 30)
	numKinds := 0
	for _, ch := range p {
		target[ch-'a']++
		if target[ch-'a'] == 1 {
			numKinds++
		}
	}
	ans := []int{}
	count := make([]int, 30)
	total := 0
	for _, ch := range s[:length2] {
		count[ch-'a']++
		if count[ch-'a'] == target[ch-'a'] {
			total++
		}
	}
	if total == numKinds {
		ans = append(ans, 0)
	}

	for idx, ch := range s[length2:] {
		count[ch-'a']++
		if count[ch-'a'] == target[ch-'a'] {
			total++
		}
		count[s[idx]-'a']--
		if count[s[idx]-'a'] == target[s[idx]-'a']-1 {
			total--
		}
		if total == numKinds {
			ans = append(ans, idx+1)
		}
	}
	return ans
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}
