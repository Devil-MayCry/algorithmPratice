/*
给定一个单链表的头节点head,节点的值类型是整型，再给定一个整数pivot。实现一个调整链表的函数，
将链表调整为左部分都是值小于pivot的节点，
中间部分都是值等于pivot的节点，
右部分都是值大于pivot的节点。
除这个要求外，对调整后的节点顺序没有更多的要求。
*/
package main

import (
	"fmt"

	"github.com/Devil-MayCry/algorithmPratice/lib/linkedlist"
)

func main() {
	// l := linkedlist.MakeChain([]int64{1, 9, 6, 5, 3, 3, 4, 3, 2})
	l := linkedlist.MakeChain([]int64{1, 2})
	pivot := int64(3)
	updatePivotLink(l, pivot)
	l.PrintlnList()
}

func updatePivotLink(h *linkedlist.Node, pivot int64) *linkedlist.Node {
	p := h
	var leftHead, leftTail, pivotHead, pivotTail, rightHead, rightTail *linkedlist.Node
	for p != nil {
		if p.V < pivot {
			if leftHead == nil {
				leftHead = p
				leftTail = p
			} else {
				leftTail.Next = p
				leftTail = p
			}
		} else if p.V == pivot {
			if pivotHead == nil {
				pivotHead = p
				pivotTail = p
			} else {
				pivotTail.Next = p
				pivotTail = p
			}
		} else {
			if rightHead == nil {
				rightHead = p
				rightTail = p
			} else {
				rightTail.Next = p
				rightTail = p
			}
		}
		next := p.Next
		p.Next = nil
		p = next
	}
	nHead, nTail := mergeTwoLink(leftHead, leftTail, pivotHead, pivotTail)
	h, _ = mergeTwoLink(nHead, nTail, rightHead, rightTail)
	return h
}

func mergeTwoLink(a *linkedlist.Node, aTail *linkedlist.Node, b *linkedlist.Node, bTail *linkedlist.Node) (*linkedlist.Node, *linkedlist.Node) {
	fmt.Println(a, aTail, b)
	if a == nil {
		return b, bTail
	}
	if b == nil {
		return a, aTail
	}
	aTail.Next = b
	return a, bTail
}
