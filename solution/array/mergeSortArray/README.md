# [88.合并两个有序数组](https://leetcode-cn.com/problems/merge-sorted-array/)

## 题目描述

给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。

初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 的空间大小等于 m + n，这样它就有足够的空间保存来自 nums2 的元素。

示例 1：

```
输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
```

示例 2：

```
输入：nums1 = [1], m = 1, nums2 = [], n = 0
输出：[1]
```

提示：

- nums1.length == m + n
- nums2.length == n
- 0 <= m, n <= 200
- 1 <= m + n <= 200
- -109 <= nums1[i], nums2[i] <= 109

## 解法

### 思路解析

j >= 0 的条件下，比较 nums[i] 和 nums[j] 的大小，比较步骤如下：

![](./images/mergeSortarray.png)


### Golang 实现及其测试代码

```go
package main

import "fmt"

/*				i	  k
	nums1: [1,2,3,0,0,0]
					  j
	nums2: 		 [2,5,6]
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for j >= 0 {
		if i >=0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

func main() {
	nums1 := []int{1,2,3,0,0,0}
	nums2 := []int{2,5,6}
	
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}
```
