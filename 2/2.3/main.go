/*
给定链表的头结点head，实现删除链表的中间节点的函数。

例如：
不删除任何节点；
1->2,           删除节点1
1->2->3,        删除节点2
1->2->3->4,     删除节点2
1->2->3->4->5,  删除节点3
给定链表的头节点head,整数a和整数b，实现删除位于a/b处节点的函数。
1->2->3->4->5，假设a/b的值为r。
如果r等于0，不删除任何节点；
如果r在区间(0,1/5]上，删除节点1；
如果r在区间(1/5,2/5]上，删除节点2；
如果r在区间(2/5,3/5]上，删除节点3；
如果r在区间(3/5,4/5]上，删除节点4；
如果r在区间(4/5,1]上，删除节点5；
如果r大于1，不删除任何节点。
*/

package main

import "math"

// Node ...
type Node struct {
	v int64
	p *Node
}

func makeChain(arr []int64) *Node {
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

func main() {
	arr := []int64{1, 2, 3, 4, 5}
	chain := makeChain(arr)
	chain.deleteMiddleNode()
	chain.deleteInsideNode(5, 5)
	for chain.p != nil {
		println(chain.v)
		chain = chain.p
		if chain.p == nil {
			println(chain.v)
		}
	}
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

func (h *Node) deleteMiddleNode() {
	if h == nil {
		return
	}
	fast := h
	slow := h
	pre := h
	for fast.p != nil && fast.p.p != nil {
		pre = slow
		slow = slow.p
		fast = fast.p.p
	}
	pre.p = pre.p.p
}

func (h *Node) deleteInsideNode(a int, b int) {
	if h == nil || a > b {
		return
	}
	len := h.findListLen()
	index := int(math.Ceil(float64(a)/float64(b)*float64(len))) - 1
	for i := 0; i < index-1; i++ {
		h = h.p
	}
	h.p = h.p.p
}
