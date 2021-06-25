# [169.多数元素](https://leetcode-cn.com/problems/majority-element)

## 题目描述

给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。


示例 1：

```
输入：[3,2,3]
输出：3
```

示例 2：

```
输入：[2,2,1,1,1,2,2]
输出：2
```

进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。


## 解法

此题有俩种解法，首先已经明确目标数的个数一定是大于n/2的。

1. 排序之后取中间值一定为众数；
2. 使用摩尔投票法，这个经过查阅学习了解到，其实利用的是相互抵消的概念；

### Golang 实现及其测试代码

```go
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
```
