# [236.二叉树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree)

## 题目描述

给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

![1](./images/binarytree1.png)

示例 1：

```
输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出：3
解释：节点 5 和节点 1 的最近公共祖先是节点 3 。
```

示例 2：

![2](./images/binarytree2.png)

```
输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出：5
解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。
```

提示：

- 树中节点数目在范围 [2, 105] 内。
- -109 <= Node.val <= 109
- 所有 Node.val 互不相同 。
- p != q
- p 和 q 均存在于给定的二叉树中。

## 解法

### 思路与算法

根据“最近公共祖先”的定义，若 root 是 p, q 的最近公共祖先 ，则只可能为以下情况之一：

如果 p 和 q 分别是 root 的左右节点，那么 root 就是我们要找的最近公共祖先；
如果 p 和 q 都是 root 的左节点，那么返回 lowestCommonAncestor(root.left, p, q)；
如果 p 和 q 都是 root 的右节点，那么返回 lowestCommonAncestor(root.right, p, q)。
边界条件讨论：

如果 root 为 null，则说明我们已经找到最底了，返回 null 表示没找到；
如果 root 与 p 相等或者与 q 相等，则返回 root；
如果左子树没找到，递归函数返回 null，证明 p 和 q 同在 root 的右侧，那么最终的公共祖先就是右子树找到的结点；
如果右子树没找到，递归函数返回 null，证明 p 和 q 同在 root 的左侧，那么最终的公共祖先就是左子树找到的结点。

### 复杂度分析

时间复杂度：O(n)，其中 nn 是给定的二叉搜索树中的节点个数。分析思路与方法一相同。

空间复杂度：O(1)。

### Golang 实现及其测试代码

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
```
