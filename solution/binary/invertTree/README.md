# [226.翻转二叉树](https://leetcode-cn.com/problems/invert-binary-tree/)

## 题目描述

翻转一棵二叉树。

示例：

输入：
```
     4
   /   \
  2     7
 / \   / \
1   3 6   9
```

输出：
```
     4
   /   \
  7     2
 / \   / \
9   6 3   1
```

## 解法

用递归方法求解比较简单。

### Golang 实现及其测试代码

```go
import (
	. "GitCode/leetcode/solution/binary/comm"
)

func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
```
