# [189.旋转数组](https://leetcode-cn.com/problems/rotate-array/)

## 题目描述

给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

进阶：

- 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
- 使用空间复杂度为 O(1) 的 原地算法解决这个问题就更好了。

示例 1：

```
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
```

示例 2：

```
输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释: 
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]
```

## 解法

### 方法一：使用额外的数组

我们可以使用额外的数组来将每个元素放至正确的位置。用 n 表示数组的长度，我们遍历原数组，将原数组下标为 i 的元素放至新数组下标为 (i+k) mod n 的位置，最后将新数组拷贝至原数组即可。

复杂度分析：

- 时间复杂度： O(n)，其中 n 为数组的长度。
- 空间复杂度： O(n)。

### 方法二：数组翻转

该方法基于如下思路：将数组的元素向右移动 k 次后，尾部 k mod n 个元素会移动至数组头部，其余元素向后移动 k mod n 个位置。

该方法为数组的翻转：可以先将所有元素翻转，这样尾部的 k mod n 个元素就被移至数组头部，然后再翻转 [0, k mod n-1] 区间的元素和 [k mod n, n-1] 区间的元素即能得到最后的答案。

以 n=7，k=3 为例进行如下展示：

| 操作 | 结果 |
| --- | --- |
| 原始数组 || 1,2,3,4,5,6,7 |
| 翻转所有元素 | 7,6,5,4,3,2,1 |
| 翻转[0,k mod n-1]区间的元素 | 5,6,7,4,3,2,1 |
| 翻转[k mode n,n-1]区间的元素 | 5,6,7,1,2,3,4 |


### Golang 实现及其测试代码

```go
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
	// 时间复杂度： O(n)，其中 n 为数组的长度。空间复杂度： O(1)。
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
```