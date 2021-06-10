package main

import (
	. "GitCode/leetcode/solution/link/comm"
	"fmt"
)

func reversePrint(head *Node) []int {
	rest := make([]int, 0)
	for head != nil {
		rest = append([]int{head.Data}, rest...)
		head = head.Next
	}
	return rest
}

func main()  {
	linkList := CreateLinkList()
	linkList.Append(11)
	linkList.Append(12)
	linkList.Append(13)

	fmt.Println(reversePrint(linkList.Header))
}
