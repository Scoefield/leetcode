# [145.二叉树后序遍历](https://leetcode-cn.com/problems/binary-tree-postorder-traversal)

## 题目描述

给你二叉树的根节点 root ，返回它节点值的 后序 遍历（左右根）。

示例 1：
![1](./images/inorder_1.jpeg)

```
输入: [1,null,2,3]  
   1
    \
     2
    /
   3 

输出: [3,2,1]
```

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
```
