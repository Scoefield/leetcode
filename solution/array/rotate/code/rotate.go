package main

import "fmt"

// 方法二用到的翻转函数
func reverse(tmp []int) {
	n := len(tmp)
	for i := 0; i < n/2; i++ {
		tmp[i], tmp[n-1-i] = tmp[n-1-i], tmp[i]
	}
}

func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 || nums == nil {
		return
	}

	// 方法一：使用额外的数组存，最后再copy 回去
	// 时间复杂度： O(n)，其中 n 为数组的长度。空间复杂度： O(n)。
	// tmp := make([]int, n)
	// for i, val := range nums {
	// 	tmp[(i+k)%n] = val
	// }
	// copy(nums, tmp)
	// fmt.Println(nums)

	// 方法二：数组翻转
	k %= n
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
	fmt.Println(nums)
}

func main() {
	nums := []int{1,2,3,4,5,6,7}
	rotate(nums, 3)
}
