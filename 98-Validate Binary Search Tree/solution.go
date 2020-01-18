package main

import (
	"fmt"
)

// TreeNode definite a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	last := -1 << 63
	n := root
	stack := []*TreeNode{}
	for n != nil || len(stack) > 0 {
		for n != nil {
			stack = append(stack, n)
			n = n.Left
		}

		n, stack = stack[len(stack)-1], stack[:len(stack)-1]

		if n.Val <= last {
			return false
		}
		last = n.Val
		n = n.Right
	}
	return true
}

func main() {
	root := TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(isValidBST(&root))
}
