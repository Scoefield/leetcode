# [144.二叉树前序遍历](https://leetcode-cn.com/problems/binary-tree-preorder-traversal)

## 题目描述

给你二叉树的根节点 root ，返回它节点值的 前序 遍历。

示例 1：
![1](./images/inorder_1.jpeg)

```
输入：root = [1,null,2,3]
输出：[1,2,3]
```

示例 2：

```
输入：root = []
输出：[]
```

示例 3：

```
输入：root = [1]
输出：[1]
```

示例 4：
![4](./images/inorder_4.jpeg)

```
输入：root = [1,2]
输出：[1,2]
```


提示：

- 树中节点数目在范围 [0, 100] 内
- 100 <= Node.val <= 100

进阶：递归算法很简单，试着通过迭代算法完成。

## 解法

### 迭代思路解析

利用栈实现非递归遍历。非递归的思路如下：

1. 定义一个栈，先将根节点压入栈
2. 若栈不为空，每次从栈中弹出一个节点
3. 处理该节点
4. 先把节点右孩子压入栈，接着把节点左孩子压入栈（如果有孩子节点）
5. 重复 2-4
6. 返回结果

### Golang 实现及其测试代码

```go

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

```
