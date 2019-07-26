/*
给定一个整型数组arr，己知其中没有重复值，判断arr是否可能是节点值类型为整型的搜索二叉树后序遍历的结果。
进阶：如果整型数组arr中没有重复值，且己知是一棵搜索二叉树的后序遍历结果，通过数组arr重构二叉树。
*/

package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/tree"
)

func main() {
	t := []int64{4, 5, 2, 6, 3, 1}
	isPostArray(t)
	h := posArrayToBST(t)
	h.PrintlnTree()
}

func isPostArray(arr []int64) bool {
	if arr == nil || len(arr) == 0 {
		return false
	}
	return isPost(arr, 0, int64(len(arr)-1))
}

func isPost(arr []int64, start int64, end int64) bool {
	if start == end {
		return true
	}
	less := int64(-1)
	more := end
	for i := start; i < end; i++ {
		if arr[end] > arr[i] {
			less = i
		} else {
			if more == end {
				more = i
			}
		}
	}
	if less == -1 || more == end {
		return isPost(arr, start, end-1)
	}
	if less != more-1 {
		return false
	}
	return isPost(arr, start, less) && isPost(arr, more, end-1)
}

func posArrayToBST(posArr []int64) *tree.Node {
	if posArr == nil {
		return nil
	}
	return posToBST(posArr, 0, int64(len(posArr)-1))
}

func posToBST(posArr []int64, start int64, end int64) *tree.Node {
	if start > end {
		return nil
	}
	head := &tree.Node{V: posArr[end]}
	less := int64(-1)
	more := end
	for i := start; i < end; i++ {
		if posArr[end] > posArr[i] {
			less = i
		} else {
			if more == end {
				more = i
			}
		}
	}
	head.Left = posToBST(posArr, start, less)
	head.Right = posToBST(posArr, more, end-1)
	return head
}
