/*
给定一棵二叉树的头节点head，按照如下两种标准分别实现二叉树边界节点的逆时针打印。

标准一：
1. 头节点为边界节点
2. 叶节点为边界节点
3. 如果节点在其所在的层中是最左或最右的，那么也是边界节点。
标准二：
1. 头节点为边界节点
2. 叶节点为边界节点
3. 树左边界延伸下去的路径为边界节点。
4. 树右边界延伸下去的路径为边界节点。
要求：

如果节点数为N，两种标准实现的时间复杂度要求都为O(N)，额外空间复杂度要求都为O(h)，h为二叉树的高度。

两种标准都要求逆时针顺序且不重复打印所有的边界节点。
*/
package main

import (
	"math"

	"github.com/Devil-MayCry/algorithmPratice/lib/tree"
	"github.com/golang-collections/collections/stack"
)

var s *stack.Stack

func main() {
	t := tree.MakeExampleTree()
	printEdge1(t)
}

func printEdge1(head *tree.Node) {
	if head == nil {
		return
	}
	height := getHight(head, 0)
	edgeMap := make([][]*tree.Node, height)
	for i := 0; i < height; i++ {
		edgeMap[i] = make([]*tree.Node, 2)
	}
	setEdgeMap(head, 0, edgeMap)
	// 打印左边界
	for i := 0; i < height; i++ {
		println(edgeMap[i][0].V)
	}
	// 打印既不是左边界，又不是右边界的叶子节点
	printLeafNotInMap(head, 0, edgeMap)
	// 打印右边界，但不是左边界的节点
	for i := height - 1; i > -1; i-- {
		if edgeMap[i][0] != edgeMap[i][1] {
			println(edgeMap[i][1].V)
		}
	}
}

func getHight(h *tree.Node, l int) int {
	if h == nil {
		return l
	}
	return int(math.Max(float64(getHight(h.Left, l+1)), float64(getHight(h.Right, l+1))))
}

func setEdgeMap(h *tree.Node, l int, edgeMap [][]*tree.Node) {
	if h == nil {
		return
	}
	if edgeMap[l][0] == nil {
		edgeMap[l][0] = h
	}
	edgeMap[l][1] = h
	setEdgeMap(h.Left, l+1, edgeMap)
	setEdgeMap(h.Right, l+1, edgeMap)
}

func printLeafNotInMap(h *tree.Node, l int, m [][]*tree.Node) {
	if h == nil {
		return
	}
	if h.Left == nil && h.Right == nil && h != m[l][0] && h != m[l][1] {
		println(h.V)
	}
	printLeafNotInMap(h.Left, l+1, m)
	printLeafNotInMap(h.Right, l+1, m)
}

/************************************* 以上为标准一*****************************************************/
