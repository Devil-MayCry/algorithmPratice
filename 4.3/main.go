package main

import (
	maths "math"

	"github.com/Devil-MayCry/algorithmPratice/lib/math"
)

/*
给定数组arr, arr中所有的值都为正数且不重复。每个值代表一种面值的货币，
每种面值的货币可以使用任意张，再给定一个整数aim代表要找的钱数，求组成aim的最少货币数。
*/

func main() {

}

func minCoins1(arr []int, aim int) int {
	if arr == nil || len(arr) == 0 || aim < 0 {
		return -1
	}
	n := len(arr)
	max := maths.MaxInt64
	dp := math.Make2DArr(n, aim+1)
	for j := 1; j <= aim; j++ {
		dp[0][j] = max
		if j-arr[0] >= 0 && dp[0][j-arr[0]] != max {
			dp[0][j] = dp[0][j-arr[0]] + 1
		}
	}
	left := 0
	for i := 1; i < n; i++ {
		for j := 1; j <= aim; j++ {
			left = max
			if j-arr[i] >= 0 && dp[i][j-arr[i]] != max {
				left = dp[i][j-arr[i]] + 1
			}
			dp[i][j] = math.MinInt(left, dp[i-1][j])
		}
	}
	if dp[n-1][aim] != max {
		return dp[n-1][aim]
	}
	return -1
}
