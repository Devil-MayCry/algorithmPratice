/*
给定一个单链表的头节点head，实现一个调整单链表的函数，使得每K个节点之间逆序，
如果最后不够K个节点，则不调制最后几个节点。

例如：
链表：1->2->3->4->5->6->7->8->null,K=3 o
调整后为：3->2->1->6->5->4->7->8->null。其中7、8不调整，因为不够一组
*/

package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/linkedlist"
)

func main() {
	l := linkedlist.MakeChain([]int64{1, 2, 3, 4, 5, 6, 7, 8})
	l = adjustChain(l, 3)
	l.PrintlnList()
}

func adjustChain(h *linkedlist.Node, k int) *linkedlist.Node {
	len := h.FindListLen()
	if len < k || k <= 1 {
		return h
	}
	var newHead, preSegNext, newTail *linkedlist.Node

	for r := k; r <= len; r = r + k {
		head, tail, next := tranChain(h, k)
		if newHead == nil {
			newHead = head
		}
		if preSegNext != nil {
			preSegNext.Next = head
		} else {
			preSegNext = tail
		}
		h = next
		newTail = tail
	}
	if h != nil {
		newTail.Next = h
	}
	return newHead
}

// 把从p开头的m个节点反转。返回头指针和尾指针
func tranChain(p *linkedlist.Node, k int) (head, tail, next *linkedlist.Node) {
	var pre, ne *linkedlist.Node
	tail = p
	for i := 0; i < k; i++ {
		ne = p.Next
		p.Next = pre
		pre = p
		p = ne
	}
	head = pre
	next = p
	return
}
