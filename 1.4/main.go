// 用一个栈排序另一个栈
package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func sort(s *stack.Stack) {
	h := stack.New()
	for {
		if s.Len() == 0 {
			break
		}
		if h.Len() == 0 {
			h.Push(s.Pop())
		}
		sp := s.Pop()
		for {
			if sp.(int) <= h.Peek().(int) {
				break
			}
			s.Push(h.Pop())
		}
		h.Push(sp)
	}

	for {
		if h.Len() == 0 {
			break
		}
		s.Push(h.Pop())
	}
}

func main() {
	// 准备数据
	data := stack.New()
	data.Push(2)
	data.Push(4)
	data.Push(3)
	data.Push(1)
	data.Push(5)
	sort(data)
	fmt.Println(data.Pop())
	fmt.Println(data.Pop())
	fmt.Println(data.Pop())
	fmt.Println(data.Pop())
	fmt.Println(data.Pop())

}
