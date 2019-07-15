/*
链表节点值类型为int型，给定一个链表中的节点node,但不给定整个链表的头节点。
如何在链表中删除node？请实现这个函数，并并分析这么会出现哪些问题。
要求：时间复杂度为O(1)。
*/
package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/linkedlist"
)

func main() {
	l := linkedlist.MakeChain([]int64{1, 2, 3, 5, 3, 4, 4, 2, 1})
	n := l.Next.Next.Next // n是5
	strangeDel(n)
	l.PrintlnList()
}

// 把从这个节点后，每个节点都改为之后的节点的数值，再删除最后一个节点，以此达到“删除”
func strangeDel(n *linkedlist.Node) {
	var pre *linkedlist.Node
	for n != nil && n.Next != nil {
		n.V = n.Next.V
		pre = n
		n = n.Next
	}
	if n != nil {
		pre.Next = nil
	}
}
