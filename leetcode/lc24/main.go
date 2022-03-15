package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}

	cur := dummy
	for {
		first := cur.Next
		if first == nil || first.Next == nil {
			break
		}
		second := first.Next
		secondNext := second.Next

		cur.Next = second
		second.Next = first
		first.Next = secondNext

		cur = first
	}

	return dummy.Next
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}
	head = swapPairs(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
