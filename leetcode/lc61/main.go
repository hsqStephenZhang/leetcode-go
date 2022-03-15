package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	length := 0
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil {
		length++
		cur = cur.Next
	}
	cur.Next = head

	k = k % length

	cur = head
	for k < length-1 {
		cur = cur.Next
		k++
	}
	res := cur.Next
	cur.Next = nil
	return res
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	newHead := rotateRight(head, 2)
	for newHead != nil {
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}
}
