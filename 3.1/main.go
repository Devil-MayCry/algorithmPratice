/*
用递归和非递归方式，分别按照二叉树先序、 中序和后序打印所有的节点。
我们约定：先序遍历顺序为根、左、右；中序遍历顺序为左、根、右；后序遍历顺序为左、右、根
*/
package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/tree"
	"github.com/golang-collections/collections/stack"
)

var s *stack.Stack

func main() {
	t := tree.MakeExampleTree()
	// recursivePreTraverse(t)
	// println("--------------")
	// recursiveInTraverse(t)
	// println("--------------")
	// recursivePostTraverse(t)
	// println("--------------")
	// recursiveInNotTraverse
	// println("--------------")
	// recursivePostNotTraverse(t)
	// println("--------------")
	recursivePostNotTraverseOneStack(t)
}

// 递归先序
func recursivePreTraverse(t *tree.Node) {
	if t == nil {
		return
	}
	println(t.V)
	if t.Left != nil {
		recursivePreTraverse(t.Left)
	}
	if t.Right != nil {
		recursivePreTraverse(t.Right)
	}
}

func recursivePreNotTraverse(t *tree.Node) {
	s = stack.New()
	if t == nil {
		return
	}
	s.Push(t)
	for s.Len() > 0 {
		n := s.Pop().(*tree.Node)
		println(n.V)
		if n.Right != nil {
			s.Push(n.Right)
		}
		if n.Left != nil {
			s.Push(n.Left)
		}
	}
}

func recursiveInNotTraverse(t *tree.Node) {
	s = stack.New()
	if t == nil {
		return
	}
	head := t
	for s.Len() > 0 || head != nil {
		if head != nil {
			s.Push(head)
			head = head.Left
		} else {
			head = s.Pop().(*tree.Node)
			println(head.V)
			head = head.Right
		}
	}
}

func recursivePostNotTraverse(t *tree.Node) {
	s1 := stack.New()
	s2 := stack.New()
	if t == nil {
		return
	}
	s1.Push(t)
	for s1.Len() > 0 {
		n := s1.Pop().(*tree.Node)
		s2.Push(n)
		if n.Left != nil {
			s1.Push(n.Left)
		}
		if n.Right != nil {
			s1.Push(n.Right)
		}
	}
	for s2.Len() > 0 {
		println(s2.Pop().(*tree.Node).V)
	}
}

func recursivePostNotTraverseOneStack(h *tree.Node) {
	s := stack.New()
	if h == nil {
		return
	}
	s.Push(h)
	var c *tree.Node
	for s.Len() > 0 {
		c = s.Peek().(*tree.Node)
		if c.Left != nil && h != c.Left && h != c.Right {
			s.Push(c.Left)
		} else if c.Right != nil && h != c.Right {
			s.Push(c.Right)
		} else {
			println(s.Pop().(*tree.Node).V)
			h = c
		}
	}
}

// 递归中序
func recursiveInTraverse(t *tree.Node) {
	if t == nil {
		return
	}
	if t.Left != nil {
		recursiveInTraverse(t.Left)
	}
	println(t.V)
	if t.Right != nil {
		recursiveInTraverse(t.Right)
	}
}

// 递归后序
func recursivePostTraverse(t *tree.Node) {
	if t == nil {
		return
	}
	if t.Left != nil {
		recursivePostTraverse(t.Left)
	}
	if t.Right != nil {
		recursivePostTraverse(t.Right)
	}
	println(t.V)
}
