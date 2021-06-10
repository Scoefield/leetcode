package main

import (
	. "GitCode/leetcode/solution/link/comm"
	"fmt"
)

// 方法一：用双指针实现
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

// 方法二：先把链表数据从尾到头保存到slice中，再从头遍历 k 个节点即可
func reverseAndPrintK(head *Node, k int) *Node {
	ret := make([]*Node, 0)
	for head != nil {
		ret = append([]*Node{head}, ret...)
		head = head.Next
	}
	fmt.Println(ret)
	for i, n := range ret {
		if i == k-1 {
			return n
		}
	}
	return nil
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

	// 测试方法一
	data := kthToList(link.Header, 2)
	fmt.Println(data)

	// 测试方法二
	data2 := reverseAndPrintK(link.Header, 2)
	fmt.Println(data2)
}
