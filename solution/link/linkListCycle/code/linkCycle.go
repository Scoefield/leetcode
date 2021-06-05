package main

import (
	. "GitCode/leetcode/solution/link/comm"
	"fmt"
)

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
