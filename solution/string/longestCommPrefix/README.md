# [最长公共前缀](https://leetcode-cn.com/problems/longest-common-prefix)

## 题目描述

编写一个函数来查找字符串数组中的最长公共前缀，如果不存在公共前缀，返回空字符串 ""。

示例 1：

```
输入：strs = ["abcdef", "abcd", "abcdafg"]
输出："abcd"
```

示例 2：

```
输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。
```

提示：

- 0 <= strs.length <= 200
- 0 <= strs[i].length <= 200
- strs[i] 仅由小写英文字母组成

## 思路及解法（横向扫描）

依次遍历字符串数组中的每个字符串，对于每个遍历到的字符串，更新最长公共前缀，当遍历完所有的字符串以后，即可得到字符串数组中的最长公共前缀。

## Golang 实现及其测试代码

```go
package main

import "fmt"

func longestCommPrefix(strSlice []string) string {
	var ret string
	for i, val := range strSlice {
		if i == 0 {
			ret = val
		}
		index := 0
		for index < len(ret) && index < len(val) && ret[index] == val[index] {
			index++
		}
		ret = ret[:index]
	}
	return ret
}

func main() {
	strSlice := []string{"abcdef", "abcd", "abcdafg"}
	commPrefix := longestCommPrefix(strSlice)
	fmt.Println(commPrefix)		// abcd

	strSlice2 := []string{"dog", "racecar", "car"}
	commPrefix2 := longestCommPrefix(strSlice2)
	fmt.Println(commPrefix2)	// ""
}
```
