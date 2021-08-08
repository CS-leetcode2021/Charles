package heap

import (
	"container/heap"
	"fmt"
)

/**
 *  @ClassName:maxheap
 *  @Description:大根堆
 *  @Author:jackey
 *  @Create:2021/8/8 下午6:30
 */

// 五个接口

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

// test
func main() {
	h := make(MaxHeap,0)
	heap.Init(&h)

	heap.Push(&h,8)
	heap.Push(&h,1)
	heap.Push(&h,4)
	heap.Push(&h,5)
	heap.Push(&h,2)

	fmt.Println(h)
	fmt.Println(heap.Pop(&h))
	fmt.Println(heap.Pop(&h))
	fmt.Println(heap.Pop(&h))
	fmt.Println(heap.Pop(&h))
	fmt.Println(heap.Pop(&h))


}