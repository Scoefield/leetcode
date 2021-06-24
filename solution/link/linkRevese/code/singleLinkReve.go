package main

import (
	. "GitCode/leetcode/solution/link/comm"
)

// 单向链表翻转
func MyReverseList(head *Node) *Node {
	current := head
	var pre *Node
	for current != nil {
		current, current.Next, pre = current.Next, pre, current
	}
	return pre
}

func main() {
	link := CreateLinkList()
	link.Append(1)
	link.Append(2)
	link.Append(3)
	link.Append(4)

	node := MyReverseList(link.Header)
	PrintNodeByHead(node)	
}
