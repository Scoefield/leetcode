# [70.爬楼梯](https://leetcode-cn.com/problems/climbing-stairs)

## 题目描述

假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

```
输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
```

示例 2：

```
输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶
```

## 解法

想上第 n 级台阶，可从第 n-1 级台阶爬一级上去，也可从第 n-2 级台阶爬两级上去，即：f(n) = f(n-1) + f(n-2)。递推求解即可。
换句话说：可以使用动态规划进行求解，dp[i] = dp[i-1] + dp[i-2]，dp[i]表示n==i时的方法数。举个例子：能够一步走到第三个台阶的只有第二个台阶和第一个台阶，所以走到第三个台阶的方法数等于走到第一个台阶的方法数加上走到第二个台阶的方法数。得出这个规律后这题和斐波那契数列的解法一模一样。

## Golang 实现及其测试代码

```go
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
```
