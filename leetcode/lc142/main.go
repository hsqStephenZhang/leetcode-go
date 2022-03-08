package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	length := 0
	for {
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		length++

		if slow == fast {
			newSlow := head
			for newSlow != slow {
				newSlow = newSlow.Next
				slow = slow.Next
			}
			return slow
		}
	}
}

func main() {
	head := &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: -4, Next: nil}}}}
	head.Next.Next.Next.Next = head.Next
	fmt.Println(detectCycle(head).Val)
}
