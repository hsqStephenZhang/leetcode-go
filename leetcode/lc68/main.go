package main

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) (ans []string) {
	i, n := 0, len(words)
	for {
		start := i   // 这一行文本在 words 的起始位置
		wordLen := 0 // 统计这一行单词字符的个数
		for ; i < n && wordLen+len(words[i])+i-start <= maxWidth; i++ {
			wordLen += len(words[i])
		}
		// 循环结束后，i 为这一行文本在 words 的结束位置 +1
		if i == n { // 最后一行
			s := strings.Join(words[start:], " ")
			ans = append(ans, s+strings.Repeat(" ", maxWidth-len(s)))
			return
		}
		space := maxWidth - wordLen // 需要插入至单词间的空格个数
		if i-start == 1 {           // 只有一个单词的情况，在末尾加上空格即可
			ans = append(ans, words[start]+strings.Repeat(" ", space))
		} else { // 计算两单词间至少要加多少个空格，以及有多少个间隔要额外加一个空格
			avg, extra := space/(i-start-1), space%(i-start-1)
			avgSpace := strings.Repeat(" ", avg)
			s1 := strings.Join(words[start:start+extra+1], avgSpace+" ")
			s2 := strings.Join(words[start+extra+1:i], avgSpace)
			ans = append(ans, s1+avgSpace+s2)
		}
	}
}

func main() {
	var ans []string
	// ans = fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16)
	// for _, v := range ans {
	// 	fmt.Println(v)
	// }
	// fmt.Println()

	// ans = fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16)
	// for _, v := range ans {
	// 	fmt.Println(v)
	// }

	// fmt.Println()

	// ans = fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain",
	// 	"to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20)
	// for _, v := range ans {
	// 	fmt.Println(v)
	// }
	ans = []string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country"}
	ans = fullJustify(ans, 20)
	for _, v := range ans {
		fmt.Println(v)
	}
}
