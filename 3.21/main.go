/*
己知一棵二叉树的所有节点值都不同，给定这棵二叉树正确的先序、中序和后序数组。
请分别用三个函数实现任意两种数组结合重构原来的二叉树，并返回重构二叉树的头节点。
*/

package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/tree"
)

func main() {
	arrPre := []int64{1, 2, 4, 5, 8, 9, 3, 6, 7}
	arrIn := []int64{4, 2, 8, 5, 9, 1, 6, 3, 7}
	// arrPost := []int64{4, 8, 9, 5, 2, 6, 7, 3, 1}
	preInToTree(arrPre, arrIn)
}

func preInToTree(arrPre []int64, arrIn []int64) *tree.Node {
	if arrPre == nil || arrIn == nil {
		return nil
	}
	m := make(map[int64]int64)
	for i := 0; i < len(arrIn); i++ {
		m[arrIn[i]] = int64(i)
	}
	return preIn(arrPre, 0, int64(len(arrPre)-1), arrIn, 0, int64(len(arrIn)-1), m)
}

func preIn(p []int64, pi int64, pj int64, n []int64, ni int64, nj int64, m map[int64]int64) *tree.Node {
	if pi > pj {
		return nil
	}
	head := &tree.Node{V: p[pi]}
	index := m[p[pi]]
	head.Left = preIn(p, pi+1, pi+index-ni, n, ni, index-1, m)
	head.Right = preIn(p, pi+index-ni+1, pj, n, index+1, nj, m)
	return head
}

/*******************************中序+后序*****************************************************************/

func inPosToTree(arrIn []int64, arrPost []int64) *tree.Node {
	if arrIn == nil || arrPost == nil {
		return nil
	}
	m := make(map[int64]int64)
	for i := 0; i < len(arrIn); i++ {
		m[arrIn[i]] = int64(i)
	}
	return inPos(arrIn, 0, int64(len(arrIn)-1), arrPost, 0, int64(len(arrPost)-1), m)
}

func inPos(n []int64, ni int64, nj int64, s []int64, si int64, sj int64, m map[int64]int64) *tree.Node {
	if si > sj {
		return nil
	}
	head := &tree.Node{V: s[si]}
	index := m[s[si]]
	head.Left = inPos(n, ni, index-1, s, si, si+index-ni-1, m)
	head.Right = inPos(n, index+1, nj, s, si+index-ni, sj-1, m)
	return head
}

/*******************************前序+后序*****************************************************************/

// 如果一颗二叉树除叶子节点外，所有的其他节点都有左孩子和右孩子，这样的树才能被先序和后序构建出来

func prePosToTree(arrPre []int64, arrPost []int64) *tree.Node {
	if arrPre == nil || arrPost == nil {
		return nil
	}
	m := make(map[int64]int64)
	for i := 0; i < len(arrPost); i++ {
		m[arrPost[i]] = int64(i)
	}
	return prePos(arrPre, 0, int64(len(arrPre)-1), arrPost, 0, int64(len(arrPost)-1), m)
}

func prePos(p []int64, pi int64, pj int64, s []int64, si int64, sj int64, m map[int64]int64) *tree.Node {
	head := &tree.Node{V: s[sj]}
	sj--
	if pi == pj {
		return head
	}
	pi++
	index := m[p[pi]]
	head.Left = prePos(p, pi, pi+index-si, s, si, index, m)
	head.Right = prePos(p, pi+index-si+1, pj, s, index+1, sj, m)
	return head
}
