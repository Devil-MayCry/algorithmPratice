// 实现一个支持获取栈最小值的栈
// 易错点： 忽略栈为空的情况

package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

type getminStack struct {
	s  stack.Stack
	ms stack.Stack
}

func (g *getminStack) Pop() int {
	if g.s.Len() == 0 {
		panic("empty")
	}
	g.ms.Pop()
	return g.s.Pop().(int)
}

func (g *getminStack) Push(v int) {
	g.s.Push(v)
	if g.ms.Len() == 0 || v <= g.ms.Peek().(int) {
		g.ms.Push(v)
	} else {
		g.ms.Push(g.ms.Peek().(int))
	}
}

func (g *getminStack) GetMin() int {
	if g.s.Len() == 0 {
		panic("empty")
	}
	return g.ms.Peek().(int)
}

func main() {
	t := &getminStack{}
	t.Push(3)
	fmt.Println(t.GetMin())
	t.Push(2)
	fmt.Println(t.GetMin())
	t.Push(1)
	fmt.Println(t.GetMin())
	t.Pop()
	fmt.Println(t.GetMin())
	t.Pop()
	fmt.Println(t.GetMin())
}
