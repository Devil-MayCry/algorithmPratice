package main

/*
给定数组arr，返回arr的最长递增子序列。
*/
import (
	"github.com/Devil-MayCry/algorithmPratice/lib/math"
)

func main() {

}

func getdp(str1 []rune, str2 []rune) [][]int {
	dp := math.Make2DArr(len(str1), len(str2))
	dp[0][0] = math.TriCalcuInt(str1[0] == str2[0], 1, 0)
	for i := 1; i < len(str1); i++ {
		dp[i][0] = math.MaxInt(dp[i-1][0], math.TriCalcuInt(str1[i] == str2[0], 1, 0))
	}
	for i := 1; i < len(str2); i++ {
		dp[0][i] = math.MaxInt(dp[0][i-1], math.TriCalcuInt(str1[0] == str2[i], 1, 0))
	}
	for i := 1; i < len(str1); i++ {
		for j := 1; j < len(str2); j++ {
			dp[i][j] = math.MaxInt(dp[i-1][j], dp[i][j-1])
			if str1[i] == str2[j] {
				dp[i][j] = math.MaxInt(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}
	return dp
}

func lcse(str1 string, str2 string) string {
	if str1 == "" || str2 == "" {
		return ""
	}
	chs1 := []rune(str1)
	chs2 := []rune(str2)
	dp := getdp(chs1, chs2)
	m := len(chs1) - 1
	n := len(chs2) - 1
	res := make([]rune, dp[m][n])
	index := len(res) - 1
	for index >= 0 {
		if n > 0 && dp[m][n] == dp[m][n-1] {
			n--
		} else if m > 0 && dp[m][n] == dp[m-1][n] {
			m--
		} else {
			res[index] = chs1[m]
			index--
			m--
			n--
		}
	}
	return string(res)
}
