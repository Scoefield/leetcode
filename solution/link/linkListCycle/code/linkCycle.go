package main

type Node struct {
	Data interface{}
	Next *Node
}

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
