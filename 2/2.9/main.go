/*
Node类中的value是节点值，next指针和正常单链表中next指针的含义一样，都指向下一个节点，
rand指针是Node类中新增的指针，
这个指针可能指向链表中的任意一个节点，也可能指向null。
给定一个由Node节点类型组成的无环单链表的头节点head，请实现一个函数完成这个链表中所有节点的复制。
*/
package main

import "fmt"

type Node struct {
	V    int64
	Next *Node
	Rand *Node
}

func makeChain() *Node {
	c := &Node{
		V: 3,
	}
	b := &Node{
		V:    2,
		Next: c,
	}
	a := &Node{
		V:    1,
		Next: b,
	}
	c.Rand = b
	b.Rand = a
	return a
}

func copyChain(h *Node) *Node {
	var p, next *Node
	p = h
	for p != nil {
		next = p.Next
		newNode := &Node{
			V: p.V,
		}
		p.Next = newNode
		newNode.Next = next
		p = next
	}

	p = h
	for p != nil {
		if p.Rand != nil {
			newRand := p.Rand.Next
			p.Next.Rand = newRand
		}
		p = p.Next.Next
	}

	p = h
	newH := p.Next
	for p != nil && p.Next != nil {
		next = p.Next
		p.Next = p.Next.Next
		p = next
	}
	return newH
}

func printChain(h *Node) {
	for h != nil {
		fmt.Println(h.V)
		fmt.Println(h.Rand)
		h = h.Next
	}
}

func main() {
	h := makeChain()
	newC := copyChain(h)
	printChain(newC)
}
