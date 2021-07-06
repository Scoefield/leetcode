# [206.反转链表](https://leetcode-cn.com/problems/reverse-linked-list/)

## 题目描述

给定一个单链表，返回翻转后的链表。

示例 1：

```
输入：linkList = [3,2,0,-4]
输出：linkList = [-4,0,2,3]
```

## 解题思路

在遍历链表时，将当前节点的 \textit{next}next 指针改为指向前一个节点。由于节点没有引用其前一个节点，因此必须事先存储其前一个节点。在更改引用之前，还需要存储后一个节点。最后返回新的头引用。

## Golang

```go
package main

import (
	. "GitCode/leetcode/solution/link/comm"
)

// 单向链表翻转
func MyReverseList(head *Node) *Node {
	current := head
	var pre *Node
	for current != nil {
		current, current.Next, pre = current.Next, pre, current
	}
	return pre
}

func main() {
	link := CreateLinkList()
	link.Append(1)
	link.Append(2)
	link.Append(3)
	link.Append(4)

	node := MyReverseList(link.Header)
	PrintNodeByHead(node)	
}
```
