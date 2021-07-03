# [20.有效的括号](https://leetcode-cn.com/problems/valid-parentheses/)

## 题目描述

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

示例 1：

```
输入：s = "()"
输出：true
```

示例 2：

```
输入：s = "()[]{}"
输出：true
```

示例 3：

```
输入：s = "(]"
输出：false
```

示例 4：

```
输入：s = "([)]"
输出：false
```


示例 5：

```
输入：s = "{[]}"
输出：true
```

## 思路及解法

方法一：栈
判断括号的有效性可以使用「栈」这一数据结构来解决。

我们遍历给定的字符串 s。当我们遇到一个左括号时，我们会期望在后续的遍历中，有一个相同类型的右括号将其闭合。由于后遇到的左括号要先闭合，因此我们可以将这个左括号放入栈顶。

当我们遇到一个右括号时，我们需要将一个相同类型的左括号闭合。此时，我们可以取出栈顶的左括号并判断它们是否是相同类型的括号。如果不是相同的类型，或者栈中并没有左括号，那么字符串 s 无效，返回 False。为了快速判断括号的类型，我们可以使用哈希表存储每一种括号。哈希表的键为右括号，值为相同类型的左括号。
 
在遍历结束后，如果栈中没有左括号，说明我们将字符串 s 中的所有左括号闭合，返回 True，否则返回 False。

注意到有效字符串的长度一定为偶数，因此如果字符串的长度为奇数，我们可以直接返回 False，省去后续的遍历判断过程。


## Golang 实现及其测试代码

```go
/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	if len(s) == 0 {
		return false
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if isLeft(s[i]) {
			stack = append(stack, s[i])
		} else {
			if len(stack) > 0 && stack[len(stack)-1] == findLeft(s[i]) {
				stack = stack[:len(stack) - 1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func isLeft(s byte) bool {
	if s == '(' || s == '{' || s == '[' {
		return true
	}
	return false
}

// 这里也可以用一个 map 来存放
func findLeft(s byte) byte {
	if s == ')' {
		return '('
	}
	if s == ']' {
		return '['
	}
	if s == '}' {
		return '{'
	}
	return ' '
}

// @lc code=end
```
