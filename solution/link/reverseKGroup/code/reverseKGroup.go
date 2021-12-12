package main

import (
	. "GitCode/leetcode/solution/link/comm"
)

func reverseKGroup(head *Node, k int) *Node {
	hair := &Node{Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = reverseNode(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func reverseNode(head, tail *Node) (*Node, *Node) {
	prev := tail.Next
	cur := head
	for prev != tail {
		cur, cur.Next, prev = cur.Next, prev, cur
	}
	return tail, head
}

func main() {
}
