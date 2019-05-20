/*
题目：

分别实现反转单向链表和反转双向链表的函数

要求：

如果链表长度为N，时间复杂度要求为O(N)，额外空间复杂度要求为O(1)
*/
package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/linkedlist"
)

func main() {
	l := linkedlist.MakeChain([]int64{1, 2, 3, 4, 5})
	rl := reverseChain(l)
	rl.PrintlnList()
	tl := linkedlist.MakeTwoWayChain([]int64{1, 2, 3, 4, 5})
	rtl := reverseDoubleChain(tl)
	rtl.PrintlnList()
}

func reverseChain(head *linkedlist.Node) *linkedlist.Node {
	if head == nil && head.Next == nil {
		return nil
	}
	pre := head
	head = head.Next
	pre.Next = nil
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func reverseDoubleChain(head *linkedlist.TwoWayNode) *linkedlist.TwoWayNode {
	if head == nil && head.Next == nil {
		return nil
	}
	pre := head
	head = head.Next
	pre.Next = nil
	for head != nil {
		next := head.Next
		head.Next = pre
		pre.Pre = head
		pre = head
		head = next
	}
	pre.Pre = nil
	return pre
}
