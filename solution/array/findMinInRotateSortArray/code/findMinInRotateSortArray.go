package main

import "fmt"

func findMinInRotateSortArray(nums []int) int {
	l, r := 0, len(nums) - 1
	if nums[l] < nums[r] {
		return nums[0]
	}
	for l < r {
		m := (l+r) >> 1  // 取中间, ==> (l+r)/2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}

func main() {
	// [3,4,5,1,2]
	// [4,5,6,7,0,1,2]
	// [11,13,15,17]
	nums := []int{3,4,5,1,2}
	fmt.Println(findMinInRotateSortArray(nums))	// 1
}
