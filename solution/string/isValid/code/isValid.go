/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
// 解法一
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

// 解法二
func isValid2(s string) bool {
	if len(s) == 0 || len(s) == 1{
		return false
	}
	mapLeft := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if mapLeft[s[i]] > 0 {
			if len(stack) > 0 && stack[len(stack)-1] == mapLeft[s[i]] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
// @lc code=end