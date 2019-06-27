/*
题目：

给定数组arr和整数num,共返回有多少个子数组满足如下情况：
max(arr[i..j])-min(arr[i..j])<=num

max(arr[i..j])表示子数组arr[i..j]中的最大值，
min(arr[i..j])表示子数组arr[i..j]中的最小值。

要求： 如果数组长度为N,请实现时间复杂度为O(N)的解法。
*/

package main

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
	arr := []int{1, 3, 5, 8, 9, 12, 15}
	num := 3
	print(getNum(arr, num))
}

func getNum(arr []int, num int) int {
	if len(arr) == 0 {
		return 0
	}
	maxQueue := newTwoWayQueue()
	minQueue := newTwoWayQueue()
	i := 0
	j := 0
	res := 0
	for i < len(arr) {
		for j < len(arr) {
			for minQueue.len() != 0 && arr[minQueue.peekTail()] >= arr[j] {
				minQueue.popTail()
			}
			minQueue.add(j)
			for maxQueue.len() != 0 && arr[maxQueue.peekTail()] <= arr[j] {
				maxQueue.popTail()
			}
			maxQueue.add(j)
			if arr[maxQueue.peekHead()]-arr[minQueue.peekHead()] > num {
				break
			}
			j++
		}
		if minQueue.peekHead() == i {
			minQueue.popHead()
		}
		if maxQueue.peekHead() == i {
			maxQueue.popHead()
		}
		res += j - i
		i++
	}
	return res
}
