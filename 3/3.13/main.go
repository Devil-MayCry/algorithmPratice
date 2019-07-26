/*
平衡二叉树的性质为：要么是一棵空树，要么任何一个节点的左右子树高度差的绝对值不超过1。
给定一棵二叉树的头节点head,判断这棵二叉树是否为平衡二叉树。
*/

package main

import (
	"math"

	m "github.com/Devil-MayCry/algorithmPratice/lib/math"
	"github.com/Devil-MayCry/algorithmPratice/lib/tree"
)

func main() {
	t := tree.MakeExampleTree()
	isBalance(t)
}

func isBalance(head *tree.Node) {
	res := make([]bool, 1)
	res[0] = true
}

func getHeight(head *tree.Node, level int64, res []bool) int64 {
	if head == nil {
		return level
	}
	lH := getHeight(head.Left, level+1, res)
	if !res[0] {
		return level
	}
	rH := getHeight(head.Right, level+1, res)
	if !res[0] {
		return level
	}
	if math.Abs(float64(lH-rH)) > 1 {
		res[0] = false
	}
	return m.Max(lH, rH)
}
