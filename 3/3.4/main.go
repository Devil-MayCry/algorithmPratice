/*
二叉树被记录成文件的过程叫作二叉树的序列化，通过文件内容重建原来二叉树的过程叫作二叉树的反序列化。
给定一棵二叉树的头节点 head，并己知二叉树节点值的类型为32位整型。
请设计一种二叉树序列化和反序列化的方案，并用代码实现。
*/
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Devil-MayCry/algorithmPratice/lib/tree"
	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
)

var s *stack.Stack

func main() {
	t := tree.MakeExampleTree()
	s := serialByPre(t)
	println(s)
	reconByString(s).PrintlnTree()
}

func serialByPre(h *tree.Node) string {
	if h == nil {
		return "#!"
	}
	s := fmt.Sprintf("%d!", h.V)
	s = s + serialByPre(h.Left)
	s = s + serialByPre(h.Right)
	return s
}

func reconByString(s string) *tree.Node {
	values := strings.Split(s, "!")
	q := queue.New()
	for i := 0; i < len(values); i++ {
		q.Enqueue(values[i])
	}
	return reconOrder(q)
}

func reconOrder(q *queue.Queue) *tree.Node {
	value := q.Dequeue()
	if value.(string) == "#" {
		return nil
	}
	v, _ := strconv.ParseInt(value.(string), 10, 64)
	head := &tree.Node{
		V: v,
	}
	head.Left = reconOrder(q)
	head.Right = reconOrder(q)
	return head
}
