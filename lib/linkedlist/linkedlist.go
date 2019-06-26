package linkedlist

// Node 单链表节点
type Node struct {
	V    int64
	Next *Node
}

// TwoWayNode 双向链表
type TwoWayNode struct {
	V    int64
	Pre  *TwoWayNode
	Next *TwoWayNode
}

// MakeChain 构造一个单向链表
func MakeChain(arr []int64) *Node {
	var pre *Node
	var head *Node
	for k, va := range arr {
		n := &Node{V: va}
		if k == 0 {
			pre = n
			head = n
		} else {
			pre.Next = n
			pre = n
		}
	}
	return head
}

// MakeTwoWayChain 构造一个双向链表
func MakeTwoWayChain(arr []int64) *TwoWayNode {
	var pre *TwoWayNode
	var head *TwoWayNode
	for k, va := range arr {
		n := &TwoWayNode{V: va}
		if k == 0 {
			pre = n
			head = n
		} else {
			pre.Next = n
			n.Pre = pre
			pre = n
		}
	}
	return head
}

// MakeCircleChain 构造循环链表
func MakeCircleChain(arr []int64) *Node {
	var pre *Node
	var head *Node
	var tail *Node
	for k, va := range arr {
		tail = &Node{V: va}
		if k == 0 {
			pre = tail
			head = tail
		} else {
			pre.Next = tail
			pre = tail
		}
	}
	tail.Next = head
	return head
}

// FindListLen 计算链表长度
func (h *Node) FindListLen() int {
	if h == nil {
		return 0
	}
	res := 0
	for h.Next != nil {
		res++
		h = h.Next
	}
	res++
	return res
}

// FindListLen 计算链表长度
func (h *TwoWayNode) FindListLen() int {
	if h == nil {
		return 0
	}
	res := 0
	for h.Next != nil {
		res++
		h = h.Next
	}
	res++
	return res
}

// PrintlnList 打印单向链表
func (h *Node) PrintlnList() {
	if h.Next == nil {
		println(h.V)
		return
	}
	for h.Next != nil {
		println(h.V)
		h = h.Next
		if h.Next == nil {
			println(h.V)
		}
	}
}

// PrintlnList 打印双向链表
func (h *TwoWayNode) PrintlnList() {
	if h.Next == nil {
		println(h.V)
		return
	}
	for h.Next != nil {
		println(h.V)
		h = h.Next
		if h.Next == nil {
			println(h.V)
		}
	}
	for h.Pre != nil {
		println(h.V)
		h = h.Pre
		if h.Pre == nil {
			println(h.V)
		}
	}
}
