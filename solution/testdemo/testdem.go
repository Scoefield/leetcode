package main

import (
	. "GitCode/leetcode/solution/link/comm"
	"fmt"
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

func swapNode(head *Node) *Node {
	dumy := &Node{}
	pre := dumy
	cur := head
	for cur != nil && cur.Next != nil {
		cur.Next, cur.Next.Next, pre.Next = cur.Next.Next, cur, cur.Next
		cur = pre.Next
		pre = cur.Next
	}
	return dumy.Next
}

func demo6(nums []int, k int) {
	if nums == nil || len(nums) == 0 || k > len(nums) {
		return
	}

	var ret []int
	ret = append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	copy(nums, ret)
	fmt.Println(nums)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	k := m + n - 1
	i := m - 1
	j := n - 1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	output := make([]int, n)
	l, r := 1, 1
	for i := 0; i < n; i++ {
		output[i] = l
		l *= nums[i]
	}
	for i := n - 1; i >= 0; i-- {
		output[i] *= r
		r *= nums[i]
	}
	return output
}

func removeDuplicates(nums []int) int {
	cnt, n := 0, len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == nums[i+1] {
			cnt++
		}
	}
	return n - cnt
}

func removeElement(nums []int, val int) int {
	cnt, n := 0, len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == val {
			cnt++
		} else {
			nums[i - cnt] = nums[i]
		}
	}
	return n - cnt
}

func main() {
	link := CreateLinkList()
	link.Append(1)
	// link.Append(2)
	// link.Append(4)

	// link2 := CreateLinkList()
	// link2.Append(1)
	// link2.Append(3)
	// link2.Append(4)

	// mergeHeadNode := demo4(link.Header, link2.Header)
	// PrintNodeByHead(mergeHeadNode)

	nums := []int{1,2,3,4,5,6,7}
	demo6(nums, 3)
}