package main

import (
	. "GitCode/leetcode/solution/link/comm"
)

func GetIntersectionNode(headA *Node, headB *Node) *Node {
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

func main() {

}