/*
一棵二叉树原本是搜索二叉树，但是其中有两个节点调换了位置，使得这棵二叉树不再是搜索二叉树，
请找到这两个错误节点并返回。
己知二叉树中所有节点的值都不一样，给定二叉树的头节点head,返回一个长度为2的二叉树节点类型的数组errs,
errs[0]表示一个错误节点，errs[1]表示另一个错误节点。

 进阶：
 如果原问题中得到了这两个错误节点，
 我们当然可以通过交换两个节点的节点值的方式让整棵二叉树重新成为搜索二叉树。
但现在要求你不能这么做，而是在结构上完全交换两个节点的位置，请实现调整的函数。
*/

package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/tree"
	"github.com/golang-collections/collections/stack"
)

func main() {
	t := tree.MakeExampleTree()
	getTwoErrNodes(t)
}

func getTwoErrNodes(head *tree.Node) []*tree.Node {
	errs := make([]*tree.Node, 2)
	if head == nil {
		return errs
	}
	s := stack.New()
	var pre *tree.Node
	for s.Len() > 0 || head != nil {
		if head != nil {
			s.Push(head)
			head = head.Left
		} else {
			head = s.Pop().(*tree.Node)
			if pre != nil && pre.V > head.V {
				if errs[0] == nil {
					errs[0] = pre
				}
				errs[1] = head
			}
			pre = head
			head = head.Right
		}
	}
	return errs
}
