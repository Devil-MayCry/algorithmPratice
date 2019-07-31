/*
己知一棵二叉树所有的节点值都不同，给定这棵树正确的先序和中序数组，
不要重建整棵树，而是通过这两个数组直接生成正确的后序数组。
*/

package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/tree"
)

func main() {
	t := tree.MakeExampleTree()
	nodeNum(t)
}

func nodeNum(h *tree.Node) int {
	if h == nil {
		return 0
	}
	return bs(h, 1, mostLeftLevel(h, 1))
}
func bs(n *tree.Node, l int, h int) int {
	if l == h {
		return 1
	}
	if mostLeftLevel(n.Right, l+1) == h {
		return 1<<uint(h-l) + bs(n.Right, l+1, h)
	} else {
		return 1<<uint(h-l-1) + bs(n.Left, l+1, h)
	}
}

func mostLeftLevel(n *tree.Node, level int) int {
	for n != nil {
		level++
		n = n.Left
	}
	return level - 1
}
