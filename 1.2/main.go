// 用两个栈来实现队列
package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

type stackQueue struct {
	s1 stack.Stack
	s2 stack.Stack
}

func (g *stackQueue) add(v int) {
	g.s1.Push(v)
}

func (g *stackQueue) poll() int {
	if g.s2.Len() == 0 && g.s1.Len() == 0 {
		panic(-1)
	}
	if g.s2.Len() == 0 {
		for {
			v := g.s1.Pop()
			if v == nil {
				break
			}
			g.s2.Push(v)
		}
	}
	return g.s2.Pop().(int)
}

func (g *stackQueue) Peek() int {
	if g.s2.Len() == 0 && g.s1.Len() == 0 {
		panic(-1)
	}
	if g.s2.Len() == 0 {
		for {
			v := g.s1.Pop()
			if v == nil {
				break
			}
			g.s2.Push(v)
		}
	}
	return g.s2.Peek().(int)
}

func main() {
	s := &stackQueue{}
	s.add(1)
	s.add(2)
	fmt.Println(s.poll())
	fmt.Println(s.poll())
	s.add(3)
	fmt.Println(s.poll())
}
