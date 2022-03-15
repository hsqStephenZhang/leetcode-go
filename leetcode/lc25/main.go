package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := head
	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}
	if length < k {
		return head
	}

	dummy := &ListNode{Next: head}

	prev := dummy
	numGroups := length / k
	for i := 0; i < numGroups; i++ {
		p := prev.Next
		for j := 1; j < k; j++ {
			next := p.Next
			p.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
		}
		prev = p
	}

	return dummy.Next
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	newHead := reverseKGroup(head, 2)
	for newHead != nil {
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}
}
