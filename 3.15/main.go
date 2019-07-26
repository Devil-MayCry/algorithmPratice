/*
给定一个二叉树的头节点head,已知其中没有重复值的节点，
实现两个函数分别判断这棵二叉树是否是搜索二叉树和完全二叉树
*/

package main

import (
	"math"

	"github.com/Devil-MayCry/algorithmPratice/lib/tree"
	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
)

func main() {
	h := tree.MakeExampleTree()
	isBST(h)
	isCBT(h)
}

// 判断是否是搜索二叉树
func isBST(t *tree.Node) bool {
	s := stack.New()
	res := true
	last := int64(math.MinInt64)
	if t == nil {
		return false
	}
	head := t
	for s.Len() > 0 || head != nil {
		if head != nil {
			s.Push(head)
			head = head.Left
		} else {
			head = s.Pop().(*tree.Node)
			if head.V <= last {
				res = false
				break
			}
			head = head.Right
		}
	}
	return res
}

// 判断是否是完全二叉树
func isCBT(head *tree.Node) bool {
	if head == nil {
		return true
	}
	q := queue.New()
	leaf := false
	var l *tree.Node
	var r *tree.Node
	q.Enqueue(head)

	for q.Len() > 0 {
		if (leaf && (l != nil || r != nil)) || (l == nil && r != nil) {
			return false
		}
		if l != nil {
			q.Enqueue(l)
		}
		if r != nil {
			q.Enqueue(r)
		} else {
			leaf = true
		}
	}
	return true
}
