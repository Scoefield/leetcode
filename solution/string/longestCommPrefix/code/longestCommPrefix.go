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
