# [25.K 个一组翻转链表](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/)

## 题目描述

给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

进阶：

你可以设计一个只使用常数额外空间的算法来解决此问题吗？
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

示例一：

![1](./images/reverse_ex1.jpeg)

```
输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]
```

示例二：

![2](./images/reverse_ex2.jpeg)

```
输入：head = [1,2,3,4,5], k = 3
输出：[3,2,1,4,5]
```

示例三：

```
输入：head = [1,2,3,4,5], k = 1
输出：[1,2,3,4,5]
```

## 思路与算法

参考：[力扣官网解题思路](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/solution/k-ge-yi-zu-fan-zhuan-lian-biao-by-leetcode-solutio/)

## Golang 实现

```go
func reverseKGroup(head *Node, k int) *Node {
	hair := &Node{Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = reverseNode(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func reverseNode(head, tail *Node) (*Node, *Node) {
	prev := tail.Next
	cur := head
	for prev != tail {
		cur, cur.Next, prev = cur.Next, prev, cur
	}
	return head, tail
}
```
