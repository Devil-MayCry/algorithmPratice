/*
己知一棵二叉树所有的节点值都不同，给定这棵树正确的先序和中序数组，
不要重建整棵树，而是通过这两个数组直接生成正确的后序数组。
*/

package main

func main() {
	arrPre := []int64{1, 2, 4, 5, 8, 9, 3, 6, 7}
	arrIn := []int64{4, 2, 8, 5, 9, 1, 6, 3, 7}
	// arrPost := []int64{4, 8, 9, 5, 2, 6, 7, 3, 1}
	getPostArr(arrPre, arrIn)
}

func getPostArr(arrPre []int64, arrIn []int64) []int64 {
	if arrPre == nil || arrIn == nil {
		return nil
	}
	l := len(arrPre)
	pos := make([]int64, l)
	m := make(map[int64]int64)
	for i := 0; i < l; i++ {
		m[arrIn[i]] = int64(i)
	}
	setPos(arrPre, 0, int64(l-1), arrIn, 0, int64(l-1), pos, int64(l-1), m)
	return pos
}

func setPos(p []int64, pi int64, pj int64, n []int64, ni int64, nj int64, s []int64, si int64, m map[int64]int64) int64 {
	if pi > pj {
		return si
	}
	s[si] = p[pi]
	si--
	i := m[p[pi]]
	si = setPos(p, pi-nj+i+1, pj, n, i+1, nj, s, si, m)
	return setPos(p, pi+1, pi+i-ni, n, ni, i-1, s, si, m)
}
