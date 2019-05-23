/*
给定一个单向链表的头节点head，以及两个链表from和to，在单向链表上把第from个节点到第to个节点这一部分进行反转。

1->2->3->4->5->null,from=2,to=4
调整结果为：1->4->3->2->5->null
*/
package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/linkedlist"
)

func main() {
	l := linkedlist.MakeChain([]int64{1, 2, 3, 4, 5, 6})
	rl := reversePartChain(l, 1, 6)
	rl.PrintlnList()
}

func reversePartChain(head *linkedlist.Node, from int, to int) *linkedlist.Node {
	if from > head.FindListLen() || (from >= to) || from < 1 {
		return head
	}
	var prePoint *linkedlist.Node
	p := head
	// 移动到指定位置
	for i := 1; i < from; i++ {
		prePoint = p
		p = p.Next
	}
	tail := p
	pre := p
	p = p.Next
	for i := from; i < to; i++ {
		next := p.Next
		p.Next = pre
		pre = p
		p = next
	}
	if prePoint != nil {
		prePoint.Next = pre
		tail.Next = p
		return head
	}
	tail.Next = p
	return pre
}
