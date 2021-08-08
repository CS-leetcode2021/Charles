package heap

import (
	"fmt"
	"math"
)

/**
 *  @ClassName:minheap
 *  @Description:小根堆
 *  @Author:jackey
 *  @Create:2021/8/8 下午6:00
 */

type MinHeap struct {
	Element []int
}

func NewMinHeap() *MinHeap {
	// 数组中第一个元素不使用，存放一个小于堆中任何数字的值用于结束循环
	h := &MinHeap{Element: []int{math.MinInt64}}
	return h
}

// 上浮操作，是二叉堆有序，如果上浮一直到根，时间复杂度为O(logN)
func (H *MinHeap) Insert(v int) {
	H.Element = append(H.Element, v)
	i := len(H.Element) - 1
	// 上浮

	for ; H.Element[i/2] > v; i /= 2 {
		H.Element[i] = H.Element[i/2]
	}
	H.Element[i] = v
}

func (H *MinHeap) DeleteMin() (int, error) {
	if len(H.Element) <= 1 {
		return 0, fmt.Errorf("MinHeap is empty")
	}
	minElement := H.Element[1]

	lastElement := H.Element[len(H.Element)-1]

	var i, child int
	// 每次将最小值换上来
	for i = 1; i*2 < len(H.Element); i = child {
		child = i * 2
		if child < len(H.Element)-1 && H.Element[child+1] < H.Element[child] {
			child++
		}
		// 下降一层
		if lastElement > H.Element[child] {
			H.Element[i] = H.Element[child]
		} else {
			break
		}
	}
	H.Element[i] = lastElement
	H.Element = H.Element[:len(H.Element)-1]
	return minElement, nil

}

func (H *MinHeap) Size() int {
	return len(H.Element) - 1
}

func (H *MinHeap) Min() (int, error) {
	if H.Size() <= 1 {
		return 0, fmt.Errorf("MinHeap is empty")
	}

	return H.Element[1], nil
}

func (H *MinHeap) String() string {
	return fmt.Sprintf("%v", H.Element[1:])
}
