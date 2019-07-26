/*
给定彼此独立的两棵树头节点分别为t1和t2，判断tl中是否有与t2树拓扑结构完全相同的子树。
*/

package main

import (
	"fmt"

	"github.com/Devil-MayCry/algorithmPratice/lib/tree"
)

func main() {
	t1 := tree.MakeExampleTree()
	t2 := tree.MakeExampleTree()
	isSubtree(t1, t2)
}

/*
O(N*M)的解法同3.11类似，不再重复
最优解是O(N*M)： 将两个二叉树进行序列化，然后通过KMP算法判断s2是否为s1子串
*/

func isSubtree(t1 *tree.Node, t2 *tree.Node) bool {
	t1Str := serialByPre(t1)
	t2Str := serialByPre(t2)
	return getIndexOf(t1Str, t2Str) != -1
}

func serialByPre(head *tree.Node) string {
	if head == nil {
		return "#!"
	}
	res := fmt.Sprintf("%d!", head.V)
	res += serialByPre(head.Left)
	res += serialByPre(head.Right)
	return res
}

// KMP
func getIndexOf(s string, m string) int {
	if s == "" || m == "" || len(m) < 1 || len(s) < len(m) {
		return -1
	}
	ss := []rune(s)
	ms := []rune(m)
	si := 0
	mi := 0
	next := getNextArray(ms)
	for si < len(ss) && mi < len(ms) {
		if ss[si] == ms[mi] {
			si++
			mi++
		} else if next[mi] == -1 {
			si++
		} else {
			mi = next[mi]
		}
	}
	if mi == len(ms) {
		return si - mi
	}
	return -1
}

func getNextArray(ms []rune) []int {
	if len(ms) == 1 {
		return []int{-1}
	}
	next := make([]int, len(ms))
	next[0] = -1
	next[1] = 0
	pos := 2
	cn := 0
	for pos < len(next) {
		if ms[pos-1] == ms[cn] {
			cn++
			next[pos] = cn
			pos++
		} else if cn > 0 {
			cn = next[cn]
		} else {
			next[pos] = 0
			pos++
		}
	}
	return next
}
