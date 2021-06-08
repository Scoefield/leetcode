# [爬楼梯](https://leetcode-cn.com/problems/climbing-stairs)

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

## Golang 实现及其测试代码

```go
package main

import "fmt"

func climbStairs(n int) int {
	a, b := 0, 1
	for n > 0 {
		a, b = b, a+b
		n--
	}
	return b
}

func main() {
	fmt.Println(climbStairs(5))
}
```
