# [1207.独一无二的出现次数](https://leetcode-cn.com/problems/unique-number-of-occurrences/)

## 题目描述

给你一个整数数组 arr，请你帮忙统计数组中每个数的出现次数。

如果每个数的出现次数都是独一无二的，就返回 true；否则返回 false。

示例 1：

```
输入：arr = [1,2,2,1,1,3]
输出：true
解释：在该数组中，1 出现了 3 次，2 出现了 2 次，3 只出现了 1 次。没有两个数的出现次数相同。
```

示例 2：

```
输入：arr = [1,2]
输出：false
```

示例 3：

```
输入：arr = [-3,0,1,-3,1,1,1,-3,10,0]
输出：true
```

## 解法

### 前后分别使用 map 来存和统计数据

此题比较简单，这里就不再赘述。

### Golang 实现及其测试代码

```go
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
```
