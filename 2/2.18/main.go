/*
一个环形单链表从头节点head开始不降序，同时由最后的节点指回头节点。
给定这样一个环形单链表的头节点head和个整数num，请生成节点值为num的新节点，并插入到这个环形链表中，
保证调整后的链表依然有序。
*/
package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/linkedlist"
)

func main() {
	l := prepareData()
	num := int64(3)
	l = insertData(l, num)
	l.PrintlnList()
}

func insertData(h *linkedlist.Node, num int64) *linkedlist.Node {
	newNode := &linkedlist.Node{V: num}
	if h == nil {
		h = newNode
		h.Next = h
		return h
	}
	var p, pre *linkedlist.Node
	pre = h
	p = h.Next
	for p != h {
		if pre.V <= num && p.V > num {
			break
		}
		pre = p
		p = p.Next
	}
	pre.Next = newNode
	newNode.Next = p
	if h.V < num {
		return h
	} else {
		return newNode
	}
}

func prepareData() *linkedlist.Node {
	n1 := &linkedlist.Node{V: 1}
	n2 := &linkedlist.Node{V: 2}
	n3 := &linkedlist.Node{V: 2}
	n4 := &linkedlist.Node{V: 4}
	n5 := &linkedlist.Node{V: 5}
	n6 := &linkedlist.Node{V: 6}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n1
	return n1
}
