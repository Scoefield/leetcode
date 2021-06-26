package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	output := make([]int, n)
	l, r := 1, 1
	
	for i := 0; i < n; i++ {
		output[i] = l
		l *= nums[i]
	}
	for i := n-1; i >= 0; i-- {
		output[i] *= r
		r *= nums[i]
	}

	return output
}

func main() {
	nums := []int{1,2,3,4}
	fmt.Println(productExceptSelf(nums))
}