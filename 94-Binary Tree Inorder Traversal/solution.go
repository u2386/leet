package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	ret := []int{}

	stack := []*TreeNode{}
	cur := root
	for {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
			continue
		}

		if len(stack) > 0 {
			cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
			ret = append(ret, cur.Val)
			cur = cur.Right
			continue
		}

		break
	}
	return ret
}

func main() {

}
