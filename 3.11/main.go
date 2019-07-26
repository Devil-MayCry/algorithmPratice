/*
给定彼此独立的两棵树头节点分别为t1和t2，判断t1树是否包含t2树全部的拓扑结构。
*/

package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/tree"
)

func main() {
	t1 := tree.MakeExampleTree()
	t2 := tree.MakeExampleTree()
	contains(t1, t2)
}

func contains(t1 *tree.Node, t2 *tree.Node) bool {
	return check(t1, t2) || contains(t1.Left, t2) || contains(t1.Right, t2)
}

func check(h *tree.Node, t2 *tree.Node) bool {
	if t2 == nil {
		return true
	}
	if h == nil || h.V != t2.V {
		return false
	}
	return check(h.Left, t2.Left) && check(h.Right, t2.Right)
}
