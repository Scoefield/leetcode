package main

import (
	. "GitCode/leetcode/solution/binary/comm/"
)

func inorderTraversal(root *TreeNode) []int {
	/*
	// 方法一：
	// 递归方法一
	res:= []int{}
	var dealFunc func(*TreeNode)
	dealFunc = func (node *TreeNode) {
		if node == nil {
			return
		}
		dealFunc(node.Left)
		res= append(res, node.Val)
		dealFunc(node.Right)
	}
	dealFunc(root)
	return res
	*/

	// 递归方法二：
	if root == nil {
		return nil
	}
	res := []int{}
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res

}

func main() {
	
}
