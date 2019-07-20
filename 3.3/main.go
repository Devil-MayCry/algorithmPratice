/*
二叉树可以用常规的三种遍历结果来描述其结构，但是不够直观，尤其是二叉树中有重复值的时候，
仅通过三种遍历的结果来构造二叉树的真实结构更是难上加难，有时则根本不可能。
给定一棵二叉树的头节点head，己知二叉树节点值的类型为32位整型，
请实现一个打印二叉树的函数，可以直观地展示树的形状，也便于画出真实的结构。
*/
package main

import (
	"fmt"
	"strings"

	"github.com/Devil-MayCry/algorithmPratice/lib/tree"
	"github.com/golang-collections/collections/stack"
)

var s *stack.Stack

func main() {
	t := tree.MakeExampleTree()
	printTree(t)
}

func printTree(head *tree.Node) {
	printInOrder(head, 0, "H", 17)
	println()
}

func printInOrder(head *tree.Node, height int, to string, len int) {
	if head == nil {
		return
	}
	printInOrder(head.Right, height+1, "v", len)
	val := fmt.Sprintf("%s%d%s", to, head.V, to)
	lenM := strings.Count(val, "")
	lenL := (len - lenM) / 2
	lenR := (len - lenM - lenL)
	val = getSpace(lenL) + val + getSpace(lenR)
	println(getSpace(height*len) + val)
	printInOrder(head.Left, height+1, "^", len)
}

func getSpace(num int) string {
	space := " "
	buf := ""
	for i := 0; i < num; i++ {
		buf = buf + space
	}
	return buf
}
