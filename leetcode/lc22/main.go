package main

import "fmt"

func generateParenthesis(n int) []string {
	res := []string{}
	s := []byte{}
	inner(n, 0, 0, &s, &res)
	return res
}

func inner(n int, left int, right int, s *[]byte, res *[]string) {
	if left == n && right == n {
		*res = append(*res, string(*s))
		return
	}

	if left < n {
		*s = append(*s, '(')
		inner(n, left+1, right, s, res)
		*s = (*s)[:len(*s)-1]
	}
	if right < left {
		*s = append(*s, ')')
		inner(n, left, right+1, s, res)
		*s = (*s)[:len(*s)-1]
	}
}

func main() {
	res := generateParenthesis(3)
	for _, v := range res {
		fmt.Println(v)
	}
}
