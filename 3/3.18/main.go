/*
从二叉树的节点A出发，可以向上或者向下走，但沿途的节点只能经过一次，
当到达节点B时，路径上的节点数叫作A到B的距离。
求整个二叉树节点间的最大距离
*/

package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/math"
	"github.com/Devil-MayCry/algorithmPratice/lib/tree"
)

func main() {
	t := tree.MakeExampleTree()
	maxDistance(t)
}

func maxDistance(h *tree.Node) int64 {
	record := make([]int64, 1)
	return posOrder(h, record)
}

func posOrder(h *tree.Node, record []int64) int64 {
	if h == nil {
		record[0] = 0
		return 0
	}
	lMax := posOrder(h.Left, record)
	maxFromLeft := record[0]
	rMax := posOrder(h.Right, record)
	maxFromRight := record[0]
	curNodeMax := maxFromLeft + maxFromRight + 1
	record[0] = math.Max(maxFromRight, maxFromLeft) + 1
	return math.Max(math.Max(lMax, rMax), curNodeMax)
}
