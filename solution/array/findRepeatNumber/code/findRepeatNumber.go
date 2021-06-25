func findRepeatNumber(nums []int) int {
    if len(nums) == 0 || nums == nil {
        return -1
    }

    //方法一：先排序在查重
    //时间复杂度：O(nlogn) 用于排序
    //空间复杂度：O(1)
    // sort.Ints(nums) // 实现是快排
	// for i := 0; i < len(nums)-1; i++ {
	// 	if nums[i] == nums[i+1] {
	// 		// 前后元素相等
	// 		return nums[i]
	// 	}
	// }
	// return -1

    //方法二：哈希表
    //时间复杂度：O(n)
    //空间复杂度：O(n)
    // m := make(map[int]int, 0)
	// for i, v := range nums {
	// 	if _, exist := m[v]; exist {
	// 		// 存在该数
	// 		return v
	// 	}
	// 	// 不存在
	// 	m[v] = i
	// }
	// return -1

    //方法三：原地置换
    //时间复杂度：O(n）
    //空间复杂度：O(1)
    // for i := 0; i < len(nums); i++ {
    //     if i != nums[i] {
    //         if nums[i] == nums[nums[i]] {
    //             return nums[i]
    //         }
    //         nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
    //     }
    // }
    // return -1

    //方法四：数组标记
    //时间复杂度：O(n)
    //空间复杂度：O(n)
    var slice [100000]bool
    for i := 0; i < len(nums); i++ {
        if slice[nums[i]] {
            return nums[i]
        }
        slice[nums[i]] = true
    }
    return -1

}