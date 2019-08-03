package main

/*
给定数组arr,arr中所有的值都为正数且不重复。每个值代表一种面值的货币，每种面值的货币可以使用任意张，
再给定一个整数aim代表要找的钱数，求换钱有多少种方法。
*/
import (
	"github.com/Devil-MayCry/algorithmPratice/lib/math"
)

func main() {

}

/******************************暴力破解*****************************************/
func coins(arr []int, aim int) int {
	if arr == nil || len(arr) == 0 || aim < 0 {
		return 0
	}
	return process1(arr, 0, aim)
}

func process1(arr []int, index int, aim int) int {
	res := 0
	if index == len(arr) {
		if aim == 0 {
			res = 1
		} else {
			res = 0
		}
	} else {
		for i := 0; arr[index]*i <= aim; i++ {
			res += process1(arr, index+1, aim-arr[index]*i)
		}
	}
	return res
}

/******************************暴力破解+记忆化搜索*****************************************/

func coins2(arr []int, aim int) int {
	if arr == nil || len(arr) == 0 || aim < 0 {
		return 0
	}
	m := math.Make2DArr(len(arr)+1, aim+1)
	return process2(arr, 0, aim, m)
}

func process2(arr []int, index int, aim int, m [][]int) int {
	res := 0
	if index == len(arr) {
		if aim == 0 {
			res = 1
		} else {
			res = 0
		}
	} else {
		mapValue := 0
		for i := 0; arr[index]*i <= aim; i++ {
			mapValue = m[index+1][aim-arr[index]*i]
			if mapValue != 0 {
				if mapValue == -1 {
					res += 0
				} else {
					res += mapValue
				}
			} else {
				res += process2(arr, index+1, aim-arr[index]*i, m)
			}
		}
	}
	if res == 0 {
		m[index][aim] = -1
	} else {
		m[index][aim] = res
	}
	return res
}

/******************************动态规划*****************************************/

func coin3(arr []int, aim int) int {
	if arr == nil || len(arr) == 0 || aim < 0 {
		return 0
	}
	dp := math.Make2DArr(len(arr), aim+1)
	for i := 0; i < len(arr); i++ {
		dp[i][0] = 1
	}
	for j := 1; arr[0]*j <= aim; j++ {
		dp[0][arr[0]*j] = 1
	}
	num := 0
	for i := 0; i < len(arr); i++ {
		for j := 1; j <= aim; j++ {
			num = 0
			for k := 0; j-arr[i]*k >= 0; k++ {
				num += dp[i-1][j-arr[i]*k]
			}
			dp[i][j] = num
		}
	}
	return dp[len(arr)-1][aim]
}

/******************************优化动态规划*****************************************/

func coin4(arr []int, aim int) int {
	if arr == nil || len(arr) == 0 || aim < 0 {
		return 0
	}
	dp := math.Make2DArr(len(arr), aim+1)
	for i := 0; i < len(arr); i++ {
		dp[i][0] = 1
	}
	for j := 1; arr[0]*j <= aim; j++ {
		dp[0][arr[0]*j] = 1
	}
	for i := 0; i < len(arr); i++ {
		for j := 1; j <= aim; j++ {
			dp[i][j] = dp[i-1][j]
			if j-arr[i] > 0 {
				dp[i][j] += dp[i][j-arr[i]]
			}
		}
	}
	return dp[len(arr)-1][aim]
}
