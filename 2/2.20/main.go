/*
给定一个单链表的头部节点head,链表长度为N，
如果N为偶数，那么前N/2个节点算作左半区，后N/2个节点算作右半区；
如果N为奇数，那么前N/2个节点算作左半区，后N/2+1个节点算作右半区。
左半区从左到右依次记为Ll->L2->…，右半区从左到右依次记为Rt->R2->…，
请将单链表调整成L1->R1->L2->R2->…的形式。
*/
package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/linkedlist"
)

func main() {
	l := linkedlist.MakeChain([]int64{1, 2, 3, 4, 5, 6, 7})
	adjustChain(l)
	l.PrintlnList()
}

func adjustChain(h *linkedlist.Node) {
	p := h
	m := findMiddle(h)
	mi := m
	leftTurn := true
	for p != mi {
		if leftTurn {
			pN := p.Next
			p.Next = m
			p = pN
		} else {
			mN := m.Next
			m.Next = p
			m = mN
		}
		leftTurn = !leftTurn
	}
}

// 通过快慢指针找到中点，
func findMiddle(h *linkedlist.Node) *linkedlist.Node {
	var slow, fast *linkedlist.Node
	if h == nil {
		return nil
	}
	if h.Next == nil {
		return h
	}
	slow = h
	fast = h
	for {
		if fast == nil || fast.Next == nil {
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
