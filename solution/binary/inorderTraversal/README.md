# [94.二叉树中序遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal)

## 题目描述

给你二叉树的根节点 root ，返回它节点值的 中序 遍历（左根右）。

示例 1：
![1](./images/inorder_1.jpeg)

```
输入：root = [1,null,2,3]
输出：[1,3,2]
```

提示：

- 树中节点数目在范围 [0, 100] 内
- -100 <= Node.val <= 100

进阶：递归算法很简单，试着通过迭代算法完成。

### Golang 实现及其测试代码

```go
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
		res= append(res, node.Val)
		dealFunc(node.Right)
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
	res = append(res, root.Val)
	res = append(res, postorderTraversal(root.Right)...)
	return res
	*/
}
```
