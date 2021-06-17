package main

import (
	"fmt"
	"sort"
)

// 方法一：排序后取中间值一定为众数
func majorityElement(nums []int) int {
	sort.Ints(nums)
	index := len(nums) / 2
	return nums[index]
}

// 方法二：摩尔投票法，其实是利用相互抵消的概念
func majorityElement2(nums []int) int {
	var cnt, target int

	for _, num := range nums {
		if cnt == 0 {
			target = num
		}

		if target == num {
			cnt++
		} else {
			cnt--
		}
	}

	return target
}

func main() {
	nums := []int{0, 6, 5, 1, 6, 6}
	fmt.Println(majorityElement(nums))
	fmt.Println(majorityElement2((nums)))
}
