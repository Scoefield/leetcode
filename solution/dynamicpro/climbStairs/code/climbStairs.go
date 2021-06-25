package main

import "fmt"

/*
	实现方法跟斐波纳契数大同小异，不同点已在下面注释说明
*/

// 常规方法实现（动态规划），时间复杂度：O(n)，空间复杂度：O(1)
func climbStairs(n int) int {
	a, b := 0, 1
	for n > 0 {
		a, b = b, a+b
		n--
	}
	return b
}

// 递归方法实现
func climbStairs2(n int) int {
	switch {
	case n <= 2: // 斐波纳契数 这里应该为 n < 2
		return n
	default:
		return climbStairs2(n-1) + climbStairs2(n-2)
	}
}

func main() {
	n := 5
	fmt.Println(climbStairs(n))  // 8
	fmt.Println(climbStairs2(n)) // 8
}
