package main

import (
	"fmt"

	"github.com/Devil-MayCry/algorithmPratice/lib/linkedlist"
)

/*
据说著名犹太历史学家Josephus有过以下故事：
在罗马人占领乔塔帕特后，39个犹太人与Josephus及他的朋友躲到一个洞中，39个犹太人决定宁愿死也不要被敌人抓到，
于是决定了一个自杀方式，41个人排成个圆圈，由第个人开始报数，报数到3的人就自杀，然后再由下一个人重新报l，
报数到3的人再自杀，这样依次下去，直到剩下最后一个人时，那个人可以自由选择自己的命运。
这就是著名的约瑟夫问题。现在请用单向环形链表 描述该结构并呈现整个自杀过程。
*/
func main() {
	l := linkedlist.MakeCircleChain(buildArr(43))
	printJosephCircle(l)
}

func buildArr(m int) []int64 {
	arr := make([]int64, 0)
	for i := 1; i <= m; i++ {
		arr = append(arr, int64(i))
	}
	return arr
}
func printJosephCircle(head *linkedlist.Node) {
	var pre *linkedlist.Node
	count := 1
	for head.Next != head {
		if count == 3 {
			fmt.Printf("kill %d\n", head.V)
			pre.Next = head.Next
			count = 0
		}
		pre = head
		head = head.Next
		count++
	}
}
