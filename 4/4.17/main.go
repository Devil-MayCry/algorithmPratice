package main

import "math"

/*
N皇后问题是指在 NxN 的棋盘上要摆N个皇后，
 要求任何两个皇后不同行、不同列，也不在同一条斜线上。
 给定一个整数n,返回n皇后的摆法有多少种。
*/

func main() {

}

func num1(n int) int {
	if n < 1 {
		return 0
	}
	record := make([]int, n)
	return process1(0, record, n)
}

func process1(i int, record []int, n int) int {
	if i == n {
		return 1
	}
	res := 0
	for j := 0; j < n; j++ {
		if isValid(record, i, j) {
			record[i] = j
			res += process1(i+1, record, n)
		}
	}
	return res
}

func isValid(record []int, i int, j int) bool {
	for k := 0; k < i; k++ {
		if j == record[k] || math.Abs(float64(record[k]-j)) == math.Abs(float64(i-k)) {
			return false
		}
	}
	return true
}
