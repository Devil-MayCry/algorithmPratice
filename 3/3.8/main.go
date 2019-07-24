/*
给定一棵二叉树的头节点head,己知所有节点的值都不一样，
返回其中最大的且符合搜索二叉树条件的最大拓扑结构的大小。
*/

package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/math"
	"github.com/Devil-MayCry/algorithmPratice/lib/tree"
)

func main() {
	t := tree.MakeExampleTree()
	println(bstTopoSize(t))
}

func bstTopoSize(head *tree.Node) int64 {
	if head == nil {
		return 0
	}
	max := maxTopo(head, head)
	max = math.Max(bstTopoSize(head.Left), max)
	max = math.Max(bstTopoSize(head.Right), max)
	return max
}

func maxTopo(h *tree.Node, n *tree.Node) int64 {
	if h != nil && n != nil && isBSTNode(h, n, n.V) {
		return maxTopo(h, n.Left) + maxTopo(h, n.Right) + 1
	}
	return 0
}

func isBSTNode(h *tree.Node, n *tree.Node, value int64) bool {
	if h == nil {
		return false
	}
	if h == n {
		return true
	}
	var t *tree.Node
	if h.V > value {
		t = h.Left
	} else {
		t = h.Right
	}
	return isBSTNode(t, n, value)
}
