package main

import (
	"container/heap"
	"fmt"
)

type Seat []int

// 需要实现5个函数实现接口
func (s Seat) Len() int {
	return len(s)
}

func (s Seat) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s *Seat) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *Seat) Push(x interface{}) {
	*s = append(*s, x.(int))
}

func (s *Seat) Pop() interface{} {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func PickSeat(a, b *Seat, keep bool) int {
	var x int
	if a.Len() > 0 { // 还有单人桌子
		x = heap.Pop(a).(int) // 取出单人桌子
		if keep {
			heap.Push(b, x) //
		}
	} else { // 无单人桌子
		x = heap.Pop(b).(int) // 取出空桌子
		if !keep {            // 放入单人桌子
			heap.Push(a, x)
		}
	}
	return x
}

func main() {
	var T int
	fmt.Scan(&T)

	var N, M int
	var seats string
	var order string

	for i := 0; i < T; i++ {
		fmt.Scan(&N)
		fmt.Scan(&seats)

		one := &Seat{}
		zero := &Seat{}

		for j, t := range seats {
			if t == '1' {
				heap.Push(one, j+1) // 放入几号桌子
			} else if t == '0' {
				heap.Push(zero, j+1)
			}
		}

		fmt.Scan(&M)
		fmt.Scanln(&order)
		for _, t := range order {
			if t == 'M' {
				fmt.Println(PickSeat(one, zero, false))
			} else {
				fmt.Println(PickSeat(zero, one, true))
			}
		}
	}
}
