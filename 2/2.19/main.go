/*
给定两个有序单链表的头节点head1和head2，请合并两个有序链表，
合并后的链表依然有序，并返回合并后链表的头节点。
*/

package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/linkedlist"
)

func main() {
	l1 := linkedlist.MakeChain([]int64{1, 3, 5, 7})
	l2 := linkedlist.MakeChain([]int64{2, 4, 6, 8})
	mergeChain(l1, l2).PrintlnList()
}

func mergeChain(l1, l2 *linkedlist.Node) *linkedlist.Node {
	if l1 == nil {
		return l2
	}
	if l1 == nil {
		return l1
	}
	var newHead, p *linkedlist.Node
	p1 := l1
	p2 := l2
	if l1.V < l2.V {
		newHead = l1
		p1 = p1.Next
	} else {
		newHead = l2
		p2 = p2.Next
	}
	p = newHead
	for p1 != nil && p2 != nil {
		if p1.V < p2.V {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}
	if p1 == nil {
		p.Next = p2
	} else {
		p.Next = p1
	}
	return newHead
}
