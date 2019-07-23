/*
给定一棵二叉树的头节点head,已知其中所有结点的值都不一样，
找到含有节点最多的搜索二叉树，并返回这棵子树的头节点。
*/

package main

import (
	"math"

	mathlib "github.com/Devil-MayCry/algorithmPratice/lib/math"
	"github.com/Devil-MayCry/algorithmPratice/lib/tree"
)

func main() {
	t := tree.MakeExampleTree()
	res := biggestSubBST(t)
	res.PrintlnTree()
}

func biggestSubBST(head *tree.Node) *tree.Node {
	record := make([]int64, 3)
	return posOrder(head, record)
}

func posOrder(head *tree.Node, record []int64) *tree.Node {
	if head == nil {
		record[0] = 0
		record[1] = math.MaxInt64
		record[2] = math.MinInt64
		return nil
	}
	value := head.V
	left := head.Left
	right := head.Right
	lBST := posOrder(left, record)
	lSize := record[0]
	lMin := record[1]
	lMax := record[2]
	rBST := posOrder(right, record)
	rSize := record[0]
	rMin := record[1]
	rMax := record[2]
	record[1] = mathlib.Min(lMin, value)
	record[2] = mathlib.Max(rMax, value)
	if left == lBST && right == rBST && lMax < value && value < rMin {
		record[0] = lSize + rSize + 1
		return head
	}
	record[0] = mathlib.Max(lSize, rSize)
	if lSize > rSize {
		return lBST
	}
	return rBST
}
