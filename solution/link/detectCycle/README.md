# [141.环形链表 II](https://leetcode-cn.com/problems/linked-list-cycle-ii/)

## 题目描述

给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。

说明：不允许修改给定的链表。


进阶：

你能用 O(1) 空间解决此问题吗？

示例 1：

![1](./images/circularlinkedlist1.png)

```
输入：head = [3,2,0,-4], pos = 1
输出：返回索引为 1 的链表节点
解释：链表中有一个环，其尾部连接到第二个节点。
```

示例 2：

![2](./images/linkcycle2.png)

```
输入：head = [1,2], pos = 0
输出：返回索引为 0 的链表节点
解释：链表中有一个环，其尾部连接到第一个节点。
```

示例 3：

![3](./images/linkcycle3.png)

```
输入：head = [1], pos = -1
输出：返回 null
解释：链表中没有环。
```

## 解法

方法一：快慢指针法

定义快慢指针 slow、fast，初始指向 head。

快指针每次走两步，慢指针每次走一步，不断循环。当相遇时，说明链表存在环。如果循环结束依然没有相遇，说明链表不存在环。

方法二：哈希表法

一个非常直观的思路是：我们遍历链表中的每个节点，并将它记录下来；一旦遇到了此前遍历过的节点，就可以判定链表中存在环。借助哈希表可以很方便地实现。

## Golang

```go
package main

import (
	. "GitCode/leetcode/solution/link/comm"
)

// 快慢指针实现
func detectCycle(head *Node) *Node {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil	
}

// 遍历链表，用一个 map 存遍历过点节点
func detectCycle2(head *Node) *Node {
	mapNode := make(map[*Node]struct{})
	for head != nil {
		if _, ok := mapNode[head]; ok {
			return head
		}
		mapNode[head] = struct{}{}
		head = head.Next
	}
	return nil
}
```
