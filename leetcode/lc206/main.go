package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Val: 0, Next: head}
	for head.Next != nil {
		cur := head.Next
		head.Next = cur.Next
		cur.Next = dummy.Next
		dummy.Next = cur
	}
	return dummy.Next
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	newHead := reverseList(head)
	for newHead != nil {
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}
}
