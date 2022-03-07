package main

func checkInclusion(s1 string, s2 string) bool {
	length1 := len(s1)
	length2 := len(s2)
	if length1 > length2 {
		return false
	}
	target := make([]int, 30)
	numKinds := 0
	for _, ch := range s1 {
		target[ch-'a']++
		if target[ch-'a'] == 1 {
			numKinds++
		}
	}

	count := make([]int, 30)
	total := 0
	for _, ch := range s2[:len(s1)] {
		count[ch-'a']++
		if count[ch-'a'] == target[ch-'a'] {
			total++
		}
	}
	if total == numKinds {
		return true
	}
	for idx, ch := range s2[len(s1):] {
		count[ch-'a']++
		if count[ch-'a'] == target[ch-'a'] {
			total++
		}
		count[s2[idx]-'a']--
		if count[s2[idx]-'a'] == target[s2[idx]-'a']-1 {
			total--
		}
		if total == numKinds {
			return true
		}
	}
	return false
}

func main() {
	println(checkInclusion("ab", "eidbaooo"))
}
