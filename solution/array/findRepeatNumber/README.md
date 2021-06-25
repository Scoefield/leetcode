# [数组中重复的数字](https://leetcode-cn.com/problems/majority-element)

## 题目描述

找出数组中重复的数字。

在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

示例 1：

```
输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3 
```

## 解法

此题有多种解法：

1. 先排序再查重，时间复杂度：O(nlogn)，用的快排，空间复杂度：O(1)；
2. 使用哈希表存放，时间复杂度：O(n)，空间复杂度：O(n)；
3. 原地置换，时间复杂度：O(n)，空间复杂度：O(1)；
4. 数组标记，时间复杂度：O(n)，空间复杂度：O(n)；

### Golang 实现及其测试代码

```go
func findRepeatNumber(nums []int) int {
    if len(nums) == 0 || nums == nil {
        return -1
    }

    //方法一：先排序再查重
    //时间复杂度：O(nlogn)
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
```
