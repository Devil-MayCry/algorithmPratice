/*
给定一个有序数组sortArr，已知其中没有重复值，
用这个有序数组生成一棵平衡搜索二叉树，
并且该搜索二叉树中序遍历的结果与sortArr一致。
*/

package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/tree"
)

func main() {
	arr := []int64{1, 2, 3, 4, 5}
	makeBST(arr)
}

// 判断是否是搜索二叉树
func makeBST(arr []int64) *tree.Node {
	if arr == nil {
		return nil
	}
	return generate(arr, 0, len(arr)-1)
}

func generate(sortArr []int64, start int, end int) *tree.Node {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	head := &tree.Node{V: sortArr[mid]}
	head.Left = generate(sortArr, start, mid-1)
	head.Right = generate(sortArr, mid+1, end)
	return head
}
