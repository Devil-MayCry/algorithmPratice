/*
在本题中，单链表可能有环，也可能无环。给定两个单链表的头节点headl和head2,这两个链表可能相交，也可能不相交。
请实现一个函数，如果两个链表相交，请返回相交的第一个节点；如果不相交，返回null即可。

要求：如果链表l的长度为N，链表2的长度为M，时间复杂度请达到O(N+M)，额外空间复杂度请达到0(1)。
*/

package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/linkedlist"
)

// 判断A和B链表是否有环
// 如果两个均无环，连接首尾正常判断
// 如果一个有环，一个无环，不可能相交
// 如果均有环：有可能三种情况 ： 相交点再环开始前 / 入口处 / 环内

// 当均无环时，怎么求出相交点？分别遍历两个链表，记录链表长度len1，len2。
// 用两个指针指向两个链表的头部，让长度较长的链表先走|len1-len2|步，然后两个指针共同走，当两个指针相等时，即为第一个相交链表。

// 当有环是，怎么怎么相交
// 对其中一个用快慢指针求相遇点，相遇点必在环上。然后检测另一个链表的环是否包含这个点

// 当有环时，怎么求出相交点？
// 如果相交点在，相交点再环开始前 / 入口处。用|len1-len2|的方法就可求出
// 如果相交点再环内，对于每个环来说，第一个相交点就是分别的第一个入环节点
func main() {

}

// 判断链表是否有环
func isCircle(h *linkedlist.Node) bool {
	if h == nil || h.Next == nil {
		return false
	}
	slow := h
	fast := h
	for fast != nil && fast.Next != nil {
		slow := slow.Next
		fast := fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

/*
用快慢指针求入口点
因为快指针速度是满指针2倍，可证明相交点必在环内，而且是在慢指针第一圈里
设入口前长度为w，入口点到第一个相交点长度为y，环长度为n，快指针相遇前已绕k圈
则有
w + y + kn = 2 (w + y)
求得 kn - y = w
所以相遇后把快指针放回头结点，速度与满指针一致，继续前进，则会在入口节点再次相遇
*/
func getCircleJoin(h *linkedlist.Node) *linkedlist.Node {
	slow := h.Next
	fast := h.Next.Next
	for fast != slow {
		slow = slow.Next
		fast = fast.Next.Next
	}
	fast = h
	for fast != slow {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}

// 判断两个无环链表是否相交，求交点
func getTwoNoCircleJoin(h1 *linkedlist.Node, h2 *linkedlist.Node) *linkedlist.Node {
	l1 := h1.FindListLen()
	l2 := h2.FindListLen()
	if l1 < l2 {
		h1, h2 = switchTwoPointer(h1, h2)
	}
	// h1 先走 l1-l2步
	for i := 0; i < l1-l2; i++ {
		h1 = h1.Next
	}
	for h1 != nil {
		if h1 == h2 {
			return h1
		}
		h1 = h1.Next
		h2 = h2.Next
	}
	return nil
}

func switchTwoPointer(h1 *linkedlist.Node, h2 *linkedlist.Node) (*linkedlist.Node, *linkedlist.Node) {
	tmp := h1
	h1 = h2
	h2 = tmp
	return h1, h2
}

// 求两个有环链表交点
func getTwoCircleJoin(h1 *linkedlist.Node, h2 *linkedlist.Node) *linkedlist.Node {
	h1Join := getCircleJoin(h1)
	h2Join := getCircleJoin(h2)

	// 判断是否在交点相交
	if h1Join == h2Join {
		return h1Join
	}

	// 判断是否在环前相交
	for h2 != h2Join {
		if h2 == h1Join {
			return h2
		}
		h2 = h2.Next
	}

	h2 = h2Join.Next
	// 判断是否在环上相交
	for h2 != h2Join {
		if h2 == h1Join {
			return h2
		}
		h2 = h2.Next
	}
	return nil
}
