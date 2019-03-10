/*
定义二叉树节点如下：

public class Node{
    public int value;
    public Node left;
    public Node right;
    public Node(int data){
        this.value = data;
    }
}
一个数组的MaxTree定义如下：

数组必须没有重复元素
MaxTree是一棵二叉树，数组的每一个值对应一个二叉树节点。
包括MaxTree树在内且在其中的每一棵子树上，值最大的节点都是树的头。

给定一个没有重复元素的数组arr，写出生成这个数组的MaxTree的函数，
要求如果数组长度为N,则时间复杂度为O(N)、额外空间复杂度为O(N)。
*/

/*
按照以下原则生成树：
1. 每个数的父节点是左边第一个比它大的，和右边第一个比它大的，较小的那个
2. 如果左右都没有比它大的，即为头结点

该原则的证明过程见书
*/
package main

import (
	"fmt"
	"math"

	"github.com/golang-collections/collections/stack"
)

type node struct {
	left  *node
	right *node
	value int
}

func findLeftHigher(arr []int) (res []int) {
	h := stack.New()
	for _, v := range arr {
		for {
			if h.Len() == 0 {
				res = append(res, math.MaxInt64)
				h.Push(v)
				break
			}
			if h.Peek().(int) > v {
				res = append(res, h.Peek().(int))
				h.Push(v)
				break
			}
			h.Pop()
		}
	}
	return res
}

func findRightHigher(arr []int) (res []int) {
	h := stack.New()
	for i := len(arr) - 1; i >= 0; i-- {
		for {
			if h.Len() == 0 {
				res = append(res, math.MaxInt64)
				h.Push(arr[i])
				break
			}
			if h.Peek().(int) > arr[i] {
				res = append(res, h.Peek().(int))
				h.Push(arr[i])
				break
			}
			h.Pop()
		}
	}
	return res
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	arr := []int{3, 6, 5, 7, 2, 1}
	lres := findLeftHigher(arr)
	rres := findRightHigher(arr)

	m := make(map[int]*node)

	// 创建index和value的对应关系表
	tmp := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		tmp[arr[i]] = i

		n := &node{}
		n.value = arr[i]
		m[i] = n
	}

	// 构建树
	for i := 0; i < len(arr); i++ {
		father := min(lres[i], rres[len(arr)-1-i])
		if m[tmp[father]].left == nil {
			m[tmp[father]].left = m[i]
		} else {
			m[tmp[father]].right = m[i]
		}
	}

	for _, v := range m {
		fmt.Printf("%+v", v)
	}
}
