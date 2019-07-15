/*
题目：
分别实现两个函数，一个可以删除单链表中倒数第k个结点，另一个可以删除双链表中倒数第K个节点
要求：
如果链表长度为N,时间复杂度达到O(N),额外空间复杂度达到O(1)。
*/
package main

// Node 单链表节点
type Node struct {
	v int64
	p *Node
}

// TwoWayNode 双链表节点
type TwoWayNode struct {
	v    int64
	pre  *TwoWayNode
	next *TwoWayNode
}

func makeSingleChain(arr []int64) *Node {
	var pre *Node
	var head *Node
	for k, va := range arr {
		n := &Node{v: va}
		if k == 0 {
			pre = n
			head = n
		} else {
			pre.p = n
			pre = n
		}
	}
	return head
}

func makeTwoWayChain(arr []int64) *TwoWayNode {
	var pre *TwoWayNode
	var head *TwoWayNode
	for k, va := range arr {
		n := &TwoWayNode{v: va}
		if k == 0 {
			pre = n
			head = n
		} else {
			pre.next = n
			n.pre = pre
			pre = n
		}
	}
	return head
}
func main() {
	arr := []int64{1, 2, 3, 4, 5, 6}
	sl := makeSingleChain(arr)
	dl := makeTwoWayChain(arr)
	sl = sl.deleteFromSingleList(6)
	dl = dl.deleteFromDoubleList(2)
	// 验证：
	// for sl.p != nil {
	// 	println(sl.v)
	// 	sl = sl.p
	// 	if sl.p == nil {
	// 		println(sl.v)
	// 	}
	// }
	// for dl.next != nil {
	// 	println(dl.v)
	// 	dl = dl.next
	// 	if dl.next == nil {
	// 		println(dl.v)
	// 	}
	// }
}

func (h *Node) deleteFromSingleList(k int) (nh *Node) {
	l := h.findListLen()
	// 错误的情况
	if l < k {
		return nil
	}
	// 删除的是头结点的情况
	index := l - k
	if index == 0 {
		return h.p
	}
	// 非头结点的情况
	pre := h
	for i := 0; i < index-1; i++ {
		pre = pre.p
	}
	pre.p = pre.p.p
	return h
}

func (h *TwoWayNode) deleteFromDoubleList(k int) (nh *TwoWayNode) {
	l := h.findListLen()
	// 错误的情况
	if l < k {
		return nil
	}
	// 删除的是头结点的情况
	index := l - k
	if index == 0 {
		h.next.pre = nil
		h = h.next
		return h
	}
	// 非头结点的情况
	pre := h
	for i := 0; i < index-1; i++ {
		pre = pre.next
	}

	if pre.next.next != nil {
		// 非最后一个节点
		po := pre.next
		pre.next = pre.next.next
		po.next.pre = pre
	} else {
		pre.next.pre = nil
		pre.next = nil
	}
	return h
}

func (h *Node) findListLen() int {
	if h == nil {
		return 0
	}
	res := 0
	for h.p != nil {
		res++
		h = h.p
	}
	res++
	return res
}

func (h *TwoWayNode) findListLen() int {
	if h == nil {
		return 0
	}
	res := 0
	for h.next != nil {
		res++
		h = h.next
	}
	res++
	return res
}
