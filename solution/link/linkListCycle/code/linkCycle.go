package main

import (
	. "GitCode/leetcode/solution/link/comm"
	"fmt"
)

func LinkHasCycle(head *Node) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	// 创建单向链表实例
	link := CreateLinkList()

	// 测试 在头节点添加元素
	link.Append(23)
	link.Append(22)
	link.Append(21)

	hasCycle := LinkHasCycle(link.Header)
	fmt.Println(hasCycle)	// false

}
