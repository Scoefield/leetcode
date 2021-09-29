package main

import . "GitCode/leetcode/solution/binary/comm"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	leftToRight := true
	var q []*TreeNode
	q = append(q, root)

	for len(q) > 0 {
		length := len(q)
		level := []int{}

		for i := 0; i < length; i++ {
			node := q[0]
			q = q[1:]

			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		var levelFinal []int
		if leftToRight {
			levelFinal = level
		} else {
			for i := len(level) - 1; i >= 0; i-- {
				levelFinal = append(levelFinal, level[i])
			}
		}
		if len(levelFinal) > 0 {
			res = append(res, levelFinal)
		}
		leftToRight = !leftToRight
	}

	return res
}

func main() {
	
}
