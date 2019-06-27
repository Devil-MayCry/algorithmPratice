/*
假设链表中每一个结点的值都在0~9之间，那么链表整体就可以代表一个整数
例如：9->3->7，可以代表整数937
给定两个这种链表的头结点head1和head2，请生成代表两个整数相加值的结果链表。
*/
package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/linkedlist"
)

func main() {
	l1 := linkedlist.MakeChain([]int64{1, 1, 5, 7, 4, 5})
	l2 := linkedlist.MakeChain([]int64{5, 4, 3, 2, 1})
	l := mergeLinks(l1, l2)
	l = reverseLink(l)
	l.PrintlnList()
}

func mergeLinks(l1 *linkedlist.Node, l2 *linkedlist.Node) *linkedlist.Node {
	// 先反转，再相加
	r1 := reverseLink(l1)
	r2 := reverseLink(l2)
	len1 := r1.FindListLen()
	len2 := r2.FindListLen()
	// 保证r2不比r1短
	if len1 > len2 {
		tmp := r1
		r1 = r2
		r2 = tmp
	}

	var jinWei int64
	var r1Num int64
	newHead := r2
	for r2 != nil {
		if r1 == nil {
			r1Num = 0
		} else {
			r1Num = r1.V
			r1 = r1.Next
		}
		v := r1Num + jinWei + r2.V
		newR2Num := v % 10
		jinWei = v / 10
		r2.V = newR2Num
		r2 = r2.Next
	}
	if jinWei != 0 {
		n := &linkedlist.Node{}
		n.V = 1
		n.Next = nil
		r2.Next = n
	}
	return newHead
}

func reverseLink(h *linkedlist.Node) *linkedlist.Node {
	if h == nil {
		return nil
	}
	if h.Next == nil {
		return h
	}
	var p1, p2 *linkedlist.Node
	p1 = h.Next
	h.Next = nil
	for p1 != nil {
		p2 = p1.Next
		p1.Next = h
		h = p1
		p1 = p2
	}
	return h
}
