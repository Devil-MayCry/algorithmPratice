/*
给定一个无序单链表的头节点head，实现单链表的选择排序。
要求：额外空间复杂度为O(1)。
*/

package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/linkedlist"
)

func main() {
	l := linkedlist.MakeChain([]int64{1, 2, 3, 3, 4, 4, 2, 1})
	sort(l).PrintlnList()
}

func sort(h *linkedlist.Node) *linkedlist.Node {
	var isFinish bool
	for !isFinish {
		h, isFinish = adjustNode(h)
	}
	return h
}

//  遍历一遍链表。将所有位置有偏差的节点向前移动一位。返回是否已经全部有序
func adjustNode(h *linkedlist.Node) (*linkedlist.Node, bool) {
	isFinish := true
	var pre, newHead *linkedlist.Node
	for h != nil && h.Next != nil {
		if newHead == nil {
			newHead = h
		}
		if h.Next.V < h.V {
			isFinish = false
			next := h.Next
			nNext := h.Next.Next
			if pre != nil {
				pre.Next = next
				next.Next = h
			} else {
				newHead = next
				newHead.Next = h
			}
			pre = next
			h.Next = nNext
		} else {
			pre = h
			h = h.Next
		}
	}
	return newHead, isFinish
}
