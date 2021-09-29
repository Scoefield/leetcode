package main

import (
	. "GitCode/leetcode/solution/link/comm"
)

func isPalindrome(head *Node) bool {
	valSlice := []int{}
	for head != nil {
		valSlice = append(valSlice, head.Data)
		head = head.Next
	}
	n := len(valSlice)
	for i, val := range valSlice[:n/2] {
		if val != valSlice[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
}
