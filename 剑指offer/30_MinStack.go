package main

import "math"

/**
 *  @ClassName:30_MinStack
 *  @Description:剑指offer-30 包含min函数的栈 同 leetcode-155
 *  @Author:jackey
 *  @Create:2021/7/15 下午6:49
 */

// 99/77
type MinStack struct {
	Val []int
	Mi []int
}

func Constructor() MinStack {
	val := make([]int,0)
	min := []int{math.MaxInt64}
	return MinStack{val,min}
}

func (this *MinStack) Push(x int)  {
	this.Val = append(this.Val,x)
	tmp := min(this.Mi[len(this.Mi)-1],x)
	this.Mi = append(this.Mi,tmp)
}

func (this *MinStack) Pop()  {
	if this.Size() == -1 {
		return
	}
	this.Val = this.Val[0:this.Size()-1]
	this.Mi = this.Mi[0:len(this.Mi)-1]
}
func (this *MinStack) Top() int {
	if this.Size() < 0 {
		return -1
	}
	return this.Val[this.Size()-1]
}

func (this *MinStack) Min() int {
	if len(this.Mi) == 1 {
		return 0
	}

	return this.Mi[len(this.Mi)-1]

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (this *MinStack) Size() int {
	if len(this.Val) == 0 {
		return -1
	}

	return len(this.Val)
}

