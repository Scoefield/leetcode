# 链表翻转

## 题目描述

给定一个单链表，返回翻转后的链表。

示例 1：

```
输入：linkList = [3,2,0,-4]
输出：linkList = [-4,0,2,3]
```

## Golang

```go
// 定义节点结构体
type Node struct {
	Data interface{}	// 接口类型的数据data
	Next *Node			// 下一个节点
}

// 定义单向链表结构体
type LinkList struct {
	Header *Node	// 头节点，指向第一个节点
	Length int		// 链表长度
}

type Node struct {
	Data interface{}
	Next *Node
}

// 单向链表翻转
func (l *LinkList) ReverseList() {
	current := l.Header
	var pre *Node
	for current != nil {
		current, current.Next, pre = current.Next, pre, current
	}
	l.Header = pre
}
```
