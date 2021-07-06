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