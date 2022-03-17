package main

import "fmt"

func removeInvalidParentheses(s string) []string {
	ans := []string{}

	leftCnt := 0
	rightCnt := 0
	targetLength := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			leftCnt++
			targetLength++
		} else if s[i] == ')' && leftCnt > rightCnt {
			rightCnt++
			targetLength++
		} else if s[i] != ')' && s[i] != '(' {
			targetLength++
		}
	}
	if leftCnt != rightCnt {
		targetLength -= leftCnt - rightCnt
	}

	path := []byte{}
	used := make(map[string]bool)
	dfs(s, 0, 0, 0, rightCnt, &path, &ans, used, targetLength)

	return ans
}

func dfs(s string, index int, leftCnt int, rightCnt int, targetCnt int, path *[]byte, ans *[]string, used map[string]bool, targetLength int) {
	if leftCnt == targetCnt && rightCnt == targetCnt && len(*path) == targetLength {
		if !used[string(*path)] {
			*ans = append(*ans, string(*path))
			used[string(*path)] = true
		}
		return
	}
	if index == len(s) {
		return
	}
	for j := index; j < len(s); j++ {
		ch := s[j]
		if ch == '(' && leftCnt < targetCnt {
			*path = append(*path, ch)
			dfs(s, j+1, leftCnt+1, rightCnt, targetCnt, path, ans, used, targetLength)
			*path = (*path)[:len(*path)-1]
		}
		if ch == ')' && rightCnt < targetCnt && rightCnt < leftCnt {
			*path = append(*path, ch)
			dfs(s, j+1, leftCnt, rightCnt+1, targetCnt, path, ans, used, targetLength)
			*path = (*path)[:len(*path)-1]
		}
		if ch != '(' && ch != ')' {
			*path = append(*path, ch)
			dfs(s, j+1, leftCnt, rightCnt, targetCnt, path, ans, used, targetLength)
			*path = (*path)[:len(*path)-1]
		}
	}
}

func main() {
	s := "n"
	fmt.Println(removeInvalidParentheses(s))
}
