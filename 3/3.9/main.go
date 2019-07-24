/*
给定一棵二叉树的头节点head,分别实现按层打印和ZigZag打印二叉树的函数。
*/

package main

import (
	"fmt"

	"github.com/Devil-MayCry/algorithmPratice/lib/tree"
	"github.com/golang-collections/collections/queue"
)

func main() {
	t := tree.MakeExampleTree()
	printByLevel(t)
}

func printByLevel(head *tree.Node) {
	if head == nil {
		return
	}
	q := queue.New()
	level := 1
	last := head
	var nLast *tree.Node
	q.Enqueue(head)
	fmt.Printf("level %d:", level)
	level++
	for q.Len() > 0 {
		head = q.Dequeue().(*tree.Node)
		fmt.Printf("%d :", head.V)
		if head.Left != nil {
			q.Enqueue(head.Left)
			nLast = head.Left
		}
		if head.Left != nil {
			q.Enqueue(head.Right)
			nLast = head.Right
		}
		if head == last && q.Len() > 0 {
			fmt.Printf("level %d:", level)
			level++
			last = nLast
		}
	}
}
