package main

import (
	. "GitCode/leetcode/solution/link/comm"
)

func MergeTwoLists(l1 *Node, l2 *Node) *Node {
	head := &Node{}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Data < l2.Data {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else if l2 != nil {
		cur.Next = l2
	}
	return head.Next
}

func main() {
	link := CreateLinkList()
	link.Append(1)
	link.Append(2)
	link.Append(4)

	link2 := CreateLinkList()
	link2.Append(1)
	link2.Append(3)
	link2.Append(4)

	mergeHeadNode := MergeTwoLists(link.Header, link2.Header)
	PrintNodeByHead(mergeHeadNode)
}