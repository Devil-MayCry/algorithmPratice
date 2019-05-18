/*
给定两个有序链表的头指针head1和head2，打印两个链表的公共部分
*/
package main

import "fmt"

// Node ...
type Node struct {
	v int64
	p *Node
}

func makeChain(arr []int64) *Node {
	var pre = &Node{}
	var head *Node
	for k, va := range arr {
		n := &Node{v: va}
		if k == 0 {
			pre = n
			head = n
		} else {
			pre.p = n
			pre = pre.p
		}
	}
	return head
}

func main() {
	arr1 := []int64{1, 2, 4, 5, 6}
	arr2 := []int64{2, 3, 5, 7}
	chain1 := makeChain(arr1)
	chain2 := makeChain(arr2)
	cal(chain1, chain2)
}

func cal(head1 *Node, head2 *Node) {
	p1 := head1
	p2 := head2
	for p1 != nil && p2 != nil {
		if p1.v < p2.v {
			p1 = p1.p
		} else if p1.v == p2.v {
			fmt.Println(p1.v)
			p1 = p1.p
			p2 = p2.p
		} else {
			p2 = p2.p
		}
	}
}
