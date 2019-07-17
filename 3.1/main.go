/*
用递归和非递归方式，分别按照二叉树先序、 中序和后序打印所有的节点。
我们约定：先序遍历顺序为根、左、右；中序遍历顺序为左、根、右；后序遍历顺序为左、右、根
*/
package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/tree"
)

func main() {
	t := tree.MakeExampleTree()
	recursivePreTraverse(t)
	println("--------------")
	recursiveInTraverse(t)
	println("--------------")
	recursivePostTraverse(t)
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
