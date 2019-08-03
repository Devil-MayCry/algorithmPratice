package main

import "github.com/Devil-MayCry/algorithmPratice/lib/math"

func main() {

}

/*
给定整数N，返回斐波那契数列第N项
*/
func f(n int) int {
	if n < 1 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}
	res := 1
	pre := 1
	tmp := 1
	for i := 3; i <= n; i++ {
		tmp = res
		res = res + pre
		pre = tmp
	}
	return res
}

/****************************************使用矩阵乘法求斐波拉契**************************************************/

func f2(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	base := [][]int{{1, 1}, {1, 0}}
	res := math.MatrixPower(base, n-2)
	return res[0][0] + res[1][0]
}
