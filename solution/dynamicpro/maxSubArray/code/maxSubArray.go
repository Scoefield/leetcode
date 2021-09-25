package main

import "fmt"

// 解法一：暴力法（不推荐）
func maxSubArray(nums []int) int {
    result := -100000

    for i := 0; i < len(nums); i++ {
        sum := 0
        for j := i; j < len(nums); j++ {
            sum += nums[j]
            result = maxNum(result, sum)
        }
    }
    return result
}

func maxNum(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// 解法二：动态规划
func maxSubArray2(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    dp[0] = nums[0]
    result := nums[0]
    for i := 1; i < n; i++ {
        dp[i] = maxNum(dp[i-1]+nums[i], nums[i])
        result = maxNum(result, dp[i])
    }
    return result
}

func main() {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
    fmt.Println(maxSubArray(nums))
    fmt.Println(maxSubArray2(nums))
}
