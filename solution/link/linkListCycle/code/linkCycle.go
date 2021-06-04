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

func LinkHasCycle(head *Node) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for {
		if fast == nil || fast.Next == nil {
			return  false
		}
		if slow == fast {
			return true
		}
		slow, fast = slow.Next, fast.Next.Next
	}
}

func main() {
	// 创建单向链表实例
	link := CreateLinkList()

	// 测试 在头节点添加元素
	link.AddInHead(23)
	link.AddInHead(22)
	link.AddInHead(21)

	hasCycle := LinkHasCycle(link.Header)
	fmt.Println(hasCycle)	// false

}
