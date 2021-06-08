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
