package main

/*
给定一个整数n，代表汉诺塔游戏中从小到大放置的n个圆盘，
假设开始时所有的圆盘都放在左边的柱子上，想按照汉诺塔游戏的要求把所有的圆盘都移到右边的柱子上。
实现函数打印最优移动轨迹。
*/

func main() {

}

func hanoi(n int) {
	if n > 0 {
		f(n, "left", "mid", "right")
	}

}

func f(n int, from string, mid string, to string) {
	if n == 1 {
		println("move from " + from + " to " + to)
	} else {
		f(n-1, from, to, mid)
		f(1, from, mid, to)
		f(n-1, mid, from, to)
	}
}
