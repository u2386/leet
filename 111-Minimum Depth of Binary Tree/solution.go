package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	depth := 0

	if root == nil {
		return depth
	}

	q := [][]*TreeNode{}
	q = append(q, []*TreeNode{root})
	for {
		tair := q[0]
		q = append(q[1:], []*TreeNode{})
		if len(tair) == 0 {
			return depth
		}
		depth++

		for _, n := range tair {
			if n == nil || (n.Left == nil && n.Right == nil) {
				return depth
			}
			if n.Left != nil {
				q[len(q)-1] = append(q[len(q)-1], n.Left)
			}
			if n.Right != nil {
				q[len(q)-1] = append(q[len(q)-1], n.Right)
			}
		}
	}

}

func main() {
}
