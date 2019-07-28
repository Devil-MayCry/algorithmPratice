/*
给定一棵二叉树的头节点head,以及这棵树中的两个节点o1和o2，请返回o1和o2的最近公共祖先节点。
*/

package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/tree"
)

func main() {

}

func findNearestParent(head *tree.Node, o1 *tree.Node, o2 *tree.Node) *tree.Node {
	if head == nil || o1 == nil || o2 == nil {
		return head
	}
	left := findNearestParent(head.Left, o1, o2)
	right := findNearestParent(head.Right, o1, o2)
	if left != nil && right != nil {
		return head
	}

	if left != nil {
		return left
	}
	return right
}

/********************************************方法二：通过保存记录*********************************************************************/
var nodeMap map[*tree.Node](map[*tree.Node]*tree.Node)

func findNearestParent2(head *tree.Node, o1 *tree.Node, o2 *tree.Node) *tree.Node {
	nodeMap := make(map[*tree.Node](map[*tree.Node]*tree.Node))
	initMap(head)
	setMap(head)
	if o1 == o2 {
		return o1
	}
	if _, ok := nodeMap[o1]; ok {
		return nodeMap[o1][o2]
	}
	if _, ok := nodeMap[o2]; ok {
		return nodeMap[o2][o1]
	}
	return nil
}

func initMap(head *tree.Node) {
	if head == nil {
		return
	}
	m := make(map[*tree.Node]*tree.Node)
	nodeMap[head] = m
	initMap(head.Left)
	initMap(head.Right)
}

func setMap(head *tree.Node) {
	if head == nil {
		return
	}
	headRecord(head.Left, head)
	headRecord(head.Right, head)
	subRecord(head)
	setMap(head.Left)
	setMap(head.Right)
}

func headRecord(n *tree.Node, h *tree.Node) {
	if n == nil {
		return
	}
	nodeMap[n][h] = h
	headRecord(n.Left, h)
	headRecord(n.Right, h)
}

func subRecord(head *tree.Node) {
	if head == nil {
		return
	}
	preLeft(head.Left, head.Right, head)
	subRecord(head.Left)
	subRecord(head.Right)
}

func preLeft(l *tree.Node, r *tree.Node, h *tree.Node) {
	if l == nil {
		return
	}
	preRight(l, r, h)
	preLeft(l.Left, r, h)
	preLeft(l.Right, r, h)
}

func preRight(l *tree.Node, r *tree.Node, h *tree.Node) {
	if l == nil {
		return
	}
	nodeMap[l][r] = h
	preRight(l, r.Left, h)
	preRight(l, r.Right, h)
}
