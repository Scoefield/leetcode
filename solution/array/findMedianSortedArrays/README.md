# [4.寻找两个正序数组的中位数](https://leetcode-cn.com/problems/median-of-two-sorted-arrays/)

## 题目描述

给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

示例 1：

```
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
```

示例 2：

```
输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
```

示例 3：

```
输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000
```



提示：

- nums1.length == m
- nums2.length == n
- 0 <= m <= 1000
- 0 <= n <= 1000
- 1 <= m + n <= 2000
- -106 <= nums1[i], nums2[i] <= 106

进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？

## 解法

### 思路解析和算法

首先合并两个有序的数组（参考88题），然后再求中位数，注意元素个数为奇数还是偶数。


### Golang 实现及其测试代码

```go
package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    // 先合并两个有序数组，参考88题
    k := len(nums1)+len(nums2)-1
    num3 := make([]int, k+1)
    copy(num3, nums1)
    i, j := len(nums1)-1, len(nums2)-1
    
    for j >= 0 {
        if i >= 0 && num3[i] > nums2[j] {
            num3[k] = num3[i]
            i--
        } else {
            num3[k] = nums2[j]
            j--
        }
        k--
    }

    // 再求中位数
    n := len(num3)
    if n % 2 != 0 {
        return float64(num3[(n-1) / 2.0])
    }
    return float64((num3[n/2] + num3[n/2 - 1])) / 2.0
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
```
