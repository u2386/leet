package main

import (
	"fmt"
)

// TreeNode is node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func index(arr []int, t int) int {
	for i, v := range arr {
		if v == t {
			return i
		}
	}
	return -1
}

func leftMost(root *TreeNode) interface{} {
	if root.Left != nil {
		if v := leftMost(root.Left); v != nil {
			return v
		}
	}
	if root.Right != nil {
		if v := leftMost(root.Right); v != nil {
			return v
		}
	}
	return root.Val
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	val := postorder[len(postorder)-1]
	inIndex := index(inorder, val)
	if inIndex == -1 {
		return &TreeNode{
			Val: val,
		}
	}
	root := &TreeNode{
		Val: val,
	}
	root.Right = buildTree(inorder[inIndex+1:], postorder[:len(postorder)-1])

	postIndex := len(postorder) - 1
	if root.Right != nil {
		if v := leftMost(root.Right); v != nil {
			postIndex = index(postorder, v.(int))
		}
	}

	root.Left = buildTree(inorder[:inIndex], postorder[:postIndex])
	return root
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := buildTree(inorder, postorder)
	fmt.Println(root)
}
