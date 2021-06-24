package main

import (
	. "GitCode/leetcode/solution/link/comm"
	"hash"
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

func linkCycle(head *Node) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil{
		if slow == fast {
			return true
		}
		slow, fast = slow.Next, fast.Next.Next
	}
	return false
}

func demo5(headA, headB *Node) *Node {
	if headA == nil || headB == nil {
		return nil
	}

	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}

		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

func kthToLast(head *Node, k int) int {
    pa, pb := head, head
	
	for k > 0 {
		pb = pb.Next
		k--
	}

	for pb != nil {
		pb = pb.Next
		pa = pa.Next
	}
	return pa.Data
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