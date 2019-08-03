package main

/*
给定两个字符串strl和str2，再给定三个整数ic、dc和rc，分别代表插入、删除和替换一个字符的代价，
返回将strl编辑成str2的最小代价。
*/
import (
	"github.com/Devil-MayCry/algorithmPratice/lib/math"
)

func main() {

}

func minCost1(str1 string, str2 string, ic int, dc int, rc int) int {
	if str1 == "" || str2 == "" {
		return 0
	}
	chs1 := []rune(str1)
	chs2 := []rune(str2)
	row := len(chs1) + 1
	col := len(chs2) + 1
	dp := math.Make2DArr(row, col)
	for i := 1; i < row; i++ {
		dp[i][0] = dc * i
	}
	for j := 1; j < col; j++ {
		dp[0][j] = ic * j
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if chs1[i-1] == chs2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j-1] + rc
			}
			dp[i][j] = math.MinInt(dp[i][j], dp[i][j-1]+ic)
			dp[i][j] = math.MinInt(dp[i][j], dp[i-1][j]+dc)
		}
	}
	return dp[row-1][col-1]
}
