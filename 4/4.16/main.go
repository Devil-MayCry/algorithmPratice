package main

/*
给定无序数组arr， 返回其中最长的连续序列的长度。
*/
import (
	"github.com/Devil-MayCry/algorithmPratice/lib/math"
)

func main() {

}
func longestConsecutive(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	max := 1
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if _, ok := m[arr[i]]; !ok {
			m[arr[i]] = 1
			if _, ok := m[arr[i]-1]; ok {
				max = math.MaxInt(max, merge(m, arr[i-1], arr[i]))
			}
			if _, ok := m[arr[i]+1]; ok {
				max = math.MaxInt(max, merge(m, arr[i], arr[i]+1))
			}
		}
	}
	return max
}

func merge(m map[int]int, less int, more int) int {
	left := less - m[less] + 1
	right := more - m[more] + 1
	l := right - left + 1
	m[left] = l
	m[right] = l
	return l
}
