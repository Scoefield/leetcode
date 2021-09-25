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