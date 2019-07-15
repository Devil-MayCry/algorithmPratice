/*
给定一个链表的头节点head,请判断该链表是否为回文结构。

1->2->1, true;
1->2->2->1, true;
1->2->3, false;
*/
package main

import (
	"fmt"

	"github.com/Devil-MayCry/algorithmPratice/lib/linkedlist"
)

func main() {
	l := linkedlist.MakeChain([]int64{1, 2, 3, 1, 3, 2, 1})
	b := isPlindrome(l)
	fmt.Println(b)
}

func isPlindrome(h *linkedlist.Node) bool {
	if h == nil {
		return false
	}
	if h.Next == nil {
		return true
	}
	// 找到中点
	fast := h
	slow := h
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// 反转后半部分
	m := slow
	var pre, post *linkedlist.Node
	pre = m
	post = m.Next
	m.Next = nil
	m = post
	for m != nil {
		post = m.Next
		m.Next = pre
		pre = m
		m = post
	}
	// 开始比较
	tail1 := h
	tail2 := pre

	for tail1 != nil && tail2 != nil {
		if tail1.V != tail2.V {
			return false
		}
		tail1 = tail1.Next
		tail2 = tail2.Next
	}
	return true
}
