package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := []int{}
	stack := []*TreeNode{root}
	var r *TreeNode
	for len(stack) > 0 {
		r, stack = stack[len(stack)-1], stack[:len(stack)-1]
		ret = append(ret, r.Val)

		if r.Left != nil {
			stack = append(stack, r.Left)
		}
		if r.Right != nil {
			stack = append(stack, r.Right)
		}
	}

	for i := 0; i < len(ret)>>1; i++ {
		ret[i], ret[len(ret)-i-1] = ret[len(ret)-i-1], ret[i]
	}
	return ret
}

func main() {
	root := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	fmt.Println(postorderTraversal(root))
}
