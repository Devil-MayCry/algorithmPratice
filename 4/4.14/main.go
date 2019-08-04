package main

/*
给定一个整型数组arr，代表数值不同的纸牌排成一条线。
玩家A和玩家B依次拿走每张纸牌，规定玩家A先拿，玩家B后拿，
但是每个玩家每次只能拿走最左或最右的纸牌，玩家A和玩家B都绝顶聪明。请返回最后获胜者的分数。
*/
import (
	"github.com/Devil-MayCry/algorithmPratice/lib/math"
)

func main() {

}

func win1(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	return math.MaxInt(f(arr, 0, len(arr)-1), s(arr, 0, len(arr)-1))
}

func f(arr []int, i int, j int) int {
	if i == j {
		return arr[i]
	}
	return math.MaxInt(arr[i]+s(arr, i+1, j), arr[j]+s(arr, i, j-1))

}

func s(arr []int, i int, j int) int {
	if i == j {
		return 0
	}
	return math.MinInt(f(arr, i+1, j), f(arr, i, j-1))
}
