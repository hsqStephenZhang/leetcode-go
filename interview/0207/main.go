package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB
	len1, len2 := 0, 0
	for p1 != nil {
		len1++
		p1 = p1.Next
	}
	for p2 != nil {
		len2++
		p2 = p2.Next
	}
	if len1 > len2 {
		headA, headB = headB, headA
		len1, len2 = len2, len1
	}
	delta := len2 - len1
	for delta > 0 {
		headB = headB.Next
		delta--
	}

	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

func main() {
	commonList := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}
	headA := &ListNode{Val: 3, Next: commonList}
	headB := &ListNode{Val: 0, Next: &ListNode{Val: 9, Next: &ListNode{Val: 1, Next: commonList}}}
	intersect := getIntersectionNode(headB, headA)
	if intersect != nil {
		fmt.Println(intersect.Val)
	}
}
