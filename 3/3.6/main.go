/*
给定一棵二叉树的头节点head和一个32位整数sum，二叉树节点值类型为整型，
求累加和为sum的最长路径长度。
路径是指从某个节点往下，每次最多选择一个孩子节点或者不选所形成的节点链。
*/

package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/math"
	"github.com/Devil-MayCry/algorithmPratice/lib/tree"
)

func main() {
	t := tree.MakeExampleTree()
	res := getMaxLength(t, 4)
	println(res)
}

func getMaxLength(head *tree.Node, sum int64) int64 {
	sumMap := make(map[int64]int64)
	sumMap[0] = 0
	return preOrder(head, sum, 0, 1, 0, sumMap)
}

func preOrder(head *tree.Node, sum int64, preSum int64, level int64, maxLen int64, sumMap map[int64]int64) int64 {
	if head == nil {
		return maxLen
	}
	curSum := preSum + head.V
	_, ok := sumMap[curSum]
	if !ok {
		sumMap[curSum] = level
	}
	_, ok = sumMap[curSum-sum]
	if ok {
		maxLen = math.Max(level-sumMap[curSum-sum], maxLen)
	}
	maxLen = preOrder(head.Left, sum, curSum, level+1, maxLen, sumMap)
	maxLen = preOrder(head.Right, sum, curSum, level+1, maxLen, sumMap)
	if level == sumMap[curSum] {
		delete(sumMap, curSum)
	}
	return maxLen
}
