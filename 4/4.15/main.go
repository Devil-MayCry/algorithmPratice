package main

import (
	"github.com/Devil-MayCry/algorithmPratice/lib/math"
)

/*
给定数组arr, arr[i]=k代表可以从位置i向右跳l~k个距离。
比如，arr[2]==3，代表从位置2可以跳到位置3、位置4或位置5。
如果从位置0出发，返回最少跳几次能跳到 arr最后的位置上。
*/

func main() {

}

func jump(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	jump := 0
	cur := 0
	next := 0
	for i := 0; i < len(arr); i++ {
		if cur < i {
			jump++
			cur = next
		}
		next = math.MaxInt(next, i+arr[i])
	}
	return jump
}
