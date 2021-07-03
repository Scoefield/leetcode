/*
 * @lc app=leetcode.cn id=1207 lang=golang
 *
 * [1207] 独一无二的出现次数
 */

// @lc code=start
func uniqueOccurrences(arr []int) bool {
	if arr == nil || len(arr) == 0 {
		return false
	}
	map1 := make(map[int]int)
	for _, val := range arr {
		map1[val]++
	}
	map2 := make(map[int]int)
	for _, val := range map1 {
		_, ok := map2[val]
		if ok {
			return false
		}
		map2[val]++
	}
	return true
}
// @lc code=end

