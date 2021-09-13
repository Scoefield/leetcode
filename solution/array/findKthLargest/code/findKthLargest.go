package main

import "fmt"

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