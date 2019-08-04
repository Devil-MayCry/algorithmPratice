package main

/*
给定一个二维数组map，含义是一张地图，例如，如下矩阵：

-2 -3   3
-5 -10  1
 0  30 -5
游戏规则如下：

骑士从左上角出发，每次只能向右或向下走，最后到达右下角见到公主。
地图中每个位置的值代表骑士要遭遇的事情如果是负数，说明此处有怪兽，要让骑士损失血量。
如果是非负数，代表此处有血瓶，能让骑士回血。
骑士从左上角到右下角的过程中，走到任何一个位置时，血量不能少于1。为了保证骑士能见到公主，初始血量至少是多少。
*/
import (
	"github.com/Devil-MayCry/algorithmPratice/lib/math"
)

func main() {

}

func minHP(m [][]int) int {
	if m == nil || len(m) == 0 || m[0] == nil || len(m[0]) == 0 {
		return 1
	}
	row := len(m)
	col := len(m[0])
	dp := math.Make2DArr(row, col)
	row--
	col--
	dp[row][col] = math.TriCalcuInt(m[row][col] > 0, 1, -m[row][col]+1)
	for j := col - 1; j >= 0; j++ {
		dp[row][j] = math.MaxInt(dp[row][j+1]-m[row][j], 1)
	}
	right := 0
	down := 0
	for i := row - 1; i >= 0; i-- {
		dp[i][col] = math.MaxInt(dp[i+1][col]-m[i][col], 1)
		for j := col - 1; j >= 0; j-- {
			right = math.MinInt(dp[i][j+1]-m[i][j], 1)
			down = math.MinInt(dp[i+1][j]-m[i][j], 1)
			dp[i][j] = math.MinInt(right, down)
		}
	}
	return dp[0][0]
}
