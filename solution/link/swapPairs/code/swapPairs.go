package main

import (
	. "GitCode/leetcode/solution/link/comm"
)

func SwapPairs(head *Node) *Node {
	dummy := &Node{}
	pre := dummy
	cur := head
	for cur != nil && cur.Next != nil {
		cur.Next, cur.Next.Next, pre.Next = cur.Next.Next, cur, cur.Next
		pre = cur
		cur = pre.Next
	}
	return dummy.Next
}

func main() {
	// 创建单向链表实例
	link := CreateLinkList()

	// 测试 在头节点添加元素
	link.AddInHead(24)
	link.AddInHead(23)
	link.AddInHead(22)
	link.AddInHead(21)
	link.AddInHead(20)

	// 遍历打印单向链表元素
	// [21, 22, 23, 24]
	link.ScanLinkList()

	data := SwapPairs(link.Header)
	PrintNodeByHead(data)
}
