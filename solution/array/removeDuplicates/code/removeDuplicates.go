package main

import "fmt"

func removeDuplicates(nums []int) int {
	cnt, n := 0, len(nums)

	for i := 0; i < n - 1; i++ {
		if nums[i] == nums[i+1] {
			cnt++
		} else {
			nums[i + 1 - cnt] = nums[i+1]
		}
	}
	return n - cnt
}

func main() {
	nums := []int{0, 1, 5, 6, 6, 6}
	fmt.Println(removeDuplicates(nums))
}
