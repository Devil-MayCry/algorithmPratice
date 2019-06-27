// 仅用递归函数和栈操作逆序一个栈
package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func reversePop(s *stack.Stack) int {
	v := s.Pop()
	if s.Len() == 0 {
		return v.(int)
	}
	l := reversePop(s)
	s.Push(v)
	return l
}

func reverse(s *stack.Stack) {
	if s.Len() == 0 {
		return
	}
	i := reversePop(s)
	reverse(s)
	s.Push(i)
}

func main() {
	// 准备数据
	data := stack.New()
	data.Push(1)
	data.Push(2)
	data.Push(3)
	data.Push(4)
	data.Push(5)

	reverse(data)
	fmt.Println(data.Pop())
	fmt.Println(data.Pop())
	fmt.Println(data.Pop())
	fmt.Println(data.Pop())
	fmt.Println(data.Pop())
}
