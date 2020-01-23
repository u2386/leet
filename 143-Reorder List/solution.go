package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow = slow.Next
	stack := []*ListNode{}
	for slow != nil {
		stack = append(stack, slow)
		slow = slow.Next
	}

	slow = head
	var prev *ListNode
	for len(stack) > 0 {
		prev, stack = stack[len(stack)-1], stack[:len(stack)-1]
		prev.Next = slow.Next
		slow.Next = prev
		slow = prev.Next
	}
	slow.Next = nil
}

func main() {
	root := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
		},
	}

	reorderList(root)
	for root != nil {
		fmt.Println(root.Val)
		root = root.Next
	}
}
