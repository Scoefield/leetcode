package main

import (
	. "GitCode/leetcode/solution/binary/comm/"
)

func postorderTraversal(root *TreeNode) []int {
	// 方法一：
	// 递归方法一
	res:= []int{}
	var dealFunc func(*TreeNode)
	dealFunc = func (node *TreeNode) {
		if node == nil {
			return
		}
		dealFunc(node.Left)
		dealFunc(node.Right)
		res= append(res, node.Val)
	}
	dealFunc(root)
	return res

	/*
	// 递归方法二：
	if root == nil {
		return nil
	}
	res := []int{}
	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
	*/
}

func main() {
	
}
