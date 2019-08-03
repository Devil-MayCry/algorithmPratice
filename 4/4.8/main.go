package main

/*
给定两个字符串strl和str2，返回两个字符串的最长公共子串。
*/
import (
	"github.com/Devil-MayCry/algorithmPratice/lib/math"
)

func main() {

}

func getdp(str1 []rune, str2 []rune) [][]int {
	dp := math.Make2DArr(len(str1), len(str2))
	for i := 1; i < len(str1); i++ {
		if str1[i] == str2[0] {
			dp[i][0] = 1
		}
	}
	for i := 1; i < len(str2); i++ {
		if str1[0] == str2[i] {
			dp[0][i] = 1
		}
	}
	for i := 1; i < len(str1); i++ {
		for j := 1; j < len(str2); j++ {
			if str1[i] == str2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}
	return dp
}

func lcst1(str1 string, str2 string) string {
	if str1 == "" || str2 == "" {
		return ""
	}
	chs1 := []rune(str1)
	chs2 := []rune(str2)
	dp := getdp(chs1, chs2)
	end := 0
	max := 0
	for i := 0; i < len(chs1); i++ {
		for j := 0; j < len(chs2); j++ {
			if dp[i][j] > max {
				end = i
				max = dp[i][j]
			}
		}
	}
	return string(chs1[end-max+1 : end+1])
}

func lcst2(str1 string, str2 string) string {
	if str1 == "" || str2 == "" {
		return ""
	}
	chs1 := []rune(str1)
	chs2 := []rune(str2)
	row := 0
	col := len(chs2) - 1
	end := 0
	max := 0
	for row < len(chs1) {
		i := row
		j := col
		l := 0
		for i < len(chs1) && j < len(chs2) {
			if chs1[i] != chs2[j] {
				l = 0
			} else {
				l++
			}
			if l > max {
				end = i
				max = l
			}
			i++
			j++
		}
		if col > 0 {
			col--
		} else {
			row++
		}
	}
	return string(chs1[end-max+1 : end+1])
}
