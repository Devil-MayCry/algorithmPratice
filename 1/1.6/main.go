/*
有一个整数数组arr和一个大小为w的窗口从数组的最左边滑到最右边，窗口每次向右边滑到一个位置。 例如，数组为[4,3,5,4,3,3,6,7],窗口大小为3时：

[4 3 5] 4 3 3 6 7     窗口中最大值为5
4 [3 5 4] 3 3 6 7     窗口中最大值为5
4 3 [5 4 3] 3 6 7     窗口中最大值为5
4 3 5 [4 3 3] 6 7     窗口中最大值为4
4 3 5 4 [3 3 6] 7     窗口中最大值为6
4 3 5 4 3 [3 6 7]     窗口中最大值为7
如果数组长度为n,窗口大小为w,则一共产生n-w+1个窗口的最大值。 请实现一个函数。

输入：整数数组arr,窗口大小为w。
输出：一个长度为n-w+1的数组res,res[i]表示每一种窗口状态下的最大值。

*/
package main

import "fmt"

type twoWayQueue struct {
	arr []int
}

func newTwoWayQueue() *twoWayQueue {
	return &twoWayQueue{}
}

func (t *twoWayQueue) add(data int) {
	t.arr = append(t.arr, data)
}

func (t *twoWayQueue) peekHead() int {
	if len(t.arr) == 0 {
		panic(-1)
	}
	return t.arr[0]
}

func (t *twoWayQueue) popHead() {
	if len(t.arr) == 0 {
		panic(-1)
	}
	t.arr = t.arr[1:]
}

func (t *twoWayQueue) peekTail() int {
	return t.arr[len(t.arr)-1]
}

func (t *twoWayQueue) popTail() {
	t.arr = t.arr[0 : len(t.arr)-1]
}

func (t *twoWayQueue) len() int {
	return len(t.arr)
}

func main() {
	var array = [9]int{5, 3, 4, 2, 1, 1, 1, 1, 1}
	var w = 5
	var res []int
	helpQueue := newTwoWayQueue()
	for k, v := range array {
		for {
			if helpQueue.len() > 0 && k-helpQueue.peekHead() >= w {
				helpQueue.popHead()
			}
			if helpQueue.len() == 0 || v < array[helpQueue.peekTail()] {
				helpQueue.add(k)
				break
			} else {
				helpQueue.popTail()
			}
		}
		res = append(res, array[helpQueue.peekHead()])
	}
	fmt.Println(res)
}
