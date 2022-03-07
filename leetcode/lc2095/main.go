package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	first, second := head, dummy
	for first != nil && first.Next != nil {
		first = first.Next.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	head = deleteMiddle(head)
	for ; head != nil; head = head.Next {
		println(head.Val, " ")
	}
}
