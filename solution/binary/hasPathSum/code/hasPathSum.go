package main

import (
	. "GitCode/leetcode/solution/binary/comm"
)

func hasPathSum(root *TreeNode, targetSum int) bool {
	 // 递归写法一：
    // var checkPathSum func(*TreeNode, int) bool
    // checkPathSum = func(node *TreeNode, sum int) bool {
    //      if node == nil {
    //         return false
    //     }
    //     if node.Val == sum && node.Left == nil && node.Right == nil {
    //         return true
    //     }
    //     return checkPathSum(node.Left, sum - node.Val) || checkPathSum(node.Right, sum - node.Val)
    // }
    // return checkPathSum(root, targetSum)

    // 递归写法二：
  if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return hasPathSum(root.Left, targetSum - root.Val) || hasPathSum(root.Right, targetSum - root.Val)
}

func main() {
	
}