/*
给定一个整数矩阵map,其中的值只有0和1两种，求其中全是1的所有矩形区域中，最大的矩形区域为1的数量。

例如：
1 1 1 0
其中，最大的矩形区域有3个1，所以返回3。
再如：
1 0 1 1
1 1 1 1
1 1 1 0
其中，最大的矩形区域有6个1，所以返回6。
*/
package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func findMaxValueInRow(arr [][]int, row int) int {
	ta := calculateHightArr(arr, row)
	s := stack.New()
	maxArea := 0
	for i := 0; i < len(ta); i++ {
		for s.Len() != 0 && ta[s.Peek().(int)] >= ta[i] {
			j := s.Pop().(int)
			var k int
			if s.Len() == 0 {
				k = -1
			} else {
				k = s.Peek().(int)
			}
			curArea := (i - k - 1) * ta[j]
			maxArea = max(maxArea, curArea)
		}
		s.Push(i)
	}
	for s.Len() > 0 {
		j := s.Pop().(int)
		var k int
		if s.Len() == 0 {
			k = -1
		} else {
			k = s.Peek().(int)
		}
		curArea := (len(arr[0]) - k - 1) * ta[j]
		maxArea = max(maxArea, curArea)
	}
	// fmt.Println(maxArea)
	return maxArea
}

// 将二维数组化为显示直方图高度的一维数组
func calculateHightArr(arr [][]int, row int) []int {
	res := make([]int, len(arr[0]))
	for j := 0; j < len(arr[0]); j++ {
		res[j] = 0
		for i := row; i >= 0; i-- {
			if arr[i][j] == 0 {
				break
			}
			res[j]++
		}
	}
	return res
}

func main() {
	arr := [][]int{
		{0, 0, 1, 0, 0, 1},
		{0, 1, 1, 1, 0, 1},
		{1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1},
	}

	res := 0
	for i := 0; i < len(arr); i++ {
		res = max(res, findMaxValueInRow(arr, i))
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
