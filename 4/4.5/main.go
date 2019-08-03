package main

/*
给定数组arr，返回arr的最长递增子序列。
*/
import (
	"github.com/Devil-MayCry/algorithmPratice/lib/math"
)

func main() {

}

/******************************O(N2)*****************************************/

func getdp1(arr []int) []int {
	dp := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				dp[i] = math.MaxInt(dp[i], dp[j]+1)
			}
		}
	}
	return dp
}

func generateLIS(arr []int, dp []int) []int {
	l := 0
	index := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] > l {
			l = dp[i]
			index = i
		}
	}
	lis := make([]int, l)
	l--
	lis[l] = arr[index]
	for i := index; i >= 0; i-- {
		if arr[i] < arr[index] && dp[i] == dp[index]-1 {
			l--
			lis[l] = arr[i]
			index = i
		}
	}
	return lis
}

func lis1(arr []int) []int {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	dp := getdp1(arr)
	return generateLIS(arr, dp)
}

/******************************O(NlogN)*****************************************/

func getdp2(arr []int) []int {
	dp := make([]int, len(arr))
	ends := make([]int, len(arr))
	ends[0] = arr[0]
	dp[0] = 1
	right := 0
	l := 0
	r := 0
	m := 0
	for i := l; i < len(arr); i++ {
		l = 0
		r = right
		for l <= r {
			m = (l + r) / 2
			if arr[i] > ends[m] {
				l = m + 1
			} else {
				r = m - 1
			}
			right = math.MaxInt(right, l)
			ends[l] = arr[i]
			dp[i] = l + 1
		}
	}
	return dp
}

func lis2(arr []int) []int {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	dp := getdp2(arr)
	return generateLIS(arr, dp)
}
