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

func SwapPairs(head *Node) *Node {
	dummy := &Node{
		Data: 0,
		Next: nil,
	}
	pre := dummy
	cur := head
	for cur != nil && cur.Next != nil {
		cur.Next, cur.Next.Next, pre.Next = cur.Next.Next, cur, cur.Next
		pre = cur
		cur = pre.Next
	}
	return dummy.Next
}

func printNodeByHead(head *Node) {
	for head != nil {
		fmt.Println(head.Data)
		head = head.Next
	}
}

func main() {
	// 创建单向链表实例
	link := CreateLinkList()

	// 测试 在头节点添加元素
	link.AddInHead(24)
	link.AddInHead(23)
	link.AddInHead(22)
	link.AddInHead(21)

	// 遍历打印单向链表元素
	link.ScanLinkList()

	data := SwapPairs(link.Header)
	printNodeByHead(data)
}
