package main

import (
	. "GitCode/leetcode/solution/link/comm"
)

// 快慢指针实现
func detectCycle(head *Node) *Node {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil	
}

// 遍历链表，用一个 map 存遍历过点节点
func detectCycle2(head *Node) *Node {
	mapNode := make(map[*Node]struct{})
	for head != nil {
		if _, ok := mapNode[head]; ok {
			return head
		}
		mapNode[head] = struct{}{}
		head = head.Next
	}
	return nil
}
