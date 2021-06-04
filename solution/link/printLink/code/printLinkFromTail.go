package main

type LinkNode struct {
	Val int
	Next *LinkNode
}

func reversePrint(head *LinkNode) []int {
	rest := make([]int, 0)
	for head != nil {
		rest = append([]int{head.Val}, rest...)
		head = head.Next
	}
	return rest
}