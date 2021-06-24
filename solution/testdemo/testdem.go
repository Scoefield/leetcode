package main

import (
	. "GitCode/leetcode/solution/link/comm"
	// "fmt"
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

func demo2(head *Node) []int {
	ret := make([]int, 0)
	for head != nil {
		ret = append([]int{head.Data}, ret...)
		head = head.Next
	}
	return ret
}

func demo3(head *Node) *Node {
	cur := head
	pre := &Node{}
	for cur != nil {
		cur, cur.Next, pre = cur.Next, pre, cur
	}
	return pre
}

func demo4(l1, l2 *Node) *Node {
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
	
	if l1 == nil {
		cur.Next = l2
	} else if l2 == nil {
		cur.Next = l1
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

	mergeHeadNode := demo4(link.Header, link2.Header)
	PrintNodeByHead(mergeHeadNode)
}