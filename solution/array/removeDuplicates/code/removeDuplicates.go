package main

import "fmt"

func removeDuplicates(nums []int) int {
	cnt, n := 0, len(nums)
	// 方法一
	// for i := 1; i < n; i++ {
	// 	if nums[i] == nums[i-1] {
	// 		cnt++
	// 	}
	// }
	// 方法二
	for i := 0; i < n - 1; i++ {
		if nums[i] == nums[i+1] {
			cnt++
		}
	}
	return n - cnt
}

func main() {
	nums := []int{0, 1, 5, 6, 6, 6}
	fmt.Println(removeDuplicates(nums))
}
