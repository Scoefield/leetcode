# [112.路径总和](https://leetcode-cn.com/problems/path-sum/)

## 题目描述

给你二叉树的根节点 root 和一个表示目标和的整数 targetSum ，判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。

叶子节点 是指没有子节点的节点。

示例1：
![1](./images/pathsum1.jpeg)
```
输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
输出：true
```

示例2：
![1](./images/pathsum2.jpeg)
```
输入：root = [1,2,3], targetSum = 5
输出：false
```

示例3：
```
输入：root = [1,2], targetSum = 0
输出：false
```

## 解法

用递归方法求解比较简单。

### Golang 实现及其测试代码

```go
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
```
