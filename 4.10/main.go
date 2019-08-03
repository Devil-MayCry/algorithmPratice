package main

/*
给定三个字符串strl、str2和aim，
如果aim包含且仅包含来自strl和str2的所有字符，而且在aim中属于strl的字符之间保持原来在strl中的顺序，
属于str2的字符之间保持原来在str2中的顺序，那么称aim是str1和str2的交错组成。
实现一个函数，判断aim是否是strl和str2交错组成。
*/

func main() {

}

func isCross1(str1 string, str2 string, aim string) bool {
	ch1 := []rune(str1)
	ch2 := []rune(str2)
	chaim := []rune(aim)
	if len(chaim) != len(ch1)+len(ch2) {
		return false
	}
	dp := make([][]bool, len(ch1))
	for i := range dp {
		t := make([]bool, len(ch2))
		dp[i] = t
	}

	dp[0][0] = true
	for i := 1; i <= len(ch1); i++ {
		if ch1[i-1] != chaim[i-1] {
			break
		}
		dp[i][0] = true
	}
	for j := 1; j <= len(ch2); j++ {
		if ch2[j-1] != chaim[j-1] {
			break
		}
		dp[0][j] = true
	}
	for i := 1; i <= len(ch1); i++ {
		for j := 1; j <= len(ch2); i++ {
			if (ch1[i-1] == chaim[i+j-1] && dp[i-1][j]) ||
				(ch2[j-1] == chaim[i+j-1] && dp[i][j-1]) {
				dp[i][j] = true
			}
		}
	}
	return dp[len(ch1)][len(ch2)]
}
