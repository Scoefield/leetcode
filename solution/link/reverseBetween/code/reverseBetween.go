package main

import (
	. "GitCode/leetcode/solution/link/comm"
)

func reverseLink(head *Node) {
	curr := head
	var pre *Node
	for curr != nil {
		curr, curr.Next, pre = curr.Next, pre, curr
	}
}

// 反转链表II
func reverseBetween(head *Node, left int, right int) *Node {
	dyNode := &Node{Data: -1}
	dyNode.Next = head

	pre := dyNode
	for i := 0; i < left - 1; i++ {
		pre = pre.Next
	}

	rNode := pre
	for i := 0; i < right - left + 1; i++ {
		rNode = rNode.Next
	}

	lNode := pre.Next
	curr := rNode.Next

	pre.Next = nil
	rNode.Next = nil

	reverseLink(lNode)

	pre.Next = rNode
	lNode.Next = curr

	return dyNode.Next
}

func main() {
	link := CreateLinkList()
	link.Append(1)
	link.Append(2)
	link.Append(3)
	link.Append(4)
	link.Append(5)

	node := reverseBetween(link.Header, 2, 4)
	PrintNodeByHead(node)	
}
