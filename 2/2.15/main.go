/*
对二叉树的节点来说，有本身的值域，有指向左孩子和右孩子的两个指针；
对双向链表的节点来说，有本身的值域，有指向上一个节点和下一个节点的指针。
结构上，两种结构由相似性，现在有一棵搜索二叉树，请将其转换为一个有序的双向链表。
*/

package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/linkedlist"
	"github.com/Devil-MayCry/algorithmPratice/lib/tree"
)

var c chan int64

func main() {
	c = make(chan int64)
	t := makeTree()
	h := transTreeToChain(t)
	h.PrintlnList()
}

func transTreeToChain(t *tree.Node) *linkedlist.TwoWayNode {
	var h, p *linkedlist.TwoWayNode
	go startTrave(t)
	for v := range c {
		if h == nil {
			h = &linkedlist.TwoWayNode{V: v}
			p = h
		} else {
			tn := &linkedlist.TwoWayNode{V: v}
			p.Next = tn
			tn.Pre = p
			p = tn
		}
	}
	return h
}

func travesTree(t *tree.Node) {
	if t == nil {
		return
	}
	if t.Left != nil {
		travesTree(t.Left)
	}
	visitNode(t)
	if t.Right != nil {
		travesTree(t.Right)
	}
}
func startTrave(t *tree.Node) {
	travesTree(t)
	end()
}

func end() {
	close(c)
}

func visitNode(t *tree.Node) {
	c <- t.V
}

func makeTree() *tree.Node {
	a := &tree.Node{V: 7}
	b := &tree.Node{V: 3}
	c := &tree.Node{V: 8}
	d := &tree.Node{V: 2}
	e := &tree.Node{V: 5}
	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	return a
}
