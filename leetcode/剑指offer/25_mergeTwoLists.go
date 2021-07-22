package main

/**
 *  @ClassName:25_mergeTwoLists
 *  @Description:剑指offer-25 合并两个排序的链表
 *  @Author:jackey
 *  @Create:2021/7/13 下午7:37
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	r := new(ListNode)

	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	p0 := r

	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			p0.Next = l2
			p0 = p0.Next
			l2 = l2.Next
		} else {
			p0.Next = l1
			p0 = p0.Next
			l1 = l1.Next
		}
	}

	if l1 == nil {
		p0.Next = l2
	} else {
		p0.Next = l1
	}

	return r.Next
}


