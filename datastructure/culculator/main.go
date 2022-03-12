package main

import (
	"fmt"
	"strconv"
	"unicode"
)

const NEG = ","

type Stack struct {
	data []interface{}
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Push(x interface{}) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() (interface{}, error) {
	if len(s.data) == 0 {
		return nil, fmt.Errorf("stack is empty")
	}
	x := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return x, nil
}

func (s *Stack) Top() (interface{}, error) {
	if len(s.data) == 0 {
		return nil, fmt.Errorf("stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func opPriorityCmp(prev, next string) bool {
	switch prev {
	case "+", "-":
		if next == "*" || next == "/" {
			return true
		}
	case "(":
		return true
	}
	return false
}

// return tokens
func Infix2Prefix(infix string) []string {
	prefix := []string{}

	opStack := Stack{}
	length := len(infix)
	i := 0
	for ; i < length; i++ {
		ch := string(infix[i])
		switch ch {
		case "(":
			opStack.Push(ch)
		case ")":
			for !opStack.Empty() {
				op, _ := opStack.Top()
				if op == "(" {
					opStack.Pop()
					break
				}
				opStack.Pop()
				prefix = append(prefix, op.(string))
			}
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			rightBound := i
			s := ""
			for ; rightBound < length && unicode.IsDigit(rune(infix[rightBound])); rightBound++ {
				s += string(infix[rightBound])
			}
			prefix = append(prefix, s)
			i = rightBound - 1
		default:
			for !opStack.Empty() {
				op, _ := opStack.Top()
				if opPriorityCmp(op.(string), ch) {
					break
				}
				opStack.Pop()
				prefix = append(prefix, op.(string))
			}
			if ch == "-" {
				idx := i - 1
				for ; idx >= 0; idx-- {
					if infix[idx] != ' ' {
						break
					}
				}
				if idx == -1 || infix[idx] == '(' {
					opStack.Push(",")
				} else {
					opStack.Push(ch)
				}
			} else {
				opStack.Push(ch)
			}
		}
	}
	for !opStack.Empty() {
		op, _ := opStack.Pop()
		prefix = append(prefix, op.(string))
	}

	return prefix
}

func calculate(input string) int {
	valStack := Stack{}
	prefix := Infix2Prefix(input)
	for _, v := range prefix {
		if unicode.IsDigit(rune(v[0])) {
			val, _ := strconv.Atoi(v)
			valStack.Push(val)
		} else {
			if v == NEG {
				val, _ := valStack.Pop()
				val = -val.(int)
				valStack.Push(val)
				continue
			}

			rightOperand, _ := valStack.Pop()
			leftOperand, _ := valStack.Pop()
			switch v {
			case "+":
				valStack.Push(leftOperand.(int) + rightOperand.(int))
			case "-":
				valStack.Push(leftOperand.(int) - rightOperand.(int))
			case "*":
				valStack.Push(leftOperand.(int) * rightOperand.(int))
			case "/":
				valStack.Push(leftOperand.(int) / rightOperand.(int))
			}
		}
	}
	ans, _ := valStack.Top()
	return ans.(int)
}

func main() {
	infix := "(-1)+1"
	fmt.Println(Infix2Prefix(infix))
	fmt.Println(calculate(infix))
}
