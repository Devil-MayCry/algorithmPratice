/*
给定一个链表的头节点head和一个整数num，请实现函数将值为num的节点全部删除。
例如，链表为1->2->3->4->null，num=3，链表调整后为：1->2->4->null。
*/

package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/linkedlist"
)

func main() {
	l := linkedlist.MakeChain([]int64{1, 2, 3, 3, 4, 4, 2, 1})
	l = removeNum(l, 1)
	l.PrintlnList()
}

func removeNum(h *linkedlist.Node, num int64) *linkedlist.Node {
	var pre, p *linkedlist.Node
	p = h
	for p != nil {
		if p.V == num {
			if pre != nil {
				pre.Next = p.Next
				p.Next = nil
				p = pre.Next
			} else {
				pre = p
				p = p.Next
				h = p
				pre.Next = nil
				pre = nil
			}
		} else {
			pre = p
			p = p.Next
		}
	}
	return h
}
