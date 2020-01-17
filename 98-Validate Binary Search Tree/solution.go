package main

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

	last := 1 << 31
	var n *TreeNode
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		n = stack[len(stack)-1]
		if n.Left != nil {
			stack = append(stack, n.Left)
			continue
		}

		n, stack = stack[len(stack)-1], stack[:len(stack)-1]

		if n.Val < last {
			return false
		}
		last = n.Val

		if n.Right != nil {
			stack = append(stack, n.Right)
		}
	}
}

func main() {

}
