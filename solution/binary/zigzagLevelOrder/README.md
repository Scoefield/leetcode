# [103.二叉树的锯齿形层序遍历](https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/)

## 题目描述

给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树[3,9,20,null,null,15,7],

```
    3
   / \
  9  20
    /  \
   15   7
```

返回锯齿形层序遍历如下：
```
[
  [3],
  [20,9],
  [15,7]
]
```

## 解法

### Golang 实现及其测试代码

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    res := [][]int{}
    if root == nil {
        return res
    }

    leftToRight := true
    var q []*TreeNode
    q = append(q, root)

    for len(q) > 0 {
        length := len(q)
        level := []int{}

        for i := 0; i < length; i++ {
            node := q[0]
            q = q[1:]

            level = append(level, node.Val)
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
        var levelFinal []int
        if leftToRight {
            levelFinal = level
        } else {
            for i := len(level) - 1; i >= 0; i-- {
                levelFinal = append(levelFinal, level[i])
            }
        }
        if len(levelFinal) > 0 {
            res = append(res, levelFinal)
        }
        leftToRight = !leftToRight
    }

    return res
}
```
