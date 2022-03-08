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

func Infix2Postfix(infix string) []string {
	opStack := Stack{}
	postfix := []string{}
	length := len(infix)

	for i := 0; i < length; i++ {
		ch := string(infix[i])

		switch ch {
		case " ":
			continue
		case "(":
			opStack.Push(ch)
		case ")":
			for !opStack.Empty() {
				op, _ := opStack.Top()
				if op == "(" {
					opStack.Pop()
					break
				}
				postfix = append(postfix, op.(string))
				opStack.Pop()
			}
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			startIdx := i
			digit := ""
			for ; startIdx < length && unicode.IsDigit(rune(infix[startIdx])); startIdx++ {
				digit += string(infix[startIdx])
			}
			// now startIdx exceeds by one
			i = startIdx - 1
			postfix = append(postfix, digit)
		default:
			for !opStack.Empty() {
				op, _ := opStack.Top()
				if opPriorityCmp(op.(string), ch) {
					break
				}
				postfix = append(postfix, op.(string))
				opStack.Pop()
			}
			// handle "(-1)" kind of negtive number
			if ch == "-" {
				idx := i - 1
				for ; idx >= 0; idx-- {
					if infix[idx] != ' ' {
						break
					}
				}
				if idx == -1 || infix[idx] == '(' {
					opStack.Push(NEG)
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
		postfix = append(postfix, op.(string))
	}
	return postfix
}

func opPriorityCmp(op1, op2 string) bool {
	switch op1 {
	case "+", "-":
		if op2 == "*" || op2 == "/" {
			return true
		}
	case "(":
		return true
	}
	// otherwise, op1's priority is less or equal than op2
	return false
}

func calculate(s string) int {
	valStack := Stack{}
	postfix := Infix2Postfix(s)
	idx := 0
	for ; idx < len(postfix); idx++ {
		ch := postfix[idx]
		// if the string is one digit
		if unicode.IsDigit(rune(ch[0])) {
			val, _ := strconv.Atoi(ch)
			valStack.Push(val)
		} else { // the string is an operator
			if ch == NEG {
				val, _ := valStack.Pop()
				valStack.Push(-val.(int))
				continue
			}
			right, _ := valStack.Pop()
			left, _ := valStack.Pop()
			rightVal, _ := right.(int)
			leftVal, _ := left.(int)
			switch string(ch) {
			case "+":
				valStack.Push(leftVal + rightVal)
			case "-":
				valStack.Push(leftVal - rightVal)
			case "*":
				valStack.Push(leftVal * rightVal)
			case "/":
				valStack.Push(leftVal / rightVal)
			}
		}
	}
	result, _ := valStack.Pop()
	return result.(int)
}

func main() {
	s := "1+2*3+(4*5+6)*7"
	fmt.Println(Infix2Postfix(s))
	fmt.Println(calculate(s))
}
