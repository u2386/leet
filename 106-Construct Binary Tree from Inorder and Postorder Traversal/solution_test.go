package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func traverse(root *TreeNode) []interface{} {
	if root == nil {
		return []interface{}{root}
	}
	ret := []interface{}{}
	q0, q1 := []*TreeNode{root}, []*TreeNode{}
	for len(q0) > 0 {
		root, q0 = q0[0], q0[1:]
		if root == nil {
			ret = append(ret, nil)
			continue
		}
		ret = append(ret, root.Val)
		q1 = append(q1, root.Left)
		q1 = append(q1, root.Right)

		if len(q0) == 0 {
			for {
				for _, v := range q1 {
					if v != nil {
						goto SWAP
					}
				}
				goto RET
			}
		SWAP:
			q0, q1 = q1, q0
		}
	}
RET:
	return ret
}

func Test01(t *testing.T) {
	inorder := []int{1, 2, 3, 4}
	postorder := []int{1, 4, 3, 2}
	assert.Equal(t, []interface{}{2, 1, 3, nil, nil, nil, 4}, traverse(buildTree(inorder, postorder)))
}

func Test02(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	assert.Equal(t, []interface{}{3, 9, 20, nil, nil, 15, 7}, traverse(buildTree(inorder, postorder)))
}
