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

type Node struct {
	Data int
	Next *Node
}

func CreateNode(data int) *Node {
	return &Node{
		Data: data,
		Next: nil,
	}
}

// 定义单向链表结构体
type LinkList struct {
	Header *Node	// 头节点，指向第一个节点
	Length int		// 链表长度
}

// 创建链表方法
func CreateLinkList() *LinkList {
	return &LinkList{
		Header: CreateNode(0),
		Length: 0,
	}
}

// 在头节点插入数据的方法
func (l *LinkList) AddInHead(data int) {
	newNode := CreateNode(data)
	defer func() {
		l.Length++
	}()

	if l.Length == 0 {
		l.Header = newNode
	}else {
		newNode.Next = l.Header
		l.Header = newNode		// 头指针指向新加的
	}
}

// 遍历单向链表并打印出来
func (l *LinkList) ScanLinkList() {
	current := l.Header
	i := 1
	for current.Next != nil {
		fmt.Printf("第 %d 个节点的值为：%d\n", i, current.Data)
		current = current.Next
		i++
	}
	fmt.Printf("第 %d 个节点的值为：%d\n", i, current.Data)
}

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
