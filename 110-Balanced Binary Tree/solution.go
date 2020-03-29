package main

// TreeNode is tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func balanceHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := balanceHeight(root.Left)
	right := balanceHeight(root.Right)

	if left == -1 || right == -1 || abs(left - right) > 1 {
		return -1
	}
	return max(left, right) + 1
}

func isBalanced(root *TreeNode) bool {
	return balanceHeight(root) != -1
}

func main() {

}
