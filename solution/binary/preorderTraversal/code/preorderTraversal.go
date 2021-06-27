package main

import (
	. "GitCode/leetcode/solution/binary/comm/"
)

func preorderTraversal(root *TreeNode) []int {
	/*
	// 方法一：
	// 递归方法一
	res:= []int{}
	var dealFunc func(*TreeNode)
	dealFunc = func (node *TreeNode) {
		if node == nil {
			return
		}
		res= append(ret, node.Val)
		dealFunc(node.Left)
		dealFunc(node.Right)
	}
	dealFunc(root)
	return res
	*/

	/*
	// 递归方法二：
	if root == nil {
		return nil
	}
	res := []int{}
	res = append(res, root.Val)
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
	*/

	// 方法二：迭代的方法（栈来实现）
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			res = append(res, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}

func main() {
	
}
