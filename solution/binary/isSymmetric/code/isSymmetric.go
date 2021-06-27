package main

import (
	. "GitCode/leetcode/solution/binary/comm"
)

func isSymmetric(root *TreeNode) bool {
    var checkFunc func(*TreeNode, *TreeNode) bool
	checkFunc = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		return p.Val == q.Val && checkFunc(p.Left, q.Right) && checkFunc(p.Right, q.Left)
	}
	return checkFunc(root, root)
}

func main() {
	
}