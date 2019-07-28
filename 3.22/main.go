/*
给定一棵完全二叉树的头节点head，返回这棵树的节点个数。
要求：时间复杂度低于O(N)
*/

package main

func main() {
	numTrees(9)
}

func numTrees(n int) int {
	if n < 2 {
		return 1
	}
	num := make([]int, n+1)
	num[0] = 1
	for i := 1; i < n+1; i++ {
		for j := 1; j < i+1; j++ {
			num[i] += num[j-1] + num[i-j]
		}
	}
	return num[n]
}
