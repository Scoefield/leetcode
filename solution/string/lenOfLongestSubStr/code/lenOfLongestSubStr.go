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
