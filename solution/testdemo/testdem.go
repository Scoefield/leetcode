package main

import (
	. "GitCode/leetcode/solution/link/comm"
)

func demo1(head *Node) *Node {
	dumy := &Node{}
	pre := dumy
	cur := head
	for cur != nil && cur.Next != nil {
		cur.Next, cur.Next.Next, pre.Next = cur.Next.Next, cur, cur.Next
		pre = cur
		cur = pre.Next
	}
	return dumy.Next
}

func main() {
	link := CreateLinkList()
	link.Append(20)
	link.Append(21)
	link.Append(22)
	link.Append(23)
	link.Append(24)

	dumy := demo1(link.Header)
	PrintNodeByHead(dumy)
}