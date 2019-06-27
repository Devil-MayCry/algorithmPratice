/*
汉诺塔问题比较经典，这里修改一下游戏规则：现在限制不能从最左侧的塔直接移动到最右侧，
也不能从最右侧直接移动到最左侧，而是必须经过中间。
求当塔有N层的时候，打印最优移动过程和最优移动总步数。
例如，当塔数为两层时，最上层的塔记为1，最下层的塔记为2，则打印：
Move 1 from left to mid
Move 1 from mid to right
Move 2 from left to mid
Move 1 from right to mid
Move 1 from mid to left
Move 2 from mid to right
Move 1 from left to mid
Move 1 from mid to right
It will move 8 steps
*/
package main

import "fmt"

var step int

func moveA2B(n int) {
	if n == 1 {
		fmt.Printf("%d A -> B", n)
		step++
		fmt.Printf("\n")
		return
	}
	moveA2B(n - 1)
	moveB2C(n - 1)
	fmt.Printf("%d A -> B", n)
	step++

	fmt.Printf("\n")

	moveC2B(n - 1)
}

func moveB2C(n int) {
	if n == 1 {
		fmt.Printf("%d B -> C", n)
		step++

		fmt.Printf("\n")
		return
	}
	moveB2A(n - 1)
	fmt.Printf("%d B -> C", n)
	step++

	fmt.Printf("\n")

	moveA2B(n - 1)
	moveB2C(n - 1)
}

func moveC2B(n int) {
	if n == 1 {
		fmt.Printf("%d C -> B", n)
		step++

		fmt.Printf("\n")

		return
	}
	moveC2B(n - 1)
	moveB2A(n - 1)
	fmt.Printf("%d C -> B", n)
	step++

	fmt.Printf("\n")

	moveA2B(n - 1)
}

func moveB2A(n int) {
	if n == 1 {
		fmt.Printf("%d B -> A", n)
		step++

		fmt.Printf("\n")

		return
	}
	moveB2C(n - 1)
	fmt.Printf("%d B -> A", n)
	step++

	fmt.Printf("\n")

	moveC2B(n - 1)
	moveB2A(n - 1)
}

func main() {
	n := 10
	moveA2B(n - 1)
	moveB2C(n - 1)
	fmt.Printf("%d A -> B", n)
	step++

	fmt.Printf("\n")

	moveC2B(n - 1)
	moveB2A(n - 1)
	fmt.Printf("%d B -> C", n)
	step++

	fmt.Printf("\n")

	moveA2B(n - 1)
	moveB2C(n - 1)
	fmt.Println(step)
}
