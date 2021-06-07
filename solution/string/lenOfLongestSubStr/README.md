# [无重复字符的最长子串](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters)

## 题目描述

给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1：

```
输入: s = "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
```

示例 2：

```
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
```

示例 3：

```
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
```


示例 4：

```
输入: s = ""
输出: 0
```

提示：

- 0 <= s.length <= 5 * 104
- s 由英文字母、数字、符号和空格组成

## 解法

- 定义一个哈希表存放字符及其出现的位置；
- 定义 i, j 分别表示不重复子串的开始位置和结束位置；
- j 向后遍历，若遇到与 [i, j] 区间内字符相同的元素，更新 i 的值，此时 [i, j] 区间内不存在重复字符，计算 res 的最大值。

## Golang 实现及其测试代码

```go
package main

import "fmt"

func lengthOfLongestSubStr(str string) int {
	mathMax := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	cache := make(map[rune]int)
	var position, max int
	for i, r := range str {
		if num, ok := cache[r]; ok {
			position = mathMax(position, num+1)
		}
		cache[r] = i
		max = mathMax(max, i-position+1)
	}
	return max
}

func main() {
	str := "ababab"
	maxLen := lengthOfLongestSubStr(str)
	fmt.Println(maxLen)

	str2 := "abcabc"
	maxLen2 := lengthOfLongestSubStr(str2)
	fmt.Println(maxLen2)
}
```
