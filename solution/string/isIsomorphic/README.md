# [205.同构字符串](https://leetcode-cn.com/problems/isomorphic-strings/)

## 题目描述

给定两个字符串 s 和 t，判断它们是否是同构的。

如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。

每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。

示例 1：

```
输入：s = "egg", t = "add"
输出：true
```

示例 2：

```
输入：s = "foo", t = "bar"
输出：false
```

示例 3：

```
输入：s = "paper", t = "title"
输出：true
```


提示：

- 可以假设 s 和 t 长度相同。

## 思路及解法（横向扫描）

依次遍历字符串数组中的每个字符串，对于每个遍历到的字符串，更新最长公共前缀，当遍历完所有的字符串以后，即可得到字符串数组中的最长公共前缀。

## Golang 实现及其测试代码

```go
/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
		return false
	}
    return isIsomorphicHelper(s, t) && isIsomorphicHelper(t, s)
}

func isIsomorphicHelper(s string, t string) bool {
	p := 0
	dic := make(map[byte]byte)
	for p <= len(s) - 1 {
		_, ok := dic[s[p]]; 
        if !ok {
			dic[s[p]] = t[p]
		} else {
			if dic[s[p]] != t[p] {
				return false
			}
		}
		p++
	}
	return true
}

// @lc code=end
```
