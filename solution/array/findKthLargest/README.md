# [153.寻找旋转排序数组中的最小值](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)

## 题目描述

给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1：

```
输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
```

示例 2：

```
输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
```


提示：

- 1 <= k <= nums.length <= 104
- -104 <= nums[i] <= 104

## 解法 - 基于堆排序的选择方法

### 思路解析和算法

我们也可以使用堆排序来解决这个问题——建立一个大根堆，做 k - 1k−1 次删除操作后堆顶元素就是我们要找的答案。在很多语言中，都有优先队列或者堆的的容器可以直接使用，但是在面试中，面试官更倾向于让更面试者自己实现一个堆。所以建议读者掌握这里大根堆的实现方法，在这道题中尤其要搞懂「建堆」、「调整」和「删除」的过程。

友情提醒：「堆排」在很多大公司的面试中都很常见，不了解的同学建议参考《算法导论》或者大家的数据结构教材，一定要学会这个知识点哦！^_^


### Golang 实现及其测试代码

```go
package main

func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	buildHeap(nums, heapSize)
	for i := len(nums); i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		adjustHeap(nums, 0, heapSize)
	}
	return nums[0]
}

func buildHeap(a []int, heapSize int) {
	for i := heapSize/2; i >= 0; i-- {
		adjustHeap(a, i, heapSize)
	}
}

func adjustHeap(a []int, i int, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}

	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		adjustHeap(a, largest, heapSize)
	}
}

func main() {
    nums := []int{3,2,1,5,6,4}
	fmt.Println(findKthLargest(nums, 2))
}
```
