package main

/**
 *  @ClassName:reversePrint
 *  @Description:剑指offer-06 从尾到头打印链表
 *  @Author:jackey
 *  @Create:2021/7/1 下午8:10
 */
/*
 *  @Description: 	三种解法：1、直接使用递归，从后往前遍历每一个结果；2、使用反转链表，在顺序读取结果；3、先顺序读取结果，将结果数组反转（本次实现）
 *  @Param:
 *  @Return:
 */


// 递归时间复杂度太高
func reversePrint1(head *ListNode) []int {
	if head == nil {
	return nil
}

	return appendData(head)
}

func appendData(head *ListNode) []int {
	if head.Next != nil{
	list := appendData(head.Next)
	list = append(list, head.Val)
	return list
}

	return []int{head.Val}
}

// 时间复杂度中等，反转链表耗费时间，比反转数组耗费的时间长
func reversePrint2(head *ListNode) []int {
	if head == nil {
		return nil
	}

	var newHead *ListNode
	res := []int{}
	for head != nil {
		node := head.Next
		head.Next = newHead
		newHead = head
		head = node
	}

	for newHead != nil {
		res = append(res, newHead.Val)
		newHead = newHead.Next
	}

	return res
}

// 最优解
func reversePrint3(head *ListNode) []int {
	if head == nil {
		return nil
	}
	r := head
	res := []int{}

	for r != nil {
		res = append(res, r.Val)
		r = r.Next
	}
	// 这个需要先排序再进行反转
	//sort.Reverse(sort.IntSlice(res))

	// 在go语言中for循环如果存在两个参数的变量，则不能只写i++，j--
	// 最简单实现的数组反转算法
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i],res[j] = res[j],res[i]
	}
	return res
}
