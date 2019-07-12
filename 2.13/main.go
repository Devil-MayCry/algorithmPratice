/*
给定一个无序单链表的头节点head,删除其中值重复出现的节点。

例如：l->2->3->3->4->4->2->1->null，删除值重复的节点之后为l->2->3->4->null。
请按以下要求实现两种方法。
方法l：如果链表长度为N，时间复杂度达到O(N)。
方法2：额外空间复杂度为O(1)。
*/
package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/linkedlist"
)

func main() {
	l1 := linkedlist.MakeChain([]int64{1, 2, 3, 3, 4, 4, 2, 1})
	l2 := linkedlist.MakeChain([]int64{1, 2, 3, 3, 4, 4, 2, 1})
	cleanChain1(l1)
	cleanChain2(l2)
}

// 时间复杂度O(N)
// 用一个map即可
func cleanChain1(h *linkedlist.Node) {
	m := make(map[int64]bool)
	var pre *linkedlist.Node
	for h != nil {
		if m[h.V] {
			pre.Next = h.Next
		} else {
			pre = h
		}
		m[h.V] = true
		h = h.Next
	}
}

// 额外空间复杂度为O(1)
func cleanChain2(h *linkedlist.Node) {
	// 先排序，扫描移动到合适的位置后，再最后扫一遍，去除重复的
	var isFinish bool
	for !isFinish {
		h, isFinish = adjustNode(h)
	}
	removeDuplicate(h)
	h.PrintlnList()
}

func removeDuplicate(h *linkedlist.Node) {
	for h != nil && h.Next != nil {
		if h.V == h.Next.V {
			h.Next = h.Next.Next
		}
		h = h.Next
	}
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
