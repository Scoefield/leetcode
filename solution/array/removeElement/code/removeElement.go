package main

import "fmt"

func removeElement(nums []int, val int) int {
	cnt, n := 0, len(nums)

	for i, num := range nums {
		if num == val {
			cnt++
		} else {
			nums[i - cnt] = nums[i]
		}
	}
	
	return n - cnt
}

func main() {
	nums := []int{0,1,2,2,3,0,4,2}
	fmt.Println(removeElement(nums, 2))
	fmt.Println(nums)	// nums 为引用，[0 1 3 0 4 0 4 2]
}
