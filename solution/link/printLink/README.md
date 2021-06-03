# [从尾到头打印链表](https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/)

## 题目描述

输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

示例 1：

```
输入：head = [1,3,2]
输出：[2,3,1]
```

限制：

- 0 <= 链表长度 <= 10000

## 解法

栈实现。或者其它方式，见题解。

## Golang

```go
type LinkNode struct {
	Val int
	Next *LinkNode
}

func reversePrint(head *LinkNode) []int {
	rest := []int{}
	for head != nil {
		rest = append([]int{head.Val}, rest...)
		head = head.Next
	}
	return rest
}
```
