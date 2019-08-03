package main

import "github.com/Devil-MayCry/algorithmPratice/lib/math"

/*
给定一个矩阵m，从左上角开始每次只能向右或者向下走，最后到达右下角的位置，
路径上所有的数字累加起来就是路径和，返回所有的路径中最小的路径和。
*/

func main() {

}

// 普通方法
func minPathSum1(m [][]int) int {
	if m == nil || len(m) == 0 || m[0] == nil || len(m[0]) == 0 {
		return 0
	}
	row := len(m)
	col := len(m[0])
	dp := math.Make2DArr(row, col)
	dp[0][0] = m[0][0]
	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][0] + m[i][0]
	}
	for j := 1; j < col; j++ {
		dp[0][j] = dp[0][j-1] + m[0][j]
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = math.MinInt(dp[i-1][j], dp[i][j-1]) + m[i][j]
		}
	}
	return dp[row-1][col-1]
}

// 压缩矩阵

// 普通方法
func minPathSum2(m [][]int) int {
	if m == nil || len(m) == 0 || m[0] == nil || len(m[0]) == 0 {
		return 0
	}
	row := len(m)
	col := len(m[0])
	more := math.MaxInt(row, col)
	less := math.MinInt(row, col)
	rowmore := more == row
	arr := make([]int, less)
	arr[0] = m[0][0]

	for i := 1; i < less; i++ {
		if rowmore {
			arr[i] = arr[i-1] + m[0][i]
		} else {
			arr[i] = arr[i-1] + m[i][0]
		}
	}
	for i := 1; i < more; i++ {
		if rowmore {
			arr[0] = arr[0] + m[i][0]
		} else {
			arr[0] = arr[0] + m[0][i]
		}
		for j := 1; j < less; j++ {
			if rowmore {
				arr[j] = math.MinInt(arr[j-1], arr[j]) + m[i][j]
			} else {
				arr[j] = math.MinInt(arr[j-1], arr[j]) + m[j][i]
			}
		}
	}

	dp := math.Make2DArr(row, col)
	dp[0][0] = m[0][0]
	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][0] + m[i][0]
	}
	for j := 1; j < col; j++ {
		dp[0][j] = dp[0][j-1] + m[0][j]
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = math.MinInt(dp[i-1][j], dp[i][j-1]) + m[i][j]
		}
	}
	return arr[less-1]
}
