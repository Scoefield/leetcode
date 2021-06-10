# [返回倒数第k个节点](https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci)

## 题目描述

实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。

注意：本题相对原题稍作改动

示例：

```
输入： 1->2->3->4->5 和 k = 2
输出： 4
```

说明：

给定的 k 保证是有效的。

## 解法

定义 slow、fast 指针指向 head。

fast 先向前走 k 步，接着 fast、slow 同时向前走，当 fast 指向 null 时，slow 指向的节点即为链表的倒数第 k 个节点。

## Golang 实现及其测试代码

```go
package main

import "fmt"

func kthToList(head *Node, k int) int {
	slow, fast := head, head
	for k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow.Data
}

func main() {
	// 创建单向链表实例
	link := CreateLinkList()

	// 测试 在头节点添加元素
	link.AddInHead(23)
	link.AddInHead(22)
	link.AddInHead(21)

	// 遍历打印单向链表元素
	link.ScanLinkList()

	data := kthToList(link.Header, 2)
	fmt.Println(data)
}
```
